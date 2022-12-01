package vnet

import "fmt"

type IMsgHandler interface {
	Add(router IRouter)
	DoMsgHandle(req IRequest)
}

type MsgHandle struct {
	Apis map[uint32]IRouter // msgId 对应的处理函数
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]IRouter),
	}
}

func (m *MsgHandle) Add(router IRouter) {
	msgId := router.GetMsgId()
	_, ok := m.Apis[msgId]
	if ok {
		panic(fmt.Sprintf("msgId: %d has already register", msgId))
	}

	m.Apis[msgId] = router
}

func (m *MsgHandle) DoMsgHandle(req IRequest) {
	handle, ok := m.Apis[req.GetMsgId()]
	if !ok {
		req.GetServer().LogErr("msgHandle not found handle id: %d", req.GetMsgId())
		return
	}

	handle.PreHandle(req)
	handle.Handle(req)
	handle.AfterHandle(req)
}
