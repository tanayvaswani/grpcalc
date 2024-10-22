package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverAddr := flag.String(
		"server",
		"localhost:8080",
		"The server is in the format host:port",
	)
	flag.Parse()

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalln("Failed to dial server", err)
	}
	defer conn.Close()
}
