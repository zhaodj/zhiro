package zhiro

import (
	"strings"
)

type ChainManager struct {
	filters map[string]Filter
}

func NewChainManager() *ChainManager {
	return &ChainManager{map[string]Filter{}}
}

func split(line string, delimiter int32, beginQuote int32, endQuote int32, retainQuote bool, trim bool) []string {
	s := []string{}
	l := strings.TrimSpace(line)
	return s
}

func addToChain(url string,pair1)

func (m *ChainManager) CreateChain(url string, chainDef string) {
	ft := split(chainDef, ',', '[', ']', true, true)
}
