package main

import (
	"net"
	"time"

	"github.com/Cavan-xu/van/vnet"
)

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
		message := &vnet.Message{
			MsgId:   1,
			ConnId:  1,
			DataLen: 9,
			Data:    []byte("hello van"),
		}
		pack := vnet.NewDataPack()
		data := pack.Pack(message)
		conn.Write(data)
		time.Sleep(10 * time.Second)
	}
}
