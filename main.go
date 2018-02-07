package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/build"
	"net/http"
	"text/template"
)

type hub struct {
	// 注册了的连接器
	connections map[*connection]bool

	// 从连接器中发入的信息
	broadcast chan []byte

	// 从连接器中注册请求
	register chan *connection

	// 从连接器中注销请求
	unregister chan *connection
}

var h = hub{
	broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}

type connection struct {
	// websocket 连接器
	ws *websocket.Conn

	// 发送信息的缓冲 channel
	send chan []byte
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		h.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	fmt.Println("write...")
	for message := range c.send {
		fmt.Println(string(message))
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("wsHandler")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}

var (
	addr      = flag.String("addr", ":8080", "http service address")
	assets    = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("gary.burd.info/go-websocket-chat", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func main() {
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles("home.html"))
	go h.run()
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/ws", wsHandler)
	// if err := http.ListenAndServe(*addr, nil); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		homeTempl.Execute(c.Writer, c.Request)
	})

	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})
	r.Run("localhost:8080")
}
