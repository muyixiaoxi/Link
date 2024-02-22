// Code generated by goctl. DO NOT EDIT.
// Source: tag.proto

package server

import (
	"context"

	"user/service/tag/service/internal/logic/taglogin"
	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"
)

type TagLoginServer struct {
	svcCtx *svc.ServiceContext
	tag.UnimplementedTagLoginServer
}

func NewTagLoginServer(svcCtx *svc.ServiceContext) *TagLoginServer {
	return &TagLoginServer{
		svcCtx: svcCtx,
	}
}

func (s *TagLoginServer) CreateTag(ctx context.Context, in *tag.CreateTagRequest) (*tag.CreateTagResponse, error) {
	l := tagloginlogic.NewCreateTagLogic(ctx, s.svcCtx)
	return l.CreateTag(in)
}

func (s *TagLoginServer) UpdateTag(ctx context.Context, in *tag.CreateTagRequest) (*tag.CreateTagResponse, error) {
	l := tagloginlogic.NewUpdateTagLogic(ctx, s.svcCtx)
	return l.UpdateTag(in)
}

func (s *TagLoginServer) DeleteTag(ctx context.Context, in *tag.DeleteTagRequest) (*tag.DeleteTagResponse, error) {
	l := tagloginlogic.NewDeleteTagLogic(ctx, s.svcCtx)
	return l.DeleteTag(in)
}

func (s *TagLoginServer) SelectGroupTag(ctx context.Context, in *tag.Empty) (*tag.GroupTagResponse, error) {
	l := tagloginlogic.NewSelectGroupTagLogic(ctx, s.svcCtx)
	return l.SelectGroupTag(in)
}

func (s *TagLoginServer) SelectAllTagsByGroup(ctx context.Context, in *tag.SelectAllTagsByGroupName) (*tag.AllTagsByGroupNameResponse, error) {
	l := tagloginlogic.NewSelectAllTagsByGroupLogic(ctx, s.svcCtx)
	return l.SelectAllTagsByGroup(in)
}

func (s *TagLoginServer) ChooseTags(ctx context.Context, in *tag.ChooseTagsRequest) (*tag.ChooseTagsResponse, error) {
	l := tagloginlogic.NewChooseTagsLogic(ctx, s.svcCtx)
	return l.ChooseTags(in)
}

func (s *TagLoginServer) SelectLinkTags(ctx context.Context, in *tag.SelectLinkTagsRequest) (*tag.SelectLinkTagsResponse, error) {
	l := tagloginlogic.NewSelectLinkTagsLogic(ctx, s.svcCtx)
	return l.SelectLinkTags(in)
}
