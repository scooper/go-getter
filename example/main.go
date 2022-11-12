package main

import (
	gg "github.com/scooper/go-getter"
	"github.com/scooper/go-getter/pkg/context"
)

func main() {
	ctx := gg.CreateContext()

	ctx.Route("/hello", "GET", func(request *context.Request) *context.Response {
		return context.Text("Hello World!")
	})

	ctx.Start()
}