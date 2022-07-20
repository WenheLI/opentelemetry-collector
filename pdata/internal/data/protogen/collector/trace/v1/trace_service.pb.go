// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/trace/v1/trace_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	v1 "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
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

type ExportTraceServiceRequest struct {
	// An array of ResourceSpans.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceSpans []*v1.ResourceSpans `protobuf:"bytes,1,rep,name=resource_spans,json=resourceSpans,proto3" json:"resource_spans,omitempty"`
}

func (m *ExportTraceServiceRequest) Reset()         { *m = ExportTraceServiceRequest{} }
func (m *ExportTraceServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportTraceServiceRequest) ProtoMessage()    {}
func (*ExportTraceServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192a962890318cf4, []int{0}
}
func (m *ExportTraceServiceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportTraceServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportTraceServiceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportTraceServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTraceServiceRequest.Merge(m, src)
}
func (m *ExportTraceServiceRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExportTraceServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTraceServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTraceServiceRequest proto.InternalMessageInfo

func (m *ExportTraceServiceRequest) GetResourceSpans() []*v1.ResourceSpans {
	if m != nil {
		return m.ResourceSpans
	}
	return nil
}

type ExportTraceServiceResponse struct {
}

func (m *ExportTraceServiceResponse) Reset()         { *m = ExportTraceServiceResponse{} }
func (m *ExportTraceServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportTraceServiceResponse) ProtoMessage()    {}
func (*ExportTraceServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192a962890318cf4, []int{1}
}
func (m *ExportTraceServiceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportTraceServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportTraceServiceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportTraceServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTraceServiceResponse.Merge(m, src)
}
func (m *ExportTraceServiceResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExportTraceServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTraceServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTraceServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportTraceServiceRequest)(nil), "opentelemetry.proto.collector.trace.v1.ExportTraceServiceRequest")
	proto.RegisterType((*ExportTraceServiceResponse)(nil), "opentelemetry.proto.collector.trace.v1.ExportTraceServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/trace/v1/trace_service.proto", fileDescriptor_192a962890318cf4)
}

var fileDescriptor_192a962890318cf4 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xbf, 0x4a, 0xc4, 0x30,
	0x18, 0x6f, 0x10, 0x6e, 0x88, 0x7f, 0xc0, 0xe2, 0xa0, 0x45, 0x82, 0xdc, 0x20, 0x27, 0x42, 0xc2,
	0x9d, 0x9b, 0x9b, 0x05, 0x47, 0xe1, 0xe8, 0x39, 0xb9, 0x48, 0xad, 0x1f, 0xa5, 0x50, 0xf3, 0xc5,
	0x24, 0x57, 0xf4, 0x2d, 0xf4, 0x15, 0x5c, 0x7c, 0x15, 0xc7, 0x1b, 0x1d, 0xa5, 0x7d, 0x11, 0x69,
	0xa2, 0xd2, 0x93, 0x0a, 0x07, 0x6e, 0xdf, 0x9f, 0xdf, 0xbf, 0x84, 0x8f, 0x9e, 0xa2, 0x02, 0x69,
	0xa1, 0x84, 0x3b, 0xb0, 0xfa, 0x51, 0x28, 0x8d, 0x16, 0x45, 0x86, 0x65, 0x09, 0x99, 0x45, 0x2d,
	0xac, 0x4e, 0x33, 0x10, 0xd5, 0xd8, 0x17, 0xd7, 0x06, 0x74, 0x55, 0x64, 0xc0, 0x1d, 0x2c, 0x3c,
	0x5c, 0xe2, 0xfa, 0x21, 0xff, 0xe1, 0x72, 0x47, 0xe1, 0xd5, 0x38, 0xda, 0xc9, 0x31, 0x47, 0xaf,
	0xdc, 0x56, 0x1e, 0x18, 0x8d, 0xfa, 0x9c, 0x97, 0xfd, 0x3c, 0x72, 0x88, 0x74, 0xef, 0xfc, 0x41,
	0xa1, 0xb6, 0x97, 0xed, 0x70, 0xe6, 0x33, 0x24, 0x70, 0x3f, 0x07, 0x63, 0xc3, 0x84, 0x6e, 0x69,
	0x30, 0x38, 0xd7, 0x6d, 0x3c, 0x95, 0x4a, 0xb3, 0x4b, 0x0e, 0xd6, 0x46, 0xeb, 0x93, 0x63, 0xde,
	0x97, 0xee, 0x3b, 0x13, 0x4f, 0xbe, 0x38, 0xb3, 0x96, 0x92, 0x6c, 0xea, 0x6e, 0x3b, 0xdc, 0xa7,
	0x51, 0x9f, 0xa1, 0x51, 0x28, 0x0d, 0x4c, 0x5e, 0x08, 0xdd, 0xe8, 0x2e, 0xc2, 0x67, 0x42, 0x07,
	0x1e, 0x1f, 0x9e, 0xf1, 0xd5, 0xfe, 0x84, 0xff, 0xf9, 0xa0, 0x28, 0xfe, 0x8f, 0x84, 0x8f, 0x38,
	0x0c, 0xe2, 0x57, 0xf2, 0x56, 0x33, 0xb2, 0xa8, 0x19, 0xf9, 0xa8, 0x19, 0x79, 0x6a, 0x58, 0xb0,
	0x68, 0x58, 0xf0, 0xde, 0xb0, 0x80, 0x1e, 0x15, 0xb8, 0xa2, 0x45, 0xbc, 0xdd, 0x55, 0x9f, 0xb6,
	0xa8, 0x29, 0xb9, 0xba, 0xc8, 0x7f, 0xf3, 0x8b, 0xee, 0xc9, 0xa8, 0xdb, 0xd4, 0xa6, 0xa2, 0x90,
	0x16, 0xb4, 0x4c, 0x4b, 0xe1, 0x3a, 0x67, 0x90, 0x83, 0xec, 0xb9, 0xac, 0x9b, 0x81, 0x5b, 0x9e,
	0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x91, 0xe9, 0x89, 0x8a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportTraceServiceRequest, opts ...grpc.CallOption) (*ExportTraceServiceResponse, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) Export(ctx context.Context, in *ExportTraceServiceRequest, opts ...grpc.CallOption) (*ExportTraceServiceResponse, error) {
	out := new(ExportTraceServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.trace.v1.TraceService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportTraceServiceRequest) (*ExportTraceServiceResponse, error)
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) Export(ctx context.Context, req *ExportTraceServiceRequest) (*ExportTraceServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportTraceServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.trace.v1.TraceService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).Export(ctx, req.(*ExportTraceServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.trace.v1.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _TraceService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/trace/v1/trace_service.proto",
}

func (m *ExportTraceServiceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportTraceServiceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportTraceServiceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResourceSpans) > 0 {
		for iNdEx := len(m.ResourceSpans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceSpans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTraceService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExportTraceServiceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportTraceServiceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportTraceServiceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTraceService(dAtA []byte, offset int, v uint64) int {
	offset -= sovTraceService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExportTraceServiceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ResourceSpans) > 0 {
		for _, e := range m.ResourceSpans {
			l = e.Size()
			n += 1 + l + sovTraceService(uint64(l))
		}
	}
	return n
}

func (m *ExportTraceServiceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTraceService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTraceService(x uint64) (n int) {
	return sovTraceService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExportTraceServiceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
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
			return fmt.Errorf("proto: ExportTraceServiceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportTraceServiceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceSpans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
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
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceSpans = append(m.ResourceSpans, &v1.ResourceSpans{})
			if err := m.ResourceSpans[len(m.ResourceSpans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraceService
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
func (m *ExportTraceServiceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
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
			return fmt.Errorf("proto: ExportTraceServiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportTraceServiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraceService
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
func skipTraceService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceService
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
					return 0, ErrIntOverflowTraceService
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
					return 0, ErrIntOverflowTraceService
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
				return 0, ErrInvalidLengthTraceService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTraceService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTraceService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTraceService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTraceService = fmt.Errorf("proto: unexpected end of group")
)
