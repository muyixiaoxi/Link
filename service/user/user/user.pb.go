// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: service/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{0}
}

type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserCreateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UserCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserCreateResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserCreateResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCreateResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserLoginResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserLoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Age      uint64 `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Gender   string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Address  string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Phone    string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfoResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoResponse) GetAge() uint64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserInfoResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfoResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserInfoResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UserUpdateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Age      uint64 `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Gender   uint64 `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Address  string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Phone    string `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserUpdateInfoRequest) Reset() {
	*x = UserUpdateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateInfoRequest) ProtoMessage() {}

func (x *UserUpdateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateInfoRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateInfoRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserUpdateInfoRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserUpdateInfoRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserUpdateInfoRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserUpdateInfoRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserUpdateInfoRequest) GetAge() uint64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserUpdateInfoRequest) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserUpdateInfoRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserUpdateInfoRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UserAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BeId    uint64 `protobuf:"varint,2,opt,name=beId,proto3" json:"beId,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Type    uint32 `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UserAddRequest) Reset() {
	*x = UserAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAddRequest) ProtoMessage() {}

func (x *UserAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAddRequest.ProtoReflect.Descriptor instead.
func (*UserAddRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserAddRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserAddRequest) GetBeId() uint64 {
	if x != nil {
		return x.BeId
	}
	return 0
}

func (x *UserAddRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserAddRequest) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UserUpdateRemarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BeId   uint64 `protobuf:"varint,2,opt,name=beId,proto3" json:"beId,omitempty"`
	Remark string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UserUpdateRemarkRequest) Reset() {
	*x = UserUpdateRemarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRemarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRemarkRequest) ProtoMessage() {}

func (x *UserUpdateRemarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRemarkRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRemarkRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserUpdateRemarkRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserUpdateRemarkRequest) GetBeId() uint64 {
	if x != nil {
		return x.BeId
	}
	return 0
}

func (x *UserUpdateRemarkRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type UserCreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupBossId   uint64 `protobuf:"varint,1,opt,name=group_boss_id,json=groupBossId,proto3" json:"group_boss_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SystemTagId   uint64 `protobuf:"varint,3,opt,name=system_tag_id,json=systemTagId,proto3" json:"system_tag_id,omitempty"`
	UserSelfTagId uint64 `protobuf:"varint,4,opt,name=user_self_tag_id,json=userSelfTagId,proto3" json:"user_self_tag_id,omitempty"`
}

func (x *UserCreateGroupRequest) Reset() {
	*x = UserCreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateGroupRequest) ProtoMessage() {}

func (x *UserCreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateGroupRequest.ProtoReflect.Descriptor instead.
func (*UserCreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_user_user_proto_rawDescGZIP(), []int{10}
}

func (x *UserCreateGroupRequest) GetGroupBossId() uint64 {
	if x != nil {
		return x.GroupBossId
	}
	return 0
}

func (x *UserCreateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserCreateGroupRequest) GetSystemTagId() uint64 {
	if x != nil {
		return x.SystemTagId
	}
	return 0
}

func (x *UserCreateGroupRequest) GetUserSelfTagId() uint64 {
	if x != nil {
		return x.UserSelfTagId
	}
	return 0
}

var File_service_user_user_proto protoreflect.FileDescriptor

var file_service_user_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x79, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0x58, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x4a, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x57, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x22, 0x21, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x62, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x62, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x32, 0xf4, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x2f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_user_user_proto_rawDescOnce sync.Once
	file_service_user_user_proto_rawDescData = file_service_user_user_proto_rawDesc
)

func file_service_user_user_proto_rawDescGZIP() []byte {
	file_service_user_user_proto_rawDescOnce.Do(func() {
		file_service_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_user_proto_rawDescData)
	})
	return file_service_user_user_proto_rawDescData
}

var file_service_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_user_user_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: user.Empty
	(*UserCreateRequest)(nil),       // 1: user.UserCreateRequest
	(*UserCreateResponse)(nil),      // 2: user.UserCreateResponse
	(*UserLoginRequest)(nil),        // 3: user.UserLoginRequest
	(*UserLoginResponse)(nil),       // 4: user.UserLoginResponse
	(*UserInfoRequest)(nil),         // 5: user.UserInfoRequest
	(*UserInfoResponse)(nil),        // 6: user.UserInfoResponse
	(*UserUpdateInfoRequest)(nil),   // 7: user.UserUpdateInfoRequest
	(*UserAddRequest)(nil),          // 8: user.UserAddRequest
	(*UserUpdateRemarkRequest)(nil), // 9: user.UserUpdateRemarkRequest
	(*UserCreateGroupRequest)(nil),  // 10: user.UserCreateGroupRequest
}
var file_user_proto_depIdxs = []int32{
	1, // 0: user.UserService.UserCreate:input_type -> user.UserCreateRequest
	3, // 1: user.UserService.UserLogin:input_type -> user.UserLoginRequest
	5, // 2: user.UserService.UserInfo:input_type -> user.UserInfoRequest
	7, // 3: user.UserService.UserUpdateInfo:input_type -> user.UserUpdateInfoRequest
	8, // 4: user.UserService.UserFlowed:input_type -> user.UserAddRequest
	9, // 5: user.UserService.UserUpdateRemark:input_type -> user.UserUpdateRemarkRequest
	2, // 6: user.UserService.UserCreate:output_type -> user.UserCreateResponse
	4, // 7: user.UserService.UserLogin:output_type -> user.UserLoginResponse
	6, // 8: user.UserService.UserInfo:output_type -> user.UserInfoResponse
	0, // 9: user.UserService.UserUpdateInfo:output_type -> user.Empty
	0, // 10: user.UserService.UserFlowed:output_type -> user.Empty
	0, // 11: user.UserService.UserUpdateRemark:output_type -> user.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_user_user_proto_init() }
func file_service_user_user_proto_init() {
	if File_service_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAddRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRemarkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_user_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_user_proto_goTypes,
		DependencyIndexes: file_service_user_user_proto_depIdxs,
		MessageInfos:      file_service_user_user_proto_msgTypes,
	}.Build()
	File_service_user_user_proto = out.File
	file_service_user_user_proto_rawDesc = nil
	file_service_user_user_proto_goTypes = nil
	file_service_user_user_proto_depIdxs = nil
}
