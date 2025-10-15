package account

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	v "github.com/mhthrh/common_pkg/pkg/model/config"
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

func (r *Request) Verify(ctx context.Context, body, method string) (*http.Response, error) {
	path := "/visadirect-connect/v1/accounts/verify"
	var err error
	client := &http.Client{Transport: r.transport}

	apiUrl := r.config.Url + path
	var request *http.Request = nil
	if body != "" {
		request, err = http.NewRequest(method, apiUrl, bytes.NewBuffer([]byte(body)))
	} else {
		request, err = http.NewRequest(method, apiUrl, nil)
	}

	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(r.config.User, r.config.Pass)
	request.Header.Set("keyId", r.config.KeyId)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	return client.Do(request)
}
func (r *Request) Validate(ctx context.Context, body, method string) (*http.Response, error) {
	path := "/visadirect-connect/v1/accounts/payout/validate"
	var err error
	client := &http.Client{Transport: r.transport}

	apiUrl := r.config.Url + path
	var request *http.Request = nil
	if body != "" {
		request, err = http.NewRequest(method, apiUrl, bytes.NewBuffer([]byte(body)))
	} else {
		request, err = http.NewRequest(method, apiUrl, nil)
	}

	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(r.config.User, r.config.Pass)
	request.Header.Set("keyId", r.config.KeyId)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	return client.Do(request)
}
func (r *Request) Payout(ctx context.Context, body, method string) (*http.Response, error) {
	path := "/visadirect-connect/v1/accounts/payout"
	var err error
	client := &http.Client{Transport: r.transport}

	apiUrl := r.config.Url + path
	var request *http.Request = nil
	if body != "" {
		request, err = http.NewRequest(method, apiUrl, bytes.NewBuffer([]byte(body)))
	} else {
		request, err = http.NewRequest(method, apiUrl, nil)
	}

	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(r.config.User, r.config.Pass)
	request.Header.Set("keyId", r.config.KeyId)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	return client.Do(request)
}
