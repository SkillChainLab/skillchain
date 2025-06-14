// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: skillchain/marketplace/job_posting.proto

package types

import (
	fmt "fmt"
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

type JobPosting struct {
	Index           string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ClientAddress   string `protobuf:"bytes,2,opt,name=clientAddress,proto3" json:"clientAddress,omitempty"`
	Title           string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description     string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	SkillsRequired  string `protobuf:"bytes,5,opt,name=skillsRequired,proto3" json:"skillsRequired,omitempty"`
	BudgetAmount    uint64 `protobuf:"varint,6,opt,name=budgetAmount,proto3" json:"budgetAmount,omitempty"`
	PaymentCurrency string `protobuf:"bytes,7,opt,name=paymentCurrency,proto3" json:"paymentCurrency,omitempty"`
	Deadline        int64  `protobuf:"varint,8,opt,name=deadline,proto3" json:"deadline,omitempty"`
	IsActive        bool   `protobuf:"varint,9,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CreatedAt       int64  `protobuf:"varint,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Creator         string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *JobPosting) Reset()         { *m = JobPosting{} }
func (m *JobPosting) String() string { return proto.CompactTextString(m) }
func (*JobPosting) ProtoMessage()    {}
func (*JobPosting) Descriptor() ([]byte, []int) {
	return fileDescriptor_8190ebda28bf52f1, []int{0}
}
func (m *JobPosting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JobPosting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JobPosting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JobPosting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobPosting.Merge(m, src)
}
func (m *JobPosting) XXX_Size() int {
	return m.Size()
}
func (m *JobPosting) XXX_DiscardUnknown() {
	xxx_messageInfo_JobPosting.DiscardUnknown(m)
}

var xxx_messageInfo_JobPosting proto.InternalMessageInfo

func (m *JobPosting) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *JobPosting) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (m *JobPosting) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *JobPosting) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *JobPosting) GetSkillsRequired() string {
	if m != nil {
		return m.SkillsRequired
	}
	return ""
}

func (m *JobPosting) GetBudgetAmount() uint64 {
	if m != nil {
		return m.BudgetAmount
	}
	return 0
}

func (m *JobPosting) GetPaymentCurrency() string {
	if m != nil {
		return m.PaymentCurrency
	}
	return ""
}

func (m *JobPosting) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

func (m *JobPosting) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *JobPosting) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *JobPosting) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*JobPosting)(nil), "skillchain.marketplace.JobPosting")
}

func init() {
	proto.RegisterFile("skillchain/marketplace/job_posting.proto", fileDescriptor_8190ebda28bf52f1)
}

var fileDescriptor_8190ebda28bf52f1 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x9b, 0xfe, 0x37, 0xfd, 0x3e, 0x85, 0x20, 0x12, 0x44, 0xc2, 0x50, 0x44, 0x66, 0xd5,
	0x2e, 0xdc, 0xb8, 0x1d, 0xdd, 0xb9, 0x92, 0x59, 0xba, 0x91, 0x99, 0xe4, 0x50, 0x63, 0xa7, 0xc9,
	0x98, 0x9c, 0x91, 0xf6, 0x2e, 0xbc, 0x2c, 0x97, 0x5d, 0xba, 0x94, 0x16, 0xbc, 0x0e, 0x69, 0x8a,
	0xfd, 0x5b, 0x3e, 0xcf, 0x79, 0xdf, 0x99, 0xc0, 0x4b, 0x63, 0x3f, 0xd1, 0x45, 0x21, 0x5f, 0x32,
	0x6d, 0x46, 0xd3, 0xcc, 0x4d, 0x00, 0xcb, 0x22, 0x93, 0x30, 0x7a, 0xb5, 0xf9, 0x73, 0x69, 0x3d,
	0x6a, 0x33, 0x1e, 0x96, 0xce, 0xa2, 0x65, 0xe7, 0xbb, 0xe4, 0x70, 0x2f, 0x39, 0xf8, 0xa9, 0x53,
	0xfa, 0x60, 0xf3, 0xc7, 0x4d, 0x98, 0x9d, 0xd1, 0x96, 0x36, 0x0a, 0x66, 0x9c, 0x44, 0x24, 0xee,
	0xa5, 0x1b, 0x60, 0x57, 0xf4, 0xbf, 0x2c, 0x34, 0x18, 0x4c, 0x94, 0x72, 0xe0, 0x3d, 0xaf, 0x87,
	0xeb, 0xa1, 0x5c, 0x77, 0x51, 0x63, 0x01, 0xbc, 0xb1, 0xe9, 0x06, 0x60, 0x11, 0xed, 0x2b, 0xf0,
	0xd2, 0xe9, 0x12, 0xb5, 0x35, 0xbc, 0x19, 0x6e, 0xfb, 0x8a, 0x5d, 0xd3, 0x93, 0xf0, 0x38, 0x9f,
	0xc2, 0x5b, 0xa5, 0x1d, 0x28, 0xde, 0x0a, 0xa1, 0x23, 0xcb, 0x06, 0xf4, 0x5f, 0x5e, 0xa9, 0x31,
	0x60, 0x32, 0xb5, 0x95, 0x41, 0xde, 0x8e, 0x48, 0xdc, 0x4c, 0x0f, 0x1c, 0x8b, 0xe9, 0x69, 0x99,
	0xcd, 0xa7, 0x60, 0xf0, 0xbe, 0x72, 0x0e, 0x8c, 0x9c, 0xf3, 0x4e, 0xf8, 0xd8, 0xb1, 0x66, 0x17,
	0xb4, 0xab, 0x20, 0x53, 0x85, 0x36, 0xc0, 0xbb, 0x11, 0x89, 0x1b, 0xe9, 0x96, 0xd7, 0x37, 0xed,
	0x13, 0x89, 0xfa, 0x1d, 0x78, 0x2f, 0x22, 0x71, 0x37, 0xdd, 0x32, 0xbb, 0xa4, 0x3d, 0xe9, 0x20,
	0x43, 0x50, 0x09, 0x72, 0x1a, 0x8a, 0x3b, 0xc1, 0x38, 0xed, 0x04, 0xb0, 0x8e, 0xf7, 0xc3, 0x7f,
	0xff, 0xf0, 0xee, 0xf6, 0x73, 0x29, 0xc8, 0x62, 0x29, 0xc8, 0xf7, 0x52, 0x90, 0x8f, 0x95, 0xa8,
	0x2d, 0x56, 0xa2, 0xf6, 0xb5, 0x12, 0xb5, 0x27, 0xb1, 0x37, 0xe2, 0xec, 0x60, 0x46, 0x9c, 0x97,
	0xe0, 0xf3, 0x76, 0x58, 0xf0, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0x63, 0x0e, 0x1d, 0x31, 0xed,
	0x01, 0x00, 0x00,
}

func (m *JobPosting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobPosting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JobPosting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if m.CreatedAt != 0 {
		i = encodeVarintJobPosting(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x50
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Deadline != 0 {
		i = encodeVarintJobPosting(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x40
	}
	if len(m.PaymentCurrency) > 0 {
		i -= len(m.PaymentCurrency)
		copy(dAtA[i:], m.PaymentCurrency)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.PaymentCurrency)))
		i--
		dAtA[i] = 0x3a
	}
	if m.BudgetAmount != 0 {
		i = encodeVarintJobPosting(dAtA, i, uint64(m.BudgetAmount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SkillsRequired) > 0 {
		i -= len(m.SkillsRequired)
		copy(dAtA[i:], m.SkillsRequired)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.SkillsRequired)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClientAddress) > 0 {
		i -= len(m.ClientAddress)
		copy(dAtA[i:], m.ClientAddress)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.ClientAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintJobPosting(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJobPosting(dAtA []byte, offset int, v uint64) int {
	offset -= sovJobPosting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *JobPosting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	l = len(m.ClientAddress)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	l = len(m.SkillsRequired)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	if m.BudgetAmount != 0 {
		n += 1 + sovJobPosting(uint64(m.BudgetAmount))
	}
	l = len(m.PaymentCurrency)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	if m.Deadline != 0 {
		n += 1 + sovJobPosting(uint64(m.Deadline))
	}
	if m.IsActive {
		n += 2
	}
	if m.CreatedAt != 0 {
		n += 1 + sovJobPosting(uint64(m.CreatedAt))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovJobPosting(uint64(l))
	}
	return n
}

func sovJobPosting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJobPosting(x uint64) (n int) {
	return sovJobPosting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JobPosting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobPosting
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
			return fmt.Errorf("proto: JobPosting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobPosting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkillsRequired", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SkillsRequired = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BudgetAmount", wireType)
			}
			m.BudgetAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BudgetAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentCurrency", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentCurrency = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
			m.IsActive = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobPosting
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
				return ErrInvalidLengthJobPosting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJobPosting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobPosting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJobPosting
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
func skipJobPosting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJobPosting
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
					return 0, ErrIntOverflowJobPosting
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
					return 0, ErrIntOverflowJobPosting
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
				return 0, ErrInvalidLengthJobPosting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJobPosting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJobPosting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJobPosting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJobPosting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJobPosting = fmt.Errorf("proto: unexpected end of group")
)
