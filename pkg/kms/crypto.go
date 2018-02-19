package kms

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sort"
)

func makeAad(ctx map[string]string) []byte {
	buf := bytes.Buffer{}

	keys := make([]string, len(ctx))
	for key := range ctx {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		buf.WriteString(key)
		buf.WriteString(ctx[key])
	}

	return buf.Bytes()
}

func writeBytes(buf *bytes.Buffer, str []byte) error {
	l := len(str)
	if l > 0xFFFF {
		return errors.New("InternalFailure")
	}

	buf.WriteByte(byte(l & 0xFF))
	buf.WriteByte(byte((l >> 8) & 0xFF))
	buf.Write(str)

	return nil
}

func makeCiphertextBlob(keyID string, nonce []byte, ciphertext []byte) ([]byte, error) {
	buf := bytes.Buffer{}

	if err := buf.WriteByte( /*version=*/ 0); err != nil {
		return nil, err
	}
	if err := writeBytes(&buf, []byte(keyID)); err != nil {
		return nil, err
	}
	if err := writeBytes(&buf, nonce); err != nil {
		return nil, err
	}
	if err := writeBytes(&buf, ciphertext); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func readLen(buf *bytes.Reader) (int, error) {
	low, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}

	high, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}

	return (int(high) << 8) | int(low), nil
}

func readBytes(buf *bytes.Reader) ([]byte, error) {
	len, err := readLen(buf)
	if err != nil {
		return nil, err
	}

	str := make([]byte, len)
	_, err = buf.Read(str)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func readString(buf *bytes.Reader) (string, error) {
	str, err := readBytes(buf)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func parseCiphertextBlob(ciphertextBlob []byte) (string, []byte, []byte, error) {
	buf := bytes.NewReader(ciphertextBlob)

	ver, err := buf.ReadByte()
	if err != nil {
		return "", nil, nil, err
	}
	if ver != 0 {
		return "", nil, nil, errors.New("InvalidCiphertextException")
	}

	keyID, err := readString(buf)
	if err != nil {
		return "", nil, nil, err
	}

	nonce, err := readBytes(buf)
	if err != nil {
		return "", nil, nil, err
	}

	ciphertext, err := readBytes(buf)
	if err != nil {
		return "", nil, nil, err
	}

	if buf.Len() != 0 {
		return "", nil, nil, errors.New("InvalidCiphertextException")
	}

	return keyID, nonce, ciphertext, nil
}

func encrypt(plaintext []byte, ctx map[string]string, key *key) ([]byte, error) {
	block, err := aes.NewCipher(key.key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	aad := makeAad(ctx)
	ciphertext := aead.Seal(nil, nonce, plaintext, aad)
	return makeCiphertextBlob(key.meta.KeyID, nonce, ciphertext)
}

func (k *kms) doGDK(req *GenerateDataKeyRequest, withPlaintext bool) (*GenerateDataKeyResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.GrantTokens != nil {
		return nil, errors.New("GrantsNotSupported")
	}

	len := req.NumberOfBytes
	if len == 0 {
		if req.KeySpec == "AES_128" {
			len = 16
		} else if req.KeySpec == "AES_256" {
			len = 32
		} else {
			return nil, errors.New("InvalidKeyUsageException")
		}
	} else if req.KeySpec != "" {
		return nil, errors.New("InvalidParameterCombination")
	}

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}
	if !key.meta.Enabled {
		return nil, errors.New("DisabledException")
	}

	plaintext := make([]byte, len)
	_, err := rand.Read(plaintext)
	if err != nil {
		return nil, err
	}

	ciphertext, err := encrypt(plaintext, req.EncryptionContext, key)
	if err != nil {
		return nil, err
	}

	encodedPlaintext := ""
	if withPlaintext {
		encodedPlaintext = base64.StdEncoding.EncodeToString(plaintext)
	}

	return &GenerateDataKeyResult{
		KeyID:          key.meta.KeyID,
		Plaintext:      encodedPlaintext,
		CiphertextBlob: base64.StdEncoding.EncodeToString(ciphertext),
	}, nil
}

func (k *kms) GenerateDataKey(req *GenerateDataKeyRequest) (*GenerateDataKeyResult, error) {
	return k.doGDK(req, true)
}

func (k *kms) GenerateDataKeyWithoutPlaintext(req *GenerateDataKeyRequest) (*GenerateDataKeyResult, error) {
	return k.doGDK(req, false)
}

func (k *kms) Encrypt(req *EncryptRequest) (*EncryptResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.GrantTokens != nil {
		return nil, errors.New("GrantsNotSupported")
	}

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}
	if !key.meta.Enabled {
		return nil, errors.New("DisabledException")
	}

	plaintext, err := base64.StdEncoding.DecodeString(req.Plaintext)
	if err != nil {
		return nil, err
	}

	ciphertext, err := encrypt(plaintext, req.EncryptionContext, key)
	if err != nil {
		return nil, err
	}

	return &EncryptResult{
		KeyID:          key.meta.KeyID,
		CiphertextBlob: base64.StdEncoding.EncodeToString(ciphertext),
	}, nil
}

func (k *kms) decrypt(ciphertextBlob []byte, ctx map[string]string) (string, []byte, error) {
	keyID, nonce, ciphertext, err := parseCiphertextBlob(ciphertextBlob)
	if err != nil {
		return "", nil, err
	}

	key := k.get(keyID)
	if key == nil {
		return "", nil, errors.New("NotFoundException")
	}
	if !key.meta.Enabled {
		return "", nil, errors.New("DisabledException")
	}

	block, err := aes.NewCipher(key.key)
	if err != nil {
		return "", nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, err
	}

	aad := makeAad(ctx)
	plaintext, err := aead.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		return "", nil, err
	}

	return keyID, plaintext, nil
}

func (k *kms) Decrypt(req *DecryptRequest) (*DecryptResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.GrantTokens != nil {
		return nil, errors.New("GrantsNotSupported")
	}

	ciphertextBlob, err := base64.StdEncoding.DecodeString(req.CiphertextBlob)
	if err != nil {
		return nil, err
	}

	keyID, plaintext, err := k.decrypt(ciphertextBlob, req.EncryptionContext)
	if err != nil {
		return nil, err
	}

	return &DecryptResult{
		KeyID:     keyID,
		Plaintext: base64.StdEncoding.EncodeToString(plaintext),
	}, nil
}

func (k *kms) ReEncrypt(req *ReEncryptRequest) (*ReEncryptResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.GrantTokens != nil {
		return nil, errors.New("GrantsNotSupported")
	}

	ciphertextBlob, err := base64.StdEncoding.DecodeString(req.CiphertextBlob)
	if err != nil {
		return nil, err
	}

	sourceKeyID, plaintext, err := k.decrypt(ciphertextBlob, req.SourceEncryptionContext)
	if err != nil {
		return nil, err
	}

	key := k.get(req.DestinationKeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}
	if !key.meta.Enabled {
		return nil, errors.New("DisabledException")
	}

	ciphertext, err := encrypt(plaintext, req.DestinationEncryptionContext, key)
	if err != nil {
		return nil, err
	}

	return &ReEncryptResult{
		KeyID:          key.meta.KeyID,
		SourceKeyID:    sourceKeyID,
		CiphertextBlob: base64.StdEncoding.EncodeToString(ciphertext),
	}, nil
}
