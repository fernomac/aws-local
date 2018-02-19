package kms

import "errors"

func (k *kms) GetKeyRotationStatus(req *GetKeyRotationStatusRequest) (*GetKeyRotationStatusResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}

	return &GetKeyRotationStatusResult{
		KeyRotationEnabled: false,
	}, nil
}

func (k *kms) EnableKeyRotation(req *EnableKeyRotationRequest) error {
	return errors.New("Unimplemented")
}

func (k *kms) DisableKeyRotation(req *DisableKeyRotationRequest) error {
	return errors.New("Unimplemented")
}
