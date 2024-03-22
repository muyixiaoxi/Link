package tagsignlogic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUserChooseTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUserChooseTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUserChooseTagLogic {
	return &SignUserChooseTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUserChooseTagLogic) SignUserChooseTag(in *tag.UserChooseTagRequest) (*tag.UserChooseTagRequest, error) {
	fmt.Println("<--------------------------正常的标签逻辑------------------------------->")
	// 获取 RawDB
	// 注册账号时,选择标签
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) (err error) {
		// 用户注册时选择标签
		var exists bool
		err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM tb_user_tag WHERE tag_id = ? and user_id = ?)", in.TagId, in.UserId).Scan(&exists)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("标签重复选择")
		}
		fmt.Println("开始插入标签")
		_, err = tx.Exec("INSERT INTO tb_user_tag (tb_user_tag.created_at , tb_user_tag.updated_at , tag_id, user_id) VALUES (?,?,?, ?)", time.Now(), time.Now(), in.TagId, in.UserId)
		if err != nil {
			return fmt.Errorf("标签选择失败")
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Println("<-------------------标签插入成功-------------------------->")
	return &tag.UserChooseTagRequest{}, nil
}
