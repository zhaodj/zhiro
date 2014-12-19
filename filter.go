package zhiro

import (
	"net/http"
)

type Filter interface {
	Handle(w http.ResponseWriter, req *http.Request, val string) bool
}

type pathProcessor struct {
	paths map[string][]string
}

type AnonFilter struct {
}

func (filter *AnonFilter) Handle(w http.ResponseWriter, req *http.Request, val string) bool {
	return true
}
