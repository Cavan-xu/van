package routers

import (
	"fmt"
	"van/vnet"
)

type PingRouter struct {
	MsgId uint32
}

func NewPingRouter() *PingRouter {
	return &PingRouter{
		MsgId: 1,
	}
}

func (r *PingRouter) GetMsgId() uint32 {
	return r.MsgId
}

func (r *PingRouter) PreHandle(req vnet.IRequest) {
	fmt.Println("preHandle: data", string(req.GetMessage().GetData()))
}

func (r *PingRouter) Handle(req vnet.IRequest) {
	fmt.Println("handle")
}

func (r *PingRouter) AfterHandle(req vnet.IRequest) {
	fmt.Println("afterHandle")
}
