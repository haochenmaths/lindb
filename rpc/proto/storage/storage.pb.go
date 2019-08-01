// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Replica struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Replica) Reset()         { *m = Replica{} }
func (m *Replica) String() string { return proto.CompactTextString(m) }
func (*Replica) ProtoMessage()    {}
func (*Replica) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}
func (m *Replica) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Replica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Replica.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Replica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Replica.Merge(m, src)
}
func (m *Replica) XXX_Size() int {
	return m.Size()
}
func (m *Replica) XXX_DiscardUnknown() {
	xxx_messageInfo_Replica.DiscardUnknown(m)
}

var xxx_messageInfo_Replica proto.InternalMessageInfo

func (m *Replica) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Replica) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteRequest struct {
	Replicas             []*Replica `protobuf:"bytes,3,rep,name=replicas,proto3" json:"replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}
func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return m.Size()
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetReplicas() []*Replica {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type WriteResponse struct {
	// Types that are valid to be assigned to Seq:
	//	*WriteResponse_ResetSeq
	//	*WriteResponse_AckSeq
	Seq                  isWriteResponse_Seq `protobuf_oneof:"seq"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}
func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return m.Size()
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type isWriteResponse_Seq interface {
	isWriteResponse_Seq()
	MarshalTo([]byte) (int, error)
	Size() int
}

type WriteResponse_ResetSeq struct {
	ResetSeq int64 `protobuf:"varint,1,opt,name=resetSeq,proto3,oneof"`
}
type WriteResponse_AckSeq struct {
	AckSeq int64 `protobuf:"varint,2,opt,name=ackSeq,proto3,oneof"`
}

func (*WriteResponse_ResetSeq) isWriteResponse_Seq() {}
func (*WriteResponse_AckSeq) isWriteResponse_Seq()   {}

func (m *WriteResponse) GetSeq() isWriteResponse_Seq {
	if m != nil {
		return m.Seq
	}
	return nil
}

func (m *WriteResponse) GetResetSeq() int64 {
	if x, ok := m.GetSeq().(*WriteResponse_ResetSeq); ok {
		return x.ResetSeq
	}
	return 0
}

func (m *WriteResponse) GetAckSeq() int64 {
	if x, ok := m.GetSeq().(*WriteResponse_AckSeq); ok {
		return x.AckSeq
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WriteResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WriteResponse_OneofMarshaler, _WriteResponse_OneofUnmarshaler, _WriteResponse_OneofSizer, []interface{}{
		(*WriteResponse_ResetSeq)(nil),
		(*WriteResponse_AckSeq)(nil),
	}
}

func _WriteResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WriteResponse)
	// seq
	switch x := m.Seq.(type) {
	case *WriteResponse_ResetSeq:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.ResetSeq))
	case *WriteResponse_AckSeq:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.AckSeq))
	case nil:
	default:
		return fmt.Errorf("WriteResponse.Seq has unexpected type %T", x)
	}
	return nil
}

func _WriteResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WriteResponse)
	switch tag {
	case 1: // seq.resetSeq
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Seq = &WriteResponse_ResetSeq{int64(x)}
		return true, err
	case 2: // seq.ackSeq
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Seq = &WriteResponse_AckSeq{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _WriteResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WriteResponse)
	// seq
	switch x := m.Seq.(type) {
	case *WriteResponse_ResetSeq:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ResetSeq))
	case *WriteResponse_AckSeq:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.AckSeq))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type QueryRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{4}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Replica)(nil), "storage.Replica")
	proto.RegisterType((*WriteRequest)(nil), "storage.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "storage.WriteResponse")
	proto.RegisterType((*QueryRequest)(nil), "storage.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "storage.QueryResponse")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x9b, 0xad, 0xbb, 0xab, 0x63, 0x0b, 0x25, 0xa0, 0x14, 0x91, 0x52, 0x73, 0xea, 0x41,
	0x56, 0x59, 0xaf, 0x7b, 0xda, 0x93, 0x07, 0x11, 0xcc, 0x1e, 0x3c, 0xc7, 0x3a, 0x2c, 0xc5, 0x3f,
	0xed, 0x26, 0x59, 0xc1, 0x37, 0xf1, 0x91, 0x3c, 0xfa, 0x08, 0x52, 0x5f, 0x44, 0x92, 0xa6, 0x25,
	0xba, 0xb7, 0xf9, 0x32, 0x33, 0xbf, 0xef, 0x1b, 0x02, 0xb1, 0xd2, 0xb5, 0x14, 0x6b, 0x9c, 0x35,
	0xb2, 0xd6, 0x35, 0x9d, 0x3a, 0xc9, 0x2e, 0x60, 0xca, 0xb1, 0x79, 0xae, 0x4a, 0x41, 0x13, 0x08,
	0x15, 0x6e, 0x52, 0x92, 0x93, 0x22, 0xe4, 0xa6, 0xa4, 0x14, 0xf6, 0x1e, 0x85, 0x16, 0xe9, 0x28,
	0x27, 0x45, 0xc4, 0x6d, 0xcd, 0x16, 0x10, 0xdd, 0xcb, 0x4a, 0x23, 0xc7, 0xcd, 0x16, 0x95, 0xa6,
	0xe7, 0xb0, 0x2f, 0x3b, 0x80, 0x4a, 0xc3, 0x3c, 0x2c, 0x0e, 0xe7, 0xc9, 0xac, 0xf7, 0x72, 0x64,
	0x3e, 0x4c, 0xb0, 0x5b, 0x88, 0xdd, 0xb6, 0x6a, 0xea, 0x57, 0x85, 0xf4, 0xd4, 0xac, 0x2b, 0xd4,
	0xab, 0xde, 0xf9, 0x3a, 0xe0, 0xc3, 0x0b, 0x4d, 0x61, 0x22, 0xca, 0x27, 0xd3, 0x1b, 0xb9, 0x9e,
	0xd3, 0xcb, 0xb1, 0x0d, 0xcb, 0x72, 0x88, 0xee, 0xb6, 0x28, 0xdf, 0xfb, 0x34, 0x09, 0x84, 0x2f,
	0x6a, 0x6d, 0x49, 0x07, 0xdc, 0x94, 0xec, 0x0c, 0x62, 0x37, 0xe1, 0x1c, 0x77, 0x46, 0xe6, 0x37,
	0xee, 0xa4, 0x15, 0xca, 0xb7, 0xaa, 0x44, 0xba, 0x80, 0xb1, 0xd5, 0xf4, 0x68, 0xb8, 0xc4, 0x3f,
	0xf9, 0xe4, 0xf8, 0xff, 0x73, 0x47, 0x66, 0x41, 0x41, 0x2e, 0x89, 0xa1, 0x59, 0x43, 0x8f, 0x66,
	0xb5, 0x47, 0xf3, 0x23, 0x7b, 0xb4, 0x3f, 0x39, 0x3b, 0xda, 0x32, 0xf9, 0x6c, 0x33, 0xf2, 0xd5,
	0x66, 0xe4, 0xbb, 0xcd, 0xc8, 0xc7, 0x4f, 0x16, 0x3c, 0x4c, 0xec, 0x0f, 0x5e, 0xfd, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xb8, 0x45, 0x64, 0xcf, 0xd2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WriteServiceClient is the client API for WriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WriteServiceClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (WriteService_WriteClient, error)
}

type writeServiceClient struct {
	cc *grpc.ClientConn
}

func NewWriteServiceClient(cc *grpc.ClientConn) WriteServiceClient {
	return &writeServiceClient{cc}
}

func (c *writeServiceClient) Write(ctx context.Context, opts ...grpc.CallOption) (WriteService_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WriteService_serviceDesc.Streams[0], "/storage.WriteService/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &writeServiceWriteClient{stream}
	return x, nil
}

type WriteService_WriteClient interface {
	Send(*WriteRequest) error
	Recv() (*WriteResponse, error)
	grpc.ClientStream
}

type writeServiceWriteClient struct {
	grpc.ClientStream
}

func (x *writeServiceWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *writeServiceWriteClient) Recv() (*WriteResponse, error) {
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WriteServiceServer is the server API for WriteService service.
type WriteServiceServer interface {
	Write(WriteService_WriteServer) error
}

func RegisterWriteServiceServer(s *grpc.Server, srv WriteServiceServer) {
	s.RegisterService(&_WriteService_serviceDesc, srv)
}

func _WriteService_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WriteServiceServer).Write(&writeServiceWriteServer{stream})
}

type WriteService_WriteServer interface {
	Send(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type writeServiceWriteServer struct {
	grpc.ServerStream
}

func (x *writeServiceWriteServer) Send(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *writeServiceWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WriteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.WriteService",
	HandlerType: (*WriteServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _WriteService_Write_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	Query(ctx context.Context, opts ...grpc.CallOption) (QueryService_QueryClient, error)
}

type queryServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueryServiceClient(cc *grpc.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Query(ctx context.Context, opts ...grpc.CallOption) (QueryService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/storage.QueryService/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryClient{stream}
	return x, nil
}

type QueryService_QueryClient interface {
	Send(*QueryRequest) error
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type queryServiceQueryClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryClient) Send(m *QueryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *queryServiceQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	Query(QueryService_QueryServer) error
}

func RegisterQueryServiceServer(s *grpc.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QueryServiceServer).Query(&queryServiceQueryServer{stream})
}

type QueryService_QueryServer interface {
	Send(*QueryResponse) error
	Recv() (*QueryRequest, error)
	grpc.ServerStream
}

type queryServiceQueryServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *queryServiceQueryServer) Recv() (*QueryRequest, error) {
	m := new(QueryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _QueryService_Query_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}

func (m *Replica) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Replica) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Seq != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Seq))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WriteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Replicas) > 0 {
		for _, msg := range m.Replicas {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintStorage(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WriteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Seq != nil {
		nn1, err := m.Seq.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WriteResponse_ResetSeq) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x8
	i++
	i = encodeVarintStorage(dAtA, i, uint64(m.ResetSeq))
	return i, nil
}
func (m *WriteResponse_AckSeq) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x10
	i++
	i = encodeVarintStorage(dAtA, i, uint64(m.AckSeq))
	return i, nil
}
func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Replica) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seq != 0 {
		n += 1 + sovStorage(uint64(m.Seq))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WriteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Replicas) > 0 {
		for _, e := range m.Replicas {
			l = e.Size()
			n += 1 + l + sovStorage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WriteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seq != nil {
		n += m.Seq.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WriteResponse_ResetSeq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStorage(uint64(m.ResetSeq))
	return n
}
func (m *WriteResponse_AckSeq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStorage(uint64(m.AckSeq))
	return n
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStorage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Replica) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Replica: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Replica: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WriteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: WriteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Replicas = append(m.Replicas, &Replica{})
			if err := m.Replicas[len(m.Replicas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WriteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: WriteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResetSeq", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Seq = &WriteResponse_ResetSeq{v}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckSeq", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Seq = &WriteResponse_AckSeq{v}
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
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
				return 0, ErrInvalidLengthStorage
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthStorage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStorage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthStorage
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStorage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage   = fmt.Errorf("proto: integer overflow")
)