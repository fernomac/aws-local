package kms

import (
	"errors"
	"fmt"
	"strings"
)

func (k *kms) ListAliases(req *ListAliasesRequest) (*ListAliasesResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.Marker != "" {
		return nil, errors.New("InvalidMarkerException")
	}
	if req.Limit != 0 {
		return nil, errors.New("LimitNotSupported")
	}

	aliases := []AliasListEntry{}
	for alias, key := range k.aliases {
		aliases = append(aliases, AliasListEntry{
			AliasArn:    fmt.Sprintf("arn:aws:kms:us-local-1:-:%v", alias),
			AliasName:   alias,
			TargetKeyID: key.meta.KeyID,
		})
	}

	return &ListAliasesResult{
		Aliases:   aliases,
		Truncated: false,
	}, nil
}

func (k *kms) CreateAlias(req *CreateAliasRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	if !strings.HasPrefix(req.AliasName, "alias/") {
		return errors.New("InvalidAliasNameException")
	}
	if _, ok := k.aliases[req.AliasName]; ok {
		return errors.New("AlreadyExistsException")
	}

	key := k.get(req.TargetKeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	k.aliases[req.AliasName] = key
	return nil
}

func (k *kms) UpdateAlias(req *UpdateAliasRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	if _, ok := k.aliases[req.AliasName]; !ok {
		return errors.New("NotFoundException")
	}

	key := k.get(req.TargetKeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	k.aliases[req.AliasName] = key
	return nil
}

func (k *kms) DeleteAlias(req *DeleteAliasRequest) error {
	defer k.lock.Unlock()

	if _, ok := k.aliases[req.AliasName]; !ok {
		return errors.New("NotFoundException")
	}

	delete(k.aliases, req.AliasName)
	return nil
}
