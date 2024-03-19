package logic

import (
	"chat/restful/internal/svc"
	"context"
	"github.com/gorilla/websocket"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatWSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatWSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWSLogic {
	return &ChatWSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Client struct {
	Id   uint64 `json:"id"`
	Conn *websocket.Conn
}

var (
	Clients = sync.Map{}
)

func (l *ChatWSLogic) ChatWS() error {
	// todo: add your logic here and delete this line

	return nil
}
