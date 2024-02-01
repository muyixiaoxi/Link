// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: service/tag/tag.proto

package tag

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
		mi := &file_service_tag_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[0]
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
	return file_service_tag_tag_proto_rawDescGZIP(), []int{0}
}

// CreateTagRequest创建标签 start
type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorId uint64 `protobuf:"varint,1,opt,name=creatorId,proto3" json:"creatorId,omitempty"`
	TagName   string `protobuf:"bytes,2,opt,name=tagName,proto3" json:"tagName,omitempty"`
	GroupName string `protobuf:"bytes,3,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTagRequest) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *CreateTagRequest) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *CreateTagRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *CreateTagRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string   `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	TagName   []string `protobuf:"bytes,3,rep,name=tagName,proto3" json:"tagName,omitempty"`
}

func (x *CreateTagResponse) Reset() {
	*x = CreateTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagResponse) ProtoMessage() {}

func (x *CreateTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagResponse.ProtoReflect.Descriptor instead.
func (*CreateTagResponse) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTagResponse) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *CreateTagResponse) GetTagName() []string {
	if x != nil {
		return x.TagName
	}
	return nil
}

// 删除标签 start
type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorId uint64   `protobuf:"varint,1,opt,name=creatorId,proto3" json:"creatorId,omitempty"`
	TagId     []uint64 `protobuf:"varint,2,rep,packed,name=tagId,proto3" json:"tagId,omitempty"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTagRequest) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *DeleteTagRequest) GetTagId() []uint64 {
	if x != nil {
		return x.TagId
	}
	return nil
}

type DeleteTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagResponse) Reset() {
	*x = DeleteTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagResponse) ProtoMessage() {}

func (x *DeleteTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagResponse.ProtoReflect.Descriptor instead.
func (*DeleteTagResponse) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{4}
}

type UserChooseTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TagId  uint64 `protobuf:"varint,2,opt,name=tagId,proto3" json:"tagId,omitempty"`
}

func (x *UserChooseTagRequest) Reset() {
	*x = UserChooseTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChooseTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChooseTagRequest) ProtoMessage() {}

func (x *UserChooseTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChooseTagRequest.ProtoReflect.Descriptor instead.
func (*UserChooseTagRequest) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{5}
}

func (x *UserChooseTagRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserChooseTagRequest) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

// 查询全部标签组 start
type GroupTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
}

func (x *GroupTag) Reset() {
	*x = GroupTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupTag) ProtoMessage() {}

func (x *GroupTag) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupTag.ProtoReflect.Descriptor instead.
func (*GroupTag) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{6}
}

func (x *GroupTag) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupTag) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

type GroupTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*GroupTag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GroupTagResponse) Reset() {
	*x = GroupTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupTagResponse) ProtoMessage() {}

func (x *GroupTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupTagResponse.ProtoReflect.Descriptor instead.
func (*GroupTagResponse) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{7}
}

func (x *GroupTagResponse) GetTags() []*GroupTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

// 根据所选标签组 查询对应小标签
type SelectAllTagsByGroupName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
}

func (x *SelectAllTagsByGroupName) Reset() {
	*x = SelectAllTagsByGroupName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectAllTagsByGroupName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectAllTagsByGroupName) ProtoMessage() {}

func (x *SelectAllTagsByGroupName) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectAllTagsByGroupName.ProtoReflect.Descriptor instead.
func (*SelectAllTagsByGroupName) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{8}
}

func (x *SelectAllTagsByGroupName) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

type AllTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsCreator bool   `protobuf:"varint,3,opt,name=isCreator,proto3" json:"isCreator,omitempty"`
}

func (x *AllTags) Reset() {
	*x = AllTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllTags) ProtoMessage() {}

func (x *AllTags) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllTags.ProtoReflect.Descriptor instead.
func (*AllTags) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{9}
}

func (x *AllTags) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AllTags) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllTags) GetIsCreator() bool {
	if x != nil {
		return x.IsCreator
	}
	return false
}

type AllTagsByGroupNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowTags []*AllTags `protobuf:"bytes,1,rep,name=lowTags,proto3" json:"lowTags,omitempty"`
}

func (x *AllTagsByGroupNameResponse) Reset() {
	*x = AllTagsByGroupNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tag_tag_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllTagsByGroupNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllTagsByGroupNameResponse) ProtoMessage() {}

func (x *AllTagsByGroupNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_tag_tag_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllTagsByGroupNameResponse.ProtoReflect.Descriptor instead.
func (*AllTagsByGroupNameResponse) Descriptor() ([]byte, []int) {
	return file_service_tag_tag_proto_rawDescGZIP(), []int{10}
}

func (x *AllTagsByGroupNameResponse) GetLowTags() []*AllTags {
	if x != nil {
		return x.LowTags
	}
	return nil
}

var File_service_tag_tag_proto protoreflect.FileDescriptor

var file_service_tag_tag_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x61, 0x67, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x7c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a,
	0x14, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x61,
	0x67, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a,
	0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0x38, 0x0a, 0x18, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b,
	0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x1a, 0x41,
	0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x6c, 0x6f, 0x77,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x52, 0x07, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x67,
	0x73, 0x32, 0xcb, 0x02, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x3a,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x54, 0x61, 0x67, 0x12, 0x0a, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x61, 0x67, 0x73, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x1f,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x0f, 0x0a, 0x0d, 0x54, 0x61, 0x67, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_tag_tag_proto_rawDescOnce sync.Once
	file_service_tag_tag_proto_rawDescData = file_service_tag_tag_proto_rawDesc
)

func file_service_tag_tag_proto_rawDescGZIP() []byte {
	file_service_tag_tag_proto_rawDescOnce.Do(func() {
		file_service_tag_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_tag_tag_proto_rawDescData)
	})
	return file_service_tag_tag_proto_rawDescData
}

var file_service_tag_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_tag_tag_proto_goTypes = []interface{}{
	(*Empty)(nil),                      // 0: tag.Empty
	(*CreateTagRequest)(nil),           // 1: tag.CreateTagRequest
	(*CreateTagResponse)(nil),          // 2: tag.CreateTagResponse
	(*DeleteTagRequest)(nil),           // 3: tag.DeleteTagRequest
	(*DeleteTagResponse)(nil),          // 4: tag.DeleteTagResponse
	(*UserChooseTagRequest)(nil),       // 5: tag.UserChooseTagRequest
	(*GroupTag)(nil),                   // 6: tag.GroupTag
	(*GroupTagResponse)(nil),           // 7: tag.GroupTagResponse
	(*SelectAllTagsByGroupName)(nil),   // 8: tag.SelectAllTagsByGroupName
	(*AllTags)(nil),                    // 9: tag.AllTags
	(*AllTagsByGroupNameResponse)(nil), // 10: tag.AllTagsByGroupNameResponse
}
var file_service_tag_tag_proto_depIdxs = []int32{
	6,  // 0: tag.GroupTagResponse.tags:type_name -> tag.GroupTag
	9,  // 1: tag.AllTagsByGroupNameResponse.lowTags:type_name -> tag.AllTags
	1,  // 2: tag.TagLogin.CreateTag:input_type -> tag.CreateTagRequest
	1,  // 3: tag.TagLogin.UpdateTag:input_type -> tag.CreateTagRequest
	3,  // 4: tag.TagLogin.DeleteTag:input_type -> tag.DeleteTagRequest
	0,  // 5: tag.TagLogin.SelectGroupTag:input_type -> tag.Empty
	8,  // 6: tag.TagLogin.SelectAllTagsByGroup:input_type -> tag.SelectAllTagsByGroupName
	2,  // 7: tag.TagLogin.CreateTag:output_type -> tag.CreateTagResponse
	2,  // 8: tag.TagLogin.UpdateTag:output_type -> tag.CreateTagResponse
	4,  // 9: tag.TagLogin.DeleteTag:output_type -> tag.DeleteTagResponse
	7,  // 10: tag.TagLogin.SelectGroupTag:output_type -> tag.GroupTagResponse
	10, // 11: tag.TagLogin.SelectAllTagsByGroup:output_type -> tag.AllTagsByGroupNameResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_tag_tag_proto_init() }
func file_service_tag_tag_proto_init() {
	if File_service_tag_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_tag_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_tag_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_service_tag_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagResponse); i {
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
		file_service_tag_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
		file_service_tag_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagResponse); i {
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
		file_service_tag_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChooseTagRequest); i {
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
		file_service_tag_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupTag); i {
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
		file_service_tag_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupTagResponse); i {
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
		file_service_tag_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectAllTagsByGroupName); i {
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
		file_service_tag_tag_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllTags); i {
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
		file_service_tag_tag_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllTagsByGroupNameResponse); i {
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
			RawDescriptor: file_service_tag_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_tag_tag_proto_goTypes,
		DependencyIndexes: file_service_tag_tag_proto_depIdxs,
		MessageInfos:      file_service_tag_tag_proto_msgTypes,
	}.Build()
	File_service_tag_tag_proto = out.File
	file_service_tag_tag_proto_rawDesc = nil
	file_service_tag_tag_proto_goTypes = nil
	file_service_tag_tag_proto_depIdxs = nil
}
