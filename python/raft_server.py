import time
import random
import grpc
import raft_pb2
import raft_pb2_grpc
from concurrent import futures

class RaftServer(raft_pb2_grpc.RaftServicer):
    def __init__(self, node_id, peers):
        self.node_id = node_id
        self.peers = peers
        self.state = "follower"
        self.current_term = 0
        self.voted_for = None
        self.logs = []  # Stores logs as list of LogEntry
        self.commit_index = 0
        self.last_applied = 0
        self.vote_count = 0
        self.leader_id = None

    def RequestVote(self, request, context):
        print(f"Node {self.node_id} runs RPC RequestVote called by Node {request.candidate_id}")

        # Handle vote request
        if request.term > self.current_term:
            self.current_term = request.term
            self.voted_for = None
            self.state = "follower"

        vote_granted = False
        if (self.voted_for is None or self.voted_for == request.candidate_id) and request.term >= self.current_term:
            vote_granted = True
            self.voted_for = request.candidate_id

        return raft_pb2.RequestVoteResponse(
            term=self.current_term,
            vote_granted=vote_granted
        )

    def AppendEntries(self, request, context):
        print(f"Node {self.node_id} runs RPC AppendEntries called by Node {request.leader_id}")

        # If the term in the AppendEntries request is greater, become a follower
        if request.term > self.current_term:
            self.current_term = request.term
            self.state = "follower"
            self.leader_id = str(request.leader_id)

        # Append entries to the log
        self.logs.extend(request.entries)
        self.commit_index = max(self.commit_index, request.leader_commit)

        return raft_pb2.AppendEntriesResponse(
            term=self.current_term,
            success=True
        )

    def start_election(self):
        self.state = "candidate"
        self.current_term += 1
        self.voted_for = self.node_id
        self.vote_count = 1  # Vote for self

        for peer in self.peers:
            if peer == self.node_id:
                continue

            # Send RequestVote to peers
            self.send_request_vote(peer)

    def send_request_vote(self, peer):
        # Assuming gRPC connection is available
        with grpc.insecure_channel(f'{peer}:50051') as channel:
            stub = raft_pb2_grpc.RaftStub(channel)
            request = raft_pb2.RequestVoteRequest(
                term=self.current_term,
                candidate_id=self.node_id
            )
            response = stub.RequestVote(request)

            if response.vote_granted:
                self.vote_count += 1
                if self.vote_count > len(self.peers) // 2:
                    self.state = "leader"
                    self.leader_id = self.node_id
                    print(f"Node {self.node_id} is now the leader")
                    break

    def run_heartbeat(self):
        while True:
            time.sleep(1)
            if self.state == "leader":
                # Send AppendEntries RPC to followers
                self.send_heartbeat()

    def send_heartbeat(self):
        for peer in self.peers:
            if peer == self.node_id:
                continue

            with grpc.insecure_channel(f'{peer}:50051') as channel:
                stub = raft_pb2_grpc.RaftStub(channel)
                request = raft_pb2.AppendEntriesRequest(
                    term=self.current_term,
                    leader_id=self.node_id,
                    entries=[],  # Empty entries for heartbeat
                    leader_commit=self.commit_index
                )
                stub.AppendEntries(request)

# Start the gRPC server
def serve(node_id, peers):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    raft_server = RaftServer(node_id, peers)
    raft_pb2_grpc.add_RaftServicer_to_server(raft_server, server)

    server.add_insecure_port('[::]:50051')
    print(f"Node {node_id} listening on port 50051")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    node_id = "node1"  # Get from environment or args
    peers = ["node1", "node2", "node3", "node4", "node5"]  # List of peer nodes
    serve(node_id, peers)
