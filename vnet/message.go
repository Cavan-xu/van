package vnet

type Message struct {
	MsgId   uint32 // 消息id
	ConnId  uint32 // 连接id
	DataLen uint32 // 数据长度
	Data    []byte // 数据
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
