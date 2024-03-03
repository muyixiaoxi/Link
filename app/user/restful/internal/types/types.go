// Code generated by goctl. DO NOT EDIT.
package types

type CancelRequest struct {
	TagIds []uint64 `json:"tagIds" validate:"required"`
}

type DisposeFlowedRequest struct {
	From   uint64 `json:"from"`
	To     uint64 `json:"to"`
	Type   uint32 `json:"type"`
	Remark string `json:"remark,optional"`
	Res    bool   `json:"res"`
}

type GroupList struct {
	Id              uint64 `json:"id"`              //群id
	Name            string `json:"name"`            //群聊名称
	SystemTagName   string `json:"systemTagName"`   //群聊所带的系统标签名称
	UserSelfTagName string `json:"userSelfTagName"` //用户自定义标签
	Avatar          string `json:"avatar"`          //用户头像
	Remark          string `json:"remark"`          //群聊备注
}

type GroupUserList struct {
	ID       uint64 `json:"id"`
	Avatar   string `json:"avatar"`
	UserName string `json:"userName"`
}

type GroupUserListRequest struct {
	GroupID uint64 `json:"groupId" validate:"required"` //群聊id
}

type GroupUserListResponse struct {
	GroupBossId   uint64          `json:"groupBossId"`
	GroupUserList []GroupUserList `json:"groupUserList"`
}

type KickOutGroupRequest struct {
	GroupId uint64   `json:"groupId"`                     //群id
	UserIDs []uint64 `json:"userIDs" validate:"required"` //被踢出群聊的人员id
}

type Message struct {
	From         uint64 `json:"from,optional"`
	To           uint64 `json:"to"`
	Type         uint8  `json:"type"`
	Time         string `json:"time"`
	Content      string `json:"content"`
	FromAvatar   string `json:"fromAvatar"`
	FromUsername string `json:"fromUsername"`
}

type QueryGroupListRequest struct {
	PageNo   uint64 `json:"pageNo" validate:"required"`
	PageSize uint64 `json:"pageSize" validate:"required"`
}

type QueryLink struct {
	Id        uint64 `json:"id"`        //标签id
	CreatorId uint64 `json:"creatorId"` //创作者id
	TagName   string `json:"tagName"`   //标签名称
}

type QueryLinkTagsResponse struct {
	LinkTags []QueryLink `json:"linkTags"`
}

type QuitGroupRequest struct {
	GroupId uint64 `json:"groupId" validate:"required"`
}

type RecommendGroupByTagRequest struct {
	PageNo   uint64   `json:"pageNo" validate:"required"`
	PageSize uint64   `json:"pageSize" validate:"required"`
	TagIds   []uint64 `json:"tagIds" validate:"max=5"` //标签id
}

type RecommendGroupByTagResponse struct {
	GroupList []GroupList `json:"groupList"`
	Total     int64       `json:"total"`
}

type RecommendUser struct {
	Id        uint64 `json:"id"`        //人员Id
	Username  string `json:"username"`  //用户名
	Avatar    string `json:"avatar"`    //用户头像
	Signature string `json:"signature"` //个性签名
}

type RecommendUserListResponse struct {
	RecommendUserList []RecommendUser `json:"recommendUserList"`
	Total             int64           `json:"total"`
}

type SelectGroupDeatilRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type SelectGroupDeatilResponse struct {
	ID              uint64  `json:"id"`
	Count           uint64  `json:"count"`
	GroupBossId     uint64  `json:"groupBossId"`
	Name            string  `json:"name"`
	UserSelfTagName string  `json:"userSelfTagName"`
	Avatar          string  `json:"avatar"`
	CreatedAt       string  `json:"createdAt"`
	SystemTagName   string  `json:"systemTagName"`
	Man             float32 `json:"man"`
	WoMan           float32 `json:"woMan"`
	Address         string  `json:"address"`
}

type SelectMyGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

type SelectStrangerRequest struct {
	PageNo   uint64 `json:"pageNo" validate:"required"`
	PageSize uint64 `json:"pageSize" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UpdateGroupInfoRequest struct {
	GroupId       uint64 `json:"groupId" validate:"required"` //群聊id
	SystemTagId   uint64 `json:"systemTagId"`                 //系统标签
	UserSelfTagId uint64 `json:"userSelfTagId"`               //用户自定义标签
	Name          string `json:"name"`                        //群聊名称
	Avatar        string `json:"avatar"`                      //群头像
}

type UpdateGroupRemarkRequest struct {
	GroupId uint64 `json:"groupId" validate:"required"`
	Remark  string `json:"remark" validate:"required"`
}

type UserAppleRequest struct {
	From    uint64 `json:"from,optional"`
	To      uint64 `json:"to"`
	Type    uint8  `json:"type"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type UserChooseTagRequst struct {
	SystemTagId uint64   `json:"systemTagId" validate:"required"'`
	TagIds      []uint64 `json:"tagIds" validate:"required,max=3"`
}

type UserCreateGroupRequset struct {
	SystemTagId   uint64 `json:"systemTagId" validate:"required"`
	UserSelfTagId uint64 `json:"userSelfTagId" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Avatar        string `json:"avatar"`
}

type UserCreateRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
	StartTagId uint64 `json:"startTagId"`
}

type UserCreateResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserDeleteFriendRequest struct {
	UserId   uint64 `json:"userId,optional"`
	FriendId uint64 `json:"friendId"`
}

type UserFlowedRequest struct {
	BeId    uint64 `json:"beId"`
	Message string `json:"message"`
	Type    uint64 `json:"type"`
}

type UserFriend struct {
	Id        uint64   `json:"id"`
	Username  string   `json:"username"`
	Avatar    string   `json:"avatar"`
	Remark    string   `json:"remark"`
	Signature string   `json:"signature"`
	TagName   []string `json:"tagName"`
}

type UserFriendsResponse struct {
	Friends []UserFriend `json:"friends"`
}

type UserInfoRequest struct {
	Id uint64 `path:"id"`
}

type UserInfoResponse struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Age       uint   `json:"age"`
	Gender    uint   `json:"gender"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Signature string `json:"signature"`
	IsFriend  bool   `json:"isFriend"`
	Remark    string `json:"remark"`
}

type UserLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token    string `json:"token"`
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserQueryFriendRequest struct {
	Param string `json:"param"`
}

type UserQueryPhoneRequest struct {
	Phone string `json:"phone"`
}

type UserQueryPhoneResponse struct {
	Id        uint64   `json:"id"`
	Username  string   `json:"username"`
	Avatar    string   `json:"avatar"`
	Address   string   `json:"address"`
	Signature string   `json:"signature"`
	TagName   []string `json:"tagName"`
}

type UserUpdateInfoRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password,optional"`
	Avatar    string `json:"avatar"`
	Age       uint   `json:"age"`
	Gender    uint   `json:"gender"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Signature string `json:"signature"`
}

type UserUpdateRemarkRequest struct {
	Friend uint64 `json:"friend"`
	Remark string `json:"remark"`
}
