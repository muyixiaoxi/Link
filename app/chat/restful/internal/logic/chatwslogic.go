package logic

import (
	"chat/restful/internal/svc"
	"chat/restful/internal/types"
	"context"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (l *ChatWSLogic) SendMessage(message types.Message) (err error) {
	if message.Type == 1 {
		err = l.SingleChat(message)
	} else if message.Type == 2 {
		//群聊
		err = l.GroupChat(message)
		fromErr, _ := status.FromError(err)
		if fromErr.Code() != codes.FailedPrecondition {
			err = nil
		}
	} else if message.Type == 100 { // 集群转发失败
		err = l.SaveMessage(message, false)
	}
	return
}
