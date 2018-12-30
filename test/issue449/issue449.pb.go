// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue449.proto

package issue449

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

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

type CodeGenMsg struct {
	Int64ReqPtr          *int64   `protobuf:"varint,1,req,name=Int64ReqPtr" json:"Int64ReqPtr,omitempty"`
	Int32OptPtr          *int32   `protobuf:"varint,2,opt,name=Int32OptPtr" json:"Int32OptPtr,omitempty"`
	Int64Req             int64    `protobuf:"varint,3,req,name=Int64Req" json:"Int64Req"`
	Int32Opt             int32    `protobuf:"varint,4,opt,name=Int32Opt" json:"Int32Opt"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeGenMsg) Reset()         { *m = CodeGenMsg{} }
func (m *CodeGenMsg) String() string { return proto.CompactTextString(m) }
func (*CodeGenMsg) ProtoMessage()    {}
func (*CodeGenMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_dece653619274e1d, []int{0}
}
func (m *CodeGenMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeGenMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeGenMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeGenMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGenMsg.Merge(m, src)
}
func (m *CodeGenMsg) XXX_Size() int {
	return m.Size()
}
func (m *CodeGenMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGenMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGenMsg proto.InternalMessageInfo

func (m *CodeGenMsg) GetInt64ReqPtr() int64 {
	if m != nil && m.Int64ReqPtr != nil {
		return *m.Int64ReqPtr
	}
	return 0
}

func (m *CodeGenMsg) GetInt32OptPtr() int32 {
	if m != nil && m.Int32OptPtr != nil {
		return *m.Int32OptPtr
	}
	return 0
}

func (m *CodeGenMsg) GetInt64Req() int64 {
	if m != nil {
		return m.Int64Req
	}
	return 0
}

func (m *CodeGenMsg) GetInt32Opt() int32 {
	if m != nil {
		return m.Int32Opt
	}
	return 0
}

type NonCodeGenMsg struct {
	Int64ReqPtr          *int64   `protobuf:"varint,1,req,name=Int64ReqPtr" json:"Int64ReqPtr,omitempty"`
	Int32OptPtr          *int32   `protobuf:"varint,2,opt,name=Int32OptPtr" json:"Int32OptPtr,omitempty"`
	Int64Req             int64    `protobuf:"varint,3,req,name=Int64Req" json:"Int64Req"`
	Int32Opt             int32    `protobuf:"varint,4,opt,name=Int32Opt" json:"Int32Opt"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonCodeGenMsg) Reset()         { *m = NonCodeGenMsg{} }
func (m *NonCodeGenMsg) String() string { return proto.CompactTextString(m) }
func (*NonCodeGenMsg) ProtoMessage()    {}
func (*NonCodeGenMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_dece653619274e1d, []int{1}
}
func (m *NonCodeGenMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonCodeGenMsg.Unmarshal(m, b)
}
func (m *NonCodeGenMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonCodeGenMsg.Marshal(b, m, deterministic)
}
func (m *NonCodeGenMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonCodeGenMsg.Merge(m, src)
}
func (m *NonCodeGenMsg) XXX_Size() int {
	return xxx_messageInfo_NonCodeGenMsg.Size(m)
}
func (m *NonCodeGenMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_NonCodeGenMsg.DiscardUnknown(m)
}

var xxx_messageInfo_NonCodeGenMsg proto.InternalMessageInfo

func (m *NonCodeGenMsg) GetInt64ReqPtr() int64 {
	if m != nil && m.Int64ReqPtr != nil {
		return *m.Int64ReqPtr
	}
	return 0
}

func (m *NonCodeGenMsg) GetInt32OptPtr() int32 {
	if m != nil && m.Int32OptPtr != nil {
		return *m.Int32OptPtr
	}
	return 0
}

func (m *NonCodeGenMsg) GetInt64Req() int64 {
	if m != nil {
		return m.Int64Req
	}
	return 0
}

func (m *NonCodeGenMsg) GetInt32Opt() int32 {
	if m != nil {
		return m.Int32Opt
	}
	return 0
}

func init() {
	proto.RegisterType((*CodeGenMsg)(nil), "issue449.CodeGenMsg")
	proto.RegisterType((*NonCodeGenMsg)(nil), "issue449.NonCodeGenMsg")
}
func (this *CodeGenMsg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodeGenMsg)
	if !ok {
		that2, ok := that.(CodeGenMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Int64ReqPtr != nil && that1.Int64ReqPtr != nil {
		if *this.Int64ReqPtr != *that1.Int64ReqPtr {
			return false
		}
	} else if this.Int64ReqPtr != nil {
		return false
	} else if that1.Int64ReqPtr != nil {
		return false
	}
	if this.Int32OptPtr != nil && that1.Int32OptPtr != nil {
		if *this.Int32OptPtr != *that1.Int32OptPtr {
			return false
		}
	} else if this.Int32OptPtr != nil {
		return false
	} else if that1.Int32OptPtr != nil {
		return false
	}
	if this.Int64Req != that1.Int64Req {
		return false
	}
	if this.Int32Opt != that1.Int32Opt {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NonCodeGenMsg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NonCodeGenMsg)
	if !ok {
		that2, ok := that.(NonCodeGenMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Int64ReqPtr != nil && that1.Int64ReqPtr != nil {
		if *this.Int64ReqPtr != *that1.Int64ReqPtr {
			return false
		}
	} else if this.Int64ReqPtr != nil {
		return false
	} else if that1.Int64ReqPtr != nil {
		return false
	}
	if this.Int32OptPtr != nil && that1.Int32OptPtr != nil {
		if *this.Int32OptPtr != *that1.Int32OptPtr {
			return false
		}
	} else if this.Int32OptPtr != nil {
		return false
	} else if that1.Int32OptPtr != nil {
		return false
	}
	if this.Int64Req != that1.Int64Req {
		return false
	}
	if this.Int32Opt != that1.Int32Opt {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *CodeGenMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeGenMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Int64ReqPtr == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("Int64ReqPtr")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintIssue449(dAtA, i, uint64(*m.Int64ReqPtr))
	}
	if m.Int32OptPtr != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIssue449(dAtA, i, uint64(*m.Int32OptPtr))
	}
	dAtA[i] = 0x18
	i++
	i = encodeVarintIssue449(dAtA, i, uint64(m.Int64Req))
	dAtA[i] = 0x20
	i++
	i = encodeVarintIssue449(dAtA, i, uint64(m.Int32Opt))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintIssue449(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CodeGenMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Int64ReqPtr != nil {
		n += 1 + sovIssue449(uint64(*m.Int64ReqPtr))
	}
	if m.Int32OptPtr != nil {
		n += 1 + sovIssue449(uint64(*m.Int32OptPtr))
	}
	n += 1 + sovIssue449(uint64(m.Int64Req))
	n += 1 + sovIssue449(uint64(m.Int32Opt))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue449(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue449(x uint64) (n int) {
	return sovIssue449(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CodeGenMsg) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue449
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
			return fmt.Errorf("proto: CodeGenMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeGenMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64ReqPtr", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue449
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int64ReqPtr = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32OptPtr", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue449
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int32OptPtr = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64Req", wireType)
			}
			m.Int64Req = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue449
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int64Req |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32Opt", wireType)
			}
			m.Int32Opt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue449
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int32Opt |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIssue449(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue449
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Int64ReqPtr")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Int64Req")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIssue449(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue449
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
					return 0, ErrIntOverflowIssue449
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
					return 0, ErrIntOverflowIssue449
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
				return 0, ErrInvalidLengthIssue449
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue449
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
				next, err := skipIssue449(dAtA[start:])
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
	ErrInvalidLengthIssue449 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue449   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("issue449.proto", fileDescriptor_dece653619274e1d) }

var fileDescriptor_dece653619274e1d = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0xb1, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0x5a, 0xc4, 0xc8,
	0xc5, 0xe5, 0x9c, 0x9f, 0x92, 0xea, 0x9e, 0x9a, 0xe7, 0x5b, 0x9c, 0x2e, 0xa4, 0xc0, 0xc5, 0xed,
	0x99, 0x57, 0x62, 0x66, 0x12, 0x94, 0x5a, 0x18, 0x50, 0x52, 0x24, 0xc1, 0xa8, 0xc0, 0xa4, 0xc1,
	0x1c, 0x84, 0x2c, 0x04, 0x55, 0x61, 0x6c, 0xe4, 0x5f, 0x50, 0x02, 0x52, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x1a, 0x84, 0x2c, 0x24, 0xa4, 0xc0, 0xc5, 0x01, 0xd3, 0x20, 0xc1, 0x0c, 0x32, 0xc0, 0x89,
	0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0xb8, 0x28, 0x54, 0x05, 0x58, 0x83, 0x04, 0x0b, 0xc8, 0x00,
	0x24, 0x15, 0x60, 0x51, 0x2b, 0x9e, 0x8e, 0x85, 0xf2, 0x8c, 0x13, 0x16, 0xca, 0x33, 0x2e, 0x58,
	0x28, 0xcf, 0xa8, 0xb4, 0x94, 0x91, 0x8b, 0xd7, 0x2f, 0x3f, 0x6f, 0x90, 0xba, 0x93, 0x61, 0xc2,
	0x42, 0x79, 0x86, 0x05, 0x0b, 0xe5, 0x19, 0x9c, 0x58, 0x56, 0x3c, 0x92, 0x63, 0x04, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0x5a, 0xe2, 0x72, 0x9c, 0x01, 0x00, 0x00,
}
