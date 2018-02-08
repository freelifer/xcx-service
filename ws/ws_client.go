package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

// websocket Client客户端
type WsClient struct {
	Id        int
	UserId    int
	wsManager *WsManager
	room      *Room
	conn      *websocket.Conn
}

func NewWsClient(userId int, conn *websocket.Conn, wsmgr *WsManager) *WsClient {
	//查询当前用户是否异常断开，当前正在断开重连
}

func (this *WsClient) Listen() {

}
