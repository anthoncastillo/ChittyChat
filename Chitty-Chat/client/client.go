package main

import (
	"Chitty-Chat/Chitty-Chat/chittychat/Chitty-Chat/chittychat"
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	// Establish connection to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	//Create a scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Create a ChittyChat client
	client := chittychat.NewChittyChatClient(conn)

	log.Println("Please enter Username:")
	scanner.Scan()
	clientInfo := chittychat.ClientInfo {
		ClientName: scanner.Text(),
	}

	log.Printf("Welcome %s", clientInfo.ClientName)

	// Create a local Lamport timestamp
	var localTime int64 = 0

	// Join the chat
	joinChat(client, &clientInfo, &localTime)

	// Start subscribing to messages in a separate goroutine
	go subscribeToMessages(client, &clientInfo, &localTime)


	for {
		scanner.Scan()
		text := scanner.Text()
		if strings.Compare(text, "leave") == 0 {
			break
		}
		if len(text) < 128 {
			publishMessage(client, &clientInfo, text, &localTime)
		} else {
			log.Print("Message is too long, sorry")
		}
	}
	leaveChat(client, &clientInfo, &localTime)

}

func joinChat(client chittychat.ChittyChatClient, clientInfo *chittychat.ClientInfo, localTime *int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	*localTime++

	joinResp, err := client.Join(ctx, clientInfo)
	if err != nil {
		log.Fatalf("Failed to join chat: %v", err)
	}

	updateLamportTime(localTime, joinResp.LamportTime)

	*&clientInfo.ClientId = joinResp.ClientId

	log.Printf("%s (Lamport time %d)", joinResp.WelcomeMessage, *localTime)
}

func publishMessage(client chittychat.ChittyChatClient, clientInfo *chittychat.ClientInfo, content string, localTime *int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	*localTime++

	pubResp, err := client.PublishMessage(ctx, &chittychat.ChatMessage{
		ClientInfo:    clientInfo,
		Content:     content,
		LamportTime: *localTime,
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	updateLamportTime(localTime, pubResp.LamportTime)

	//log.Printf("Message published at Lamport time: %d", *localTime)
}

func subscribeToMessages(client chittychat.ChittyChatClient, clientInfo *chittychat.ClientInfo, localTime *int64) {
	// Create an empty context for the subscription
	stream, err := client.Subscribe(context.Background(), clientInfo)
	if err != nil {
		log.Fatalf("Failed to subscribe to messages: %v", err)
	}

	// Receive messages from the server
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		updateLamportTime(localTime, msg.LamportTime)
		log.Printf("%s : %s (Lamport time: %d)", msg.ClientInfo.ClientName, msg.Content, *localTime)
	}
}

func leaveChat(client chittychat.ChittyChatClient, clientInfo *chittychat.ClientInfo, localTime *int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	*localTime++

	leaveResp, err := client.Leave(ctx, clientInfo)
	if err != nil {
		log.Fatalf("Failed to leave chat: %v", err)
	}

	updateLamportTime(localTime, leaveResp.LamportTime)

	log.Printf("Left chat at Lamport time: %d", *localTime)
}

func updateLamportTime(local *int64, remote int64) {
	*local = max(*local, remote) + 1
}
