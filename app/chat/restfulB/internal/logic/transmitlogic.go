package logic

import (
	"bufio"
	"chat/common/proto"
	"chat/restfulB/internal/types"
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

func Producer(user map[string][]uint64, mes types.Message) (err error) {
	transmit := types.TransmitMap{
		Users:   user,
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
		transmit := types.TransmitMap{}
		json.Unmarshal([]byte(m), &transmit)
		// 读到消息后，进行转发
		for _, uIds := range transmit.Users {
			for _, id := range uIds {
				L.SingleChat(id, transmit.Message)
			}
		}
	}
}
