// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/cluster.proto

package protos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// RPCMsg message to be sent using rpc
type RPCMsg struct {
	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (m *RPCMsg) Reset()         { *m = RPCMsg{} }
func (m *RPCMsg) String() string { return proto.CompactTextString(m) }
func (*RPCMsg) ProtoMessage()    {}
func (*RPCMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_d7c099fddc97e85a, []int{0}
}
func (m *RPCMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RPCMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RPCMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RPCMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCMsg.Merge(dst, src)
}
func (m *RPCMsg) XXX_Size() int {
	return m.Size()
}
func (m *RPCMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RPCMsg proto.InternalMessageInfo

func (m *RPCMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// RPCRes is the rpc response
type RPCRes struct {
	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (m *RPCRes) Reset()         { *m = RPCRes{} }
func (m *RPCRes) String() string { return proto.CompactTextString(m) }
func (*RPCRes) ProtoMessage()    {}
func (*RPCRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_d7c099fddc97e85a, []int{1}
}
func (m *RPCRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RPCRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RPCRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RPCRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCRes.Merge(dst, src)
}
func (m *RPCRes) XXX_Size() int {
	return m.Size()
}
func (m *RPCRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCRes.DiscardUnknown(m)
}

var xxx_messageInfo_RPCRes proto.InternalMessageInfo

func (m *RPCRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RPCMsg)(nil), "protos.RPCMsg")
	proto.RegisterType((*RPCRes)(nil), "protos.RPCRes")
}
func (m *RPCMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RPCMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *RPCRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RPCRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RPCMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	return n
}

func (m *RPCRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	return n
}

func sovCluster(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCluster(x uint64) (n int) {
	return sovCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RPCMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: RPCMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RPCMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
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
func (m *RPCRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: RPCRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RPCRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
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
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
				return 0, ErrInvalidLengthCluster
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCluster
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
				next, err := skipCluster(dAtA[start:])
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
	ErrInvalidLengthCluster = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos/cluster.proto", fileDescriptor_cluster_d7c099fddc97e85a) }

var fileDescriptor_cluster_d7c099fddc97e85a = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x03, 0x73, 0x85, 0xd8, 0x20,
	0xa2, 0x4a, 0x52, 0x5c, 0x6c, 0x41, 0x01, 0xce, 0xbe, 0xc5, 0xe9, 0x42, 0x02, 0x5c, 0xcc, 0xbe,
	0xc5, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0x26, 0x54, 0x2e, 0x28, 0xb5, 0x18,
	0x53, 0xce, 0x49, 0xe2, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0x20, 0x26,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x07, 0xed, 0xba, 0x27, 0x78, 0x00, 0x00, 0x00,
}
