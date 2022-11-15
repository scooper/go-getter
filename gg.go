package gg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/scooper/go-getter/pkg/context"
	"github.com/scooper/go-getter/pkg/settings"
	"github.com/scooper/go-getter/pkg/utils"
)

type ggcontext struct {
	routes map[string]func(w http.ResponseWriter, request *http.Request)
	settings *settings.Settings // this repeating is awful, think about better names
	logger utils.Logger
}

type GG interface {
	Route(r string, methods string, f func(request *context.Request) *context.Response)
	Start()
}

func (ctx *ggcontext) Route(r string, methods string, f func(request *context.Request) *context.Response) {
	wrapped := func(w http.ResponseWriter, request *http.Request) {
		methodSlice := strings.Split(methods, ",")
		validMethod := true

		// check if method is valid
		for _, m := range methodSlice {
			if m == request.Method {
				validMethod = false
				break
			}
		}

		ggrequest := context.CreateRequest(request)
		var ggresponse *context.Response

		if validMethod {
			ggresponse = f(ggrequest)
		} else {
			// default invalid response
			ggresponse.StatusCode = 403
			ggresponse.Body = "Forbidden"
			ggresponse.Headers["Content-Type"] = "text/plain"
		}
		
		if ggresponse.Error != nil {
			errorStr := ggresponse.Error.Error()
			ctx.logger.Error(fmt.Sprintf("Server Error: %s", errorStr))

			if !ctx.settings.Debug {
				errorStr = ""
			}

			ggresponse.Body = fmt.Sprintf("<!DOCTYPE><html><body><h1>Internal Server Error (%d)</h1><p>%s</p></body></html>",
			ggresponse.StatusCode,
			errorStr)

			ggresponse.Headers["Content-Type"] = "text/html"
		}
		
		for header, value := range ggresponse.Headers {
			w.Header().Add(header, value)
		}
		
		w.WriteHeader(ggresponse.StatusCode)
		w.Write([]byte(ggresponse.Body))

		ctx.logger.Info(fmt.Sprintf("%s: %s - %d", request.Method, request.RequestURI, ggresponse.StatusCode))
	}

	ctx.routes[r] = wrapped
}

func (ctx *ggcontext) Start() {

	// register handlers
	for route, action := range ctx.routes {
		http.HandleFunc(route, action)
	}

	// start server
	ctx.logger.Info("Starting Server")
	ctx.logger.Info("Listening on :"+ctx.settings.Port)
	err := http.ListenAndServe(":"+ctx.settings.Port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		ctx.logger.Error("Server Closed")
	}
}

func CreateContext() GG {
	logger := utils.CreateLogger()

	s, settingsErr := settings.Get()

	if settingsErr != nil {
		logger.Error("Problem opening settings.json: " + settingsErr.Error())
	}

	return &ggcontext{
		routes: make(map[string]func(w http.ResponseWriter, request *http.Request)),
		settings: s,
		logger: logger,
	}
}
