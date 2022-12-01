package main

import (
	"path/filepath"

	"github.com/Cavan-xu/van/core/conf"
	"github.com/Cavan-xu/van/core/log"
	"github.com/Cavan-xu/van/example/server/routers"
	"github.com/Cavan-xu/van/vnet"
)

var (
	ConfigFileName = filepath.Join("example/server/conf", "conf.json")
)

func main() {
	var (
		config = &vnet.Config{}
	)

	err := conf.LoadConfig(ConfigFileName, config)
	if err != nil {
		panic(err)
	}
	server, err := vnet.NewServer(config, vnet.WithLog(&log.Log{}))
	if err != nil {
		panic(err)
	}

	server.AddRouter(routers.NewPingRouter())
	server.Server()
}
