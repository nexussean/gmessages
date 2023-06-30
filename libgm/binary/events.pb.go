// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: events.proto

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*Event_ConversationEvent
	//	*Event_MessageEvent
	//	*Event_UserAlertEvent
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Event) GetConversationEvent() *ConversationEvent {
	if x, ok := x.GetEvent().(*Event_ConversationEvent); ok {
		return x.ConversationEvent
	}
	return nil
}

func (x *Event) GetMessageEvent() *MessageEvent {
	if x, ok := x.GetEvent().(*Event_MessageEvent); ok {
		return x.MessageEvent
	}
	return nil
}

func (x *Event) GetUserAlertEvent() *UserAlertEvent {
	if x, ok := x.GetEvent().(*Event_UserAlertEvent); ok {
		return x.UserAlertEvent
	}
	return nil
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_ConversationEvent struct {
	ConversationEvent *ConversationEvent `protobuf:"bytes,2,opt,name=conversationEvent,proto3,oneof"`
}

type Event_MessageEvent struct {
	MessageEvent *MessageEvent `protobuf:"bytes,3,opt,name=messageEvent,proto3,oneof"`
}

type Event_UserAlertEvent struct {
	UserAlertEvent *UserAlertEvent `protobuf:"bytes,6,opt,name=userAlertEvent,proto3,oneof"`
}

func (*Event_ConversationEvent) isEvent_Event() {}

func (*Event_MessageEvent) isEvent_Event() {}

func (*Event_UserAlertEvent) isEvent_Event() {}

type ConversationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Conversation `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConversationEvent) Reset() {
	*x = ConversationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationEvent) ProtoMessage() {}

func (x *ConversationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationEvent.ProtoReflect.Descriptor instead.
func (*ConversationEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *ConversationEvent) GetData() *Conversation {
	if x != nil {
		return x.Data
	}
	return nil
}

type MessageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Message `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MessageEvent) Reset() {
	*x = MessageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageEvent) ProtoMessage() {}

func (x *MessageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageEvent.ProtoReflect.Descriptor instead.
func (*MessageEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *MessageEvent) GetData() *Message {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserAlertEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 2 = BROWSER_ACTIVE (new session created in other browser)
	//
	// 1 = ?
	//
	// 8 = INACTIVE_LACK_OF_ACTIVITY
	//
	// 7 = INACTIVE_TIMEOUT
	//
	// 5|6 = BATTERY (tf?)
	//
	// 3|4 = DATA_CONNECTION (tf?)
	//
	// 10 = OBSERVER_REGISTERED (tf?)
	AlertType int64 `protobuf:"varint,2,opt,name=alert_type,json=alertType,proto3" json:"alert_type,omitempty"`
}

func (x *UserAlertEvent) Reset() {
	*x = UserAlertEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAlertEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAlertEvent) ProtoMessage() {}

func (x *UserAlertEvent) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAlertEvent.ProtoReflect.Descriptor instead.
func (*UserAlertEvent) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *UserAlertEvent) GetAlertType() int64 {
	if x != nil {
		return x.AlertType
	}
	return 0
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x11, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0e,
	0x75, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e,
	0x75, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a,
	0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_events_proto_goTypes = []interface{}{
	(*Event)(nil),             // 0: events.Event
	(*ConversationEvent)(nil), // 1: events.ConversationEvent
	(*MessageEvent)(nil),      // 2: events.MessageEvent
	(*UserAlertEvent)(nil),    // 3: events.UserAlertEvent
	(*Conversation)(nil),      // 4: conversations.Conversation
	(*Message)(nil),           // 5: conversations.Message
}
var file_events_proto_depIdxs = []int32{
	1, // 0: events.Event.conversationEvent:type_name -> events.ConversationEvent
	2, // 1: events.Event.messageEvent:type_name -> events.MessageEvent
	3, // 2: events.Event.userAlertEvent:type_name -> events.UserAlertEvent
	4, // 3: events.ConversationEvent.data:type_name -> conversations.Conversation
	5, // 4: events.MessageEvent.data:type_name -> conversations.Message
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	file_conversations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationEvent); i {
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
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageEvent); i {
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
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAlertEvent); i {
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
	file_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_ConversationEvent)(nil),
		(*Event_MessageEvent)(nil),
		(*Event_UserAlertEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}
