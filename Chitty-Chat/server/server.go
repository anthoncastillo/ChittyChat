package main

//"fmt"
//"log"

import (
	"Chitty-Chat/Chitty-Chat/chittychat/Chitty-Chat/chittychat"
	"context"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	chittychat.UnimplementedChittyChatServer // Embed this to satisfy the interface
	clients                                  map[string]chittychat.ChittyChat_SubscribeServer
	clientId                                 int64
	lamportTime                              int64
	mutex                                    sync.Mutex
}

// Implement the Join method
func (s *ChittyChatServer) Join(ctx context.Context, info *chittychat.ClientInfo) (*chittychat.JoinResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clientId++
	
	//Technically client should pass its own timestamp here
	//But currently it only passes protobuf clientinfo which has no timestamp
	//Maybe fix for future, currently assume it's 1 and just increment localtime by 1
	//Also set localtime to be 1 at start so it ends up as 2 after the first join

	s.lamportTime++	//SHOULD be call to update

	log.Printf("Client %d joined at Lamport time %d", s.clientId, s.lamportTime)

	// Prepare the join message to be broadcasted
	joinMessage := &chittychat.ChatMessage{
		ClientInfo:    info,
		Content:     "Participant " + *&info.ClientName + " joined Chitty-Chat",
		LamportTime: s.lamportTime,
	}

	s.PublishToAll(joinMessage)

	s.lamportTime++

	// Return response to the joining client
	return &chittychat.JoinResponse{		//Move this to before publishToAll?
		Success:        true,
		LamportTime:    s.lamportTime,
		WelcomeMessage: "Welcome to ChittyChat, " + info.ClientName,
		ClientId:       s.clientId,
	}, nil
}

// Implement the Leave method
func (s *ChittyChatServer) Leave(ctx context.Context, info *chittychat.ClientInfo) (*chittychat.LeaveResponse, error) {
	// Implementation of the Leave method
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.lamportTime++
	log.Printf("Client %d left at Lamport time %d", info.ClientId, s.lamportTime)

	delete(s.clients, info.ClientName)

	leaveMessage := &chittychat.ChatMessage{
		ClientInfo:    info,
		Content:     info.ClientName + " left Chitty-Chat",
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

	updateLamportTime(&s.lamportTime, msg.LamportTime)
	log.Printf("Message published by client %d %s at Lamport time %d: %s", msg.ClientInfo.ClientId, msg.ClientInfo.ClientName, s.lamportTime, msg.Content)

	message := &chittychat.ChatMessage{
		ClientInfo:    msg.ClientInfo,
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
func (s *ChittyChatServer) Subscribe(clientInfo *chittychat.ClientInfo, stream chittychat.ChittyChat_SubscribeServer) error {
	
	clientID := "ID " + strconv.FormatInt(int64(clientInfo.ClientId),10) // This should be generated or passed when the client joins

	s.mutex.Lock()
	s.clients[clientID] = stream
	s.mutex.Unlock()

	// Keep the stream open and simulate broadcasting messages
	for {
		time.Sleep(5 * time.Second)
		/**
		err := stream.Send(&chittychat.ChatMessage{
			ClientInfo:    clientInfo,
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
			*/
	}
}

func (s *ChittyChatServer) PublishToAll(message *chittychat.ChatMessage) {
	// Broadcast the join message to all connected clients
	for clientID, stream := range s.clients {
		s.lamportTime++
		if err := stream.Send(message); err != nil {
			log.Printf("Error sending join message to client %s: %v", clientID, err)
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

	// Create an instance of ChittyChatServer
	chittychatServer := &ChittyChatServer{

		clients:     make(map[string]chittychat.ChittyChat_SubscribeServer),
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
