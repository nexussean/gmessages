// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: rpc.proto

package gmproto

import (
	reflect "reflect"
	sync "sync"

	_ "go.mau.fi/util/pblite"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BugleRoute int32

const (
	BugleRoute_Unknown   BugleRoute = 0
	BugleRoute_DataEvent BugleRoute = 19
	BugleRoute_PairEvent BugleRoute = 14
	BugleRoute_GaiaEvent BugleRoute = 7
)

// Enum value maps for BugleRoute.
var (
	BugleRoute_name = map[int32]string{
		0:  "Unknown",
		19: "DataEvent",
		14: "PairEvent",
		7:  "GaiaEvent",
	}
	BugleRoute_value = map[string]int32{
		"Unknown":   0,
		"DataEvent": 19,
		"PairEvent": 14,
		"GaiaEvent": 7,
	}
)

func (x BugleRoute) Enum() *BugleRoute {
	p := new(BugleRoute)
	*p = x
	return p
}

func (x BugleRoute) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BugleRoute) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[0].Descriptor()
}

func (BugleRoute) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[0]
}

func (x BugleRoute) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BugleRoute.Descriptor instead.
func (BugleRoute) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

type ActionType int32

const (
	ActionType_UNSPECIFIED                            ActionType = 0
	ActionType_LIST_CONVERSATIONS                     ActionType = 1
	ActionType_LIST_MESSAGES                          ActionType = 2
	ActionType_SEND_MESSAGE                           ActionType = 3
	ActionType_MESSAGE_UPDATES                        ActionType = 4
	ActionType_LIST_CONTACTS                          ActionType = 6
	ActionType_CONVERSATION_UPDATES                   ActionType = 7
	ActionType_GET_OR_CREATE_CONVERSATION             ActionType = 9
	ActionType_MESSAGE_READ                           ActionType = 10
	ActionType_BROWSER_PRESENCE_CHECK                 ActionType = 11
	ActionType_TYPING_UPDATES                         ActionType = 12
	ActionType_SETTINGS_UPDATE                        ActionType = 13
	ActionType_USER_ALERT                             ActionType = 14
	ActionType_UPDATE_CONVERSATION                    ActionType = 15
	ActionType_GET_UPDATES                            ActionType = 16
	ActionType_ACK_BROWSER_PRESENCE                   ActionType = 17
	ActionType_LIST_STICKER_SETS                      ActionType = 18
	ActionType_LEAVE_RCS_GROUP                        ActionType = 19
	ActionType_ADD_PARTICIPANT_TO_RCS_GROUP           ActionType = 20
	ActionType_GET_CONVERSATION_TYPE                  ActionType = 21
	ActionType_NOTIFY_DITTO_ACTIVITY                  ActionType = 22
	ActionType_DELETE_MESSAGE                         ActionType = 23
	ActionType_INSTALL_STICKER_SET                    ActionType = 24
	ActionType_RESEND_MESSAGE                         ActionType = 25
	ActionType_GET_CONTACT_RCS_GROUP_STATUS           ActionType = 26
	ActionType_DOWNLOAD_MESSAGE                       ActionType = 27
	ActionType_LIST_TOP_CONTACTS                      ActionType = 28
	ActionType_GET_CONTACTS_THUMBNAIL                 ActionType = 29
	ActionType_CHANGE_PARTICIPANT_COLOR               ActionType = 30
	ActionType_IS_BUGLE_DEFAULT                       ActionType = 31
	ActionType_STICKER_USER_CONTEXT                   ActionType = 32
	ActionType_FAVORITE_STICKER_PACKS                 ActionType = 33
	ActionType_RECENT_STICKERS                        ActionType = 34
	ActionType_UPDATE_RECENT_STICKERS                 ActionType = 35
	ActionType_GET_FULL_SIZE_IMAGE                    ActionType = 36
	ActionType_GET_PARTICIPANTS_THUMBNAIL             ActionType = 37
	ActionType_SEND_REACTION                          ActionType = 38
	ActionType_SEND_REPLY                             ActionType = 39
	ActionType_GET_BLOB_FOR_ATTACHMENT                ActionType = 40
	ActionType_GET_DEVICES_AVAILABLE_FOR_GAIA_PAIRING ActionType = 41
	ActionType_CREATE_GAIA_PAIRING                    ActionType = 42
	ActionType_GET_CONVERSATION                       ActionType = 43
	ActionType_CREATE_GAIA_PAIRING_CLIENT_INIT        ActionType = 44
	ActionType_CREATE_GAIA_PAIRING_CLIENT_FINISHED    ActionType = 45
	ActionType_UNPAIR_GAIA_PAIRING                    ActionType = 46
	ActionType_CANCEL_GAIA_PAIRING                    ActionType = 47
	ActionType_PREWARM                                ActionType = 48
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "LIST_CONVERSATIONS",
		2:  "LIST_MESSAGES",
		3:  "SEND_MESSAGE",
		4:  "MESSAGE_UPDATES",
		6:  "LIST_CONTACTS",
		7:  "CONVERSATION_UPDATES",
		9:  "GET_OR_CREATE_CONVERSATION",
		10: "MESSAGE_READ",
		11: "BROWSER_PRESENCE_CHECK",
		12: "TYPING_UPDATES",
		13: "SETTINGS_UPDATE",
		14: "USER_ALERT",
		15: "UPDATE_CONVERSATION",
		16: "GET_UPDATES",
		17: "ACK_BROWSER_PRESENCE",
		18: "LIST_STICKER_SETS",
		19: "LEAVE_RCS_GROUP",
		20: "ADD_PARTICIPANT_TO_RCS_GROUP",
		21: "GET_CONVERSATION_TYPE",
		22: "NOTIFY_DITTO_ACTIVITY",
		23: "DELETE_MESSAGE",
		24: "INSTALL_STICKER_SET",
		25: "RESEND_MESSAGE",
		26: "GET_CONTACT_RCS_GROUP_STATUS",
		27: "DOWNLOAD_MESSAGE",
		28: "LIST_TOP_CONTACTS",
		29: "GET_CONTACTS_THUMBNAIL",
		30: "CHANGE_PARTICIPANT_COLOR",
		31: "IS_BUGLE_DEFAULT",
		32: "STICKER_USER_CONTEXT",
		33: "FAVORITE_STICKER_PACKS",
		34: "RECENT_STICKERS",
		35: "UPDATE_RECENT_STICKERS",
		36: "GET_FULL_SIZE_IMAGE",
		37: "GET_PARTICIPANTS_THUMBNAIL",
		38: "SEND_REACTION",
		39: "SEND_REPLY",
		40: "GET_BLOB_FOR_ATTACHMENT",
		41: "GET_DEVICES_AVAILABLE_FOR_GAIA_PAIRING",
		42: "CREATE_GAIA_PAIRING",
		43: "GET_CONVERSATION",
		44: "CREATE_GAIA_PAIRING_CLIENT_INIT",
		45: "CREATE_GAIA_PAIRING_CLIENT_FINISHED",
		46: "UNPAIR_GAIA_PAIRING",
		47: "CANCEL_GAIA_PAIRING",
		48: "PREWARM",
	}
	ActionType_value = map[string]int32{
		"UNSPECIFIED":                            0,
		"LIST_CONVERSATIONS":                     1,
		"LIST_MESSAGES":                          2,
		"SEND_MESSAGE":                           3,
		"MESSAGE_UPDATES":                        4,
		"LIST_CONTACTS":                          6,
		"CONVERSATION_UPDATES":                   7,
		"GET_OR_CREATE_CONVERSATION":             9,
		"MESSAGE_READ":                           10,
		"BROWSER_PRESENCE_CHECK":                 11,
		"TYPING_UPDATES":                         12,
		"SETTINGS_UPDATE":                        13,
		"USER_ALERT":                             14,
		"UPDATE_CONVERSATION":                    15,
		"GET_UPDATES":                            16,
		"ACK_BROWSER_PRESENCE":                   17,
		"LIST_STICKER_SETS":                      18,
		"LEAVE_RCS_GROUP":                        19,
		"ADD_PARTICIPANT_TO_RCS_GROUP":           20,
		"GET_CONVERSATION_TYPE":                  21,
		"NOTIFY_DITTO_ACTIVITY":                  22,
		"DELETE_MESSAGE":                         23,
		"INSTALL_STICKER_SET":                    24,
		"RESEND_MESSAGE":                         25,
		"GET_CONTACT_RCS_GROUP_STATUS":           26,
		"DOWNLOAD_MESSAGE":                       27,
		"LIST_TOP_CONTACTS":                      28,
		"GET_CONTACTS_THUMBNAIL":                 29,
		"CHANGE_PARTICIPANT_COLOR":               30,
		"IS_BUGLE_DEFAULT":                       31,
		"STICKER_USER_CONTEXT":                   32,
		"FAVORITE_STICKER_PACKS":                 33,
		"RECENT_STICKERS":                        34,
		"UPDATE_RECENT_STICKERS":                 35,
		"GET_FULL_SIZE_IMAGE":                    36,
		"GET_PARTICIPANTS_THUMBNAIL":             37,
		"SEND_REACTION":                          38,
		"SEND_REPLY":                             39,
		"GET_BLOB_FOR_ATTACHMENT":                40,
		"GET_DEVICES_AVAILABLE_FOR_GAIA_PAIRING": 41,
		"CREATE_GAIA_PAIRING":                    42,
		"GET_CONVERSATION":                       43,
		"CREATE_GAIA_PAIRING_CLIENT_INIT":        44,
		"CREATE_GAIA_PAIRING_CLIENT_FINISHED":    45,
		"UNPAIR_GAIA_PAIRING":                    46,
		"CANCEL_GAIA_PAIRING":                    47,
		"PREWARM":                                48,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

type MessageType int32

const (
	MessageType_UNKNOWN_MESSAGE_TYPE MessageType = 0
	MessageType_BUGLE_MESSAGE        MessageType = 2
	MessageType_GAIA_1               MessageType = 3
	MessageType_BUGLE_ANNOTATION     MessageType = 16
	MessageType_GAIA_2               MessageType = 20
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:  "UNKNOWN_MESSAGE_TYPE",
		2:  "BUGLE_MESSAGE",
		3:  "GAIA_1",
		16: "BUGLE_ANNOTATION",
		20: "GAIA_2",
	}
	MessageType_value = map[string]int32{
		"UNKNOWN_MESSAGE_TYPE": 0,
		"BUGLE_MESSAGE":        2,
		"GAIA_1":               3,
		"BUGLE_ANNOTATION":     16,
		"GAIA_2":               20,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[2].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[2]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

type StartAckMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         *int32                 `protobuf:"varint,1,opt,name=count,proto3,oneof" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartAckMessage) Reset() {
	*x = StartAckMessage{}
	mi := &file_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartAckMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartAckMessage) ProtoMessage() {}

func (x *StartAckMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartAckMessage.ProtoReflect.Descriptor instead.
func (*StartAckMessage) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *StartAckMessage) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type LongPollingPayload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *IncomingRPCMessage    `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Heartbeat     *EmptyArr              `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof" json:"heartbeat,omitempty"`
	Ack           *StartAckMessage       `protobuf:"bytes,4,opt,name=ack,proto3,oneof" json:"ack,omitempty"`
	StartRead     *EmptyArr              `protobuf:"bytes,5,opt,name=startRead,proto3,oneof" json:"startRead,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LongPollingPayload) Reset() {
	*x = LongPollingPayload{}
	mi := &file_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LongPollingPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongPollingPayload) ProtoMessage() {}

func (x *LongPollingPayload) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongPollingPayload.ProtoReflect.Descriptor instead.
func (*LongPollingPayload) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *LongPollingPayload) GetData() *IncomingRPCMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LongPollingPayload) GetHeartbeat() *EmptyArr {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

func (x *LongPollingPayload) GetAck() *StartAckMessage {
	if x != nil {
		return x.Ack
	}
	return nil
}

func (x *LongPollingPayload) GetStartRead() *EmptyArr {
	if x != nil {
		return x.StartRead
	}
	return nil
}

type IncomingRPCMessage struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ResponseID        string                 `protobuf:"bytes,1,opt,name=responseID,proto3" json:"responseID,omitempty"`
	BugleRoute        BugleRoute             `protobuf:"varint,2,opt,name=bugleRoute,proto3,enum=rpc.BugleRoute" json:"bugleRoute,omitempty"`
	StartExecute      uint64                 `protobuf:"varint,3,opt,name=startExecute,proto3" json:"startExecute,omitempty"`
	MessageType       MessageType            `protobuf:"varint,5,opt,name=messageType,proto3,enum=rpc.MessageType" json:"messageType,omitempty"`
	FinishExecute     uint64                 `protobuf:"varint,6,opt,name=finishExecute,proto3" json:"finishExecute,omitempty"`
	MicrosecondsTaken uint64                 `protobuf:"varint,7,opt,name=microsecondsTaken,proto3" json:"microsecondsTaken,omitempty"`
	Mobile            *Device                `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Browser           *Device                `protobuf:"bytes,9,opt,name=browser,proto3" json:"browser,omitempty"`
	// Either a RPCMessageData or a RPCPairData encoded as bytes
	MessageData []byte `protobuf:"bytes,12,opt,name=messageData,proto3" json:"messageData,omitempty"`
	SignatureID string `protobuf:"bytes,17,opt,name=signatureID,proto3" json:"signatureID,omitempty"`
	Timestamp   string `protobuf:"bytes,21,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Completely unsure about this, but it seems to be present for weird intermediate responses
	GdittoSource  *IncomingRPCMessage_GDittoSource `protobuf:"bytes,23,opt,name=gdittoSource,proto3" json:"gdittoSource,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncomingRPCMessage) Reset() {
	*x = IncomingRPCMessage{}
	mi := &file_rpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncomingRPCMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingRPCMessage) ProtoMessage() {}

func (x *IncomingRPCMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingRPCMessage.ProtoReflect.Descriptor instead.
func (*IncomingRPCMessage) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *IncomingRPCMessage) GetResponseID() string {
	if x != nil {
		return x.ResponseID
	}
	return ""
}

func (x *IncomingRPCMessage) GetBugleRoute() BugleRoute {
	if x != nil {
		return x.BugleRoute
	}
	return BugleRoute_Unknown
}

func (x *IncomingRPCMessage) GetStartExecute() uint64 {
	if x != nil {
		return x.StartExecute
	}
	return 0
}

func (x *IncomingRPCMessage) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (x *IncomingRPCMessage) GetFinishExecute() uint64 {
	if x != nil {
		return x.FinishExecute
	}
	return 0
}

func (x *IncomingRPCMessage) GetMicrosecondsTaken() uint64 {
	if x != nil {
		return x.MicrosecondsTaken
	}
	return 0
}

func (x *IncomingRPCMessage) GetMobile() *Device {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *IncomingRPCMessage) GetBrowser() *Device {
	if x != nil {
		return x.Browser
	}
	return nil
}

func (x *IncomingRPCMessage) GetMessageData() []byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *IncomingRPCMessage) GetSignatureID() string {
	if x != nil {
		return x.SignatureID
	}
	return ""
}

func (x *IncomingRPCMessage) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *IncomingRPCMessage) GetGdittoSource() *IncomingRPCMessage_GDittoSource {
	if x != nil {
		return x.GdittoSource
	}
	return nil
}

type RPCMessageData struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SessionID       string                 `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Timestamp       int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Action          ActionType             `protobuf:"varint,4,opt,name=action,proto3,enum=rpc.ActionType" json:"action,omitempty"`
	UnencryptedData []byte                 `protobuf:"bytes,5,opt,name=unencryptedData,proto3" json:"unencryptedData,omitempty"`
	Bool1           bool                   `protobuf:"varint,6,opt,name=bool1,proto3" json:"bool1,omitempty"`
	Bool2           bool                   `protobuf:"varint,7,opt,name=bool2,proto3" json:"bool2,omitempty"`
	EncryptedData   []byte                 `protobuf:"bytes,8,opt,name=encryptedData,proto3" json:"encryptedData,omitempty"`
	Bool3           bool                   `protobuf:"varint,9,opt,name=bool3,proto3" json:"bool3,omitempty"`
	EncryptedData2  []byte                 `protobuf:"bytes,11,opt,name=encryptedData2,proto3" json:"encryptedData2,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RPCMessageData) Reset() {
	*x = RPCMessageData{}
	mi := &file_rpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RPCMessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMessageData) ProtoMessage() {}

func (x *RPCMessageData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMessageData.ProtoReflect.Descriptor instead.
func (*RPCMessageData) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *RPCMessageData) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *RPCMessageData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RPCMessageData) GetAction() ActionType {
	if x != nil {
		return x.Action
	}
	return ActionType_UNSPECIFIED
}

func (x *RPCMessageData) GetUnencryptedData() []byte {
	if x != nil {
		return x.UnencryptedData
	}
	return nil
}

func (x *RPCMessageData) GetBool1() bool {
	if x != nil {
		return x.Bool1
	}
	return false
}

func (x *RPCMessageData) GetBool2() bool {
	if x != nil {
		return x.Bool2
	}
	return false
}

func (x *RPCMessageData) GetEncryptedData() []byte {
	if x != nil {
		return x.EncryptedData
	}
	return nil
}

func (x *RPCMessageData) GetBool3() bool {
	if x != nil {
		return x.Bool3
	}
	return false
}

func (x *RPCMessageData) GetEncryptedData2() []byte {
	if x != nil {
		return x.EncryptedData2
	}
	return nil
}

type OutgoingRPCMessage struct {
	state               protoimpl.MessageState   `protogen:"open.v1"`
	Mobile              *Device                  `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Data                *OutgoingRPCMessage_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Auth                *OutgoingRPCMessage_Auth `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	TTL                 int64                    `protobuf:"varint,5,opt,name=TTL,proto3" json:"TTL,omitempty"`
	DestRegistrationIDs []string                 `protobuf:"bytes,9,rep,name=destRegistrationIDs,proto3" json:"destRegistrationIDs,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *OutgoingRPCMessage) Reset() {
	*x = OutgoingRPCMessage{}
	mi := &file_rpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCMessage) ProtoMessage() {}

func (x *OutgoingRPCMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCMessage.ProtoReflect.Descriptor instead.
func (*OutgoingRPCMessage) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *OutgoingRPCMessage) GetMobile() *Device {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *OutgoingRPCMessage) GetData() *OutgoingRPCMessage_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OutgoingRPCMessage) GetAuth() *OutgoingRPCMessage_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *OutgoingRPCMessage) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}

func (x *OutgoingRPCMessage) GetDestRegistrationIDs() []string {
	if x != nil {
		return x.DestRegistrationIDs
	}
	return nil
}

type OutgoingRPCData struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	RequestID            string                 `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Action               ActionType             `protobuf:"varint,2,opt,name=action,proto3,enum=rpc.ActionType" json:"action,omitempty"`
	UnencryptedProtoData []byte                 `protobuf:"bytes,3,opt,name=unencryptedProtoData,proto3" json:"unencryptedProtoData,omitempty"`
	EncryptedProtoData   []byte                 `protobuf:"bytes,5,opt,name=encryptedProtoData,proto3" json:"encryptedProtoData,omitempty"`
	SessionID            string                 `protobuf:"bytes,6,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *OutgoingRPCData) Reset() {
	*x = OutgoingRPCData{}
	mi := &file_rpc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCData) ProtoMessage() {}

func (x *OutgoingRPCData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCData.ProtoReflect.Descriptor instead.
func (*OutgoingRPCData) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *OutgoingRPCData) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *OutgoingRPCData) GetAction() ActionType {
	if x != nil {
		return x.Action
	}
	return ActionType_UNSPECIFIED
}

func (x *OutgoingRPCData) GetUnencryptedProtoData() []byte {
	if x != nil {
		return x.UnencryptedProtoData
	}
	return nil
}

func (x *OutgoingRPCData) GetEncryptedProtoData() []byte {
	if x != nil {
		return x.EncryptedProtoData
	}
	return nil
}

func (x *OutgoingRPCData) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type OutgoingRPCResponse struct {
	state          protoimpl.MessageState              `protogen:"open.v1"`
	SomeIdentifier *OutgoingRPCResponse_SomeIdentifier `protobuf:"bytes,1,opt,name=someIdentifier,proto3" json:"someIdentifier,omitempty"`
	// This is not present for AckMessage responses, only for SendMessage
	Timestamp     *string `protobuf:"bytes,2,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutgoingRPCResponse) Reset() {
	*x = OutgoingRPCResponse{}
	mi := &file_rpc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCResponse) ProtoMessage() {}

func (x *OutgoingRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCResponse.ProtoReflect.Descriptor instead.
func (*OutgoingRPCResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *OutgoingRPCResponse) GetSomeIdentifier() *OutgoingRPCResponse_SomeIdentifier {
	if x != nil {
		return x.SomeIdentifier
	}
	return nil
}

func (x *OutgoingRPCResponse) GetTimestamp() string {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return ""
}

type IncomingRPCMessage_GDittoSource struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceID      int32                  `protobuf:"varint,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncomingRPCMessage_GDittoSource) Reset() {
	*x = IncomingRPCMessage_GDittoSource{}
	mi := &file_rpc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncomingRPCMessage_GDittoSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingRPCMessage_GDittoSource) ProtoMessage() {}

func (x *IncomingRPCMessage_GDittoSource) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingRPCMessage_GDittoSource.ProtoReflect.Descriptor instead.
func (*IncomingRPCMessage_GDittoSource) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2, 0}
}

func (x *IncomingRPCMessage_GDittoSource) GetDeviceID() int32 {
	if x != nil {
		return x.DeviceID
	}
	return 0
}

type OutgoingRPCMessage_Auth struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RequestID        string                 `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	TachyonAuthToken []byte                 `protobuf:"bytes,6,opt,name=tachyonAuthToken,proto3" json:"tachyonAuthToken,omitempty"`
	ConfigVersion    *ConfigVersion         `protobuf:"bytes,7,opt,name=configVersion,proto3" json:"configVersion,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *OutgoingRPCMessage_Auth) Reset() {
	*x = OutgoingRPCMessage_Auth{}
	mi := &file_rpc_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCMessage_Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCMessage_Auth) ProtoMessage() {}

func (x *OutgoingRPCMessage_Auth) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCMessage_Auth.ProtoReflect.Descriptor instead.
func (*OutgoingRPCMessage_Auth) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4, 0}
}

func (x *OutgoingRPCMessage_Auth) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *OutgoingRPCMessage_Auth) GetTachyonAuthToken() []byte {
	if x != nil {
		return x.TachyonAuthToken
	}
	return nil
}

func (x *OutgoingRPCMessage_Auth) GetConfigVersion() *ConfigVersion {
	if x != nil {
		return x.ConfigVersion
	}
	return nil
}

type OutgoingRPCMessage_Data struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	RequestID  string                 `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	BugleRoute BugleRoute             `protobuf:"varint,2,opt,name=bugleRoute,proto3,enum=rpc.BugleRoute" json:"bugleRoute,omitempty"`
	// OutgoingRPCData encoded as bytes
	MessageData     []byte                        `protobuf:"bytes,12,opt,name=messageData,proto3" json:"messageData,omitempty"`
	MessageTypeData *OutgoingRPCMessage_Data_Type `protobuf:"bytes,23,opt,name=messageTypeData,proto3" json:"messageTypeData,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OutgoingRPCMessage_Data) Reset() {
	*x = OutgoingRPCMessage_Data{}
	mi := &file_rpc_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCMessage_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCMessage_Data) ProtoMessage() {}

func (x *OutgoingRPCMessage_Data) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCMessage_Data.ProtoReflect.Descriptor instead.
func (*OutgoingRPCMessage_Data) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4, 1}
}

func (x *OutgoingRPCMessage_Data) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *OutgoingRPCMessage_Data) GetBugleRoute() BugleRoute {
	if x != nil {
		return x.BugleRoute
	}
	return BugleRoute_Unknown
}

func (x *OutgoingRPCMessage_Data) GetMessageData() []byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *OutgoingRPCMessage_Data) GetMessageTypeData() *OutgoingRPCMessage_Data_Type {
	if x != nil {
		return x.MessageTypeData
	}
	return nil
}

type OutgoingRPCMessage_Data_Type struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EmptyArr      *EmptyArr              `protobuf:"bytes,1,opt,name=emptyArr,proto3" json:"emptyArr,omitempty"`
	MessageType   MessageType            `protobuf:"varint,2,opt,name=messageType,proto3,enum=rpc.MessageType" json:"messageType,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutgoingRPCMessage_Data_Type) Reset() {
	*x = OutgoingRPCMessage_Data_Type{}
	mi := &file_rpc_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCMessage_Data_Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCMessage_Data_Type) ProtoMessage() {}

func (x *OutgoingRPCMessage_Data_Type) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCMessage_Data_Type.ProtoReflect.Descriptor instead.
func (*OutgoingRPCMessage_Data_Type) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4, 1, 0}
}

func (x *OutgoingRPCMessage_Data_Type) GetEmptyArr() *EmptyArr {
	if x != nil {
		return x.EmptyArr
	}
	return nil
}

func (x *OutgoingRPCMessage_Data_Type) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

type OutgoingRPCResponse_SomeIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 1 -> unknown
	SomeNumber    string `protobuf:"bytes,2,opt,name=someNumber,proto3" json:"someNumber,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutgoingRPCResponse_SomeIdentifier) Reset() {
	*x = OutgoingRPCResponse_SomeIdentifier{}
	mi := &file_rpc_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutgoingRPCResponse_SomeIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingRPCResponse_SomeIdentifier) ProtoMessage() {}

func (x *OutgoingRPCResponse_SomeIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingRPCResponse_SomeIdentifier.ProtoReflect.Descriptor instead.
func (*OutgoingRPCResponse_SomeIdentifier) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{6, 0}
}

func (x *OutgoingRPCResponse_SomeIdentifier) GetSomeNumber() string {
	if x != nil {
		return x.SomeNumber
	}
	return ""
}

var File_rpc_proto protoreflect.FileDescriptor

//go:embed rpc.pb.raw
var file_rpc_proto_rawDesc []byte

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_rpc_proto_goTypes = []any{
	(BugleRoute)(0),                            // 0: rpc.BugleRoute
	(ActionType)(0),                            // 1: rpc.ActionType
	(MessageType)(0),                           // 2: rpc.MessageType
	(*StartAckMessage)(nil),                    // 3: rpc.StartAckMessage
	(*LongPollingPayload)(nil),                 // 4: rpc.LongPollingPayload
	(*IncomingRPCMessage)(nil),                 // 5: rpc.IncomingRPCMessage
	(*RPCMessageData)(nil),                     // 6: rpc.RPCMessageData
	(*OutgoingRPCMessage)(nil),                 // 7: rpc.OutgoingRPCMessage
	(*OutgoingRPCData)(nil),                    // 8: rpc.OutgoingRPCData
	(*OutgoingRPCResponse)(nil),                // 9: rpc.OutgoingRPCResponse
	(*IncomingRPCMessage_GDittoSource)(nil),    // 10: rpc.IncomingRPCMessage.GDittoSource
	(*OutgoingRPCMessage_Auth)(nil),            // 11: rpc.OutgoingRPCMessage.Auth
	(*OutgoingRPCMessage_Data)(nil),            // 12: rpc.OutgoingRPCMessage.Data
	(*OutgoingRPCMessage_Data_Type)(nil),       // 13: rpc.OutgoingRPCMessage.Data.Type
	(*OutgoingRPCResponse_SomeIdentifier)(nil), // 14: rpc.OutgoingRPCResponse.SomeIdentifier
	(*EmptyArr)(nil),                           // 15: util.EmptyArr
	(*Device)(nil),                             // 16: authentication.Device
	(*ConfigVersion)(nil),                      // 17: authentication.ConfigVersion
}
var file_rpc_proto_depIdxs = []int32{
	5,  // 0: rpc.LongPollingPayload.data:type_name -> rpc.IncomingRPCMessage
	15, // 1: rpc.LongPollingPayload.heartbeat:type_name -> util.EmptyArr
	3,  // 2: rpc.LongPollingPayload.ack:type_name -> rpc.StartAckMessage
	15, // 3: rpc.LongPollingPayload.startRead:type_name -> util.EmptyArr
	0,  // 4: rpc.IncomingRPCMessage.bugleRoute:type_name -> rpc.BugleRoute
	2,  // 5: rpc.IncomingRPCMessage.messageType:type_name -> rpc.MessageType
	16, // 6: rpc.IncomingRPCMessage.mobile:type_name -> authentication.Device
	16, // 7: rpc.IncomingRPCMessage.browser:type_name -> authentication.Device
	10, // 8: rpc.IncomingRPCMessage.gdittoSource:type_name -> rpc.IncomingRPCMessage.GDittoSource
	1,  // 9: rpc.RPCMessageData.action:type_name -> rpc.ActionType
	16, // 10: rpc.OutgoingRPCMessage.mobile:type_name -> authentication.Device
	12, // 11: rpc.OutgoingRPCMessage.data:type_name -> rpc.OutgoingRPCMessage.Data
	11, // 12: rpc.OutgoingRPCMessage.auth:type_name -> rpc.OutgoingRPCMessage.Auth
	1,  // 13: rpc.OutgoingRPCData.action:type_name -> rpc.ActionType
	14, // 14: rpc.OutgoingRPCResponse.someIdentifier:type_name -> rpc.OutgoingRPCResponse.SomeIdentifier
	17, // 15: rpc.OutgoingRPCMessage.Auth.configVersion:type_name -> authentication.ConfigVersion
	0,  // 16: rpc.OutgoingRPCMessage.Data.bugleRoute:type_name -> rpc.BugleRoute
	13, // 17: rpc.OutgoingRPCMessage.Data.messageTypeData:type_name -> rpc.OutgoingRPCMessage.Data.Type
	15, // 18: rpc.OutgoingRPCMessage.Data.Type.emptyArr:type_name -> util.EmptyArr
	2,  // 19: rpc.OutgoingRPCMessage.Data.Type.messageType:type_name -> rpc.MessageType
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_authentication_proto_init()
	file_util_proto_init()
	file_rpc_proto_msgTypes[0].OneofWrappers = []any{}
	file_rpc_proto_msgTypes[1].OneofWrappers = []any{}
	file_rpc_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		EnumInfos:         file_rpc_proto_enumTypes,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
