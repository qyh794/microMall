// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cart/cart.proto

/*
Package go_micro_service_cart is a generated protocol buffer package.

It is generated from these files:
	proto/cart/cart.proto

It has these top-level messages:
	RequestAddCartInfo
	ResponseAddCart
	RequestCleanCart
	Response
	RequestItem
	RequestCartID
	RequestCartFindAll
	ResponseCartAll
*/
package go_micro_service_cart

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

type RequestAddCartInfo struct {
	Id        int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,4,opt,name=size_id,json=sizeId" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,5,opt,name=num" json:"num,omitempty"`
}

func (m *RequestAddCartInfo) Reset()                    { *m = RequestAddCartInfo{} }
func (m *RequestAddCartInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestAddCartInfo) ProtoMessage()               {}
func (*RequestAddCartInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestAddCartInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestAddCartInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RequestAddCartInfo) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *RequestAddCartInfo) GetSizeId() int64 {
	if m != nil {
		return m.SizeId
	}
	return 0
}

func (m *RequestAddCartInfo) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ResponseAddCart struct {
	CartId int64  `protobuf:"varint,1,opt,name=cart_id,json=cartId" json:"cart_id,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ResponseAddCart) Reset()                    { *m = ResponseAddCart{} }
func (m *ResponseAddCart) String() string            { return proto.CompactTextString(m) }
func (*ResponseAddCart) ProtoMessage()               {}
func (*ResponseAddCart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseAddCart) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *ResponseAddCart) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestCleanCart struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *RequestCleanCart) Reset()                    { *m = RequestCleanCart{} }
func (m *RequestCleanCart) String() string            { return proto.CompactTextString(m) }
func (*RequestCleanCart) ProtoMessage()               {}
func (*RequestCleanCart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestCleanCart) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestItem struct {
	Id        int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ChangeNum int64 `protobuf:"varint,2,opt,name=change_num,json=changeNum" json:"change_num,omitempty"`
}

func (m *RequestItem) Reset()                    { *m = RequestItem{} }
func (m *RequestItem) String() string            { return proto.CompactTextString(m) }
func (*RequestItem) ProtoMessage()               {}
func (*RequestItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestItem) GetChangeNum() int64 {
	if m != nil {
		return m.ChangeNum
	}
	return 0
}

type RequestCartID struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *RequestCartID) Reset()                    { *m = RequestCartID{} }
func (m *RequestCartID) String() string            { return proto.CompactTextString(m) }
func (*RequestCartID) ProtoMessage()               {}
func (*RequestCartID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestCartID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RequestCartFindAll struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *RequestCartFindAll) Reset()                    { *m = RequestCartFindAll{} }
func (m *RequestCartFindAll) String() string            { return proto.CompactTextString(m) }
func (*RequestCartFindAll) ProtoMessage()               {}
func (*RequestCartFindAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RequestCartFindAll) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ResponseCartAll struct {
	CartInfo []*RequestAddCartInfo `protobuf:"bytes,1,rep,name=cart_info,json=cartInfo" json:"cart_info,omitempty"`
}

func (m *ResponseCartAll) Reset()                    { *m = ResponseCartAll{} }
func (m *ResponseCartAll) String() string            { return proto.CompactTextString(m) }
func (*ResponseCartAll) ProtoMessage()               {}
func (*ResponseCartAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResponseCartAll) GetCartInfo() []*RequestAddCartInfo {
	if m != nil {
		return m.CartInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestAddCartInfo)(nil), "go.micro.service.cart.RequestAddCartInfo")
	proto.RegisterType((*ResponseAddCart)(nil), "go.micro.service.cart.ResponseAddCart")
	proto.RegisterType((*RequestCleanCart)(nil), "go.micro.service.cart.RequestCleanCart")
	proto.RegisterType((*Response)(nil), "go.micro.service.cart.Response")
	proto.RegisterType((*RequestItem)(nil), "go.micro.service.cart.RequestItem")
	proto.RegisterType((*RequestCartID)(nil), "go.micro.service.cart.RequestCartID")
	proto.RegisterType((*RequestCartFindAll)(nil), "go.micro.service.cart.RequestCartFindAll")
	proto.RegisterType((*ResponseCartAll)(nil), "go.micro.service.cart.ResponseCartAll")
}

func init() { proto.RegisterFile("proto/cart/cart.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x63, 0xe3, 0xd6, 0x53, 0x51, 0xaa, 0x95, 0x2a, 0xac, 0x8a, 0xaa, 0xd5, 0x0a, 0x41,
	0x2b, 0x84, 0x91, 0xca, 0x35, 0x97, 0x10, 0x2b, 0xc8, 0x07, 0x38, 0xf8, 0x82, 0x72, 0x80, 0xc8,
	0x78, 0x27, 0xc1, 0x92, 0xbd, 0x1b, 0xd6, 0x36, 0x12, 0xfc, 0x01, 0x7e, 0x02, 0x7f, 0x17, 0xcd,
	0xda, 0x4e, 0x80, 0x90, 0x38, 0x87, 0x5e, 0xac, 0x9d, 0x8f, 0xf7, 0xf6, 0xed, 0x9b, 0x91, 0xe1,
	0x7c, 0xa5, 0x55, 0xa5, 0x5e, 0xa5, 0x89, 0xae, 0xcc, 0x27, 0x30, 0x31, 0x3b, 0x5f, 0xaa, 0xa0,
	0xc8, 0x52, 0xad, 0x82, 0x12, 0xf5, 0xb7, 0x2c, 0xc5, 0x80, 0x8a, 0xfc, 0xa7, 0x05, 0x2c, 0xc6,
	0xaf, 0x35, 0x96, 0xd5, 0x58, 0x88, 0x49, 0xa2, 0xab, 0x48, 0x2e, 0x14, 0x3b, 0x85, 0x61, 0x26,
	0x7c, 0xeb, 0xda, 0xba, 0xb1, 0xe3, 0x61, 0x26, 0xd8, 0x63, 0x38, 0xaa, 0x4b, 0xd4, 0xf3, 0x4c,
	0xf8, 0x43, 0x93, 0x74, 0x29, 0x8c, 0x04, 0xbb, 0x04, 0x58, 0x69, 0x25, 0xea, 0xb4, 0xa2, 0x9a,
	0x6d, 0x6a, 0x5e, 0x9b, 0x89, 0x0c, 0xae, 0xcc, 0x7e, 0x20, 0xd5, 0x9c, 0x06, 0x47, 0x61, 0x24,
	0xd8, 0x19, 0xd8, 0xb2, 0x2e, 0xfc, 0x07, 0x26, 0x49, 0x47, 0x3e, 0x82, 0x47, 0x31, 0x96, 0x2b,
	0x25, 0x4b, 0x6c, 0x95, 0x10, 0x9a, 0x44, 0xce, 0xd7, 0x52, 0x5c, 0x0a, 0x1b, 0x74, 0x51, 0x2e,
	0x8d, 0x14, 0x2f, 0xa6, 0x23, 0x7f, 0x01, 0x67, 0xed, 0x33, 0x26, 0x39, 0x26, 0xb2, 0x83, 0x77,
	0xa2, 0xad, 0x3f, 0x45, 0xf3, 0x27, 0x70, 0xdc, 0x5d, 0xd5, 0x51, 0x59, 0x1b, 0xaa, 0x11, 0x9c,
	0xb4, 0x54, 0x51, 0x85, 0xc5, 0x96, 0x15, 0x97, 0x00, 0xe9, 0x97, 0x44, 0x2e, 0x71, 0x4e, 0x0f,
	0x68, 0xdc, 0xf0, 0x9a, 0xcc, 0xfb, 0xba, 0xe0, 0x57, 0xf0, 0xb0, 0x13, 0x42, 0x5a, 0xc3, 0x7f,
	0xf1, 0xfc, 0xe5, 0xda, 0x70, 0x6a, 0x98, 0x66, 0x52, 0x8c, 0xf3, 0x7c, 0xb7, 0xd6, 0xd9, 0xc6,
	0x16, 0xea, 0xa7, 0xde, 0x29, 0x78, 0x8d, 0x2d, 0x72, 0xa1, 0x7c, 0xeb, 0xda, 0xbe, 0x39, 0xb9,
	0xbb, 0x0d, 0xfe, 0x3b, 0xde, 0x60, 0x7b, 0xb4, 0xf1, 0x71, 0xda, 0x9e, 0xee, 0x7e, 0x39, 0xe0,
	0x18, 0xa3, 0x3e, 0xc1, 0x51, 0x67, 0xf9, 0xe1, 0x44, 0x17, 0xcf, 0x76, 0xb6, 0xfe, 0x35, 0x45,
	0x3e, 0x60, 0x1f, 0xc0, 0xdb, 0x4c, 0xe5, 0xf9, 0xfe, 0x1b, 0xd6, 0x8d, 0x17, 0x57, 0x3d, 0xfc,
	0x7c, 0xc0, 0xde, 0x81, 0x13, 0xc9, 0x54, 0x33, 0xbe, 0x9f, 0x93, 0xe6, 0x78, 0x20, 0x5d, 0x88,
	0xf7, 0x47, 0x37, 0x83, 0xd3, 0x10, 0x73, 0xac, 0x90, 0x00, 0x6f, 0xbe, 0x47, 0x21, 0x7b, 0xda,
	0xf3, 0x76, 0xb3, 0x31, 0x87, 0x50, 0x7f, 0x04, 0xf7, 0x2d, 0x9a, 0x65, 0xb8, 0xed, 0xa7, 0x6c,
	0x77, 0xac, 0x77, 0x60, 0xed, 0x7e, 0xf1, 0xc1, 0x67, 0xd7, 0xfc, 0x33, 0x5e, 0xff, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x18, 0x00, 0x44, 0x4c, 0x04, 0x00, 0x00,
}
