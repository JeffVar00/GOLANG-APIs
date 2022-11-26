package server

import "net/http"

/*
 */

type Country struct {
	Name     string
	Language string
}

// variable global
var countries = []*Country{}

func New(addr string) *http.Server { //*http.Server indica que esto esta devolviendo (puntero)
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
