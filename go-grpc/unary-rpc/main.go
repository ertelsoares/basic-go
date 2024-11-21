package main

import (
	"grpc/server"
)

func main() {

	// Marca esta goroutine como concluída
	server.Run()

}
