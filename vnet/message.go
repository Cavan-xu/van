package vnet

type IMessage interface {
	GetMsgId() uint32
	GetConnId() uint32
	GetDataLen() uint32
	GetData() []byte
	SetData(data []byte)
}

type Message struct {
	MsgId   uint32 // 消息id
	ConnId  uint32 // 连接id
	DataLen uint32 // 数据长度
	Data    []byte // 数据
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) GetMsgId() uint32 {
	return m.MsgId
}

func (m *Message) GetConnId() uint32 {
	return m.ConnId
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetData(data []byte) {
	m.Data = data
	m.DataLen = uint32(len(data))
}
