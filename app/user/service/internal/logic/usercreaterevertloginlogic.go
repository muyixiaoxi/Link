package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
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
	fmt.Println("用户标签回滚开始--->")
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		fmt.Println("注册事务走入了补偿")
		//删除插入的标签数据 和 用户数据
		_, err = tx.Exec("DELETE FROM tb_user_tag where user_id = ?", in.Id)
		_, err = tx.Exec("DELETE FROM users where id = ?", in.Id)
		//返回子事务执行失败
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("failed---->", err)
		return nil, err
	}
	fmt.Println("删除成功")
	fmt.Println("用户标签回滚结束--->")
	return &user.UserCreateResponse{}, nil
}
