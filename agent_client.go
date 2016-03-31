package main

import (
	zmq "github.com/pebbe/zmq4"
	
	"fmt"
	"os"
	"net"
)

func checkError(err error){
	if 	err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)	
	}
}

func parseServerIP() net.IP {
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
	
	return addr
}

func connectZMQServer(serverIP net.IP) *zmq.Socket {
	fmt.Println("Connect to server")
	
	requester, err := zmq.NewSocket(zmq.REQ)
	checkError(err)
	
	peer_addr := "tcp://" + serverIP.String() + ":5555"
	fmt.Println("peer_addr is", peer_addr)
	
	requester.Connect(peer_addr)
	
	return requester
}

func initUDPServer() *net.UDPConn{
	udp_addr, err := net.ResolveUDPAddr("udp", ":11110")
	checkError(err)	
	
	conn, err := net.ListenUDP("udp", udp_addr)
	checkError(err)
	
	return conn
}

func recvUDPMsg(ZMQ_requester *zmq.Socket, conn *net.UDPConn) {
	var buf [20]byte
	
	for {
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return	
		}
		
		str_msg := string(buf[0:n])
		fmt.Println("recvUDP msg, msg is", str_msg)
		
		ZMQ_requester.Send(str_msg, 0)
	}
}

func main() {
	server_IP := parseServerIP()
	ZMQ_requester := connectZMQServer(server_IP)
	defer ZMQ_requester.Close()
	
	UDP_conn := initUDPServer()
	defer UDP_conn.Close()
	
	go recvUDPMsg(ZMQ_requester, UDP_conn)
/*	
	for {
		reply_msg, err := ZMQ_requester.Recv(0)
		if err != nil {
			fmt.Println("Recv:", reply_msg)	
		}
	}
	*/
    
    poller := zmq.NewPoller()
    poller.Add(ZMQ_requester, zmq.POLLIN)
    //poller.Add(UDP_conn, zmq.POLLIN)
    
    for {
    	sockets, _ := poller.Poll(-1) 
    	for _, socket := range sockets {
    		switch s := socket.Socket; s {
    		case ZMQ_requester:
    			reply_msg, err := s.Recv(0)
    			if err == nil {
    				fmt.Println("Recv:", reply_msg)	
    			}	
    		}
    	}	
    }
}