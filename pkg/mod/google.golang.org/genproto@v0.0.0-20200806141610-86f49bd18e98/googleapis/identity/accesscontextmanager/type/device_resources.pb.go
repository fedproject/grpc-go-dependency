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
// source: google/identity/accesscontextmanager/type/device_resources.proto

package _type

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

// The encryption state of the device.
type DeviceEncryptionStatus int32

const (
	// The encryption status of the device is not specified or not known.
	DeviceEncryptionStatus_ENCRYPTION_UNSPECIFIED DeviceEncryptionStatus = 0
	// The device does not support encryption.
	DeviceEncryptionStatus_ENCRYPTION_UNSUPPORTED DeviceEncryptionStatus = 1
	// The device supports encryption, but is currently unencrypted.
	DeviceEncryptionStatus_UNENCRYPTED DeviceEncryptionStatus = 2
	// The device is encrypted.
	DeviceEncryptionStatus_ENCRYPTED DeviceEncryptionStatus = 3
)

// Enum value maps for DeviceEncryptionStatus.
var (
	DeviceEncryptionStatus_name = map[int32]string{
		0: "ENCRYPTION_UNSPECIFIED",
		1: "ENCRYPTION_UNSUPPORTED",
		2: "UNENCRYPTED",
		3: "ENCRYPTED",
	}
	DeviceEncryptionStatus_value = map[string]int32{
		"ENCRYPTION_UNSPECIFIED": 0,
		"ENCRYPTION_UNSUPPORTED": 1,
		"UNENCRYPTED":            2,
		"ENCRYPTED":              3,
	}
)

func (x DeviceEncryptionStatus) Enum() *DeviceEncryptionStatus {
	p := new(DeviceEncryptionStatus)
	*p = x
	return p
}

func (x DeviceEncryptionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceEncryptionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes[0].Descriptor()
}

func (DeviceEncryptionStatus) Type() protoreflect.EnumType {
	return &file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes[0]
}

func (x DeviceEncryptionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceEncryptionStatus.Descriptor instead.
func (DeviceEncryptionStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescGZIP(), []int{0}
}

// The operating system type of the device.
// Next id: 7
type OsType int32

const (
	// The operating system of the device is not specified or not known.
	OsType_OS_UNSPECIFIED OsType = 0
	// A desktop Mac operating system.
	OsType_DESKTOP_MAC OsType = 1
	// A desktop Windows operating system.
	OsType_DESKTOP_WINDOWS OsType = 2
	// A desktop Linux operating system.
	OsType_DESKTOP_LINUX OsType = 3
	// A desktop ChromeOS operating system.
	OsType_DESKTOP_CHROME_OS OsType = 6
	// An Android operating system.
	OsType_ANDROID OsType = 4
	// An iOS operating system.
	OsType_IOS OsType = 5
)

// Enum value maps for OsType.
var (
	OsType_name = map[int32]string{
		0: "OS_UNSPECIFIED",
		1: "DESKTOP_MAC",
		2: "DESKTOP_WINDOWS",
		3: "DESKTOP_LINUX",
		6: "DESKTOP_CHROME_OS",
		4: "ANDROID",
		5: "IOS",
	}
	OsType_value = map[string]int32{
		"OS_UNSPECIFIED":    0,
		"DESKTOP_MAC":       1,
		"DESKTOP_WINDOWS":   2,
		"DESKTOP_LINUX":     3,
		"DESKTOP_CHROME_OS": 6,
		"ANDROID":           4,
		"IOS":               5,
	}
)

func (x OsType) Enum() *OsType {
	p := new(OsType)
	*p = x
	return p
}

func (x OsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OsType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes[1].Descriptor()
}

func (OsType) Type() protoreflect.EnumType {
	return &file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes[1]
}

func (x OsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OsType.Descriptor instead.
func (OsType) EnumDescriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescGZIP(), []int{1}
}

// The degree to which the device is managed by the Cloud organization.
type DeviceManagementLevel int32

const (
	// The device's management level is not specified or not known.
	DeviceManagementLevel_MANAGEMENT_UNSPECIFIED DeviceManagementLevel = 0
	// The device is not managed.
	DeviceManagementLevel_NONE DeviceManagementLevel = 1
	// Basic management is enabled, which is generally limited to monitoring and
	// wiping the corporate account.
	DeviceManagementLevel_BASIC DeviceManagementLevel = 2
	// Complete device management. This includes more thorough monitoring and the
	// ability to directly manage the device (such as remote wiping). This can be
	// enabled through the Android Enterprise Platform.
	DeviceManagementLevel_COMPLETE DeviceManagementLevel = 3
)

// Enum value maps for DeviceManagementLevel.
var (
	DeviceManagementLevel_name = map[int32]string{
		0: "MANAGEMENT_UNSPECIFIED",
		1: "NONE",
		2: "BASIC",
		3: "COMPLETE",
	}
	DeviceManagementLevel_value = map[string]int32{
		"MANAGEMENT_UNSPECIFIED": 0,
		"NONE":                   1,
		"BASIC":                  2,
		"COMPLETE":               3,
	}
)

func (x DeviceManagementLevel) Enum() *DeviceManagementLevel {
	p := new(DeviceManagementLevel)
	*p = x
	return p
}

func (x DeviceManagementLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceManagementLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes[2].Descriptor()
}

func (DeviceManagementLevel) Type() protoreflect.EnumType {
	return &file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes[2]
}

func (x DeviceManagementLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceManagementLevel.Descriptor instead.
func (DeviceManagementLevel) EnumDescriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescGZIP(), []int{2}
}

var File_google_identity_accesscontextmanager_type_device_resources_proto protoreflect.FileDescriptor

var file_google_identity_accesscontextmanager_type_device_resources_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x70, 0x0a, 0x16, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x82, 0x01,
	0x0a, 0x06, 0x4f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x49,
	0x4e, 0x55, 0x58, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50,
	0x5f, 0x43, 0x48, 0x52, 0x4f, 0x4d, 0x45, 0x5f, 0x4f, 0x53, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53,
	0x10, 0x05, 0x2a, 0x56, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x16, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x42, 0x92, 0x02, 0x0a, 0x2d, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x09, 0x54, 0x79,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x3b, 0x74, 0x79, 0x70, 0x65, 0xaa, 0x02, 0x29, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x29, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x54, 0x79, 0x70, 0x65,
	0xea, 0x02, 0x2c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescOnce sync.Once
	file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescData = file_google_identity_accesscontextmanager_type_device_resources_proto_rawDesc
)

func file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescGZIP() []byte {
	file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescOnce.Do(func() {
		file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescData)
	})
	return file_google_identity_accesscontextmanager_type_device_resources_proto_rawDescData
}

var file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_identity_accesscontextmanager_type_device_resources_proto_goTypes = []interface{}{
	(DeviceEncryptionStatus)(0), // 0: google.identity.accesscontextmanager.type.DeviceEncryptionStatus
	(OsType)(0),                 // 1: google.identity.accesscontextmanager.type.OsType
	(DeviceManagementLevel)(0),  // 2: google.identity.accesscontextmanager.type.DeviceManagementLevel
}
var file_google_identity_accesscontextmanager_type_device_resources_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_identity_accesscontextmanager_type_device_resources_proto_init() }
func file_google_identity_accesscontextmanager_type_device_resources_proto_init() {
	if File_google_identity_accesscontextmanager_type_device_resources_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_identity_accesscontextmanager_type_device_resources_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_identity_accesscontextmanager_type_device_resources_proto_goTypes,
		DependencyIndexes: file_google_identity_accesscontextmanager_type_device_resources_proto_depIdxs,
		EnumInfos:         file_google_identity_accesscontextmanager_type_device_resources_proto_enumTypes,
	}.Build()
	File_google_identity_accesscontextmanager_type_device_resources_proto = out.File
	file_google_identity_accesscontextmanager_type_device_resources_proto_rawDesc = nil
	file_google_identity_accesscontextmanager_type_device_resources_proto_goTypes = nil
	file_google_identity_accesscontextmanager_type_device_resources_proto_depIdxs = nil
}
