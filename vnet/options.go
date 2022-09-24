package vnet

type Option func(s *Server)

func WithIp(ip string) Option {
	return func(s *Server) {
		s.Config.Ip = ip
	}
}
