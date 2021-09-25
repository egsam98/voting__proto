// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: voting.proto

package votingpb

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

type VoteStatus int32

const (
	VoteStatus_OK   VoteStatus = 0
	VoteStatus_FAIL VoteStatus = 1
)

// Enum value maps for VoteStatus.
var (
	VoteStatus_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	VoteStatus_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x VoteStatus) Enum() *VoteStatus {
	p := new(VoteStatus)
	*p = x
	return p
}

func (x VoteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_voting_proto_enumTypes[0].Descriptor()
}

func (VoteStatus) Type() protoreflect.EnumType {
	return &file_voting_proto_enumTypes[0]
}

func (x VoteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoteStatus.Descriptor instead.
func (VoteStatus) EnumDescriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{0}
}

type Votes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votes []*Vote `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (x *Votes) Reset() {
	*x = Votes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Votes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Votes) ProtoMessage() {}

func (x *Votes) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Votes.ProtoReflect.Descriptor instead.
func (*Votes) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{0}
}

func (x *Votes) GetVotes() []*Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Voter       *Voter     `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	CampaignId  int64      `protobuf:"varint,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	CandidateId int64      `protobuf:"varint,3,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Status      VoteStatus `protobuf:"varint,4,opt,name=status,proto3,enum=voting.VoteStatus" json:"status,omitempty"`
	FailReason  *string    `protobuf:"bytes,5,opt,name=fail_reason,json=failReason,proto3,oneof" json:"fail_reason,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{1}
}

func (x *Vote) GetVoter() *Voter {
	if x != nil {
		return x.Voter
	}
	return nil
}

func (x *Vote) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *Vote) GetCandidateId() int64 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *Vote) GetStatus() VoteStatus {
	if x != nil {
		return x.Status
	}
	return VoteStatus_OK
}

func (x *Vote) GetFailReason() string {
	if x != nil && x.FailReason != nil {
		return *x.FailReason
	}
	return ""
}

type Voter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Passport  string `protobuf:"bytes,3,opt,name=passport,proto3" json:"passport,omitempty"`
	Fullname  string `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	BirthDate int64  `protobuf:"varint,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	DeathDate *int64 `protobuf:"varint,6,opt,name=death_date,json=deathDate,proto3,oneof" json:"death_date,omitempty"`
}

func (x *Voter) Reset() {
	*x = Voter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voter) ProtoMessage() {}

func (x *Voter) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voter.ProtoReflect.Descriptor instead.
func (*Voter) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{2}
}

func (x *Voter) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Voter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Voter) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *Voter) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Voter) GetBirthDate() int64 {
	if x != nil {
		return x.BirthDate
	}
	return 0
}

func (x *Voter) GetDeathDate() int64 {
	if x != nil && x.DeathDate != nil {
		return *x.DeathDate
	}
	return 0
}

type VotePersisted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoterPersisted *VoterPersisted `protobuf:"bytes,1,opt,name=voter_persisted,json=voterPersisted,proto3" json:"voter_persisted,omitempty"`
	CampaignId     int64           `protobuf:"varint,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	CandidateId    int64           `protobuf:"varint,3,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Status         VoteStatus      `protobuf:"varint,4,opt,name=status,proto3,enum=voting.VoteStatus" json:"status,omitempty"`
	FailReason     *string         `protobuf:"bytes,5,opt,name=fail_reason,json=failReason,proto3,oneof" json:"fail_reason,omitempty"`
}

func (x *VotePersisted) Reset() {
	*x = VotePersisted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VotePersisted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VotePersisted) ProtoMessage() {}

func (x *VotePersisted) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VotePersisted.ProtoReflect.Descriptor instead.
func (*VotePersisted) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{3}
}

func (x *VotePersisted) GetVoterPersisted() *VoterPersisted {
	if x != nil {
		return x.VoterPersisted
	}
	return nil
}

func (x *VotePersisted) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *VotePersisted) GetCandidateId() int64 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *VotePersisted) GetStatus() VoteStatus {
	if x != nil {
		return x.Status
	}
	return VoteStatus_OK
}

func (x *VotePersisted) GetFailReason() string {
	if x != nil && x.FailReason != nil {
		return *x.FailReason
	}
	return ""
}

type VoterPersisted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKeyRsa []byte `protobuf:"bytes,1,opt,name=public_key_rsa,json=publicKeyRsa,proto3" json:"public_key_rsa,omitempty"`
	PassportRsa  []byte `protobuf:"bytes,2,opt,name=passport_rsa,json=passportRsa,proto3" json:"passport_rsa,omitempty"`
}

func (x *VoterPersisted) Reset() {
	*x = VoterPersisted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoterPersisted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoterPersisted) ProtoMessage() {}

func (x *VoterPersisted) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoterPersisted.ProtoReflect.Descriptor instead.
func (*VoterPersisted) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{4}
}

func (x *VoterPersisted) GetPublicKeyRsa() []byte {
	if x != nil {
		return x.PublicKeyRsa
	}
	return nil
}

func (x *VoterPersisted) GetPassportRsa() []byte {
	if x != nil {
		return x.PassportRsa
	}
	return nil
}

var File_voting_proto protoreflect.FileDescriptor

var file_voting_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2b, 0x0a, 0x05, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x05,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x24, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x05, 0x56, 0x6f, 0x74, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x65, 0x61, 0x74, 0x68,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x64,
	0x65, 0x61, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x64, 0x65, 0x61, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x0d, 0x56,
	0x6f, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0f,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x6f, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x52, 0x0e, 0x76,
	0x6f, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a,
	0x0b, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x0e, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x72, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x73, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x73, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x73, 0x61, 0x2a, 0x1e,
	0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x73,
	0x61, 0x6d, 0x39, 0x38, 0x2f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_voting_proto_rawDescOnce sync.Once
	file_voting_proto_rawDescData = file_voting_proto_rawDesc
)

func file_voting_proto_rawDescGZIP() []byte {
	file_voting_proto_rawDescOnce.Do(func() {
		file_voting_proto_rawDescData = protoimpl.X.CompressGZIP(file_voting_proto_rawDescData)
	})
	return file_voting_proto_rawDescData
}

var file_voting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_voting_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_voting_proto_goTypes = []interface{}{
	(VoteStatus)(0),        // 0: voting.VoteStatus
	(*Votes)(nil),          // 1: voting.Votes
	(*Vote)(nil),           // 2: voting.Vote
	(*Voter)(nil),          // 3: voting.Voter
	(*VotePersisted)(nil),  // 4: voting.VotePersisted
	(*VoterPersisted)(nil), // 5: voting.VoterPersisted
}
var file_voting_proto_depIdxs = []int32{
	2, // 0: voting.Votes.votes:type_name -> voting.Vote
	3, // 1: voting.Vote.voter:type_name -> voting.Voter
	0, // 2: voting.Vote.status:type_name -> voting.VoteStatus
	5, // 3: voting.VotePersisted.voter_persisted:type_name -> voting.VoterPersisted
	0, // 4: voting.VotePersisted.status:type_name -> voting.VoteStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_voting_proto_init() }
func file_voting_proto_init() {
	if File_voting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Votes); i {
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
		file_voting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_voting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voter); i {
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
		file_voting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VotePersisted); i {
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
		file_voting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoterPersisted); i {
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
	file_voting_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_voting_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_voting_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_voting_proto_goTypes,
		DependencyIndexes: file_voting_proto_depIdxs,
		EnumInfos:         file_voting_proto_enumTypes,
		MessageInfos:      file_voting_proto_msgTypes,
	}.Build()
	File_voting_proto = out.File
	file_voting_proto_rawDesc = nil
	file_voting_proto_goTypes = nil
	file_voting_proto_depIdxs = nil
}
