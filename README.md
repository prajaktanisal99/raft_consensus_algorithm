# Raft Consensus Algorithm Implementation

This repository contains an implementation of the Raft Consensus Algorithm. Raft is used to manage a replicated log for distributed systems, providing a way to ensure consistency across a cluster of nodes. This implementation is written in **Go** (for Leader Election) and **Python** (for Log Replication).

## Table of Contents

- [Overview](#overview)
- [Components](#components)
- [Setup and Running the System](#setup-and-running-the-system)
- [Testing the System](#testing-the-system)
- [Log Output](#log-output)
- [Known Issues](#known-issues)
- [License](#license)

---

## Overview

Raft is a distributed consensus algorithm designed to manage a replicated log across a cluster of nodes. The algorithm ensures that each node in the cluster agrees on a single order of log entries. This implementation uses **gRPC** for communication and Docker for containerization.

The system consists of:
- **Leader Election:** Handles electing a leader in the cluster.
- **Log Replication:** Ensures log consistency across all follower nodes by sending log entries from the leader to the followers.

---

## Components

1. **Leader Election (Go):**
    - Each node starts as a follower.
    - A node times out and starts an election.
    - Nodes vote for a candidate.
    - The candidate becomes the leader if it receives votes from the majority.
    - The leader sends `AppendEntries` to maintain its leadership.

2. **Log Replication (Python):**
    - Leader sends log entries to followers using `AppendEntries`.
    - Followers acknowledge the entries and update their log.
    - This ensures all nodes in the cluster have the same log entries, guaranteeing consistency.

3. **Dockerized Setup:**
    - Each node runs in its own container.
    - Docker Compose is used to configure and manage the containers.

---

## Setup and Running the System

1. **Prerequisites:**
    - Install [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) if you don't have them installed.
    - Make sure you have Go and Python installed for local development.

2. **Clone the Repository:**
    ```bash
    git clone https://github.com/your-username/raft-consensus.git
    cd raft-consensus
    ```

3. **Building the Docker Images:**
    - Build the Docker images for the Go and Python components:
    ```bash
    docker-compose build
    ```

4. **Starting the Cluster:**
    - Start the nodes using Docker Compose:
    ```bash
    docker-compose up
    ```

5. **Viewing Logs:**
    - You can view the logs of any specific node using:
    ```bash
    docker logs -f <node-name>
    ```

6. **Stop the Cluster:**
    - To stop the cluster, use:
    ```bash
    docker-compose down
    ```

---

## Testing the System

### Test Cases:

1. **Basic Leader Election:**
    - When all nodes start as followers, one node will timeout and initiate an election, becoming the leader.
    - Verify that the election process works and the leader sends heartbeats to the followers.

2. **Leader Failure and Re-election:**
    - Simulate leader failure by stopping a node (e.g., `docker stop go-node1`).
    - Verify that a new leader is elected and the cluster continues to function.

3. **Leader Instability Under Rapid Timeouts:**
    - Simulate rapid timeouts to trigger multiple leader elections in a short time.
    - Ensure that the system correctly handles leadership transitions.

4. **New Node Joins After Election:**
    - Add a new node to the cluster after an election has occurred.
    - Verify that the new node joins and synchronizes with the current leader.

5. **Log Inconsistency Recovery:**
    - Simulate a situation where the leader has log entries that were not replicated to all followers.
    - Verify that the new leader resolves the log inconsistency by syncing the logs with followers.

### Running the Tests:
Each test can be simulated by stopping and starting specific nodes in Docker, monitoring the logs, and verifying the expected behavior. For example:
- **Leader Election Test:**
    - Stop a node to trigger the leader election.
    - Observe the logs to confirm a new leader is elected.

---

## Log Output

The log output is critical for understanding the state of the system. Common log entries include:
- **Election Start:** `Node <node-name> timed out, starting election`
- **Vote Request:** `Node <node-name> sends RPC RequestVote to Node <target-node>`
- **Log Replication:** `Node <node-name> sends RPC AppendEntries to Node <target-node>`
- **Leader Change:** `Node <node-name> becomes leader`

You can view the logs of each node to track its state during the simulation.

---

## Known Issues

1. **Log Inconsistency Under Heavy Load:**
    - In rare cases, a heavy load can cause some inconsistencies, which the system automatically handles during leader transitions.
  
2. **Leader Instability:**
    - When multiple nodes timeout simultaneously, multiple elections can occur in a very short time, potentially causing instability.

---
