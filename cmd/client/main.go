package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/RafaelPereiraSantos/grpc-test/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := grpc.NewClient("localhost:"+os.Getenv("PROTO_SERVER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	connectWithServerAndIdentifyUser(conn)
}

func connectWithServerAndIdentifyUser(conn *grpc.ClientConn) {
	client := pb.NewCDPIntegrationClient(conn)

	userID := "123"

	resp, err := client.Identify(context.Background(), &pb.IdentifyPayload{
		UserId: userID,
	})

	fmt.Println(resp, err)
}
