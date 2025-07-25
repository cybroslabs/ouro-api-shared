// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: crypto/management.proto

package crypto

import (
	common "github.com/cybroslabs/ouro-api-shared/gen/go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DLMS crypto mode
type SecretDataDesryptionMethod int32

const (
	SecretDataDesryptionMethod_PLAIN     SecretDataDesryptionMethod = 0 // Plain data, no decryption needed
	SecretDataDesryptionMethod_AES256CBC SecretDataDesryptionMethod = 1 // AES-256-CBC encryption method
)

// Enum value maps for SecretDataDesryptionMethod.
var (
	SecretDataDesryptionMethod_name = map[int32]string{
		0: "PLAIN",
		1: "AES256CBC",
	}
	SecretDataDesryptionMethod_value = map[string]int32{
		"PLAIN":     0,
		"AES256CBC": 1,
	}
)

func (x SecretDataDesryptionMethod) Enum() *SecretDataDesryptionMethod {
	p := new(SecretDataDesryptionMethod)
	*p = x
	return p
}

func (x SecretDataDesryptionMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecretDataDesryptionMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_crypto_management_proto_enumTypes[0].Descriptor()
}

func (SecretDataDesryptionMethod) Type() protoreflect.EnumType {
	return &file_crypto_management_proto_enumTypes[0]
}

func (x SecretDataDesryptionMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type GetCryptoSecretRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ObjectType  common.ObjectType      `protobuf:"varint,1,opt,name=object_type,json=objectType,enum=io.clbs.openhes.models.common.ObjectType"`
	xxx_hidden_DriverType  *string                `protobuf:"bytes,2,opt,name=driver_type,json=driverType"`
	xxx_hidden_CryptoId    *string                `protobuf:"bytes,3,opt,name=crypto_id,json=cryptoId"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetCryptoSecretRequest) Reset() {
	*x = GetCryptoSecretRequest{}
	mi := &file_crypto_management_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCryptoSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptoSecretRequest) ProtoMessage() {}

func (x *GetCryptoSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_management_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetCryptoSecretRequest) GetObjectType() common.ObjectType {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			return x.xxx_hidden_ObjectType
		}
	}
	return common.ObjectType(0)
}

func (x *GetCryptoSecretRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *GetCryptoSecretRequest) GetCryptoId() string {
	if x != nil {
		if x.xxx_hidden_CryptoId != nil {
			return *x.xxx_hidden_CryptoId
		}
		return ""
	}
	return ""
}

func (x *GetCryptoSecretRequest) SetObjectType(v common.ObjectType) {
	x.xxx_hidden_ObjectType = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *GetCryptoSecretRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *GetCryptoSecretRequest) SetCryptoId(v string) {
	x.xxx_hidden_CryptoId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *GetCryptoSecretRequest) HasObjectType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetCryptoSecretRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *GetCryptoSecretRequest) HasCryptoId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *GetCryptoSecretRequest) ClearObjectType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_ObjectType = common.ObjectType_BULK
}

func (x *GetCryptoSecretRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_DriverType = nil
}

func (x *GetCryptoSecretRequest) ClearCryptoId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_CryptoId = nil
}

type GetCryptoSecretRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	ObjectType *common.ObjectType
	DriverType *string
	CryptoId   *string
}

func (b0 GetCryptoSecretRequest_builder) Build() *GetCryptoSecretRequest {
	m0 := &GetCryptoSecretRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.ObjectType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_ObjectType = *b.ObjectType
	}
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_DriverType = b.DriverType
	}
	if b.CryptoId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_CryptoId = b.CryptoId
	}
	return m0
}

type CryptoSecrets struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Secrets *[]*CryptoSecret       `protobuf:"bytes,1,rep,name=secrets"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CryptoSecrets) Reset() {
	*x = CryptoSecrets{}
	mi := &file_crypto_management_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CryptoSecrets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoSecrets) ProtoMessage() {}

func (x *CryptoSecrets) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_management_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CryptoSecrets) GetSecrets() []*CryptoSecret {
	if x != nil {
		if x.xxx_hidden_Secrets != nil {
			return *x.xxx_hidden_Secrets
		}
	}
	return nil
}

func (x *CryptoSecrets) SetSecrets(v []*CryptoSecret) {
	x.xxx_hidden_Secrets = &v
}

type CryptoSecrets_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Secrets []*CryptoSecret
}

func (b0 CryptoSecrets_builder) Build() *CryptoSecrets {
	m0 := &CryptoSecrets{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Secrets = &b.Secrets
	return m0
}

type CryptoSecret struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_AccessLevel *string                `protobuf:"bytes,1,opt,name=access_level,json=accessLevel"`
	xxx_hidden_KeyId       *string                `protobuf:"bytes,2,opt,name=key_id,json=keyId"`
	xxx_hidden_CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt"`
	xxx_hidden_UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt"`
	xxx_hidden_Data        []byte                 `protobuf:"bytes,15,opt,name=data"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *CryptoSecret) Reset() {
	*x = CryptoSecret{}
	mi := &file_crypto_management_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CryptoSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoSecret) ProtoMessage() {}

func (x *CryptoSecret) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_management_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CryptoSecret) GetAccessLevel() string {
	if x != nil {
		if x.xxx_hidden_AccessLevel != nil {
			return *x.xxx_hidden_AccessLevel
		}
		return ""
	}
	return ""
}

func (x *CryptoSecret) GetKeyId() string {
	if x != nil {
		if x.xxx_hidden_KeyId != nil {
			return *x.xxx_hidden_KeyId
		}
		return ""
	}
	return ""
}

func (x *CryptoSecret) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_CreatedAt
	}
	return nil
}

func (x *CryptoSecret) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_UpdatedAt
	}
	return nil
}

func (x *CryptoSecret) GetData() []byte {
	if x != nil {
		return x.xxx_hidden_Data
	}
	return nil
}

func (x *CryptoSecret) SetAccessLevel(v string) {
	x.xxx_hidden_AccessLevel = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 5)
}

func (x *CryptoSecret) SetKeyId(v string) {
	x.xxx_hidden_KeyId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 5)
}

func (x *CryptoSecret) SetCreatedAt(v *timestamppb.Timestamp) {
	x.xxx_hidden_CreatedAt = v
}

func (x *CryptoSecret) SetUpdatedAt(v *timestamppb.Timestamp) {
	x.xxx_hidden_UpdatedAt = v
}

func (x *CryptoSecret) SetData(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Data = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 5)
}

func (x *CryptoSecret) HasAccessLevel() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *CryptoSecret) HasKeyId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *CryptoSecret) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_CreatedAt != nil
}

func (x *CryptoSecret) HasUpdatedAt() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_UpdatedAt != nil
}

func (x *CryptoSecret) HasData() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *CryptoSecret) ClearAccessLevel() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_AccessLevel = nil
}

func (x *CryptoSecret) ClearKeyId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_KeyId = nil
}

func (x *CryptoSecret) ClearCreatedAt() {
	x.xxx_hidden_CreatedAt = nil
}

func (x *CryptoSecret) ClearUpdatedAt() {
	x.xxx_hidden_UpdatedAt = nil
}

func (x *CryptoSecret) ClearData() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Data = nil
}

type CryptoSecret_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	AccessLevel *string
	KeyId       *string
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
	Data        []byte
}

func (b0 CryptoSecret_builder) Build() *CryptoSecret {
	m0 := &CryptoSecret{}
	b, x := &b0, m0
	_, _ = b, x
	if b.AccessLevel != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 5)
		x.xxx_hidden_AccessLevel = b.AccessLevel
	}
	if b.KeyId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 5)
		x.xxx_hidden_KeyId = b.KeyId
	}
	x.xxx_hidden_CreatedAt = b.CreatedAt
	x.xxx_hidden_UpdatedAt = b.UpdatedAt
	if b.Data != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 5)
		x.xxx_hidden_Data = b.Data
	}
	return m0
}

type SetCryptoSecretRequest struct {
	state                             protoimpl.MessageState     `protogen:"opaque.v1"`
	xxx_hidden_ObjectType             common.ObjectType          `protobuf:"varint,1,opt,name=object_type,json=objectType,enum=io.clbs.openhes.models.common.ObjectType"`
	xxx_hidden_DriverType             *string                    `protobuf:"bytes,2,opt,name=driver_type,json=driverType"`
	xxx_hidden_CryptoId               *string                    `protobuf:"bytes,3,opt,name=crypto_id,json=cryptoId"`
	xxx_hidden_AccessLevel            *string                    `protobuf:"bytes,4,opt,name=access_level,json=accessLevel"`
	xxx_hidden_KeyId                  *string                    `protobuf:"bytes,5,opt,name=key_id,json=keyId"`
	xxx_hidden_DataDecryptionSecretId *string                    `protobuf:"bytes,6,opt,name=data_decryption_secret_id,json=dataDecryptionSecretId"`
	xxx_hidden_DataDecryptionMethod   SecretDataDesryptionMethod `protobuf:"varint,7,opt,name=data_decryption_method,json=dataDecryptionMethod,enum=io.clbs.openhes.models.crypto.SecretDataDesryptionMethod"`
	xxx_hidden_DataDecryptionIv       []byte                     `protobuf:"bytes,8,opt,name=data_decryption_iv,json=dataDecryptionIv"`
	xxx_hidden_Data                   []byte                     `protobuf:"bytes,15,opt,name=data"`
	XXX_raceDetectHookData            protoimpl.RaceDetectHookData
	XXX_presence                      [1]uint32
	unknownFields                     protoimpl.UnknownFields
	sizeCache                         protoimpl.SizeCache
}

func (x *SetCryptoSecretRequest) Reset() {
	*x = SetCryptoSecretRequest{}
	mi := &file_crypto_management_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetCryptoSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCryptoSecretRequest) ProtoMessage() {}

func (x *SetCryptoSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_management_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SetCryptoSecretRequest) GetObjectType() common.ObjectType {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			return x.xxx_hidden_ObjectType
		}
	}
	return common.ObjectType(0)
}

func (x *SetCryptoSecretRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *SetCryptoSecretRequest) GetCryptoId() string {
	if x != nil {
		if x.xxx_hidden_CryptoId != nil {
			return *x.xxx_hidden_CryptoId
		}
		return ""
	}
	return ""
}

func (x *SetCryptoSecretRequest) GetAccessLevel() string {
	if x != nil {
		if x.xxx_hidden_AccessLevel != nil {
			return *x.xxx_hidden_AccessLevel
		}
		return ""
	}
	return ""
}

func (x *SetCryptoSecretRequest) GetKeyId() string {
	if x != nil {
		if x.xxx_hidden_KeyId != nil {
			return *x.xxx_hidden_KeyId
		}
		return ""
	}
	return ""
}

func (x *SetCryptoSecretRequest) GetDataDecryptionSecretId() string {
	if x != nil {
		if x.xxx_hidden_DataDecryptionSecretId != nil {
			return *x.xxx_hidden_DataDecryptionSecretId
		}
		return ""
	}
	return ""
}

func (x *SetCryptoSecretRequest) GetDataDecryptionMethod() SecretDataDesryptionMethod {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 6) {
			return x.xxx_hidden_DataDecryptionMethod
		}
	}
	return SecretDataDesryptionMethod_PLAIN
}

func (x *SetCryptoSecretRequest) GetDataDecryptionIv() []byte {
	if x != nil {
		return x.xxx_hidden_DataDecryptionIv
	}
	return nil
}

func (x *SetCryptoSecretRequest) GetData() []byte {
	if x != nil {
		return x.xxx_hidden_Data
	}
	return nil
}

func (x *SetCryptoSecretRequest) SetObjectType(v common.ObjectType) {
	x.xxx_hidden_ObjectType = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 9)
}

func (x *SetCryptoSecretRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 9)
}

func (x *SetCryptoSecretRequest) SetCryptoId(v string) {
	x.xxx_hidden_CryptoId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 9)
}

func (x *SetCryptoSecretRequest) SetAccessLevel(v string) {
	x.xxx_hidden_AccessLevel = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 9)
}

func (x *SetCryptoSecretRequest) SetKeyId(v string) {
	x.xxx_hidden_KeyId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 9)
}

func (x *SetCryptoSecretRequest) SetDataDecryptionSecretId(v string) {
	x.xxx_hidden_DataDecryptionSecretId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 9)
}

func (x *SetCryptoSecretRequest) SetDataDecryptionMethod(v SecretDataDesryptionMethod) {
	x.xxx_hidden_DataDecryptionMethod = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 6, 9)
}

func (x *SetCryptoSecretRequest) SetDataDecryptionIv(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_DataDecryptionIv = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 7, 9)
}

func (x *SetCryptoSecretRequest) SetData(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Data = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 8, 9)
}

func (x *SetCryptoSecretRequest) HasObjectType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *SetCryptoSecretRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *SetCryptoSecretRequest) HasCryptoId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *SetCryptoSecretRequest) HasAccessLevel() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *SetCryptoSecretRequest) HasKeyId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *SetCryptoSecretRequest) HasDataDecryptionSecretId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *SetCryptoSecretRequest) HasDataDecryptionMethod() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 6)
}

func (x *SetCryptoSecretRequest) HasDataDecryptionIv() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 7)
}

func (x *SetCryptoSecretRequest) HasData() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 8)
}

func (x *SetCryptoSecretRequest) ClearObjectType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_ObjectType = common.ObjectType_BULK
}

func (x *SetCryptoSecretRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_DriverType = nil
}

func (x *SetCryptoSecretRequest) ClearCryptoId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_CryptoId = nil
}

func (x *SetCryptoSecretRequest) ClearAccessLevel() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_AccessLevel = nil
}

func (x *SetCryptoSecretRequest) ClearKeyId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_KeyId = nil
}

func (x *SetCryptoSecretRequest) ClearDataDecryptionSecretId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_DataDecryptionSecretId = nil
}

func (x *SetCryptoSecretRequest) ClearDataDecryptionMethod() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 6)
	x.xxx_hidden_DataDecryptionMethod = SecretDataDesryptionMethod_PLAIN
}

func (x *SetCryptoSecretRequest) ClearDataDecryptionIv() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 7)
	x.xxx_hidden_DataDecryptionIv = nil
}

func (x *SetCryptoSecretRequest) ClearData() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 8)
	x.xxx_hidden_Data = nil
}

type SetCryptoSecretRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	ObjectType             *common.ObjectType
	DriverType             *string
	CryptoId               *string
	AccessLevel            *string
	KeyId                  *string
	DataDecryptionSecretId *string
	DataDecryptionMethod   *SecretDataDesryptionMethod
	DataDecryptionIv       []byte
	Data                   []byte
}

func (b0 SetCryptoSecretRequest_builder) Build() *SetCryptoSecretRequest {
	m0 := &SetCryptoSecretRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.ObjectType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 9)
		x.xxx_hidden_ObjectType = *b.ObjectType
	}
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 9)
		x.xxx_hidden_DriverType = b.DriverType
	}
	if b.CryptoId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 9)
		x.xxx_hidden_CryptoId = b.CryptoId
	}
	if b.AccessLevel != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 9)
		x.xxx_hidden_AccessLevel = b.AccessLevel
	}
	if b.KeyId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 9)
		x.xxx_hidden_KeyId = b.KeyId
	}
	if b.DataDecryptionSecretId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 9)
		x.xxx_hidden_DataDecryptionSecretId = b.DataDecryptionSecretId
	}
	if b.DataDecryptionMethod != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 6, 9)
		x.xxx_hidden_DataDecryptionMethod = *b.DataDecryptionMethod
	}
	if b.DataDecryptionIv != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 7, 9)
		x.xxx_hidden_DataDecryptionIv = b.DataDecryptionIv
	}
	if b.Data != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 8, 9)
		x.xxx_hidden_Data = b.Data
	}
	return m0
}

var File_crypto_management_proto protoreflect.FileDescriptor

const file_crypto_management_proto_rawDesc = "" +
	"\n" +
	"\x17crypto/management.proto\x12\x1dio.clbs.openhes.models.crypto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x14common/objects.proto\"\xa2\x01\n" +
	"\x16GetCryptoSecretRequest\x12J\n" +
	"\vobject_type\x18\x01 \x01(\x0e2).io.clbs.openhes.models.common.ObjectTypeR\n" +
	"objectType\x12\x1f\n" +
	"\vdriver_type\x18\x02 \x01(\tR\n" +
	"driverType\x12\x1b\n" +
	"\tcrypto_id\x18\x03 \x01(\tR\bcryptoId\"V\n" +
	"\rCryptoSecrets\x12E\n" +
	"\asecrets\x18\x01 \x03(\v2+.io.clbs.openhes.models.crypto.CryptoSecretR\asecrets\"\x8e\x02\n" +
	"\fCryptoSecret\x12!\n" +
	"\faccess_level\x18\x01 \x01(\tR\vaccessLevel\x12\x15\n" +
	"\x06key_id\x18\x02 \x01(\tR\x05keyId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x12\n" +
	"\x04data\x18\x0f \x01(\fR\x04dataJ\x04\b\x05\x10\x06J\x04\b\x06\x10\aJ\x04\b\a\x10\bJ\x04\b\b\x10\tJ\x04\b\t\x10\n" +
	"J\x04\b\n" +
	"\x10\vJ\x04\b\v\x10\fJ\x04\b\f\x10\rJ\x04\b\r\x10\x0eJ\x04\b\x0e\x10\x0f\"\xee\x03\n" +
	"\x16SetCryptoSecretRequest\x12J\n" +
	"\vobject_type\x18\x01 \x01(\x0e2).io.clbs.openhes.models.common.ObjectTypeR\n" +
	"objectType\x12\x1f\n" +
	"\vdriver_type\x18\x02 \x01(\tR\n" +
	"driverType\x12\x1b\n" +
	"\tcrypto_id\x18\x03 \x01(\tR\bcryptoId\x12!\n" +
	"\faccess_level\x18\x04 \x01(\tR\vaccessLevel\x12\x15\n" +
	"\x06key_id\x18\x05 \x01(\tR\x05keyId\x129\n" +
	"\x19data_decryption_secret_id\x18\x06 \x01(\tR\x16dataDecryptionSecretId\x12o\n" +
	"\x16data_decryption_method\x18\a \x01(\x0e29.io.clbs.openhes.models.crypto.SecretDataDesryptionMethodR\x14dataDecryptionMethod\x12,\n" +
	"\x12data_decryption_iv\x18\b \x01(\fR\x10dataDecryptionIv\x12\x12\n" +
	"\x04data\x18\x0f \x01(\fR\x04dataJ\x04\b\t\x10\n" +
	"J\x04\b\n" +
	"\x10\vJ\x04\b\v\x10\fJ\x04\b\f\x10\rJ\x04\b\r\x10\x0eJ\x04\b\x0e\x10\x0f*6\n" +
	"\x1aSecretDataDesryptionMethod\x12\t\n" +
	"\x05PLAIN\x10\x00\x12\r\n" +
	"\tAES256CBC\x10\x01B5Z3github.com/cybroslabs/ouro-api-shared/gen/go/cryptob\beditionsp\xe8\a"

var file_crypto_management_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crypto_management_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crypto_management_proto_goTypes = []any{
	(SecretDataDesryptionMethod)(0), // 0: io.clbs.openhes.models.crypto.SecretDataDesryptionMethod
	(*GetCryptoSecretRequest)(nil),  // 1: io.clbs.openhes.models.crypto.GetCryptoSecretRequest
	(*CryptoSecrets)(nil),           // 2: io.clbs.openhes.models.crypto.CryptoSecrets
	(*CryptoSecret)(nil),            // 3: io.clbs.openhes.models.crypto.CryptoSecret
	(*SetCryptoSecretRequest)(nil),  // 4: io.clbs.openhes.models.crypto.SetCryptoSecretRequest
	(common.ObjectType)(0),          // 5: io.clbs.openhes.models.common.ObjectType
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_crypto_management_proto_depIdxs = []int32{
	5, // 0: io.clbs.openhes.models.crypto.GetCryptoSecretRequest.object_type:type_name -> io.clbs.openhes.models.common.ObjectType
	3, // 1: io.clbs.openhes.models.crypto.CryptoSecrets.secrets:type_name -> io.clbs.openhes.models.crypto.CryptoSecret
	6, // 2: io.clbs.openhes.models.crypto.CryptoSecret.created_at:type_name -> google.protobuf.Timestamp
	6, // 3: io.clbs.openhes.models.crypto.CryptoSecret.updated_at:type_name -> google.protobuf.Timestamp
	5, // 4: io.clbs.openhes.models.crypto.SetCryptoSecretRequest.object_type:type_name -> io.clbs.openhes.models.common.ObjectType
	0, // 5: io.clbs.openhes.models.crypto.SetCryptoSecretRequest.data_decryption_method:type_name -> io.clbs.openhes.models.crypto.SecretDataDesryptionMethod
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_crypto_management_proto_init() }
func file_crypto_management_proto_init() {
	if File_crypto_management_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_crypto_management_proto_rawDesc), len(file_crypto_management_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_management_proto_goTypes,
		DependencyIndexes: file_crypto_management_proto_depIdxs,
		EnumInfos:         file_crypto_management_proto_enumTypes,
		MessageInfos:      file_crypto_management_proto_msgTypes,
	}.Build()
	File_crypto_management_proto = out.File
	file_crypto_management_proto_goTypes = nil
	file_crypto_management_proto_depIdxs = nil
}
