package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuitGroupLogic) QuitGroup(in *user.QuitGroupRequest) (*user.QuitGroupResponse, error) {
	// 退出群聊 或者 解散群聊
	// 判断用户是否是群聊的群主 , 如果是则解散 , 不是 则退出
	tx := l.svcCtx.DB.Begin()
	var groupBossId uint64
	err := l.svcCtx.DB.Model(&types.GroupChat{}).Select("group_boss_id").Where("id = ?", in.GroupId).Find(&groupBossId).Error
	if err != nil {
		return nil, err
	}
	if groupBossId == in.UserId {
		//群主解散群聊 删除群聊 , 并且删除中间表中的记录
		err := tx.Where("id = ?", in.GroupId).Delete(&types.GroupChat{}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		//删除中间表内此群聊的所有记录
		err = tx.Unscoped().Where("group_chat_id = ?", in.GroupId).Delete(&types.UserGroupChat{}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		//普通成员退出群聊 删除中间表中的记录
		err := tx.Unscoped().Where("group_chat_id = ? AND user_id = ?", in.GroupId, in.UserId).Delete(&types.UserGroupChat{}).Error
		if err != nil {
			return nil, err
		}
	}
	return &user.QuitGroupResponse{}, tx.Commit().Error
}
