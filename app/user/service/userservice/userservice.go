// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"user/service/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DetailGroupRequest            = user.DetailGroupRequest
	DetailGroupResponse           = user.DetailGroupResponse
	DisposeFlowedRequest          = user.DisposeFlowedRequest
	Empty                         = user.Empty
	GetOffsetRequest              = user.GetOffsetRequest
	GetOffsetResponse             = user.GetOffsetResponse
	GroupChat                     = user.GroupChat
	GroupInformation              = user.GroupInformation
	KickOutUserGroupRequest       = user.KickOutUserGroupRequest
	KickOutUserGroupResponse      = user.KickOutUserGroupResponse
	MyGroupResponse               = user.MyGroupResponse
	NextUserIDResponse            = user.NextUserIDResponse
	QueryMyGroupListRequest       = user.QueryMyGroupListRequest
	QuitGroupRequest              = user.QuitGroupRequest
	QuitGroupResponse             = user.QuitGroupResponse
	RecommendUser                 = user.RecommendUser
	RecommendUsersRequest         = user.RecommendUsersRequest
	RecommendUsersResponse        = user.RecommendUsersResponse
	SearchMyGroupByNameRequest    = user.SearchMyGroupByNameRequest
	SearchStrangerGroupRequest    = user.SearchStrangerGroupRequest
	SelectUserListByGroup         = user.SelectUserListByGroup
	SelectUserListByGroupRequest  = user.SelectUserListByGroupRequest
	SelectUserListByGroupResponse = user.SelectUserListByGroupResponse
	SetOffsetRequest              = user.SetOffsetRequest
	UpdateGroupInfoRequest        = user.UpdateGroupInfoRequest
	UpdateGroupInfoResponse       = user.UpdateGroupInfoResponse
	UpdateGroupRemarkRequest      = user.UpdateGroupRemarkRequest
	UpdateGroupRemarkResponse     = user.UpdateGroupRemarkResponse
	UserAddRequest                = user.UserAddRequest
	UserCreateGroupRequest        = user.UserCreateGroupRequest
	UserCreateRequest             = user.UserCreateRequest
	UserCreateResponse            = user.UserCreateResponse
	UserDeleteFriendRequest       = user.UserDeleteFriendRequest
	UserFriend                    = user.UserFriend
	UserFriendRequest             = user.UserFriendRequest
	UserFriendResponse            = user.UserFriendResponse
	UserInfoRequest               = user.UserInfoRequest
	UserInfoResponse              = user.UserInfoResponse
	UserLoginRequest              = user.UserLoginRequest
	UserLoginResponse             = user.UserLoginResponse
	UserQueryFriendRequest        = user.UserQueryFriendRequest
	UserQueryPhoneRequest         = user.UserQueryPhoneRequest
	UserQueryPhoneResponse        = user.UserQueryPhoneResponse
	UserSelectGroupsRequest       = user.UserSelectGroupsRequest
	UserSelectGroupsResponse      = user.UserSelectGroupsResponse
	UserSelfGroupRequest          = user.UserSelfGroupRequest
	UserUpdateInfoRequest         = user.UserUpdateInfoRequest
	UserUpdateRemarkRequest       = user.UserUpdateRemarkRequest

	UserService interface {
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserCreateRevertLogin(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserIsExists(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Empty, error)
		UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UserUpdateInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*Empty, error)
		// 好友
		UserFlowed(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*Empty, error)
		UserDisposeFlowed(ctx context.Context, in *DisposeFlowedRequest, opts ...grpc.CallOption) (*Empty, error)
		UserUpdateRemark(ctx context.Context, in *UserUpdateRemarkRequest, opts ...grpc.CallOption) (*Empty, error)
		RecommendUsers(ctx context.Context, in *RecommendUsersRequest, opts ...grpc.CallOption) (*RecommendUsersResponse, error)
		UserFriendList(ctx context.Context, in *UserFriendRequest, opts ...grpc.CallOption) (*UserFriendResponse, error)
		UserQueryFriend(ctx context.Context, in *UserQueryFriendRequest, opts ...grpc.CallOption) (*UserFriendResponse, error)
		UserQueryPhone(ctx context.Context, in *UserQueryPhoneRequest, opts ...grpc.CallOption) (*UserQueryPhoneResponse, error)
		UserDeleteFriend(ctx context.Context, in *UserDeleteFriendRequest, opts ...grpc.CallOption) (*Empty, error)
		// 偏移量
		GetOffset(ctx context.Context, in *GetOffsetRequest, opts ...grpc.CallOption) (*GetOffsetResponse, error)
		SetOffset(ctx context.Context, in *SetOffsetRequest, opts ...grpc.CallOption) (*Empty, error)
		// redis自增id
		NextUserID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NextUserIDResponse, error)
		AddUserId(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
		DecUserID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
		// 群聊
		UserCreateGroup(ctx context.Context, in *UserCreateGroupRequest, opts ...grpc.CallOption) (*Empty, error)
		UserSelectGroup(ctx context.Context, in *UserSelectGroupsRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error)
		UserUserSelfGroup(ctx context.Context, in *UserSelfGroupRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error)
		UserSelectDetailGroup(ctx context.Context, in *DetailGroupRequest, opts ...grpc.CallOption) (*DetailGroupResponse, error)
		QuitGroup(ctx context.Context, in *QuitGroupRequest, opts ...grpc.CallOption) (*QuitGroupResponse, error)
		UserListByGroup(ctx context.Context, in *SelectUserListByGroupRequest, opts ...grpc.CallOption) (*SelectUserListByGroupResponse, error)
		KickOutUserGroup(ctx context.Context, in *KickOutUserGroupRequest, opts ...grpc.CallOption) (*KickOutUserGroupResponse, error)
		UpdateGroupInformation(ctx context.Context, in *UpdateGroupInfoRequest, opts ...grpc.CallOption) (*UpdateGroupInfoResponse, error)
		UpdateGroupRemark(ctx context.Context, in *UpdateGroupRemarkRequest, opts ...grpc.CallOption) (*UpdateGroupRemarkResponse, error)
		QueryMyGroupList(ctx context.Context, in *QueryMyGroupListRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error)
		SearchStrangerGroup(ctx context.Context, in *SearchStrangerGroupRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error)
		SearchMyGroupByName(ctx context.Context, in *SearchMyGroupByNameRequest, opts ...grpc.CallOption) (*MyGroupResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}

func (m *defaultUserService) UserCreateRevertLogin(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserCreateRevertLogin(ctx, in, opts...)
}

func (m *defaultUserService) UserIsExists(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserIsExists(ctx, in, opts...)
}

func (m *defaultUserService) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (m *defaultUserService) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUserService) UserUpdateInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserUpdateInfo(ctx, in, opts...)
}

// 好友
func (m *defaultUserService) UserFlowed(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserFlowed(ctx, in, opts...)
}

func (m *defaultUserService) UserDisposeFlowed(ctx context.Context, in *DisposeFlowedRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserDisposeFlowed(ctx, in, opts...)
}

func (m *defaultUserService) UserUpdateRemark(ctx context.Context, in *UserUpdateRemarkRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserUpdateRemark(ctx, in, opts...)
}

func (m *defaultUserService) RecommendUsers(ctx context.Context, in *RecommendUsersRequest, opts ...grpc.CallOption) (*RecommendUsersResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.RecommendUsers(ctx, in, opts...)
}

func (m *defaultUserService) UserFriendList(ctx context.Context, in *UserFriendRequest, opts ...grpc.CallOption) (*UserFriendResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserFriendList(ctx, in, opts...)
}

func (m *defaultUserService) UserQueryFriend(ctx context.Context, in *UserQueryFriendRequest, opts ...grpc.CallOption) (*UserFriendResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserQueryFriend(ctx, in, opts...)
}

func (m *defaultUserService) UserQueryPhone(ctx context.Context, in *UserQueryPhoneRequest, opts ...grpc.CallOption) (*UserQueryPhoneResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserQueryPhone(ctx, in, opts...)
}

func (m *defaultUserService) UserDeleteFriend(ctx context.Context, in *UserDeleteFriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserDeleteFriend(ctx, in, opts...)
}

// 偏移量
func (m *defaultUserService) GetOffset(ctx context.Context, in *GetOffsetRequest, opts ...grpc.CallOption) (*GetOffsetResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetOffset(ctx, in, opts...)
}

func (m *defaultUserService) SetOffset(ctx context.Context, in *SetOffsetRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.SetOffset(ctx, in, opts...)
}

// redis自增id
func (m *defaultUserService) NextUserID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NextUserIDResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.NextUserID(ctx, in, opts...)
}

func (m *defaultUserService) AddUserId(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.AddUserId(ctx, in, opts...)
}

func (m *defaultUserService) DecUserID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.DecUserID(ctx, in, opts...)
}

// 群聊
func (m *defaultUserService) UserCreateGroup(ctx context.Context, in *UserCreateGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserCreateGroup(ctx, in, opts...)
}

func (m *defaultUserService) UserSelectGroup(ctx context.Context, in *UserSelectGroupsRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserSelectGroup(ctx, in, opts...)
}

func (m *defaultUserService) UserUserSelfGroup(ctx context.Context, in *UserSelfGroupRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserUserSelfGroup(ctx, in, opts...)
}

func (m *defaultUserService) UserSelectDetailGroup(ctx context.Context, in *DetailGroupRequest, opts ...grpc.CallOption) (*DetailGroupResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserSelectDetailGroup(ctx, in, opts...)
}

func (m *defaultUserService) QuitGroup(ctx context.Context, in *QuitGroupRequest, opts ...grpc.CallOption) (*QuitGroupResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.QuitGroup(ctx, in, opts...)
}

func (m *defaultUserService) UserListByGroup(ctx context.Context, in *SelectUserListByGroupRequest, opts ...grpc.CallOption) (*SelectUserListByGroupResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserListByGroup(ctx, in, opts...)
}

func (m *defaultUserService) KickOutUserGroup(ctx context.Context, in *KickOutUserGroupRequest, opts ...grpc.CallOption) (*KickOutUserGroupResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.KickOutUserGroup(ctx, in, opts...)
}

func (m *defaultUserService) UpdateGroupInformation(ctx context.Context, in *UpdateGroupInfoRequest, opts ...grpc.CallOption) (*UpdateGroupInfoResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UpdateGroupInformation(ctx, in, opts...)
}

func (m *defaultUserService) UpdateGroupRemark(ctx context.Context, in *UpdateGroupRemarkRequest, opts ...grpc.CallOption) (*UpdateGroupRemarkResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UpdateGroupRemark(ctx, in, opts...)
}

func (m *defaultUserService) QueryMyGroupList(ctx context.Context, in *QueryMyGroupListRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.QueryMyGroupList(ctx, in, opts...)
}

func (m *defaultUserService) SearchStrangerGroup(ctx context.Context, in *SearchStrangerGroupRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.SearchStrangerGroup(ctx, in, opts...)
}

func (m *defaultUserService) SearchMyGroupByName(ctx context.Context, in *SearchMyGroupByNameRequest, opts ...grpc.CallOption) (*MyGroupResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.SearchMyGroupByName(ctx, in, opts...)
}
