// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: skillchain/skillchain/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	BurnEnabled          bool   `protobuf:"varint,1,opt,name=burn_enabled,json=burnEnabled,proto3" json:"burn_enabled,omitempty"`
	VusdEnabled          bool   `protobuf:"varint,2,opt,name=vusd_enabled,json=vusdEnabled,proto3" json:"vusd_enabled,omitempty"`
	PriceUpdateAuthority string `protobuf:"bytes,3,opt,name=price_update_authority,json=priceUpdateAuthority,proto3" json:"price_update_authority,omitempty"`
	VusdMockPrice        string `protobuf:"bytes,4,opt,name=vusd_mock_price,json=vusdMockPrice,proto3" json:"vusd_mock_price,omitempty"`
	TokenName            string `protobuf:"bytes,5,opt,name=token_name,json=tokenName,proto3" json:"token_name,omitempty"`
	TokenSymbol          string `protobuf:"bytes,6,opt,name=token_symbol,json=tokenSymbol,proto3" json:"token_symbol,omitempty"`
	TokenDecimals        uint32 `protobuf:"varint,7,opt,name=token_decimals,json=tokenDecimals,proto3" json:"token_decimals,omitempty"`
	TokenDescription     string `protobuf:"bytes,8,opt,name=token_description,json=tokenDescription,proto3" json:"token_description,omitempty"`
	MaxSupply            string `protobuf:"bytes,9,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	ChainDescription     string `protobuf:"bytes,10,opt,name=chain_description,json=chainDescription,proto3" json:"chain_description,omitempty"`
	WebsiteUrl           string `protobuf:"bytes,11,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	MinCollateralRatio   string `protobuf:"bytes,12,opt,name=min_collateral_ratio,json=minCollateralRatio,proto3" json:"min_collateral_ratio,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a950e2169aaf775, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBurnEnabled() bool {
	if m != nil {
		return m.BurnEnabled
	}
	return false
}

func (m *Params) GetVusdEnabled() bool {
	if m != nil {
		return m.VusdEnabled
	}
	return false
}

func (m *Params) GetPriceUpdateAuthority() string {
	if m != nil {
		return m.PriceUpdateAuthority
	}
	return ""
}

func (m *Params) GetVusdMockPrice() string {
	if m != nil {
		return m.VusdMockPrice
	}
	return ""
}

func (m *Params) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *Params) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *Params) GetTokenDecimals() uint32 {
	if m != nil {
		return m.TokenDecimals
	}
	return 0
}

func (m *Params) GetTokenDescription() string {
	if m != nil {
		return m.TokenDescription
	}
	return ""
}

func (m *Params) GetMaxSupply() string {
	if m != nil {
		return m.MaxSupply
	}
	return ""
}

func (m *Params) GetChainDescription() string {
	if m != nil {
		return m.ChainDescription
	}
	return ""
}

func (m *Params) GetWebsiteUrl() string {
	if m != nil {
		return m.WebsiteUrl
	}
	return ""
}

func (m *Params) GetMinCollateralRatio() string {
	if m != nil {
		return m.MinCollateralRatio
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "skillchain.skillchain.Params")
}

func init() {
	proto.RegisterFile("skillchain/skillchain/params.proto", fileDescriptor_0a950e2169aaf775)
}

var fileDescriptor_0a950e2169aaf775 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x6a, 0x53, 0x41,
	0x14, 0xc7, 0x33, 0x5a, 0x63, 0x33, 0x69, 0xd4, 0x0e, 0x51, 0x86, 0x42, 0x6f, 0x63, 0x41, 0x0d,
	0x0a, 0x8d, 0xa0, 0x20, 0xb8, 0xf3, 0x6b, 0xa9, 0x94, 0x94, 0x6e, 0xdc, 0x0c, 0x73, 0x6f, 0x86,
	0x76, 0xc8, 0x7c, 0x31, 0x33, 0x57, 0x93, 0x57, 0xe8, 0xca, 0x47, 0xf0, 0x11, 0x7c, 0x0c, 0x97,
	0x5d, 0xba, 0x94, 0x64, 0xa1, 0x8f, 0x21, 0x73, 0x26, 0x49, 0xaf, 0x8b, 0x6e, 0x2e, 0x87, 0xdf,
	0xff, 0x77, 0xff, 0x1c, 0x98, 0x83, 0x0f, 0xc3, 0x54, 0x2a, 0x55, 0x9d, 0x73, 0x69, 0x46, 0x8d,
	0xd1, 0x71, 0xcf, 0x75, 0x38, 0x72, 0xde, 0x46, 0x4b, 0xee, 0x5f, 0x05, 0x47, 0x57, 0xe3, 0xde,
	0x2e, 0xd7, 0xd2, 0xd8, 0x11, 0x7c, 0xb3, 0xb9, 0xd7, 0x3f, 0xb3, 0x67, 0x16, 0xc6, 0x51, 0x9a,
	0x32, 0x3d, 0xbc, 0xd8, 0xc2, 0xed, 0x63, 0x28, 0x24, 0x0f, 0xf1, 0x4e, 0x59, 0x7b, 0xc3, 0x84,
	0xe1, 0xa5, 0x12, 0x13, 0x8a, 0x06, 0x68, 0xb8, 0x3d, 0xee, 0x26, 0xf6, 0x21, 0xa3, 0xa4, 0x7c,
	0xa9, 0xc3, 0x64, 0xa3, 0xdc, 0xc8, 0x4a, 0x62, 0x6b, 0xe5, 0x25, 0x7e, 0xe0, 0xbc, 0xac, 0x04,
	0xab, 0xdd, 0x84, 0x47, 0xc1, 0x78, 0x1d, 0xcf, 0xad, 0x97, 0x71, 0x4e, 0x6f, 0x0e, 0xd0, 0xb0,
	0x33, 0xee, 0x43, 0x7a, 0x0a, 0xe1, 0x9b, 0x75, 0x46, 0x1e, 0xe3, 0xbb, 0x50, 0xac, 0x6d, 0x35,
	0x65, 0x60, 0xd0, 0x2d, 0xd0, 0x7b, 0x09, 0x7f, 0xb4, 0xd5, 0xf4, 0x38, 0x41, 0xb2, 0x8f, 0x71,
	0xb4, 0x53, 0x61, 0x98, 0xe1, 0x5a, 0xd0, 0x5b, 0xa0, 0x74, 0x80, 0x7c, 0xe2, 0x5a, 0xa4, 0xfd,
	0x72, 0x1c, 0xe6, 0xba, 0xb4, 0x8a, 0xb6, 0x41, 0xe8, 0x02, 0x3b, 0x01, 0x44, 0x1e, 0xe1, 0x3b,
	0x59, 0x99, 0x88, 0x4a, 0x6a, 0xae, 0x02, 0xbd, 0x3d, 0x40, 0xc3, 0xde, 0xb8, 0x07, 0xf4, 0xfd,
	0x0a, 0x92, 0x67, 0x78, 0x77, 0xad, 0x85, 0xca, 0x4b, 0x17, 0xa5, 0x35, 0x74, 0x1b, 0xea, 0xee,
	0xad, 0xcc, 0x0d, 0x4f, 0x5b, 0x69, 0x3e, 0x63, 0xa1, 0x76, 0x4e, 0xcd, 0x69, 0x27, 0x6f, 0xa5,
	0xf9, 0xec, 0x04, 0x40, 0xea, 0x82, 0x57, 0xf9, 0xaf, 0x0b, 0xe7, 0x2e, 0x08, 0x9a, 0x5d, 0x07,
	0xb8, 0xfb, 0x55, 0x94, 0x41, 0x46, 0xc1, 0x6a, 0xaf, 0x68, 0x17, 0x34, 0xbc, 0x42, 0xa7, 0x5e,
	0x91, 0xe7, 0xb8, 0xaf, 0xa5, 0x61, 0x95, 0x55, 0x8a, 0x47, 0xe1, 0xb9, 0x62, 0x9e, 0x47, 0x69,
	0xe9, 0x0e, 0x98, 0x44, 0x4b, 0xf3, 0x6e, 0x13, 0x8d, 0x53, 0xf2, 0xfa, 0xc9, 0xdf, 0xef, 0x07,
	0xe8, 0xe2, 0xcf, 0x8f, 0xa7, 0x45, 0xe3, 0x8a, 0x66, 0xcd, 0x93, 0xca, 0x17, 0xf0, 0xf6, 0xd5,
	0xcf, 0x45, 0x81, 0x2e, 0x17, 0x05, 0xfa, 0xbd, 0x28, 0xd0, 0xb7, 0x65, 0xd1, 0xba, 0x5c, 0x16,
	0xad, 0x5f, 0xcb, 0xa2, 0xf5, 0x79, 0xff, 0xba, 0x3f, 0xe3, 0xdc, 0x89, 0x50, 0xb6, 0xe1, 0x98,
	0x5e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x27, 0x03, 0x4b, 0xb2, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.BurnEnabled != that1.BurnEnabled {
		return false
	}
	if this.VusdEnabled != that1.VusdEnabled {
		return false
	}
	if this.PriceUpdateAuthority != that1.PriceUpdateAuthority {
		return false
	}
	if this.VusdMockPrice != that1.VusdMockPrice {
		return false
	}
	if this.TokenName != that1.TokenName {
		return false
	}
	if this.TokenSymbol != that1.TokenSymbol {
		return false
	}
	if this.TokenDecimals != that1.TokenDecimals {
		return false
	}
	if this.TokenDescription != that1.TokenDescription {
		return false
	}
	if this.MaxSupply != that1.MaxSupply {
		return false
	}
	if this.ChainDescription != that1.ChainDescription {
		return false
	}
	if this.WebsiteUrl != that1.WebsiteUrl {
		return false
	}
	if this.MinCollateralRatio != that1.MinCollateralRatio {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinCollateralRatio) > 0 {
		i -= len(m.MinCollateralRatio)
		copy(dAtA[i:], m.MinCollateralRatio)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MinCollateralRatio)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.WebsiteUrl) > 0 {
		i -= len(m.WebsiteUrl)
		copy(dAtA[i:], m.WebsiteUrl)
		i = encodeVarintParams(dAtA, i, uint64(len(m.WebsiteUrl)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ChainDescription) > 0 {
		i -= len(m.ChainDescription)
		copy(dAtA[i:], m.ChainDescription)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ChainDescription)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.MaxSupply) > 0 {
		i -= len(m.MaxSupply)
		copy(dAtA[i:], m.MaxSupply)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MaxSupply)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.TokenDescription) > 0 {
		i -= len(m.TokenDescription)
		copy(dAtA[i:], m.TokenDescription)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenDescription)))
		i--
		dAtA[i] = 0x42
	}
	if m.TokenDecimals != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TokenDecimals))
		i--
		dAtA[i] = 0x38
	}
	if len(m.TokenSymbol) > 0 {
		i -= len(m.TokenSymbol)
		copy(dAtA[i:], m.TokenSymbol)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenSymbol)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenName) > 0 {
		i -= len(m.TokenName)
		copy(dAtA[i:], m.TokenName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.VusdMockPrice) > 0 {
		i -= len(m.VusdMockPrice)
		copy(dAtA[i:], m.VusdMockPrice)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VusdMockPrice)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PriceUpdateAuthority) > 0 {
		i -= len(m.PriceUpdateAuthority)
		copy(dAtA[i:], m.PriceUpdateAuthority)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceUpdateAuthority)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VusdEnabled {
		i--
		if m.VusdEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.BurnEnabled {
		i--
		if m.BurnEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BurnEnabled {
		n += 2
	}
	if m.VusdEnabled {
		n += 2
	}
	l = len(m.PriceUpdateAuthority)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.VusdMockPrice)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TokenName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TokenSymbol)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.TokenDecimals != 0 {
		n += 1 + sovParams(uint64(m.TokenDecimals))
	}
	l = len(m.TokenDescription)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MaxSupply)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ChainDescription)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.WebsiteUrl)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.MinCollateralRatio)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.BurnEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VusdEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.VusdEnabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceUpdateAuthority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceUpdateAuthority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VusdMockPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VusdMockPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenDecimals", wireType)
			}
			m.TokenDecimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenDecimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenDescription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenDescription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxSupply = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainDescription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainDescription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WebsiteUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WebsiteUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCollateralRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinCollateralRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
