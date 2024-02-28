package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"user/service/user"

	"google.golang.org/grpc/status"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/tag/service/tag"
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
	empty := user.Empty{}
	//dtm服务的etcd注册地址
	var dtmServer = "etcd://114.55.135.211:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	//创建一个自增id
	if _, err := l.svcCtx.UserRpc.AddUserId(l.ctx, &empty); err != nil {
		logx.Error(err)
		return nil, fmt.Errorf("CREATE user id error:%v", err)
	}
	userID, _ := l.svcCtx.UserRpc.NextUserID(l.ctx, &empty)

	// 创建一个saga协议
	userCreateRequest := &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
		Avatar:   req.Avatar,
		Phone:    req.Phone,
		Id:       userID.NextUserId,
	}
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).Add(tagRpcBuildServer+"/tag.TagSign/SignUserChooseTag", tagRpcBuildServer+"/tag.TagSign/SignUserChooseTagRevert", &tag.UserChooseTagRequest{
		UserId: userID.NextUserId,
		TagId:  req.StartTagId,
	}).
		Add(userRpcBuildServer+"/user.UserService/UserCreate", userRpcBuildServer+"/user.UserService/UserCreateRevertLogin", userCreateRequest)

	//事务提交
	if err := saga.Submit(); err != nil {
		//自增主键减少1
		if _, err := l.svcCtx.UserRpc.DecUserID(l.ctx, &empty); err != nil {
			logx.Error(err)
		}
		logx.Error(err)
		return nil, fmt.Errorf("saga submit error:%v", err)
	}
	return &types.UserCreateResponse{}, nil
}
