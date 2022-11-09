package gg

import (
	"github.com/scooper/go-getter/pkg/app"
	"github.com/scooper/go-getter/pkg/serve"
)

func CreateApp() app.App {
	return app.CreateApp()
}

func Start() {
	serve.Serve()
}
