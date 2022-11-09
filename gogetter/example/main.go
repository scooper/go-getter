package main

import (
	"fmt"

	gg "github.com/scooper/go-getter"
)

func main() {
	app := gg.CreateApp()
	fmt.Println(app)

	// maybe it should be app.start?
	// maybe we should be creating a context, or gogetter instance, rather than an app?
	gg.Start()
}