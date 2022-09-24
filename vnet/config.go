package vnet

import (
	"errors"
	"strconv"
)

const (
	MsgChanSize = 2048
)

type Config struct {
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	ServerName  string `json:"server_name"`   // Server name of the tcp Server
	Network     string `json:"network"`       // tcp、tcp4 or tcp6
	ReadBuffer  int    `json:"read_buffer"`   // tcp read buffer
	WriteBuffer int    `json:"write_buffer"`  // tcp write buffer
	MsgChanSize int    `json:"msg_chan_size"` // chan size
	Log         struct {
		FileName string `json:"file_name"` // log output file name
		LogLevel int    `json:"log_level"`
	} `json:"log"`
}

func (c *Config) check() error {
	if c.Ip == "" {
		return errors.New("ip cannot be empty")
	}
	if c.Port <= 0 {
		return errors.New("port cannot be negative")
	}
	if c.ServerName == "" {
		return errors.New("servername cannot be empty")
	}
	if c.Log.FileName == "" {
		c.Log.FileName = c.ServerName
	}
	if c.ReadBuffer < 0 {
		c.ReadBuffer = 0
	}
	if c.WriteBuffer < 0 {
		c.WriteBuffer = 0
	}
	if c.MsgChanSize == 0 {
		c.MsgChanSize = MsgChanSize
	}

	return nil
}

func (c *Config) Address() string {
	return c.Ip + ":" + strconv.Itoa(c.Port)
}
