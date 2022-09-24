package vnet

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Connection struct {
	// 唯一id
	id     int64
	ctx    context.Context
	cancel context.CancelFunc

	server *Server
	// tcp 的连接
	conn *net.TCPConn
	// 上次数据传输时间
	lastActiveTime time.Time
	// 是否关闭
	isClose bool
	// 消息管道，有缓冲
	msgChan chan []byte
}

func NewConnection(id int64, conn *net.TCPConn, s *Server) *Connection {
	c := &Connection{
		id:             id,
		server:         s,
		conn:           conn,
		lastActiveTime: time.Now(),
		msgChan:        make(chan []byte, s.MsgChanSize),
	}

	// 将连接加入 server 的全局连接管理器中
	c.server.GetConnectionMgr().Add(c)
	return c
}

func (c *Connection) start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.server.LogInfo("connection: %d start", c.id)
	go c.read()
	go c.write()
}

func (c *Connection) read() {
	c.server.LogInfo("%s connection start read goroutine", c.RemoteAddr().String())
	defer c.server.LogInfo("%s connection read exit", c.RemoteAddr().String())
	defer c.Stop()

	recvBuff := make([]byte, 0x10000)
	recvBytes := 0
	for {
		select {
		case <-c.ctx.Done():
			c.server.LogInfo("%s connection stop read", c.RemoteAddr().String())
			return
		default:
			n, err := c.conn.Read(recvBuff[recvBytes:])
			if err != nil {
				c.server.LogErr(err)
				return
			}
			fmt.Println(n, err)
		}
	}

}

func (c *Connection) write() {
	c.server.LogInfo("%s connection start write goroutine", c.RemoteAddr().String())
	defer c.server.LogInfo("%s connection write exit", c.RemoteAddr().String())

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.conn.Write(data); err != nil {
				c.server.LogErr(err)
				return
			}
		case <-c.ctx.Done():
			c.server.LogInfo("%s connection stop write", c.RemoteAddr().String())
			return
		}
	}
}

func (c *Connection) Stop() {
	c.cancel()
	close(c.msgChan)

	c.server.GetConnectionMgr().Delete(c.id)
}

func (c *Connection) SendMsg(data []byte) {
	c.msgChan <- data
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
