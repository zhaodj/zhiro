package zhiro

import(
)

type Subject struct{
    principal Principal
    prinLoaded bool
    authorization Authorization
    authLoaded bool
}

type Principal interface{
    Key()string
}

type Authorization interface{
    Roles()[]string
    Perms()[]string
}
