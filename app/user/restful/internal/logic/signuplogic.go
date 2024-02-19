package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/common/jwt"

	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.UserCreateRequest) (resp *types.UserCreateResponse, err error) {
	response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserCreate failed: ", err)
		return
	}
	//注册tag_rpc服务
	//clientconf := zrpc.RpcClientConf{Etcd: discov.EtcdConf{
	//	Hosts: []string{"39.101.77.206:2379"},
	//	Key:   "tag.rpc",
	//}}
	//
	////注册之后,必须选择一个标签
	//_, err = l.svcCtx.TagRpc.SignUserChooseTag(context.Background(), &tag.UserChooseTagRequest{
	//	UserId: response.Id,
	//	TagId:  req.StartTagId,
	//})
	if err != nil {
		fmt.Println(err)
		fromErr, _ := status.FromError(err)
		if fromErr.Code() == codes.AlreadyExists {
			return nil, err
		}
		logc.Error(context.Background(), "l.svcCtx.TagRpc.SignUserChooseTag is failed", err)
		return nil, status.Error(codes.DeadlineExceeded, "标签选择失败")
	}
	auth := l.svcCtx.Config.Auth
	claims := jwt.UserClaims{
		Username: response.Username,
		UserID:   response.Id,
	}
	token, _ := jwts.GenToken(claims, auth.AccessSecret, int64(response.Id))
	resp = &types.UserCreateResponse{
		Token:    token,
		Username: response.Username,
		Avatar:   response.Avatar,
	}
	return
}
