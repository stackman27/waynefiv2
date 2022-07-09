// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loan/deposit_earned.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type DepositEarned struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlockHeight int32  `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Asset       string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount      int32  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,5,opt,name=denom,proto3" json:"denom,omitempty"`
	Creator     string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *DepositEarned) Reset()         { *m = DepositEarned{} }
func (m *DepositEarned) String() string { return proto.CompactTextString(m) }
func (*DepositEarned) ProtoMessage()    {}
func (*DepositEarned) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f5c77dbfd220db, []int{0}
}
func (m *DepositEarned) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositEarned) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositEarned.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositEarned) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositEarned.Merge(m, src)
}
func (m *DepositEarned) XXX_Size() int {
	return m.Size()
}
func (m *DepositEarned) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositEarned.DiscardUnknown(m)
}

var xxx_messageInfo_DepositEarned proto.InternalMessageInfo

func (m *DepositEarned) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DepositEarned) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *DepositEarned) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *DepositEarned) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DepositEarned) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DepositEarned) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*DepositEarned)(nil), "loan.loan.DepositEarned")
}

func init() { proto.RegisterFile("loan/deposit_earned.proto", fileDescriptor_c2f5c77dbfd220db) }

var fileDescriptor_c2f5c77dbfd220db = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xc9, 0x4f, 0xcc,
	0xd3, 0x4f, 0x49, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0x89, 0x4f, 0x4d, 0x2c, 0xca, 0x4b, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x49, 0xe9, 0x81, 0x08, 0xa5, 0xf9, 0x8c, 0x5c,
	0xbc, 0x2e, 0x10, 0x35, 0xae, 0x60, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x0a, 0x5c, 0xdc, 0x49, 0x39, 0xf9, 0xc9, 0xd9,
	0x1e, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0xc8, 0x42, 0x42,
	0x22, 0x5c, 0xac, 0x89, 0xc5, 0xc5, 0xa9, 0x25, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10,
	0x8e, 0x90, 0x18, 0x17, 0x5b, 0x62, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x0b, 0x58, 0x0b, 0x94,
	0x07, 0x52, 0x9d, 0x92, 0x9a, 0x97, 0x9f, 0x2b, 0xc1, 0x0a, 0x51, 0x0d, 0xe6, 0x08, 0x49, 0x70,
	0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0xb0, 0x81, 0xc5, 0x61, 0x5c, 0x27, 0xed,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39,
	0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x04, 0xfb, 0xb0, 0x42, 0x1f,
	0x4c, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x3d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x93, 0x73, 0xee, 0xf8, 0xfd, 0x00, 0x00, 0x00,
}

func (m *DepositEarned) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositEarned) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositEarned) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDepositEarned(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDepositEarned(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Amount != 0 {
		i = encodeVarintDepositEarned(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintDepositEarned(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintDepositEarned(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintDepositEarned(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepositEarned(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepositEarned(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositEarned) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDepositEarned(uint64(m.Id))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovDepositEarned(uint64(m.BlockHeight))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovDepositEarned(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovDepositEarned(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDepositEarned(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDepositEarned(uint64(l))
	}
	return n
}

func sovDepositEarned(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepositEarned(x uint64) (n int) {
	return sovDepositEarned(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositEarned) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepositEarned
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
			return fmt.Errorf("proto: DepositEarned: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositEarned: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositEarned
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositEarned
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositEarned
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
				return ErrInvalidLengthDepositEarned
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositEarned
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositEarned
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositEarned
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
				return ErrInvalidLengthDepositEarned
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositEarned
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositEarned
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
				return ErrInvalidLengthDepositEarned
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositEarned
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDepositEarned(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDepositEarned
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
func skipDepositEarned(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepositEarned
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
					return 0, ErrIntOverflowDepositEarned
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
					return 0, ErrIntOverflowDepositEarned
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
				return 0, ErrInvalidLengthDepositEarned
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepositEarned
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepositEarned
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepositEarned        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepositEarned          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepositEarned = fmt.Errorf("proto: unexpected end of group")
)