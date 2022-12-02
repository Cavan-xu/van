package vnet

// 同一个玩家来自客户端和其他服务器的的请求统一先发给分发器，保证玩家的请求都是串行执行的
type ConnectionDispatch struct {
}
