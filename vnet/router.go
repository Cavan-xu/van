package vnet

type IRouter interface {
	GetMsgId() uint32
	PreHandle(request IRequest)
	Handle(request IRequest)
	AfterHandle(request IRequest)
}
