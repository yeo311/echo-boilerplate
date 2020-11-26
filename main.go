package main

import "github.com/yeo311/echo-boilerplate/route"

func main() {
	r := route.Route()

	r.Logger.Fatal(r.Start(":8080"))
}
