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
)

type ChittyChatServer struct {
	chittychat.UnimplementedChittyChatServer // Embed this to satisfy the interface
	clients                                  map[int64]chittychat.ChittyChat_SubscribeServer
	clientId                                 int64
	lamportTime                              int64
	mutex                                    sync.Mutex
}

func (s *ChittyChatServer) Join(ctx context.Context, info *chittychat.ClientInfo) (*chittychat.JoinResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clientId++
	updateLamportTime(&s.lamportTime, info.LamportTime)

	log.Printf("Client %d joined at Lamport time %d", s.clientId, s.lamportTime)

	// Prepare the join message to be broadcasted
	joinMessage := &chittychat.ChatMessage{
		ClientInfo: info,
		Content:    "Participant " + *&info.ClientName + " joined Chitty-Chat",
	}

	s.PublishToAll(joinMessage)

	s.lamportTime++

	// Return response to the joining client
	return &chittychat.JoinResponse{ //Move this to before publishToAll?
		Success:        true,
		LamportTime:    s.lamportTime,
		WelcomeMessage: "Welcome to ChittyChat, " + info.ClientName,
		ClientId:       s.clientId,
	}, nil
}

func (s *ChittyChatServer) Leave(ctx context.Context, info *chittychat.ClientInfo) (*chittychat.LeaveResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	updateLamportTime(&s.lamportTime, info.LamportTime)

	log.Printf("Client %d left at Lamport time %d", info.ClientId, s.lamportTime)

	delete(s.clients, info.ClientId)

	leaveMessage := &chittychat.ChatMessage{
		ClientInfo: info,
		Content:    info.ClientName + " left Chitty-Chat",
	}

	s.PublishToAll(leaveMessage)

	s.lamportTime++

	return &chittychat.LeaveResponse{
		Success:     true,
		LamportTime: s.lamportTime,
	}, nil
}

func (s *ChittyChatServer) PublishMessage(ctx context.Context, msg *chittychat.ChatMessage) (*chittychat.PublishResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	updateLamportTime(&s.lamportTime, msg.ClientInfo.LamportTime)
	log.Printf("Message published by client %d %s at Lamport time %d: %s", msg.ClientInfo.ClientId, msg.ClientInfo.ClientName, s.lamportTime, msg.Content)

	s.PublishToAll(msg)

	s.lamportTime++

	return &chittychat.PublishResponse{
		Success:     true,
		LamportTime: s.lamportTime,
	}, nil
}

func (s *ChittyChatServer) Subscribe(clientInfo *chittychat.ClientInfo, stream chittychat.ChittyChat_SubscribeServer) error {

	clientID := clientInfo.ClientId

	s.mutex.Lock()
	s.clients[clientID] = stream
	s.mutex.Unlock()

	// Keep the stream open and simulate broadcasting messages
	for {
		time.Sleep(5 * time.Second)
	}
}

func (s *ChittyChatServer) PublishToAll(message *chittychat.ChatMessage) {
	// Broadcast the join message to all connected clients
	for clientID, stream := range s.clients {
		s.lamportTime++
		message.ClientInfo.LamportTime = s.lamportTime
		if err := stream.Send(message); err != nil {
			log.Printf("Error sending join message to client %d: %v", clientID, err)
			delete(s.clients, clientID) // Remove client if there's an error
		}
	}
}

func updateLamportTime(local *int64, remote int64) {
	*local = max(*local, remote) + 1
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Creates an instance of ChittyChatServer
	chittychatServer := &ChittyChatServer{

		clients:     make(map[int64]chittychat.ChittyChat_SubscribeServer),
		clientId:    0,
		lamportTime: 0,
	}

	// Register the ChittyChatServer with the gRPC server
	chittychat.RegisterChittyChatServer(grpcServer, chittychatServer)

	log.Println("ChittyChat server started at port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
