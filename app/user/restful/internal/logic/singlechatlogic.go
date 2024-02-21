package logic

import (
	"context"
	"github.com/gorilla/websocket"
	"sync"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"user/restful/internal/svc"
)

type ChatWSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type Client struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	conn     *websocket.Conn
}

var (
	Clients     = make(map[uint64]*types.Client)
	ClientsLock sync.Mutex
)

func NewChatWSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWSLogic {
	return &ChatWSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SingleChat 单聊
func SingleChat(message types.Message) {
	// 向redis存储消息
	if c, has := Clients[message.To]; has {
		c.Conn.WriteJSON(message)
	}
}

// GroupChat 群聊
func GroupChat(message types.Message) {

}
