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

func (r *PingRouter) PreHandle(req *vnet.Request) {
	fmt.Println("preHandle: data", string(req.Data))
}

func (r *PingRouter) Handle(req *vnet.Request) {
	fmt.Println("handle")
}

func (r *PingRouter) AfterHandle(req *vnet.Request) {
	fmt.Println("afterHandle")
}
