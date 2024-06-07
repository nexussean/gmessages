// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: ukey.proto

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

type Ukey2HandshakeCipher int32

const (
	Ukey2HandshakeCipher_RESERVED          Ukey2HandshakeCipher = 0
	Ukey2HandshakeCipher_P256_SHA512       Ukey2HandshakeCipher = 100 // NIST P-256 used for ECDH, SHA512 used for commitment
	Ukey2HandshakeCipher_CURVE25519_SHA512 Ukey2HandshakeCipher = 200 // Curve 25519 used for ECDH, SHA512 used for commitment
)

// Enum value maps for Ukey2HandshakeCipher.
var (
	Ukey2HandshakeCipher_name = map[int32]string{
		0:   "RESERVED",
		100: "P256_SHA512",
		200: "CURVE25519_SHA512",
	}
	Ukey2HandshakeCipher_value = map[string]int32{
		"RESERVED":          0,
		"P256_SHA512":       100,
		"CURVE25519_SHA512": 200,
	}
)

func (x Ukey2HandshakeCipher) Enum() *Ukey2HandshakeCipher {
	p := new(Ukey2HandshakeCipher)
	*p = x
	return p
}

func (x Ukey2HandshakeCipher) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ukey2HandshakeCipher) Descriptor() protoreflect.EnumDescriptor {
	return file_ukey_proto_enumTypes[0].Descriptor()
}

func (Ukey2HandshakeCipher) Type() protoreflect.EnumType {
	return &file_ukey_proto_enumTypes[0]
}

func (x Ukey2HandshakeCipher) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ukey2HandshakeCipher.Descriptor instead.
func (Ukey2HandshakeCipher) EnumDescriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{0}
}

// A list of supported public key types
type PublicKeyType int32

const (
	PublicKeyType_UNKNOWN_PUBLIC_KEY_TYPE PublicKeyType = 0
	PublicKeyType_EC_P256                 PublicKeyType = 1
	PublicKeyType_RSA2048                 PublicKeyType = 2
	// 2048-bit MODP group 14, from RFC 3526
	PublicKeyType_DH2048_MODP PublicKeyType = 3
)

// Enum value maps for PublicKeyType.
var (
	PublicKeyType_name = map[int32]string{
		0: "UNKNOWN_PUBLIC_KEY_TYPE",
		1: "EC_P256",
		2: "RSA2048",
		3: "DH2048_MODP",
	}
	PublicKeyType_value = map[string]int32{
		"UNKNOWN_PUBLIC_KEY_TYPE": 0,
		"EC_P256":                 1,
		"RSA2048":                 2,
		"DH2048_MODP":             3,
	}
)

func (x PublicKeyType) Enum() *PublicKeyType {
	p := new(PublicKeyType)
	*p = x
	return p
}

func (x PublicKeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublicKeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_ukey_proto_enumTypes[1].Descriptor()
}

func (PublicKeyType) Type() protoreflect.EnumType {
	return &file_ukey_proto_enumTypes[1]
}

func (x PublicKeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublicKeyType.Descriptor instead.
func (PublicKeyType) EnumDescriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{1}
}

type Ukey2Message_Type int32

const (
	Ukey2Message_UNKNOWN_DO_NOT_USE Ukey2Message_Type = 0
	Ukey2Message_ALERT              Ukey2Message_Type = 1
	Ukey2Message_CLIENT_INIT        Ukey2Message_Type = 2
	Ukey2Message_SERVER_INIT        Ukey2Message_Type = 3
	Ukey2Message_CLIENT_FINISH      Ukey2Message_Type = 4
)

// Enum value maps for Ukey2Message_Type.
var (
	Ukey2Message_Type_name = map[int32]string{
		0: "UNKNOWN_DO_NOT_USE",
		1: "ALERT",
		2: "CLIENT_INIT",
		3: "SERVER_INIT",
		4: "CLIENT_FINISH",
	}
	Ukey2Message_Type_value = map[string]int32{
		"UNKNOWN_DO_NOT_USE": 0,
		"ALERT":              1,
		"CLIENT_INIT":        2,
		"SERVER_INIT":        3,
		"CLIENT_FINISH":      4,
	}
)

func (x Ukey2Message_Type) Enum() *Ukey2Message_Type {
	p := new(Ukey2Message_Type)
	*p = x
	return p
}

func (x Ukey2Message_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ukey2Message_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_ukey_proto_enumTypes[2].Descriptor()
}

func (Ukey2Message_Type) Type() protoreflect.EnumType {
	return &file_ukey_proto_enumTypes[2]
}

func (x Ukey2Message_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ukey2Message_Type.Descriptor instead.
func (Ukey2Message_Type) EnumDescriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{0, 0}
}

type Ukey2Alert_AlertType int32

const (
	Ukey2Alert_UNKNOWN_ALERT_TYPE Ukey2Alert_AlertType = 0
	// Framing errors
	Ukey2Alert_BAD_MESSAGE       Ukey2Alert_AlertType = 1 // The message could not be deserialized
	Ukey2Alert_BAD_MESSAGE_TYPE  Ukey2Alert_AlertType = 2 // message_type has an undefined value
	Ukey2Alert_INCORRECT_MESSAGE Ukey2Alert_AlertType = 3 // message_type received does not correspond to expected type at this stage of the protocol
	Ukey2Alert_BAD_MESSAGE_DATA  Ukey2Alert_AlertType = 4 // Could not deserialize message_data as per value in message_type
	// ClientInit and ServerInit errors
	Ukey2Alert_BAD_VERSION          Ukey2Alert_AlertType = 100 // version is invalid; server cannot find suitable version to speak with client.
	Ukey2Alert_BAD_RANDOM           Ukey2Alert_AlertType = 101 // Random data is missing or of incorrect length
	Ukey2Alert_BAD_HANDSHAKE_CIPHER Ukey2Alert_AlertType = 102 // No suitable handshake ciphers were found
	Ukey2Alert_BAD_NEXT_PROTOCOL    Ukey2Alert_AlertType = 103 // The next protocol is missing, unknown, or unsupported
	Ukey2Alert_BAD_PUBLIC_KEY       Ukey2Alert_AlertType = 104 // The public key could not be parsed
	// Other errors
	Ukey2Alert_INTERNAL_ERROR Ukey2Alert_AlertType = 200 // An internal error has occurred.  error_message may contain additional details for logging and debugging.
)

// Enum value maps for Ukey2Alert_AlertType.
var (
	Ukey2Alert_AlertType_name = map[int32]string{
		0:   "UNKNOWN_ALERT_TYPE",
		1:   "BAD_MESSAGE",
		2:   "BAD_MESSAGE_TYPE",
		3:   "INCORRECT_MESSAGE",
		4:   "BAD_MESSAGE_DATA",
		100: "BAD_VERSION",
		101: "BAD_RANDOM",
		102: "BAD_HANDSHAKE_CIPHER",
		103: "BAD_NEXT_PROTOCOL",
		104: "BAD_PUBLIC_KEY",
		200: "INTERNAL_ERROR",
	}
	Ukey2Alert_AlertType_value = map[string]int32{
		"UNKNOWN_ALERT_TYPE":   0,
		"BAD_MESSAGE":          1,
		"BAD_MESSAGE_TYPE":     2,
		"INCORRECT_MESSAGE":    3,
		"BAD_MESSAGE_DATA":     4,
		"BAD_VERSION":          100,
		"BAD_RANDOM":           101,
		"BAD_HANDSHAKE_CIPHER": 102,
		"BAD_NEXT_PROTOCOL":    103,
		"BAD_PUBLIC_KEY":       104,
		"INTERNAL_ERROR":       200,
	}
)

func (x Ukey2Alert_AlertType) Enum() *Ukey2Alert_AlertType {
	p := new(Ukey2Alert_AlertType)
	*p = x
	return p
}

func (x Ukey2Alert_AlertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ukey2Alert_AlertType) Descriptor() protoreflect.EnumDescriptor {
	return file_ukey_proto_enumTypes[3].Descriptor()
}

func (Ukey2Alert_AlertType) Type() protoreflect.EnumType {
	return &file_ukey_proto_enumTypes[3]
}

func (x Ukey2Alert_AlertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ukey2Alert_AlertType.Descriptor instead.
func (Ukey2Alert_AlertType) EnumDescriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{1, 0}
}

type Ukey2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType Ukey2Message_Type `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=ukey.Ukey2Message_Type" json:"message_type,omitempty"` // Identifies message type
	MessageData []byte            `protobuf:"bytes,2,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`                              // Actual message, to be parsed according to message_type
}

func (x *Ukey2Message) Reset() {
	*x = Ukey2Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ukey2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ukey2Message) ProtoMessage() {}

func (x *Ukey2Message) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ukey2Message.ProtoReflect.Descriptor instead.
func (*Ukey2Message) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{0}
}

func (x *Ukey2Message) GetMessageType() Ukey2Message_Type {
	if x != nil {
		return x.MessageType
	}
	return Ukey2Message_UNKNOWN_DO_NOT_USE
}

func (x *Ukey2Message) GetMessageData() []byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

type Ukey2Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         Ukey2Alert_AlertType `protobuf:"varint,1,opt,name=type,proto3,enum=ukey.Ukey2Alert_AlertType" json:"type,omitempty"`
	ErrorMessage string               `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *Ukey2Alert) Reset() {
	*x = Ukey2Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ukey2Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ukey2Alert) ProtoMessage() {}

func (x *Ukey2Alert) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ukey2Alert.ProtoReflect.Descriptor instead.
func (*Ukey2Alert) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{1}
}

func (x *Ukey2Alert) GetType() Ukey2Alert_AlertType {
	if x != nil {
		return x.Type
	}
	return Ukey2Alert_UNKNOWN_ALERT_TYPE
}

func (x *Ukey2Alert) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type Ukey2ClientInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           int32                               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // highest supported version for rollback protection
	Random            []byte                              `protobuf:"bytes,2,opt,name=random,proto3" json:"random,omitempty"`    // random bytes for replay/reuse protection
	CipherCommitments []*Ukey2ClientInit_CipherCommitment `protobuf:"bytes,3,rep,name=cipher_commitments,json=cipherCommitments,proto3" json:"cipher_commitments,omitempty"`
	// Next protocol that the client wants to speak.
	NextProtocol string `protobuf:"bytes,4,opt,name=next_protocol,json=nextProtocol,proto3" json:"next_protocol,omitempty"`
}

func (x *Ukey2ClientInit) Reset() {
	*x = Ukey2ClientInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ukey2ClientInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ukey2ClientInit) ProtoMessage() {}

func (x *Ukey2ClientInit) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ukey2ClientInit.ProtoReflect.Descriptor instead.
func (*Ukey2ClientInit) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{2}
}

func (x *Ukey2ClientInit) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ukey2ClientInit) GetRandom() []byte {
	if x != nil {
		return x.Random
	}
	return nil
}

func (x *Ukey2ClientInit) GetCipherCommitments() []*Ukey2ClientInit_CipherCommitment {
	if x != nil {
		return x.CipherCommitments
	}
	return nil
}

func (x *Ukey2ClientInit) GetNextProtocol() string {
	if x != nil {
		return x.NextProtocol
	}
	return ""
}

type Ukey2ServerInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"` // highest supported version for rollback protection
	Random  []byte `protobuf:"bytes,2,opt,name=random,proto3" json:"random,omitempty"`    // random bytes for replay/reuse protection
	// Selected Cipher and corresponding public key
	HandshakeCipher Ukey2HandshakeCipher `protobuf:"varint,3,opt,name=handshake_cipher,json=handshakeCipher,proto3,enum=ukey.Ukey2HandshakeCipher" json:"handshake_cipher,omitempty"`
	PublicKey       *GenericPublicKey    `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Ukey2ServerInit) Reset() {
	*x = Ukey2ServerInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ukey2ServerInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ukey2ServerInit) ProtoMessage() {}

func (x *Ukey2ServerInit) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ukey2ServerInit.ProtoReflect.Descriptor instead.
func (*Ukey2ServerInit) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{3}
}

func (x *Ukey2ServerInit) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ukey2ServerInit) GetRandom() []byte {
	if x != nil {
		return x.Random
	}
	return nil
}

func (x *Ukey2ServerInit) GetHandshakeCipher() Ukey2HandshakeCipher {
	if x != nil {
		return x.HandshakeCipher
	}
	return Ukey2HandshakeCipher_RESERVED
}

func (x *Ukey2ServerInit) GetPublicKey() *GenericPublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type Ukey2ClientFinished struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey *GenericPublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"` // public key matching selected handshake cipher
}

func (x *Ukey2ClientFinished) Reset() {
	*x = Ukey2ClientFinished{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ukey2ClientFinished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ukey2ClientFinished) ProtoMessage() {}

func (x *Ukey2ClientFinished) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ukey2ClientFinished.ProtoReflect.Descriptor instead.
func (*Ukey2ClientFinished) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{4}
}

func (x *Ukey2ClientFinished) GetPublicKey() *GenericPublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// A convenience proto for encoding NIST P-256 elliptic curve public keys
type EcP256PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// x and y are encoded in big-endian two's complement (slightly wasteful)
	// Client MUST verify (x,y) is a valid point on NIST P256
	X []byte `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y []byte `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *EcP256PublicKey) Reset() {
	*x = EcP256PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcP256PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcP256PublicKey) ProtoMessage() {}

func (x *EcP256PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcP256PublicKey.ProtoReflect.Descriptor instead.
func (*EcP256PublicKey) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{5}
}

func (x *EcP256PublicKey) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *EcP256PublicKey) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

// A convenience proto for encoding RSA public keys with small exponents
type SimpleRsaPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encoded in big-endian two's complement
	N []byte `protobuf:"bytes,1,opt,name=n,proto3" json:"n,omitempty"`
	E int32  `protobuf:"varint,2,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *SimpleRsaPublicKey) Reset() {
	*x = SimpleRsaPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRsaPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRsaPublicKey) ProtoMessage() {}

func (x *SimpleRsaPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRsaPublicKey.ProtoReflect.Descriptor instead.
func (*SimpleRsaPublicKey) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{6}
}

func (x *SimpleRsaPublicKey) GetN() []byte {
	if x != nil {
		return x.N
	}
	return nil
}

func (x *SimpleRsaPublicKey) GetE() int32 {
	if x != nil {
		return x.E
	}
	return 0
}

// A convenience proto for encoding Diffie-Hellman public keys,
// for use only when Elliptic Curve based key exchanges are not possible.
// (Note that the group parameters must be specified separately)
type DhPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Big-endian two's complement encoded group element
	Y []byte `protobuf:"bytes,1,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *DhPublicKey) Reset() {
	*x = DhPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DhPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DhPublicKey) ProtoMessage() {}

func (x *DhPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DhPublicKey.ProtoReflect.Descriptor instead.
func (*DhPublicKey) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{7}
}

func (x *DhPublicKey) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

type GenericPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PublicKeyType `protobuf:"varint,1,opt,name=type,proto3,enum=ukey.PublicKeyType" json:"type,omitempty"`
	// Types that are assignable to PublicKey:
	//
	//	*GenericPublicKey_EcP256PublicKey
	//	*GenericPublicKey_Rsa2048PublicKey
	//	*GenericPublicKey_Dh2048PublicKey
	PublicKey isGenericPublicKey_PublicKey `protobuf_oneof:"public_key"`
}

func (x *GenericPublicKey) Reset() {
	*x = GenericPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPublicKey) ProtoMessage() {}

func (x *GenericPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPublicKey.ProtoReflect.Descriptor instead.
func (*GenericPublicKey) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{8}
}

func (x *GenericPublicKey) GetType() PublicKeyType {
	if x != nil {
		return x.Type
	}
	return PublicKeyType_UNKNOWN_PUBLIC_KEY_TYPE
}

func (m *GenericPublicKey) GetPublicKey() isGenericPublicKey_PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (x *GenericPublicKey) GetEcP256PublicKey() *EcP256PublicKey {
	if x, ok := x.GetPublicKey().(*GenericPublicKey_EcP256PublicKey); ok {
		return x.EcP256PublicKey
	}
	return nil
}

func (x *GenericPublicKey) GetRsa2048PublicKey() *SimpleRsaPublicKey {
	if x, ok := x.GetPublicKey().(*GenericPublicKey_Rsa2048PublicKey); ok {
		return x.Rsa2048PublicKey
	}
	return nil
}

func (x *GenericPublicKey) GetDh2048PublicKey() *DhPublicKey {
	if x, ok := x.GetPublicKey().(*GenericPublicKey_Dh2048PublicKey); ok {
		return x.Dh2048PublicKey
	}
	return nil
}

type isGenericPublicKey_PublicKey interface {
	isGenericPublicKey_PublicKey()
}

type GenericPublicKey_EcP256PublicKey struct {
	EcP256PublicKey *EcP256PublicKey `protobuf:"bytes,2,opt,name=ec_p256_public_key,json=ecP256PublicKey,proto3,oneof"`
}

type GenericPublicKey_Rsa2048PublicKey struct {
	Rsa2048PublicKey *SimpleRsaPublicKey `protobuf:"bytes,3,opt,name=rsa2048_public_key,json=rsa2048PublicKey,proto3,oneof"`
}

type GenericPublicKey_Dh2048PublicKey struct {
	// Use only as a last resort
	Dh2048PublicKey *DhPublicKey `protobuf:"bytes,4,opt,name=dh2048_public_key,json=dh2048PublicKey,proto3,oneof"`
}

func (*GenericPublicKey_EcP256PublicKey) isGenericPublicKey_PublicKey() {}

func (*GenericPublicKey_Rsa2048PublicKey) isGenericPublicKey_PublicKey() {}

func (*GenericPublicKey_Dh2048PublicKey) isGenericPublicKey_PublicKey() {}

// One commitment (hash of ClientFinished containing public key) per supported cipher
type Ukey2ClientInit_CipherCommitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HandshakeCipher Ukey2HandshakeCipher `protobuf:"varint,1,opt,name=handshake_cipher,json=handshakeCipher,proto3,enum=ukey.Ukey2HandshakeCipher" json:"handshake_cipher,omitempty"`
	Commitment      []byte               `protobuf:"bytes,2,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (x *Ukey2ClientInit_CipherCommitment) Reset() {
	*x = Ukey2ClientInit_CipherCommitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ukey_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ukey2ClientInit_CipherCommitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ukey2ClientInit_CipherCommitment) ProtoMessage() {}

func (x *Ukey2ClientInit_CipherCommitment) ProtoReflect() protoreflect.Message {
	mi := &file_ukey_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ukey2ClientInit_CipherCommitment.ProtoReflect.Descriptor instead.
func (*Ukey2ClientInit_CipherCommitment) Descriptor() ([]byte, []int) {
	return file_ukey_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Ukey2ClientInit_CipherCommitment) GetHandshakeCipher() Ukey2HandshakeCipher {
	if x != nil {
		return x.HandshakeCipher
	}
	return Ukey2HandshakeCipher_RESERVED
}

func (x *Ukey2ClientInit_CipherCommitment) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

var File_ukey_proto protoreflect.FileDescriptor

//go:embed ukey.pb.raw
var file_ukey_proto_rawDesc []byte

var (
	file_ukey_proto_rawDescOnce sync.Once
	file_ukey_proto_rawDescData = file_ukey_proto_rawDesc
)

func file_ukey_proto_rawDescGZIP() []byte {
	file_ukey_proto_rawDescOnce.Do(func() {
		file_ukey_proto_rawDescData = protoimpl.X.CompressGZIP(file_ukey_proto_rawDescData)
	})
	return file_ukey_proto_rawDescData
}

var file_ukey_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_ukey_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ukey_proto_goTypes = []interface{}{
	(Ukey2HandshakeCipher)(0),                // 0: ukey.Ukey2HandshakeCipher
	(PublicKeyType)(0),                       // 1: ukey.PublicKeyType
	(Ukey2Message_Type)(0),                   // 2: ukey.Ukey2Message.Type
	(Ukey2Alert_AlertType)(0),                // 3: ukey.Ukey2Alert.AlertType
	(*Ukey2Message)(nil),                     // 4: ukey.Ukey2Message
	(*Ukey2Alert)(nil),                       // 5: ukey.Ukey2Alert
	(*Ukey2ClientInit)(nil),                  // 6: ukey.Ukey2ClientInit
	(*Ukey2ServerInit)(nil),                  // 7: ukey.Ukey2ServerInit
	(*Ukey2ClientFinished)(nil),              // 8: ukey.Ukey2ClientFinished
	(*EcP256PublicKey)(nil),                  // 9: ukey.EcP256PublicKey
	(*SimpleRsaPublicKey)(nil),               // 10: ukey.SimpleRsaPublicKey
	(*DhPublicKey)(nil),                      // 11: ukey.DhPublicKey
	(*GenericPublicKey)(nil),                 // 12: ukey.GenericPublicKey
	(*Ukey2ClientInit_CipherCommitment)(nil), // 13: ukey.Ukey2ClientInit.CipherCommitment
}
var file_ukey_proto_depIdxs = []int32{
	2,  // 0: ukey.Ukey2Message.message_type:type_name -> ukey.Ukey2Message.Type
	3,  // 1: ukey.Ukey2Alert.type:type_name -> ukey.Ukey2Alert.AlertType
	13, // 2: ukey.Ukey2ClientInit.cipher_commitments:type_name -> ukey.Ukey2ClientInit.CipherCommitment
	0,  // 3: ukey.Ukey2ServerInit.handshake_cipher:type_name -> ukey.Ukey2HandshakeCipher
	12, // 4: ukey.Ukey2ServerInit.public_key:type_name -> ukey.GenericPublicKey
	12, // 5: ukey.Ukey2ClientFinished.public_key:type_name -> ukey.GenericPublicKey
	1,  // 6: ukey.GenericPublicKey.type:type_name -> ukey.PublicKeyType
	9,  // 7: ukey.GenericPublicKey.ec_p256_public_key:type_name -> ukey.EcP256PublicKey
	10, // 8: ukey.GenericPublicKey.rsa2048_public_key:type_name -> ukey.SimpleRsaPublicKey
	11, // 9: ukey.GenericPublicKey.dh2048_public_key:type_name -> ukey.DhPublicKey
	0,  // 10: ukey.Ukey2ClientInit.CipherCommitment.handshake_cipher:type_name -> ukey.Ukey2HandshakeCipher
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_ukey_proto_init() }
func file_ukey_proto_init() {
	if File_ukey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ukey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ukey2Message); i {
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
		file_ukey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ukey2Alert); i {
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
		file_ukey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ukey2ClientInit); i {
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
		file_ukey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ukey2ServerInit); i {
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
		file_ukey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ukey2ClientFinished); i {
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
		file_ukey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcP256PublicKey); i {
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
		file_ukey_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRsaPublicKey); i {
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
		file_ukey_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DhPublicKey); i {
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
		file_ukey_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericPublicKey); i {
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
		file_ukey_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ukey2ClientInit_CipherCommitment); i {
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
	file_ukey_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*GenericPublicKey_EcP256PublicKey)(nil),
		(*GenericPublicKey_Rsa2048PublicKey)(nil),
		(*GenericPublicKey_Dh2048PublicKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ukey_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ukey_proto_goTypes,
		DependencyIndexes: file_ukey_proto_depIdxs,
		EnumInfos:         file_ukey_proto_enumTypes,
		MessageInfos:      file_ukey_proto_msgTypes,
	}.Build()
	File_ukey_proto = out.File
	file_ukey_proto_rawDesc = nil
	file_ukey_proto_goTypes = nil
	file_ukey_proto_depIdxs = nil
}
