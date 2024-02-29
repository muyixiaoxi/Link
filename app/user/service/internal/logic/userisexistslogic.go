package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIsExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIsExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIsExistsLogic {
	return &UserIsExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserIsExistsLogic) UserIsExists(in *user.UserCreateRequest) (*user.Empty, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	// 判断用户是否存在
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE phone = ?)", in.Phone).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "用户已存在")
	}
	return &user.Empty{}, nil
}
