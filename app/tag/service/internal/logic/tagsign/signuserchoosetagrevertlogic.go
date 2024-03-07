package tagsignlogic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUserChooseTagRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUserChooseTagRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUserChooseTagRevertLogic {
	return &SignUserChooseTagRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUserChooseTagRevertLogic) SignUserChooseTagRevert(in *tag.UserChooseTagRequest) (*tag.UserChooseTagRequest, error) {
	fmt.Println("用户标签SignUserChooseTagRevert--->开始")
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) (err error) {
		fmt.Println("注册时选择标签进入了补偿")
		logc.Info(l.ctx)
		//删除记录
		_, err = tx.Exec("DELETE FROM tb_user_tag where tag_id = ? and user_id = ?", in.TagId, in.UserId)
		return err
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("用户标签SignUserChooseTagRevert--->结束")

	return &tag.UserChooseTagRequest{}, nil
}
