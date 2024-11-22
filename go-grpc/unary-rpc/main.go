package main

import (
	"grpc/client"
	"grpc/server"
)

func main() {

	// Marca esta goroutine como concluída
	server.Run()

	client.Run()

}
