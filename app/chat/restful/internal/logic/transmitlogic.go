package logic

import (
	"bufio"
	"chat/common/proto"
	"chat/restful/internal/types"
	"encoding/json"
	"net"
	"time"
)

var L *ChatWSLogic

var Conn net.Conn

func InitConnect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "127.0.0.1:8333")
	return
}

// Producer 生产者 发送消息
func Producer(ip string, mes types.Message) (err error) {
	transmit := types.Transmit{
		Ip:      ip,
		Message: mes,
	}
	message, _ := json.Marshal(transmit)
	m, _ := proto.Encode(string(message))
	_, err = Conn.Write(m)
	if err != nil {
		// 重试三次，一次休眠一秒
		for i := 0; i < 3 && err != nil; i++ {
			time.Sleep(1 * time.Second)
			_, err = Conn.Write(m)
		}
	}
	return
}

// Consumer 消费者 读消息
func Consumer() {
	defer Conn.Close()
	reader := bufio.NewReader(Conn)
	for {
		m, err := proto.Decode(reader)
		if err != nil {
			continue
		}
		message := types.Message{}
		json.Unmarshal([]byte(m), &message)
		// 读到消息后，进行转发
		L.SendMessage(message)
	}
}
