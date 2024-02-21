package logic

import (
	"context"
	"database/sql"
	"errors"
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
	"user/service/internal/types"
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
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 判断用户是否存在
		var users types.User
		if err := tx.QueryRow("SELECT * FROM users WHERE username = ?", in.Username).Scan(&users); err == nil {
			return errors.New("用户存在")
		}
		fmt.Println("新用户")
		// 加密密码
		pwd, _ := bcrypt.GetPwd(in.Password)
		// 插入用户数据
		stmt, err := tx.Prepare("INSERT INTO users (created_at, updated_at, username, password, avatar, phone) VALUES (?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		res, err := stmt.Exec(time.Now(), time.Now(), in.Username, pwd, in.Avatar, in.Phone)

		//返回子事务执行失败
		if err != nil {
			return dtmcli.ErrFailure
		}
		// 获取刚插入行的 ID
		lastInsertID, _ := res.LastInsertId()
		// 返回创建成功的用户信息
		pd = &user.UserCreateResponse{
			Id:       uint64(lastInsertID),
			Username: in.Username,
			Avatar:   in.Avatar,
		}
		return nil
	})
	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if err != nil {
		return nil, err
	}
	return
}
