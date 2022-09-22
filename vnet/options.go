package vnet

type option func(s *server)

func WithIp(ip string) option {
	return func(s *server) {
		s.Config.Ip = ip
	}
}
