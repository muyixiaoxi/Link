package tagsignlogic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"user/service/tag/service/internal/svc"
	"user/service/tag/service/internal/types"
	"user/service/tag/service/tag"

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
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	var chooseTag types.UserTagFollow
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) (err error) {
		// 用户注册时选择标签
		err = tx.QueryRow("SELECT * FROM tb_user_tag WHERE tag_id = ? AND user_id = ?", in.TagId, in.UserId).Scan(&chooseTag)
		if err == nil {
			return status.Error(codes.AlreadyExists, "禁止重复选择标签")
		} else if err != sql.ErrNoRows {
			return err
		}
		chooseTag = types.UserTagFollow{
			TagId:  in.TagId,
			UserId: in.UserId,
		}
		_, err = tx.Exec("INSERT INTO tb_user_tag (tb_user_tag.created_at , tb_user_tag.updated_at , tag_id, user_id) VALUES (?,?,?, ?)", time.Now(), time.Now(), chooseTag.TagId, chooseTag.UserId)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &tag.UserChooseTagRequest{}, nil
}
