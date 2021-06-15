package auth

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/apache/pulsar-client-go/pulsar"
)

type tuYaAuthProvider struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	accessID  string
	accessKey string
}

func md5Hex(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Provider is a interface of authentication providers.
func (a *tuYaAuthProvider) Init() error {
	key := md5Hex(a.accessID + md5Hex(a.accessKey))
	key = key[8:24]
	a.UserName = a.accessID
	a.Password = key
	return nil
}

func (a *tuYaAuthProvider) Name() string {
	return "auth1"
}

func (a *tuYaAuthProvider) GetTLSCertificate() (*tls.Certificate, error) {
	return nil, nil
}

func (a *tuYaAuthProvider) GetData() ([]byte, error) {
	return json.Marshal(a)
}

func (a *tuYaAuthProvider) Close() error {
	return nil
}

func (p *tuYaAuthProvider) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func (p *tuYaAuthProvider) Transport() http.RoundTripper {
	return nil
}

func (p *tuYaAuthProvider) WithTransport(tripper http.RoundTripper) error {
	return nil
}

func newAuthenticationWithTuYa(accessId, accessKey string) (pulsar.Authentication, error) {
	return &tuYaAuthProvider{accessKey: accessKey, accessID: accessId}, nil
}

func NewAuthenticationDataProviderWithTuYa(accessId, accessKey string) (pulsar.Authentication, error) {

	return newAuthenticationWithTuYa(accessId, accessKey)
}
