// Code generated by goctl. DO NOT EDIT.
// Source: tag.proto

package tagloginfront

import (
	"context"

	"Link/service/tag/tag"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AllTags                    = tag.AllTags
	AllTagsByGroupNameResponse = tag.AllTagsByGroupNameResponse
	CreateTagRequest           = tag.CreateTagRequest
	CreateTagResponse          = tag.CreateTagResponse
	DeleteTagRequest           = tag.DeleteTagRequest
	DeleteTagResponse          = tag.DeleteTagResponse
	Empty                      = tag.Empty
	GroupTag                   = tag.GroupTag
	GroupTagResponse           = tag.GroupTagResponse
	SelectAllTagsByGroupName   = tag.SelectAllTagsByGroupName
	UserChooseTagRequest       = tag.UserChooseTagRequest

	TagLoginFront interface {
	}

	defaultTagLoginFront struct {
		cli zrpc.Client
	}
)

func NewTagLoginFront(cli zrpc.Client) TagLoginFront {
	return &defaultTagLoginFront{
		cli: cli,
	}
}