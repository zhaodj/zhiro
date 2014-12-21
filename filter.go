package zhiro

import (
	"net/http"
    "strings"
)

type Filter interface {
	Handle(w http.ResponseWriter, req *http.Request, val string) bool
    ProcessPathConfig(path string, config string)
}

type PathProcessor struct {
	paths map[string][]string
}

func (proc *PathProcessor) ProcessPathConfig(path string,config string){
    values := []string{}
    if len(config)>0{
        values = strings.Split(config,",")
    }
    proc.paths[path]=values
}

type FilterList struct{
    name string
    Filters []Filter
}

type AnonFilter struct {
    PathProcessor
}

func (filter *AnonFilter) Handle(w http.ResponseWriter, req *http.Request, val string) bool {
	return true
}
