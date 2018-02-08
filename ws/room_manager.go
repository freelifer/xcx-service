package ws

import (
	"fmt"
	"time"
)

// websocket room管理者
type Room struct {
	Id        int
	Name      string
	wsManager *WsManager
	// 加入房间的客户端
	clients map[*WsClient]bool

	// 从连接器中发入的信息
	// broadcast chan []byte

	// 从连接器中注册请求
	// register chan *connection

	// 从连接器中注销请求
	// unregister chan *connection
}
