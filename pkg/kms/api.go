package kms

// KMS is the service interface for AWS KMS.
type KMS interface {
	GenerateRandom(*GenerateRandomRequest) (*GenerateRandomResult, error)

	ListGrants(*ListGrantsRequest) (*ListGrantsResult, error)
	ListRetireableGrants(*ListRetireableGrantsRequest) (*ListRetireableGrantsResult, error)
	CreateGrant(*CreateGrantRequest) (*CreateGrantResult, error)
	RetireGrant(*RetireGrantRequest) error
	RevokeGrant(*RevokeGrantRequest) error

	ListResourceTags(*ListResourceTagsRequest) (*ListResourceTagsResult, error)
	TagResource(*TagResourceRequest) error
	UntagResource(*UntagResourceRequest) error

	ListKeys(*ListKeysRequest) (*ListKeysResult, error)
	CreateKey(*CreateKeyRequest) (*CreateKeyResult, error)
	DescribeKey(*DescribeKeyRequest) (*DescribeKeyResult, error)
	UpdateKeyDescription(*UpdateKeyDescriptionRequest) error
	EnableKey(*EnableKeyRequest) error
	DisableKey(*DisableKeyRequest) error
	ScheduleKeyDeletion(*ScheduleKeyDeletionRequest) (*ScheduleKeyDeletionResult, error)
	CancelKeyDeletion(*CancelKeyDeletionRequest) (*CancelKeyDeletionResult, error)

	GetParametersForImport(*GetParametersForImportRequest) (*GetParametersForImportResult, error)
	ImportKeyMaterial(*ImportKeyMaterialRequest) error
	DeleteImportedKeyMaterial(*DeleteImportedKeyMaterialRequest) error

	GetKeyRotationStatus(*GetKeyRotationStatusRequest) (*GetKeyRotationStatusResult, error)
	EnableKeyRotation(*EnableKeyRotationRequest) error
	DisableKeyRotation(*DisableKeyRotationRequest) error

	ListKeyPolicies(*ListKeyPoliciesRequest) (*ListKeyPoliciesResult, error)
	GetKeyPolicy(*GetKeyPolicyRequest) (*GetKeyPolicyResult, error)
	PutKeyPolicy(*PutKeyPolicyRequest) error

	ListAliases(*ListAliasesRequest) (*ListAliasesResult, error)
	CreateAlias(*CreateAliasRequest) error
	UpdateAlias(*UpdateAliasRequest) error
	DeleteAlias(*DeleteAliasRequest) error

	GenerateDataKey(*GenerateDataKeyRequest) (*GenerateDataKeyResult, error)
	GenerateDataKeyWithoutPlaintext(*GenerateDataKeyRequest) (*GenerateDataKeyResult, error)
	Encrypt(*EncryptRequest) (*EncryptResult, error)
	Decrypt(*DecryptRequest) (*DecryptResult, error)
	ReEncrypt(*ReEncryptRequest) (*ReEncryptResult, error)
}

//
// API shapes for general requests.
//

// GenerateRandomRequest is a request to GenerateRandom.
type GenerateRandomRequest struct {
	NumberOfBytes int `json:"NumberOfBytes"`
}

// GenerateRandomResult is the result of GenerateRandom.
type GenerateRandomResult struct {
	Plaintext string `json:"Plaintext,omitempty"`
}

// GrantConstraint is a constraint on a grant.
type GrantConstraint struct {
	EncryptionContextEquals map[string]string `json:"EncryptionContextEquals,omitempty"`
	EncryptionContextSubset map[string]string `json:"EncryptionContextSubset,omitempty"`
}

// GrantListEntry is an entry in a list of grants.
type GrantListEntry struct {
	Constraints       *GrantConstraint `json:"Constraints,omitempty"`
	CreationDate      int64            `json:"CreationDate,omitempty"`
	GranteePrincipal  string           `json:"GranteePrincipal,omitempty"`
	GrantID           string           `json:"GrantId,omitempty"`
	IssuingAccount    string           `json:"IssuingAccount,omitempty"`
	KeyID             string           `json:"KeyId,omitempty"`
	Name              string           `json:"Name,omitempty"`
	Operations        []string         `json:"Operations,omitempty"`
	RetiringPrincipal string           `json:"RetiringPrincipal,omitempty"`
}

// ListGrantsRequest is a request to ListGrants.
type ListGrantsRequest struct {
	KeyID  string `json:"KeyId"`
	Limit  int    `json:"Limit"`
	Marker string `json:"Marker"`
}

// ListGrantsResult is the result of ListGrants.
type ListGrantsResult struct {
	Grants     []GrantListEntry `json:"Grants"`
	NextMarker string           `json:"NextMarker,omitempty"`
	Truncated  bool             `json:"Truncated,omitempty"`
}

// ListRetireableGrantsRequest is a request to ListRetireableGrants.
type ListRetireableGrantsRequest struct {
	Limit             int    `json:"Limit"`
	Marker            string `json:"Marker"`
	RetiringPrincipal string `json:"RetiringPrincipal"`
}

// ListRetireableGrantsResult is the result of ListRetireableGrants.
type ListRetireableGrantsResult struct {
	Grants     []GrantListEntry `json:"Grants"`
	NextMarker string           `json:"NextMarker,omitempty"`
	Truncated  bool             `json:"Truncated,omitempty"`
}

// CreateGrantRequest is a request to CreateGrant.
type CreateGrantRequest struct {
	Constraints       *GrantConstraint `json:"Constraints"`
	GranteePrincipal  string           `json:"GranteePrincipal"`
	GrantTokens       []string         `json:"GrantTokens"`
	KeyID             string           `json:"KeyId"`
	Name              string           `json:"Name"`
	Operations        []string         `json:"Operations"`
	RetiringPrincipal string           `json:"RetiringPrincipal"`
}

// CreateGrantResult is the result of CreateGrant.
type CreateGrantResult struct {
	GrantID    string `json:"GrantId"`
	GrantToken string `json:"GrantToken"`
}

// RetireGrantRequest is a request to RetireGrant.
type RetireGrantRequest struct {
	GrantID    string `json:"GrantId"`
	GrantToken string `json:"GrantToken"`
	KeyID      string `json:"KeyId"`
}

// RevokeGrantRequest is a request to RevokeGrant.
type RevokeGrantRequest struct {
	GrantID string `json:"GrantId"`
	KeyID   string `json:"KeyId"`
}

// Tag is a tag key/value pair.
type Tag struct {
	TagKey   string `json:"TagKey"`
	TagValue string `json:"TagValue"`
}

// ListResourceTagsRequest is a request to ListResourceTags.
type ListResourceTagsRequest struct {
	KeyID  string `json:"KeyId"`
	Limit  int    `json:"Limit"`
	Marker string `json:"Marker"`
}

// ListResourceTagsResult is the result of ListResourceTags.
type ListResourceTagsResult struct {
	Tags       []Tag  `json:"Tags"`
	NextMarker string `json:"NextMarker,omitempty"`
	Truncated  bool   `json:"Truncated,omitempty"`
}

// TagResourceRequest is a request to TagResource.
type TagResourceRequest struct {
	KeyID string `json:"KeyId"`
	Tags  []Tag  `json:"Tags"`
}

// UntagResourceRequest is a request to UntagResource.
type UntagResourceRequest struct {
	KeyID   string   `json:"KeyId"`
	TagKeys []string `json:"TagKeys"`
}

//
// API shapes for key metadata requests.
//

// ListKeysRequest is a request to ListKeys.
type ListKeysRequest struct {
	Limit  int    `json:"Limit"`
	Marker string `json:"Marker"`
}

// KeyListEntry is an entry in a list of keys.
type KeyListEntry struct {
	KeyArn string `json:"KeyArn,omitempty"`
	KeyID  string `json:"KeyId,omitempty"`
}

// ListKeysResult is the result of ListKeys.
type ListKeysResult struct {
	Keys       []KeyListEntry `json:"Keys"`
	NextMarker string         `json:"NextMarker,omitempty"`
	Truncated  bool           `json:"Truncated,omitempty"`
}

// CreateKeyRequest is a request to CreateKey.
type CreateKeyRequest struct {
	BypassPolicyLockoutSafetyCheck bool   `json:"BypassPolicyLockoutSafetyCheck"`
	Description                    string `json:"Description"`
	KeyUsage                       string `json:"KeyUsage"`
	Origin                         string `json:"Origin"`
	Policy                         string `json:"Policy"`
	Tags                           []Tag  `json:"Tags"`
}

// KeyMetadata is metadata about a key.
type KeyMetadata struct {
	Arn             string `json:"Arn"`
	AWSAccountID    string `json:"AWSAccountId"`
	CreationDate    int64  `json:"CreationDate"`
	DeletionDate    int64  `json:"DeletionDate,omitempty"`
	Description     string `json:"Description"`
	Enabled         bool   `json:"Enabled"`
	ExpirationModel string `json:"ExpirationModel,omitempty"`
	KeyID           string `json:"KeyId"`
	KeyManager      string `json:"KeyManager"`
	KeyState        string `json:"KeyState"`
	KeyUsage        string `json:"KeyUsage"`
	Origin          string `json:"Origin"`
	ValidTo         int64  `json:"ValidTo,omitempty"`
}

// CreateKeyResult is the result of CreateKey.
type CreateKeyResult struct {
	KeyMetadata *KeyMetadata `json:"KeyMetadata"`
}

// DescribeKeyRequest is a request to DescribeKey.
type DescribeKeyRequest struct {
	KeyID       string   `json:"KeyId"`
	GrantTokens []string `json:"GrantTokens"`
}

// DescribeKeyResult is the result of DescribeKey.
type DescribeKeyResult struct {
	KeyMetadata *KeyMetadata `json:"KeyMetadata"`
}

// UpdateKeyDescriptionRequest is a request to UpdateKeyDescription.
type UpdateKeyDescriptionRequest struct {
	Description string `json:"Description"`
	KeyID       string `json:"KeyId"`
}

// EnableKeyRequest is a request to EnableKey.
type EnableKeyRequest struct {
	KeyID string `json:"KeyId"`
}

// DisableKeyRequest is a request to DisableKey.
type DisableKeyRequest struct {
	KeyID string `json:"KeyId"`
}

// ScheduleKeyDeletionRequest is a request to ScheduleKeyDeletion.
type ScheduleKeyDeletionRequest struct {
	KeyID               string `json:"KeyId"`
	PendingWindowInDays int    `json:"PendingWindowInDays"`
}

// ScheduleKeyDeletionResult is the result of ScheduleKeyDeletion.
type ScheduleKeyDeletionResult struct {
	DeletionDate int64  `json:"DeletionDate"`
	KeyID        string `json:"KeyId"`
}

// CancelKeyDeletionRequest is a request to CancelKeyDeletion.
type CancelKeyDeletionRequest struct {
	KeyID string `json:"KeyId"`
}

// CancelKeyDeletionResult is the result of CancelKeyDeletion.
type CancelKeyDeletionResult struct {
	KeyID string `json:"KeyId"`
}

// GetParametersForImportRequest is a request to GetParametersForImport.
type GetParametersForImportRequest struct {
	KeyID             string `json:"KeyId"`
	WrappingAlgorithm string `json:"WrappingAlgorithm"`
	WrappingKeySpec   string `json:"WrappingKeySpec"`
}

// GetParametersForImportResult is the result of GetParametersForImport.
type GetParametersForImportResult struct {
	ImportToken       string `json:"ImportToken"`
	KeyID             string `json:"KeyID"`
	ParametersValidTo int64  `json:"ParametersValidTo"`
	PublicKey         string `json:"PublicKey"`
}

// ImportKeyMaterialRequest is a request to ImportKeyMaterial.
type ImportKeyMaterialRequest struct {
	EncryptedKeyMaterial string `json:"EncryptedKeyMaterial"`
	ExpirationModel      string `json:"ExpirationModel"`
	ImportToken          string `json:"ImportToken"`
	KeyID                string `json:"KeyId"`
	ValidTo              int64  `json:"ValidTo"`
}

// DeleteImportedKeyMaterialRequest is a request to DeleteImportedKeyMaterial.
type DeleteImportedKeyMaterialRequest struct {
	KeyID string `json:"KeyId"`
}

// GetKeyRotationStatusRequest is a request to GetKeyRotationStatus.
type GetKeyRotationStatusRequest struct {
	KeyID string `json:"KeyId"`
}

// GetKeyRotationStatusResult is the result of GetKeyRotationStatus.
type GetKeyRotationStatusResult struct {
	KeyRotationEnabled bool `json:"KeyRotationEnabled"`
}

// EnableKeyRotationRequest is a request to EnableKeyRotation.
type EnableKeyRotationRequest struct {
	KeyID string `json:"KeyId"`
}

// DisableKeyRotationRequest is a request to DisableKeyRotation.
type DisableKeyRotationRequest struct {
	KeyID string `json:"KeyId"`
}

// ListKeyPoliciesRequest is a request to ListKeyPolicies.
type ListKeyPoliciesRequest struct {
	KeyID  string `json:"KeyId"`
	Limit  int    `json:"Limit"`
	Marker string `json:"Marker"`
}

// ListKeyPoliciesResult is the result of ListKeyPolicies.
type ListKeyPoliciesResult struct {
	PolicyNames []string `json:"PolicyNames"`
	NextMarker  string   `json:"NextMarker,omitempty"`
	Truncated   bool     `json:"Truncated,omitempty"`
}

// GetKeyPolicyRequest is a request to GetKeyPolicy.
type GetKeyPolicyRequest struct {
	KeyID      string `json:"KeyId"`
	PolicyName string `json:"PolicyName"`
}

// GetKeyPolicyResult is the result of GetKeyPolicy.
type GetKeyPolicyResult struct {
	Policy string `json:"Policy"`
}

// PutKeyPolicyRequest is a request to PutKeyPolicy.
type PutKeyPolicyRequest struct {
	BypassPolicyLockoutSafetyCheck bool   `json:"BypassPolicyLockoutSafetyCheck"`
	KeyID                          string `json:"KeyId"`
	Policy                         string `json:"Policy"`
	PolicyName                     string `json:"PolicyName"`
}

//
// API shapes for dealing with aliases.
//

// ListAliasesRequest is a request to ListAliases.
type ListAliasesRequest struct {
	Limit  int    `json:"Limit"`
	Marker string `json:"Marker"`
}

// AliasListEntry is an entry in an alias list.
type AliasListEntry struct {
	AliasArn    string `json:"AliasArn"`
	AliasName   string `json:"AliasName"`
	TargetKeyID string `json:"TargetKeyId"`
}

// ListAliasesResult is the result of ListAliases.
type ListAliasesResult struct {
	Aliases    []AliasListEntry `json:"Aliases"`
	NextMarker string           `json:"NextMarker,omitempty"`
	Truncated  bool             `json:"Truncated,omitempty"`
}

// CreateAliasRequest is a request to CreateAlias.
type CreateAliasRequest struct {
	AliasName   string `json:"AliasName"`
	TargetKeyID string `json:"TargetKeyId"`
}

// UpdateAliasRequest is a request to UpdateAlias.
type UpdateAliasRequest struct {
	AliasName   string `json:"AliasName"`
	TargetKeyID string `json:"TargetKeyId"`
}

// DeleteAliasRequest is a request to DeleteAlias.
type DeleteAliasRequest struct {
	AliasName string `json:"AliasName"`
}

//
// API shapes for crypto operations.
//

// GenerateDataKeyRequest is a request to GenerateDataKey.
type GenerateDataKeyRequest struct {
	EncryptionContext map[string]string `json:"EncryptionContext"`
	GrantTokens       []string          `json:"GrantTokens"`
	KeyID             string            `json:"KeyId"`
	KeySpec           string            `json:"KeySpec"`
	NumberOfBytes     int               `json:"NumberOfBytes"`
}

// GenerateDataKeyResult is the result of GenerateDataKey.
type GenerateDataKeyResult struct {
	CiphertextBlob string `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
	Plaintext      string `json:"Plaintext,omitempty"`
}

// EncryptRequest is a request to Encrypt.
type EncryptRequest struct {
	EncryptionContext map[string]string `json:"EncryptionContext"`
	GrantTokens       []string          `json:"GrantTokens"`
	KeyID             string            `json:"KeyId"`
	Plaintext         string            `json:"Plaintext"`
}

// EncryptResult is the result of Encrypt.
type EncryptResult struct {
	CiphertextBlob string `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
}

// DecryptRequest is a request to Decrypt.
type DecryptRequest struct {
	CiphertextBlob    string            `json:"CiphertextBlob"`
	EncryptionContext map[string]string `json:"EncryptionContext"`
	GrantTokens       []string          `json:"GrantTokens"`
}

// DecryptResult is the result of Decrypt.
type DecryptResult struct {
	KeyID     string `json:"KeyId,omitempty"`
	Plaintext string `json:"Plaintext,omitempty"`
}

// ReEncryptRequest is a request to ReEncrypt.
type ReEncryptRequest struct {
	CiphertextBlob               string            `json:"CiphertextBlob"`
	DestinationEncryptionContext map[string]string `json:"DestinationEncryptionContext"`
	DestinationKeyID             string            `json:"DestinationKeyId"`
	GrantTokens                  []string          `json:"GrantTokens"`
	SourceEncryptionContext      map[string]string `json:"SourceEncryptionContext"`
}

// ReEncryptResult is the result of ReEncrypt.
type ReEncryptResult struct {
	CiphertextBlob string `json:"CiphertextBlob,omitempty"`
	KeyID          string `json:"KeyId,omitempty"`
	SourceKeyID    string `json:"SourceKeyId,omitempty"`
}
