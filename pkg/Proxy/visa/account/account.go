package account

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	v "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/util/xCrypto"
)

const (
	url          = " "
	AcquiringBin = ""
)

type Service interface {
	Verify(context.Context, string, string) (*http.Response, error)
	Validate(context.Context, string, string) (*http.Response, error)
	Payout(context.Context, string, string) (*http.Response, error)
}

type Request struct {
	transport *http.Transport
	config    v.Visa
}

func New(config v.Visa) (*Request, error) {
	clientCACert, err := os.ReadFile(config.CaCertificateFile)
	if err != nil {
		return nil, err
	}

	//Load Client Key Pair
	clientKeyPair, err := tls.LoadX509KeyPair(config.ClientCertificateFile, config.ClientCertificateKeyFile)

	clientCertPool, _ := x509.SystemCertPool()
	if clientCertPool == nil {
		clientCertPool = x509.NewCertPool()
	}

	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientKeyPair},
		RootCAs:      clientCertPool,
	}

	return &Request{&http.Transport{
		TLSClientConfig: tlsConfig,
	}, config}, nil
}

func (r *Request) Verify(ctx context.Context, body, method string) (int, string, error) {
	path := "/visadirect-connect/v1/accounts/verify"
	var err error
	client := &http.Client{Transport: r.transport}

	apiUrl := r.config.Url + path
	crypto := xCrypto.New(nil)

	encData := map[string]string{"encData": crypto.CreateJWE(body, r.config.KeyId, r.config.MleServerPublicCertificatePath)}

	encryptedPayload, err := json.Marshal(encData)
	if err != nil {
		return -1, "", fmt.Errorf("error marshalling encrypted payload")
	}

	var request *http.Request = nil
	if body != "" {
		request, err = http.NewRequest(method, apiUrl, bytes.NewBuffer(encryptedPayload))
	} else {
		request, err = http.NewRequest(method, apiUrl, nil)
	}

	if err != nil {
		return -1, "", err
	}
	request.SetBasicAuth(r.config.User, r.config.Pass)
	request.Header.Set("keyId", r.config.KeyId)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return -1, "", err
	}
	encryptedResponsePayload, err := ioutil.ReadAll(response.Body)

	decryptedData, err := crypto.DecryptJWE(string(encryptedResponsePayload), r.config.MleClientPrivateKeyPath)
	if err != nil {
		decryptedData = string(encryptedResponsePayload)
	}

	return response.StatusCode, string(decryptedData), err
}
func (r *Request) Validate(ctx context.Context, body, method string) (int, string, error) {
	path := "/visadirect-connect/v1/accounts/payout/validate"
	var err error
	client := &http.Client{Transport: r.transport}

	apiUrl := r.config.Url + path

	crypto := xCrypto.New(nil)

	encData := map[string]string{"encData": crypto.CreateJWE(body, r.config.KeyId, r.config.MleServerPublicCertificatePath)}

	encryptedPayload, err := json.Marshal(encData)
	if err != nil {
		return -1, "", fmt.Errorf("error marshalling encrypted payload")
	}

	var request *http.Request = nil
	if body != "" {
		request, err = http.NewRequest(method, apiUrl, bytes.NewBuffer(encryptedPayload))
	} else {
		request, err = http.NewRequest(method, apiUrl, nil)
	}

	if err != nil {
		return -1, "", err
	}
	request.SetBasicAuth(r.config.User, r.config.Pass)
	request.Header.Set("keyId", r.config.KeyId)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return -1, "", err
	}
	encryptedResponsePayload, err := ioutil.ReadAll(response.Body)

	decryptedData, err := crypto.DecryptJWE(string(encryptedResponsePayload), r.config.MleClientPrivateKeyPath)
	if err != nil {
		decryptedData = string(encryptedResponsePayload)
	}

	return response.StatusCode, string(decryptedData), err
}
func (r *Request) Payout(ctx context.Context, body, method string) (int, string, error) {
	path := "/visadirect-connect/v1/accounts/payout"
	var err error
	client := &http.Client{Transport: r.transport}

	apiUrl := r.config.Url + path

	crypto := xCrypto.New(nil)

	encData := map[string]string{"encData": crypto.CreateJWE(body, r.config.KeyId, r.config.MleServerPublicCertificatePath)}

	encryptedPayload, err := json.Marshal(encData)
	if err != nil {
		return -1, "", fmt.Errorf("error marshalling encrypted payload")
	}

	var request *http.Request = nil
	if body != "" {
		request, err = http.NewRequest(method, apiUrl, bytes.NewBuffer(encryptedPayload))
	} else {
		request, err = http.NewRequest(method, apiUrl, nil)
	}

	if err != nil {
		return -1, "", err
	}
	request.SetBasicAuth(r.config.User, r.config.Pass)
	request.Header.Set("keyId", r.config.KeyId)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return -1, "", err
	}
	encryptedResponsePayload, err := ioutil.ReadAll(response.Body)

	decryptedData, err := crypto.DecryptJWE(string(encryptedResponsePayload), r.config.MleClientPrivateKeyPath)
	if err != nil {
		decryptedData = string(encryptedResponsePayload)
	}

	return response.StatusCode, string(decryptedData), err
}
