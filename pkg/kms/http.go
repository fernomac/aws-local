package kms

import (
	"encoding/json"
	"net/http"

	"github.com/fernomac/aws-local/pkg/awsjson11"
)

// NewHandler creates a new HTTP handler.
func NewHandler(kms KMS) http.Handler {
	rval := awsjson11.NewHandler("TrentService")

	rval.HandleWith("GenerateRandom", func(body []byte) (interface{}, error) {
		req := GenerateRandomRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.GenerateRandom(&req)
	})

	//
	// Grants.
	//

	rval.HandleWith("ListGrants", func(body []byte) (interface{}, error) {
		req := ListGrantsRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ListGrants(&req)
	})

	rval.HandleWith("ListRetireableGrants", func(body []byte) (interface{}, error) {
		req := ListRetireableGrantsRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ListRetireableGrants(&req)
	})

	rval.HandleWith("CreateGrant", func(body []byte) (interface{}, error) {
		req := CreateGrantRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.CreateGrant(&req)
	})

	rval.HandleWith("RetireGrant", func(body []byte) (interface{}, error) {
		req := RetireGrantRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.RetireGrant(&req)
	})

	rval.HandleWith("RevokeGrant", func(body []byte) (interface{}, error) {
		req := RevokeGrantRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.RevokeGrant(&req)
	})

	//
	// Tags.
	//

	rval.HandleWith("ListResourceTags", func(body []byte) (interface{}, error) {
		req := ListResourceTagsRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ListResourceTags(&req)
	})

	rval.HandleWith("TagResource", func(body []byte) (interface{}, error) {
		req := TagResourceRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.TagResource(&req)
	})

	rval.HandleWith("UntagResource", func(body []byte) (interface{}, error) {
		req := UntagResourceRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.UntagResource(&req)
	})

	//
	// Keys.
	//

	rval.HandleWith("ListKeys", func(body []byte) (interface{}, error) {
		req := ListKeysRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ListKeys(&req)
	})

	rval.HandleWith("CreateKey", func(body []byte) (interface{}, error) {
		req := CreateKeyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.CreateKey(&req)
	})

	rval.HandleWith("DescribeKey", func(body []byte) (interface{}, error) {
		req := DescribeKeyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.DescribeKey(&req)
	})

	rval.HandleWith("UpdateKeyDescription", func(body []byte) (interface{}, error) {
		req := UpdateKeyDescriptionRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.UpdateKeyDescription(&req)
	})

	rval.HandleWith("EnableKey", func(body []byte) (interface{}, error) {
		req := EnableKeyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.EnableKey(&req)
	})

	rval.HandleWith("DisableKey", func(body []byte) (interface{}, error) {
		req := DisableKeyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.DisableKey(&req)
	})

	rval.HandleWith("ScheduleKeyDeletion", func(body []byte) (interface{}, error) {
		req := ScheduleKeyDeletionRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ScheduleKeyDeletion(&req)
	})

	rval.HandleWith("CancelKeyDeletion", func(body []byte) (interface{}, error) {
		req := CancelKeyDeletionRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.CancelKeyDeletion(&req)
	})

	//
	// Import.
	//

	rval.HandleWith("GetParametersForImport", func(body []byte) (interface{}, error) {
		req := GetParametersForImportRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.GetParametersForImport(&req)
	})

	rval.HandleWith("ImportKeyMaterial", func(body []byte) (interface{}, error) {
		req := ImportKeyMaterialRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.ImportKeyMaterial(&req)
	})

	rval.HandleWith("DeleteImportedKeyMaterial", func(body []byte) (interface{}, error) {
		req := DeleteImportedKeyMaterialRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.DeleteImportedKeyMaterial(&req)
	})

	//
	// Rotation.
	//

	rval.HandleWith("GetKeyRotationStatus", func(body []byte) (interface{}, error) {
		req := GetKeyRotationStatusRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.GetKeyRotationStatus(&req)
	})

	rval.HandleWith("EnableKeyRotation", func(body []byte) (interface{}, error) {
		req := EnableKeyRotationRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.EnableKeyRotation(&req)
	})

	rval.HandleWith("DisableKeyRotation", func(body []byte) (interface{}, error) {
		req := DisableKeyRotationRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.DisableKeyRotation(&req)
	})

	//
	// Policies.
	//

	rval.HandleWith("ListKeyPolicies", func(body []byte) (interface{}, error) {
		req := ListKeyPoliciesRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ListKeyPolicies(&req)
	})

	rval.HandleWith("GetKeyPolicy", func(body []byte) (interface{}, error) {
		req := GetKeyPolicyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.GetKeyPolicy(&req)
	})

	rval.HandleWith("PutKeyPolicy", func(body []byte) (interface{}, error) {
		req := PutKeyPolicyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.PutKeyPolicy(&req)
	})

	//
	// Aliases.
	//

	rval.HandleWith("ListAliases", func(body []byte) (interface{}, error) {
		req := ListAliasesRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ListAliases(&req)
	})

	rval.HandleWith("CreateAlias", func(body []byte) (interface{}, error) {
		req := CreateAliasRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.CreateAlias(&req)
	})

	rval.HandleWith("UpdateAlias", func(body []byte) (interface{}, error) {
		req := UpdateAliasRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.UpdateAlias(&req)
	})

	rval.HandleWith("DeleteAlias", func(body []byte) (interface{}, error) {
		req := DeleteAliasRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return nil, kms.DeleteAlias(&req)
	})

	//
	// Crypto.
	//

	rval.HandleWith("GenerateDataKey", func(body []byte) (interface{}, error) {
		req := GenerateDataKeyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.GenerateDataKey(&req)
	})

	rval.HandleWith("GenerateDataKeyWithoutPlaintext", func(body []byte) (interface{}, error) {
		req := GenerateDataKeyRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.GenerateDataKeyWithoutPlaintext(&req)
	})

	rval.HandleWith("Encrypt", func(body []byte) (interface{}, error) {
		req := EncryptRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.Encrypt(&req)
	})

	rval.HandleWith("Decrypt", func(body []byte) (interface{}, error) {
		req := DecryptRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.Decrypt(&req)
	})

	rval.HandleWith("ReEncrypt", func(body []byte) (interface{}, error) {
		req := ReEncryptRequest{}
		if err := json.Unmarshal(body, &req); err != nil {
			return nil, err
		}
		return kms.ReEncrypt(&req)
	})

	return rval
}
