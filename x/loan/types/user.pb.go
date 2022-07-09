// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loan/user.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type User struct {
	Creator        string           `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id             uint64           `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Collateral     []bool           `protobuf:"varint,3,rep,packed,name=collateral,proto3" json:"collateral,omitempty"`
	Deposit        []*Deposit       `protobuf:"bytes,4,rep,name=deposit,proto3" json:"deposit,omitempty"`
	Borrow         []*Borrow        `protobuf:"bytes,5,rep,name=borrow,proto3" json:"borrow,omitempty"`
	AssetBalances  []int32          `protobuf:"varint,6,rep,packed,name=assetBalances,proto3" json:"assetBalances,omitempty"`
	DepositEarneds []*DepositEarned `protobuf:"bytes,7,rep,name=depositEarneds,proto3" json:"depositEarneds,omitempty"`
	BorrowAccured  []*BorrowAccured `protobuf:"bytes,8,rep,name=borrowAccured,proto3" json:"borrowAccured,omitempty"`
	TxHistories    []*TxHistory     `protobuf:"bytes,9,rep,name=txHistories,proto3" json:"txHistories,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_90629940cb7ec090, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetCollateral() []bool {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *User) GetDeposit() []*Deposit {
	if m != nil {
		return m.Deposit
	}
	return nil
}

func (m *User) GetBorrow() []*Borrow {
	if m != nil {
		return m.Borrow
	}
	return nil
}

func (m *User) GetAssetBalances() []int32 {
	if m != nil {
		return m.AssetBalances
	}
	return nil
}

func (m *User) GetDepositEarneds() []*DepositEarned {
	if m != nil {
		return m.DepositEarneds
	}
	return nil
}

func (m *User) GetBorrowAccured() []*BorrowAccured {
	if m != nil {
		return m.BorrowAccured
	}
	return nil
}

func (m *User) GetTxHistories() []*TxHistory {
	if m != nil {
		return m.TxHistories
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "loan.loan.User")
}

func init() { proto.RegisterFile("loan/user.proto", fileDescriptor_90629940cb7ec090) }

var fileDescriptor_90629940cb7ec090 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xd7, 0x75, 0x7f, 0x33, 0xb6, 0x1f, 0x7b, 0xd8, 0x0f, 0xe2, 0x0e, 0xa1, 0x88, 0x87,
	0x8a, 0xd2, 0x81, 0x82, 0x47, 0xd1, 0xa2, 0xe0, 0x39, 0xe8, 0xc5, 0xcb, 0xc8, 0xda, 0x30, 0x0b,
	0x65, 0x19, 0x49, 0x86, 0xdb, 0x4b, 0xf0, 0xe6, 0xcb, 0xf2, 0xb8, 0xa3, 0x47, 0xd9, 0xde, 0x88,
	0x2c, 0xe9, 0x46, 0xea, 0x2e, 0xa5, 0xfd, 0x7c, 0x3f, 0xcf, 0xd3, 0xa7, 0x4f, 0x8a, 0xfe, 0xe5,
	0x82, 0xcd, 0x46, 0x0b, 0xc5, 0x65, 0x34, 0x97, 0x42, 0x0b, 0x68, 0xef, 0x40, 0xb4, 0xbb, 0x0c,
	0x07, 0x53, 0x31, 0x15, 0x86, 0x8e, 0x76, 0x77, 0x56, 0x18, 0x82, 0xa9, 0x48, 0xf9, 0x5c, 0xa8,
	0x4c, 0x17, 0xac, 0x6f, 0xd8, 0x44, 0x48, 0x29, 0xde, 0x0b, 0x74, 0xe2, 0x6a, 0x63, 0xce, 0xe4,
	0x8c, 0xa7, 0xa5, 0xc8, 0xda, 0x63, 0x96, 0x24, 0x0b, 0x79, 0x88, 0xfe, 0x9b, 0x48, 0x2f, 0xc7,
	0x6f, 0x99, 0xd2, 0x42, 0xae, 0x2c, 0x3e, 0xfd, 0xf0, 0x51, 0xed, 0x45, 0x71, 0x09, 0x18, 0x35,
	0x13, 0xc9, 0x99, 0x16, 0x12, 0x7b, 0x81, 0x17, 0xb6, 0xe9, 0xfe, 0x11, 0x7a, 0xa8, 0x9a, 0xa5,
	0xb8, 0x1a, 0x78, 0x61, 0x8d, 0x56, 0xb3, 0x14, 0x08, 0x42, 0x89, 0xc8, 0x73, 0xa6, 0xb9, 0x64,
	0x39, 0xf6, 0x03, 0x3f, 0x6c, 0x51, 0x87, 0xc0, 0x25, 0x6a, 0x16, 0xc3, 0xe1, 0x5a, 0xe0, 0x87,
	0x9d, 0x2b, 0x88, 0x0e, 0x5f, 0x1e, 0x3d, 0xd8, 0x84, 0xee, 0x15, 0x38, 0x47, 0x0d, 0x3b, 0x2f,
	0xae, 0x1b, 0xb9, 0xef, 0xc8, 0xb1, 0x09, 0x68, 0x21, 0xc0, 0x19, 0xea, 0x32, 0xa5, 0xb8, 0x8e,
	0x59, 0xce, 0x66, 0x09, 0x57, 0xb8, 0x11, 0xf8, 0x61, 0x9d, 0x96, 0x21, 0xdc, 0xa1, 0x5e, 0xd1,
	0xfb, 0xd1, 0xac, 0x46, 0xe1, 0xa6, 0x69, 0x8c, 0x8f, 0xa7, 0xb0, 0x02, 0xfd, 0xe3, 0xc3, 0x2d,
	0xea, 0xda, 0x37, 0xde, 0xdb, 0x0d, 0xe2, 0xd6, 0x51, 0x83, 0xd8, 0xcd, 0x69, 0x59, 0x87, 0x1b,
	0xd4, 0xd1, 0xcb, 0x27, 0xb3, 0xe6, 0x8c, 0x2b, 0xdc, 0x36, 0xd5, 0x03, 0xa7, 0xfa, 0xb9, 0x48,
	0x57, 0xd4, 0x15, 0xe3, 0x8b, 0xaf, 0x0d, 0xf1, 0xd6, 0x1b, 0xe2, 0xfd, 0x6c, 0x88, 0xf7, 0xb9,
	0x25, 0x95, 0xf5, 0x96, 0x54, 0xbe, 0xb7, 0xa4, 0xf2, 0x6a, 0xff, 0x82, 0xe5, 0xc8, 0x9e, 0xe1,
	0x6a, 0xce, 0xd5, 0xa4, 0x61, 0xce, 0xef, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x41, 0xb2, 0x87,
	0x0a, 0x67, 0x02, 0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHistories) > 0 {
		for iNdEx := len(m.TxHistories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxHistories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.BorrowAccured) > 0 {
		for iNdEx := len(m.BorrowAccured) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BorrowAccured[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.DepositEarneds) > 0 {
		for iNdEx := len(m.DepositEarneds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositEarneds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AssetBalances) > 0 {
		dAtA2 := make([]byte, len(m.AssetBalances)*10)
		var j1 int
		for _, num1 := range m.AssetBalances {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintUser(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Borrow) > 0 {
		for iNdEx := len(m.Borrow) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Borrow[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Deposit) > 0 {
		for iNdEx := len(m.Deposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Collateral) > 0 {
		for iNdEx := len(m.Collateral) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.Collateral[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintUser(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovUser(uint64(m.Id))
	}
	if len(m.Collateral) > 0 {
		n += 1 + sovUser(uint64(len(m.Collateral))) + len(m.Collateral)*1
	}
	if len(m.Deposit) > 0 {
		for _, e := range m.Deposit {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.Borrow) > 0 {
		for _, e := range m.Borrow {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.AssetBalances) > 0 {
		l = 0
		for _, e := range m.AssetBalances {
			l += sovUser(uint64(e))
		}
		n += 1 + sovUser(uint64(l)) + l
	}
	if len(m.DepositEarneds) > 0 {
		for _, e := range m.DepositEarneds {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.BorrowAccured) > 0 {
		for _, e := range m.BorrowAccured {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.TxHistories) > 0 {
		for _, e := range m.TxHistories {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
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
					return ErrIntOverflowUser
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
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Collateral = append(m.Collateral, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthUser
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthUser
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.Collateral) == 0 {
					m.Collateral = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUser
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Collateral = append(m.Collateral, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = append(m.Deposit, &Deposit{})
			if err := m.Deposit[len(m.Deposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrow = append(m.Borrow, &Borrow{})
			if err := m.Borrow[len(m.Borrow)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AssetBalances = append(m.AssetBalances, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthUser
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthUser
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.AssetBalances) == 0 {
					m.AssetBalances = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUser
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AssetBalances = append(m.AssetBalances, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetBalances", wireType)
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositEarneds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositEarneds = append(m.DepositEarneds, &DepositEarned{})
			if err := m.DepositEarneds[len(m.DepositEarneds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowAccured", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorrowAccured = append(m.BorrowAccured, &BorrowAccured{})
			if err := m.BorrowAccured[len(m.BorrowAccured)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHistories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHistories = append(m.TxHistories, &TxHistory{})
			if err := m.TxHistories[len(m.TxHistories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
