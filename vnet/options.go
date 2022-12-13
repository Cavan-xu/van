package vnet

import "github.com/Cavan-xu/van/core/vlog"

type Option func(s *Server)

func WithLog(log vlog.ILog) Option {
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
