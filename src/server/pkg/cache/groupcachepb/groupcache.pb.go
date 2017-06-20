// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/pkg/cache/groupcachepb/groupcache.proto

/*
	Package groupcachepb is a generated protocol buffer package.

	It is generated from these files:
		server/pkg/cache/groupcachepb/groupcache.proto

	It has these top-level messages:
		GetRequest
		GetResponse
*/
package groupcachepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Group            *string `protobuf:"bytes,1,req,name=group" json:"group,omitempty"`
	Key              *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptorGroupcache, []int{0} }

func (m *GetRequest) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *GetRequest) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type GetResponse struct {
	Value            []byte   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	MinuteQps        *float64 `protobuf:"fixed64,2,opt,name=minute_qps,json=minuteQps" json:"minute_qps,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptorGroupcache, []int{1} }

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GetResponse) GetMinuteQps() float64 {
	if m != nil && m.MinuteQps != nil {
		return *m.MinuteQps
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "groupcachepb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "groupcachepb.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GroupCache service

type GroupCacheClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type groupCacheClient struct {
	cc *grpc.ClientConn
}

func NewGroupCacheClient(cc *grpc.ClientConn) GroupCacheClient {
	return &groupCacheClient{cc}
}

func (c *groupCacheClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/groupcachepb.GroupCache/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupCache service

type GroupCacheServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterGroupCacheServer(s *grpc.Server, srv GroupCacheServer) {
	s.RegisterService(&_GroupCache_serviceDesc, srv)
}

func _GroupCache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupCacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/groupcachepb.GroupCache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupCacheServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "groupcachepb.GroupCache",
	HandlerType: (*GroupCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GroupCache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pkg/cache/groupcachepb/groupcache.proto",
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Group == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGroupcache(dAtA, i, uint64(len(*m.Group)))
		i += copy(dAtA[i:], *m.Group)
	}
	if m.Key == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGroupcache(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGroupcache(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.MinuteQps != nil {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64Groupcache(dAtA, i, uint64(math.Float64bits(float64(*m.MinuteQps))))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Groupcache(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Groupcache(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGroupcache(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetRequest) Size() (n int) {
	var l int
	_ = l
	if m.Group != nil {
		l = len(*m.Group)
		n += 1 + l + sovGroupcache(uint64(l))
	}
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovGroupcache(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovGroupcache(uint64(l))
	}
	if m.MinuteQps != nil {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGroupcache(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGroupcache(x uint64) (n int) {
	return sovGroupcache(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupcache
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupcache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroupcache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Group = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupcache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroupcache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipGroupcache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroupcache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupcache
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupcache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroupcache
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinuteQps", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.MinuteQps = &v2
		default:
			iNdEx = preIndex
			skippy, err := skipGroupcache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroupcache
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
func skipGroupcache(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroupcache
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
					return 0, ErrIntOverflowGroupcache
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
					return 0, ErrIntOverflowGroupcache
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGroupcache
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGroupcache
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
				next, err := skipGroupcache(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthGroupcache = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroupcache   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("server/pkg/cache/groupcachepb/groupcache.proto", fileDescriptorGroupcache)
}

var fileDescriptorGroupcache = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0x4e, 0x4c, 0xce, 0x48, 0xd5, 0x4f, 0x2f, 0xca,
	0x2f, 0x2d, 0x00, 0x33, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x3c,
	0xc8, 0xd2, 0x4a, 0x26, 0x5c, 0x5c, 0xee, 0xa9, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x22, 0x5c, 0xac, 0x60, 0x59, 0x09, 0x46, 0x05, 0x26, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48,
	0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x09, 0x2c, 0x06, 0x62, 0x2a, 0x39, 0x71, 0x71, 0x83,
	0x75, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82, 0xb4, 0x95, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38, 0x42, 0xb2, 0x5c, 0x5c, 0xb9, 0x99, 0x79, 0xa5, 0x25,
	0xa9, 0xf1, 0x85, 0x05, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x9c, 0x10, 0x91, 0xc0,
	0x82, 0x62, 0x23, 0x2f, 0x2e, 0x2e, 0x77, 0x90, 0xf1, 0xce, 0x20, 0x97, 0x08, 0xd9, 0x70, 0x31,
	0xbb, 0xa7, 0x96, 0x08, 0x49, 0xe8, 0x21, 0xbb, 0x4e, 0x0f, 0xe1, 0x34, 0x29, 0x49, 0x2c, 0x32,
	0x10, 0xeb, 0x95, 0x18, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x71, 0x24,
	0x37, 0x16, 0x01, 0x00, 0x00,
}
