## van是什么？

一款 **go** 语言开发的轻量级框架，适合用于游戏服务器开发
具有以下特点：

> 基于 **TCP** 协议稳定传输数据，
> 每个 **socket** 连接创建工作协程，无需考虑消息的并发问题，
> 用户自定义消息路由，并支持编写路由前、路由中、路由后钩子函数，
> 框架集成连接管理器、数据装箱拆箱、路由管理、日志管理，同时也支持用户自定义

## 开箱即用

**客户端示例**

````
package main

import (
	"net"
	"time"

	"github.com/Cavan-xu/van/vnet"
)

func main() {
	rAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:12310")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp4", nil, rAddr)
	if err != nil {
		panic(err)
	}

	for {
		message := &vnet.Message{
			MsgId:   1,
			ConnId:  1,
			DataLen: 9,
			Data:    []byte("hello van"),
		}
		pack := vnet.NewDataPack()
		data := pack.Pack(message)
		conn.Write(data)
		time.Sleep(10 * time.Second)
	}
}
````

**服务器示例**

自定义router

```
package routers

import (
	"fmt"

	"github.com/Cavan-xu/van/vnet"
)

type PingRouter struct {
	MsgId uint32
}

func NewPingRouter() *PingRouter {
	return &PingRouter{
		MsgId: 1,
	}
}

func (r *PingRouter) GetMsgId() uint32 {
	return r.MsgId
}

func (r *PingRouter) PreHandle(req vnet.IRequest) {
	fmt.Println("preHandle: data", string(req.GetMessage().GetData()))
}

func (r *PingRouter) Handle(req vnet.IRequest) {
	fmt.Println("handle")
}

func (r *PingRouter) AfterHandle(req vnet.IRequest) {
	fmt.Println("afterHandle")
}
```

服务启动

```
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
```

