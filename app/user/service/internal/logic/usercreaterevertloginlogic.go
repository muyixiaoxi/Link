package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"time"
	"user/common/bcrypt"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateRevertLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateRevertLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateRevertLoginLogic {
	return &UserCreateRevertLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateRevertLoginLogic) UserCreateRevertLogin(in *user.UserCreateRequest) (pd *user.UserCreateResponse, err error) {
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	// 判断用户是否存在
	var users types.User
	if err := db.QueryRow("SELECT * FROM users WHERE username = ?", in.Username).Scan(&users); err == nil {
		return nil, errors.New("用户存在")
	}
	// 加密密码
	pwd, _ := bcrypt.GetPwd(in.Password)
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 插入用户数据
		stmt, err := tx.Prepare("INSERT INTO users (created_at, updated_at, username, password, avatar, phone) VALUES (?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		res, err := stmt.Exec(time.Now(), time.Now(), in.Username, pwd, in.Avatar, in.Phone)
		//返回子事务执行失败
		if err != nil {
			return err
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
	return
}
