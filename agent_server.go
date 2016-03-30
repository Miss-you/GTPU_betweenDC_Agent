package main

import (
	zmq "github.com/pebbe/zmq4"

	"os"
	"fmt"
	"time"
)

func main() {
	response, err := zmq.NewSocket(zmq.REP)
	defer response.Close()
	if err != nil {
		os.Exit(1)	
	}
	
	response.Bind("tcp://*:5555")
	
	for {
		msg, _ := response.Recv(0)
		fmt.Println("Recv", msg)
		
		time.Sleep(time.Second)
		
		reply := "World"
		response.Send(reply) 	
	}
}