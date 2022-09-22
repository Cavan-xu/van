package vnet

import (
	"net"
	"van/core/log"
)

type server struct {
	*Config
	*log.Log
}

func NewServer(config *Config, opts ...option) (*server, error) {
	s := &server{
		Config: config,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

func (s *server) setUp() error {
	if err := s.check(); err != nil {
		return err
	}

	return nil
}

func (s *server) start() error {
	tcpAddr, err := net.ResolveTCPAddr(s.Network, s.Address())
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP(s.Network, tcpAddr)
	if err != nil {
		return err
	}

	s.LogInfo("listen tpc on: %s", s.Address())

	go func() {
		conn, err := listener.AcceptTCP()
		if err != nil {
			s.LogErr(err)
			return
		}
		s.LogInfo("receive a tcp conn from: %s", conn.RemoteAddr())
		_ = conn.SetReadBuffer(s.ReadBuffer)
		_ = conn.SetWriteBuffer(s.WriteBuffer)
	}()

	return nil
}

func (s *server) Server() {
	if err := s.start(); err != nil {
		s.LogErr(err)
		return
	}

	select {}
}

func (s *server) Stop() {
	s.LogInfo("stop Server")
}
