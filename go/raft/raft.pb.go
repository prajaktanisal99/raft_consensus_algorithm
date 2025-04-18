// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/raft.proto

package raft

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestVoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Term          int32                  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId   int32                  `protobuf:"varint,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	mi := &file_proto_raft_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetCandidateId() int32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Term          int32                  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted   bool                   `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted,proto3" json:"vote_granted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	mi := &file_proto_raft_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{1}
}

func (x *RequestVoteResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type LogEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operation     string                 `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Term          int32                  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Index         int32                  `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	mi := &file_proto_raft_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{2}
}

func (x *LogEntry) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *LogEntry) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Term          int32                  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId      int32                  `protobuf:"varint,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Entries       []*LogEntry            `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	CommitIndex   int32                  `protobuf:"varint,4,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppendEntriesRequest) Reset() {
	*x = AppendEntriesRequest{}
	mi := &file_proto_raft_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppendEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesRequest) ProtoMessage() {}

func (x *AppendEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesRequest.ProtoReflect.Descriptor instead.
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{3}
}

func (x *AppendEntriesRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesRequest) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *AppendEntriesRequest) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntriesRequest) GetCommitIndex() int32 {
	if x != nil {
		return x.CommitIndex
	}
	return 0
}

type AppendEntriesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Term          int32                  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppendEntriesResponse) Reset() {
	*x = AppendEntriesResponse{}
	mi := &file_proto_raft_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppendEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResponse) ProtoMessage() {}

func (x *AppendEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResponse.ProtoReflect.Descriptor instead.
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{4}
}

func (x *AppendEntriesResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_raft_proto protoreflect.FileDescriptor

const file_proto_raft_proto_rawDesc = "" +
	"\n" +
	"\x10proto/raft.proto\x12\x04raft\"K\n" +
	"\x12RequestVoteRequest\x12\x12\n" +
	"\x04term\x18\x01 \x01(\x05R\x04term\x12!\n" +
	"\fcandidate_id\x18\x02 \x01(\x05R\vcandidateId\"L\n" +
	"\x13RequestVoteResponse\x12\x12\n" +
	"\x04term\x18\x01 \x01(\x05R\x04term\x12!\n" +
	"\fvote_granted\x18\x02 \x01(\bR\vvoteGranted\"R\n" +
	"\bLogEntry\x12\x1c\n" +
	"\toperation\x18\x01 \x01(\tR\toperation\x12\x12\n" +
	"\x04term\x18\x02 \x01(\x05R\x04term\x12\x14\n" +
	"\x05index\x18\x03 \x01(\x05R\x05index\"\x94\x01\n" +
	"\x14AppendEntriesRequest\x12\x12\n" +
	"\x04term\x18\x01 \x01(\x05R\x04term\x12\x1b\n" +
	"\tleader_id\x18\x02 \x01(\x05R\bleaderId\x12(\n" +
	"\aentries\x18\x03 \x03(\v2\x0e.raft.LogEntryR\aentries\x12!\n" +
	"\fcommit_index\x18\x04 \x01(\x05R\vcommitIndex\"E\n" +
	"\x15AppendEntriesResponse\x12\x12\n" +
	"\x04term\x18\x01 \x01(\x05R\x04term\x12\x18\n" +
	"\asuccess\x18\x02 \x01(\bR\asuccess2\x94\x01\n" +
	"\x04Raft\x12B\n" +
	"\vRequestVote\x12\x18.raft.RequestVoteRequest\x1a\x19.raft.RequestVoteResponse\x12H\n" +
	"\rAppendEntries\x12\x1a.raft.AppendEntriesRequest\x1a\x1b.raft.AppendEntriesResponseB\aZ\x05/raftb\x06proto3"

var (
	file_proto_raft_proto_rawDescOnce sync.Once
	file_proto_raft_proto_rawDescData []byte
)

func file_proto_raft_proto_rawDescGZIP() []byte {
	file_proto_raft_proto_rawDescOnce.Do(func() {
		file_proto_raft_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_raft_proto_rawDesc), len(file_proto_raft_proto_rawDesc)))
	})
	return file_proto_raft_proto_rawDescData
}

var file_proto_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_raft_proto_goTypes = []any{
	(*RequestVoteRequest)(nil),    // 0: raft.RequestVoteRequest
	(*RequestVoteResponse)(nil),   // 1: raft.RequestVoteResponse
	(*LogEntry)(nil),              // 2: raft.LogEntry
	(*AppendEntriesRequest)(nil),  // 3: raft.AppendEntriesRequest
	(*AppendEntriesResponse)(nil), // 4: raft.AppendEntriesResponse
}
var file_proto_raft_proto_depIdxs = []int32{
	2, // 0: raft.AppendEntriesRequest.entries:type_name -> raft.LogEntry
	0, // 1: raft.Raft.RequestVote:input_type -> raft.RequestVoteRequest
	3, // 2: raft.Raft.AppendEntries:input_type -> raft.AppendEntriesRequest
	1, // 3: raft.Raft.RequestVote:output_type -> raft.RequestVoteResponse
	4, // 4: raft.Raft.AppendEntries:output_type -> raft.AppendEntriesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_raft_proto_init() }
func file_proto_raft_proto_init() {
	if File_proto_raft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_raft_proto_rawDesc), len(file_proto_raft_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raft_proto_goTypes,
		DependencyIndexes: file_proto_raft_proto_depIdxs,
		MessageInfos:      file_proto_raft_proto_msgTypes,
	}.Build()
	File_proto_raft_proto = out.File
	file_proto_raft_proto_goTypes = nil
	file_proto_raft_proto_depIdxs = nil
}
