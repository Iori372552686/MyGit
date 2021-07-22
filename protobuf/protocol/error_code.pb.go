// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error_code.proto

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

type ErrorCode int32

const (
	// 公共错误码，范围：[-99,0]
	ErrorCode_ERR_OK                              ErrorCode = 0
	ErrorCode_ERR_FAIL                            ErrorCode = -1
	ErrorCode_ERR_TIMEOUT                         ErrorCode = -2
	ErrorCode_ERR_MARSHAL                         ErrorCode = -3
	ErrorCode_ERR_NOT_EXIST                       ErrorCode = -4
	ErrorCode_ERR_DB                              ErrorCode = -5
	ErrorCode_ERR_CONF                            ErrorCode = -6
	ErrorCode_ERR_ARGV                            ErrorCode = -7
	ErrorCode_ERR_ZONE                            ErrorCode = -8
	ErrorCode_ERR_INSTANCE_NOT_AVAILABLE          ErrorCode = -10001
	ErrorCode_ERR_ITEM_NOT_ENOUGH                 ErrorCode = -10002
	ErrorCode_ERR_ITEM_ADD_ERROR                  ErrorCode = -10003
	ErrorCode_ERR_ITEM_CARD_NOT_EXIST             ErrorCode = -10004
	ErrorCode_ERR_INSTANCE_NOT_PASS               ErrorCode = -10005
	ErrorCode_ERR_REWARD_ALREADY_GET              ErrorCode = -10006
	ErrorCode_ERR_SINEW_NOT_ENOUGH                ErrorCode = -10007
	ErrorCode_ERR_GOLD_NOT_ENOUGH                 ErrorCode = -10008
	ErrorCode_ERR_DIAMOND_NOT_ENOUGH              ErrorCode = -10009
	ErrorCode_ERR_RES_INSTANCE_MAX_TIMES          ErrorCode = -10010
	ErrorCode_ERR_RES_INSTANCE_NOT_OPEN           ErrorCode = -10011
	ErrorCode_ERR_RES_INSTANCE_MAX_BUY_TIMES      ErrorCode = -10012
	ErrorCode_ERR_FURY_COIN_NOT_ENOUGH            ErrorCode = -10013
	ErrorCode_ERR_SEMICONDUCTOR_NOT_ENOUGH        ErrorCode = -10014
	ErrorCode_ERR_ICON_NOT_HAVE                   ErrorCode = -10020
	ErrorCode_ERR_FRAME_NOT_HAVE                  ErrorCode = -10021
	ErrorCode_ERR_IMAGE_NOT_HAVE                  ErrorCode = -10022
	ErrorCode_ERR_INVALID_NAME                    ErrorCode = -10023
	ErrorCode_ERR_SENSITIVE_WORD                  ErrorCode = -10024
	ErrorCode_ERR_DUPLICATE_NAME                  ErrorCode = -10025
	ErrorCode_ERR_HERO_MAX_STAR                   ErrorCode = -10030
	ErrorCode_ERR_HERO_MAX_LEVEL                  ErrorCode = -10031
	ErrorCode_ERR_FRIEND_ALREADY_FRIEND           ErrorCode = -10050
	ErrorCode_ERR_FRIEND_NOT_EXIST                ErrorCode = -10051
	ErrorCode_ERR_FRIEND_MAX_INVITATION_NUM       ErrorCode = -10052
	ErrorCode_ERR_FRIEND_OTHER_MAX_INVITATION_NUM ErrorCode = -10053
	ErrorCode_ERR_FRIEND_INVITATION_NOT_EXIST     ErrorCode = -10054
	ErrorCode_ERR_FRIEND_MAX_FRIEND_NUM           ErrorCode = -10055
	ErrorCode_ERR_FRIEND_OTHER_MAX_FRIEND_NUM     ErrorCode = -10056
	ErrorCode_ERR_FRIEND_SHIELD                   ErrorCode = -10057
	ErrorCode_ERR_FRIEND_MAX_SHIELD               ErrorCode = -10058
	ErrorCode_ERR_MAIN_TASK_NOT_FINISH            ErrorCode = -10150
	ErrorCode_ERR_GUIDE_IS_EXIST                  ErrorCode = -10187
	ErrorCode_ERR_FUNCTION_NOT_OPEN               ErrorCode = -10188
	// 排行榜
	ErrorCode_ERR_RANK_TYPE_ERROR ErrorCode = -90000
)

var ErrorCode_name = map[int32]string{
	0:      "ERR_OK",
	-1:     "ERR_FAIL",
	-2:     "ERR_TIMEOUT",
	-3:     "ERR_MARSHAL",
	-4:     "ERR_NOT_EXIST",
	-5:     "ERR_DB",
	-6:     "ERR_CONF",
	-7:     "ERR_ARGV",
	-8:     "ERR_ZONE",
	-10001: "ERR_INSTANCE_NOT_AVAILABLE",
	-10002: "ERR_ITEM_NOT_ENOUGH",
	-10003: "ERR_ITEM_ADD_ERROR",
	-10004: "ERR_ITEM_CARD_NOT_EXIST",
	-10005: "ERR_INSTANCE_NOT_PASS",
	-10006: "ERR_REWARD_ALREADY_GET",
	-10007: "ERR_SINEW_NOT_ENOUGH",
	-10008: "ERR_GOLD_NOT_ENOUGH",
	-10009: "ERR_DIAMOND_NOT_ENOUGH",
	-10010: "ERR_RES_INSTANCE_MAX_TIMES",
	-10011: "ERR_RES_INSTANCE_NOT_OPEN",
	-10012: "ERR_RES_INSTANCE_MAX_BUY_TIMES",
	-10013: "ERR_FURY_COIN_NOT_ENOUGH",
	-10014: "ERR_SEMICONDUCTOR_NOT_ENOUGH",
	-10020: "ERR_ICON_NOT_HAVE",
	-10021: "ERR_FRAME_NOT_HAVE",
	-10022: "ERR_IMAGE_NOT_HAVE",
	-10023: "ERR_INVALID_NAME",
	-10024: "ERR_SENSITIVE_WORD",
	-10025: "ERR_DUPLICATE_NAME",
	-10030: "ERR_HERO_MAX_STAR",
	-10031: "ERR_HERO_MAX_LEVEL",
	-10050: "ERR_FRIEND_ALREADY_FRIEND",
	-10051: "ERR_FRIEND_NOT_EXIST",
	-10052: "ERR_FRIEND_MAX_INVITATION_NUM",
	-10053: "ERR_FRIEND_OTHER_MAX_INVITATION_NUM",
	-10054: "ERR_FRIEND_INVITATION_NOT_EXIST",
	-10055: "ERR_FRIEND_MAX_FRIEND_NUM",
	-10056: "ERR_FRIEND_OTHER_MAX_FRIEND_NUM",
	-10057: "ERR_FRIEND_SHIELD",
	-10058: "ERR_FRIEND_MAX_SHIELD",
	-10150: "ERR_MAIN_TASK_NOT_FINISH",
	-10187: "ERR_GUIDE_IS_EXIST",
	-10188: "ERR_FUNCTION_NOT_OPEN",
	-90000: "ERR_RANK_TYPE_ERROR",
}

var ErrorCode_value = map[string]int32{
	"ERR_OK":                              0,
	"ERR_FAIL":                            -1,
	"ERR_TIMEOUT":                         -2,
	"ERR_MARSHAL":                         -3,
	"ERR_NOT_EXIST":                       -4,
	"ERR_DB":                              -5,
	"ERR_CONF":                            -6,
	"ERR_ARGV":                            -7,
	"ERR_ZONE":                            -8,
	"ERR_INSTANCE_NOT_AVAILABLE":          -10001,
	"ERR_ITEM_NOT_ENOUGH":                 -10002,
	"ERR_ITEM_ADD_ERROR":                  -10003,
	"ERR_ITEM_CARD_NOT_EXIST":             -10004,
	"ERR_INSTANCE_NOT_PASS":               -10005,
	"ERR_REWARD_ALREADY_GET":              -10006,
	"ERR_SINEW_NOT_ENOUGH":                -10007,
	"ERR_GOLD_NOT_ENOUGH":                 -10008,
	"ERR_DIAMOND_NOT_ENOUGH":              -10009,
	"ERR_RES_INSTANCE_MAX_TIMES":          -10010,
	"ERR_RES_INSTANCE_NOT_OPEN":           -10011,
	"ERR_RES_INSTANCE_MAX_BUY_TIMES":      -10012,
	"ERR_FURY_COIN_NOT_ENOUGH":            -10013,
	"ERR_SEMICONDUCTOR_NOT_ENOUGH":        -10014,
	"ERR_ICON_NOT_HAVE":                   -10020,
	"ERR_FRAME_NOT_HAVE":                  -10021,
	"ERR_IMAGE_NOT_HAVE":                  -10022,
	"ERR_INVALID_NAME":                    -10023,
	"ERR_SENSITIVE_WORD":                  -10024,
	"ERR_DUPLICATE_NAME":                  -10025,
	"ERR_HERO_MAX_STAR":                   -10030,
	"ERR_HERO_MAX_LEVEL":                  -10031,
	"ERR_FRIEND_ALREADY_FRIEND":           -10050,
	"ERR_FRIEND_NOT_EXIST":                -10051,
	"ERR_FRIEND_MAX_INVITATION_NUM":       -10052,
	"ERR_FRIEND_OTHER_MAX_INVITATION_NUM": -10053,
	"ERR_FRIEND_INVITATION_NOT_EXIST":     -10054,
	"ERR_FRIEND_MAX_FRIEND_NUM":           -10055,
	"ERR_FRIEND_OTHER_MAX_FRIEND_NUM":     -10056,
	"ERR_FRIEND_SHIELD":                   -10057,
	"ERR_FRIEND_MAX_SHIELD":               -10058,
	"ERR_MAIN_TASK_NOT_FINISH":            -10150,
	"ERR_GUIDE_IS_EXIST":                  -10187,
	"ERR_FUNCTION_NOT_OPEN":               -10188,
	"ERR_RANK_TYPE_ERROR":                 -90000,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5513ac0a8e17e40, []int{0}
}

func init() {
	proto.RegisterEnum("g1.protocol.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() {
	proto.RegisterFile("error_code.proto", fileDescriptor_c5513ac0a8e17e40)
}

var fileDescriptor_c5513ac0a8e17e40 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xc9, 0x6e, 0x14, 0x3d,
	0x10, 0xc7, 0xbf, 0x5c, 0xa2, 0x0f, 0x47, 0x48, 0xc5, 0x84, 0x40, 0x88, 0x48, 0x02, 0x84, 0x45,
	0x04, 0x84, 0x40, 0x3c, 0x41, 0xa5, 0x5d, 0x33, 0x6d, 0xa5, 0xdb, 0x1e, 0xd9, 0xee, 0x4e, 0x86,
	0x4b, 0x4b, 0x24, 0x11, 0x17, 0xa4, 0x41, 0x23, 0x1e, 0x2c, 0x91, 0xd8, 0x37, 0xf1, 0x0a, 0x70,
	0x01, 0x02, 0x27, 0xf6, 0x9d, 0x1c, 0xd9, 0x41, 0xbd, 0x64, 0xdc, 0x8c, 0x93, 0x39, 0xf5, 0x54,
	0xfd, 0x5c, 0xf5, 0xaf, 0xc5, 0x66, 0xb0, 0xd2, 0xeb, 0x75, 0x7b, 0xd9, 0x52, 0x77, 0x79, 0xe5,
	0xd4, 0xa5, 0x5e, 0xf7, 0x72, 0xb7, 0x31, 0x72, 0xe1, 0x4c, 0xf9, 0xb5, 0xd4, 0xbd, 0x38, 0xbb,
	0x31, 0xc2, 0x76, 0x50, 0x4e, 0x04, 0xdd, 0xe5, 0x95, 0x06, 0x63, 0xc3, 0xa4, 0x75, 0xa6, 0xe6,
	0xe1, 0xbf, 0xc6, 0x18, 0xfb, 0x3f, 0xff, 0x6e, 0xa2, 0x88, 0xe0, 0xcf, 0xe6, 0x6f, 0xa8, 0x31,
	0xce, 0x46, 0x72, 0xb3, 0x15, 0x31, 0xa9, 0xc4, 0xc2, 0x6f, 0xcf, 0x13, 0xa3, 0x36, 0x21, 0x46,
	0xf0, 0xcb, 0x79, 0x26, 0xd8, 0xce, 0xdc, 0x23, 0x95, 0xcd, 0x68, 0x51, 0x18, 0x0b, 0x3f, 0x9d,
	0x6f, 0xb4, 0x4c, 0xc9, 0xe7, 0xe0, 0x87, 0x33, 0x56, 0xb9, 0x03, 0x25, 0x9b, 0xf0, 0xdd, 0x33,
	0xa3, 0x6e, 0xa5, 0xf0, 0xcd, 0x33, 0x9f, 0x53, 0x92, 0xe0, 0xab, 0x33, 0x1f, 0x63, 0x13, 0xb9,
	0x59, 0x48, 0x63, 0x51, 0x06, 0x54, 0xa4, 0xc7, 0x14, 0x45, 0x84, 0x73, 0x11, 0xc1, 0x97, 0xb5,
	0x3e, 0x78, 0x80, 0x8d, 0x16, 0xa0, 0xa5, 0xb8, 0xd4, 0x28, 0x55, 0xd2, 0x0a, 0xe1, 0xb3, 0x23,
	0xa6, 0x59, 0xa3, 0x4f, 0x20, 0xe7, 0x19, 0x69, 0xad, 0x34, 0x7c, 0x72, 0xc0, 0x61, 0xb6, 0xb7,
	0x0f, 0x04, 0xa8, 0x79, 0xad, 0xd6, 0x8f, 0x8e, 0x3a, 0xc4, 0xc6, 0x3c, 0x45, 0x6d, 0x34, 0x06,
	0x3e, 0x38, 0x66, 0x86, 0xed, 0xc9, 0x19, 0x4d, 0x0b, 0x79, 0x18, 0x8c, 0x34, 0x21, 0xef, 0x64,
	0x2d, 0xb2, 0xf0, 0xde, 0x41, 0x07, 0xd9, 0xee, 0x1c, 0x32, 0x42, 0xd2, 0x42, 0x5d, 0xf2, 0x3b,
	0xaf, 0xa8, 0x96, 0x8a, 0x78, 0x9d, 0x78, 0xeb, 0x65, 0xe2, 0x02, 0x63, 0x25, 0xff, 0x81, 0xde,
	0xac, 0x0d, 0x36, 0x51, 0x93, 0x71, 0xb2, 0x63, 0x5c, 0x2c, 0xe6, 0x6f, 0xe0, 0xb5, 0x03, 0x8f,
	0xb2, 0x7d, 0x1e, 0x98, 0x87, 0x54, 0x6d, 0x92, 0xf0, 0xca, 0x71, 0x27, 0xd8, 0xd4, 0x96, 0x01,
	0xe7, 0x92, 0x4e, 0x15, 0xf4, 0xa5, 0x83, 0x8f, 0xb0, 0xf1, 0x62, 0x07, 0x13, 0xdd, 0xc9, 0x02,
	0x25, 0x64, 0x5d, 0xe4, 0x0b, 0x87, 0x1d, 0x67, 0xfb, 0x8b, 0x76, 0x50, 0x2c, 0x02, 0x25, 0x79,
	0x12, 0x58, 0xa5, 0xeb, 0xe8, 0x73, 0x87, 0x4e, 0xb1, 0x5d, 0xc5, 0x08, 0x02, 0x55, 0x06, 0x0b,
	0x31, 0x25, 0x78, 0xe6, 0x4d, 0xba, 0xa9, 0x31, 0x26, 0x07, 0x3c, 0xf5, 0x57, 0x21, 0xc6, 0x56,
	0x0d, 0x58, 0x77, 0xc0, 0x24, 0x83, 0x72, 0xc8, 0x29, 0x46, 0x82, 0x67, 0x12, 0x63, 0x82, 0x27,
	0xde, 0x79, 0x43, 0xd2, 0x08, 0x2b, 0x52, 0xca, 0x16, 0x94, 0xe6, 0xf0, 0xd8, 0x03, 0x78, 0xd2,
	0x8e, 0x44, 0x80, 0x96, 0xca, 0x08, 0x8f, 0xbc, 0x12, 0x42, 0xd2, 0xaa, 0xe8, 0x9c, 0xb1, 0xa8,
	0xe1, 0xa1, 0x17, 0xa0, 0xef, 0x8f, 0x28, 0xa5, 0x08, 0x1e, 0x78, 0xa3, 0x6a, 0x6a, 0x41, 0xd2,
	0xad, 0x58, 0xf9, 0x17, 0xee, 0x79, 0x5b, 0x56, 0x71, 0x6e, 0xa3, 0xef, 0x3a, 0x64, 0x96, 0x4d,
	0xd6, 0x90, 0x3c, 0x9b, 0x90, 0xa9, 0xb0, 0x68, 0x45, 0xde, 0xdf, 0x24, 0x86, 0x3b, 0x8e, 0x3d,
	0xcd, 0x66, 0x6a, 0xac, 0xb2, 0x21, 0xe9, 0xad, 0x4e, 0xdc, 0x76, 0x27, 0x4e, 0xb2, 0xe9, 0xda,
	0x89, 0x3a, 0xd7, 0xd7, 0x72, 0x6b, 0xbb, 0xb2, 0xf2, 0xc8, 0x9b, 0xca, 0x93, 0x18, 0x6e, 0x6e,
	0x17, 0xd5, 0xe9, 0xa8, 0xd1, 0x37, 0xbc, 0x6e, 0x57, 0x4e, 0x13, 0x0a, 0x8a, 0x38, 0x5c, 0xf7,
	0xee, 0x74, 0x2d, 0x6b, 0xc5, 0x5c, 0xf3, 0xd6, 0x38, 0x46, 0x21, 0x33, 0x8b, 0x66, 0xbe, 0xd0,
	0xdf, 0x14, 0x52, 0x98, 0x10, 0xd6, 0x57, 0x07, 0x07, 0xd7, 0x4a, 0x04, 0xa7, 0x4c, 0x98, 0xaa,
	0xc2, 0xab, 0xab, 0x5e, 0xae, 0x44, 0x06, 0xfd, 0x36, 0x14, 0xf7, 0xeb, 0xca, 0xea, 0xe0, 0xbd,
	0xd7, 0x28, 0xe7, 0x33, 0xdb, 0x69, 0x53, 0xf5, 0x56, 0x6d, 0xdc, 0xaf, 0xde, 0xd1, 0xa1, 0xf3,
	0xc3, 0xc5, 0xe3, 0x7f, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x88, 0x54, 0xab, 0x1a,
	0x06, 0x00, 0x00,
}
