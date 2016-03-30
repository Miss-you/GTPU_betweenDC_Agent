package main

import (
	zmq "github.com/pebbe/zmq4"
	
	"fmt"
	"os"
)

func main() {
	fmt.Println("Connect to server")
	
	requester, err := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	if err != nil {
		os.Exit(1)	
	}
	
	requester.Connect("tcp://localhost:5555")
	
	for i := 1; i < 10; ++i {
		msg := fmt.Sprintf("Hello %d", i)
		fmt.Println("Send :", msg)
		requester.Send(msg)
		
		//wait for reply
		reply_msg, _ := requester.Recv(0)
		fmt.Println("Recv:", reply_msg)	
	}
}