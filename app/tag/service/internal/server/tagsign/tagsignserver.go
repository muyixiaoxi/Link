// Code generated by goctl. DO NOT EDIT.
// Source: tag.proto

package server

import (
	"context"

	"tag/service/internal/logic/tagsign"
	"tag/service/internal/svc"
	"tag/service/tag"
)

type TagSignServer struct {
	svcCtx *svc.ServiceContext
	tag.UnimplementedTagSignServer
}

func NewTagSignServer(svcCtx *svc.ServiceContext) *TagSignServer {
	return &TagSignServer{
		svcCtx: svcCtx,
	}
}

func (s *TagSignServer) SignUserChooseTag(ctx context.Context, in *tag.UserChooseTagRequest) (*tag.UserChooseTagRequest, error) {
	l := tagsignlogic.NewSignUserChooseTagLogic(ctx, s.svcCtx)
	return l.SignUserChooseTag(in)
}

func (s *TagSignServer) SignUserChooseTagRevert(ctx context.Context, in *tag.UserChooseTagRequest) (*tag.UserChooseTagRequest, error) {
	l := tagsignlogic.NewSignUserChooseTagRevertLogic(ctx, s.svcCtx)
	return l.SignUserChooseTagRevert(in)
}
