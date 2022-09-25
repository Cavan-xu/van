package vnet

import "van/core/codeengine"

var (
	MessageHeadLen = 12
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 消息头的长度：MsgId uint32 + ConnId uint32 + DataLen uint32
func (p *DataPack) GetHeadLen() int {
	return MessageHeadLen
}

// 数据装包
func (p *DataPack) Pack(message *Message) []byte {
	engine := codeengine.NewCodeEngine()
	engine.EncodeUint32(message.MsgId)
	engine.EncodeUint32(message.ConnId)
	engine.EncodeUint32(message.DataLen)
	engine.AppendBuff(message.Data)

	return engine.GetBuff()
}

// 数据解包
func (p *DataPack) UnPack(data []byte) (*Message, error) {
	message := NewMessage()
	engine := codeengine.NewCodeEngine()
	engine.SetBuff(data)

	var (
		pos int
		err error
	)

	message.MsgId, pos, err = engine.DecodeUint32(0)
	if err != nil {
		return nil, err
	}
	message.ConnId, pos, err = engine.DecodeUint32(pos)
	if err != nil {
		return nil, err
	}
	message.DataLen, pos, err = engine.DecodeUint32(pos)
	if err != nil {
		return nil, err
	}

	return message, nil
}
