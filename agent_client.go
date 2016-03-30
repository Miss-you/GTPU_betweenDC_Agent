package main

import (
	zmq "github.com/pebbe/zmq4"
	
	"fmt"
	"os"
	"net"
)

func main() {
	
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(0)
	}	
	
	IPstr := os.Args[1]
	addr := net.ParseIP(IPstr)
	if addr == nil{
		fmt.Println("Invalid address")	
		os.Exit(0)
	} else {
		fmt.Println("Connect Address is", addr)	
	}
	
	fmt.Println("Connect to server")
	
	requester, err := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	if err != nil {
		os.Exit(1)	
	}
	
	peer_addr := "tcp://" + addr.String() + ":5555"
	fmt.Println("peer_addr is", peer_addr)
	
	requester.Connect(peer_addr)
	//requester.Connect("tcp://localhost:5555")
	
	for i := 1; i < 10; i++ {
		msg := fmt.Sprintf("Hello %d", i)
		fmt.Println("Send :", msg)
		requester.Send(msg, 0)
		
		//wait for reply
		reply_msg, _ := requester.Recv(0)
		fmt.Println("Recv:", reply_msg)	
	}
}