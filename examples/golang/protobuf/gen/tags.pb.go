// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tags.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type Outside struct {
	Inside *Inside `protobuf:"bytes,1,opt,name=Inside,proto3" json:"Inside,omitempty" validate:"required"`
	Field2 string  `protobuf:"bytes,2,opt,name=Field2,proto3" json:"Field2,omitempty" validate:"min=2"`
	// Types that are valid to be assigned to Filed:
	//	*Outside_Field3
	Filed                isOutside_Filed `protobuf_oneof:"filed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Outside) Reset()         { *m = Outside{} }
func (m *Outside) String() string { return proto.CompactTextString(m) }
func (*Outside) ProtoMessage()    {}
func (*Outside) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d9cbcae1e528f6, []int{0}
}
func (m *Outside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outside.Unmarshal(m, b)
}
func (m *Outside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outside.Marshal(b, m, deterministic)
}
func (m *Outside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outside.Merge(m, src)
}
func (m *Outside) XXX_Size() int {
	return xxx_messageInfo_Outside.Size(m)
}
func (m *Outside) XXX_DiscardUnknown() {
	xxx_messageInfo_Outside.DiscardUnknown(m)
}

var xxx_messageInfo_Outside proto.InternalMessageInfo

type isOutside_Filed interface {
	isOutside_Filed()
}

type Outside_Field3 struct {
	Field3 string `protobuf:"bytes,3,opt,name=Field3,proto3,oneof" json:"Field3,omitempty" valiate:"max=8"`
}

func (*Outside_Field3) isOutside_Filed() {}

func (m *Outside) GetFiled() isOutside_Filed {
	if m != nil {
		return m.Filed
	}
	return nil
}

func (m *Outside) GetInside() *Inside {
	if m != nil {
		return m.Inside
	}
	return nil
}

func (m *Outside) GetField2() string {
	if m != nil {
		return m.Field2
	}
	return ""
}

func (m *Outside) GetField3() string {
	if x, ok := m.GetFiled().(*Outside_Field3); ok {
		return x.Field3
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Outside) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Outside_Field3)(nil),
	}
}

type Inside struct {
	Field1               string   `protobuf:"bytes,1,opt,name=Field1,proto3" json:"Field1,omitempty" validate:"max=8"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inside) Reset()         { *m = Inside{} }
func (m *Inside) String() string { return proto.CompactTextString(m) }
func (*Inside) ProtoMessage()    {}
func (*Inside) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d9cbcae1e528f6, []int{1}
}
func (m *Inside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inside.Unmarshal(m, b)
}
func (m *Inside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inside.Marshal(b, m, deterministic)
}
func (m *Inside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inside.Merge(m, src)
}
func (m *Inside) XXX_Size() int {
	return xxx_messageInfo_Inside.Size(m)
}
func (m *Inside) XXX_DiscardUnknown() {
	xxx_messageInfo_Inside.DiscardUnknown(m)
}

var xxx_messageInfo_Inside proto.InternalMessageInfo

func (m *Inside) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func init() {
	proto.RegisterType((*Outside)(nil), "proto.Outside")
	proto.RegisterType((*Inside)(nil), "proto.Inside")
}

func init() { proto.RegisterFile("tags.proto", fileDescriptor_e7d9cbcae1e528f6) }

var fileDescriptor_e7d9cbcae1e528f6 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4c, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe1, 0xa4, 0xd2,
	0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba, 0x94, 0xd6, 0x33, 0x72, 0xb1, 0xfb, 0x97, 0x96,
	0x14, 0x67, 0xa6, 0xa4, 0x0a, 0xd9, 0x71, 0xb1, 0x79, 0xe6, 0x81, 0x58, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0xbc, 0x10, 0x35, 0x7a, 0x10, 0x41, 0x27, 0xf1, 0x4f, 0xf7, 0xe4, 0x85, 0xcb,
	0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xad, 0x94, 0x8a, 0x52, 0x0b, 0x4b, 0x33, 0x8b, 0x52,
	0x53, 0x94, 0x82, 0xa0, 0xba, 0x84, 0x74, 0xb8, 0xd8, 0xdc, 0x32, 0x53, 0x73, 0x52, 0x8c, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x9d, 0x44, 0x3e, 0xdd, 0x93, 0x17, 0x40, 0x68, 0xc8, 0xcd, 0xcc,
	0xb3, 0x35, 0x52, 0x0a, 0x82, 0xaa, 0x11, 0xd2, 0x85, 0xaa, 0x36, 0x96, 0x60, 0x06, 0xab, 0x16,
	0xfe, 0x74, 0x4f, 0x9e, 0x1f, 0xa4, 0x1a, 0xa2, 0x38, 0xb1, 0xc2, 0xd6, 0x42, 0xc9, 0x83, 0x01,
	0xaa, 0xdc, 0xd8, 0x89, 0x9d, 0x8b, 0x35, 0x2d, 0x33, 0x27, 0x35, 0x45, 0xc9, 0x8c, 0x0b, 0xdd,
	0x3e, 0x43, 0xb0, 0x7b, 0x31, 0xed, 0x03, 0x1b, 0x01, 0x35, 0xc0, 0xd0, 0x89, 0xe5, 0xc7, 0x43,
	0x39, 0xc6, 0x24, 0x36, 0xb0, 0x97, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0x3c, 0x77,
	0xb0, 0x3a, 0x01, 0x00, 0x00,
}

func NewPopulatedOutside(r randyTags, easy bool) *Outside {
	this := &Outside{}
	if r.Intn(5) != 0 {
		this.Inside = NewPopulatedInside(r, easy)
	}
	this.Field2 = string(randStringTags(r))
	oneofNumber_Filed := []int32{3}[r.Intn(1)]
	switch oneofNumber_Filed {
	case 3:
		this.Filed = NewPopulatedOutside_Field3(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 4)
	}
	return this
}

func NewPopulatedOutside_Field3(r randyTags, easy bool) *Outside_Field3 {
	this := &Outside_Field3{}
	this.Field3 = string(randStringTags(r))
	return this
}
func NewPopulatedInside(r randyTags, easy bool) *Inside {
	this := &Inside{}
	this.Field1 = string(randStringTags(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 2)
	}
	return this
}

type randyTags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTags(r randyTags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTags(r randyTags) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTags(r)
	}
	return string(tmps)
}
func randUnrecognizedTags(r randyTags, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTags(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTags(dAtA []byte, r randyTags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTags(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTags(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTags(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
