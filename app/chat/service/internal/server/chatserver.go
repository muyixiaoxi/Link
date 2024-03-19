// Code generated by goctl. DO NOT EDIT.
// Source: service.proto

package server

import (
	"context"

	"chat/service/chat"
	"chat/service/internal/logic"
	"chat/service/internal/svc"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) Ping(ctx context.Context, in *chat.Request) (*chat.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}