// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hetu/checkpointing/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_hetu_project_hetu_v1_crypto_bls12381 "github.com/hetu-project/hetu-checkpoint/crypto/bls12381"
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

// MsgRegistValidator defines a message to regist a validator
type MsgRegistValidator struct {
	// pubkey is the BLS public key of a validator
	BlsPubkey *github_com_hetu_project_hetu_v1_crypto_bls12381.PublicKey `protobuf:"bytes,1,opt,name=bls_pubkey,json=blsPubkey,proto3,customtype=github.com/hetu-project/hetu-checkpoint/crypto/bls12381.PublicKey" json:"bls_pubkey,omitempty"`
	// validator_address is the address of the validator
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// sender is the cosmos bech32 address from the owner of the given Cosmos coins
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgRegistValidator) Reset()         { *m = MsgRegistValidator{} }
func (m *MsgRegistValidator) String() string { return proto.CompactTextString(m) }
func (*MsgRegistValidator) ProtoMessage()    {}
func (*MsgRegistValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_74c41b14ed2cffe3, []int{0}
}
func (m *MsgRegistValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegistValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegistValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegistValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegistValidator.Merge(m, src)
}
func (m *MsgRegistValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegistValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegistValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegistValidator proto.InternalMessageInfo

// MsgRegistValidatorResponse defines the MsgRegistValidator
// response type
type MsgRegistValidatorResponse struct {
}

func (m *MsgRegistValidatorResponse) Reset()         { *m = MsgRegistValidatorResponse{} }
func (m *MsgRegistValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegistValidatorResponse) ProtoMessage()    {}
func (*MsgRegistValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74c41b14ed2cffe3, []int{1}
}
func (m *MsgRegistValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegistValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegistValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegistValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegistValidatorResponse.Merge(m, src)
}
func (m *MsgRegistValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegistValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegistValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegistValidatorResponse proto.InternalMessageInfo

// MsgRegistStakeContract defines a message to register a validator's staking contract address
type MsgRegistStakeContract struct {
	// contract_address is the staking contract address of the validator
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// sender is the cosmos bech32 address from the owner of the given Cosmos coins
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgRegistStakeContract) Reset()         { *m = MsgRegistStakeContract{} }
func (m *MsgRegistStakeContract) String() string { return proto.CompactTextString(m) }
func (*MsgRegistStakeContract) ProtoMessage()    {}
func (*MsgRegistStakeContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_74c41b14ed2cffe3, []int{2}
}
func (m *MsgRegistStakeContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegistStakeContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegistStakeContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegistStakeContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegistStakeContract.Merge(m, src)
}
func (m *MsgRegistStakeContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegistStakeContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegistStakeContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegistStakeContract proto.InternalMessageInfo

// MsgRegistStakeContractResponse defines the MsgRegistStakeContract response type
type MsgRegistStakeContractResponse struct {
}

func (m *MsgRegistStakeContractResponse) Reset()         { *m = MsgRegistStakeContractResponse{} }
func (m *MsgRegistStakeContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegistStakeContractResponse) ProtoMessage()    {}
func (*MsgRegistStakeContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74c41b14ed2cffe3, []int{3}
}
func (m *MsgRegistStakeContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegistStakeContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegistStakeContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegistStakeContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegistStakeContractResponse.Merge(m, src)
}
func (m *MsgRegistStakeContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegistStakeContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegistStakeContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegistStakeContractResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegistValidator)(nil), "hetu.checkpointing.v1.MsgRegistValidator")
	proto.RegisterType((*MsgRegistValidatorResponse)(nil), "hetu.checkpointing.v1.MsgRegistValidatorResponse")
	proto.RegisterType((*MsgRegistStakeContract)(nil), "hetu.checkpointing.v1.MsgRegistStakeContract")
	proto.RegisterType((*MsgRegistStakeContractResponse)(nil), "hetu.checkpointing.v1.MsgRegistStakeContractResponse")
}

func init() { proto.RegisterFile("hetu/checkpointing/v1/tx.proto", fileDescriptor_74c41b14ed2cffe3) }

var fileDescriptor_74c41b14ed2cffe3 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0x14, 0x41,
	0x10, 0x9d, 0x4e, 0x30, 0x90, 0x46, 0x48, 0xec, 0x68, 0x5c, 0x06, 0xe9, 0x5d, 0xf6, 0x94, 0x44,
	0x32, 0xc3, 0x24, 0x04, 0x54, 0xf0, 0x60, 0x3c, 0x4a, 0x30, 0xac, 0xe0, 0x41, 0x84, 0x30, 0xdd,
	0xd3, 0xf4, 0x8e, 0xf3, 0xd1, 0xcd, 0x54, 0xcf, 0x92, 0xc1, 0x8b, 0x78, 0xf2, 0xe8, 0x4f, 0xc8,
	0x4f, 0xc8, 0xcf, 0xf0, 0xe0, 0x21, 0x47, 0xf1, 0x20, 0xb2, 0x7b, 0x88, 0xbf, 0xc1, 0x93, 0xcc,
	0xd7, 0x4a, 0x76, 0x17, 0x56, 0x6f, 0x55, 0xf5, 0x1e, 0xef, 0xbd, 0x9a, 0xa9, 0xc6, 0x74, 0x28,
	0x4c, 0xee, 0xf2, 0xa1, 0xe0, 0x91, 0x56, 0x61, 0x6a, 0xc2, 0x54, 0xba, 0x23, 0xcf, 0x35, 0xe7,
	0x8e, 0xce, 0x94, 0x51, 0xe4, 0x5e, 0x89, 0x3b, 0x37, 0x70, 0x67, 0xe4, 0xd9, 0x77, 0xa5, 0x92,
	0xaa, 0x62, 0xb8, 0x65, 0x55, 0x93, 0xed, 0x2e, 0x57, 0x90, 0x28, 0x70, 0xc1, 0xf8, 0x51, 0x2d,
	0xc4, 0x84, 0xf1, 0xff, 0xaa, 0xd9, 0xf7, 0x1b, 0x42, 0x02, 0x95, 0x4b, 0x02, 0xb2, 0x06, 0xfa,
	0x5f, 0x11, 0x26, 0x27, 0x20, 0x07, 0x42, 0x86, 0x60, 0x5e, 0xfb, 0x71, 0x18, 0xf8, 0x46, 0x65,
	0xe4, 0x2d, 0xc6, 0x2c, 0x86, 0x33, 0x9d, 0xb3, 0x48, 0x14, 0x1d, 0xd4, 0x43, 0x3b, 0xb7, 0x8f,
	0x9f, 0x7e, 0xff, 0xd1, 0x7d, 0x2c, 0x43, 0x33, 0xcc, 0x99, 0xc3, 0x55, 0xe2, 0x96, 0x01, 0xf7,
	0x75, 0xa6, 0xde, 0x09, 0x6e, 0xaa, 0xa6, 0x54, 0xe6, 0x59, 0xa1, 0x8d, 0x72, 0x59, 0x0c, 0xde,
	0xc1, 0xe1, 0x23, 0xcf, 0x39, 0xcd, 0x59, 0x1c, 0xf2, 0x17, 0xa2, 0x18, 0xac, 0xb3, 0x18, 0x4e,
	0x2b, 0x3d, 0xf2, 0x10, 0xdf, 0x19, 0xb5, 0x56, 0x67, 0x7e, 0x10, 0x64, 0x02, 0xa0, 0xb3, 0xd2,
	0x43, 0x3b, 0xeb, 0x83, 0xcd, 0x29, 0xf0, 0xac, 0x9e, 0x93, 0x6d, 0xbc, 0x06, 0x22, 0x0d, 0x44,
	0xd6, 0x59, 0xad, 0x18, 0x4d, 0xf7, 0x64, 0xeb, 0xd3, 0x45, 0xd7, 0xfa, 0x75, 0xd1, 0xb5, 0x3e,
	0x5e, 0x5f, 0xee, 0x35, 0xc3, 0xfe, 0x03, 0x6c, 0xcf, 0x6f, 0x33, 0x10, 0xa0, 0x55, 0x0a, 0xa2,
	0xaf, 0xf1, 0xf6, 0x14, 0x7d, 0x65, 0xfc, 0x48, 0x3c, 0x57, 0xa9, 0xc9, 0x7c, 0x6e, 0xc8, 0x2e,
	0xde, 0xe4, 0x4d, 0x3d, 0x0d, 0x84, 0x2a, 0xbb, 0x8d, 0x76, 0x3e, 0x9f, 0x67, 0x65, 0x79, 0x9e,
	0x1e, 0xa6, 0x8b, 0x1d, 0xdb, 0x4c, 0x07, 0xbf, 0x11, 0x5e, 0x3d, 0x01, 0x49, 0x14, 0xde, 0x98,
	0xfd, 0x09, 0xbb, 0xce, 0xc2, 0x1b, 0x70, 0xe6, 0x37, 0xb4, 0xbd, 0x7f, 0xa6, 0xb6, 0xc6, 0xe4,
	0x3d, 0xde, 0x5a, 0xf4, 0x25, 0xf6, 0x97, 0x29, 0xdd, 0xa0, 0xdb, 0x47, 0xff, 0x45, 0x6f, 0xcd,
	0xed, 0x5b, 0x1f, 0xae, 0x2f, 0xf7, 0xd0, 0xf1, 0xcb, 0x2f, 0x63, 0x8a, 0xae, 0xc6, 0x14, 0xfd,
	0x1c, 0x53, 0xf4, 0x79, 0x42, 0xad, 0xab, 0x09, 0xb5, 0xbe, 0x4d, 0xa8, 0xf5, 0xe6, 0x68, 0xd9,
	0xa1, 0x9d, 0xcf, 0xbc, 0x1d, 0x53, 0x68, 0x01, 0x6c, 0xad, 0xba, 0xea, 0xc3, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0x7b, 0x4e, 0xc7, 0x5e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RegistValidator defines a method for registering a new validator of the checkpoint network
	RegistValidator(ctx context.Context, in *MsgRegistValidator, opts ...grpc.CallOption) (*MsgRegistValidatorResponse, error)
	// RegistStakeContract defines a method for registering a validator's staking contract address
	RegistStakeContract(ctx context.Context, in *MsgRegistStakeContract, opts ...grpc.CallOption) (*MsgRegistStakeContractResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegistValidator(ctx context.Context, in *MsgRegistValidator, opts ...grpc.CallOption) (*MsgRegistValidatorResponse, error) {
	out := new(MsgRegistValidatorResponse)
	err := c.cc.Invoke(ctx, "/hetu.checkpointing.v1.Msg/RegistValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegistStakeContract(ctx context.Context, in *MsgRegistStakeContract, opts ...grpc.CallOption) (*MsgRegistStakeContractResponse, error) {
	out := new(MsgRegistStakeContractResponse)
	err := c.cc.Invoke(ctx, "/hetu.checkpointing.v1.Msg/RegistStakeContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegistValidator defines a method for registering a new validator of the checkpoint network
	RegistValidator(context.Context, *MsgRegistValidator) (*MsgRegistValidatorResponse, error)
	// RegistStakeContract defines a method for registering a validator's staking contract address
	RegistStakeContract(context.Context, *MsgRegistStakeContract) (*MsgRegistStakeContractResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegistValidator(ctx context.Context, req *MsgRegistValidator) (*MsgRegistValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistValidator not implemented")
}
func (*UnimplementedMsgServer) RegistStakeContract(ctx context.Context, req *MsgRegistStakeContract) (*MsgRegistStakeContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistStakeContract not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegistValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegistValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegistValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hetu.checkpointing.v1.Msg/RegistValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegistValidator(ctx, req.(*MsgRegistValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegistStakeContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegistStakeContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegistStakeContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hetu.checkpointing.v1.Msg/RegistStakeContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegistStakeContract(ctx, req.(*MsgRegistStakeContract))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hetu.checkpointing.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistValidator",
			Handler:    _Msg_RegistValidator_Handler,
		},
		{
			MethodName: "RegistStakeContract",
			Handler:    _Msg_RegistStakeContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hetu/checkpointing/v1/tx.proto",
}

func (m *MsgRegistValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegistValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegistValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlsPubkey != nil {
		{
			size := m.BlsPubkey.Size()
			i -= size
			if _, err := m.BlsPubkey.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegistValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegistValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegistValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRegistStakeContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegistStakeContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegistStakeContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegistStakeContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegistStakeContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegistStakeContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegistValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlsPubkey != nil {
		l = m.BlsPubkey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegistValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRegistStakeContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegistStakeContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegistValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegistValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegistValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hetu_project_hetu_v1_crypto_bls12381.PublicKey
			m.BlsPubkey = &v
			if err := m.BlsPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegistValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegistValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegistValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegistStakeContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegistStakeContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegistStakeContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegistStakeContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegistStakeContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegistStakeContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
