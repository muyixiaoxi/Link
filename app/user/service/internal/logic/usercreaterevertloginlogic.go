package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"user/common/bcrypt"
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
	// 在UserCreate方法中,插入了一条记录 , 在这里给它删除
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	// 加密密码
	pwd, _ := bcrypt.GetPwd(in.Password)
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		//删除用户数据
		stmt, err := tx.Prepare("DELETE FROM users where phone = ? and password = ?")
		//返回事务执行失败
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(in.Phone, pwd)
		//返回子事务执行失败
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
