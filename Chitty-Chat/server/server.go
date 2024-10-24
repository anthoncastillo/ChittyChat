package main

//"fmt"
//"log"

import (
	"Chitty-Chat/Chitty-Chat/chittychat/Chitty-Chat/chittychat"
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChittyChatServer struct {
	chittychat.UnimplementedChittyChatServer // Embed this to satisfy the interface
	clients                                  map[string]chittychat.ChittyChat_SubscribeServer
	lamportTime                              int64
	mutex                                    sync.Mutex
}

// Implement the Join method
func (s *ChittyChatServer) Join(ctx context.Context, info *chittychat.ClientInfo) (*chittychat.JoinResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Increment Lamport time for the join event
	s.lamportTime++

	// Log the join event
	log.Printf("Client %s joined at Lamport time %d", info.ClientId, s.lamportTime)

	// Prepare the join message to be broadcasted
	joinMessage := &chittychat.ChatMessage{
		ClientId:    info.ClientId,
		Content:     "Participant " + info.ClientId + " joined Chitty-Chat",
		LamportTime: s.lamportTime,
	}

	s.PublishToAll(joinMessage)

	// Return response to the joining client
	return &chittychat.JoinResponse{
		Success:        true,
		LamportTime:    s.lamportTime,
		WelcomeMessage: "Welcome to ChittyChat!",
	}, nil
}

// Implement the Leave method
func (s *ChittyChatServer) Leave(ctx context.Context, info *chittychat.ClientInfo) (*chittychat.LeaveResponse, error) {
	// Implementation of the Leave method
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lamportTime++
	log.Printf("Client %s left at Lamport time %d", info.ClientId, s.lamportTime)

	leaveMessage := &chittychat.ChatMessage{
		ClientId:    info.ClientId,
		Content:     "Participant " + info.ClientId + " left Chitty-Chat",
		LamportTime: s.lamportTime,
	}

	s.PublishToAll(leaveMessage)

	return &chittychat.LeaveResponse{
		Success:     true,
		LamportTime: s.lamportTime,
	}, nil
}

// Implement the PublishMessage method
func (s *ChittyChatServer) PublishMessage(ctx context.Context, msg *chittychat.ChatMessage) (*chittychat.PublishResponse, error) {
	// Implementation of the PublishMessage method
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lamportTime++
	log.Printf("Message published by %s at Lamport time %d: %s", msg.ClientId, s.lamportTime, msg.Content)

	message := &chittychat.ChatMessage{
		ClientId:    msg.ClientId,
		Content:     msg.Content,
		LamportTime: s.lamportTime,
	}

	s.PublishToAll(message)

	return &chittychat.PublishResponse{
		Success:     true,
		LamportTime: s.lamportTime,
	}, nil
}

// Implement the Subscribe method
func (s *ChittyChatServer) Subscribe(empty *emptypb.Empty, stream chittychat.ChittyChat_SubscribeServer) error {
	clientID := "some_unique_id" // This should be generated or passed when the client joins

	s.mutex.Lock()
	s.clients[clientID] = stream
	s.mutex.Unlock()

	// Keep the stream open and simulate broadcasting messages
	for {
		time.Sleep(5 * time.Second)
		err := stream.Send(&chittychat.ChatMessage{
			ClientId:    clientID,
			Content:     "This is a broadcast message",
			LamportTime: s.lamportTime,
		})
		if err != nil {
			log.Printf("Error sending to client %s: %v", clientID, err)
			s.mutex.Lock()
			delete(s.clients, clientID)
			s.mutex.Unlock()
			return err
		}
	}
}

func (s *ChittyChatServer) PublishToAll(message *chittychat.ChatMessage) {
	// Broadcast the join message to all connected clients
	for clientID, stream := range s.clients {
		if err := stream.Send(message); err != nil {
			log.Printf("Error sending join message to client %s: %v", clientID, err)
			delete(s.clients, clientID) // Remove client if there's an error
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Create an instance of ChittyChatServer
	chittychatServer := &ChittyChatServer{

		clients:     make(map[string]chittychat.ChittyChat_SubscribeServer),
		lamportTime: 0,
	}

	// Register the ChittyChatServer with the gRPC server
	chittychat.RegisterChittyChatServer(grpcServer, chittychatServer)

	log.Println("ChittyChat server started at port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
