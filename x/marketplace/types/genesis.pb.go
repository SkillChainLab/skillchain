// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: skillchain/marketplace/genesis.proto

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

// GenesisState defines the marketplace module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params         Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	JobPostingList []JobPosting `protobuf:"bytes,2,rep,name=jobPostingList,proto3" json:"jobPostingList"`
	ProposalList   []Proposal   `protobuf:"bytes,3,rep,name=proposalList,proto3" json:"proposalList"`
	ProjectList    []Project    `protobuf:"bytes,4,rep,name=projectList,proto3" json:"projectList"`
	MilestoneList  []Milestone  `protobuf:"bytes,5,rep,name=milestoneList,proto3" json:"milestoneList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0debdc95c7d5ed89, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetJobPostingList() []JobPosting {
	if m != nil {
		return m.JobPostingList
	}
	return nil
}

func (m *GenesisState) GetProposalList() []Proposal {
	if m != nil {
		return m.ProposalList
	}
	return nil
}

func (m *GenesisState) GetProjectList() []Project {
	if m != nil {
		return m.ProjectList
	}
	return nil
}

func (m *GenesisState) GetMilestoneList() []Milestone {
	if m != nil {
		return m.MilestoneList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "skillchain.marketplace.GenesisState")
}

func init() {
	proto.RegisterFile("skillchain/marketplace/genesis.proto", fileDescriptor_0debdc95c7d5ed89)
}

var fileDescriptor_0debdc95c7d5ed89 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4a, 0x03, 0x31,
	0x18, 0xc7, 0x2f, 0xb6, 0x16, 0x4c, 0xab, 0x60, 0x10, 0x29, 0x1d, 0xd2, 0x5a, 0xab, 0x14, 0x87,
	0x3b, 0xa8, 0x8b, 0xab, 0x5d, 0x0a, 0xc5, 0x42, 0xa9, 0x9b, 0x8b, 0xa4, 0x25, 0x9c, 0x69, 0xef,
	0x2e, 0xe1, 0x92, 0x41, 0x9f, 0xc0, 0xd5, 0xc7, 0x70, 0xf4, 0x31, 0x3a, 0x76, 0x74, 0x12, 0xe9,
	0x0d, 0xbe, 0x86, 0x5c, 0x2e, 0xad, 0x77, 0x62, 0x5c, 0x8e, 0x70, 0xfc, 0xfe, 0xbf, 0x7c, 0xf9,
	0x7f, 0xb0, 0x23, 0x17, 0x2c, 0x08, 0x66, 0x0f, 0x84, 0x45, 0x5e, 0x48, 0xe2, 0x05, 0x55, 0x22,
	0x20, 0x33, 0xea, 0xf9, 0x34, 0xa2, 0x92, 0x49, 0x57, 0xc4, 0x5c, 0x71, 0x74, 0xfc, 0x43, 0xb9,
	0x39, 0xaa, 0x71, 0x48, 0x42, 0x16, 0x71, 0x4f, 0x7f, 0x33, 0xb4, 0x71, 0xe4, 0x73, 0x9f, 0xeb,
	0xa3, 0x97, 0x9e, 0xcc, 0xdf, 0x53, 0xcb, 0x35, 0x82, 0xc4, 0x24, 0x34, 0xb7, 0x34, 0xba, 0x16,
	0x68, 0xce, 0xa7, 0xf7, 0x82, 0x4b, 0xc5, 0x22, 0xdf, 0x90, 0x67, 0x36, 0x5d, 0xcc, 0x05, 0x97,
	0x24, 0x30, 0x58, 0xc7, 0x8e, 0xcd, 0xe9, 0x4c, 0x19, 0xea, 0xdc, 0x42, 0x85, 0x2c, 0xa0, 0x52,
	0xf1, 0x88, 0x66, 0x5c, 0xfb, 0xb9, 0x04, 0x6b, 0x83, 0xac, 0x96, 0x5b, 0x45, 0x14, 0x45, 0xd7,
	0xb0, 0x92, 0xcd, 0x5f, 0x07, 0x2d, 0xd0, 0xad, 0xf6, 0xb0, 0xfb, 0x77, 0x4d, 0xee, 0x58, 0x53,
	0xfd, 0xbd, 0xe5, 0x47, 0xd3, 0x79, 0xfd, 0x7a, 0xbb, 0x00, 0x13, 0x13, 0x44, 0x63, 0x78, 0x30,
	0xe7, 0xd3, 0x71, 0xf6, 0xb8, 0x1b, 0x26, 0x55, 0x7d, 0xa7, 0x55, 0xea, 0x56, 0x7b, 0x6d, 0x9b,
	0x6a, 0xb8, 0xa5, 0xfb, 0xe5, 0x54, 0x37, 0xf9, 0x95, 0x47, 0x43, 0x58, 0xdb, 0xb4, 0xa0, 0x7d,
	0x25, 0xed, 0x6b, 0x59, 0x47, 0x33, 0xac, 0xb1, 0x15, 0xb2, 0x68, 0x00, 0xab, 0xa6, 0x2a, 0xad,
	0x2a, 0x6b, 0x55, 0xf3, 0x1f, 0x55, 0x8a, 0x1a, 0x53, 0x3e, 0x89, 0x46, 0x70, 0x7f, 0xdb, 0xa6,
	0x56, 0xed, 0x6a, 0xd5, 0x89, 0x4d, 0x35, 0xda, 0xc0, 0x46, 0x56, 0x4c, 0xf7, 0xaf, 0x96, 0x6b,
	0x0c, 0x56, 0x6b, 0x0c, 0x3e, 0xd7, 0x18, 0xbc, 0x24, 0xd8, 0x59, 0x25, 0xd8, 0x79, 0x4f, 0xb0,
	0x73, 0x87, 0x73, 0xbb, 0x7c, 0x2c, 0x6c, 0x53, 0x3d, 0x09, 0x2a, 0xa7, 0x15, 0xbd, 0xca, 0xcb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x8d, 0xa4, 0x4e, 0xf7, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MilestoneList) > 0 {
		for iNdEx := len(m.MilestoneList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MilestoneList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ProjectList) > 0 {
		for iNdEx := len(m.ProjectList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProposalList) > 0 {
		for iNdEx := len(m.ProposalList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposalList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.JobPostingList) > 0 {
		for iNdEx := len(m.JobPostingList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.JobPostingList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.JobPostingList) > 0 {
		for _, e := range m.JobPostingList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposalList) > 0 {
		for _, e := range m.ProposalList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProjectList) > 0 {
		for _, e := range m.ProjectList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MilestoneList) > 0 {
		for _, e := range m.MilestoneList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobPostingList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobPostingList = append(m.JobPostingList, JobPosting{})
			if err := m.JobPostingList[len(m.JobPostingList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalList = append(m.ProposalList, Proposal{})
			if err := m.ProposalList[len(m.ProposalList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectList = append(m.ProjectList, Project{})
			if err := m.ProjectList[len(m.ProjectList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MilestoneList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MilestoneList = append(m.MilestoneList, Milestone{})
			if err := m.MilestoneList[len(m.MilestoneList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
