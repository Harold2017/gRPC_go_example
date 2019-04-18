// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Authentication.proto

package authentication

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_119b0f6404520e77, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignupRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupRequest) Reset()         { *m = SignupRequest{} }
func (m *SignupRequest) String() string { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()    {}
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_119b0f6404520e77, []int{1}
}

func (m *SignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupRequest.Unmarshal(m, b)
}
func (m *SignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupRequest.Marshal(b, m, deterministic)
}
func (m *SignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupRequest.Merge(m, src)
}
func (m *SignupRequest) XXX_Size() int {
	return xxx_messageInfo_SignupRequest.Size(m)
}
func (m *SignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupRequest proto.InternalMessageInfo

func (m *SignupRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignupRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Time                 float32  `protobuf:"fixed32,2,opt,name=time,proto3" json:"time,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_119b0f6404520e77, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Response) GetTime() float32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type QueryRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_119b0f6404520e77, []int{3}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type QueryResponse struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Time                 float32  `protobuf:"fixed32,2,opt,name=time,proto3" json:"time,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_119b0f6404520e77, []int{4}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func (m *QueryResponse) GetTime() float32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *QueryResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "authentication.LoginRequest")
	proto.RegisterType((*SignupRequest)(nil), "authentication.SignupRequest")
	proto.RegisterType((*Response)(nil), "authentication.Response")
	proto.RegisterType((*QueryRequest)(nil), "authentication.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "authentication.QueryResponse")
}

func init() { proto.RegisterFile("Authentication.proto", fileDescriptor_119b0f6404520e77) }

var fileDescriptor_119b0f6404520e77 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x62, 0x42, 0x1d, 0xda, 0x52, 0x96, 0x22, 0x21, 0x58, 0x28, 0x39, 0x79, 0xca,
	0x41, 0x1f, 0x40, 0x0a, 0xe2, 0xa9, 0x1e, 0x8c, 0x67, 0x0f, 0xab, 0x0e, 0x75, 0xd1, 0xee, 0xc6,
	0x9d, 0x5d, 0xc4, 0x57, 0xf5, 0x69, 0x24, 0x9b, 0xdd, 0xd2, 0x8d, 0x58, 0xf0, 0x36, 0x7f, 0xfe,
	0xe1, 0x9b, 0xc9, 0xfc, 0x0b, 0x8b, 0xb5, 0x35, 0xaf, 0x28, 0x8d, 0x78, 0xe6, 0x46, 0x28, 0x59,
	0xb7, 0x5a, 0x19, 0xc5, 0x66, 0x3c, 0xfa, 0x5a, 0xdd, 0xc2, 0x64, 0xa3, 0xb6, 0x42, 0x36, 0xf8,
	0x61, 0x91, 0x0c, 0x2b, 0x61, 0x6c, 0x09, 0xb5, 0xe4, 0x3b, 0x2c, 0x46, 0xab, 0xd1, 0xc5, 0x69,
	0xb3, 0xd7, 0x9d, 0xd7, 0x72, 0xa2, 0x4f, 0xa5, 0x5f, 0x8a, 0xa4, 0xf7, 0x82, 0xae, 0x1e, 0x61,
	0xfa, 0x20, 0xb6, 0xd2, 0xb6, 0x01, 0xb4, 0x80, 0x0c, 0x77, 0x5c, 0xbc, 0x7b, 0x4a, 0x2f, 0x22,
	0x7c, 0x72, 0x04, 0x9f, 0x0e, 0xf0, 0x1b, 0x18, 0x37, 0x48, 0xad, 0x92, 0x84, 0x1d, 0xd9, 0xa8,
	0x37, 0x94, 0x81, 0xec, 0x04, 0x63, 0x70, 0x62, 0x84, 0xa7, 0x26, 0x8d, 0xab, 0xd9, 0x19, 0xe4,
	0x64, 0xb8, 0xb1, 0xe4, 0x78, 0x59, 0xe3, 0x55, 0xb5, 0x82, 0xc9, 0xbd, 0x45, 0xfd, 0x15, 0x76,
	0x9d, 0x43, 0xca, 0x5b, 0xe1, 0x79, 0x5d, 0x59, 0xdd, 0xc1, 0xd4, 0x77, 0xf8, 0xa1, 0x73, 0x48,
	0x35, 0x52, 0x68, 0xd1, 0x48, 0xff, 0x19, 0x78, 0xf9, 0x3d, 0x82, 0x59, 0x1c, 0x07, 0xbb, 0x86,
	0xcc, 0x1d, 0x9e, 0x9d, 0xd7, 0x71, 0x24, 0xf5, 0x61, 0x1e, 0x65, 0x31, 0x74, 0xf7, 0x1b, 0xad,
	0x21, 0xef, 0x2f, 0xce, 0x96, 0xc3, 0x9e, 0x28, 0x89, 0x23, 0x88, 0x1b, 0xc8, 0xdc, 0x5f, 0xfe,
	0xde, 0xe1, 0xf0, 0x3c, 0xe5, 0xf2, 0x0f, 0xb7, 0xa7, 0x3c, 0xe5, 0xee, 0x65, 0x5d, 0xfd, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x58, 0x2d, 0xdd, 0x32, 0x71, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*Response, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	Login(context.Context, *LoginRequest) (*Response, error)
	Signup(context.Context, *SignupRequest) (*Response, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _Authentication_Signup_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Authentication_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Authentication.proto",
}