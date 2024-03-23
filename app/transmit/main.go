package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	msg := []byte{}
	for {
		m, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("与客户端", conn.LocalAddr(), "断开连接")
			return
		}
		transmit := types.Transmit{}
		json.Unmarshal([]byte(m), &transmit)
		// 读到消息后，进行转发
		c, ok := Connects.Load(transmit.Ip)
		if !ok {
			message := types.Message{
				Id:          "",
				From:        0,
				To:          0,
				Type:        100,
				ContentType: 0,
				Time:        "",
				Content:     "客户端离线",
			}
			s, _ := json.Marshal(message)
			msg, _ = proto.Encode(string(s))
			conn.Write(msg)
			fmt.Println("客户端离线")
			continue
		}
		message := transmit.Message
		j, _ := json.Marshal(message)
		msg, _ = proto.Encode(string(j))
		fmt.Println("ip:", transmit.Ip, "msg:", string(msg))
		c.(net.Conn).Write(msg)
	}
}
