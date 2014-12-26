package zhiro

import (
	"crypto/sha256"
	"net/http"
)

type Zhiro struct {
	denyFunc           http.HandlerFunc
	manager            *ChainManager
	realm              Realm
	credentialsService CredentialsService
}

type Realm interface {
	GetPrincipal(token AuthenticationToken) Principal
	GetAuthorization(princ Principal) Authorization
}

type CredentialsService interface {
	Match(token AuthenticationToken, princ Principal) bool
	Encrypt(plain string) string
}

type SimpleCredentialsService struct {
}

func (s *SimpleCredentialsService) Match(token AuthenticationToken, princ Principal) bool {
	return token.Credentials() == princ.Credentials()
}
func (s *SimpleCredentialsService) Encrypt(plain string) string {
	return plain
}

type SHA256CredentialsService struct {
	hashTimes int
}

func (s *SHA256CredentialsService) Match(token AuthenticationToken, princ Principal) bool {
	return s.Encrypt(token.Credentials()) == princ.Credentials()
}

func (s *SHA256CredentialsService) Encrypt(plain string) string {
	hash := sha256.New()
	hashed := hash.Sum([]byte(plain))
	for i := 0; i < s.hashTimes; i++ {
		hash.Reset()
		hashed = hash.Sum(hashed)
	}
	return string(hashed)
}

func NewZhiro(chain map[string]string, realm Realm) *Zhiro {
	zhiro := &Zhiro{func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/login", 302)
	}, NewChainManager(), realm, &SHA256CredentialsService{1000}}
	for key, value := range chain {
		zhiro.manager.CreateChain(key, value)
	}
	return zhiro
}

func (zh *Zhiro) SetCredentialsService(service CredentialsService) {
	zh.credentialsService = service
}

func (zh *Zhiro) DenyFunc() http.HandlerFunc {
	return zh.denyFunc
}

func (zh *Zhiro) Check(w http.ResponseWriter, req *http.Request) bool {
	return false
}

func (zh *Zhiro) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if zh.Check(w, req) {
		zh.DenyFunc()(w, req)
		return
	}
	next(w, req)
}
