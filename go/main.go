package main

import (
	"log"
	"net"
	"os"
	"strings"

	pb "example/go/raft" // gRPC-generated code
	"example/go/server"  // your own raft logic (raft_server.go etc.)

	"google.golang.org/grpc"
)

func main() {
	nodeID := os.Getenv("NODE_ID")
	peersEnv := os.Getenv("PEERS")
	peers := strings.Split(peersEnv, ",")

	raft := server.NewRaftServer(nodeID, peers)

	go raft.RunElectionLoop()
	go raft.RunHeartbeatLoop()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRaftServer(grpcServer, raft)
	log.Printf("Node %s gRPC server running on port 50051", nodeID)
	grpcServer.Serve(lis)
}
