package vnet

type Request struct {
	*Connection
	*Message
}

func NewRequest(connection *Connection, message *Message) *Request {
	return &Request{Connection: connection, Message: message}
}

func (r *Request) GetMsgId() uint32 {
	return r.MsgId
}
