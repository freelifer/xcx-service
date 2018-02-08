package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// websocket 全局管理者
type WsManager struct {
	pattern string

	rooms map[int]*Room

	clients map[int]*WsClient

	addCh  chan *WsClient
	delCh  chan *WsClient
	doneCh chan bool
	errCh  chan error
	// 注册了的连接器
	// connections map[*connection]bool

	// 从连接器中发入的信息
	// broadcast chan []byte

	// 从连接器中注册请求
	// register chan *connection

	// 从连接器中注销请求
	// unregister chan *connection
}

type WsService interface {
	Add(c *WsClient)
	Del(c *WsClient)
	Create() *Room
	Destroy()
	Insert()
	Send()
	Err(err error)
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func NewWsManager(pattern string) *WsManager {
	return &WsManager{
		pattern,
	}
}

func (this *WsManager) Listen(r *gin.Engine) {
	log.Println("Websocket Listening server...")

	// websocket handler
	onConnected := func(w http.ResponseWriter, r *http.Request) {
		// token 验证

		// wesocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				this.errCh <- err
			}
		}()

		client := NewWsClient(conn, this)
		this.Add(client)
		client.Listen()
	}
	r.GET(this.pattern, func(c *gin.Context) {
		onConnected(c.Writer, c.Request)
	})
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-this.addCh:

		// del a client
		case c := <-this.delCh:

		// broadcast message for all clients
		case msg := <-this.sendAllCh:

		case err := <-this.errCh:
			log.Println("Error:", err.Error())

		case <-this.doneCh:
			return
		}
	}
}

func (this *WsManager) Add(c *WsClient) {

}

func (this *WsManager) Err(err *error) {
	if err != nil {
		this.errCh <- err
	}
}

func StartTime() {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
	}
}
