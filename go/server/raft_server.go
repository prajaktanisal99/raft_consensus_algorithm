package server

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	pb "example/go/raft"
)

type RaftServer struct {
	pb.UnimplementedRaftServer

	NodeID      string
	Peers       []string
	CurrentTerm int32
	VotedFor    string
	State       string // "follower", "candidate", "leader"
	VoteCount   int
	LeaderID    string

	Clients map[string]pb.RaftClient
	Mutex   sync.Mutex
}

func NewRaftServer(nodeID string, peers []string) *RaftServer {
	return &RaftServer{
		NodeID:      nodeID,
		Peers:       peers,
		CurrentTerm: 0,
		VotedFor:    "",
		State:       "follower",
		VoteCount:   0,
		LeaderID:    "",
		Clients:     make(map[string]pb.RaftClient),
	}
}

// RPC Handlers

func (r *RaftServer) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	candidateID := strconv.Itoa(int(req.CandidateId))
	fmt.Printf("Node %s runs RPC RequestVote called by Node %s\n", r.NodeID, candidateID)

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	if req.Term > r.CurrentTerm {
		r.CurrentTerm = req.Term
		r.VotedFor = ""
		r.State = "follower"
	}

	voteGranted := false
	if (r.VotedFor == "" || r.VotedFor == candidateID) && req.Term >= r.CurrentTerm {
		voteGranted = true
		r.VotedFor = candidateID
	}

	return &pb.RequestVoteResponse{
		Term:        r.CurrentTerm,
		VoteGranted: voteGranted,
	}, nil
}

func (r *RaftServer) AppendEntries(ctx context.Context, req *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	leaderID := strconv.Itoa(int(req.LeaderId))
	fmt.Printf("Node %s runs RPC AppendEntries called by Node %s\n", r.NodeID, leaderID)

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	if req.Term >= r.CurrentTerm {
		r.CurrentTerm = req.Term
		r.State = "follower"
		r.VotedFor = ""
		r.LeaderID = leaderID
	}

	return &pb.AppendEntriesResponse{
		Term:    r.CurrentTerm,
		Success: true,
	}, nil
}
