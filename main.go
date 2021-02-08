package main

import (
	helloworld "hello-world/helloworld"
	"hello-world/router"
)

func main() {
	var routers = router.NewEcho()

	helloworldService := helloworld.NewService()
	helloworldHandler := helloworld.NewHandler(helloworldService)
	routers.GET("/hello-world", helloworldHandler.HelloWorld)

	go router.Runs(routers)

	router.Shutdown(routers)
}
