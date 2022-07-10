// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loan/pool.proto

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

type Pool struct {
	Creator          string  `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id               uint64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Asset            string  `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	Denom            string  `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty"`
	CollatoralFactor int32   `protobuf:"varint,5,opt,name=collatoralFactor,proto3" json:"collatoralFactor,omitempty"`
	DepositBalance   int32   `protobuf:"varint,6,opt,name=depositBalance,proto3" json:"depositBalance,omitempty"`
	BorrowBalance    int32   `protobuf:"varint,7,opt,name=borrowBalance,proto3" json:"borrowBalance,omitempty"`
	Users            []*User `protobuf:"bytes,9,rep,name=users,proto3" json:"users,omitempty"`
	Aprs             []*Apr  `protobuf:"bytes,10,rep,name=aprs,proto3" json:"aprs,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c20d1b0042bdddb, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Pool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Pool) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Pool) GetCollatoralFactor() int32 {
	if m != nil {
		return m.CollatoralFactor
	}
	return 0
}

func (m *Pool) GetDepositBalance() int32 {
	if m != nil {
		return m.DepositBalance
	}
	return 0
}

func (m *Pool) GetBorrowBalance() int32 {
	if m != nil {
		return m.BorrowBalance
	}
	return 0
}

func (m *Pool) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Pool) GetAprs() []*Apr {
	if m != nil {
		return m.Aprs
	}
	return nil
}

func init() {
	proto.RegisterType((*Pool)(nil), "loan.loan.Pool")
}

func init() { proto.RegisterFile("loan/pool.proto", fileDescriptor_4c20d1b0042bdddb) }

var fileDescriptor_4c20d1b0042bdddb = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x9b, 0x4e, 0x3b, 0x43, 0xaf, 0xd8, 0xd1, 0xe0, 0x22, 0xcc, 0x22, 0x94, 0x41, 0xa5,
	0x28, 0x74, 0x40, 0x9f, 0xc0, 0x59, 0xb8, 0x96, 0x82, 0x1b, 0x77, 0x99, 0x36, 0x8b, 0x42, 0xec,
	0x0d, 0x49, 0x45, 0x7d, 0x0b, 0x5f, 0xc1, 0xb7, 0x71, 0x39, 0x4b, 0x97, 0xd2, 0xbe, 0x88, 0x24,
	0x75, 0xc4, 0x9f, 0x4d, 0xcb, 0xf9, 0xce, 0xc7, 0x85, 0x1c, 0x98, 0x2b, 0x14, 0xed, 0x4a, 0x23,
	0xaa, 0x42, 0x1b, 0xec, 0x90, 0x26, 0x0e, 0x14, 0xee, 0xb3, 0x18, 0xbb, 0x07, 0x2b, 0xcd, 0xd8,
	0x2d, 0x52, 0x0f, 0x84, 0xfe, 0xca, 0xcb, 0xd7, 0x10, 0xa2, 0x1b, 0x44, 0x45, 0x19, 0xcc, 0x2a,
	0x23, 0x45, 0x87, 0x86, 0x91, 0x8c, 0xe4, 0x49, 0xb9, 0x8b, 0x34, 0x85, 0xb0, 0xa9, 0x59, 0x98,
	0x91, 0x3c, 0x2a, 0xc3, 0xa6, 0xa6, 0x47, 0x10, 0x0b, 0x6b, 0x65, 0xc7, 0x26, 0xde, 0x1b, 0x83,
	0xa3, 0xb5, 0x6c, 0xf1, 0x9e, 0x45, 0x23, 0xf5, 0x81, 0x9e, 0xc1, 0x41, 0x85, 0x4a, 0xb9, 0x3b,
	0x42, 0x5d, 0x8b, 0xca, 0x9d, 0x8f, 0x33, 0x92, 0xc7, 0xe5, 0x3f, 0x4e, 0x4f, 0x21, 0xad, 0xa5,
	0x46, 0xdb, 0x74, 0x6b, 0xa1, 0x44, 0x5b, 0x49, 0x36, 0xf5, 0xe6, 0x1f, 0x4a, 0x8f, 0x61, 0x7f,
	0x83, 0xc6, 0xe0, 0xe3, 0x4e, 0x9b, 0x79, 0xed, 0x37, 0xa4, 0x27, 0x10, 0xbb, 0x67, 0x5b, 0x96,
	0x64, 0x93, 0x7c, 0xef, 0x62, 0x5e, 0x7c, 0x8f, 0x52, 0xdc, 0x5a, 0x69, 0xca, 0xb1, 0xa5, 0x4b,
	0x88, 0x84, 0x36, 0x96, 0x81, 0xb7, 0xd2, 0x1f, 0xd6, 0x95, 0x36, 0xa5, 0xef, 0xd6, 0xe7, 0x6f,
	0x3d, 0x27, 0xdb, 0x9e, 0x93, 0x8f, 0x9e, 0x93, 0x97, 0x81, 0x07, 0xdb, 0x81, 0x07, 0xef, 0x03,
	0x0f, 0xee, 0x0e, 0xfd, 0x9a, 0x4f, 0x2b, 0xff, 0xeb, 0x9e, 0xb5, 0xb4, 0x9b, 0xa9, 0xdf, 0xf5,
	0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xeb, 0xa9, 0x18, 0x96, 0x01, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Aprs) > 0 {
		for iNdEx := len(m.Aprs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Aprs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Users[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.BorrowBalance != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.BorrowBalance))
		i--
		dAtA[i] = 0x38
	}
	if m.DepositBalance != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.DepositBalance))
		i--
		dAtA[i] = 0x30
	}
	if m.CollatoralFactor != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.CollatoralFactor))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.CollatoralFactor != 0 {
		n += 1 + sovPool(uint64(m.CollatoralFactor))
	}
	if m.DepositBalance != 0 {
		n += 1 + sovPool(uint64(m.DepositBalance))
	}
	if m.BorrowBalance != 0 {
		n += 1 + sovPool(uint64(m.BorrowBalance))
	}
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	if len(m.Aprs) > 0 {
		for _, e := range m.Aprs {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollatoralFactor", wireType)
			}
			m.CollatoralFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CollatoralFactor |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositBalance", wireType)
			}
			m.DepositBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositBalance |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowBalance", wireType)
			}
			m.BorrowBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BorrowBalance |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &User{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aprs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aprs = append(m.Aprs, &Apr{})
			if err := m.Aprs[len(m.Aprs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)