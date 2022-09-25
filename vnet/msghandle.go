package vnet

import "fmt"

type MsgHandle struct {
	Apis map[uint32]Router // msgId 对应的处理函数
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]Router),
	}
}

func (m *MsgHandle) Add(router Router) {
	msgId := router.GetMsgId()
	_, ok := m.Apis[msgId]
	if ok {
		panic(fmt.Sprintf("msgId: %d has already register", msgId))
	}

	m.Apis[msgId] = router
}

func (m *MsgHandle) DoMsgHandle(req *Request) {
	handle, ok := m.Apis[req.GetMsgId()]
	if !ok {
		req.server.LogErr("msgHandle not found handle id: %d", req.GetMsgId())
		return
	}

	handle.PreHandle(req)
	handle.Handle(req)
	handle.AfterHandle(req)
}
