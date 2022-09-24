package main

import (
	"path/filepath"
	"van/core/conf"
	"van/vnet"
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
	server, err := vnet.NewServer(config)
	if err != nil {
		panic(err)
	}

	server.Server()
}
