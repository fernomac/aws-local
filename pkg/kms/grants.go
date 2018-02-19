package kms

import "errors"

func (k *kms) ListGrants(req *ListGrantsRequest) (*ListGrantsResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	grants := []GrantListEntry{}
	for _, grant := range k.grants {
		if grant.KeyID == req.KeyID {
			grants = append(grants, *grant)
		}
	}

	return &ListGrantsResult{
		Grants:    grants,
		Truncated: false,
	}, nil
}

func (k *kms) ListRetireableGrants(req *ListRetireableGrantsRequest) (*ListRetireableGrantsResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	grants := []GrantListEntry{}
	for _, grant := range k.grants {
		if grant.RetiringPrincipal == req.RetiringPrincipal {
			grants = append(grants, *grant)
		}
	}

	return &ListRetireableGrantsResult{
		Grants:    grants,
		Truncated: false,
	}, nil
}

func (k *kms) CreateGrant(req *CreateGrantRequest) (*CreateGrantResult, error) {
	// TODO: Implement me?
	return nil, errors.New("GrantsNotSupported")
}

func (k *kms) findGrantToken(keyID string, grantID string) string {
	for token, grant := range k.grants {
		if grant.KeyID == keyID && grant.GrantID == grantID {
			return token
		}
	}
	return ""
}

func (k *kms) RetireGrant(req *RetireGrantRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.GrantToken != "" {
		delete(k.grants, req.GrantToken)
	} else {
		token := k.findGrantToken(req.KeyID, req.GrantID)
		if token == "" {
			return errors.New("NotFoundException")
		}
		delete(k.grants, token)
	}

	return nil
}

func (k *kms) RevokeGrant(req *RevokeGrantRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	token := k.findGrantToken(req.KeyID, req.GrantID)
	if token == "" {
		return errors.New("NotFoundException")
	}
	delete(k.grants, token)

	return nil
}
