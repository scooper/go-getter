package main

import (
	"fmt"

	gg "github.com/scooper/go-getter"
	"github.com/scooper/go-getter/pkg/context"
)

func main() {
	ctx := gg.CreateContext()
	fmt.Println(ctx)

	ctx.Route("/hello", "GET", func(request *context.Request) *context.Response {
		return context.Text("Hello World!")
	})

	// maybe we should be creating a context, or gogetter instance, rather than an app?
	ctx.Start()
}