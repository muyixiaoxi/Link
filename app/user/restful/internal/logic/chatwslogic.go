package logic

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
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
	Conn     *websocket.Conn
}

var (
	Clients = sync.Map{}
)

func NewChatWSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWSLogic {
	return &ChatWSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatWSLogic) ChatWS() error {
	// todo: add your logic here and delete this line

	return nil
}
