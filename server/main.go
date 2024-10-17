package main

import (
	"context"
	"log"
	"net"

	"github.com/tanayvaswani/grpcalc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create listener", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterCalculatorServer(grpcServer, &server{})
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
