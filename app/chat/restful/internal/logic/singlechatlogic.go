package logic

import (
	"chat/restful/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

// SingleChat 单聊
func (l *ChatWSLogic) SingleChat(message types.Message) (err error) {
	// 在线存储+转发
	if c, has := Clients.Load(message.To); has {
		// 如果存储失败返回
		if err = l.SaveMessage(message, true); err != nil {
			return
		}
		// 如果转发失败，当作离线存储
		if err := c.(*Client).Conn.WriteJSON(message); err != nil {
			Clients.Delete(message.To)
			// 不管离线存储成功与否都返回
			err = l.SaveMessage(message, false)
			return
		}
		return
	}
	// 离线存储消息队列
	err = l.SaveMessage(message, false)
	return
}

// WriteByConn 基于conn发送消息
func WriteByConn(message types.Message, userId uint64) {
	topic := fmt.Sprintf("link_user_%d", userId)

	host := "114.55.135.211:9092"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", host, topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	jsonMessage, _ := json.Marshal(message)
	_, err := conn.WriteMessages(kafka.Message{Value: jsonMessage})

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
}

// ReadByConn 连接至kafka后接收消息
func ReadByConn(id uint64) (messages []*types.Message) {
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

		msg, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		message := &types.Message{}
		json.Unmarshal(msg.Value, message)

		messages = append(messages, message)
	}

	return
}
