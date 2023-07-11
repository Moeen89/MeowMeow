// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/biz.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthKey              int32    `protobuf:"varint,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	MessageId            int32    `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e749db24d5e0fb, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRequest) GetAuthKey() int32 {
	if m != nil {
		return m.AuthKey
	}
	return 0
}

func (m *UserRequest) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type UserRequestWithSqlInject struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthKey              int32    `protobuf:"varint,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	MessageId            int32    `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequestWithSqlInject) Reset()         { *m = UserRequestWithSqlInject{} }
func (m *UserRequestWithSqlInject) String() string { return proto.CompactTextString(m) }
func (*UserRequestWithSqlInject) ProtoMessage()    {}
func (*UserRequestWithSqlInject) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e749db24d5e0fb, []int{1}
}

func (m *UserRequestWithSqlInject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequestWithSqlInject.Unmarshal(m, b)
}
func (m *UserRequestWithSqlInject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequestWithSqlInject.Marshal(b, m, deterministic)
}
func (m *UserRequestWithSqlInject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequestWithSqlInject.Merge(m, src)
}
func (m *UserRequestWithSqlInject) XXX_Size() int {
	return xxx_messageInfo_UserRequestWithSqlInject.Size(m)
}
func (m *UserRequestWithSqlInject) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequestWithSqlInject.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequestWithSqlInject proto.InternalMessageInfo

func (m *UserRequestWithSqlInject) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserRequestWithSqlInject) GetAuthKey() int32 {
	if m != nil {
		return m.AuthKey
	}
	return 0
}

func (m *UserRequestWithSqlInject) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type UserResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e749db24d5e0fb, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserResponse) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Family               string   `protobuf:"bytes,2,opt,name=family,proto3" json:"family,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  string   `protobuf:"bytes,5,opt,name=sex,proto3" json:"sex,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e749db24d5e0fb, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Key struct {
	AuthKey              int32    `protobuf:"varint,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e749db24d5e0fb, []int{4}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetAuthKey() int32 {
	if m != nil {
		return m.AuthKey
	}
	return 0
}

type Val struct {
	IsTrue               int32    `protobuf:"varint,1,opt,name=isTrue,proto3" json:"isTrue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Val) Reset()         { *m = Val{} }
func (m *Val) String() string { return proto.CompactTextString(m) }
func (*Val) ProtoMessage()    {}
func (*Val) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e749db24d5e0fb, []int{5}
}

func (m *Val) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Val.Unmarshal(m, b)
}
func (m *Val) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Val.Marshal(b, m, deterministic)
}
func (m *Val) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Val.Merge(m, src)
}
func (m *Val) XXX_Size() int {
	return xxx_messageInfo_Val.Size(m)
}
func (m *Val) XXX_DiscardUnknown() {
	xxx_messageInfo_Val.DiscardUnknown(m)
}

var xxx_messageInfo_Val proto.InternalMessageInfo

func (m *Val) GetIsTrue() int32 {
	if m != nil {
		return m.IsTrue
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "biz.v1.UserRequest")
	proto.RegisterType((*UserRequestWithSqlInject)(nil), "biz.v1.UserRequest_with_sql_inject")
	proto.RegisterType((*UserResponse)(nil), "biz.v1.UserResponse")
	proto.RegisterType((*User)(nil), "biz.v1.User")
	proto.RegisterType((*Key)(nil), "biz.v1.key")
	proto.RegisterType((*Val)(nil), "biz.v1.val")
}

func init() {
	proto.RegisterFile("pb/biz.proto", fileDescriptor_49e749db24d5e0fb)
}

var fileDescriptor_49e749db24d5e0fb = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xfd, 0xb5, 0x85, 0x42, 0x07, 0xf2, 0x8b, 0x59, 0x8d, 0x16, 0x94, 0x48, 0xea, 0x41, 0x4e,
	0x10, 0x30, 0xf1, 0xae, 0x37, 0xe2, 0x49, 0xa2, 0x1e, 0xbc, 0xd4, 0xa5, 0x1d, 0x4b, 0x6d, 0x29,
	0xa5, 0xdb, 0x56, 0xeb, 0xd9, 0x8f, 0xe6, 0x07, 0x33, 0xbb, 0x6d, 0x91, 0x3f, 0x7a, 0xf3, 0x36,
	0xef, 0xcd, 0xcc, 0x9b, 0xd9, 0x79, 0x0b, 0xcd, 0x70, 0x3a, 0x98, 0xba, 0xef, 0xfd, 0x30, 0x5a,
	0xc4, 0x0b, 0xa2, 0xf2, 0x30, 0x1d, 0x1a, 0x4f, 0xd0, 0xb8, 0x67, 0x18, 0x4d, 0x70, 0x99, 0x20,
	0x8b, 0xc9, 0x11, 0xd4, 0x12, 0x86, 0x91, 0xe9, 0xda, 0xba, 0xd4, 0x95, 0x7a, 0xd5, 0x89, 0xca,
	0xe1, 0xd8, 0x26, 0x2d, 0xa8, 0xd3, 0x24, 0x9e, 0x99, 0x1e, 0x66, 0xba, 0x2c, 0x32, 0x35, 0x8e,
	0x6f, 0x30, 0x23, 0x1d, 0x80, 0x39, 0x32, 0x46, 0x1d, 0xe4, 0x6d, 0x8a, 0x48, 0x6a, 0x05, 0x33,
	0xb6, 0x8d, 0x10, 0x8e, 0xd7, 0x26, 0x98, 0xaf, 0x6e, 0x3c, 0x33, 0xd9, 0xd2, 0x37, 0xdd, 0xe0,
	0x05, 0xad, 0x9d, 0x89, 0xda, 0x1f, 0x4c, 0xbc, 0x85, 0x66, 0x3e, 0x91, 0x85, 0x8b, 0x80, 0x21,
	0x31, 0xa0, 0xca, 0x35, 0x99, 0x2e, 0x75, 0x95, 0x5e, 0x63, 0xd4, 0xec, 0xe7, 0x6f, 0xef, 0x8b,
	0xa2, 0x3c, 0xb5, 0x25, 0x29, 0x6f, 0x4b, 0x7e, 0x48, 0x50, 0xe1, 0xe5, 0x84, 0x40, 0x25, 0xa0,
	0x73, 0x2c, 0x76, 0x15, 0x31, 0x39, 0x04, 0xf5, 0x99, 0xce, 0x5d, 0x3f, 0xdf, 0x53, 0x9b, 0x14,
	0x88, 0xfc, 0x07, 0x79, 0xb5, 0x9e, 0xec, 0xda, 0x64, 0x0f, 0x14, 0xea, 0xa0, 0x5e, 0x11, 0x04,
	0x0f, 0x39, 0xc3, 0xf0, 0x4d, 0xaf, 0x8a, 0x36, 0x1e, 0x92, 0x13, 0xd0, 0xac, 0x08, 0x69, 0x8c,
	0xf6, 0x55, 0xac, 0xab, 0x82, 0xff, 0x26, 0x8c, 0x53, 0x50, 0x3c, 0xcc, 0x88, 0x0e, 0xe5, 0x29,
	0x0a, 0x97, 0x4a, 0x68, 0x74, 0x40, 0x49, 0xa9, 0xcf, 0x37, 0x72, 0xd9, 0x5d, 0x94, 0x60, 0xe9,
	0x62, 0x8e, 0x46, 0x9f, 0x12, 0x68, 0x0e, 0xc6, 0x66, 0xfe, 0xe6, 0xcb, 0x75, 0xb0, 0xbf, 0x71,
	0x95, 0xdc, 0xac, 0xf6, 0xc1, 0x26, 0x59, 0xdc, 0xf3, 0x01, 0x5a, 0xab, 0xbe, 0x1d, 0x3f, 0xcf,
	0x7e, 0xd0, 0xd9, 0x2e, 0xfa, 0x45, 0xf7, 0x1c, 0x34, 0x6b, 0x86, 0x96, 0xc7, 0x2d, 0x27, 0x8d,
	0xb2, 0xc4, 0xc3, 0xac, 0xbd, 0x02, 0x29, 0xf5, 0x8d, 0x7f, 0xd7, 0xf5, 0x47, 0xd5, 0xc1, 0x60,
	0x90, 0x0e, 0xa7, 0xaa, 0xf8, 0xcd, 0x17, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x9b, 0xc4,
	0x26, 0xdd, 0x02, 0x00, 0x00,
}