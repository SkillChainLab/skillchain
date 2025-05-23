// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: skillchain/job/application.proto

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

type Application struct {
	JobId          uint64 `protobuf:"varint,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	Applicant      string `protobuf:"bytes,2,opt,name=applicant,proto3" json:"applicant,omitempty"`
	CoverLetter    string `protobuf:"bytes,3,opt,name=coverLetter,proto3" json:"coverLetter,omitempty"`
	Status         string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	JobTitle       string `protobuf:"bytes,5,opt,name=jobTitle,proto3" json:"jobTitle,omitempty"`
	JobDescription string `protobuf:"bytes,6,opt,name=jobDescription,proto3" json:"jobDescription,omitempty"`
	JobBudget      string `protobuf:"bytes,7,opt,name=jobBudget,proto3" json:"jobBudget,omitempty"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_dda493dc9bbebcea, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Application.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return m.Size()
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetJobId() uint64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *Application) GetApplicant() string {
	if m != nil {
		return m.Applicant
	}
	return ""
}

func (m *Application) GetCoverLetter() string {
	if m != nil {
		return m.CoverLetter
	}
	return ""
}

func (m *Application) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Application) GetJobTitle() string {
	if m != nil {
		return m.JobTitle
	}
	return ""
}

func (m *Application) GetJobDescription() string {
	if m != nil {
		return m.JobDescription
	}
	return ""
}

func (m *Application) GetJobBudget() string {
	if m != nil {
		return m.JobBudget
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "skillchain.job.Application")
}

func init() { proto.RegisterFile("skillchain/job/application.proto", fileDescriptor_dda493dc9bbebcea) }

var fileDescriptor_dda493dc9bbebcea = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x1b, 0xdd, 0xaa, 0xcb, 0x60, 0x60, 0x18, 0x12, 0x8a, 0x84, 0xe2, 0x41, 0x76, 0x6a,
	0x0e, 0x3e, 0x81, 0xd3, 0xcb, 0x60, 0xa7, 0xe9, 0xc9, 0x5b, 0xd2, 0x85, 0x2e, 0xb1, 0xeb, 0xaf,
	0xb4, 0xa9, 0xe8, 0x5b, 0xf8, 0x58, 0x1e, 0x77, 0xf4, 0x22, 0x48, 0xfb, 0x22, 0xd2, 0x74, 0xac,
	0x63, 0x97, 0xf2, 0xfb, 0xfe, 0xa1, 0xf9, 0xf2, 0xc1, 0x61, 0xf9, 0xa6, 0xd3, 0x34, 0xde, 0x08,
	0x9d, 0x71, 0x03, 0x92, 0x8b, 0x3c, 0x4f, 0x75, 0x2c, 0xac, 0x86, 0x2c, 0xca, 0x0b, 0xb0, 0x40,
	0x26, 0x7d, 0x23, 0x32, 0x20, 0x83, 0x2b, 0xb1, 0xd5, 0x19, 0x70, 0xf7, 0xed, 0x2a, 0xc1, 0x34,
	0x81, 0x04, 0xdc, 0xc9, 0xdb, 0xab, 0x73, 0x6f, 0x7f, 0x11, 0x1e, 0x3f, 0xf4, 0xbf, 0x23, 0x53,
	0x3c, 0x34, 0x20, 0x17, 0x6b, 0x8a, 0x42, 0x34, 0x1b, 0xac, 0x3a, 0x41, 0x6e, 0xf0, 0x68, 0xff,
	0x66, 0x66, 0xe9, 0x59, 0x88, 0x66, 0xa3, 0x55, 0x6f, 0x90, 0x10, 0x8f, 0x63, 0x78, 0x57, 0xc5,
	0x52, 0x59, 0xab, 0x0a, 0x7a, 0xee, 0xf2, 0x63, 0x8b, 0x5c, 0x63, 0xbf, 0xb4, 0xc2, 0x56, 0x25,
	0x1d, 0xb8, 0x70, 0xaf, 0x48, 0x80, 0x2f, 0x0d, 0xc8, 0x17, 0x6d, 0x53, 0x45, 0x87, 0x2e, 0x39,
	0x68, 0x72, 0x87, 0x27, 0x06, 0xe4, 0x93, 0x2a, 0xe3, 0x42, 0xe7, 0xed, 0x36, 0xea, 0xbb, 0xc6,
	0x89, 0xdb, 0x6e, 0x33, 0x20, 0xe7, 0xd5, 0x3a, 0x51, 0x96, 0x5e, 0x74, 0xdb, 0x0e, 0xc6, 0x7c,
	0xf1, 0x5d, 0x33, 0xb4, 0xab, 0x19, 0xfa, 0xab, 0x19, 0xfa, 0x6a, 0x98, 0xb7, 0x6b, 0x98, 0xf7,
	0xd3, 0x30, 0xef, 0x95, 0x27, 0xda, 0x6e, 0x2a, 0x19, 0xc5, 0xb0, 0xe5, 0xcf, 0x2d, 0xbd, 0xc7,
	0x96, 0xde, 0x52, 0x48, 0x7e, 0x44, 0xfb, 0xc3, 0xf1, 0xb6, 0x9f, 0xb9, 0x2a, 0xa5, 0xef, 0x88,
	0xdd, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x15, 0x47, 0x9c, 0xdc, 0x8e, 0x01, 0x00, 0x00,
}

func (m *Application) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Application) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Application) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JobBudget) > 0 {
		i -= len(m.JobBudget)
		copy(dAtA[i:], m.JobBudget)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.JobBudget)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.JobDescription) > 0 {
		i -= len(m.JobDescription)
		copy(dAtA[i:], m.JobDescription)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.JobDescription)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.JobTitle) > 0 {
		i -= len(m.JobTitle)
		copy(dAtA[i:], m.JobTitle)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.JobTitle)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CoverLetter) > 0 {
		i -= len(m.CoverLetter)
		copy(dAtA[i:], m.CoverLetter)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.CoverLetter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Applicant) > 0 {
		i -= len(m.Applicant)
		copy(dAtA[i:], m.Applicant)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Applicant)))
		i--
		dAtA[i] = 0x12
	}
	if m.JobId != 0 {
		i = encodeVarintApplication(dAtA, i, uint64(m.JobId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	offset -= sovApplication(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Application) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JobId != 0 {
		n += 1 + sovApplication(uint64(m.JobId))
	}
	l = len(m.Applicant)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.CoverLetter)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.JobTitle)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.JobDescription)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.JobBudget)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	return n
}

func sovApplication(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Application) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: Application: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Application: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			m.JobId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Applicant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Applicant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoverLetter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoverLetter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobTitle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobTitle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobDescription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobDescription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobBudget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobBudget = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApplication
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
func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
				return 0, ErrInvalidLengthApplication
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApplication
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApplication
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApplication        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApplication = fmt.Errorf("proto: unexpected end of group")
)
