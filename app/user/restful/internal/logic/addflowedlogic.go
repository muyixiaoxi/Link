package logic

import (
	"context"
	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFlowedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlowedLogic {
	return &AddFlowedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFlowedLogic) AddFlowed(req *types.UserAppleRequest) (err error) {
	/*	jid := l.ctx.Value("user_id").(json.Number)
		id, _ := jid.Int64()
		req.From = uint64(id)
		_, err = l.svcCtx.UserRpc.UserFlowed(l.ctx, &user.UserAddRequest{
			Id:      uint64(id),
			BeId:    req.To,
			Message: req.Content,
			Type:    uint32(req.Type),
			Remark:  req.Remark,
		})
		if err != nil {
			logx.Error("l.svcCtx.UserRpc.UserFlowed failed: ", err)
			return
		}
		if req.Type == 3 {
			// 如果该用户登录
			if client, has := Clients.Load(req.To); has {
				client.(*Client).Conn.WriteJSON(req)
				return
			}
			// 存储到kafka
			message := types.Message{
				From:    uint64(id),
				To:      req.To,
				Time:    time.Now().Format("2006/01/02 15:04:05"),
				Type:    req.Type,
				Content: req.Content,
			}
			WriteByConn(message, req.To)
			return
		}

		if req.Type == 4 {
			//添加群聊
			//判断该用户加入 创建的群聊数是否超过了最大限制
			groupCount, err := l.svcCtx.UserRpc.SelectMyGroupCount(l.ctx, &user.SelectMyGroupCountRequest{UserId: uint64(id)})
			if err != nil {
				return err
			}
			if groupCount.Count > 50 {
				return status.Error(899, "创建或者加入的群聊数据达到了最大限制")
			}
			// 获取群主id
			resp, _ := l.svcCtx.UserRpc.QueryGroupHost(l.ctx, &user.QueryGroupHostRequest{
				GroupId: req.To,
			})

			hId := resp.GroupHostId

			// 如果该用户登录
			if client, has := Clients.Load(hId); has {
				client.(*Client).Conn.WriteJSON(req)
				return err
			}
			// 存储到kafka
			message := types.Message{
				From:    uint64(id),
				To:      req.To,
				Time:    time.Now().Format("2006/01/02 15:04:05"),
				Type:    req.Type,
				Content: req.Content,
			}
			WriteByConn(message, hId)
			return err
		}*/
	return err
}
