// Code generated by goctl. DO NOT EDIT.
package types

type DisposeFlowedRequest struct {
	From   uint64 `json:"from"`
	To     uint64 `json:"to"`
	Type   uint32 `json:"type"`
	Remark string `json:"remark,optional"`
	Res    bool   `json:"res"`
}

type Message struct {
	From    uint64 `json:"from,optional"`
	To      uint64 `json:"to"`
	Type    uint8  `json:"type"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

type QueryLink struct {
	Id        uint64 `json:"id"`        //标签id
	CreatorId uint64 `json:"creatorId"` //创作者id
	TagName   string `json:"tagName"`   //标签名称
}

type QueryLinkTagsResponse struct {
	LinkTags []QueryLink `json:"linkTags"`
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

type UserFlowedRequest struct {
	BeId    uint64 `json:"beId"`
	Message string `json:"message"`
	Type    uint64 `json:"type"`
}

type UserFriend struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
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
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Signature string `json:"signature"`
	IsFriend  bool   `json:"isFriend"`
	Remark    string `json:"remark"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token    string `json:"token"`
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
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
