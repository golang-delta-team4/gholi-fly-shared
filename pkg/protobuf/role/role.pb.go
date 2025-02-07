// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: role.proto

package pb

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

type AccessStatus int32

const (
	AccessStatus_ACCESS_FAILED  AccessStatus = 0
	AccessStatus_ACCESS_GRANTED AccessStatus = 1
)

// Enum value maps for AccessStatus.
var (
	AccessStatus_name = map[int32]string{
		0: "ACCESS_FAILED",
		1: "ACCESS_GRANTED",
	}
	AccessStatus_value = map[string]int32{
		"ACCESS_FAILED":  0,
		"ACCESS_GRANTED": 1,
	}
)

func (x AccessStatus) Enum() *AccessStatus {
	p := new(AccessStatus)
	*p = x
	return p
}

func (x AccessStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_role_proto_enumTypes[0].Descriptor()
}

func (AccessStatus) Type() protoreflect.EnumType {
	return &file_role_proto_enumTypes[0]
}

func (x AccessStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessStatus.Descriptor instead.
func (AccessStatus) EnumDescriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

type ResourcePermission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Route         string                 `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Method        string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourcePermission) Reset() {
	*x = ResourcePermission{}
	mi := &file_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourcePermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePermission) ProtoMessage() {}

func (x *ResourcePermission) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePermission.ProtoReflect.Descriptor instead.
func (*ResourcePermission) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *ResourcePermission) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *ResourcePermission) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type GrantResourceAccessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        AccessStatus           `protobuf:"varint,1,opt,name=status,proto3,enum=AccessStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantResourceAccessResponse) Reset() {
	*x = GrantResourceAccessResponse{}
	mi := &file_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantResourceAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantResourceAccessResponse) ProtoMessage() {}

func (x *GrantResourceAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantResourceAccessResponse.ProtoReflect.Descriptor instead.
func (*GrantResourceAccessResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *GrantResourceAccessResponse) GetStatus() AccessStatus {
	if x != nil {
		return x.Status
	}
	return AccessStatus_ACCESS_FAILED
}

type GrantResourceAccessRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerUUID     string                 `protobuf:"bytes,1,opt,name=ownerUUID,proto3" json:"ownerUUID,omitempty"`
	Permissions   []*ResourcePermission  `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	RoleName      string                 `protobuf:"bytes,3,opt,name=roleName,proto3" json:"roleName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantResourceAccessRequest) Reset() {
	*x = GrantResourceAccessRequest{}
	mi := &file_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantResourceAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantResourceAccessRequest) ProtoMessage() {}

func (x *GrantResourceAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantResourceAccessRequest.ProtoReflect.Descriptor instead.
func (*GrantResourceAccessRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *GrantResourceAccessRequest) GetOwnerUUID() string {
	if x != nil {
		return x.OwnerUUID
	}
	return ""
}

func (x *GrantResourceAccessRequest) GetPermissions() []*ResourcePermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *GrantResourceAccessRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0x44, 0x0a, 0x1b, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x1a, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x35, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x01, 0x32, 0x5f, 0x0a,
	0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x13,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e,
	0x5a, 0x0c, 0x67, 0x68, 0x6f, 0x6c, 0x69, 0x2d, 0x66, 0x6c, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_role_proto_goTypes = []any{
	(AccessStatus)(0),                   // 0: AccessStatus
	(*ResourcePermission)(nil),          // 1: ResourcePermission
	(*GrantResourceAccessResponse)(nil), // 2: GrantResourceAccessResponse
	(*GrantResourceAccessRequest)(nil),  // 3: GrantResourceAccessRequest
}
var file_role_proto_depIdxs = []int32{
	0, // 0: GrantResourceAccessResponse.status:type_name -> AccessStatus
	1, // 1: GrantResourceAccessRequest.permissions:type_name -> ResourcePermission
	3, // 2: RoleService.GrantResourceAccess:input_type -> GrantResourceAccessRequest
	2, // 3: RoleService.GrantResourceAccess:output_type -> GrantResourceAccessResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		EnumInfos:         file_role_proto_enumTypes,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
