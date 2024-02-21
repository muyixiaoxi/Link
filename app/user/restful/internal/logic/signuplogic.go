package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"

	"google.golang.org/grpc/status"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/tag/service/tag"
	"user/service/user"
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
	//response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
	//	Username: req.Username,
	//	Password: req.Password,
	//})

	// 获取UserRpc 的BuildTarget
	userRpcBuildServer, err := l.svcCtx.Config.UserRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "用户注册异常")
	}
	// 获取TagRpc 的BuildTarget
	tagRpcBuildServer, err := l.svcCtx.Config.TagRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "标签选择异常")
	}

	//dtm服务的etcd注册地址
	var dtmServer = "etcd://114.55.135.211:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid)
	// 第一个 Add 方法
	userCreateRequest := &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
	}
	saga.Add(userRpcBuildServer+"user.UserService/UserCreate", userRpcBuildServer+"user.UserService/UserCreateRevertLogin", userCreateRequest)
	//第二个add方法
	saga.Add(tagRpcBuildServer+"tag.TagSign/SignUserChooseTag", tagRpcBuildServer+"tag.TagSign/SignUserChooseTagRevert", &tag.UserChooseTagRequest{
		UserId: 0,
		TagId:  req.StartTagId,
	})

	//事务提交
	err = saga.Submit()

	if err != nil {
		return nil, err
	}
	//auth := l.svcCtx.Config.Auth
	//claims := jwt.UserClaims{
	//	Username: response.Username,
	//	UserID:   response.Id,
	//}
	//token, _ := jwts.GenToken(claims, auth.AccessSecret, int64(response.Id))
	//resp = &types.UserCreateResponse{
	//	Token:    token,
	//	Username: response.Username,
	//	Avatar:   response.Avatar,
	//}
	return
}
