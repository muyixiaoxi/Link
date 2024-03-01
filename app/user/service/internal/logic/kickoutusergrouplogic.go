package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickOutUserGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickOutUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickOutUserGroupLogic {
	return &KickOutUserGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KickOutUserGroupLogic) KickOutUserGroup(in *user.KickOutUserGroupRequest) (*user.KickOutUserGroupResponse, error) {
	// 踢出群聊
	// 先判断此人是否是群主
	var groupBossId uint64
	err := l.svcCtx.DB.Model(&types.GroupChat{}).Select("group_boss_id").Where("id = ?", in.GroupId).First(&groupBossId).Error
	if err != nil {
		return nil, err
	}
	if groupBossId != in.UserId {
		return nil, status.Error(codes.Unavailable, "无权操作")
	}
	//如果是群主,则删除中间表中的数据 群主不能踢出自己 其它人也无法踢出群主
	//去除数组中的指定的元素
	rightKickOutUserIds := make([]uint64, 0, len(in.KickOutUsersId))
	for _, userId := range in.KickOutUsersId {
		if userId != groupBossId { //在此处去除群主id
			rightKickOutUserIds = append(rightKickOutUserIds, userId)
		}
	}
	tx := l.svcCtx.DB.Begin()
	err = tx.Unscoped().Where("group_chat_id = ? and user_id in (?)", in.GroupId, rightKickOutUserIds).Delete(&types.UserGroupChat{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user.KickOutUserGroupResponse{}, tx.Commit().Error
}
