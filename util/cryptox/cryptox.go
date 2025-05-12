package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Crypto struct {
	key string
}

func New(k string) (ICrypto, error) {
	l := len(k)
	if l == 0 {
		return nil, errors.New("key is nil")
	}
	return Crypto{key: k}, nil
}

func (c Crypto) Encrypt(s string) (string, error) {
	cr, err := aes.NewCipher([]byte(c.key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(cr)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err

	}
	return base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(s), nil)), nil
}

func (c Crypto) Decrypt(s string) (string, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(s)
	cr, err := aes.NewCipher([]byte(c.key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(cr)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return "", err
	}

	return bytes.NewBuffer(t).String(), nil
}

func (c Crypto) Md5Sum(s string) (string, error) {

	hash := md5.New()
	if _, err := io.Copy(hash, strings.NewReader(s)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (c Crypto) Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
