package main

import (
	"net"
	"time"
	"van/core/codeengine"
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

		engine := codeengine.NewCodeEngine()
		engine.EncodeInt32(10)
		engine.EncodeInt32(10)
		engine.EncodeInt32(9)
		data := []byte("hello van")
		engine.Buff = append(engine.Buff, data...)

		conn.Write(engine.GetBuff())

		time.Sleep(10 * time.Second)
	}
}
