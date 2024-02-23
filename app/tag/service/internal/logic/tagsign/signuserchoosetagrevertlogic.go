package tagsignlogic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
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
	//reids 实现自增主键减一,如果补偿了
	key := "next_user_id"
	_, err := l.svcCtx.RDB.Decr(key)
	if err != nil {
		return nil, err
	}
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) (err error) {
		//删除记录
		_, err = tx.Exec("DELETE FROM tb_user_tag where tag_id = ? and user_id = ?", in.TagId, in.UserId)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &tag.UserChooseTagRequest{}, nil
}
