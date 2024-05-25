package pkg

import "os"

type authorizationFunc func() bool

type Db struct {
	AuthorizationFn authorizationFunc
}

func NewDB() *Db {
	return &Db{
		AuthorizationFn: argsAuthorization,
	}
}

func (d *Db) IsAuthorized() bool {
	return d.AuthorizationFn()
}

func argsAuthorization() bool {
	user := os.Args[1]
	// super secure authorization layer
	// in a real application, this would be a database call
	if user == "admin" {
		return true
	}
	return false
}
