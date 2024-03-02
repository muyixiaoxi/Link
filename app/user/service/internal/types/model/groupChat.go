package model

import "github.com/gorilla/websocket"

// Message 结构体用于表示消息内容
type Message struct {
	Content interface{} `json:"content"` // 消息内容，可以是文字、图片链接等
	Type    string      `json:"type"`    // 消息类型，可以是 text、image、emoji 等
}

// Client 连接到服务器的客户端
type Client struct {
	ID   string `json:"id"`
	Conn *websocket.Conn
}

// Hub 用于管理客户端连接和广播消息
type Hub struct {
	clients    map[string]*Client //所有连接到服务器的客户端
	broadcast  chan []byte        //广播消息通道
	register   chan *Client       //注册新客户端的通道
	unregister chan string        //注销客户端的通道
}

// NewHub 初始化
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan string),
	}
}

// Run 启动Hub实例的运行循环
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: //注册新的客户端
			h.clients[client.ID] = client
		case clientID := <-h.unregister: //注销客户端
			if _, ok := h.clients[clientID]; ok {
				delete(h.clients, clientID)
				h.clients[clientID].Conn.Close() //关闭websocket连接
			}
		case message := <-h.broadcast: //广播消息给所有客户端
			for clientID, client := range h.clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					// 如果发送消息失败，关闭 WebSocket 连接并从列表中删除客户端
					client.Conn.Close()
					delete(h.clients, clientID)
				}
			}
		}
	}
}
