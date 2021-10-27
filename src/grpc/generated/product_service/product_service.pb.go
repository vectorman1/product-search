// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: product_service.proto

package product_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *ProductSearchRequest) Reset() {
	*x = ProductSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchRequest) ProtoMessage() {}

func (x *ProductSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchRequest.ProtoReflect.Descriptor instead.
func (*ProductSearchRequest) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProductSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProductSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Product `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProductSearchResponse) Reset() {
	*x = ProductSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchResponse) ProtoMessage() {}

func (x *ProductSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchResponse.ProtoReflect.Descriptor instead.
func (*ProductSearchResponse) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSearchResponse) GetItems() []*Product {
	if x != nil {
		return x.Items
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category            string   `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
	Subcategory         string   `protobuf:"bytes,2,opt,name=Subcategory,proto3" json:"Subcategory,omitempty"`
	Name                string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	CurrentPrice        float32  `protobuf:"fixed32,4,opt,name=CurrentPrice,proto3" json:"CurrentPrice,omitempty"`
	RawPrice            float32  `protobuf:"fixed32,5,opt,name=RawPrice,proto3" json:"RawPrice,omitempty"`
	Currency            string   `protobuf:"bytes,6,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Discount            int32    `protobuf:"varint,7,opt,name=Discount,proto3" json:"Discount,omitempty"`
	LikesCount          int32    `protobuf:"varint,8,opt,name=LikesCount,proto3" json:"LikesCount,omitempty"`
	IsNew               bool     `protobuf:"varint,9,opt,name=IsNew,proto3" json:"IsNew,omitempty"`
	Brand               string   `protobuf:"bytes,10,opt,name=Brand,proto3" json:"Brand,omitempty"`
	BrandUrl            string   `protobuf:"bytes,11,opt,name=BrandUrl,proto3" json:"BrandUrl,omitempty"`
	CodCountry          []string `protobuf:"bytes,12,rep,name=CodCountry,proto3" json:"CodCountry,omitempty"`
	Variation0Color     string   `protobuf:"bytes,13,opt,name=Variation0Color,proto3" json:"Variation0Color,omitempty"`
	Variation1Color     string   `protobuf:"bytes,14,opt,name=Variation1Color,proto3" json:"Variation1Color,omitempty"`
	Variation0Thumbnail string   `protobuf:"bytes,15,opt,name=Variation0Thumbnail,proto3" json:"Variation0Thumbnail,omitempty"`
	Variation0Image     string   `protobuf:"bytes,16,opt,name=Variation0Image,proto3" json:"Variation0Image,omitempty"`
	Variation1Thumbnail string   `protobuf:"bytes,17,opt,name=Variation1Thumbnail,proto3" json:"Variation1Thumbnail,omitempty"`
	Variation1Image     string   `protobuf:"bytes,18,opt,name=Variation1Image,proto3" json:"Variation1Image,omitempty"`
	ImageUrl            string   `protobuf:"bytes,19,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	Url                 string   `protobuf:"bytes,20,opt,name=Url,proto3" json:"Url,omitempty"`
	Id                  int32    `protobuf:"varint,21,opt,name=Id,proto3" json:"Id,omitempty"`
	Model               string   `protobuf:"bytes,22,opt,name=Model,proto3" json:"Model,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetSubcategory() string {
	if x != nil {
		return x.Subcategory
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetCurrentPrice() float32 {
	if x != nil {
		return x.CurrentPrice
	}
	return 0
}

func (x *Product) GetRawPrice() float32 {
	if x != nil {
		return x.RawPrice
	}
	return 0
}

func (x *Product) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Product) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Product) GetLikesCount() int32 {
	if x != nil {
		return x.LikesCount
	}
	return 0
}

func (x *Product) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *Product) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Product) GetBrandUrl() string {
	if x != nil {
		return x.BrandUrl
	}
	return ""
}

func (x *Product) GetCodCountry() []string {
	if x != nil {
		return x.CodCountry
	}
	return nil
}

func (x *Product) GetVariation0Color() string {
	if x != nil {
		return x.Variation0Color
	}
	return ""
}

func (x *Product) GetVariation1Color() string {
	if x != nil {
		return x.Variation1Color
	}
	return ""
}

func (x *Product) GetVariation0Thumbnail() string {
	if x != nil {
		return x.Variation0Thumbnail
	}
	return ""
}

func (x *Product) GetVariation0Image() string {
	if x != nil {
		return x.Variation0Image
	}
	return ""
}

func (x *Product) GetVariation1Thumbnail() string {
	if x != nil {
		return x.Variation1Thumbnail
	}
	return ""
}

func (x *Product) GetVariation1Image() string {
	if x != nil {
		return x.Variation1Image
	}
	return ""
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Product) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Product) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

var File_product_service_proto protoreflect.FileDescriptor

var file_product_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xbb, 0x05,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x75, 0x62, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x61, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x52, 0x61, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x4e, 0x65, 0x77, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x31, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x30, 0x0a, 0x13, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x31, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a,
	0x0f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x31, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x32, 0x6b, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_service_proto_rawDescOnce sync.Once
	file_product_service_proto_rawDescData = file_product_service_proto_rawDesc
)

func file_product_service_proto_rawDescGZIP() []byte {
	file_product_service_proto_rawDescOnce.Do(func() {
		file_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_service_proto_rawDescData)
	})
	return file_product_service_proto_rawDescData
}

var file_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_product_service_proto_goTypes = []interface{}{
	(*ProductSearchRequest)(nil),  // 0: product_service.ProductSearchRequest
	(*ProductSearchResponse)(nil), // 1: product_service.ProductSearchResponse
	(*Product)(nil),               // 2: product_service.Product
}
var file_product_service_proto_depIdxs = []int32{
	2, // 0: product_service.ProductSearchResponse.items:type_name -> product_service.Product
	0, // 1: product_service.ProductService.Search:input_type -> product_service.ProductSearchRequest
	1, // 2: product_service.ProductService.Search:output_type -> product_service.ProductSearchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_service_proto_init() }
func file_product_service_proto_init() {
	if File_product_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_service_proto_goTypes,
		DependencyIndexes: file_product_service_proto_depIdxs,
		MessageInfos:      file_product_service_proto_msgTypes,
	}.Build()
	File_product_service_proto = out.File
	file_product_service_proto_rawDesc = nil
	file_product_service_proto_goTypes = nil
	file_product_service_proto_depIdxs = nil
}
