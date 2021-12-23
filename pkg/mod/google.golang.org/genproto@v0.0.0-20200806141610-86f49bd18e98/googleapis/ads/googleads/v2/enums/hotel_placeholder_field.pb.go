// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v2/enums/hotel_placeholder_field.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Possible values for Hotel placeholder fields.
type HotelPlaceholderFieldEnum_HotelPlaceholderField int32

const (
	// Not specified.
	HotelPlaceholderFieldEnum_UNSPECIFIED HotelPlaceholderFieldEnum_HotelPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	HotelPlaceholderFieldEnum_UNKNOWN HotelPlaceholderFieldEnum_HotelPlaceholderField = 1
	// Data Type: STRING. Required. Unique ID.
	HotelPlaceholderFieldEnum_PROPERTY_ID HotelPlaceholderFieldEnum_HotelPlaceholderField = 2
	// Data Type: STRING. Required. Main headline with property name to be shown
	// in dynamic ad.
	HotelPlaceholderFieldEnum_PROPERTY_NAME HotelPlaceholderFieldEnum_HotelPlaceholderField = 3
	// Data Type: STRING. Name of destination to be shown in dynamic ad.
	HotelPlaceholderFieldEnum_DESTINATION_NAME HotelPlaceholderFieldEnum_HotelPlaceholderField = 4
	// Data Type: STRING. Description of destination to be shown in dynamic ad.
	HotelPlaceholderFieldEnum_DESCRIPTION HotelPlaceholderFieldEnum_HotelPlaceholderField = 5
	// Data Type: STRING. Complete property address, including postal code.
	HotelPlaceholderFieldEnum_ADDRESS HotelPlaceholderFieldEnum_HotelPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	HotelPlaceholderFieldEnum_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	HotelPlaceholderFieldEnum_FORMATTED_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	HotelPlaceholderFieldEnum_SALE_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	HotelPlaceholderFieldEnum_FORMATTED_SALE_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	HotelPlaceholderFieldEnum_IMAGE_URL HotelPlaceholderFieldEnum_HotelPlaceholderField = 11
	// Data Type: STRING. Category of property used to group like items together
	// for recommendation engine.
	HotelPlaceholderFieldEnum_CATEGORY HotelPlaceholderFieldEnum_HotelPlaceholderField = 12
	// Data Type: INT64. Star rating (1 to 5) used to group like items
	// together for recommendation engine.
	HotelPlaceholderFieldEnum_STAR_RATING HotelPlaceholderFieldEnum_HotelPlaceholderField = 13
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	HotelPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS HotelPlaceholderFieldEnum_HotelPlaceholderField = 14
	// Data Type: URL_LIST. Required. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific flight for ads that show multiple
	// flights.
	HotelPlaceholderFieldEnum_FINAL_URLS HotelPlaceholderFieldEnum_HotelPlaceholderField = 15
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	HotelPlaceholderFieldEnum_FINAL_MOBILE_URLS HotelPlaceholderFieldEnum_HotelPlaceholderField = 16
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	HotelPlaceholderFieldEnum_TRACKING_URL HotelPlaceholderFieldEnum_HotelPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	HotelPlaceholderFieldEnum_ANDROID_APP_LINK HotelPlaceholderFieldEnum_HotelPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended property IDs to show together
	// with this item.
	HotelPlaceholderFieldEnum_SIMILAR_PROPERTY_IDS HotelPlaceholderFieldEnum_HotelPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	HotelPlaceholderFieldEnum_IOS_APP_LINK HotelPlaceholderFieldEnum_HotelPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	HotelPlaceholderFieldEnum_IOS_APP_STORE_ID HotelPlaceholderFieldEnum_HotelPlaceholderField = 21
)

// Enum value maps for HotelPlaceholderFieldEnum_HotelPlaceholderField.
var (
	HotelPlaceholderFieldEnum_HotelPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "PROPERTY_ID",
		3:  "PROPERTY_NAME",
		4:  "DESTINATION_NAME",
		5:  "DESCRIPTION",
		6:  "ADDRESS",
		7:  "PRICE",
		8:  "FORMATTED_PRICE",
		9:  "SALE_PRICE",
		10: "FORMATTED_SALE_PRICE",
		11: "IMAGE_URL",
		12: "CATEGORY",
		13: "STAR_RATING",
		14: "CONTEXTUAL_KEYWORDS",
		15: "FINAL_URLS",
		16: "FINAL_MOBILE_URLS",
		17: "TRACKING_URL",
		18: "ANDROID_APP_LINK",
		19: "SIMILAR_PROPERTY_IDS",
		20: "IOS_APP_LINK",
		21: "IOS_APP_STORE_ID",
	}
	HotelPlaceholderFieldEnum_HotelPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"PROPERTY_ID":          2,
		"PROPERTY_NAME":        3,
		"DESTINATION_NAME":     4,
		"DESCRIPTION":          5,
		"ADDRESS":              6,
		"PRICE":                7,
		"FORMATTED_PRICE":      8,
		"SALE_PRICE":           9,
		"FORMATTED_SALE_PRICE": 10,
		"IMAGE_URL":            11,
		"CATEGORY":             12,
		"STAR_RATING":          13,
		"CONTEXTUAL_KEYWORDS":  14,
		"FINAL_URLS":           15,
		"FINAL_MOBILE_URLS":    16,
		"TRACKING_URL":         17,
		"ANDROID_APP_LINK":     18,
		"SIMILAR_PROPERTY_IDS": 19,
		"IOS_APP_LINK":         20,
		"IOS_APP_STORE_ID":     21,
	}
)

func (x HotelPlaceholderFieldEnum_HotelPlaceholderField) Enum() *HotelPlaceholderFieldEnum_HotelPlaceholderField {
	p := new(HotelPlaceholderFieldEnum_HotelPlaceholderField)
	*p = x
	return p
}

func (x HotelPlaceholderFieldEnum_HotelPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HotelPlaceholderFieldEnum_HotelPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (HotelPlaceholderFieldEnum_HotelPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_enumTypes[0]
}

func (x HotelPlaceholderFieldEnum_HotelPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HotelPlaceholderFieldEnum_HotelPlaceholderField.Descriptor instead.
func (HotelPlaceholderFieldEnum_HotelPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Hotel placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type HotelPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HotelPlaceholderFieldEnum) Reset() {
	*x = HotelPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelPlaceholderFieldEnum) ProtoMessage() {}

func (x *HotelPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*HotelPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v2_enums_hotel_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x03, 0x0a, 0x19, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xaf, 0x03, 0x0a, 0x15, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43,
	0x45, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44,
	0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x41, 0x4c, 0x45,
	0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45,
	0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x0c, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x52, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0d,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x4b,
	0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x10,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c,
	0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x12, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x49, 0x4d, 0x49,
	0x4c, 0x41, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x53,
	0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49,
	0x4e, 0x4b, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x15, 0x42, 0xf0, 0x01, 0x0a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x42, 0x1b, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescData = file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_goTypes = []interface{}{
	(HotelPlaceholderFieldEnum_HotelPlaceholderField)(0), // 0: google.ads.googleads.v2.enums.HotelPlaceholderFieldEnum.HotelPlaceholderField
	(*HotelPlaceholderFieldEnum)(nil),                    // 1: google.ads.googleads.v2.enums.HotelPlaceholderFieldEnum
}
var file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_init() }
func file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_init() {
	if File_google_ads_googleads_v2_enums_hotel_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_enums_hotel_placeholder_field_proto = out.File
	file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v2_enums_hotel_placeholder_field_proto_depIdxs = nil
}
