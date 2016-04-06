package main

import (
	"os"
	"fmt"
	"net"
)

//  Structure of our class
type GTPU_ICMP struct {
	//GTPU head 8 bytes
	gtpuhdr_flags 		int8
	gtpuhdr_msgtype 	int8
	gtpuhdr_length 	int16
	gtpuhdr_tunid		int32
	
	//IP head 20 bytes
	iphdr_ihl_version 	int8
	iphdr_tos			int8
	iphdr_tot_len		int16
	iphdr_id			int16
	iphdr_frag_off		int16
	iphdr_ttl			int8
	iphdr_protocol		int8
	iphdr_check			int16
	iphdr_saddr			int32
	iphdr_daddr			int32
	
	//ICMP pre
}

func checkError(err error){
	if 	err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)	
	}
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:11110")
	defer conn.Close()
	checkError(err)
	
	conn.Write([]byte("Hello world!"))	
	fmt.Println("send msg")
	
	var msg [20]byte
	conn.Read(msg[0:])
	
	fmt.Println("msg is", string(msg[0:10]))
}