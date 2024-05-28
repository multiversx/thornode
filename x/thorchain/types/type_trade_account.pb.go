// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_trade_account.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/thorchain/thornode/common"
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

type TradeAccount struct {
	Asset              common.Asset                                  `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset"`
	Units              github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,2,opt,name=units,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"units"`
	Owner              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
	LastAddHeight      int64                                         `protobuf:"varint,4,opt,name=last_add_height,json=lastAddHeight,proto3" json:"last_add_height,omitempty"`
	LastWithdrawHeight int64                                         `protobuf:"varint,5,opt,name=last_withdraw_height,json=lastWithdrawHeight,proto3" json:"last_withdraw_height,omitempty"`
}

func (m *TradeAccount) Reset()         { *m = TradeAccount{} }
func (m *TradeAccount) String() string { return proto.CompactTextString(m) }
func (*TradeAccount) ProtoMessage()    {}
func (*TradeAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_68dec3a28ee06bd4, []int{0}
}
func (m *TradeAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradeAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradeAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradeAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeAccount.Merge(m, src)
}
func (m *TradeAccount) XXX_Size() int {
	return m.Size()
}
func (m *TradeAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeAccount.DiscardUnknown(m)
}

var xxx_messageInfo_TradeAccount proto.InternalMessageInfo

func (m *TradeAccount) GetAsset() common.Asset {
	if m != nil {
		return m.Asset
	}
	return common.Asset{}
}

func (m *TradeAccount) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *TradeAccount) GetLastAddHeight() int64 {
	if m != nil {
		return m.LastAddHeight
	}
	return 0
}

func (m *TradeAccount) GetLastWithdrawHeight() int64 {
	if m != nil {
		return m.LastWithdrawHeight
	}
	return 0
}

type TradeUnit struct {
	Asset common.Asset                            `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset"`
	Units github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=units,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"units"`
	Depth github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=depth,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"depth"`
}

func (m *TradeUnit) Reset()         { *m = TradeUnit{} }
func (m *TradeUnit) String() string { return proto.CompactTextString(m) }
func (*TradeUnit) ProtoMessage()    {}
func (*TradeUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_68dec3a28ee06bd4, []int{1}
}
func (m *TradeUnit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradeUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradeUnit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradeUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeUnit.Merge(m, src)
}
func (m *TradeUnit) XXX_Size() int {
	return m.Size()
}
func (m *TradeUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeUnit.DiscardUnknown(m)
}

var xxx_messageInfo_TradeUnit proto.InternalMessageInfo

func (m *TradeUnit) GetAsset() common.Asset {
	if m != nil {
		return m.Asset
	}
	return common.Asset{}
}

func init() {
	proto.RegisterType((*TradeAccount)(nil), "types.TradeAccount")
	proto.RegisterType((*TradeUnit)(nil), "types.TradeUnit")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_trade_account.proto", fileDescriptor_68dec3a28ee06bd4)
}

var fileDescriptor_68dec3a28ee06bd4 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x52, 0xbf, 0x6a, 0xe3, 0x30,
	0x18, 0xb7, 0x92, 0xf8, 0x20, 0xba, 0x84, 0x03, 0x93, 0xc1, 0x64, 0x70, 0x4c, 0x86, 0x3b, 0xdf,
	0x90, 0xf8, 0x72, 0x1d, 0x3a, 0xdb, 0x10, 0xda, 0xae, 0xa6, 0xa1, 0xd0, 0xc5, 0x28, 0x96, 0xb0,
	0x44, 0x13, 0x29, 0x58, 0x4a, 0xd3, 0xbe, 0x45, 0xdf, 0xa1, 0xaf, 0xd2, 0x21, 0x63, 0xc6, 0xd2,
	0x21, 0x94, 0xe4, 0x2d, 0x3a, 0x15, 0x4b, 0x2e, 0x6d, 0xe9, 0x52, 0xba, 0x74, 0xf1, 0xf7, 0xf9,
	0xf7, 0x0f, 0xf4, 0xf1, 0x83, 0x87, 0x8a, 0x8a, 0x22, 0xa3, 0x88, 0xf1, 0xf0, 0x72, 0x14, 0x5e,
	0x85, 0xaf, 0xbf, 0xea, 0x7a, 0x41, 0xa4, 0xfe, 0xa6, 0xaa, 0x40, 0x98, 0xa4, 0x28, 0xcb, 0xc4,
	0x92, 0xab, 0xe1, 0xa2, 0x10, 0x4a, 0x38, 0xb6, 0xe6, 0xbb, 0xfe, 0x3b, 0x7f, 0x26, 0xe6, 0x73,
	0xc1, 0xab, 0x61, 0x84, 0xdd, 0x4e, 0x2e, 0x72, 0xa1, 0xd7, 0xb0, 0xdc, 0x0c, 0xda, 0xbf, 0xad,
	0xc1, 0xd6, 0x69, 0x19, 0x1b, 0x99, 0x54, 0xe7, 0x2f, 0xb4, 0x91, 0x94, 0x44, 0xb9, 0xc0, 0x07,
	0xc1, 0xcf, 0xff, 0xed, 0x61, 0x15, 0x12, 0x95, 0x60, 0xdc, 0x58, 0x6f, 0x7b, 0x56, 0x62, 0x14,
	0xce, 0x18, 0xda, 0x4b, 0xce, 0x94, 0x74, 0x6b, 0x3e, 0x08, 0x9a, 0x71, 0x58, 0x72, 0x0f, 0xdb,
	0xde, 0x9f, 0x9c, 0x29, 0xba, 0x9c, 0x96, 0xc6, 0x30, 0x13, 0x72, 0x2e, 0x64, 0x35, 0x06, 0x12,
	0x5f, 0x98, 0xe7, 0x0c, 0x27, 0x8c, 0xab, 0xc4, 0xb8, 0x9d, 0x23, 0x68, 0x8b, 0x15, 0x27, 0x85,
	0x5b, 0xf7, 0x41, 0xd0, 0x8a, 0x47, 0x4f, 0xdb, 0xde, 0xe0, 0x13, 0x11, 0x51, 0x96, 0x45, 0x18,
	0x17, 0x44, 0xca, 0xc4, 0xf8, 0x9d, 0xdf, 0xf0, 0xd7, 0x0c, 0x49, 0x95, 0x22, 0x8c, 0x53, 0x4a,
	0x58, 0x4e, 0x95, 0xdb, 0xf0, 0x41, 0x50, 0x4f, 0xda, 0x25, 0x1c, 0x61, 0x7c, 0xac, 0x41, 0xe7,
	0x1f, 0xec, 0x68, 0xdd, 0x8a, 0x29, 0x8a, 0x0b, 0xb4, 0x7a, 0x11, 0xdb, 0x5a, 0xec, 0x94, 0xdc,
	0x59, 0x45, 0x19, 0x47, 0xff, 0x0e, 0xc0, 0xa6, 0xbe, 0xd2, 0x84, 0xb3, 0xef, 0x38, 0xd1, 0x18,
	0xda, 0x98, 0x2c, 0x14, 0xd5, 0x27, 0xfa, 0x4a, 0x8c, 0x76, 0xc7, 0x27, 0xeb, 0x9d, 0x07, 0x36,
	0x3b, 0x0f, 0x3c, 0xee, 0x3c, 0x70, 0xb3, 0xf7, 0xac, 0xcd, 0xde, 0xb3, 0xee, 0xf7, 0x9e, 0x75,
	0x1e, 0xe6, 0x4c, 0xcd, 0x90, 0x49, 0x7a, 0x53, 0x3d, 0x2a, 0x0a, 0x2e, 0x30, 0xf9, 0xd8, 0xc7,
	0xe9, 0x0f, 0x5d, 0x9f, 0x83, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0xd0, 0xb5, 0xe4, 0xb8,
	0x02, 0x00, 0x00,
}

func (m *TradeAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradeAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradeAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastWithdrawHeight != 0 {
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(m.LastWithdrawHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LastAddHeight != 0 {
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(m.LastAddHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Units.Size()
		i -= size
		if _, err := m.Units.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TradeUnit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradeUnit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradeUnit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Depth.Size()
		i -= size
		if _, err := m.Depth.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Units.Size()
		i -= size
		if _, err := m.Units.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeTradeAccount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypeTradeAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeTradeAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TradeAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Asset.Size()
	n += 1 + l + sovTypeTradeAccount(uint64(l))
	l = m.Units.Size()
	n += 1 + l + sovTypeTradeAccount(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypeTradeAccount(uint64(l))
	}
	if m.LastAddHeight != 0 {
		n += 1 + sovTypeTradeAccount(uint64(m.LastAddHeight))
	}
	if m.LastWithdrawHeight != 0 {
		n += 1 + sovTypeTradeAccount(uint64(m.LastWithdrawHeight))
	}
	return n
}

func (m *TradeUnit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Asset.Size()
	n += 1 + l + sovTypeTradeAccount(uint64(l))
	l = m.Units.Size()
	n += 1 + l + sovTypeTradeAccount(uint64(l))
	l = m.Depth.Size()
	n += 1 + l + sovTypeTradeAccount(uint64(l))
	return n
}

func sovTypeTradeAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeTradeAccount(x uint64) (n int) {
	return sovTypeTradeAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TradeAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTradeAccount
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
			return fmt.Errorf("proto: TradeAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradeAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
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
				return ErrInvalidLengthTypeTradeAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
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
				return ErrInvalidLengthTypeTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Units.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
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
				return ErrInvalidLengthTypeTradeAccount
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAddHeight", wireType)
			}
			m.LastAddHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastAddHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastWithdrawHeight", wireType)
			}
			m.LastWithdrawHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastWithdrawHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTradeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTradeAccount
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
func (m *TradeUnit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTradeAccount
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
			return fmt.Errorf("proto: TradeUnit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradeUnit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
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
				return ErrInvalidLengthTypeTradeAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
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
				return ErrInvalidLengthTypeTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Units.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTradeAccount
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
				return ErrInvalidLengthTypeTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Depth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTradeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTradeAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTradeAccount
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
func skipTypeTradeAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeTradeAccount
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
					return 0, ErrIntOverflowTypeTradeAccount
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
					return 0, ErrIntOverflowTypeTradeAccount
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
				return 0, ErrInvalidLengthTypeTradeAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeTradeAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeTradeAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeTradeAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeTradeAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeTradeAccount = fmt.Errorf("proto: unexpected end of group")
)
