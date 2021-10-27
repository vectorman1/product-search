package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"product_service/generated/product_service"
	"product_service/presentation"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	productService := presentation.NewProductServiceServer()
	product_service.RegisterProductServiceServer(server, productService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC grpc-server...")

			server.GracefulStop()
		}
	}()

	server.Serve(listen)
}
