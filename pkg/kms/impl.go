package kms

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
	"sync"
)

type key struct {
	key      []byte
	meta     *KeyMetadata
	tags     map[string]string
	policies map[string]string
}

type kms struct {
	lock    sync.Mutex
	counter int64
	keys    map[string]*key
	arns    map[string]*key
	aliases map[string]*key
	grants  map[string]*GrantListEntry
}

// New creates a new KMS object.
func New() KMS {
	return &kms{
		counter: 0,
		keys:    make(map[string]*key),
		arns:    make(map[string]*key),
		aliases: make(map[string]*key),
		grants:  make(map[string]*GrantListEntry),
	}
}

func (k *kms) get(keyID string) *key {
	if strings.HasPrefix(keyID, "alias/") {
		return k.aliases[keyID]
	}
	if strings.HasPrefix(keyID, "arn:") {
		return k.arns[keyID]
	}
	return k.keys[keyID]
}

func (k *kms) GenerateRandom(req *GenerateRandomRequest) (*GenerateRandomResult, error) {
	if req.NumberOfBytes < 1 || req.NumberOfBytes > 1024 {
		return nil, errors.New("InvalidParameterValue")
	}

	out := make([]byte, req.NumberOfBytes)
	if _, err := rand.Read(out); err != nil {
		return nil, err
	}

	return &GenerateRandomResult{
		Plaintext: base64.StdEncoding.EncodeToString(out),
	}, nil
}
