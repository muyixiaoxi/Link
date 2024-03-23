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

func client() {
	conn, _ := InitConnect()
	go Consumer(conn)
	var ip string
	message := types.Message{
		Id:          "123",
		From:        1,
		To:          2,
		Type:        1,
		ContentType: 1,
		Time:        "123",
		Content:     "你好",
	}
	for {
		fmt.Scan(&ip)
		users := map[string][]uint64{}
		users[ip] = []uint64{1, 2}
		time.Sleep(2 * time.Second)
		if err := Producer(conn, users, message); err != nil {
			fmt.Println("Producer(conn, ip, message) failed", err)
		}
	}

}

func InitConnect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "127.0.0.1:8333")
	fmt.Println(conn.LocalAddr())
	return
}

func Producer(conn net.Conn, user map[string][]uint64, mes types.Message) (err error) {
	transmit := types.TransmitMap{
		Users:   user,
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
		transmit := types.TransmitMap{}
		json.Unmarshal([]byte(m), &transmit)
		// 读到消息后，进行转发
		for _, uIds := range transmit.Users {
			for _, id := range uIds {
				fmt.Println(id, transmit.Message)
			}
		}
	}
}
