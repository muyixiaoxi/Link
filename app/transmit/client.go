package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
	"transmit/common/proto"
	"transmit/types"
)

func main() {
	conn, _ := InitConnect()
	message := types.Message{
		Id:          "123",
		From:        1,
		To:          2,
		Type:        1,
		ContentType: 1,
		Time:        "123",
		Content:     "你好",
	}
	Consumer(conn)
	var ip string
	for {
		time.Sleep(2 * time.Second)
		fmt.Scan(&ip)
		if err := Producer(conn, ip, message); err != nil {
			fmt.Println("Producer(conn, ip, message) failed", err)
		}
	}
}

func InitConnect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "127.0.0.1:8333")
	fmt.Println(conn.LocalAddr())
	return
}

// Producer 生产者 发送消息
func Producer(conn net.Conn, ip string, mes types.Message) (err error) {
	transmit := types.Transmit{
		Ip:      ip,
		Message: mes,
	}
	message, _ := json.Marshal(transmit)
	m, _ := proto.Encode(string(message))
	_, err = conn.Write(m)
	if err != nil {
		// 重试三次，一次休眠一秒
		for i := 0; i < 3 && err != nil; i++ {
			time.Sleep(1 * time.Second)
			_, err = conn.Write(m)
		}
	}
	return
}

// Consumer 消费者 读消息
func Consumer(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		m, err := proto.Decode(reader)
		if err != nil {
			continue
		}
		message := types.Message{}
		json.Unmarshal([]byte(m), &message)
		fmt.Println(message)
	}
}
