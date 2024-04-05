package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "github.com/RafaelPereiraSantos/grpc-test/server"
)

type CDPIntegrationServer struct {
	cdpImplementation *CDPImplementation
	pb.UnimplementedCDPIntegrationServer
}

func NewCDPIntegrationServer(cdpImplementation *CDPImplementation) *CDPIntegrationServer {
	return &CDPIntegrationServer{cdpImplementation: cdpImplementation}
}

func (proto *CDPIntegrationServer) Identify(ctx context.Context, in *pb.IdentifyPayload) (*pb.IdentifyResponse, error) {
	fmt.Println("Identify called with", in)

	if in == nil {
		return nil, nil
	}

	err := proto.cdpImplementation.Identify(ctx, IdentifyUser{UserID: in.UserId})

	if err != nil {
		return &pb.IdentifyResponse{Status: "FAILED"}, err
	}

	return &pb.IdentifyResponse{Status: "SUCCESS"}, nil
}

func StartProtoServer(cdpImplementation *CDPImplementation) {
	fmt.Println("Starting proto server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lis, err := net.Listen("tcp", "localhost:"+os.Getenv("PROTO_SERVER_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCDPIntegrationServer(grpcServer, NewCDPIntegrationServer(cdpImplementation))
	grpcServer.Serve(lis)
}
