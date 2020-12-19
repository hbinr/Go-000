// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_dto.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// 用户DTO UserDTO
type UserDTO struct {
	UserID               int64    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Passport             string   `protobuf:"bytes,2,opt,name=Passport,proto3" json:"Passport,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Nickname             string   `protobuf:"bytes,5,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	CreateTime           int64    `protobuf:"varint,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDTO) Reset()         { *m = UserDTO{} }
func (m *UserDTO) String() string { return proto.CompactTextString(m) }
func (*UserDTO) ProtoMessage()    {}
func (*UserDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_e188e24f0a46067a, []int{0}
}

func (m *UserDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDTO.Unmarshal(m, b)
}
func (m *UserDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDTO.Marshal(b, m, deterministic)
}
func (m *UserDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDTO.Merge(m, src)
}
func (m *UserDTO) XXX_Size() int {
	return xxx_messageInfo_UserDTO.Size(m)
}
func (m *UserDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDTO.DiscardUnknown(m)
}

var xxx_messageInfo_UserDTO proto.InternalMessageInfo

func (m *UserDTO) GetUseID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserDTO) GetPassport() string {
	if m != nil {
		return m.Passport
	}
	return ""
}

func (m *UserDTO) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserDTO) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDTO) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserDTO) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UserDTO) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserDTO)(nil), "pb.UserDTO")
}

func init() { proto.RegisterFile("user_dto.proto", fileDescriptor_e188e24f0a46067a) }

var fileDescriptor_e188e24f0a46067a = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0x8a, 0x4f, 0x29, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x3a,
	0xc9, 0xc8, 0xc5, 0x1e, 0x5a, 0x9c, 0x5a, 0xe4, 0x12, 0xe2, 0x2f, 0x24, 0xc2, 0xc5, 0x1a, 0x5a,
	0x9c, 0xea, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe1, 0x08, 0x49, 0x71, 0x71,
	0x04, 0x24, 0x16, 0x17, 0x17, 0xe4, 0x17, 0x95, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1,
	0xf9, 0x30, 0xb9, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x66, 0x84, 0x1c, 0x88, 0x0f, 0x32, 0xcd, 0x35,
	0x37, 0x31, 0x33, 0x47, 0x82, 0x05, 0x2c, 0x01, 0xe1, 0x80, 0x74, 0xf8, 0x65, 0x26, 0x67, 0xe7,
	0x25, 0xe6, 0xa6, 0x4a, 0xb0, 0x42, 0x74, 0xc0, 0xf8, 0x42, 0x72, 0x5c, 0x5c, 0xce, 0x45, 0xa9,
	0x89, 0x25, 0xa9, 0x21, 0x99, 0xb9, 0xa9, 0x12, 0x6c, 0x60, 0x47, 0x20, 0x89, 0x80, 0xe4, 0x43,
	0x0b, 0x52, 0x60, 0xf2, 0xec, 0x10, 0x79, 0x84, 0x48, 0x12, 0x1b, 0xd8, 0x5b, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x83, 0xbf, 0x08, 0xe8, 0x00, 0x00, 0x00,
}