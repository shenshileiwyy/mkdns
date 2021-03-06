// Code generated by protoc-gen-go.
// source: zone.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	zone.proto

It has these top-level messages:
	Record
	Records
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Record_Type int32

const (
	Record_UNKNOWN Record_Type = 0
	Record_SOA     Record_Type = 1
	Record_NS      Record_Type = 2
	Record_A       Record_Type = 3
	Record_CNAME   Record_Type = 4
	Record_MX      Record_Type = 5
	Record_TXT     Record_Type = 6
	Record_AAAA    Record_Type = 7
	Record_SRV     Record_Type = 8
	Record_PTR     Record_Type = 9
	Record_SPF     Record_Type = 10
	Record_ALIAS   Record_Type = 11
	Record_CAA     Record_Type = 12
)

var Record_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "SOA",
	2:  "NS",
	3:  "A",
	4:  "CNAME",
	5:  "MX",
	6:  "TXT",
	7:  "AAAA",
	8:  "SRV",
	9:  "PTR",
	10: "SPF",
	11: "ALIAS",
	12: "CAA",
}
var Record_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"SOA":     1,
	"NS":      2,
	"A":       3,
	"CNAME":   4,
	"MX":      5,
	"TXT":     6,
	"AAAA":    7,
	"SRV":     8,
	"PTR":     9,
	"SPF":     10,
	"ALIAS":   11,
	"CAA":     12,
}

func (x Record_Type) String() string {
	return proto.EnumName(Record_Type_name, int32(x))
}
func (Record_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Record struct {
	Name  string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ttl   int32           `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	Type  string          `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State int32           `protobuf:"varint,4,opt,name=state" json:"state,omitempty"`
	Value []*Record_Value `protobuf:"bytes,5,rep,name=value" json:"value,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetValue() []*Record_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Record_Value struct {
	Record    []string          `protobuf:"bytes,1,rep,name=record" json:"record,omitempty"`
	View      string            `protobuf:"bytes,2,opt,name=view" json:"view,omitempty"`
	Weight    int32             `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Continent string            `protobuf:"bytes,4,opt,name=continent" json:"continent,omitempty"`
	Country   string            `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	Soa       *Record_Value_Soa `protobuf:"bytes,6,opt,name=soa" json:"soa,omitempty"`
}

func (m *Record_Value) Reset()                    { *m = Record_Value{} }
func (m *Record_Value) String() string            { return proto.CompactTextString(m) }
func (*Record_Value) ProtoMessage()               {}
func (*Record_Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Record_Value) GetSoa() *Record_Value_Soa {
	if m != nil {
		return m.Soa
	}
	return nil
}

type Record_Value_Soa struct {
	Mname   string `protobuf:"bytes,1,opt,name=mname" json:"mname,omitempty"`
	Nname   string `protobuf:"bytes,2,opt,name=nname" json:"nname,omitempty"`
	Serial  uint32 `protobuf:"varint,3,opt,name=serial" json:"serial,omitempty"`
	Refresh uint32 `protobuf:"varint,4,opt,name=refresh" json:"refresh,omitempty"`
	Retry   uint32 `protobuf:"varint,5,opt,name=retry" json:"retry,omitempty"`
	Expire  uint32 `protobuf:"varint,6,opt,name=expire" json:"expire,omitempty"`
	Minttl  uint32 `protobuf:"varint,7,opt,name=minttl" json:"minttl,omitempty"`
}

func (m *Record_Value_Soa) Reset()                    { *m = Record_Value_Soa{} }
func (m *Record_Value_Soa) String() string            { return proto.CompactTextString(m) }
func (*Record_Value_Soa) ProtoMessage()               {}
func (*Record_Value_Soa) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0, 0} }

type Records struct {
	Domain  string    `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Records []*Record `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
}

func (m *Records) Reset()                    { *m = Records{} }
func (m *Records) String() string            { return proto.CompactTextString(m) }
func (*Records) ProtoMessage()               {}
func (*Records) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Records) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "types.Record")
	proto.RegisterType((*Record_Value)(nil), "types.Record.Value")
	proto.RegisterType((*Record_Value_Soa)(nil), "types.Record.Value.Soa")
	proto.RegisterType((*Records)(nil), "types.Records")
	proto.RegisterEnum("types.Record_Type", Record_Type_name, Record_Type_value)
}

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xc9, 0x26, 0x4e, 0x9a, 0x59, 0x22, 0x59, 0x06, 0x41, 0x54, 0x71, 0xa8, 0xf6, 0x42,
	0xb9, 0xe4, 0x50, 0x9e, 0x20, 0xaa, 0x40, 0xe2, 0x4f, 0xd3, 0xca, 0x1b, 0x4a, 0xaf, 0x61, 0x6b,
	0x68, 0xa4, 0x8d, 0xbd, 0x72, 0xdc, 0x96, 0xe5, 0xc4, 0xeb, 0xf0, 0x4a, 0x5c, 0x79, 0x11, 0x3c,
	0xe3, 0x44, 0x80, 0x44, 0x4e, 0xf3, 0x7d, 0xfe, 0x32, 0xf3, 0xcb, 0x38, 0x00, 0xdf, 0x8c, 0x56,
	0xd5, 0xce, 0x1a, 0x67, 0x04, 0x73, 0xfb, 0x9d, 0x1a, 0x57, 0xbf, 0x12, 0x48, 0xa5, 0xda, 0x18,
	0x7b, 0x2d, 0x04, 0x24, 0xba, 0x1b, 0x54, 0x19, 0x1d, 0x45, 0xc7, 0xb9, 0xa4, 0x5a, 0x70, 0x88,
	0x9d, 0xdb, 0x96, 0x0b, 0x6f, 0x31, 0x89, 0x25, 0xa6, 0xf0, 0xcd, 0x32, 0x0e, 0x29, 0xac, 0xc5,
	0x63, 0x60, 0xa3, 0xeb, 0x9c, 0x2a, 0x13, 0xca, 0x05, 0x21, 0x5e, 0x00, 0xbb, 0xeb, 0xb6, 0xb7,
	0xaa, 0x64, 0x47, 0xf1, 0xf1, 0xf2, 0xe4, 0x51, 0x45, 0x13, 0xab, 0x30, 0xad, 0xba, 0xc4, 0x23,
	0x19, 0x12, 0x87, 0x3f, 0x17, 0xc0, 0xc8, 0x10, 0x4f, 0x20, 0xb5, 0x14, 0xf0, 0x18, 0xb1, 0x1f,
	0x30, 0x29, 0x1c, 0x7b, 0xd7, 0xab, 0x7b, 0x22, 0xf1, 0x63, 0xb1, 0xc6, 0xec, 0xbd, 0xea, 0xbf,
	0xdc, 0x38, 0x82, 0x61, 0x72, 0x52, 0xe2, 0x19, 0xe4, 0x1b, 0xa3, 0x5d, 0xaf, 0x95, 0x76, 0x84,
	0x94, 0xcb, 0x3f, 0x86, 0x28, 0x21, 0xdb, 0x98, 0x5b, 0xed, 0xec, 0xde, 0x83, 0xe1, 0xd9, 0x2c,
	0x3d, 0x70, 0x3c, 0x9a, 0xae, 0x4c, 0xbd, 0xbb, 0x3c, 0x79, 0xfa, 0x1f, 0xdc, 0x6a, 0x6d, 0x3a,
	0x89, 0x99, 0xc3, 0x1f, 0x11, 0xc4, 0x5e, 0xe0, 0x97, 0x0f, 0x7f, 0x2d, 0x2d, 0x08, 0x74, 0x35,
	0xb9, 0x81, 0x36, 0x08, 0xc4, 0x1d, 0x95, 0xed, 0xbb, 0x2d, 0xe1, 0x16, 0x72, 0x52, 0x08, 0x64,
	0xd5, 0x67, 0xab, 0xc6, 0x1b, 0x82, 0x2d, 0xe4, 0x2c, 0xb1, 0x8f, 0x55, 0x33, 0x68, 0x21, 0x83,
	0xc0, 0x3e, 0xea, 0xeb, 0xae, 0xb7, 0x8a, 0x48, 0x7d, 0x9f, 0xa0, 0xd0, 0x1f, 0x7a, 0x8d, 0xd7,
	0x95, 0x05, 0x3f, 0xa8, 0xd5, 0xf7, 0x08, 0x92, 0x16, 0xaf, 0x69, 0x09, 0xd9, 0x87, 0xe6, 0x5d,
	0x73, 0xfe, 0xb1, 0xe1, 0x0f, 0x44, 0xe6, 0x3f, 0xe0, 0xbc, 0xe6, 0x91, 0x48, 0x61, 0xd1, 0xac,
	0xf9, 0x42, 0x30, 0x88, 0x6a, 0x1e, 0x8b, 0x1c, 0xd8, 0x69, 0x53, 0x9f, 0xbd, 0xe2, 0x09, 0x9e,
	0x9c, 0x5d, 0x71, 0x86, 0xd1, 0xf6, 0xaa, 0xe5, 0xa9, 0x38, 0x80, 0xa4, 0xf6, 0x0f, 0xcf, 0xe8,
	0x6d, 0x79, 0xc9, 0x0f, 0xb0, 0xb8, 0x68, 0x25, 0xcf, 0xc9, 0xb9, 0x78, 0xcd, 0x01, 0x1b, 0xd4,
	0xef, 0xdf, 0xd4, 0x6b, 0xbe, 0x44, 0xef, 0xd4, 0xc7, 0x1f, 0xae, 0xde, 0x42, 0x16, 0xf6, 0x38,
	0x22, 0xe5, 0xb5, 0x19, 0xba, 0x5e, 0x4f, 0x2b, 0x9b, 0x94, 0x78, 0x8e, 0x5b, 0xa0, 0x88, 0xdf,
	0x1a, 0xfe, 0x2f, 0xc5, 0x3f, 0x17, 0x20, 0xe7, 0xd3, 0x4f, 0x29, 0xfd, 0xbf, 0x2f, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x56, 0x82, 0x72, 0x9f, 0xcd, 0x02, 0x00, 0x00,
}
