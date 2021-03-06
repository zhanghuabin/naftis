// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/authn/v2alpha1/config.proto

/*
	Package v2alpha1 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/authn/v2alpha1/config.proto
		envoy/config/filter/http/jwt_auth/v2alpha1/config.proto

	It has these top-level messages:
		FilterConfig
		HttpUri
		DataSource
		JwtRule
		RemoteJwks
		JwtHeader
		JwtAuthentication
*/
package v2alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import istio_authentication_v1alpha1 "istio.io/api/authentication/v1alpha1"

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

// FilterConfig is the config for Istio-specific filter that is used to enforce
// authentication policy on Envoy.
type FilterConfig struct {
	// Policy is the original copy of the policy.
	Policy *istio_authentication_v1alpha1.Policy `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// Map from issuer to location of the payload that is emitted by Jwt filter.
	// This information is added by pilot when construct and add Jwt and
	// authN filters.
	JwtOutputPayloadLocations map[string]string `protobuf:"bytes,2,rep,name=jwt_output_payload_locations,json=jwtOutputPayloadLocations" json:"jwt_output_payload_locations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *FilterConfig) Reset()                    { *m = FilterConfig{} }
func (m *FilterConfig) String() string            { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()               {}
func (*FilterConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *FilterConfig) GetPolicy() *istio_authentication_v1alpha1.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *FilterConfig) GetJwtOutputPayloadLocations() map[string]string {
	if m != nil {
		return m.JwtOutputPayloadLocations
	}
	return nil
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig")
}
func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Policy.Size()))
		n1, err := m.Policy.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.JwtOutputPayloadLocations) > 0 {
		for k, _ := range m.JwtOutputPayloadLocations {
			dAtA[i] = 0x12
			i++
			v := m.JwtOutputPayloadLocations[k]
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FilterConfig) Size() (n int) {
	var l int
	_ = l
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.JwtOutputPayloadLocations) > 0 {
		for k, v := range m.JwtOutputPayloadLocations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &istio_authentication_v1alpha1.Policy{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwtOutputPayloadLocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.JwtOutputPayloadLocations == nil {
				m.JwtOutputPayloadLocations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.JwtOutputPayloadLocations[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/authn/v2alpha1/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0x87, 0x03, 0x33, 0x0f, 0x52, 0x3c, 0xcc, 0x21, 0x65, 0x88, 0xc2, 0x2e, 0xbe,
	0xb0, 0xb9, 0x83, 0x08, 0x5e, 0x26, 0x7a, 0x90, 0x81, 0xa3, 0xc7, 0x5d, 0x4a, 0xac, 0x9d, 0xcd,
	0x0c, 0x49, 0xe8, 0x5e, 0x3b, 0xfa, 0x59, 0xfc, 0x24, 0x7e, 0x03, 0x8f, 0x7e, 0x04, 0xe9, 0x27,
	0x91, 0x26, 0x2d, 0xe8, 0x41, 0xd1, 0x5b, 0x5e, 0xde, 0xfb, 0xfd, 0x92, 0x3f, 0x8f, 0x4e, 0x13,
	0x55, 0xe8, 0x92, 0xc5, 0x5a, 0xad, 0xc4, 0x13, 0x5b, 0x09, 0x89, 0x49, 0xc6, 0x52, 0x44, 0xc3,
	0x78, 0x8e, 0xa9, 0x62, 0xc5, 0x84, 0x4b, 0x93, 0xf2, 0x71, 0x33, 0x01, 0x26, 0xd3, 0xa8, 0xfd,
	0x33, 0xb1, 0x41, 0xa1, 0xc1, 0xb2, 0xd0, 0x74, 0x1c, 0x0b, 0x35, 0x0b, 0x96, 0x85, 0x96, 0x1d,
	0x9c, 0xd4, 0x75, 0xa2, 0x50, 0xc4, 0x1c, 0x85, 0x56, 0xac, 0x18, 0x37, 0x52, 0xa3, 0xa5, 0x88,
	0x4b, 0x27, 0x3d, 0x7e, 0xf5, 0xe8, 0xde, 0xad, 0x95, 0x5c, 0x5b, 0xa3, 0x7f, 0x45, 0xbb, 0x6e,
	0xa0, 0x4f, 0x86, 0x64, 0xd4, 0x9b, 0x9c, 0x82, 0x7b, 0xf6, 0xbb, 0x0d, 0x5a, 0x1b, 0x2c, 0xec,
	0x70, 0xd8, 0x40, 0xfe, 0x0b, 0xa1, 0x47, 0xeb, 0x2d, 0x46, 0x3a, 0x47, 0x93, 0x63, 0x64, 0x78,
	0x29, 0x35, 0x7f, 0x8c, 0xa4, 0x76, 0xdc, 0xa6, 0xef, 0x0d, 0x3b, 0xa3, 0xde, 0x64, 0x09, 0xff,
	0x0a, 0x03, 0x5f, 0xbf, 0x08, 0x77, 0x5b, 0xbc, 0xb7, 0xfa, 0x85, 0xb3, 0xcf, 0x5b, 0xf9, 0x8d,
	0xc2, 0xac, 0x0c, 0x0f, 0xd7, 0x3f, 0xf5, 0x07, 0x73, 0x1a, 0xfc, 0x0e, 0xfb, 0xfb, 0xb4, 0xf3,
	0x9c, 0xb8, 0xec, 0xbb, 0x61, 0x7d, 0xf4, 0x0f, 0xe8, 0x4e, 0xc1, 0x65, 0x9e, 0xf4, 0x3d, 0x7b,
	0xe7, 0x8a, 0x4b, 0xef, 0x82, 0xcc, 0x66, 0x6f, 0x55, 0x40, 0xde, 0xab, 0x80, 0x7c, 0x54, 0x01,
	0x59, 0x4e, 0x5d, 0x22, 0xa1, 0x19, 0x37, 0x82, 0xfd, 0x71, 0xc3, 0x0f, 0x5d, 0xbb, 0x86, 0xf3,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x7c, 0xfe, 0x70, 0x13, 0x02, 0x00, 0x00,
}
