package vnet

type Router interface {
	GetMsgId() uint32
	PreHandle(request *Request)
	Handle(request *Request)
	AfterHandle(request *Request)
}
