package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "example/go/raft"

	"google.golang.org/grpc"
)

// Start election timeout loop
func (r *RaftServer) RunElectionLoop() {
	for {
		timeout := time.Duration(1500+rand.Intn(1500)) * time.Millisecond
		time.Sleep(timeout)

		r.Mutex.Lock()
		if r.State != "leader" {
			log.Printf("Node %s timed out, starting election", r.NodeID)
			r.startElection()
		}
		r.Mutex.Unlock()
	}
}

func (r *RaftServer) startElection() {
	r.State = "candidate"
	r.CurrentTerm++
	r.VotedFor = r.NodeID
	r.VoteCount = 1 // vote for self

	for _, peer := range r.Peers {
		if peer == r.NodeID {
			continue
		}

		go func(peer string) {
			client := r.getClient(peer)
			if client == nil {
				return
			}

			idInt, _ := strconv.Atoi(r.NodeID)
			req := &pb.RequestVoteRequest{
				Term:        r.CurrentTerm,
				CandidateId: int32(idInt),
			}

			fmt.Printf("Node %s sends RPC RequestVote to Node %s\n", r.NodeID, peer)
			resp, err := client.RequestVote(context.Background(), req)
			if err != nil {
				return
			}

			r.Mutex.Lock()
			defer r.Mutex.Unlock()

			if resp.VoteGranted {
				r.VoteCount++
				if r.VoteCount > len(r.Peers)/2 && r.State == "candidate" {
					log.Printf("Node %s becomes leader (term %d)", r.NodeID, r.CurrentTerm)
					r.State = "leader"
					r.LeaderID = r.NodeID
				}
			} else if resp.Term > r.CurrentTerm {
				r.State = "follower"
				r.CurrentTerm = resp.Term
				r.VotedFor = ""
			}
		}(peer)
	}
}

// Start sending heartbeats if leader
func (r *RaftServer) RunHeartbeatLoop() {
	for {
		time.Sleep(1 * time.Second)

		r.Mutex.Lock()
		if r.State != "leader" {
			r.Mutex.Unlock()
			continue
		}

		for _, peer := range r.Peers {
			if peer == r.NodeID {
				continue
			}

			go func(peer string) {
				client := r.getClient(peer)
				if client == nil {
					return
				}

				idInt, _ := strconv.Atoi(r.NodeID)
				req := &pb.AppendEntriesRequest{
					Term:     r.CurrentTerm,
					LeaderId: int32(idInt),
				}

				fmt.Printf("Node %s sends RPC AppendEntries to Node %s\n", r.NodeID, peer)
				_, _ = client.AppendEntries(context.Background(), req)
			}(peer)
		}

		r.Mutex.Unlock()
	}
}

// gRPC connection cache
func (r *RaftServer) getClient(peerID string) pb.RaftClient {
	if client, ok := r.Clients[peerID]; ok {
		return client
	}

	host := fmt.Sprintf("%s:50051", peerID)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Printf("Node %s failed to connect to peer %s: %v", r.NodeID, peerID, err)
		return nil
	}

	client := pb.NewRaftClient(conn)
	r.Clients[peerID] = client
	return client
}
