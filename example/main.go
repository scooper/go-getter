package main

import (
	gg "github.com/scooper/go-getter"
	"github.com/scooper/go-getter/pkg/context"
)

type IndexData struct {
	Title string
	Heading string
	Content string
}

func main() {
	ctx := gg.CreateContext()

	ctx.Route("/hello", "GET", func(request *context.Request) *context.Response {
		return context.Text("Hello World!")
	})

	ctx.Route("/", "GET", func(request *context.Request) *context.Response {
		// TODO: init index data and pass into function
		return context.Template("index.html", &IndexData{
			Title: "Example Index",
			Heading: "Example Home",
			Content: "Lorem ispum.....",
		})
	})

	ctx.Start()
}