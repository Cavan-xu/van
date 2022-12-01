package vnet

import "van/core/log"

type Option func(s *Server)

func WithLog(log log.ILog) Option {
	return func(s *Server) {
		s.ILog = log
	}
}

func WithConnectionMgr(connectionMgr IConnectionMgr) Option {
	return func(s *Server) {
		s.ConnectionMgr = connectionMgr
	}
}

func WithDataPack(dataPack IDataPack) Option {
	return func(s *Server) {
		s.DataPack = dataPack
	}
}

func WithMsgHandler(msgHandler IMsgHandler) Option {
	return func(s *Server) {
		s.MsgHandle = msgHandler
	}
}
