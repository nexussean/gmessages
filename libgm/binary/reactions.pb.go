// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: reactions.proto

package binary

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

type Reaction int32

const (
	Reaction_UNSPECIFIED Reaction = 0
	Reaction_ADD         Reaction = 1
	Reaction_REMOVE      Reaction = 2
	Reaction_SWITCH      Reaction = 3
)

// Enum value maps for Reaction.
var (
	Reaction_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ADD",
		2: "REMOVE",
		3: "SWITCH",
	}
	Reaction_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ADD":         1,
		"REMOVE":      2,
		"SWITCH":      3,
	}
)

func (x Reaction) Enum() *Reaction {
	p := new(Reaction)
	*p = x
	return p
}

func (x Reaction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reaction) Descriptor() protoreflect.EnumDescriptor {
	return file_reactions_proto_enumTypes[0].Descriptor()
}

func (Reaction) Type() protoreflect.EnumType {
	return &file_reactions_proto_enumTypes[0]
}

func (x Reaction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reaction.Descriptor instead.
func (Reaction) EnumDescriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{0}
}

type EmojiType int32

const (
	EmojiType_REACTION_TYPE_UNSPECIFIED EmojiType = 0
	EmojiType_LIKE                      EmojiType = 1
	EmojiType_LOVE                      EmojiType = 2
	EmojiType_LAUGH                     EmojiType = 3
	EmojiType_SURPRISED                 EmojiType = 4
	EmojiType_SAD                       EmojiType = 5
	EmojiType_ANGRY                     EmojiType = 6
	EmojiType_DISLIKE                   EmojiType = 7
	EmojiType_CUSTOM                    EmojiType = 8
	EmojiType_QUESTIONING               EmojiType = 9
	EmojiType_CRYING_FACE               EmojiType = 10
	EmojiType_POUTING_FACE              EmojiType = 11
	EmojiType_RED_HEART                 EmojiType = 12
)

// Enum value maps for EmojiType.
var (
	EmojiType_name = map[int32]string{
		0:  "REACTION_TYPE_UNSPECIFIED",
		1:  "LIKE",
		2:  "LOVE",
		3:  "LAUGH",
		4:  "SURPRISED",
		5:  "SAD",
		6:  "ANGRY",
		7:  "DISLIKE",
		8:  "CUSTOM",
		9:  "QUESTIONING",
		10: "CRYING_FACE",
		11: "POUTING_FACE",
		12: "RED_HEART",
	}
	EmojiType_value = map[string]int32{
		"REACTION_TYPE_UNSPECIFIED": 0,
		"LIKE":                      1,
		"LOVE":                      2,
		"LAUGH":                     3,
		"SURPRISED":                 4,
		"SAD":                       5,
		"ANGRY":                     6,
		"DISLIKE":                   7,
		"CUSTOM":                    8,
		"QUESTIONING":               9,
		"CRYING_FACE":               10,
		"POUTING_FACE":              11,
		"RED_HEART":                 12,
	}
)

func (x EmojiType) Enum() *EmojiType {
	p := new(EmojiType)
	*p = x
	return p
}

func (x EmojiType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmojiType) Descriptor() protoreflect.EnumDescriptor {
	return file_reactions_proto_enumTypes[1].Descriptor()
}

func (EmojiType) Type() protoreflect.EnumType {
	return &file_reactions_proto_enumTypes[1]
}

func (x EmojiType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmojiType.Descriptor instead.
func (EmojiType) EnumDescriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{1}
}

type SendReactionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID    string        `protobuf:"bytes,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
	ReactionData *ReactionData `protobuf:"bytes,2,opt,name=reactionData,proto3" json:"reactionData,omitempty"`
	Action       Reaction      `protobuf:"varint,3,opt,name=action,proto3,enum=reactions.Reaction" json:"action,omitempty"`
}

func (x *SendReactionPayload) Reset() {
	*x = SendReactionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReactionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReactionPayload) ProtoMessage() {}

func (x *SendReactionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReactionPayload.ProtoReflect.Descriptor instead.
func (*SendReactionPayload) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{0}
}

func (x *SendReactionPayload) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *SendReactionPayload) GetReactionData() *ReactionData {
	if x != nil {
		return x.ReactionData
	}
	return nil
}

func (x *SendReactionPayload) GetAction() Reaction {
	if x != nil {
		return x.Action
	}
	return Reaction_UNSPECIFIED
}

type SendReactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendReactionResponse) Reset() {
	*x = SendReactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReactionResponse) ProtoMessage() {}

func (x *SendReactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReactionResponse.ProtoReflect.Descriptor instead.
func (*SendReactionResponse) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{1}
}

func (x *SendReactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unicode string    `protobuf:"bytes,1,opt,name=unicode,proto3" json:"unicode,omitempty"`
	Type    EmojiType `protobuf:"varint,2,opt,name=type,proto3,enum=reactions.EmojiType" json:"type,omitempty"`
}

func (x *ReactionData) Reset() {
	*x = ReactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionData) ProtoMessage() {}

func (x *ReactionData) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionData.ProtoReflect.Descriptor instead.
func (*ReactionData) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{2}
}

func (x *ReactionData) GetUnicode() string {
	if x != nil {
		return x.Unicode
	}
	return ""
}

func (x *ReactionData) GetType() EmojiType {
	if x != nil {
		return x.Type
	}
	return EmojiType_REACTION_TYPE_UNSPECIFIED
}

type ReactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data           *ReactionData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ParticipantIDs []string      `protobuf:"bytes,2,rep,name=participantIDs,proto3" json:"participantIDs,omitempty"`
}

func (x *ReactionResponse) Reset() {
	*x = ReactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionResponse) ProtoMessage() {}

func (x *ReactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionResponse.ProtoReflect.Descriptor instead.
func (*ReactionResponse) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{3}
}

func (x *ReactionResponse) GetData() *ReactionData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReactionResponse) GetParticipantIDs() []string {
	if x != nil {
		return x.ParticipantIDs
	}
	return nil
}

type EmojiMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmojiMetaData []*EmojiMetaData `protobuf:"bytes,1,rep,name=emojiMetaData,proto3" json:"emojiMetaData,omitempty"`
}

func (x *EmojiMeta) Reset() {
	*x = EmojiMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmojiMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmojiMeta) ProtoMessage() {}

func (x *EmojiMeta) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmojiMeta.ProtoReflect.Descriptor instead.
func (*EmojiMeta) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{4}
}

func (x *EmojiMeta) GetEmojiMetaData() []*EmojiMetaData {
	if x != nil {
		return x.EmojiMetaData
	}
	return nil
}

type EmojiMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unicode string   `protobuf:"bytes,1,opt,name=unicode,proto3" json:"unicode,omitempty"`
	Names   []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *EmojiMetaData) Reset() {
	*x = EmojiMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmojiMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmojiMetaData) ProtoMessage() {}

func (x *EmojiMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmojiMetaData.ProtoReflect.Descriptor instead.
func (*EmojiMetaData) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{5}
}

func (x *EmojiMetaData) GetUnicode() string {
	if x != nil {
		return x.Unicode
	}
	return ""
}

func (x *EmojiMetaData) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_reactions_proto protoreflect.FileDescriptor

var file_reactions_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9d, 0x01, 0x0a,
	0x13, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x12, 0x3b, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x14,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x52,
	0x0a, 0x0c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x6e, 0x69, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x6e, 0x69, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x67, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x22, 0x4b, 0x0a, 0x09, 0x45,
	0x6d, 0x6f, 0x6a, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0d, 0x65, 0x6d, 0x6f, 0x6a,
	0x69, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6d, 0x6f, 0x6a,
	0x69, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x65, 0x6d, 0x6f, 0x6a, 0x69,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x0d, 0x45, 0x6d, 0x6f, 0x6a,
	0x69, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2a, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53,
	0x57, 0x49, 0x54, 0x43, 0x48, 0x10, 0x03, 0x2a, 0xc8, 0x01, 0x0a, 0x09, 0x45, 0x6d, 0x6f, 0x6a,
	0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x4c, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x41, 0x55, 0x47,
	0x48, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x52, 0x50, 0x52, 0x49, 0x53, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x41, 0x44, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x4e, 0x47, 0x52, 0x59, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x4c, 0x49, 0x4b,
	0x45, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x08, 0x12,
	0x0f, 0x0a, 0x0b, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x09,
	0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x41, 0x43, 0x45, 0x10,
	0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x41, 0x43,
	0x45, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x44, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54,
	0x10, 0x0c, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reactions_proto_rawDescOnce sync.Once
	file_reactions_proto_rawDescData = file_reactions_proto_rawDesc
)

func file_reactions_proto_rawDescGZIP() []byte {
	file_reactions_proto_rawDescOnce.Do(func() {
		file_reactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_reactions_proto_rawDescData)
	})
	return file_reactions_proto_rawDescData
}

var file_reactions_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_reactions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_reactions_proto_goTypes = []interface{}{
	(Reaction)(0),                // 0: reactions.Reaction
	(EmojiType)(0),               // 1: reactions.EmojiType
	(*SendReactionPayload)(nil),  // 2: reactions.SendReactionPayload
	(*SendReactionResponse)(nil), // 3: reactions.SendReactionResponse
	(*ReactionData)(nil),         // 4: reactions.ReactionData
	(*ReactionResponse)(nil),     // 5: reactions.ReactionResponse
	(*EmojiMeta)(nil),            // 6: reactions.EmojiMeta
	(*EmojiMetaData)(nil),        // 7: reactions.EmojiMetaData
}
var file_reactions_proto_depIdxs = []int32{
	4, // 0: reactions.SendReactionPayload.reactionData:type_name -> reactions.ReactionData
	0, // 1: reactions.SendReactionPayload.action:type_name -> reactions.Reaction
	1, // 2: reactions.ReactionData.type:type_name -> reactions.EmojiType
	4, // 3: reactions.ReactionResponse.data:type_name -> reactions.ReactionData
	7, // 4: reactions.EmojiMeta.emojiMetaData:type_name -> reactions.EmojiMetaData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_reactions_proto_init() }
func file_reactions_proto_init() {
	if File_reactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReactionPayload); i {
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
		file_reactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReactionResponse); i {
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
		file_reactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionData); i {
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
		file_reactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionResponse); i {
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
		file_reactions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmojiMeta); i {
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
		file_reactions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmojiMetaData); i {
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
			RawDescriptor: file_reactions_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reactions_proto_goTypes,
		DependencyIndexes: file_reactions_proto_depIdxs,
		EnumInfos:         file_reactions_proto_enumTypes,
		MessageInfos:      file_reactions_proto_msgTypes,
	}.Build()
	File_reactions_proto = out.File
	file_reactions_proto_rawDesc = nil
	file_reactions_proto_goTypes = nil
	file_reactions_proto_depIdxs = nil
}
