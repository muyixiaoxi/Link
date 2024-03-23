package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"sync"
	"transmit/common/proto"
	"transmit/types"
)

var Connects sync.Map

func main() {
	listenClient()
}

// listenClient 监听
func listenClient() {
	lister, err := net.Listen("tcp", "127.0.0.1:8333")
	if err != nil {
		fmt.Println("net.Listen failed:", err)
	}
	for {
		conn, err := lister.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String())
		Connects.Swap(conn.RemoteAddr().String(), conn)
		go addReceiver(conn)
	}
}

// addReceiver 向连接添加接收器
func addReceiver(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		m, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("与客户端", conn.LocalAddr(), "断开连接")
			return
		}
		transmit := types.TransmitMap{}
		json.Unmarshal([]byte(m), &transmit)
		// 读到消息后，根据服务器进行转发
		for connect := range transmit.Users {
			transmitMessage(conn, connect, transmit)
		}
	}
}

func transmitMessage(conn net.Conn, ip string, transmit types.TransmitMap) {
	c, ok := Connects.Load(ip)
	message := types.TransmitMap{
		Users: map[string][]uint64{},
	}
	if !ok {
		message = types.TransmitMap{
			Message: types.Message{
				Id:          "",
				From:        0,
				To:          0,
				Type:        100,
				ContentType: 0,
				Time:        "",
				Content:     "客户端离线",
			},
		}
		s, _ := json.Marshal(message)
		msg, _ := proto.Encode(string(s))
		conn.Write(msg)
		fmt.Println("客户端离线：", ip)
		logx.Error("connect ip offline:", ip)
		return
	}
	message.Users[ip] = transmit.Users[ip]
	message.Message = transmit.Message
	j, _ := json.Marshal(message)
	msg, _ := proto.Encode(string(j))
	fmt.Println("ip:", ip, "msg:", string(msg))
	c.(net.Conn).Write(msg)
}
