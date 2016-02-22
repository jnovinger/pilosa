// Code generated by protoc-gen-go.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Bitmap
	Chunk
	Pair
	Bit
	Profile
	Attr
	AttrMap
	QueryRequest
	QueryResponse
	ImportRequest
	ImportResponse
	Cache
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Bitmap struct {
	Chunks           []*Chunk `protobuf:"bytes,1,rep,name=Chunks" json:"Chunks,omitempty"`
	Attrs            []*Attr  `protobuf:"bytes,2,rep,name=Attrs" json:"Attrs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Bitmap) Reset()                    { *m = Bitmap{} }
func (m *Bitmap) String() string            { return proto.CompactTextString(m) }
func (*Bitmap) ProtoMessage()               {}
func (*Bitmap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bitmap) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *Bitmap) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Chunk struct {
	Key              *uint64  `protobuf:"varint,1,req,name=Key" json:"Key,omitempty"`
	Value            []uint64 `protobuf:"varint,2,rep,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Chunk) GetKey() uint64 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *Chunk) GetValue() []uint64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Pair struct {
	Key              *uint64 `protobuf:"varint,1,req,name=Key" json:"Key,omitempty"`
	Count            *uint64 `protobuf:"varint,2,req,name=Count" json:"Count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Pair) GetKey() uint64 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *Pair) GetCount() uint64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type Bit struct {
	BitmapID         *uint64 `protobuf:"varint,1,req,name=BitmapID" json:"BitmapID,omitempty"`
	ProfileID        *uint64 `protobuf:"varint,2,req,name=ProfileID" json:"ProfileID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Bit) Reset()                    { *m = Bit{} }
func (m *Bit) String() string            { return proto.CompactTextString(m) }
func (*Bit) ProtoMessage()               {}
func (*Bit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Bit) GetBitmapID() uint64 {
	if m != nil && m.BitmapID != nil {
		return *m.BitmapID
	}
	return 0
}

func (m *Bit) GetProfileID() uint64 {
	if m != nil && m.ProfileID != nil {
		return *m.ProfileID
	}
	return 0
}

type Profile struct {
	ID               *uint64 `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	Attrs            []*Attr `protobuf:"bytes,2,rep,name=Attrs" json:"Attrs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Profile) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Profile) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Attr struct {
	Key              *string `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	StringValue      *string `protobuf:"bytes,2,opt,name=StringValue" json:"StringValue,omitempty"`
	UintValue        *uint64 `protobuf:"varint,3,opt,name=UintValue" json:"UintValue,omitempty"`
	BoolValue        *bool   `protobuf:"varint,4,opt,name=BoolValue" json:"BoolValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Attr) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Attr) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Attr) GetUintValue() uint64 {
	if m != nil && m.UintValue != nil {
		return *m.UintValue
	}
	return 0
}

func (m *Attr) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

type AttrMap struct {
	Attrs            []*Attr `protobuf:"bytes,1,rep,name=Attrs" json:"Attrs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AttrMap) Reset()                    { *m = AttrMap{} }
func (m *AttrMap) String() string            { return proto.CompactTextString(m) }
func (*AttrMap) ProtoMessage()               {}
func (*AttrMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AttrMap) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type QueryRequest struct {
	DB               *string  `protobuf:"bytes,1,req,name=DB" json:"DB,omitempty"`
	Query            *string  `protobuf:"bytes,2,req,name=Query" json:"Query,omitempty"`
	Slices           []uint64 `protobuf:"varint,3,rep,name=Slices" json:"Slices,omitempty"`
	Profiles         *bool    `protobuf:"varint,4,opt,name=Profiles" json:"Profiles,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryRequest) GetDB() string {
	if m != nil && m.DB != nil {
		return *m.DB
	}
	return ""
}

func (m *QueryRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *QueryRequest) GetSlices() []uint64 {
	if m != nil {
		return m.Slices
	}
	return nil
}

func (m *QueryRequest) GetProfiles() bool {
	if m != nil && m.Profiles != nil {
		return *m.Profiles
	}
	return false
}

type QueryResponse struct {
	Err              *string    `protobuf:"bytes,1,opt,name=Err" json:"Err,omitempty"`
	Bitmap           *Bitmap    `protobuf:"bytes,2,opt,name=Bitmap" json:"Bitmap,omitempty"`
	N                *uint64    `protobuf:"varint,3,opt,name=N" json:"N,omitempty"`
	Pairs            []*Pair    `protobuf:"bytes,4,rep,name=Pairs" json:"Pairs,omitempty"`
	Profiles         []*Profile `protobuf:"bytes,5,rep,name=Profiles" json:"Profiles,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryResponse) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func (m *QueryResponse) GetBitmap() *Bitmap {
	if m != nil {
		return m.Bitmap
	}
	return nil
}

func (m *QueryResponse) GetN() uint64 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *QueryResponse) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func (m *QueryResponse) GetProfiles() []*Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type ImportRequest struct {
	DB               *string  `protobuf:"bytes,1,req,name=DB" json:"DB,omitempty"`
	Frame            *string  `protobuf:"bytes,2,req,name=Frame" json:"Frame,omitempty"`
	Slice            *uint64  `protobuf:"varint,3,req,name=Slice" json:"Slice,omitempty"`
	BitmapIDs        []uint64 `protobuf:"varint,4,rep,name=BitmapIDs" json:"BitmapIDs,omitempty"`
	ProfileIDs       []uint64 `protobuf:"varint,5,rep,name=ProfileIDs" json:"ProfileIDs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ImportRequest) Reset()                    { *m = ImportRequest{} }
func (m *ImportRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()               {}
func (*ImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ImportRequest) GetDB() string {
	if m != nil && m.DB != nil {
		return *m.DB
	}
	return ""
}

func (m *ImportRequest) GetFrame() string {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return ""
}

func (m *ImportRequest) GetSlice() uint64 {
	if m != nil && m.Slice != nil {
		return *m.Slice
	}
	return 0
}

func (m *ImportRequest) GetBitmapIDs() []uint64 {
	if m != nil {
		return m.BitmapIDs
	}
	return nil
}

func (m *ImportRequest) GetProfileIDs() []uint64 {
	if m != nil {
		return m.ProfileIDs
	}
	return nil
}

type ImportResponse struct {
	Err              *string `protobuf:"bytes,1,opt,name=Err" json:"Err,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImportResponse) Reset()                    { *m = ImportResponse{} }
func (m *ImportResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportResponse) ProtoMessage()               {}
func (*ImportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ImportResponse) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

type Cache struct {
	BitmapIDs        []uint64 `protobuf:"varint,1,rep,name=BitmapIDs" json:"BitmapIDs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Cache) Reset()                    { *m = Cache{} }
func (m *Cache) String() string            { return proto.CompactTextString(m) }
func (*Cache) ProtoMessage()               {}
func (*Cache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Cache) GetBitmapIDs() []uint64 {
	if m != nil {
		return m.BitmapIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*Bitmap)(nil), "internal.Bitmap")
	proto.RegisterType((*Chunk)(nil), "internal.Chunk")
	proto.RegisterType((*Pair)(nil), "internal.Pair")
	proto.RegisterType((*Bit)(nil), "internal.Bit")
	proto.RegisterType((*Profile)(nil), "internal.Profile")
	proto.RegisterType((*Attr)(nil), "internal.Attr")
	proto.RegisterType((*AttrMap)(nil), "internal.AttrMap")
	proto.RegisterType((*QueryRequest)(nil), "internal.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "internal.QueryResponse")
	proto.RegisterType((*ImportRequest)(nil), "internal.ImportRequest")
	proto.RegisterType((*ImportResponse)(nil), "internal.ImportResponse")
	proto.RegisterType((*Cache)(nil), "internal.Cache")
}

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x8f, 0xda, 0x30,
	0x14, 0x54, 0x88, 0x03, 0xe4, 0xa5, 0xa4, 0xe0, 0x5e, 0x50, 0x25, 0x54, 0x64, 0x2e, 0xa8, 0x07,
	0x0e, 0xa8, 0x7f, 0xa0, 0x40, 0xab, 0x22, 0x54, 0x44, 0x8b, 0xda, 0x73, 0x23, 0xe4, 0x96, 0xa8,
	0x21, 0xce, 0x3a, 0xce, 0x81, 0x1f, 0xb1, 0xff, 0x79, 0x9f, 0x3f, 0x12, 0xd8, 0x5d, 0x56, 0x7b,
	0x8a, 0x3c, 0x1e, 0xbf, 0x99, 0x79, 0x13, 0x88, 0xd3, 0x5c, 0x71, 0x99, 0x27, 0xd9, 0xac, 0x90,
	0x42, 0x09, 0xda, 0xad, 0xcf, 0xec, 0x1b, 0xb4, 0x17, 0xa9, 0x3a, 0x25, 0x05, 0xfd, 0x00, 0xed,
	0xe5, 0xb1, 0xca, 0xff, 0x97, 0x43, 0x6f, 0xec, 0x4f, 0xa3, 0xf9, 0xdb, 0x59, 0xf3, 0xc8, 0xe0,
	0x74, 0x04, 0xc1, 0x67, 0xa5, 0x64, 0x39, 0x6c, 0x99, 0xfb, 0xf8, 0x72, 0xaf, 0x61, 0x36, 0x81,
	0xc0, 0xf2, 0x22, 0xf0, 0x37, 0xfc, 0x8c, 0x53, 0x5a, 0x53, 0x42, 0x7b, 0x10, 0xfc, 0x4e, 0xb2,
	0x8a, 0x9b, 0x47, 0x84, 0x31, 0x20, 0xbb, 0x24, 0x95, 0xcf, 0x38, 0x4b, 0x51, 0xe5, 0x0a, 0x39,
	0x78, 0x64, 0x1f, 0xc1, 0x47, 0x4b, 0xb4, 0x0f, 0x5d, 0xeb, 0x6c, 0xbd, 0x72, 0xbc, 0x01, 0x84,
	0x3b, 0x29, 0xfe, 0xa6, 0x19, 0x47, 0xc8, 0x72, 0x3f, 0x41, 0xc7, 0x41, 0x14, 0xa0, 0xd5, 0x30,
	0x5f, 0xb1, 0xba, 0x05, 0xa2, 0xbf, 0xd7, 0x2e, 0x42, 0xfa, 0x0e, 0xa2, 0xbd, 0x92, 0x69, 0xfe,
	0xaf, 0xf6, 0xeb, 0x21, 0x88, 0x92, 0xbf, 0xf0, 0xad, 0x85, 0x7c, 0x84, 0x8c, 0x8b, 0x85, 0x10,
	0x99, 0x85, 0x08, 0x42, 0x5d, 0x36, 0x85, 0x8e, 0x9e, 0xf7, 0x1d, 0xb7, 0xd8, 0x28, 0x7b, 0x37,
	0x95, 0x37, 0xf0, 0xe6, 0x47, 0xc5, 0xe5, 0xf9, 0x27, 0xbf, 0xab, 0x78, 0xa9, 0xb4, 0xe9, 0xd5,
	0xc2, 0x19, 0xc0, 0x35, 0x98, 0x3b, 0x13, 0x2d, 0xa4, 0x31, 0xb4, 0xf7, 0x59, 0x7a, 0xe0, 0x25,
	0xea, 0xe2, 0xea, 0xf4, 0x3e, 0x5c, 0xd4, 0xd2, 0xc9, 0xde, 0x7b, 0xd0, 0x73, 0xd3, 0xca, 0x42,
	0xe4, 0x25, 0xd7, 0x81, 0xbe, 0x48, 0x89, 0xf3, 0xb4, 0xf7, 0x71, 0x5d, 0xad, 0xc9, 0x12, 0xcd,
	0xfb, 0x17, 0x2f, 0xae, 0xf2, 0x10, 0xbc, 0xad, 0x4b, 0x85, 0xbe, 0x75, 0x31, 0x7a, 0xf4, 0x13,
	0xdf, 0xa6, 0xaf, 0xc9, 0x95, 0x78, 0x60, 0x18, 0x83, 0x2b, 0x86, 0xbd, 0x61, 0x7f, 0xa0, 0xb7,
	0x3e, 0x15, 0x42, 0xaa, 0x17, 0xd2, 0x7d, 0x95, 0xc9, 0x89, 0xbb, 0x74, 0x78, 0x34, 0xe9, 0x50,
	0xde, 0x55, 0x5b, 0x97, 0x6d, 0x2d, 0x10, 0x8a, 0xaf, 0x9b, 0xb6, 0xad, 0x28, 0x61, 0x23, 0x88,
	0x6b, 0x85, 0x1b, 0x89, 0xd9, 0x7b, 0xfc, 0x91, 0x92, 0xc3, 0x91, 0x3f, 0x1e, 0xa7, 0x9b, 0x20,
	0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xbb, 0x74, 0xfe, 0x03, 0x03, 0x00, 0x00,
}