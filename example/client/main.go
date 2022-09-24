package main

import "net"

func main() {
	rAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:12310")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp4", nil, rAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn.Read([]byte{})
	}
}
