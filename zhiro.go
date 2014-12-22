package zhiro

import (
	"net/http"
)

type Zhiro struct {
	denyFunc http.HandlerFunc
	manager  *ChainManager
	realm    Realm
}

type Realm interface {
	GetAuthorization(princ Principal) Authorization
}

func NewZhiro(chain map[string]string) *Zhiro {
	zhiro := &Zhiro{func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/login", 302)
	}, NewChainManager()}
	for key, value := range chain {
		zhiro.manager.CreateChain(key, value)
	}
	return zhiro
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
