// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product/product.proto

/*
Package go_micro_service_product is a generated protocol buffer package.

It is generated from these files:
	proto/product/product.proto

It has these top-level messages:
	ProductInfo
	ProductImage
	ProductSize
	ProductSeo
	ResponseProduct
	RequestID
	Response
	RequestAll
	AllProduct
*/
package go_micro_service_product

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

type ProductInfo struct {
	Id                 int64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ProductName        string          `protobuf:"bytes,2,opt,name=product_name,json=productName" json:"product_name,omitempty"`
	ProductSku         string          `protobuf:"bytes,3,opt,name=product_sku,json=productSku" json:"product_sku,omitempty"`
	ProductPrice       float64         `protobuf:"fixed64,4,opt,name=product_price,json=productPrice" json:"product_price,omitempty"`
	ProductDescription string          `protobuf:"bytes,5,opt,name=product_description,json=productDescription" json:"product_description,omitempty"`
	ProductCategoryId  int64           `protobuf:"varint,6,opt,name=product_category_id,json=productCategoryId" json:"product_category_id,omitempty"`
	ProductImage       []*ProductImage `protobuf:"bytes,7,rep,name=product_image,json=productImage" json:"product_image,omitempty"`
	ProductSize        []*ProductSize  `protobuf:"bytes,8,rep,name=product_size,json=productSize" json:"product_size,omitempty"`
	ProductSeo         *ProductSeo     `protobuf:"bytes,9,opt,name=product_seo,json=productSeo" json:"product_seo,omitempty"`
}

func (m *ProductInfo) Reset()                    { *m = ProductInfo{} }
func (m *ProductInfo) String() string            { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()               {}
func (*ProductInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductInfo) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *ProductInfo) GetProductPrice() float64 {
	if m != nil {
		return m.ProductPrice
	}
	return 0
}

func (m *ProductInfo) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *ProductInfo) GetProductCategoryId() int64 {
	if m != nil {
		return m.ProductCategoryId
	}
	return 0
}

func (m *ProductInfo) GetProductImage() []*ProductImage {
	if m != nil {
		return m.ProductImage
	}
	return nil
}

func (m *ProductInfo) GetProductSize() []*ProductSize {
	if m != nil {
		return m.ProductSize
	}
	return nil
}

func (m *ProductInfo) GetProductSeo() *ProductSeo {
	if m != nil {
		return m.ProductSeo
	}
	return nil
}

type ProductImage struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	ImageCode string `protobuf:"bytes,3,opt,name=image_code,json=imageCode" json:"image_code,omitempty"`
	ImageUrl  string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
}

func (m *ProductImage) Reset()                    { *m = ProductImage{} }
func (m *ProductImage) String() string            { return proto.CompactTextString(m) }
func (*ProductImage) ProtoMessage()               {}
func (*ProductImage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProductImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductImage) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ProductImage) GetImageCode() string {
	if m != nil {
		return m.ImageCode
	}
	return ""
}

func (m *ProductImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type ProductSize struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SizeName string `protobuf:"bytes,2,opt,name=size_name,json=sizeName" json:"size_name,omitempty"`
	SizeCode string `protobuf:"bytes,3,opt,name=size_code,json=sizeCode" json:"size_code,omitempty"`
}

func (m *ProductSize) Reset()                    { *m = ProductSize{} }
func (m *ProductSize) String() string            { return proto.CompactTextString(m) }
func (*ProductSize) ProtoMessage()               {}
func (*ProductSize) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProductSize) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSize) GetSizeName() string {
	if m != nil {
		return m.SizeName
	}
	return ""
}

func (m *ProductSize) GetSizeCode() string {
	if m != nil {
		return m.SizeCode
	}
	return ""
}

type ProductSeo struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SeoTitle       string `protobuf:"bytes,2,opt,name=seo_title,json=seoTitle" json:"seo_title,omitempty"`
	SeoKeywords    string `protobuf:"bytes,3,opt,name=seo_keywords,json=seoKeywords" json:"seo_keywords,omitempty"`
	SeoDescription string `protobuf:"bytes,4,opt,name=seo_description,json=seoDescription" json:"seo_description,omitempty"`
	SeoCode        string `protobuf:"bytes,6,opt,name=seo_code,json=seoCode" json:"seo_code,omitempty"`
}

func (m *ProductSeo) Reset()                    { *m = ProductSeo{} }
func (m *ProductSeo) String() string            { return proto.CompactTextString(m) }
func (*ProductSeo) ProtoMessage()               {}
func (*ProductSeo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProductSeo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSeo) GetSeoTitle() string {
	if m != nil {
		return m.SeoTitle
	}
	return ""
}

func (m *ProductSeo) GetSeoKeywords() string {
	if m != nil {
		return m.SeoKeywords
	}
	return ""
}

func (m *ProductSeo) GetSeoDescription() string {
	if m != nil {
		return m.SeoDescription
	}
	return ""
}

func (m *ProductSeo) GetSeoCode() string {
	if m != nil {
		return m.SeoCode
	}
	return ""
}

type ResponseProduct struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *ResponseProduct) Reset()                    { *m = ResponseProduct{} }
func (m *ResponseProduct) String() string            { return proto.CompactTextString(m) }
func (*ResponseProduct) ProtoMessage()               {}
func (*ResponseProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseProduct) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type RequestID struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *RequestID) Reset()                    { *m = RequestID{} }
func (m *RequestID) String() string            { return proto.CompactTextString(m) }
func (*RequestID) ProtoMessage()               {}
func (*RequestID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestID) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestAll struct {
}

func (m *RequestAll) Reset()                    { *m = RequestAll{} }
func (m *RequestAll) String() string            { return proto.CompactTextString(m) }
func (*RequestAll) ProtoMessage()               {}
func (*RequestAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AllProduct struct {
	ProductInfo []*ProductInfo `protobuf:"bytes,1,rep,name=product_info,json=productInfo" json:"product_info,omitempty"`
}

func (m *AllProduct) Reset()                    { *m = AllProduct{} }
func (m *AllProduct) String() string            { return proto.CompactTextString(m) }
func (*AllProduct) ProtoMessage()               {}
func (*AllProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllProduct) GetProductInfo() []*ProductInfo {
	if m != nil {
		return m.ProductInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "go.micro.service.product.ProductInfo")
	proto.RegisterType((*ProductImage)(nil), "go.micro.service.product.ProductImage")
	proto.RegisterType((*ProductSize)(nil), "go.micro.service.product.ProductSize")
	proto.RegisterType((*ProductSeo)(nil), "go.micro.service.product.ProductSeo")
	proto.RegisterType((*ResponseProduct)(nil), "go.micro.service.product.ResponseProduct")
	proto.RegisterType((*RequestID)(nil), "go.micro.service.product.RequestID")
	proto.RegisterType((*Response)(nil), "go.micro.service.product.Response")
	proto.RegisterType((*RequestAll)(nil), "go.micro.service.product.RequestAll")
	proto.RegisterType((*AllProduct)(nil), "go.micro.service.product.AllProduct")
}

func init() { proto.RegisterFile("proto/product/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x6b, 0x68, 0xe3, 0x49, 0x7f, 0xe8, 0xf2, 0x62, 0x1a, 0x10, 0x61, 0x5b, 0x20, 0xf0,
	0xe0, 0xa2, 0x70, 0x82, 0xd0, 0x80, 0x88, 0x2a, 0xa1, 0xca, 0xa5, 0xf0, 0x82, 0x70, 0x83, 0x77,
	0x1a, 0xad, 0xe2, 0x78, 0x8d, 0xd7, 0x01, 0xa5, 0xa7, 0xe1, 0x32, 0xdc, 0x86, 0x43, 0xa0, 0xdd,
	0xec, 0xda, 0x56, 0xa1, 0x4d, 0x90, 0x78, 0x8a, 0x67, 0xe6, 0xf3, 0x37, 0xdf, 0x7c, 0x33, 0x8a,
	0xa1, 0x9d, 0xe5, 0xa2, 0x10, 0x87, 0x59, 0x2e, 0xd8, 0x2c, 0x2e, 0xec, 0x6f, 0xa0, 0xb3, 0xc4,
	0x1f, 0x8b, 0x60, 0xca, 0xe3, 0x5c, 0x04, 0x12, 0xf3, 0x6f, 0x3c, 0xc6, 0xc0, 0xd4, 0xe9, 0x4f,
	0x17, 0x5a, 0x27, 0x8b, 0xe7, 0x61, 0x7a, 0x21, 0xc8, 0x36, 0xac, 0x71, 0xe6, 0x3b, 0x1d, 0xa7,
	0xeb, 0x86, 0x6b, 0x9c, 0x91, 0x47, 0xb0, 0x69, 0xa0, 0x51, 0x3a, 0x9a, 0xa2, 0xbf, 0xd6, 0x71,
	0xba, 0x5e, 0xd8, 0x32, 0xb9, 0x77, 0xa3, 0x29, 0x92, 0x87, 0x60, 0xc3, 0x48, 0x4e, 0x66, 0xbe,
	0xab, 0x11, 0x60, 0x52, 0xa7, 0x93, 0x19, 0xd9, 0x87, 0x2d, 0x0b, 0xc8, 0x72, 0x1e, 0xa3, 0x7f,
	0xab, 0xe3, 0x74, 0x9d, 0xd0, 0x12, 0x9f, 0xa8, 0x1c, 0x39, 0x84, 0xbb, 0x16, 0xc4, 0x50, 0xc6,
	0x39, 0xcf, 0x0a, 0x2e, 0x52, 0xff, 0xb6, 0x66, 0x23, 0xa6, 0x34, 0xa8, 0x2a, 0x24, 0xa8, 0x5e,
	0x88, 0x47, 0x05, 0x8e, 0x45, 0x3e, 0x8f, 0x38, 0xf3, 0xd7, 0xb5, 0xf4, 0x5d, 0x53, 0x3a, 0x32,
	0x95, 0x21, 0x23, 0xc7, 0x95, 0x0a, 0x3e, 0x1d, 0x8d, 0xd1, 0xdf, 0xe8, 0xb8, 0xdd, 0x56, 0xef,
	0x49, 0x70, 0x9d, 0x37, 0x81, 0xf5, 0x45, 0xa1, 0x4b, 0xb5, 0x3a, 0x22, 0x6f, 0x2b, 0x5b, 0x24,
	0xbf, 0x44, 0xbf, 0xa9, 0xb9, 0x1e, 0x2f, 0xe5, 0x3a, 0xe5, 0x97, 0x58, 0xba, 0xa7, 0x02, 0xf2,
	0xba, 0xe6, 0x1e, 0x0a, 0xdf, 0xeb, 0x38, 0xdd, 0x56, 0xef, 0x60, 0x39, 0x11, 0x8a, 0xca, 0x63,
	0x14, 0x74, 0x0e, 0x9b, 0x75, 0xb9, 0x7f, 0xec, 0xf1, 0x01, 0x80, 0x9e, 0xba, 0xbe, 0x45, 0x4f,
	0x67, 0xf4, 0x0e, 0xcb, 0x72, 0x2c, 0x18, 0x9a, 0x15, 0x2e, 0xca, 0x47, 0x82, 0x21, 0x69, 0xc3,
	0x22, 0x88, 0x66, 0x79, 0xa2, 0xb7, 0xe7, 0x85, 0x4d, 0x9d, 0x38, 0xcb, 0x13, 0xfa, 0xb1, 0xbc,
	0x20, 0x3d, 0xd0, 0xd5, 0xce, 0x6d, 0xf0, 0x94, 0x45, 0xf5, 0xc6, 0x4d, 0x95, 0xd0, 0x7d, 0x6d,
	0xb1, 0xd6, 0x56, 0x17, 0x55, 0x57, 0xfa, 0xc3, 0x01, 0xa8, 0xc6, 0xfd, 0x2b, 0x31, 0x8a, 0xa8,
	0xe0, 0x45, 0x52, 0x11, 0xa3, 0x78, 0xaf, 0x62, 0x75, 0xb7, 0xaa, 0x38, 0xc1, 0xf9, 0x77, 0x91,
	0x33, 0x69, 0xb8, 0x5b, 0x12, 0xc5, 0xb1, 0x49, 0x91, 0xa7, 0xb0, 0xa3, 0x20, 0xf5, 0x6b, 0x5b,
	0x8c, 0xb6, 0x2d, 0x51, 0xd4, 0x2f, 0xed, 0x1e, 0x28, 0xde, 0x85, 0xc6, 0x75, 0x8d, 0xd8, 0x90,
	0x28, 0xb4, 0xc4, 0x17, 0xb0, 0x13, 0xa2, 0xcc, 0x44, 0x2a, 0xd1, 0x28, 0x55, 0x56, 0x96, 0x77,
	0x66, 0xe5, 0x7a, 0xf6, 0x78, 0x18, 0x7d, 0x0e, 0x5e, 0x88, 0x5f, 0x67, 0x28, 0x8b, 0xe1, 0x60,
	0x19, 0xf6, 0x3e, 0x34, 0x2d, 0x3b, 0xb9, 0x03, 0xee, 0x54, 0x8e, 0x35, 0xc6, 0x0b, 0xd5, 0x23,
	0xdd, 0x04, 0x30, 0x4c, 0xfd, 0x24, 0xa1, 0x1f, 0x00, 0xfa, 0x49, 0x62, 0x45, 0xd4, 0xee, 0x93,
	0xa7, 0x17, 0xc2, 0x77, 0x56, 0xbc, 0x4f, 0xf5, 0x1f, 0x50, 0xde, 0xa7, 0x0a, 0x7a, 0xbf, 0x5c,
	0xd8, 0xb0, 0xac, 0xe7, 0x00, 0x7d, 0xc6, 0x6c, 0xb4, 0x1a, 0xdb, 0xde, 0xb3, 0xeb, 0x61, 0x57,
	0xac, 0xa3, 0x0d, 0x12, 0xc1, 0xce, 0x1b, 0x9e, 0xda, 0x16, 0xaf, 0xe6, 0xc3, 0x01, 0xd9, 0xbf,
	0xe9, 0x7d, 0x63, 0xe4, 0xde, 0x6a, 0x5a, 0x68, 0x83, 0x7c, 0x82, 0xad, 0xb3, 0x8c, 0x8d, 0x0a,
	0xfc, 0xc7, 0x29, 0xe8, 0xf2, 0x29, 0x68, 0x83, 0x9c, 0xc3, 0xee, 0x00, 0x13, 0x2c, 0xd9, 0xf5,
	0x00, 0xff, 0xb5, 0xc3, 0x67, 0xd8, 0x56, 0x06, 0xd5, 0x56, 0x7d, 0xb0, 0xd4, 0x9f, 0x7e, 0x92,
	0xec, 0xdd, 0x80, 0xaa, 0xb8, 0x68, 0xe3, 0xcb, 0xba, 0xfe, 0x60, 0xbc, 0xfc, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0xf7, 0x51, 0xdc, 0x4f, 0x06, 0x00, 0x00,
}