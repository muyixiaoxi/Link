// Code generated by goctl. DO NOT EDIT.
// Source: tag.proto

package tagsign

import (
	"context"

	"tag/service/tag"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AllTags                    = tag.AllTags
	AllTagsByGroupNameResponse = tag.AllTagsByGroupNameResponse
	CheckTagCountRequest       = tag.CheckTagCountRequest
	CheckTagCountResponse      = tag.CheckTagCountResponse
	ChooseTagsRequest          = tag.ChooseTagsRequest
	ChooseTagsResponse         = tag.ChooseTagsResponse
	CreateTagRequest           = tag.CreateTagRequest
	CreateTagResponse          = tag.CreateTagResponse
	CreateTagResponse_LowTags  = tag.CreateTagResponse_LowTags
	DeleteTagRequest           = tag.DeleteTagRequest
	DeleteTagResponse          = tag.DeleteTagResponse
	Empty                      = tag.Empty
	GroupTag                   = tag.GroupTag
	GroupTagResponse           = tag.GroupTagResponse
	SelectAllTagsByGroupName   = tag.SelectAllTagsByGroupName
	SelectLinkTag              = tag.SelectLinkTag
	SelectLinkTagsRequest      = tag.SelectLinkTagsRequest
	SelectLinkTagsResponse     = tag.SelectLinkTagsResponse
	SelectMyTagsRequest        = tag.SelectMyTagsRequest
	UserChooseTagRequest       = tag.UserChooseTagRequest

	TagSign interface {
		SignUserChooseTag(ctx context.Context, in *UserChooseTagRequest, opts ...grpc.CallOption) (*UserChooseTagRequest, error)
		SignUserChooseTagRevert(ctx context.Context, in *UserChooseTagRequest, opts ...grpc.CallOption) (*UserChooseTagRequest, error)
	}

	defaultTagSign struct {
		cli zrpc.Client
	}
)

func NewTagSign(cli zrpc.Client) TagSign {
	return &defaultTagSign{
		cli: cli,
	}
}

func (m *defaultTagSign) SignUserChooseTag(ctx context.Context, in *UserChooseTagRequest, opts ...grpc.CallOption) (*UserChooseTagRequest, error) {
	client := tag.NewTagSignClient(m.cli.Conn())
	return client.SignUserChooseTag(ctx, in, opts...)
}

func (m *defaultTagSign) SignUserChooseTagRevert(ctx context.Context, in *UserChooseTagRequest, opts ...grpc.CallOption) (*UserChooseTagRequest, error) {
	client := tag.NewTagSignClient(m.cli.Conn())
	return client.SignUserChooseTagRevert(ctx, in, opts...)
}
