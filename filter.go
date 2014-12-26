package zhiro

import (
	"net/http"
	"strings"
)

type Filter interface {
	Handle(w http.ResponseWriter, req *http.Request, val []string) bool
	ProcessPathConfig(path string, config string)
	PathConfig(path string) []string
}

type PathProcessor struct {
	paths map[string][]string
}

func (proc *PathProcessor) ProcessPathConfig(path string, config string) {
	values := []string{}
	if len(config) > 0 {
		values = strings.Split(config, ",")
	}
	proc.paths[path] = values
}

func (proc *PathProcessor) PathConfig(path string) []string {
	return proc.paths[path]
}

type FilterList struct {
	name    string
	Filters []Filter
}

type AnonFilter struct {
	PathProcessor
}

func (filter *AnonFilter) Handle(w http.ResponseWriter, req *http.Request, val []string) bool {
	return true
}

type FormAuthenticationFilter struct {
	PathProcessor
	usernameParam string
	passwordParam string
}

func (filter *FormAuthenticationFilter) Handle(w http.ResponseWriter, req *http.Request, val []string) bool {
	req.ParseForm()
	params := req.PostForm
	username, password := params.Get(filter.usernameParam), params.Get(filter.passwordParam)
	if len(username) > 0 && len(password) > 0 {
		token := &UsernamePasswordToken{username, password}
	}
	return false
}
