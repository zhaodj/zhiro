package zhiro

import(
    "strings"
)

type Matcher interface{
    Matches(pattern string, source string)bool
}

type AntMatcher struct{

}

func startWith(input string,char byte) bool{
    return input[0] == char
}

func endWith(input string,char byte) bool{
    return input[len(input)-1] == char
}

func matchString(pattern string, str string) bool{
    pattIdxStart,strIdxStart := 0,0
    pattIdxEnd,strIdxEnd := len(pattern)-1, len(str)-1

    hasStar := false
    for i:=0;i<=pattIdxEnd;i++{
        if pattern[i] == '*'{
            hasStar = true
            break
        }
    }

    if !hasStar{
        if pattIdxEnd != strIdxEnd{
            return false
        }
        for i:=0;i<=pattIdxEnd;i++{
            if pattern[i] != '?' && pattern[i] != str[i]{
                return false
            }
        }
        return true
    }

    if pattIdxEnd == 0{
        return true
    }

    for ch:=pattern[pattIdxStart];ch != '*' && strIdxStart <= strIdxEnd;ch=pattern[pattIdxStart]{
        if ch != '?' && ch != str[strIdxStart]{
            return false
        }
        pattIdxStart++
        strIdxStart++
    }

    if strIdxStart>strIdxEnd{
        for i := pattIdxStart; i<=pattIdxEnd;i++{
            if pattern[i] != '*'{
                return false
            }
        }
        return true
    }

    for ch:=pattern[pattIdxEnd];ch != '*' && strIdxStart <= strIdxEnd;ch=pattern[pattIdxEnd]{
        if ch != '?' && ch != str[strIdxEnd]{
            return false
        }
        pattIdxEnd--
        strIdxEnd--
    }

    if strIdxStart>strIdxEnd{
        for i := pattIdxStart; i<=pattIdxEnd;i++{
            if pattern[i] != '*'{
                return false
            }
        }
        return true
    }

    for pattIdxStart != pattIdxEnd && strIdxStart <= strIdxEnd{
        pattIdxTmp := -1
        for i:=pattIdxStart+1;i<=pattIdxEnd;i++{
            if pattern[i] == '*'{
                pattIdxTmp = i
                break
            }
        }

        if pattIdxTmp == pattIdxStart+1{
            pattIdxStart++
            continue
        }

        pattLen := pattIdxTmp - pattIdxStart -1
        strLen := strIdxEnd - strIdxStart + 1
        foundIdx := -1
        strLoop:
        for i:=0;i<=strLen-pattLen;i++{
            for j:=0;j<pattLen;j++{
                ch:=pattern[pattIdxStart+j+1]
                if ch!='?' && ch != str[strIdxStart+i+j]{
                    continue strLoop
                }
            }
            foundIdx = strIdxStart+i
            break
        }

        if foundIdx == -1{
            return false
        }

        pattIdxStart = pattIdxTmp
        strIdxStart = foundIdx + pattLen
    }

    for i:=pattIdxStart;i<=pattIdxEnd;i++{
        if pattern[i] != '*'{
            return false
        }
    }

    return true
}

func (m *AntMatcher)Matches(pattern string,source string) bool{
    if startWith(pattern,'/') != startWith(source,'/'){
        return false
    }
    patts := strings.Split(pattern, "/")
    sours := strings.Split(source, "/")

    pattIdxStart :=0
    pattIdxEnd := len(patts) - 1
    sourIdxStart :=0
    sourIdxEnd := len(sours) - 1

    for pattIdxStart<=pattIdxEnd && sourIdxStart<=sourIdxEnd{
        patDir := patts[pattIdxStart]
        if patDir == "**"{
            break
        }
        if !matchString(patDir, sours[sourIdxStart]){
            return false
        }
        pattIdxStart++
        sourIdxStart++
    }

    if sourIdxStart > sourIdxEnd{
        if pattIdxStart > pattIdxEnd{
            if endWith(pattern,'/'){
                return endWith(source,'/')
            }
            return !endWith(source,'/')
        }
        if pattIdxStart == pattIdxEnd && patts[pattIdxStart]=="*" && endWith(source,'/'){
            return true
        }
        for i:=pattIdxStart;i<=pattIdxEnd;i++{
            if patts[i] != "**"{
                return false
            }
        }
        return true
    }else if pattIdxStart > pattIdxEnd{
        return false
    }

    for pattIdxStart <= pattIdxEnd && sourIdxStart <= sourIdxEnd{
        patDir := patts[pattIdxEnd]
        if patDir == "**"{
            break
        }
        if !matchString(patDir,sours[sourIdxEnd]){
            return false
        }
        pattIdxEnd--
        sourIdxEnd--
    }

    if sourIdxStart > sourIdxEnd{
        for i:=pattIdxStart;i<=pattIdxEnd;i++{
            if patts[i] != "**"{
                return false
            }
        }
        return true
    }

    for pattIdxStart != pattIdxEnd && sourIdxStart <= sourIdxEnd{
        pattIdxTmp := -1
        for i:=pattIdxStart+1;i<=pattIdxEnd;i++{
            if patts[i] == "**"{
                pattIdxTmp=i
                break
            }
        }

        if pattIdxTmp == pattIdxStart + 1{
            pattIdxStart++
            continue
        }

        pattLen := pattIdxTmp - pattIdxStart - 1
        sourLen := sourIdxEnd - sourIdxStart + 1
        foundIdx := -1

        sourLoop:
        for i:=0;i<=sourLen-pattLen;i++{
            for j:=0;i<pattLen;j++{
                if !matchString(patts[pattIdxStart+j+1],sours[sourIdxStart+i+j]){
                    continue sourLoop
                }
            }
            foundIdx = sourIdxStart + i
            break
        }

        if foundIdx == -1{
            return false
        }

        pattIdxStart = pattIdxTmp
        sourIdxStart = foundIdx + pattLen

    }

    for i:=pattIdxStart;i<=pattIdxEnd;i++{
        if patts[i] != "**"{
            return false
        }
    }

    return true
}
