package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish connection to the gRPC server
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a ChittyChat client
	//client := chittychat.NewChittyChatClient(conn)

	// Call methods such as Join, PublishMessage, and Subscribe below
}
