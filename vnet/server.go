package vnet

import (
	"net"
	"sync/atomic"
	"van/core/log"
)

type Server struct {
	// 为每个连接分配的 id
	connId *int64

	*Config
	log.ILog
	ConnectionMgr IConnectionMgr
	DataPack      IDataPack
	MsgHandle     IMsgHandler
}

func NewServer(config *Config, opts ...Option) (*Server, error) {
	s := &Server{
		connId: new(int64),
		Config: config,
		//ConnectionMgr: NewConnectionMgr(),
		//DataPack:      NewDataPack(),
		//MsgHandle:     NewMsgHandle(),
	}

	for _, opt := range opts {
		opt(s)
	}
	if err := s.setUp(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) setUp() error {
	if err := s.check(); err != nil {
		return err
	}
	if s.ILog == nil {
		s.ILog = &log.Log{}
	}
	if s.ConnectionMgr == nil {
		s.ConnectionMgr = NewConnectionMgr()
	}
	if s.DataPack == nil {
		s.DataPack = NewDataPack()
	}
	if s.MsgHandle == nil {
		s.MsgHandle = NewMsgHandle()
	}

	return nil
}

func (s *Server) start() error {
	tcpAddr, err := net.ResolveTCPAddr(s.Network, s.Address())
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP(s.Network, tcpAddr)
	if err != nil {
		return err
	}

	s.LogInfo("listen tcp on: %s", s.Address())

	go func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				s.LogErr("lister accept tcp err: %v", err)
				return
			}
			s.LogInfo("receive a tcp conn from: %s", conn.RemoteAddr())
			_ = conn.SetReadBuffer(s.ReadBuffer)
			_ = conn.SetWriteBuffer(s.WriteBuffer)
			workConn := NewConnection(s.autoIncrConnId(), conn, s)
			go workConn.Start()
		}
	}()

	return nil
}

func (s *Server) autoIncrConnId() int64 {
	return atomic.AddInt64(s.connId, 1)
}

func (s *Server) Server() {
	if err := s.start(); err != nil {
		s.LogErr("server Start err: %v", err)
		return
	}

	select {}
}

func (s *Server) Stop() {
	s.LogInfo("stop Server")
}

func (s *Server) GetConnectionMgr() IConnectionMgr {
	return s.ConnectionMgr
}

func (s *Server) GetDataPack() IDataPack {
	return s.DataPack
}

func (s *Server) AddRouter(router IRouter) {
	s.MsgHandle.Add(router)
}
