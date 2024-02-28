package logic

import (
	"context"
	"database/sql"
	"fmt"
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

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateResponse, err error) {
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
		// 判断用户是否存在
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", in.Username).Scan(&exists)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("用户已经存在")
		}
		// 加密密码
		pwd, _ := bcrypt.GetPwd(in.Password)
		// 插入用户数据
		stmt, err := tx.Prepare("INSERT INTO users (id , created_at, updated_at, username, password, avatar, phone) VALUES (?,?, ?, ?, ?, ?, ?)")
		if err != nil {
			//此处如果发生错误,不再重试,直接回滚
			return dtmcli.ErrFailure
		}
		defer stmt.Close()
		_, err = stmt.Exec(in.Id, time.Now(), time.Now(), in.Username, pwd, in.Avatar, in.Phone)
		//返回子事务执行失败
		if err != nil {
			return err
		}
		return nil
	})
	// 这种情况是库存不足，不再重试，走回滚
	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return
}
