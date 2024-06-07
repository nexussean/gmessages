// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: settings.proto

package gmproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

import _ "embed"

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SIMCards        []*SIMCard      `protobuf:"bytes,2,rep,name=SIMCards,proto3" json:"SIMCards,omitempty"`
	OpCodeData      *SomeData       `protobuf:"bytes,3,opt,name=opCodeData,proto3" json:"opCodeData,omitempty"`
	RCSSettings     *RCSSettings    `protobuf:"bytes,4,opt,name=RCSSettings,proto3" json:"RCSSettings,omitempty"`
	BugleVersion    string          `protobuf:"bytes,5,opt,name=bugleVersion,proto3" json:"bugleVersion,omitempty"`
	Bool1           bool            `protobuf:"varint,7,opt,name=bool1,proto3" json:"bool1,omitempty"`
	BoolFields2     *BooleanFields2 `protobuf:"bytes,8,opt,name=boolFields2,proto3" json:"boolFields2,omitempty"`
	MysteriousBytes []byte          `protobuf:"bytes,9,opt,name=mysteriousBytes,proto3" json:"mysteriousBytes,omitempty"`
	BoolFields3     *BooleanFields3 `protobuf:"bytes,10,opt,name=boolFields3,proto3" json:"boolFields3,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetSIMCards() []*SIMCard {
	if x != nil {
		return x.SIMCards
	}
	return nil
}

func (x *Settings) GetOpCodeData() *SomeData {
	if x != nil {
		return x.OpCodeData
	}
	return nil
}

func (x *Settings) GetRCSSettings() *RCSSettings {
	if x != nil {
		return x.RCSSettings
	}
	return nil
}

func (x *Settings) GetBugleVersion() string {
	if x != nil {
		return x.BugleVersion
	}
	return ""
}

func (x *Settings) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

func (x *Settings) GetBoolFields2() *BooleanFields2 {
	if x != nil {
		return x.BoolFields2
	}
	return nil
}

func (x *Settings) GetMysteriousBytes() []byte {
	if x != nil {
		return x.MysteriousBytes
	}
	return nil
}

func (x *Settings) GetBoolFields3() *BooleanFields3 {
	if x != nil {
		return x.BoolFields3
	}
	return nil
}

type SIMCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RCSChats       *RCSChats       `protobuf:"bytes,3,opt,name=RCSChats,proto3,oneof" json:"RCSChats,omitempty"`
	SIMData        *SIMData        `protobuf:"bytes,5,opt,name=SIMData,proto3" json:"SIMData,omitempty"`
	Bool1          bool            `protobuf:"varint,6,opt,name=bool1,proto3" json:"bool1,omitempty"`
	SIMParticipant *SIMParticipant `protobuf:"bytes,7,opt,name=SIMParticipant,proto3" json:"SIMParticipant,omitempty"`
}

func (x *SIMCard) Reset() {
	*x = SIMCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SIMCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SIMCard) ProtoMessage() {}

func (x *SIMCard) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SIMCard.ProtoReflect.Descriptor instead.
func (*SIMCard) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{1}
}

func (x *SIMCard) GetRCSChats() *RCSChats {
	if x != nil {
		return x.RCSChats
	}
	return nil
}

func (x *SIMCard) GetSIMData() *SIMData {
	if x != nil {
		return x.SIMData
	}
	return nil
}

func (x *SIMCard) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

func (x *SIMCard) GetSIMParticipant() *SIMParticipant {
	if x != nil {
		return x.SIMParticipant
	}
	return nil
}

type RCSChats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *RCSChats) Reset() {
	*x = RCSChats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RCSChats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RCSChats) ProtoMessage() {}

func (x *RCSChats) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RCSChats.ProtoReflect.Descriptor instead.
func (*RCSChats) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{2}
}

func (x *RCSChats) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type BoolMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bool1 bool `protobuf:"varint,1,opt,name=bool1,proto3" json:"bool1,omitempty"`
}

func (x *BoolMsg) Reset() {
	*x = BoolMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolMsg) ProtoMessage() {}

func (x *BoolMsg) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolMsg.ProtoReflect.Descriptor instead.
func (*BoolMsg) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{3}
}

func (x *BoolMsg) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

type SIMPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Two       int32 `protobuf:"varint,1,opt,name=two,proto3" json:"two,omitempty"`
	SIMNumber int32 `protobuf:"varint,2,opt,name=SIMNumber,proto3" json:"SIMNumber,omitempty"`
}

func (x *SIMPayload) Reset() {
	*x = SIMPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SIMPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SIMPayload) ProtoMessage() {}

func (x *SIMPayload) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SIMPayload.ProtoReflect.Descriptor instead.
func (*SIMPayload) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{4}
}

func (x *SIMPayload) GetTwo() int32 {
	if x != nil {
		return x.Two
	}
	return 0
}

func (x *SIMPayload) GetSIMNumber() int32 {
	if x != nil {
		return x.SIMNumber
	}
	return 0
}

type SIMData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SIMPayload  *SIMPayload `protobuf:"bytes,1,opt,name=SIMPayload,proto3" json:"SIMPayload,omitempty"`
	Bool1       bool        `protobuf:"varint,2,opt,name=bool1,proto3" json:"bool1,omitempty"`
	CarrierName string      `protobuf:"bytes,3,opt,name=carrierName,proto3" json:"carrierName,omitempty"`
	HexHash     string      `protobuf:"bytes,4,opt,name=hexHash,proto3" json:"hexHash,omitempty"`
	Int1        int64       `protobuf:"varint,5,opt,name=int1,proto3" json:"int1,omitempty"`
}

func (x *SIMData) Reset() {
	*x = SIMData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SIMData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SIMData) ProtoMessage() {}

func (x *SIMData) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SIMData.ProtoReflect.Descriptor instead.
func (*SIMData) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{5}
}

func (x *SIMData) GetSIMPayload() *SIMPayload {
	if x != nil {
		return x.SIMPayload
	}
	return nil
}

func (x *SIMData) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

func (x *SIMData) GetCarrierName() string {
	if x != nil {
		return x.CarrierName
	}
	return ""
}

func (x *SIMData) GetHexHash() string {
	if x != nil {
		return x.HexHash
	}
	return ""
}

func (x *SIMData) GetInt1() int64 {
	if x != nil {
		return x.Int1
	}
	return 0
}

type UnknownMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int1 int64 `protobuf:"varint,1,opt,name=int1,proto3" json:"int1,omitempty"`
	Int2 int64 `protobuf:"varint,2,opt,name=int2,proto3" json:"int2,omitempty"`
}

func (x *UnknownMessage) Reset() {
	*x = UnknownMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnknownMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownMessage) ProtoMessage() {}

func (x *UnknownMessage) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownMessage.ProtoReflect.Descriptor instead.
func (*UnknownMessage) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{6}
}

func (x *UnknownMessage) GetInt1() int64 {
	if x != nil {
		return x.Int1
	}
	return 0
}

func (x *UnknownMessage) GetInt2() int64 {
	if x != nil {
		return x.Int2
	}
	return 0
}

type SIMParticipant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *SIMParticipant) Reset() {
	*x = SIMParticipant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SIMParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SIMParticipant) ProtoMessage() {}

func (x *SIMParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SIMParticipant.ProtoReflect.Descriptor instead.
func (*SIMParticipant) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{7}
}

func (x *SIMParticipant) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type SomeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field7     bool     `protobuf:"varint,7,opt,name=field7,proto3" json:"field7,omitempty"`
	Field12    bool     `protobuf:"varint,12,opt,name=field12,proto3" json:"field12,omitempty"`
	SomeEmojis [][]byte `protobuf:"bytes,15,rep,name=someEmojis,proto3" json:"someEmojis,omitempty"`
	JsonData   string   `protobuf:"bytes,16,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	SomeString string   `protobuf:"bytes,17,opt,name=someString,proto3" json:"someString,omitempty"`
}

func (x *SomeData) Reset() {
	*x = SomeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeData) ProtoMessage() {}

func (x *SomeData) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeData.ProtoReflect.Descriptor instead.
func (*SomeData) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{8}
}

func (x *SomeData) GetField7() bool {
	if x != nil {
		return x.Field7
	}
	return false
}

func (x *SomeData) GetField12() bool {
	if x != nil {
		return x.Field12
	}
	return false
}

func (x *SomeData) GetSomeEmojis() [][]byte {
	if x != nil {
		return x.SomeEmojis
	}
	return nil
}

func (x *SomeData) GetJsonData() string {
	if x != nil {
		return x.JsonData
	}
	return ""
}

func (x *SomeData) GetSomeString() string {
	if x != nil {
		return x.SomeString
	}
	return ""
}

type RCSSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnabled            bool `protobuf:"varint,1,opt,name=isEnabled,proto3" json:"isEnabled,omitempty"`
	SendReadReceipts     bool `protobuf:"varint,2,opt,name=sendReadReceipts,proto3" json:"sendReadReceipts,omitempty"`
	ShowTypingIndicators bool `protobuf:"varint,3,opt,name=showTypingIndicators,proto3" json:"showTypingIndicators,omitempty"`
	IsDefaultSMSApp      bool `protobuf:"varint,4,opt,name=isDefaultSMSApp,proto3" json:"isDefaultSMSApp,omitempty"` // uncertain, but this field seems to disappear when gmessages is un-defaulted
}

func (x *RCSSettings) Reset() {
	*x = RCSSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RCSSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RCSSettings) ProtoMessage() {}

func (x *RCSSettings) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RCSSettings.ProtoReflect.Descriptor instead.
func (*RCSSettings) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{9}
}

func (x *RCSSettings) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *RCSSettings) GetSendReadReceipts() bool {
	if x != nil {
		return x.SendReadReceipts
	}
	return false
}

func (x *RCSSettings) GetShowTypingIndicators() bool {
	if x != nil {
		return x.ShowTypingIndicators
	}
	return false
}

func (x *RCSSettings) GetIsDefaultSMSApp() bool {
	if x != nil {
		return x.IsDefaultSMSApp
	}
	return false
}

type BooleanFields2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bool1    bool     `protobuf:"varint,1,opt,name=bool1,proto3" json:"bool1,omitempty"`
	Bool2    bool     `protobuf:"varint,2,opt,name=bool2,proto3" json:"bool2,omitempty"`
	BoolMsg1 *BoolMsg `protobuf:"bytes,3,opt,name=boolMsg1,proto3" json:"boolMsg1,omitempty"`
	BoolMsg2 *BoolMsg `protobuf:"bytes,5,opt,name=boolMsg2,proto3" json:"boolMsg2,omitempty"`
	Bool3    bool     `protobuf:"varint,6,opt,name=bool3,proto3" json:"bool3,omitempty"`
}

func (x *BooleanFields2) Reset() {
	*x = BooleanFields2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanFields2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanFields2) ProtoMessage() {}

func (x *BooleanFields2) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanFields2.ProtoReflect.Descriptor instead.
func (*BooleanFields2) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{10}
}

func (x *BooleanFields2) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

func (x *BooleanFields2) GetBool2() bool {
	if x != nil {
		return x.Bool2
	}
	return false
}

func (x *BooleanFields2) GetBoolMsg1() *BoolMsg {
	if x != nil {
		return x.BoolMsg1
	}
	return nil
}

func (x *BooleanFields2) GetBoolMsg2() *BoolMsg {
	if x != nil {
		return x.BoolMsg2
	}
	return nil
}

func (x *BooleanFields2) GetBool3() bool {
	if x != nil {
		return x.Bool3
	}
	return false
}

type BooleanFields3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bool1 bool `protobuf:"varint,1,opt,name=bool1,proto3" json:"bool1,omitempty"`
	Bool3 bool `protobuf:"varint,3,opt,name=bool3,proto3" json:"bool3,omitempty"`
	Bool4 bool `protobuf:"varint,4,opt,name=bool4,proto3" json:"bool4,omitempty"`
	Bool5 bool `protobuf:"varint,5,opt,name=bool5,proto3" json:"bool5,omitempty"`
	Bool6 bool `protobuf:"varint,6,opt,name=bool6,proto3" json:"bool6,omitempty"`
	Bool7 bool `protobuf:"varint,7,opt,name=bool7,proto3" json:"bool7,omitempty"`
	Bool8 bool `protobuf:"varint,8,opt,name=bool8,proto3" json:"bool8,omitempty"`
	Bool9 bool `protobuf:"varint,9,opt,name=bool9,proto3" json:"bool9,omitempty"`
}

func (x *BooleanFields3) Reset() {
	*x = BooleanFields3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanFields3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanFields3) ProtoMessage() {}

func (x *BooleanFields3) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanFields3.ProtoReflect.Descriptor instead.
func (*BooleanFields3) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{11}
}

func (x *BooleanFields3) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

func (x *BooleanFields3) GetBool3() bool {
	if x != nil {
		return x.Bool3
	}
	return false
}

func (x *BooleanFields3) GetBool4() bool {
	if x != nil {
		return x.Bool4
	}
	return false
}

func (x *BooleanFields3) GetBool5() bool {
	if x != nil {
		return x.Bool5
	}
	return false
}

func (x *BooleanFields3) GetBool6() bool {
	if x != nil {
		return x.Bool6
	}
	return false
}

func (x *BooleanFields3) GetBool7() bool {
	if x != nil {
		return x.Bool7
	}
	return false
}

func (x *BooleanFields3) GetBool8() bool {
	if x != nil {
		return x.Bool8
	}
	return false
}

func (x *BooleanFields3) GetBool9() bool {
	if x != nil {
		return x.Bool9
	}
	return false
}

var File_settings_proto protoreflect.FileDescriptor

//go:embed settings.pb.raw
var file_settings_proto_rawDesc []byte

var (
	file_settings_proto_rawDescOnce sync.Once
	file_settings_proto_rawDescData = file_settings_proto_rawDesc
)

func file_settings_proto_rawDescGZIP() []byte {
	file_settings_proto_rawDescOnce.Do(func() {
		file_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_proto_rawDescData)
	})
	return file_settings_proto_rawDescData
}

var file_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_settings_proto_goTypes = []interface{}{
	(*Settings)(nil),       // 0: settings.Settings
	(*SIMCard)(nil),        // 1: settings.SIMCard
	(*RCSChats)(nil),       // 2: settings.RCSChats
	(*BoolMsg)(nil),        // 3: settings.BoolMsg
	(*SIMPayload)(nil),     // 4: settings.SIMPayload
	(*SIMData)(nil),        // 5: settings.SIMData
	(*UnknownMessage)(nil), // 6: settings.UnknownMessage
	(*SIMParticipant)(nil), // 7: settings.SIMParticipant
	(*SomeData)(nil),       // 8: settings.SomeData
	(*RCSSettings)(nil),    // 9: settings.RCSSettings
	(*BooleanFields2)(nil), // 10: settings.BooleanFields2
	(*BooleanFields3)(nil), // 11: settings.BooleanFields3
}
var file_settings_proto_depIdxs = []int32{
	1,  // 0: settings.Settings.SIMCards:type_name -> settings.SIMCard
	8,  // 1: settings.Settings.opCodeData:type_name -> settings.SomeData
	9,  // 2: settings.Settings.RCSSettings:type_name -> settings.RCSSettings
	10, // 3: settings.Settings.boolFields2:type_name -> settings.BooleanFields2
	11, // 4: settings.Settings.boolFields3:type_name -> settings.BooleanFields3
	2,  // 5: settings.SIMCard.RCSChats:type_name -> settings.RCSChats
	5,  // 6: settings.SIMCard.SIMData:type_name -> settings.SIMData
	7,  // 7: settings.SIMCard.SIMParticipant:type_name -> settings.SIMParticipant
	4,  // 8: settings.SIMData.SIMPayload:type_name -> settings.SIMPayload
	3,  // 9: settings.BooleanFields2.boolMsg1:type_name -> settings.BoolMsg
	3,  // 10: settings.BooleanFields2.boolMsg2:type_name -> settings.BoolMsg
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_settings_proto_init() }
func file_settings_proto_init() {
	if File_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SIMCard); i {
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
		file_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RCSChats); i {
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
		file_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolMsg); i {
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
		file_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SIMPayload); i {
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
		file_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SIMData); i {
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
		file_settings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnknownMessage); i {
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
		file_settings_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SIMParticipant); i {
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
		file_settings_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeData); i {
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
		file_settings_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RCSSettings); i {
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
		file_settings_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanFields2); i {
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
		file_settings_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanFields3); i {
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
	file_settings_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_settings_proto_goTypes,
		DependencyIndexes: file_settings_proto_depIdxs,
		MessageInfos:      file_settings_proto_msgTypes,
	}.Build()
	File_settings_proto = out.File
	file_settings_proto_rawDesc = nil
	file_settings_proto_goTypes = nil
	file_settings_proto_depIdxs = nil
}
