package main

import (
	"gin-plus/routes"
)

func main() {
	router := routes.Api()
	router.Run()
}
