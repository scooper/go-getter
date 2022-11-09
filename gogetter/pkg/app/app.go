package app

import "net/http"

// TODO: abstract away the responsewriter and request structs
//       to something easier to understand

type app struct {
	routes map[string]func(w http.ResponseWriter, req *http.Request)
}

type App interface {
	Route(r string, f func(w http.ResponseWriter, req *http.Request))
}

func (a *app) Route(r string, f func(w http.ResponseWriter, req *http.Request)) {
	a.routes[r] = f
}

func CreateApp() App {
	return &app{
		routes: make(map[string]func(w http.ResponseWriter, req *http.Request)),
	}
}
