package gg

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/scooper/go-getter/pkg/context"
)

type ggcontext struct {
	routes map[string]func(w http.ResponseWriter, request *http.Request)
}

type GG interface {
	Route(r string, methods string, f func(request *context.Request) *context.Response)
	Start()
}

func (ctx *ggcontext) Route(r string, methods string, f func(request *context.Request) *context.Response) {
	wrapped := func(w http.ResponseWriter, request *http.Request) {
		ggrequest := context.CreateRequest(request)
		ggresponse := f(ggrequest)
		ggresponse.Response.Write(w)
	}

	ctx.routes[r] = wrapped
}

func (ctx *ggcontext) Start() {

	// register handlers
	for route, action := range ctx.routes {
		http.HandleFunc(route, action)
	}

	// start server
	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	}
}

func CreateContext() GG {
	return &ggcontext{
		routes: make(map[string]func(w http.ResponseWriter, request *http.Request)),
	}
}
