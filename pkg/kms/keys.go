package kms

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"
)

func (k *kms) ListKeys(req *ListKeysRequest) (*ListKeysResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.Marker != "" {
		return nil, errors.New("InvalidMarkerException")
	}
	if req.Limit != 0 {
		return nil, errors.New("LimitNotSupported")
	}

	keys := []KeyListEntry{}
	for _, key := range k.keys {
		keys = append(keys, KeyListEntry{
			KeyArn: key.meta.Arn,
			KeyID:  key.meta.KeyID,
		})
	}

	return &ListKeysResult{
		Keys:      keys,
		Truncated: false,
	}, nil
}

func (k *kms) CreateKey(req *CreateKeyRequest) (*CreateKeyResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	keyUsage := req.KeyUsage
	if keyUsage == "" {
		keyUsage = "ENCRYPT_DECRYPT"
	}
	if keyUsage != "ENCRYPT_DECRYPT" {
		return nil, errors.New("InvalidParameterValue")
	}

	origin := req.Origin
	if origin == "" {
		origin = "AWS_KMS"
	}
	if origin != "AWS_KMS" && origin != "EXTERNAL" {
		return nil, errors.New("InvalidParameterValue")
	}

	if req.Policy != "" {
		return nil, errors.New("PolicyNotSupported")
	}

	// Generate a key.
	var raw []byte
	var state string

	if origin == "AWS_KMS" {
		raw = make([]byte, 32)
		_, err := rand.Read(raw)
		if err != nil {
			return nil, err
		}
		state = "Enabled"
	} else {
		state = "PendingImport"
	}

	id := fmt.Sprintf("%v", k.counter)
	k.counter++

	tags := map[string]string{}
	for _, tag := range req.Tags {
		tags[tag.TagKey] = tag.TagValue
	}

	key := &key{
		meta: &KeyMetadata{
			AWSAccountID: "x",
			Arn:          fmt.Sprintf("arn:aws:kms:us-local-1:x:key/%v", id),
			CreationDate: time.Now().Unix(),
			Description:  req.Description,
			Enabled:      true,
			KeyID:        id,
			KeyManager:   "CUSTOMER",
			KeyState:     state,
			KeyUsage:     keyUsage,
			Origin:       origin,
		},
		key:  raw,
		tags: tags,
	}

	k.keys[key.meta.KeyID] = key
	k.arns[key.meta.Arn] = key

	return &CreateKeyResult{key.meta}, nil
}

func (k *kms) DescribeKey(req *DescribeKeyRequest) (*DescribeKeyResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.GrantTokens != nil {
		return nil, errors.New("GrantsNotSupported")
	}

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}

	return &DescribeKeyResult{
		KeyMetadata: key.meta,
	}, nil
}

func (k *kms) UpdateKeyDescription(req *UpdateKeyDescriptionRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	key.meta.Description = req.Description
	return nil
}

func (k *kms) EnableKey(req *EnableKeyRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	if key.meta.KeyState == "PendingDeletion" || key.meta.KeyState == "PendingImport" {
		return errors.New("KMSInvalidStateException")
	}

	key.meta.Enabled = true
	key.meta.KeyState = "Enabled"
	return nil
}

func (k *kms) DisableKey(req *DisableKeyRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	if key.meta.KeyState == "PendingDeletion" || key.meta.KeyState == "PendingImport" {
		return errors.New("KMSInvalidStateException")
	}

	key.meta.Enabled = false
	key.meta.KeyState = "Disabled"
	return nil
}

func (k *kms) ScheduleKeyDeletion(req *ScheduleKeyDeletionRequest) (*ScheduleKeyDeletionResult, error) {
	return nil, errors.New("Unimplemented")
}

func (k *kms) CancelKeyDeletion(req *CancelKeyDeletionRequest) (*CancelKeyDeletionResult, error) {
	return nil, errors.New("Unimplemented")
}
