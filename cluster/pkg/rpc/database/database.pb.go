// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: database.proto

package database

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// transaction contains our units of work.
type TransactionMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver  string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount    int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	SessionId int64  `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *TransactionMsg) Reset() {
	*x = TransactionMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionMsg) ProtoMessage() {}

func (x *TransactionMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionMsg.ProtoReflect.Descriptor instead.
func (*TransactionMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionMsg) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *TransactionMsg) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *TransactionMsg) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionMsg) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

// request is used for inter-shard transactions.
type RequestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction   *TransactionMsg `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ReturnAddress string          `protobuf:"bytes,2,opt,name=return_address,json=returnAddress,proto3" json:"return_address,omitempty"`
}

func (x *RequestMsg) Reset() {
	*x = RequestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMsg) ProtoMessage() {}

func (x *RequestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMsg.ProtoReflect.Descriptor instead.
func (*RequestMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{1}
}

func (x *RequestMsg) GetTransaction() *TransactionMsg {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *RequestMsg) GetReturnAddress() string {
	if x != nil {
		return x.ReturnAddress
	}
	return ""
}

// reply is used for returning a response to the client.
type ReplyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int64  `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ReplyMsg) Reset() {
	*x = ReplyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMsg) ProtoMessage() {}

func (x *ReplyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMsg.ProtoReflect.Descriptor instead.
func (*ReplyMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyMsg) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *ReplyMsg) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// prepare is used in the first phase of 2pc.
type PrepareMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction   *TransactionMsg `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client        string          `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	ReturnAddress string          `protobuf:"bytes,3,opt,name=return_address,json=returnAddress,proto3" json:"return_address,omitempty"`
}

func (x *PrepareMsg) Reset() {
	*x = PrepareMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareMsg) ProtoMessage() {}

func (x *PrepareMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareMsg.ProtoReflect.Descriptor instead.
func (*PrepareMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{3}
}

func (x *PrepareMsg) GetTransaction() *TransactionMsg {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PrepareMsg) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *PrepareMsg) GetReturnAddress() string {
	if x != nil {
		return x.ReturnAddress
	}
	return ""
}

// ack is returned by nodes for 2pc.
type AckMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int64  `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	IsAborted bool   `protobuf:"varint,2,opt,name=is_aborted,json=isAborted,proto3" json:"is_aborted,omitempty"`
	NodeId    string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *AckMsg) Reset() {
	*x = AckMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMsg) ProtoMessage() {}

func (x *AckMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMsg.ProtoReflect.Descriptor instead.
func (*AckMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{4}
}

func (x *AckMsg) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *AckMsg) GetIsAborted() bool {
	if x != nil {
		return x.IsAborted
	}
	return false
}

func (x *AckMsg) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

// commit is used in the second phase of 2pc if no node aborts prepare.
type CommitMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     int64  `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ReturnAddress string `protobuf:"bytes,2,opt,name=return_address,json=returnAddress,proto3" json:"return_address,omitempty"`
}

func (x *CommitMsg) Reset() {
	*x = CommitMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitMsg) ProtoMessage() {}

func (x *CommitMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitMsg.ProtoReflect.Descriptor instead.
func (*CommitMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{5}
}

func (x *CommitMsg) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CommitMsg) GetReturnAddress() string {
	if x != nil {
		return x.ReturnAddress
	}
	return ""
}

// abort is used in the second phase of 2pc if at least one node aborts prepare.
type AbortMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *AbortMsg) Reset() {
	*x = AbortMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortMsg) ProtoMessage() {}

func (x *AbortMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortMsg.ProtoReflect.Descriptor instead.
func (*AbortMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{6}
}

func (x *AbortMsg) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

// rebalance is used in the rebalancing phase.
type RebalanceMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record  string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Balance int64  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Add     bool   `protobuf:"varint,3,opt,name=add,proto3" json:"add,omitempty"`
}

func (x *RebalanceMsg) Reset() {
	*x = RebalanceMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebalanceMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebalanceMsg) ProtoMessage() {}

func (x *RebalanceMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebalanceMsg.ProtoReflect.Descriptor instead.
func (*RebalanceMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{7}
}

func (x *RebalanceMsg) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *RebalanceMsg) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *RebalanceMsg) GetAdd() bool {
	if x != nil {
		return x.Add
	}
	return false
}

// rebalance response is used in rebalancing phase.
type RebalanceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record  string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Balance int64  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *RebalanceRsp) Reset() {
	*x = RebalanceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebalanceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebalanceRsp) ProtoMessage() {}

func (x *RebalanceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebalanceRsp.ProtoReflect.Descriptor instead.
func (*RebalanceRsp) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{8}
}

func (x *RebalanceRsp) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *RebalanceRsp) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

// client helper functions.
type PrintBalanceMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *PrintBalanceMsg) Reset() {
	*x = PrintBalanceMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintBalanceMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintBalanceMsg) ProtoMessage() {}

func (x *PrintBalanceMsg) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintBalanceMsg.ProtoReflect.Descriptor instead.
func (*PrintBalanceMsg) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{9}
}

func (x *PrintBalanceMsg) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

type PrintBalanceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client  string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Balance int64  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *PrintBalanceRsp) Reset() {
	*x = PrintBalanceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintBalanceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintBalanceRsp) ProtoMessage() {}

func (x *PrintBalanceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintBalanceRsp.ProtoReflect.Descriptor instead.
func (*PrintBalanceRsp) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{10}
}

func (x *PrintBalanceRsp) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *PrintBalanceRsp) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type LogRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record    string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	SessionId int64  `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	NewValue  int64  `protobuf:"varint,4,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
}

func (x *LogRsp) Reset() {
	*x = LogRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRsp) ProtoMessage() {}

func (x *LogRsp) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRsp.ProtoReflect.Descriptor instead.
func (*LogRsp) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{11}
}

func (x *LogRsp) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *LogRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogRsp) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *LogRsp) GetNewValue() int64 {
	if x != nil {
		return x.NewValue
	}
	return 0
}

type DatastoreRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BallotNumberPid      string `protobuf:"bytes,1,opt,name=ballot_number_pid,json=ballotNumberPid,proto3" json:"ballot_number_pid,omitempty"`
	BallotNumberSequence int64  `protobuf:"varint,2,opt,name=ballot_number_sequence,json=ballotNumberSequence,proto3" json:"ballot_number_sequence,omitempty"`
	Record               string `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
	SessionId            int64  `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DatastoreRsp) Reset() {
	*x = DatastoreRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatastoreRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatastoreRsp) ProtoMessage() {}

func (x *DatastoreRsp) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatastoreRsp.ProtoReflect.Descriptor instead.
func (*DatastoreRsp) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{12}
}

func (x *DatastoreRsp) GetBallotNumberPid() string {
	if x != nil {
		return x.BallotNumberPid
	}
	return ""
}

func (x *DatastoreRsp) GetBallotNumberSequence() int64 {
	if x != nil {
		return x.BallotNumberSequence
	}
	return 0
}

func (x *DatastoreRsp) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *DatastoreRsp) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

var File_database_proto protoreflect.FileDescriptor

var file_database_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x73, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73,
	0x67, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x73,
	0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x4d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x73, 0x67, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5f,
	0x0a, 0x06, 0x41, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x61, 0x62,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22,
	0x51, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x29, 0x0a, 0x08, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x52, 0x0a,
	0x0c, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x64,
	0x64, 0x22, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x43,
	0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x76, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x52, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0c,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x11,
	0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x62, 0x61, 0x6c, 0x6c,
	0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xda, 0x05, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12,
	0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x41, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x73,
	0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x41, 0x62, 0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x19, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x07, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_proto_rawDescOnce sync.Once
	file_database_proto_rawDescData = file_database_proto_rawDesc
)

func file_database_proto_rawDescGZIP() []byte {
	file_database_proto_rawDescOnce.Do(func() {
		file_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_proto_rawDescData)
	})
	return file_database_proto_rawDescData
}

var file_database_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_database_proto_goTypes = []any{
	(*TransactionMsg)(nil),  // 0: database.TransactionMsg
	(*RequestMsg)(nil),      // 1: database.RequestMsg
	(*ReplyMsg)(nil),        // 2: database.ReplyMsg
	(*PrepareMsg)(nil),      // 3: database.PrepareMsg
	(*AckMsg)(nil),          // 4: database.AckMsg
	(*CommitMsg)(nil),       // 5: database.CommitMsg
	(*AbortMsg)(nil),        // 6: database.AbortMsg
	(*RebalanceMsg)(nil),    // 7: database.RebalanceMsg
	(*RebalanceRsp)(nil),    // 8: database.RebalanceRsp
	(*PrintBalanceMsg)(nil), // 9: database.PrintBalanceMsg
	(*PrintBalanceRsp)(nil), // 10: database.PrintBalanceRsp
	(*LogRsp)(nil),          // 11: database.LogRsp
	(*DatastoreRsp)(nil),    // 12: database.DatastoreRsp
	(*emptypb.Empty)(nil),   // 13: google.protobuf.Empty
}
var file_database_proto_depIdxs = []int32{
	0,  // 0: database.RequestMsg.transaction:type_name -> database.TransactionMsg
	0,  // 1: database.PrepareMsg.transaction:type_name -> database.TransactionMsg
	1,  // 2: database.Database.Request:input_type -> database.RequestMsg
	2,  // 3: database.Database.Reply:input_type -> database.ReplyMsg
	3,  // 4: database.Database.Prepare:input_type -> database.PrepareMsg
	4,  // 5: database.Database.Ack:input_type -> database.AckMsg
	5,  // 6: database.Database.Commit:input_type -> database.CommitMsg
	6,  // 7: database.Database.Abort:input_type -> database.AbortMsg
	7,  // 8: database.Database.Rebalance:input_type -> database.RebalanceMsg
	9,  // 9: database.Database.PrintBalance:input_type -> database.PrintBalanceMsg
	13, // 10: database.Database.PrintLogs:input_type -> google.protobuf.Empty
	13, // 11: database.Database.PrintDatastore:input_type -> google.protobuf.Empty
	13, // 12: database.Database.Block:input_type -> google.protobuf.Empty
	13, // 13: database.Database.Unblock:input_type -> google.protobuf.Empty
	13, // 14: database.Database.Request:output_type -> google.protobuf.Empty
	13, // 15: database.Database.Reply:output_type -> google.protobuf.Empty
	13, // 16: database.Database.Prepare:output_type -> google.protobuf.Empty
	13, // 17: database.Database.Ack:output_type -> google.protobuf.Empty
	13, // 18: database.Database.Commit:output_type -> google.protobuf.Empty
	13, // 19: database.Database.Abort:output_type -> google.protobuf.Empty
	8,  // 20: database.Database.Rebalance:output_type -> database.RebalanceRsp
	10, // 21: database.Database.PrintBalance:output_type -> database.PrintBalanceRsp
	11, // 22: database.Database.PrintLogs:output_type -> database.LogRsp
	12, // 23: database.Database.PrintDatastore:output_type -> database.DatastoreRsp
	13, // 24: database.Database.Block:output_type -> google.protobuf.Empty
	13, // 25: database.Database.Unblock:output_type -> google.protobuf.Empty
	14, // [14:26] is the sub-list for method output_type
	2,  // [2:14] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_database_proto_init() }
func file_database_proto_init() {
	if File_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_database_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionMsg); i {
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
		file_database_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RequestMsg); i {
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
		file_database_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReplyMsg); i {
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
		file_database_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PrepareMsg); i {
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
		file_database_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AckMsg); i {
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
		file_database_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CommitMsg); i {
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
		file_database_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AbortMsg); i {
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
		file_database_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RebalanceMsg); i {
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
		file_database_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RebalanceRsp); i {
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
		file_database_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*PrintBalanceMsg); i {
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
		file_database_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*PrintBalanceRsp); i {
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
		file_database_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*LogRsp); i {
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
		file_database_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*DatastoreRsp); i {
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
			RawDescriptor: file_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_database_proto_goTypes,
		DependencyIndexes: file_database_proto_depIdxs,
		MessageInfos:      file_database_proto_msgTypes,
	}.Build()
	File_database_proto = out.File
	file_database_proto_rawDesc = nil
	file_database_proto_goTypes = nil
	file_database_proto_depIdxs = nil
}
