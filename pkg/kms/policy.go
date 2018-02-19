package kms

import (
	"encoding/json"
	"errors"
)

type statement struct {
	Sid          string      `json:"Sid"`
	Effect       string      `json:"Effect"`
	Principal    interface{} `json:"Principal"`
	NotPrincipal interface{} `json:"NotPrincipal"`
	Action       interface{} `json:"Action"`
	NotAction    interface{} `json:"NotAction"`
	Resource     interface{} `json:"Resource"`
	NotResource  interface{} `json:"NotResource"`
	Condition    interface{} `json:"Condition"`
}

type policy struct {
	Version   string       `json:"Version"`
	Statement []*statement `json:"Statement"`
}

func parsePolicy(str string) (*policy, error) {
	out := &policy{}
	if err := json.Unmarshal([]byte(str), out); err != nil {
		return nil, err
	}

	// TODO: Otherwise sanity check the policy?

	return out, nil
}

func (k *kms) ListKeyPolicies(req *ListKeyPoliciesRequest) (*ListKeyPoliciesResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.Marker != "" {
		return nil, errors.New("InvalidMarkerException")
	}
	if req.Limit != 0 {
		return nil, errors.New("LimitNotSupported")
	}

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}

	names := []string{}
	for name := range key.policies {
		names = append(names, name)
	}

	return &ListKeyPoliciesResult{
		PolicyNames: names,
		Truncated:   false,
	}, nil
}

func (k *kms) GetKeyPolicy(req *GetKeyPolicyRequest) (*GetKeyPolicyResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}

	policy, ok := key.policies[req.PolicyName]
	if !ok {
		return nil, errors.New("NotFoundException")
	}

	return &GetKeyPolicyResult{
		Policy: policy,
	}, nil
}

func (k *kms) PutKeyPolicy(req *PutKeyPolicyRequest) error {
	return errors.New("Unsupported")
}
