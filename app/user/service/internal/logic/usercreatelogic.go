package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"user/common/bcrypt"
	"user/service/internal/svc"
	"user/service/user"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateResponse, endErr error) {
	// 获取 RawDB
	// 注册
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 加密密码
		pwd, _ := bcrypt.GetPwd(in.Password)
		// 插入用户数据
		_, err = tx.Exec("INSERT INTO users (id , created_at, updated_at, username, password, avatar, phone) VALUES (?,?, ?, ?, ?, ?, ?)", in.Id, time.Now(), time.Now(), in.Username, pwd, in.Avatar, in.Phone)
		//返回子事务执行失败
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	return &user.UserCreateResponse{}, endErr
}
