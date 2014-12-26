package zhiro

import ()

type Subject struct {
	principal     Principal
	prinLoaded    bool
	authorization Authorization
	authLoaded    bool
}

type Principal interface {
	Key() string
	Credentials() string
}

type AuthenticationToken interface {
	Credentials() string
}

type Authorization interface {
	Roles() []string
	Perms() []string
}

type UsernamePasswordToken struct {
	username string
	password string
}

func (t *UsernamePasswordToken) Credentials() string {
	return t.password
}
