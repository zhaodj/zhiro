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

func split(line string, delimiter uint8, beginQuote uint8, endQuote uint8, retainQuote bool, trim bool) []string {
	tokens := []string{}
	l := strings.TrimSpace(line)
    s := []uint8{}
    inQuotes := false
    llen := len(l)
    for i:=0;i<llen;i++{
        c:=l[i]
        if c == beginQuote{
            if inQuotes && llen>(i+1) && l[i+1]==beginQuote{
                s = append(s,l[i+1])
                i++
            }else{
                inQuotes = !inQuotes
                if retainQuote{
                    s = append(s,c)
                }
            }
        }else if c == endQuote{
            inQuotes = !inQuotes
            if retainQuote{
                s = append(s,c)
            }
        }else if c == delimiter && !inQuotes{
            ss := string(s)
            if trim{
                ss = strings.TrimSpace(ss)
            }
            tokens = append(tokens,ss)
            s = []uint8{}
        }else{
            s = append(s,c)
        }
    }
    ss := string(s)
    if trim{
        ss = strings.TrimSpace(ss)
    }
    tokens = append(tokens,ss)
	return tokens
}

func addToChain(url string, filterName string, chainConf string){

}

func (m *ChainManager) CreateChain(url string, chainDef string) {
	ft := split(chainDef, ',', '[', ']', true, true)
}
