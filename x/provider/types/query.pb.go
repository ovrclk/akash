// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/provider/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryProvidersRequest is request type for the Query/Providers RPC method
type QueryProvidersRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryProvidersRequest) Reset()         { *m = QueryProvidersRequest{} }
func (m *QueryProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProvidersRequest) ProtoMessage()    {}
func (*QueryProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_236609e2c9e515d5, []int{0}
}
func (m *QueryProvidersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProvidersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProvidersRequest.Merge(m, src)
}
func (m *QueryProvidersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProvidersRequest proto.InternalMessageInfo

func (m *QueryProvidersRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryProvidersResponse is response type for the Query/Providers RPC method
type QueryProvidersResponse struct {
	Providers  Providers           `protobuf:"bytes,1,rep,name=providers,proto3,castrepeated=Providers" json:"providers"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryProvidersResponse) Reset()         { *m = QueryProvidersResponse{} }
func (m *QueryProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProvidersResponse) ProtoMessage()    {}
func (*QueryProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_236609e2c9e515d5, []int{1}
}
func (m *QueryProvidersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProvidersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProvidersResponse.Merge(m, src)
}
func (m *QueryProvidersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProvidersResponse proto.InternalMessageInfo

func (m *QueryProvidersResponse) GetProviders() Providers {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *QueryProvidersResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryProviderRequest is request type for the Query/Provider RPC method
type QueryProviderRequest struct {
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
}

func (m *QueryProviderRequest) Reset()         { *m = QueryProviderRequest{} }
func (m *QueryProviderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProviderRequest) ProtoMessage()    {}
func (*QueryProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_236609e2c9e515d5, []int{2}
}
func (m *QueryProviderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProviderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProviderRequest.Merge(m, src)
}
func (m *QueryProviderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProviderRequest proto.InternalMessageInfo

func (m *QueryProviderRequest) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

// QueryProviderResponse is response type for the Query/Provider RPC method
type QueryProviderResponse struct {
	Provider Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider"`
}

func (m *QueryProviderResponse) Reset()         { *m = QueryProviderResponse{} }
func (m *QueryProviderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProviderResponse) ProtoMessage()    {}
func (*QueryProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_236609e2c9e515d5, []int{3}
}
func (m *QueryProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProviderResponse.Merge(m, src)
}
func (m *QueryProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProviderResponse proto.InternalMessageInfo

func (m *QueryProviderResponse) GetProvider() Provider {
	if m != nil {
		return m.Provider
	}
	return Provider{}
}

func init() {
	proto.RegisterType((*QueryProvidersRequest)(nil), "akash.provider.v1beta1.QueryProvidersRequest")
	proto.RegisterType((*QueryProvidersResponse)(nil), "akash.provider.v1beta1.QueryProvidersResponse")
	proto.RegisterType((*QueryProviderRequest)(nil), "akash.provider.v1beta1.QueryProviderRequest")
	proto.RegisterType((*QueryProviderResponse)(nil), "akash.provider.v1beta1.QueryProviderResponse")
}

func init() {
	proto.RegisterFile("akash/provider/v1beta1/query.proto", fileDescriptor_236609e2c9e515d5)
}

var fileDescriptor_236609e2c9e515d5 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xb1, 0x6e, 0xd3, 0x40,
	0x1c, 0xc6, 0x7d, 0x81, 0xa2, 0xf6, 0xca, 0xc2, 0xa9, 0x54, 0x55, 0x84, 0x9c, 0x60, 0x04, 0xa4,
	0x80, 0xef, 0xe4, 0xf0, 0x04, 0xf5, 0x40, 0xd7, 0xd6, 0x23, 0x0c, 0xd5, 0xd9, 0x3e, 0xb9, 0x56,
	0x5a, 0x9f, 0xeb, 0x73, 0x02, 0x15, 0x62, 0xe1, 0x09, 0x90, 0x10, 0x0b, 0x8f, 0xc0, 0xcc, 0xc8,
	0x03, 0x74, 0x8c, 0xc4, 0xc2, 0x14, 0x50, 0xc2, 0x53, 0x30, 0x21, 0xdf, 0x9d, 0xed, 0x24, 0x4a,
	0xe4, 0x4c, 0x89, 0xec, 0xef, 0xfb, 0xfe, 0xbf, 0xfb, 0xfe, 0x67, 0x68, 0xd1, 0x01, 0x15, 0xe7,
	0x24, 0xcd, 0xf8, 0x28, 0x0e, 0x59, 0x46, 0x46, 0x8e, 0xcf, 0x72, 0xea, 0x90, 0xab, 0x21, 0xcb,
	0xae, 0x71, 0x9a, 0xf1, 0x9c, 0xa3, 0x7d, 0xa9, 0xc1, 0xa5, 0x06, 0x6b, 0x4d, 0x7b, 0x2f, 0xe2,
	0x11, 0x97, 0x12, 0x52, 0xfc, 0x53, 0xea, 0xf6, 0x83, 0x88, 0xf3, 0xe8, 0x82, 0x11, 0x9a, 0xc6,
	0x84, 0x26, 0x09, 0xcf, 0x69, 0x1e, 0xf3, 0x44, 0xe8, 0xb7, 0xcf, 0x02, 0x2e, 0x2e, 0xb9, 0x20,
	0x3e, 0x15, 0x4c, 0x0d, 0xa9, 0x46, 0xa6, 0x34, 0x8a, 0x13, 0x29, 0xd6, 0xda, 0xc7, 0x6b, 0xd8,
	0x2a, 0x10, 0x29, 0xb3, 0xce, 0xe0, 0xfd, 0xd3, 0x22, 0xe8, 0x44, 0x3f, 0x16, 0x1e, 0xbb, 0x1a,
	0x32, 0x91, 0xa3, 0x57, 0x10, 0xd6, 0x99, 0x07, 0xa0, 0x0b, 0x7a, 0xbb, 0xfd, 0x27, 0x58, 0x01,
	0xe0, 0x02, 0x00, 0xab, 0x53, 0xea, 0x5c, 0x7c, 0x42, 0x23, 0xa6, 0xbd, 0xde, 0x9c, 0xd3, 0xfa,
	0x0e, 0xe0, 0xfe, 0xf2, 0x04, 0x91, 0xf2, 0x44, 0x30, 0x74, 0x0a, 0x77, 0x4a, 0x1a, 0x71, 0x00,
	0xba, 0xb7, 0x7a, 0xbb, 0xfd, 0x2e, 0x5e, 0x5d, 0x17, 0x2e, 0xdd, 0xee, 0xbd, 0x9b, 0x49, 0xc7,
	0xf8, 0xf6, 0xbb, 0xb3, 0x53, 0xe7, 0xd5, 0x29, 0xe8, 0x78, 0x81, 0xba, 0x25, 0xa9, 0x9f, 0x36,
	0x52, 0x2b, 0x9e, 0x05, 0xec, 0x33, 0xb8, 0xb7, 0x40, 0x5d, 0xd6, 0x72, 0x0c, 0xb7, 0xf8, 0xdb,
	0x84, 0x65, 0xb2, 0x91, 0xbb, 0xae, 0xf3, 0x6f, 0xd2, 0xb1, 0xa3, 0x38, 0x3f, 0x1f, 0xfa, 0x38,
	0xe0, 0x97, 0x44, 0x2f, 0x48, 0xfd, 0xd8, 0x22, 0x1c, 0x90, 0xfc, 0x3a, 0x65, 0x02, 0x1f, 0x05,
	0xc1, 0x51, 0x18, 0x66, 0x4c, 0x08, 0x4f, 0xf9, 0xad, 0x37, 0x4b, 0xc5, 0x57, 0xad, 0xb8, 0x70,
	0xbb, 0x3c, 0x8f, 0xae, 0xbd, 0xb9, 0x94, 0xdb, 0x45, 0x29, 0x5e, 0xe5, 0xeb, 0xff, 0x68, 0xc1,
	0x2d, 0x99, 0x8e, 0xbe, 0x00, 0x58, 0x37, 0x85, 0xec, 0x75, 0x49, 0x2b, 0xef, 0x40, 0x1b, 0x6f,
	0x2a, 0x57, 0xe8, 0xd6, 0xe1, 0xc7, 0x9f, 0x7f, 0x3f, 0xb7, 0x1e, 0xa1, 0x87, 0xa4, 0xe1, 0xf2,
	0x09, 0xf4, 0x15, 0xc0, 0xed, 0x32, 0x00, 0xbd, 0xd8, 0x68, 0x4e, 0x49, 0x65, 0x6f, 0xa8, 0xd6,
	0x50, 0x8e, 0x84, 0x7a, 0x8e, 0x0e, 0x1b, 0xa1, 0xc8, 0x7b, 0xb9, 0x9a, 0x0f, 0xae, 0x7b, 0x33,
	0x35, 0xc1, 0x78, 0x6a, 0x82, 0x3f, 0x53, 0x13, 0x7c, 0x9a, 0x99, 0xc6, 0x78, 0x66, 0x1a, 0xbf,
	0x66, 0xa6, 0xf1, 0xba, 0x37, 0xb7, 0x6b, 0x3e, 0xca, 0x82, 0x8b, 0x81, 0x4e, 0x7d, 0x57, 0xe7,
	0xca, 0x8d, 0xfb, 0x77, 0xe4, 0xf7, 0xf5, 0xf2, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xdd,
	0xef, 0x05, 0x24, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Providers queries providers
	Providers(ctx context.Context, in *QueryProvidersRequest, opts ...grpc.CallOption) (*QueryProvidersResponse, error)
	// Provider queries provider details
	Provider(ctx context.Context, in *QueryProviderRequest, opts ...grpc.CallOption) (*QueryProviderResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Providers(ctx context.Context, in *QueryProvidersRequest, opts ...grpc.CallOption) (*QueryProvidersResponse, error) {
	out := new(QueryProvidersResponse)
	err := c.cc.Invoke(ctx, "/akash.provider.v1beta1.Query/Providers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Provider(ctx context.Context, in *QueryProviderRequest, opts ...grpc.CallOption) (*QueryProviderResponse, error) {
	out := new(QueryProviderResponse)
	err := c.cc.Invoke(ctx, "/akash.provider.v1beta1.Query/Provider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Providers queries providers
	Providers(context.Context, *QueryProvidersRequest) (*QueryProvidersResponse, error)
	// Provider queries provider details
	Provider(context.Context, *QueryProviderRequest) (*QueryProviderResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Providers(ctx context.Context, req *QueryProvidersRequest) (*QueryProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Providers not implemented")
}
func (*UnimplementedQueryServer) Provider(ctx context.Context, req *QueryProviderRequest) (*QueryProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provider not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Providers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Providers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.provider.v1beta1.Query/Providers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Providers(ctx, req.(*QueryProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Provider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Provider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.provider.v1beta1.Query/Provider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Provider(ctx, req.(*QueryProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.provider.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Providers",
			Handler:    _Query_Providers_Handler,
		},
		{
			MethodName: "Provider",
			Handler:    _Query_Provider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/provider/v1beta1/query.proto",
}

func (m *QueryProvidersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProvidersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProvidersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProvidersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProvidersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProvidersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Providers) > 0 {
		for iNdEx := len(m.Providers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Providers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryProviderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProviderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProviderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Provider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryProvidersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProvidersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Providers) > 0 {
		for _, e := range m.Providers {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProviderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Provider.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProvidersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProvidersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProvidersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryProvidersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProvidersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProvidersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Providers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Providers = append(m.Providers, Provider{})
			if err := m.Providers[len(m.Providers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryProviderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProviderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProviderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryProviderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Provider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
