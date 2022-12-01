package vnet

type IRequest interface {
	GetMsgId() uint32
	GetConnection() IConnection
	GetMessage() IMessage
	GetServer() *Server
}

type Request struct {
	IConnection
	IMessage
}

func NewRequest(connection IConnection, message IMessage) *Request {
	return &Request{IConnection: connection, IMessage: message}
}

func (r *Request) GetMsgId() uint32 {
	return r.IMessage.GetMsgId()
}

func (r *Request) GetConnection() IConnection {
	return r.IConnection
}

func (r *Request) GetMessage() IMessage {
	return r.IMessage
}

func (r *Request) GetServer() *Server {
	return r.IConnection.GetServer()
}
