package handler

import (
	"Link/internal/response"
	"Link/restful/chat/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"

	"Link/restful/chat/internal/svc"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id   uint64
	conn *websocket.Conn
	send chan []byte
}

// 在服务端保存所有客户端连接的地方
var clients = make(map[uint64]*Client)

func WebSocketHandler(svcCtx *svc.ServiceContext) func(w http.ResponseWriter, r *http.Request) {
	// Upgrader 用于将Http连接升级为 WebSocket 连接
	upgader := websocket.Upgrader{
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			// 允许所有来源的连接
			return true
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 将http 连接升级为WebSocket 连接
		conn, err := upgader.Upgrade(w, r, nil)
		if err != nil {
			response.Response(w, nil, response.CodeWebSocketLinkFail)
			return
		}

		// 处理 WebSocket 连接的逻辑
		go handleWebSocket(conn, svcCtx)
	}
}
func handleWebSocket(conn *websocket.Conn, svcCtx *svc.ServiceContext) {
	client := &Client{}
	// 从连接中读取用户id
	if err := conn.ReadJSON(&client); err != nil {
		logc.Error(context.Background(), "conn.ReadJSON(&username) failed: ", err)
		return
	}
	fmt.Println(client.Id)
	// 将客户端添加到 clients 映射中
	client = &Client{
		Id:   client.Id,
		conn: conn,
		send: make(chan []byte),
	}
	clients[client.Id] = client
	defer func() {
		// 当函数返回时，关闭连接并从clients映射中删除客户端
		conn.Close()
		delete(clients, client.Id)
	}()
	for {
		var msg types.Message
		// 从连接中读取消息
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Printf("被玩坏了：", err)
			break
		}
		fmt.Println("信息：", msg)
	}
}
