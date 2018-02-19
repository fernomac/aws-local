package kms

import "errors"

func (k *kms) GetParametersForImport(req *GetParametersForImportRequest) (*GetParametersForImportResult, error) {
	return nil, errors.New("Unimplemented")
}

func (k *kms) ImportKeyMaterial(req *ImportKeyMaterialRequest) error {
	return errors.New("Unimplemented")
}

func (k *kms) DeleteImportedKeyMaterial(req *DeleteImportedKeyMaterialRequest) error {
	return errors.New("Unimplemented")
}
