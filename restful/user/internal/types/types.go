// Code generated by goctl. DO NOT EDIT.
package types

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

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserInfoRequest struct {
	Id uint `json:"id"`
}

type UserInfoResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Age      uint   `json:"age"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserUpdateInfoRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Age      uint   `json:"age"`
	Gender   uint   `json:"gender"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserUpdateRemarkRequest struct {
	Friend uint64 `json:"friend"`
	Remark string `json:"remark"`
}

type UserFlowedRequest struct {
	BeId    uint64 `json:"beId"`
	Message string `json:"message"`
	Type    uint32 `json:"type"`
}
