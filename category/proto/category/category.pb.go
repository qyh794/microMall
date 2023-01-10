// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/category/category.proto

/*
Package go_micro_service_category is a generated protocol buffer package.

It is generated from these files:
	proto/category/category.proto

It has these top-level messages:
	CategoryRequest
	CreateCategoryResponse
	UpdateCategoryResponse
	DeleteCategoryRequest
	DeleteCategoryResponse
	FindByNameRequest
	CategoryResponse
	FindByIDRequest
	FindByLevelRequest
	FindAllResponse
	FindByParentRequest
	FindAllRequest
*/
package go_micro_service_category

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

type CategoryRequest struct {
	CategoryName        string `protobuf:"bytes,1,opt,name=categoryName" json:"categoryName,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=categoryLevel" json:"categoryLevel,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=categoryParent" json:"categoryParent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=categoryImage" json:"categoryImage,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=categoryDescription" json:"categoryDescription,omitempty"`
}

func (m *CategoryRequest) Reset()                    { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()               {}
func (*CategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CategoryRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryRequest) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryRequest) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryRequest) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *CategoryRequest) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type CreateCategoryResponse struct {
	Message    string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	CategoryId int64  `protobuf:"varint,2,opt,name=categoryId" json:"categoryId,omitempty"`
}

func (m *CreateCategoryResponse) Reset()                    { *m = CreateCategoryResponse{} }
func (m *CreateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCategoryResponse) ProtoMessage()               {}
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateCategoryResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type UpdateCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateCategoryResponse) Reset()                    { *m = UpdateCategoryResponse{} }
func (m *UpdateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateCategoryResponse) ProtoMessage()               {}
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteCategoryRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=categoryId" json:"categoryId,omitempty"`
}

func (m *DeleteCategoryRequest) Reset()                    { *m = DeleteCategoryRequest{} }
func (m *DeleteCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryRequest) ProtoMessage()               {}
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteCategoryRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type DeleteCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteCategoryResponse) Reset()                    { *m = DeleteCategoryResponse{} }
func (m *DeleteCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryResponse) ProtoMessage()               {}
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type FindByNameRequest struct {
	CategoryName string `protobuf:"bytes,1,opt,name=categoryName" json:"categoryName,omitempty"`
}

func (m *FindByNameRequest) Reset()                    { *m = FindByNameRequest{} }
func (m *FindByNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByNameRequest) ProtoMessage()               {}
func (*FindByNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FindByNameRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

type CategoryResponse struct {
	Id                  int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CategoryName        string `protobuf:"bytes,2,opt,name=categoryName" json:"categoryName,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,3,opt,name=categoryLevel" json:"categoryLevel,omitempty"`
	CategoryParent      int64  `protobuf:"varint,4,opt,name=categoryParent" json:"categoryParent,omitempty"`
	CategoryImages      string `protobuf:"bytes,5,opt,name=categoryImages" json:"categoryImages,omitempty"`
	CategoryDescription string `protobuf:"bytes,6,opt,name=categoryDescription" json:"categoryDescription,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CategoryResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryResponse) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryResponse) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryResponse) GetCategoryImages() string {
	if m != nil {
		return m.CategoryImages
	}
	return ""
}

func (m *CategoryResponse) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type FindByIDRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=categoryId" json:"categoryId,omitempty"`
}

func (m *FindByIDRequest) Reset()                    { *m = FindByIDRequest{} }
func (m *FindByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByIDRequest) ProtoMessage()               {}
func (*FindByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FindByIDRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type FindByLevelRequest struct {
	Level uint32 `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
}

func (m *FindByLevelRequest) Reset()                    { *m = FindByLevelRequest{} }
func (m *FindByLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByLevelRequest) ProtoMessage()               {}
func (*FindByLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FindByLevelRequest) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type FindAllResponse struct {
	Category []*CategoryResponse `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
}

func (m *FindAllResponse) Reset()                    { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string            { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()               {}
func (*FindAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FindAllResponse) GetCategory() []*CategoryResponse {
	if m != nil {
		return m.Category
	}
	return nil
}

type FindByParentRequest struct {
	ParentId int64 `protobuf:"varint,1,opt,name=parentId" json:"parentId,omitempty"`
}

func (m *FindByParentRequest) Reset()                    { *m = FindByParentRequest{} }
func (m *FindByParentRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByParentRequest) ProtoMessage()               {}
func (*FindByParentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FindByParentRequest) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type FindAllRequest struct {
}

func (m *FindAllRequest) Reset()                    { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()               {}
func (*FindAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*CategoryRequest)(nil), "go.micro.service.category.CategoryRequest")
	proto.RegisterType((*CreateCategoryResponse)(nil), "go.micro.service.category.CreateCategoryResponse")
	proto.RegisterType((*UpdateCategoryResponse)(nil), "go.micro.service.category.UpdateCategoryResponse")
	proto.RegisterType((*DeleteCategoryRequest)(nil), "go.micro.service.category.DeleteCategoryRequest")
	proto.RegisterType((*DeleteCategoryResponse)(nil), "go.micro.service.category.DeleteCategoryResponse")
	proto.RegisterType((*FindByNameRequest)(nil), "go.micro.service.category.FindByNameRequest")
	proto.RegisterType((*CategoryResponse)(nil), "go.micro.service.category.CategoryResponse")
	proto.RegisterType((*FindByIDRequest)(nil), "go.micro.service.category.FindByIDRequest")
	proto.RegisterType((*FindByLevelRequest)(nil), "go.micro.service.category.FindByLevelRequest")
	proto.RegisterType((*FindAllResponse)(nil), "go.micro.service.category.FindAllResponse")
	proto.RegisterType((*FindByParentRequest)(nil), "go.micro.service.category.FindByParentRequest")
	proto.RegisterType((*FindAllRequest)(nil), "go.micro.service.category.FindAllRequest")
}

func init() { proto.RegisterFile("proto/category/category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xcd, 0x7e, 0xd4, 0xa3, 0xdb, 0xad, 0xb3, 0xeb, 0x12, 0x0b, 0x4a, 0x19, 0x44,
	0xea, 0xaa, 0x71, 0x5b, 0x2f, 0xbc, 0xd6, 0x2d, 0x4a, 0x41, 0x44, 0x02, 0xde, 0x78, 0x17, 0xdb,
	0x43, 0x89, 0xe4, 0x6b, 0x33, 0xb1, 0x20, 0x3e, 0x8f, 0x4f, 0xe5, 0x3b, 0xf8, 0x0c, 0xd2, 0x99,
	0xce, 0x34, 0x93, 0xa6, 0x69, 0x22, 0x7b, 0xb7, 0x73, 0x32, 0x73, 0x7e, 0xe7, 0xe3, 0xff, 0xdf,
	0xc2, 0xa3, 0x24, 0x8d, 0xb3, 0xf8, 0xd5, 0xcc, 0xcb, 0x70, 0x11, 0xa7, 0x3f, 0xf5, 0x1f, 0x8e,
	0x88, 0xd3, 0x87, 0x8b, 0xd8, 0x09, 0xfd, 0x59, 0x1a, 0x3b, 0x1c, 0xd3, 0xa5, 0x3f, 0x43, 0x47,
	0x5d, 0x60, 0x7f, 0x08, 0x9c, 0x5e, 0xaf, 0x0f, 0x2e, 0xde, 0xfc, 0x40, 0x9e, 0x51, 0x06, 0xf7,
	0xd4, 0xf7, 0x4f, 0x5e, 0x88, 0x36, 0x19, 0x90, 0xe1, 0x1d, 0xd7, 0x88, 0xd1, 0x27, 0x70, 0xa2,
	0xce, 0x1f, 0x71, 0x89, 0x81, 0xdd, 0x1e, 0x90, 0xe1, 0x89, 0x6b, 0x06, 0xe9, 0x53, 0xe8, 0xaa,
	0xc0, 0x67, 0x2f, 0xc5, 0x28, 0xb3, 0xad, 0x01, 0x19, 0x5a, 0x6e, 0x21, 0x9a, 0xcf, 0x36, 0x0d,
	0xbd, 0x05, 0xda, 0x07, 0x02, 0x69, 0x06, 0xe9, 0x15, 0x9c, 0xa9, 0xc0, 0x04, 0xf9, 0x2c, 0xf5,
	0x93, 0xcc, 0x8f, 0x23, 0xfb, 0x50, 0xdc, 0x2d, 0xfb, 0xc4, 0x5c, 0xb8, 0xb8, 0x4e, 0xd1, 0xcb,
	0x70, 0xd3, 0x22, 0x4f, 0xe2, 0x88, 0x23, 0xb5, 0xe1, 0x38, 0x44, 0xce, 0x57, 0x2c, 0xd9, 0x9e,
	0x3a, 0xd2, 0xc7, 0x00, 0x1a, 0x3b, 0x17, 0x6d, 0x59, 0x6e, 0x2e, 0xc2, 0xc6, 0x70, 0xf1, 0x25,
	0x99, 0x37, 0xca, 0xc9, 0xde, 0xc0, 0x83, 0x09, 0x06, 0x98, 0x7f, 0x23, 0x47, 0x6d, 0xc2, 0x48,
	0x19, 0xac, 0xf8, 0xb0, 0x06, 0xec, 0xfe, 0x7b, 0x3f, 0x9a, 0xbf, 0x13, 0x8b, 0x6a, 0xb0, 0x53,
	0xf6, 0x97, 0x40, 0x6f, 0x8b, 0xd3, 0x85, 0xb6, 0xaf, 0x2a, 0x6b, 0xfb, 0xf3, 0xad, 0x44, 0xed,
	0x3a, 0xe2, 0xb0, 0xea, 0x89, 0xe3, 0xa0, 0x54, 0x1c, 0xb9, 0x7b, 0x42, 0x07, 0x7c, 0xbd, 0xf1,
	0x42, 0x74, 0x97, 0x3c, 0x8e, 0x76, 0xcb, 0x63, 0x04, 0xa7, 0x72, 0x52, 0xd3, 0x49, 0xdd, 0x85,
	0x5c, 0x02, 0x95, 0x4f, 0x44, 0x0f, 0xea, 0xd5, 0x39, 0x1c, 0x06, 0xa2, 0x51, 0x22, 0x1a, 0x95,
	0x07, 0xf6, 0x55, 0xa6, 0x7f, 0x1b, 0x04, 0x7a, 0x9a, 0x1f, 0xa0, 0xa3, 0x92, 0xd9, 0x64, 0x60,
	0x0d, 0xef, 0x8e, 0x9f, 0x3b, 0x3b, 0xcd, 0xe9, 0x14, 0x97, 0xe1, 0xea, 0xc7, 0x6c, 0x04, 0x67,
	0xb2, 0x0e, 0x39, 0x24, 0x55, 0x48, 0x1f, 0x3a, 0x89, 0x08, 0xe8, 0xe2, 0xf5, 0x99, 0xf5, 0xa0,
	0xab, 0xcb, 0x11, 0xb7, 0xc7, 0xbf, 0x8f, 0xa1, 0xa3, 0x18, 0xf4, 0x06, 0xba, 0xa6, 0x57, 0xe8,
	0x65, 0xad, 0xd2, 0x44, 0xaa, 0xfe, 0xa8, 0xea, 0x6e, 0xa9, 0x05, 0x59, 0x6b, 0x85, 0x34, 0xad,
	0x74, 0x6b, 0xc8, 0x72, 0x87, 0xb2, 0x16, 0xfd, 0x05, 0xd4, 0x34, 0xd4, 0x6a, 0xf9, 0xf4, 0xaa,
	0x22, 0x55, 0xa9, 0x71, 0x2b, 0xe1, 0xe5, 0x8e, 0x15, 0xfd, 0x0a, 0xf1, 0x6c, 0xd0, 0xc2, 0x2d,
	0x2f, 0x2a, 0x52, 0x6d, 0x19, 0xb9, 0xdf, 0x44, 0x2f, 0xac, 0x45, 0x43, 0xe8, 0x99, 0xc8, 0xe9,
	0xa4, 0x72, 0xc8, 0x05, 0x3f, 0x34, 0xc5, 0xa5, 0x52, 0x96, 0x1b, 0x9c, 0xb4, 0xfa, 0xcb, 0xbd,
	0xc4, 0xbc, 0x9d, 0xfa, 0xfb, 0x0a, 0xcc, 0x39, 0x8a, 0xb5, 0x68, 0x06, 0xe7, 0x26, 0x73, 0xfd,
	0x7f, 0xc3, 0xd9, 0x0b, 0x35, 0xbc, 0xd3, 0x90, 0xfa, 0x5d, 0x9b, 0x5b, 0x8b, 0xf7, 0x59, 0x9d,
	0x04, 0xff, 0xc1, 0xfa, 0x76, 0x24, 0x7e, 0xc6, 0x5f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xde,
	0x42, 0x61, 0xa8, 0xe7, 0x07, 0x00, 0x00,
}