package main

import (
	"gin-api/route"
)

func main() {
	router := route.NewRoute()
	router.Run()
}