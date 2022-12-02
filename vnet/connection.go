package vnet

import (
	"context"
	"io"
	"net"
	"time"
)

type IConnection interface {
	GetId() int64
	Start()
	Read()
	Write()
	Stop()
	SendMsg(IMessage)
	GetServer() *Server
}

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

func (c *Connection) GetId() int64 {
	return c.id
}

func (c *Connection) GetServer() *Server {
	return c.server
}

func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.server.LogInfo("connection: %d Start", c.id)
	go c.Read()
	go c.Write()
}

func (c *Connection) Read() {
	c.server.LogInfo("%s connection Start Read goroutine", c.RemoteAddr().String())
	defer c.server.LogInfo("%s connection Read exit", c.RemoteAddr().String())
	defer c.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.server.LogInfo("%s connection stop Read", c.RemoteAddr().String())
			return
		default:
			head := make([]byte, c.server.GetDataPack().GetHeadLen())
			if _, err := io.ReadFull(c.conn, head); err != nil {
				c.server.LogErr("%s connection Read data error: %v", c.RemoteAddr().String(), err)
				return
			}
			message, err := c.server.GetDataPack().UnPack(head)
			if err != nil {
				c.server.LogErr("unpack head err: %v", err)
				return
			}
			var data []byte
			if message.GetDataLen() > 0 {
				data = make([]byte, message.GetDataLen())
				if _, err := io.ReadFull(c.conn, data); err != nil {
					c.server.LogErr("%s connection Read data error: %v", c.RemoteAddr().String(), err)
					return
				}
			}
			message.SetData(data)
			go c.server.MsgHandle.DoMsgHandle(NewRequest(c, message))
		}
	}

}

func (c *Connection) Write() {
	c.server.LogInfo("%s connection Start Write goroutine", c.RemoteAddr().String())
	defer c.server.LogInfo("%s connection Write exit", c.RemoteAddr().String())

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.conn.Write(data); err != nil {
				c.server.LogErr("%s connection Write data error: %v", c.RemoteAddr().String(), err)
				return
			}
		case <-c.ctx.Done():
			c.server.LogInfo("%s connection stop Write", c.RemoteAddr().String())
			return
		}
	}
}

func (c *Connection) Stop() {
	c.cancel()
	close(c.msgChan)

	c.server.GetConnectionMgr().Delete(c.id)
}

func (c *Connection) SendMsg(message IMessage) {
	c.msgChan <- c.server.GetDataPack().Pack(message)
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
