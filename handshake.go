package main

import (
	"fmt"
	"sync"
)

var serverCh chan int
var clientCh chan int

var wg sync.WaitGroup

type Node struct {
	serverSeq int
	clientSeq int
	data      []int
}

func syn(client *Node, clientCh chan int) {
	clientCh <- client.clientSeq // client sending syn
}

func syn_ack(server *Node, clientCh chan int, serverCh chan int) {
	server.clientSeq = <-clientCh // server receiving syn

	serverCh <- server.serverSeq // server sending syn

	clientCh <- server.clientSeq + 1 // server sending ack

}

func ack(client *Node, clientCh chan int, serverCh chan int) {

	serverSyn := <-serverCh // client receiving server syn

	serverAck := <-clientCh // client receiving server ack

	if serverAck == client.clientSeq+1 { // checking if ack is correct
		serverCh <- serverSyn + 1 // client sending ack
	} else {
		fmt.Println("Connection failed, incorrect sequence received from Server") // printing error if not correct
	}
}

func successOrFailure(server *Node, serverCh chan int) {
	clientAck := <-serverCh

	if server.serverSeq+1 == clientAck {
		fmt.Println("Connection established, correct sequence received from Client") // printing error if not correct
	} else {
		fmt.Println("Connection failed, incorrect sequence received from Client") // printing error if not correct
	}

}

func runClient(client *Node, clientCh chan int, serverCh chan int) {
	syn(client, clientCh)

	defer ack(client, clientCh, serverCh)
}

func runServer(server *Node, clientCh chan int, serverCh chan int) {
	syn_ack(server, clientCh, serverCh)
	defer successOrFailure(server, serverCh)
}

func main() {
	wg.Add(2)

	serverCh = make(chan int, 1)
	clientCh = make(chan int, 1)

	var server = Node{serverSeq: 30, clientSeq: 0}
	var client = Node{serverSeq: 0, clientSeq: 78}

	go func() {
		runClient(&client, clientCh, serverCh)
		defer wg.Done()
	}()

	go func() {
		runServer(&server, clientCh, serverCh)
		defer wg.Done()
	}()

	wg.Wait()
}
