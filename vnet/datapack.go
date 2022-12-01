package vnet

import "van/core/codeengine"

var (
	MessageHeadLen = 12
)

type IDataPack interface {
	GetHeadLen() int
	Pack(message IMessage) []byte
	UnPack(data []byte) (IMessage, error)
}

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
func (p *DataPack) Pack(message IMessage) []byte {
	engine := codeengine.NewCodeEngine()
	engine.EncodeUint32(message.GetMsgId())
	engine.EncodeUint32(message.GetConnId())
	engine.EncodeUint32(message.GetDataLen())
	engine.AppendBuff(message.GetData())

	return engine.GetBuff()
}

// 数据解包
func (p *DataPack) UnPack(data []byte) (IMessage, error) {
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
