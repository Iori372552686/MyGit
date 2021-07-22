// Code generated by protoc-gen-go. DO NOT EDIT.
// source: friendsvr.proto

package g1_protocol

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

type PbFriend struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbFriend) Reset()         { *m = PbFriend{} }
func (m *PbFriend) String() string { return proto.CompactTextString(m) }
func (*PbFriend) ProtoMessage()    {}
func (*PbFriend) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{0}
}

func (m *PbFriend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbFriend.Unmarshal(m, b)
}
func (m *PbFriend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbFriend.Marshal(b, m, deterministic)
}
func (m *PbFriend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbFriend.Merge(m, src)
}
func (m *PbFriend) XXX_Size() int {
	return xxx_messageInfo_PbFriend.Size(m)
}
func (m *PbFriend) XXX_DiscardUnknown() {
	xxx_messageInfo_PbFriend.DiscardUnknown(m)
}

var xxx_messageInfo_PbFriend proto.InternalMessageInfo

func (m *PbFriend) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type PbFriendInvitation struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	InviteTime           int32    `protobuf:"varint,2,opt,name=invite_time,json=inviteTime,proto3" json:"invite_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbFriendInvitation) Reset()         { *m = PbFriendInvitation{} }
func (m *PbFriendInvitation) String() string { return proto.CompactTextString(m) }
func (*PbFriendInvitation) ProtoMessage()    {}
func (*PbFriendInvitation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{1}
}

func (m *PbFriendInvitation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbFriendInvitation.Unmarshal(m, b)
}
func (m *PbFriendInvitation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbFriendInvitation.Marshal(b, m, deterministic)
}
func (m *PbFriendInvitation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbFriendInvitation.Merge(m, src)
}
func (m *PbFriendInvitation) XXX_Size() int {
	return xxx_messageInfo_PbFriendInvitation.Size(m)
}
func (m *PbFriendInvitation) XXX_DiscardUnknown() {
	xxx_messageInfo_PbFriendInvitation.DiscardUnknown(m)
}

var xxx_messageInfo_PbFriendInvitation proto.InternalMessageInfo

func (m *PbFriendInvitation) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PbFriendInvitation) GetInviteTime() int32 {
	if m != nil {
		return m.InviteTime
	}
	return 0
}

type RoleFriendInfo struct {
	Uid                  uint64                `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FriendList           []*PbFriend           `protobuf:"bytes,2,rep,name=friend_list,json=friendList,proto3" json:"friend_list,omitempty"`
	BlackList            []uint64              `protobuf:"varint,3,rep,packed,name=black_list,json=blackList,proto3" json:"black_list,omitempty"`
	InvitationList       []*PbFriendInvitation `protobuf:"bytes,4,rep,name=invitation_list,json=invitationList,proto3" json:"invitation_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RoleFriendInfo) Reset()         { *m = RoleFriendInfo{} }
func (m *RoleFriendInfo) String() string { return proto.CompactTextString(m) }
func (*RoleFriendInfo) ProtoMessage()    {}
func (*RoleFriendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{2}
}

func (m *RoleFriendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleFriendInfo.Unmarshal(m, b)
}
func (m *RoleFriendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleFriendInfo.Marshal(b, m, deterministic)
}
func (m *RoleFriendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleFriendInfo.Merge(m, src)
}
func (m *RoleFriendInfo) XXX_Size() int {
	return xxx_messageInfo_RoleFriendInfo.Size(m)
}
func (m *RoleFriendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleFriendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoleFriendInfo proto.InternalMessageInfo

func (m *RoleFriendInfo) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RoleFriendInfo) GetFriendList() []*PbFriend {
	if m != nil {
		return m.FriendList
	}
	return nil
}

func (m *RoleFriendInfo) GetBlackList() []uint64 {
	if m != nil {
		return m.BlackList
	}
	return nil
}

func (m *RoleFriendInfo) GetInvitationList() []*PbFriendInvitation {
	if m != nil {
		return m.InvitationList
	}
	return nil
}

// FRIEND_INNER_ADD_FRIEND_REQ
type FriendInnerAddFriendReq struct {
	FromUid              uint64   `protobuf:"varint,1,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`
	IsDelFriend          bool     `protobuf:"varint,2,opt,name=is_del_friend,json=isDelFriend,proto3" json:"is_del_friend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendInnerAddFriendReq) Reset()         { *m = FriendInnerAddFriendReq{} }
func (m *FriendInnerAddFriendReq) String() string { return proto.CompactTextString(m) }
func (*FriendInnerAddFriendReq) ProtoMessage()    {}
func (*FriendInnerAddFriendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{3}
}

func (m *FriendInnerAddFriendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendInnerAddFriendReq.Unmarshal(m, b)
}
func (m *FriendInnerAddFriendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendInnerAddFriendReq.Marshal(b, m, deterministic)
}
func (m *FriendInnerAddFriendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendInnerAddFriendReq.Merge(m, src)
}
func (m *FriendInnerAddFriendReq) XXX_Size() int {
	return xxx_messageInfo_FriendInnerAddFriendReq.Size(m)
}
func (m *FriendInnerAddFriendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendInnerAddFriendReq.DiscardUnknown(m)
}

var xxx_messageInfo_FriendInnerAddFriendReq proto.InternalMessageInfo

func (m *FriendInnerAddFriendReq) GetFromUid() uint64 {
	if m != nil {
		return m.FromUid
	}
	return 0
}

func (m *FriendInnerAddFriendReq) GetIsDelFriend() bool {
	if m != nil {
		return m.IsDelFriend
	}
	return false
}

type FriendInnerAddFriendRsp struct {
	Ret                  *Ret     `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendInnerAddFriendRsp) Reset()         { *m = FriendInnerAddFriendRsp{} }
func (m *FriendInnerAddFriendRsp) String() string { return proto.CompactTextString(m) }
func (*FriendInnerAddFriendRsp) ProtoMessage()    {}
func (*FriendInnerAddFriendRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{4}
}

func (m *FriendInnerAddFriendRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendInnerAddFriendRsp.Unmarshal(m, b)
}
func (m *FriendInnerAddFriendRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendInnerAddFriendRsp.Marshal(b, m, deterministic)
}
func (m *FriendInnerAddFriendRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendInnerAddFriendRsp.Merge(m, src)
}
func (m *FriendInnerAddFriendRsp) XXX_Size() int {
	return xxx_messageInfo_FriendInnerAddFriendRsp.Size(m)
}
func (m *FriendInnerAddFriendRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendInnerAddFriendRsp.DiscardUnknown(m)
}

var xxx_messageInfo_FriendInnerAddFriendRsp proto.InternalMessageInfo

func (m *FriendInnerAddFriendRsp) GetRet() *Ret {
	if m != nil {
		return m.Ret
	}
	return nil
}

// FRIEND_INNER_CONFIRM_FRIEND_REQ
type FriendInnerConfirmFriendReq struct {
	FromUid              uint64   `protobuf:"varint,1,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`
	FromName             string   `protobuf:"bytes,2,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	IsReject             bool     `protobuf:"varint,3,opt,name=is_reject,json=isReject,proto3" json:"is_reject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendInnerConfirmFriendReq) Reset()         { *m = FriendInnerConfirmFriendReq{} }
func (m *FriendInnerConfirmFriendReq) String() string { return proto.CompactTextString(m) }
func (*FriendInnerConfirmFriendReq) ProtoMessage()    {}
func (*FriendInnerConfirmFriendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{5}
}

func (m *FriendInnerConfirmFriendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendInnerConfirmFriendReq.Unmarshal(m, b)
}
func (m *FriendInnerConfirmFriendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendInnerConfirmFriendReq.Marshal(b, m, deterministic)
}
func (m *FriendInnerConfirmFriendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendInnerConfirmFriendReq.Merge(m, src)
}
func (m *FriendInnerConfirmFriendReq) XXX_Size() int {
	return xxx_messageInfo_FriendInnerConfirmFriendReq.Size(m)
}
func (m *FriendInnerConfirmFriendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendInnerConfirmFriendReq.DiscardUnknown(m)
}

var xxx_messageInfo_FriendInnerConfirmFriendReq proto.InternalMessageInfo

func (m *FriendInnerConfirmFriendReq) GetFromUid() uint64 {
	if m != nil {
		return m.FromUid
	}
	return 0
}

func (m *FriendInnerConfirmFriendReq) GetFromName() string {
	if m != nil {
		return m.FromName
	}
	return ""
}

func (m *FriendInnerConfirmFriendReq) GetIsReject() bool {
	if m != nil {
		return m.IsReject
	}
	return false
}

type FriendInnerConfirmFriendRsp struct {
	Ret                  *Ret     `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendInnerConfirmFriendRsp) Reset()         { *m = FriendInnerConfirmFriendRsp{} }
func (m *FriendInnerConfirmFriendRsp) String() string { return proto.CompactTextString(m) }
func (*FriendInnerConfirmFriendRsp) ProtoMessage()    {}
func (*FriendInnerConfirmFriendRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65d963bc74d42e8, []int{6}
}

func (m *FriendInnerConfirmFriendRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendInnerConfirmFriendRsp.Unmarshal(m, b)
}
func (m *FriendInnerConfirmFriendRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendInnerConfirmFriendRsp.Marshal(b, m, deterministic)
}
func (m *FriendInnerConfirmFriendRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendInnerConfirmFriendRsp.Merge(m, src)
}
func (m *FriendInnerConfirmFriendRsp) XXX_Size() int {
	return xxx_messageInfo_FriendInnerConfirmFriendRsp.Size(m)
}
func (m *FriendInnerConfirmFriendRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendInnerConfirmFriendRsp.DiscardUnknown(m)
}

var xxx_messageInfo_FriendInnerConfirmFriendRsp proto.InternalMessageInfo

func (m *FriendInnerConfirmFriendRsp) GetRet() *Ret {
	if m != nil {
		return m.Ret
	}
	return nil
}

func init() {
	proto.RegisterType((*PbFriend)(nil), "g1.protocol.PbFriend")
	proto.RegisterType((*PbFriendInvitation)(nil), "g1.protocol.PbFriendInvitation")
	proto.RegisterType((*RoleFriendInfo)(nil), "g1.protocol.RoleFriendInfo")
	proto.RegisterType((*FriendInnerAddFriendReq)(nil), "g1.protocol.FriendInnerAddFriendReq")
	proto.RegisterType((*FriendInnerAddFriendRsp)(nil), "g1.protocol.FriendInnerAddFriendRsp")
	proto.RegisterType((*FriendInnerConfirmFriendReq)(nil), "g1.protocol.FriendInnerConfirmFriendReq")
	proto.RegisterType((*FriendInnerConfirmFriendRsp)(nil), "g1.protocol.FriendInnerConfirmFriendRsp")
}

func init() {
	proto.RegisterFile("friendsvr.proto", fileDescriptor_f65d963bc74d42e8)
}

var fileDescriptor_f65d963bc74d42e8 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x5f, 0x4b, 0xeb, 0x30,
	0x18, 0xc6, 0xe9, 0xba, 0x73, 0x4e, 0xf7, 0xf6, 0x9c, 0x6d, 0x04, 0x0e, 0x4e, 0xa7, 0xac, 0xe4,
	0xaa, 0x57, 0x03, 0x27, 0x78, 0xe7, 0xc5, 0x50, 0xfc, 0x03, 0x22, 0x12, 0x14, 0xbc, 0x2b, 0x5d,
	0x9b, 0xca, 0xab, 0x6d, 0x33, 0x93, 0xb8, 0x2f, 0xe8, 0x17, 0x93, 0x24, 0xab, 0x9b, 0x38, 0x65,
	0x77, 0xc9, 0xf3, 0x3c, 0xef, 0xd3, 0x5f, 0xde, 0x42, 0xaf, 0x90, 0xc8, 0xeb, 0x5c, 0x2d, 0xe4,
	0x78, 0x2e, 0x85, 0x16, 0x24, 0x7c, 0x3c, 0x74, 0xa7, 0x4c, 0x94, 0x7b, 0x7f, 0x33, 0x51, 0x55,
	0xa2, 0x76, 0x02, 0xdd, 0x87, 0xe0, 0x76, 0x76, 0x6e, 0xf3, 0xa4, 0x0f, 0xfe, 0x2b, 0xe6, 0x03,
	0x2f, 0xf2, 0xe2, 0x36, 0x33, 0x47, 0x7a, 0x01, 0xa4, 0x71, 0xaf, 0xea, 0x05, 0xea, 0x54, 0xa3,
	0xa8, 0xbf, 0xe6, 0xc8, 0x08, 0x42, 0x34, 0x3e, 0x4f, 0x34, 0x56, 0x7c, 0xd0, 0x8a, 0xbc, 0xf8,
	0x17, 0x03, 0x27, 0xdd, 0x61, 0xc5, 0xe9, 0x9b, 0x07, 0x5d, 0x26, 0x4a, 0xde, 0x74, 0x15, 0x62,
	0x43, 0xcb, 0x31, 0x84, 0x8e, 0x3c, 0x29, 0x51, 0xe9, 0x41, 0x2b, 0xf2, 0xe3, 0x70, 0xf2, 0x7f,
	0xbc, 0x06, 0x3f, 0x6e, 0x68, 0x18, 0xb8, 0xe4, 0x35, 0x2a, 0x4d, 0x0e, 0x00, 0x66, 0x65, 0x9a,
	0x3d, 0xbb, 0x31, 0x3f, 0xf2, 0xe3, 0x36, 0xeb, 0x58, 0xc5, 0xda, 0x97, 0xd0, 0xc3, 0x0f, 0x78,
	0x97, 0x69, 0xdb, 0xea, 0xd1, 0xc6, 0xea, 0xd5, 0x43, 0x59, 0x77, 0x35, 0x67, 0x9a, 0xe8, 0x03,
	0xec, 0x34, 0x99, 0x9a, 0xcb, 0x69, 0x9e, 0x2f, 0x61, 0xf8, 0x0b, 0xd9, 0x85, 0xa0, 0x90, 0xa2,
	0x4a, 0x56, 0x4f, 0xfa, 0x63, 0xee, 0xf7, 0x98, 0x13, 0x0a, 0xff, 0x50, 0x25, 0x39, 0x2f, 0x13,
	0xc7, 0x6c, 0xd7, 0x13, 0xb0, 0x10, 0xd5, 0x19, 0x2f, 0x5d, 0x03, 0x3d, 0xf9, 0xa6, 0x59, 0xcd,
	0x09, 0x05, 0x5f, 0x72, 0x6d, 0x4b, 0xc3, 0x49, 0xff, 0x13, 0x32, 0xe3, 0x9a, 0x19, 0x93, 0x4a,
	0x18, 0xae, 0x8d, 0x9f, 0x8a, 0xba, 0x40, 0x59, 0x6d, 0x05, 0x37, 0x84, 0x8e, 0xb5, 0xea, 0x74,
	0xf9, 0xdf, 0x3a, 0xcc, 0x66, 0x6f, 0xd2, 0x8a, 0x1b, 0x13, 0x55, 0x22, 0xf9, 0x13, 0xcf, 0xcc,
	0x5e, 0x0d, 0x75, 0x80, 0x8a, 0xd9, 0x3b, 0x9d, 0xfe, 0xf0, 0xcd, 0xed, 0xb0, 0x67, 0xbf, 0xad,
	0x74, 0xf4, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xe0, 0x96, 0x71, 0xb1, 0x02, 0x00, 0x00,
}
