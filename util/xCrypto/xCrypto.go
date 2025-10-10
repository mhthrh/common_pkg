package xCrypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"golang.org/x/tools/go/loader"
	"gopkg.in/square/go-jose.v2"
)

type Crypto struct {
	cnfg *loader.Config
}

func New(c *loader.Config) *Crypto {
	return &Crypto{c}
}

func (c *Crypto) CreateJWE(payload string, keyId string, mleServerPublicCertificatePath string) string {
	publicKey := loadPublicKey(mleServerPublicCertificatePath)
	opts := new(jose.EncrypterOptions)

	iat := currentMillis()

	opts.WithHeader("kid", keyId)
	opts.WithHeader("iat", iat)
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: publicKey}, opts)
	if err != nil {
		panic(err)
	}

	object, err := encrypter.Encrypt([]byte(payload))
	if err != nil {
		panic(err)
	}

	serialized, err := object.CompactSerialize()
	if err != nil {
		panic(err)
	}
	return serialized
}
func (c *Crypto) DecryptJWE(encryptedPayload string, mleClientPrivateKeyPath string) (string, error) {

	encryptedData := parseEncryptedResponse(encryptedPayload)
	if strings.Trim(encryptedData.EncData, " ") == "" {
		return "", errors.New(fmt.Sprintf("encrypted payload is empty"))

	}
	object, err := jose.ParseEncrypted(encryptedData.EncData)
	if err != nil {
		return "", err
	}

	privateKey := loadPrivateKey(mleClientPrivateKeyPath)
	decrypted, err := object.Decrypt(privateKey)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func loadPrivateKey(keyFilePath string) *rsa.PrivateKey {
	keyPem, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(keyPem)

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return priv
}

type EncryptedResponse struct {
	EncData string
}

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func loadPublicKey(certFilePath string) *rsa.PublicKey {
	certificate, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(certificate)
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	return cert.PublicKey.(*rsa.PublicKey)
}

func parseEncryptedResponse(encryptedPayload string) EncryptedResponse {
	if encryptedPayload == "" {
		return EncryptedResponse{}
	}
	var encryptedResponse EncryptedResponse
	err := json.Unmarshal([]byte(encryptedPayload), &encryptedResponse)

	if err != nil {
		return EncryptedResponse{encryptedPayload}
	}
	return encryptedResponse
}
