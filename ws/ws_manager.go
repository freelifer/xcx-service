package ws

import (
	"fmt"
	"time"
)

// websocket 全局管理者
type WsManager struct {
	// 注册了的连接器
	// connections map[*connection]bool

	// 从连接器中发入的信息
	// broadcast chan []byte

	// 从连接器中注册请求
	// register chan *connection

	// 从连接器中注销请求
	// unregister chan *connection
}

func StartTime() {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
	}
}
