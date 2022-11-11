package gg

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/scooper/go-getter/pkg/context"
	"github.com/scooper/go-getter/pkg/settings"
)

type ggcontext struct {
	routes map[string]func(w http.ResponseWriter, request *http.Request)
	settings settings.Settings // this repeating is awful, think about better names
}

type GG interface {
	Route(r string, methods string, f func(request *context.Request) *context.Response)
	Start()
}

func (ctx *ggcontext) Route(r string, methods string, f func(request *context.Request) *context.Response) {
	wrapped := func(w http.ResponseWriter, request *http.Request) {
		ggrequest := context.CreateRequest(request)
		ggresponse := f(ggrequest)
		
		for header, value := range ggresponse.Headers {
			w.Header().Add(header, value)
		}

		w.Write([]byte(ggresponse.Body))
	}

	ctx.routes[r] = wrapped
}

func (ctx *ggcontext) Start() {

	// register handlers
	for route, action := range ctx.routes {
		http.HandleFunc(route, action)
	}

	// start server
	err := http.ListenAndServe(":"+ctx.settings.Port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	}
}

func CreateContext() GG {
	return &ggcontext{
		routes: make(map[string]func(w http.ResponseWriter, request *http.Request)),
		settings: settings.Get(),
	}
}
