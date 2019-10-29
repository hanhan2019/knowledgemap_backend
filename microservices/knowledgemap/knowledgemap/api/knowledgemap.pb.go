// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/knowledgemap.proto

/*
	Package api is a generated protocol buffer package.

	It is generated from these files:
		api/knowledgemap.proto

	It has these top-level messages:
		Empty
		UserReq
		CRqQueryMapBySubject
		KnowledegeMapInfo
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{0} }

type UserReq struct {
	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"uid"`
}

func (m *UserReq) Reset()                    { *m = UserReq{} }
func (m *UserReq) String() string            { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()               {}
func (*UserReq) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{1} }

func (m *UserReq) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

type CRqQueryMapBySubject struct {
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
}

func (m *CRqQueryMapBySubject) Reset()                    { *m = CRqQueryMapBySubject{} }
func (m *CRqQueryMapBySubject) String() string            { return proto.CompactTextString(m) }
func (*CRqQueryMapBySubject) ProtoMessage()               {}
func (*CRqQueryMapBySubject) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{2} }

func (m *CRqQueryMapBySubject) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type KnowledegeMapInfo struct {
	Knowledgemap string `protobuf:"bytes,1,opt,name=knowledgemap,proto3" json:"knowledgemap"`
}

func (m *KnowledegeMapInfo) Reset()                    { *m = KnowledegeMapInfo{} }
func (m *KnowledegeMapInfo) String() string            { return proto.CompactTextString(m) }
func (*KnowledegeMapInfo) ProtoMessage()               {}
func (*KnowledegeMapInfo) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{3} }

func (m *KnowledegeMapInfo) GetKnowledgemap() string {
	if m != nil {
		return m.Knowledgemap
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*UserReq)(nil), "api.UserReq")
	proto.RegisterType((*CRqQueryMapBySubject)(nil), "api.CRqQueryMapBySubject")
	proto.RegisterType((*KnowledegeMapInfo)(nil), "api.KnowledegeMapInfo")
}
func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Userid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Userid)))
		i += copy(dAtA[i:], m.Userid)
	}
	return i, nil
}

func (m *CRqQueryMapBySubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CRqQueryMapBySubject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Subject) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	return i, nil
}

func (m *KnowledegeMapInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KnowledegeMapInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Knowledgemap) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Knowledgemap)))
		i += copy(dAtA[i:], m.Knowledgemap)
	}
	return i, nil
}

func encodeVarintKnowledgemap(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *UserReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Userid)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *CRqQueryMapBySubject) Size() (n int) {
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *KnowledegeMapInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Knowledgemap)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func sovKnowledgemap(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKnowledgemap(x uint64) (n int) {
	return sovKnowledgemap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *UserReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
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
			return fmt.Errorf("proto: UserReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
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
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *CRqQueryMapBySubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
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
			return fmt.Errorf("proto: CRqQueryMapBySubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CRqQueryMapBySubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
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
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *KnowledegeMapInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
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
			return fmt.Errorf("proto: KnowledegeMapInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KnowledegeMapInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Knowledgemap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
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
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Knowledgemap = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func skipKnowledgemap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKnowledgemap
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
					return 0, ErrIntOverflowKnowledgemap
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
					return 0, ErrIntOverflowKnowledgemap
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
				return 0, ErrInvalidLengthKnowledgemap
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKnowledgemap
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
				next, err := skipKnowledgemap(dAtA[start:])
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
	ErrInvalidLengthKnowledgemap = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKnowledgemap   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/knowledgemap.proto", fileDescriptorKnowledgemap) }

var fileDescriptorKnowledgemap = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x49, 0x4d, 0x49, 0x4f, 0xcd, 0x4d, 0x2c, 0xd0, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x25, 0x95, 0xa6, 0x81, 0x79,
	0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa3, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0xa9, 0xa4,
	0xc5, 0xc5, 0x1e, 0x5a, 0x9c, 0x5a, 0x14, 0x94, 0x5a, 0x28, 0x24, 0xcf, 0xc5, 0x56, 0x5a, 0x9c,
	0x5a, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xfe, 0xea, 0x9e, 0x3c, 0x73,
	0x69, 0x66, 0x4a, 0x10, 0x54, 0x58, 0xc9, 0x96, 0x4b, 0xc4, 0x39, 0xa8, 0x30, 0xb0, 0x34, 0xb5,
	0xa8, 0xd2, 0x37, 0xb1, 0xc0, 0xa9, 0x32, 0xb8, 0x34, 0x29, 0x2b, 0x35, 0xb9, 0x44, 0x48, 0x95,
	0x8b, 0xbd, 0x18, 0xc2, 0x84, 0xea, 0xe4, 0x7e, 0x75, 0x4f, 0x1e, 0x26, 0x14, 0x04, 0x63, 0x28,
	0x79, 0x72, 0x09, 0x7a, 0x43, 0x5c, 0x9f, 0x9a, 0x9e, 0xea, 0x9b, 0x58, 0xe0, 0x99, 0x97, 0x96,
	0x2f, 0x64, 0xc2, 0xc5, 0x83, 0xec, 0x25, 0xa8, 0x01, 0x02, 0xaf, 0xee, 0xc9, 0xa3, 0x88, 0x07,
	0xa1, 0xf0, 0x8c, 0x12, 0xb9, 0x78, 0x51, 0x8c, 0x12, 0x0a, 0xe0, 0x92, 0x74, 0x4f, 0x2d, 0x41,
	0x11, 0x43, 0xb8, 0x4f, 0x52, 0x2f, 0xb1, 0x20, 0x53, 0x0f, 0x9b, 0xd3, 0xa5, 0xc4, 0xc0, 0x52,
	0x18, 0xce, 0x52, 0x62, 0x70, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x07, 0x9d, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x5f, 0x58, 0x32, 0x9f, 0x88, 0x01, 0x00, 0x00,
}