// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta1/group.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/ovrclk/akash/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// State is an enum which refers to state of group
type Group_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	GroupStateInvalid Group_State = 0
	// GroupOpen denotes state for group open
	GroupOpen Group_State = 1
	// GroupOrdered denotes state for group ordered
	GroupOrdered Group_State = 2
	// GroupMatched denotes state for group matched
	GroupMatched Group_State = 3
	// GroupInsufficientFunds denotes state for group insufficient_funds
	GroupInsufficientFunds Group_State = 4
	// GroupClosed denotes state for group closed
	GroupClosed Group_State = 5
)

var Group_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "ordered",
	3: "matched",
	4: "insufficient_funds",
	5: "closed",
}

var Group_State_value = map[string]int32{
	"invalid":            0,
	"open":               1,
	"ordered":            2,
	"matched":            3,
	"insufficient_funds": 4,
	"closed":             5,
}

func (x Group_State) String() string {
	return proto.EnumName(Group_State_name, int32(x))
}

func (Group_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92581ef27257da99, []int{3, 0}
}

// MsgCloseGroup defines SDK message to close a single Group within a Deployment.
type MsgCloseGroup struct {
	ID GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseGroup) Reset()         { *m = MsgCloseGroup{} }
func (m *MsgCloseGroup) String() string { return proto.CompactTextString(m) }
func (*MsgCloseGroup) ProtoMessage()    {}
func (*MsgCloseGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_92581ef27257da99, []int{0}
}
func (m *MsgCloseGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseGroup.Merge(m, src)
}
func (m *MsgCloseGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseGroup proto.InternalMessageInfo

func (m *MsgCloseGroup) GetID() GroupID {
	if m != nil {
		return m.ID
	}
	return GroupID{}
}

// GroupID stores owner, deployment sequence number and group sequence number
type GroupID struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
}

func (m *GroupID) Reset()      { *m = GroupID{} }
func (*GroupID) ProtoMessage() {}
func (*GroupID) Descriptor() ([]byte, []int) {
	return fileDescriptor_92581ef27257da99, []int{1}
}
func (m *GroupID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupID.Merge(m, src)
}
func (m *GroupID) XXX_Size() int {
	return m.Size()
}
func (m *GroupID) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupID.DiscardUnknown(m)
}

var xxx_messageInfo_GroupID proto.InternalMessageInfo

func (m *GroupID) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GroupID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *GroupID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

// GroupSpec stores group specifications
type GroupSpec struct {
	Name             string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Requirements     []types.Attribute `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements" yaml:"requirements"`
	Resources        []Resource        `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources" yaml:"resources"`
	OrderBidDuration int64             `protobuf:"varint,4,opt,name=order_bid_duration,json=orderBidDuration,proto3" json:"order-bid-duration" yaml:"order-bid-duration"`
}

func (m *GroupSpec) Reset()         { *m = GroupSpec{} }
func (m *GroupSpec) String() string { return proto.CompactTextString(m) }
func (*GroupSpec) ProtoMessage()    {}
func (*GroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_92581ef27257da99, []int{2}
}
func (m *GroupSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSpec.Merge(m, src)
}
func (m *GroupSpec) XXX_Size() int {
	return m.Size()
}
func (m *GroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSpec proto.InternalMessageInfo

// Group stores group id, state and specifications of group
type Group struct {
	GroupID   GroupID     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"id" yaml:"id"`
	State     Group_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta1.Group_State" json:"state" yaml:"state"`
	GroupSpec GroupSpec   `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3" json:"spec" yaml:"spec"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_92581ef27257da99, []int{3}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGroupID() GroupID {
	if m != nil {
		return m.GroupID
	}
	return GroupID{}
}

func (m *Group) GetState() Group_State {
	if m != nil {
		return m.State
	}
	return GroupStateInvalid
}

func (m *Group) GetGroupSpec() GroupSpec {
	if m != nil {
		return m.GroupSpec
	}
	return GroupSpec{}
}

// Resource stores unit, total count and price of resource
type Resource struct {
	Resources types.ResourceUnits `protobuf:"bytes,1,opt,name=resources,proto3" json:"unit" yaml:"unit"`
	Count     uint32              `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price     types1.Coin         `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_92581ef27257da99, []int{4}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetResources() types.ResourceUnits {
	if m != nil {
		return m.Resources
	}
	return types.ResourceUnits{}
}

func (m *Resource) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Resource) GetPrice() types1.Coin {
	if m != nil {
		return m.Price
	}
	return types1.Coin{}
}

func init() {
	proto.RegisterEnum("akash.deployment.v1beta1.Group_State", Group_State_name, Group_State_value)
	proto.RegisterType((*MsgCloseGroup)(nil), "akash.deployment.v1beta1.MsgCloseGroup")
	proto.RegisterType((*GroupID)(nil), "akash.deployment.v1beta1.GroupID")
	proto.RegisterType((*GroupSpec)(nil), "akash.deployment.v1beta1.GroupSpec")
	proto.RegisterType((*Group)(nil), "akash.deployment.v1beta1.Group")
	proto.RegisterType((*Resource)(nil), "akash.deployment.v1beta1.Resource")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta1/group.proto", fileDescriptor_92581ef27257da99)
}

var fileDescriptor_92581ef27257da99 = []byte{
	// 873 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0x8e, 0x13, 0xa7, 0xbb, 0x99, 0xec, 0x82, 0x19, 0x7e, 0x34, 0xeb, 0x6a, 0x33, 0xee, 0x40,
	0xa5, 0x88, 0xaa, 0x89, 0x9a, 0xde, 0xf6, 0x86, 0x1b, 0xb1, 0xda, 0x43, 0x29, 0xf2, 0x0a, 0x0e,
	0x08, 0x29, 0x38, 0x9e, 0x59, 0xef, 0xa8, 0x89, 0xc7, 0x6b, 0x8f, 0x17, 0xf6, 0xce, 0xa1, 0xca,
	0x09, 0x71, 0xe2, 0x12, 0xa9, 0x12, 0x77, 0xfe, 0x8e, 0x1e, 0x7b, 0x84, 0x8b, 0x85, 0xb2, 0x17,
	0x94, 0x0b, 0x52, 0xfe, 0x02, 0x34, 0x33, 0x76, 0x9d, 0x40, 0xab, 0x15, 0xa7, 0xe4, 0x7d, 0xef,
	0xfb, 0xde, 0x3c, 0xbf, 0xf7, 0x79, 0x0c, 0x3e, 0xf1, 0x9f, 0xf9, 0xe9, 0xf9, 0x80, 0xd0, 0x78,
	0xca, 0xaf, 0x66, 0x34, 0x12, 0x83, 0xcb, 0x87, 0x13, 0x2a, 0xfc, 0x87, 0x83, 0x30, 0xe1, 0x59,
	0xdc, 0x8f, 0x13, 0x2e, 0x38, 0xec, 0x28, 0x56, 0xbf, 0x62, 0xf5, 0x0b, 0x96, 0xfd, 0x41, 0xc8,
	0x43, 0xae, 0x48, 0x03, 0xf9, 0x4f, 0xf3, 0xed, 0xbb, 0xba, 0xea, 0xc4, 0x4f, 0xe9, 0xeb, 0x7a,
	0x09, 0x4d, 0x79, 0x96, 0x04, 0xb4, 0xa0, 0xe0, 0x37, 0x50, 0x7c, 0x21, 0x12, 0x36, 0xc9, 0x44,
	0xc9, 0xe9, 0x06, 0x3c, 0x9d, 0xf1, 0x74, 0x9b, 0x14, 0x70, 0x16, 0xe9, 0x3c, 0x0e, 0xc1, 0xfe,
	0x93, 0x34, 0x7c, 0x3c, 0xe5, 0x29, 0x3d, 0x96, 0xdd, 0xc2, 0x2f, 0x41, 0x9d, 0x91, 0x8e, 0xe1,
	0x18, 0xbd, 0xf6, 0xf0, 0x6e, 0xff, 0x6d, 0x4d, 0xf7, 0x15, 0xf9, 0x64, 0xe4, 0x1e, 0xbe, 0xcc,
	0x51, 0x6d, 0x99, 0xa3, 0xfa, 0xc9, 0x68, 0x95, 0xa3, 0x3a, 0x23, 0xeb, 0x1c, 0xb5, 0xae, 0xfc,
	0xd9, 0xf4, 0x08, 0x33, 0x82, 0xbd, 0x3a, 0x23, 0x47, 0xe6, 0x5f, 0x2f, 0x50, 0x0d, 0xff, 0x66,
	0x80, 0x9d, 0x42, 0x04, 0x07, 0xa0, 0xc9, 0xbf, 0x8f, 0x68, 0xa2, 0x8e, 0x69, 0xb9, 0x07, 0xab,
	0x1c, 0x69, 0x60, 0x9d, 0xa3, 0x3d, 0x2d, 0x56, 0x21, 0xf6, 0x34, 0x0c, 0x1f, 0x01, 0x93, 0xa4,
	0xf4, 0xa2, 0x53, 0x77, 0x8c, 0x9e, 0xe9, 0xa2, 0x65, 0x8e, 0xcc, 0xd1, 0x29, 0xbd, 0x58, 0xe5,
	0x48, 0xe1, 0xeb, 0x1c, 0xb5, 0xb5, 0x4c, 0x46, 0xd8, 0x53, 0xa0, 0x14, 0x85, 0x52, 0xd4, 0x70,
	0x8c, 0xde, 0xbe, 0x16, 0x1d, 0x17, 0xa2, 0x70, 0x4b, 0x14, 0x6a, 0x91, 0xfc, 0x39, 0xda, 0xfd,
	0xe5, 0x05, 0xaa, 0xa9, 0x86, 0x7f, 0x6e, 0x80, 0x96, 0x6a, 0xf8, 0x34, 0xa6, 0x01, 0xbc, 0x0f,
	0xcc, 0xc8, 0x9f, 0xd1, 0xa2, 0xe3, 0xdb, 0xb2, 0x88, 0x8c, 0xab, 0x22, 0x32, 0xc2, 0x9e, 0x02,
	0x61, 0x04, 0xf6, 0x12, 0x7a, 0x91, 0xb1, 0x84, 0xca, 0x99, 0xa5, 0x9d, 0xba, 0xd3, 0xe8, 0xb5,
	0x87, 0x87, 0xc5, 0x34, 0xe5, 0x2a, 0x5e, 0xcf, 0xf1, 0xb3, 0x72, 0x5f, 0xee, 0x7d, 0x39, 0xc9,
	0x55, 0x8e, 0xb6, 0xa4, 0xeb, 0x1c, 0xbd, 0xaf, 0xeb, 0x6f, 0xa2, 0xd8, 0xdb, 0x22, 0xc1, 0x10,
	0xb4, 0x4a, 0x6b, 0xa4, 0x9d, 0x86, 0x3a, 0x0c, 0xbf, 0x7d, 0x75, 0x5e, 0x41, 0x75, 0xef, 0x15,
	0x27, 0x56, 0xe2, 0x75, 0x8e, 0xac, 0xf2, 0xb8, 0x02, 0xc2, 0x5e, 0x95, 0x86, 0x19, 0x80, 0x3c,
	0x21, 0x34, 0x19, 0x4f, 0x18, 0x19, 0x93, 0x2c, 0xf1, 0x05, 0xe3, 0x51, 0xc7, 0x74, 0x8c, 0x5e,
	0xc3, 0x3d, 0x5e, 0xe6, 0xc8, 0x7a, 0x2a, 0xb3, 0x2e, 0x23, 0xa3, 0x22, 0xb7, 0xca, 0x91, 0x56,
	0x3c, 0x98, 0x30, 0xf2, 0xa0, 0x54, 0xac, 0x73, 0x74, 0x50, 0xac, 0xf9, 0x3f, 0x39, 0xec, 0x59,
	0xfc, 0x5f, 0x45, 0x8e, 0x76, 0x9f, 0x97, 0x4b, 0xf9, 0xd1, 0x04, 0x4d, 0xed, 0xd3, 0xef, 0xc0,
	0xae, 0x7a, 0xbd, 0xc6, 0xff, 0xc7, 0xad, 0xb8, 0x70, 0x6b, 0xe9, 0xc4, 0x37, 0x59, 0x76, 0x47,
	0x95, 0x3d, 0x21, 0xf0, 0x6b, 0xd0, 0x4c, 0x85, 0x2f, 0xa8, 0x72, 0xdd, 0x3b, 0xc3, 0x7b, 0x37,
	0x94, 0xef, 0x9f, 0x4a, 0xb2, 0x36, 0xb3, 0xd2, 0x55, 0x66, 0x56, 0x21, 0xf6, 0x34, 0x0c, 0xc7,
	0x00, 0xe8, 0xce, 0xd3, 0x98, 0x06, 0xca, 0x9d, 0xed, 0xe1, 0xc7, 0x37, 0x14, 0x97, 0x1e, 0x74,
	0xef, 0x14, 0xfb, 0x32, 0xa5, 0xb0, 0x72, 0x9e, 0x8c, 0xb0, 0xd7, 0x0a, 0x4b, 0x1e, 0xfe, 0xc3,
	0x00, 0x4d, 0xd5, 0x0c, 0xc4, 0x60, 0x87, 0x45, 0x97, 0xfe, 0x94, 0x11, 0xab, 0x66, 0x7f, 0x38,
	0x5f, 0x38, 0xef, 0xe9, 0x6a, 0x32, 0x79, 0xa2, 0x13, 0xf0, 0x36, 0x30, 0x79, 0x4c, 0x23, 0xcb,
	0xb0, 0xf7, 0xe7, 0x0b, 0x47, 0x5b, 0xfe, 0x69, 0x4c, 0x23, 0x78, 0x08, 0x76, 0xd4, 0x26, 0x28,
	0xb1, 0xea, 0xb6, 0x35, 0x5f, 0x38, 0x7b, 0x3a, 0xa7, 0x31, 0x99, 0x9e, 0xf9, 0x22, 0x38, 0xa7,
	0xc4, 0x6a, 0x6c, 0xa4, 0x9f, 0x68, 0x0c, 0x0e, 0x01, 0x64, 0x51, 0x9a, 0x9d, 0x9d, 0xb1, 0x80,
	0xd1, 0x48, 0x8c, 0xcf, 0xb2, 0x88, 0xa4, 0x96, 0x69, 0xdb, 0xf3, 0x85, 0xf3, 0x91, 0x1e, 0xff,
	0x46, 0xfa, 0x73, 0x99, 0x85, 0x77, 0xc0, 0xad, 0x40, 0xde, 0x44, 0xc4, 0x6a, 0xda, 0xef, 0xce,
	0x17, 0x4e, 0x5b, 0xf1, 0xd4, 0xe5, 0x44, 0x6c, 0xf3, 0xf9, 0xaf, 0xdd, 0x5a, 0x71, 0x99, 0xfc,
	0x6d, 0x80, 0xdd, 0xd2, 0xc6, 0xf0, 0xdb, 0x4d, 0xf7, 0x6f, 0x5b, 0x61, 0xeb, 0x55, 0x2b, 0x05,
	0x5f, 0x45, 0x4c, 0xa4, 0xd5, 0x30, 0xb3, 0x88, 0x89, 0x6a, 0x98, 0x32, 0xda, 0xb2, 0xfc, 0x00,
	0x34, 0x03, 0x9e, 0x45, 0x42, 0xb9, 0x60, 0x5f, 0xaf, 0x57, 0x01, 0xd5, 0x7a, 0x55, 0x88, 0x3d,
	0x0d, 0xc3, 0x2f, 0x40, 0x33, 0x4e, 0x58, 0x40, 0x8b, 0xcd, 0x1e, 0xf4, 0xf5, 0x0d, 0xbc, 0xdd,
	0xcb, 0x63, 0xce, 0x22, 0x7d, 0x77, 0xca, 0x7a, 0x8a, 0x5f, 0xd5, 0x53, 0x21, 0xf6, 0x34, 0xac,
	0x9f, 0xd8, 0x1d, 0xbd, 0x5c, 0x76, 0x8d, 0x57, 0xcb, 0xae, 0xf1, 0xe7, 0xb2, 0x6b, 0xfc, 0x74,
	0xdd, 0xad, 0xbd, 0xba, 0xee, 0xd6, 0x7e, 0xbf, 0xee, 0xd6, 0xbe, 0xf9, 0x34, 0x64, 0xe2, 0x3c,
	0x9b, 0xf4, 0x03, 0x3e, 0x1b, 0xf0, 0xcb, 0x24, 0x98, 0x3e, 0x1b, 0xe8, 0xef, 0xc2, 0x0f, 0x9b,
	0x9f, 0x24, 0x71, 0x15, 0xd3, 0x74, 0x72, 0x4b, 0x5d, 0xfa, 0x8f, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xed, 0x4c, 0x8e, 0x60, 0xb3, 0x06, 0x00, 0x00,
}

func (m *MsgCloseGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GroupID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GSeq != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GroupSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderBidDuration != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.OrderBidDuration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Requirements) > 0 {
		for iNdEx := len(m.Requirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GroupSpec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.GroupID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Count != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCloseGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func (m *GroupID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovGroup(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovGroup(uint64(m.GSeq))
	}
	return n
}

func (m *GroupSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	if m.OrderBidDuration != 0 {
		n += 1 + sovGroup(uint64(m.OrderBidDuration))
	}
	return n
}

func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GroupID.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.State != 0 {
		n += 1 + sovGroup(uint64(m.State))
	}
	l = m.GroupSpec.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Resources.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.Count != 0 {
		n += 1 + sovGroup(uint64(m.Count))
	}
	l = m.Price.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCloseGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCloseGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, types.Attribute{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, Resource{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBidDuration", wireType)
			}
			m.OrderBidDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderBidDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Group_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
