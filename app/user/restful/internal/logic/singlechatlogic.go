package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
	"user/restful/internal/types"
)

// SingleChat 单聊
func (l *ChatWSLogic) SingleChat(message types.Message) {
	// 在线直接转发
	if c, has := Clients[message.To]; has {
		c.Conn.WriteJSON(message)
		return
	}
	// 离线存储消息队列
	l.WriteByConn(message)
}

// WriteByConn 基于conn发送消息
func (l *ChatWSLogic) WriteByConn(message types.Message) {
	fmt.Println("<---------------消息开始写入kafka---------------->", message.To)
	topic := fmt.Sprintf("link_user_%d", message.To)
	host := "114.55.135.211:9092"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", host, topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	jsonMessage, _ := json.Marshal(message)
	_, err := conn.WriteMessages(kafka.Message{Value: jsonMessage})
	fmt.Println("<---------------消息写入kafka结束------------------->")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
}

// ReadByConn 连接至kafka后接收消息
func (l *ChatWSLogic) ReadByConn(id uint64) (messages []*types.Message) {
	topic := fmt.Sprintf("link_user_%d", id)
	brokers := []string{"114.55.135.211:9092"}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		GroupID:   "link",
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6,
		MaxWait:   10 * time.Second,
	})

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	defer r.Close()
	for {
		fmt.Println("<--------------------kafka开始读取消息------------------->")
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		message := &types.Message{}
		json.Unmarshal(msg.Value, message)
		fmt.Println("<------------------kafka读取消息结束------------------->", message)
		messages = append(messages, message)
	}

	return
}
