package main

import (
	"Chitty-Chat/Chitty-Chat/chittychat/Chitty-Chat/chittychat"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	// Establish connection to the gRPC server
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a ChittyChat client
	client := chittychat.NewChittyChatClient(conn)

	// Define client ID and name
	clientId := "client123"
	clientName := "Client 123"

	// Join the chat
	joinChat(client, clientId, clientName)

	// Start subscribing to messages in a separate goroutine
	go subscribeToMessages(client)

	// Publish a few messages
	for i := 0; i < 3; i++ {
		publishMessage(client, clientId, "Hello, ChittyChat!")
		time.Sleep(2 * time.Second)
	}

	// Leave the chat after sending a few messages
	leaveChat(client, clientId)
}

func joinChat(client chittychat.ChittyChatClient, clientId, clientName string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	joinResp, err := client.Join(ctx, &chittychat.ClientInfo{
		ClientId:   clientId,
		ClientName: clientName,
	})
	if err != nil {
		log.Fatalf("Failed to join chat: %v", err)
	}

	log.Printf("Joined ChittyChat: %s", joinResp.WelcomeMessage)
	log.Printf("Joined at Lamport time: %d", joinResp.LamportTime)
}

func publishMessage(client chittychat.ChittyChatClient, clientId string, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pubResp, err := client.PublishMessage(ctx, &chittychat.ChatMessage{
		ClientId:    clientId,
		Content:     content,
		LamportTime: time.Now().Unix(), // Could also use a real Lamport clock
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	log.Printf("Message published at Lamport time: %d", pubResp.LamportTime)
}

func subscribeToMessages(client chittychat.ChittyChatClient) {
	// Create an empty context for the subscription
	stream, err := client.Subscribe(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to subscribe to messages: %v", err)
	}

	// Receive messages from the server
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		log.Printf("Received message from %s at Lamport time %d: %s",
			msg.ClientId, msg.LamportTime, msg.Content)
	}
}

func leaveChat(client chittychat.ChittyChatClient, clientId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	leaveResp, err := client.Leave(ctx, &chittychat.ClientInfo{
		ClientId: clientId,
	})
	if err != nil {
		log.Fatalf("Failed to leave chat: %v", err)
	}

	log.Printf("Left chat at Lamport time: %d", leaveResp.LamportTime)
}
