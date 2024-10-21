package main

import (
	"flux-version/internals/container"
	"fmt"
	"log"
)

func main() {

	fmt.Println("starting....")
	server, err := container.NewContainer()
	if err != nil {
		log.Panic(err)
	}
	if err = server.Start(); err != nil {
		log.Panic(err)
	}
	fmt.Println("shutdown")

}
