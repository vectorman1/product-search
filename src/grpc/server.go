package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/vectorman1/product-search/proto/generated/product_service"

	"github.com/vectorman1/product-search/grpc/presentation"

	"github.com/elastic/go-elasticsearch/v8"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":10001")
	if err != nil {
		panic(err)
	}

	esConfig := elasticsearch.Config{
		Username: "elastic",
		Password: "changeme",
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}
	es, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	productService := presentation.NewProductServiceServer(es)
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
