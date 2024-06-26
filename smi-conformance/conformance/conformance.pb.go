// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: learn-khulnasoft/smi-conformance/conformance/conformance.proto

package conformance

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	service "github.com/khulnasoft/service-mesh-performance/service"
	spec "github.com/khulnasoft/service-mesh-performance/spec"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Capability int32

const (
	Capability_FULL Capability = 0
	Capability_HALF Capability = 1
	Capability_NONE Capability = 2
)

// Enum value maps for Capability.
var (
	Capability_name = map[int32]string{
		0: "FULL",
		1: "HALF",
		2: "NONE",
	}
	Capability_value = map[string]int32{
		"FULL": 0,
		"HALF": 1,
		"NONE": 2,
	}
)

func (x Capability) Enum() *Capability {
	p := new(Capability)
	*p = x
	return p
}

func (x Capability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Capability) Descriptor() protoreflect.EnumDescriptor {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes[0].Descriptor()
}

func (Capability) Type() protoreflect.EnumType {
	return &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes[0]
}

func (x Capability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Capability.Descriptor instead.
func (Capability) EnumDescriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

type TestStatus int32

const (
	TestStatus_COMPLETED  TestStatus = 0
	TestStatus_INPROGRESS TestStatus = 1
	TestStatus_CRASHED    TestStatus = 2
)

// Enum value maps for TestStatus.
var (
	TestStatus_name = map[int32]string{
		0: "COMPLETED",
		1: "INPROGRESS",
		2: "CRASHED",
	}
	TestStatus_value = map[string]int32{
		"COMPLETED":  0,
		"INPROGRESS": 1,
		"CRASHED":    2,
	}
)

func (x TestStatus) Enum() *TestStatus {
	p := new(TestStatus)
	*p = x
	return p
}

func (x TestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes[1].Descriptor()
}

func (TestStatus) Type() protoreflect.EnumType {
	return &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes[1]
}

func (x TestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestStatus.Descriptor instead.
func (TestStatus) EnumDescriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

type ResultStatus int32

const (
	ResultStatus_PASSED ResultStatus = 0
	ResultStatus_FAILED ResultStatus = 1
)

// Enum value maps for ResultStatus.
var (
	ResultStatus_name = map[int32]string{
		0: "PASSED",
		1: "FAILED",
	}
	ResultStatus_value = map[string]int32{
		"PASSED": 0,
		"FAILED": 1,
	}
)

func (x ResultStatus) Enum() *ResultStatus {
	p := new(ResultStatus)
	*p = x
	return p
}

func (x ResultStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes[2].Descriptor()
}

func (ResultStatus) Type() protoreflect.EnumType {
	return &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes[2]
}

func (x ResultStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultStatus.Descriptor instead.
func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{2}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mesh *spec.ServiceMesh `protobuf:"bytes,1,opt,name=mesh,proto3" json:"mesh,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMesh() *spec.ServiceMesh {
	if x != nil {
		return x.Mesh
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*Result_Message
	//	*Result_Error
	Result isResult_Result `protobuf_oneof:"result"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

func (m *Result) GetResult() isResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *Result) GetMessage() string {
	if x, ok := x.GetResult().(*Result_Message); ok {
		return x.Message
	}
	return ""
}

func (x *Result) GetError() *service.CommonError {
	if x, ok := x.GetResult().(*Result_Error); ok {
		return x.Error
	}
	return nil
}

type isResult_Result interface {
	isResult_Result()
}

type Result_Message struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

type Result_Error struct {
	Error *service.CommonError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*Result_Message) isResult_Result() {}

func (*Result_Error) isResult_Result() {}

type Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Smispec     string       `protobuf:"bytes,1,opt,name=smispec,proto3" json:"smispec,omitempty"`
	Specversion string       `protobuf:"bytes,2,opt,name=specversion,proto3" json:"specversion,omitempty"`
	Assertion   string       `protobuf:"bytes,3,opt,name=assertion,proto3" json:"assertion,omitempty"`
	Duration    string       `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Result      *Result      `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	Capability  Capability   `protobuf:"varint,6,opt,name=capability,proto3,enum=smi_conformance.Capability" json:"capability,omitempty"`
	Status      ResultStatus `protobuf:"varint,7,opt,name=status,proto3,enum=smi_conformance.ResultStatus" json:"status,omitempty"`
}

func (x *Detail) Reset() {
	*x = Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detail) ProtoMessage() {}

func (x *Detail) ProtoReflect() protoreflect.Message {
	mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detail.ProtoReflect.Descriptor instead.
func (*Detail) Descriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{2}
}

func (x *Detail) GetSmispec() string {
	if x != nil {
		return x.Smispec
	}
	return ""
}

func (x *Detail) GetSpecversion() string {
	if x != nil {
		return x.Specversion
	}
	return ""
}

func (x *Detail) GetAssertion() string {
	if x != nil {
		return x.Assertion
	}
	return ""
}

func (x *Detail) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Detail) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Detail) GetCapability() Capability {
	if x != nil {
		return x.Capability
	}
	return Capability_FULL
}

func (x *Detail) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_PASSED
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passpercent string            `protobuf:"bytes,1,opt,name=passpercent,proto3" json:"passpercent,omitempty"`
	Casespassed string            `protobuf:"bytes,2,opt,name=casespassed,proto3" json:"casespassed,omitempty"`
	Mesh        *spec.ServiceMesh `protobuf:"bytes,3,opt,name=mesh,proto3" json:"mesh,omitempty"`
	Details     []*Detail         `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetPasspercent() string {
	if x != nil {
		return x.Passpercent
	}
	return ""
}

func (x *Response) GetCasespassed() string {
	if x != nil {
		return x.Casespassed
	}
	return ""
}

func (x *Response) GetMesh() *spec.ServiceMesh {
	if x != nil {
		return x.Mesh
	}
	return nil
}

func (x *Response) GetDetails() []*Detail {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_learn_khulnasoft_smi_conformance_conformance_conformance_proto protoreflect.FileDescriptor

var file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2d, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x35, 0x2f, 0x73,
	0x6d, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6d,
	0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2d, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x6d, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x73, 0x68, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x22, 0x5c, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa3, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6d, 0x69, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x69, 0x73, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x70, 0x65, 0x63, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa7, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x73, 0x73, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x73, 0x65, 0x73, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x73, 0x65, 0x73, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x24,
	0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x6d, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x04,
	0x6d, 0x65, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x2a, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x02, 0x2a, 0x38, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x41, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x26, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0xc4, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x38, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x07,
	0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescOnce sync.Once
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescData = file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDesc
)

func file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescGZIP() []byte {
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescOnce.Do(func() {
		file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescData)
	})
	return file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDescData
}

var file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_goTypes = []interface{}{
	(Capability)(0),               // 0: smi_conformance.Capability
	(TestStatus)(0),               // 1: smi_conformance.TestStatus
	(ResultStatus)(0),             // 2: smi_conformance.ResultStatus
	(*Request)(nil),               // 3: smi_conformance.Request
	(*Result)(nil),                // 4: smi_conformance.Result
	(*Detail)(nil),                // 5: smi_conformance.Detail
	(*Response)(nil),              // 6: smi_conformance.Response
	(*spec.ServiceMesh)(nil),      // 7: smp.ServiceMesh
	(*service.CommonError)(nil),   // 8: service.CommonError
	(*empty.Empty)(nil),           // 9: google.protobuf.Empty
	(*service.ServiceInfo)(nil),   // 10: service.ServiceInfo
	(*service.ServiceHealth)(nil), // 11: service.ServiceHealth
}
var file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_depIdxs = []int32{
	7,  // 0: smi_conformance.Request.mesh:type_name -> smp.ServiceMesh
	8,  // 1: smi_conformance.Result.error:type_name -> service.CommonError
	4,  // 2: smi_conformance.Detail.result:type_name -> smi_conformance.Result
	0,  // 3: smi_conformance.Detail.capability:type_name -> smi_conformance.Capability
	2,  // 4: smi_conformance.Detail.status:type_name -> smi_conformance.ResultStatus
	7,  // 5: smi_conformance.Response.mesh:type_name -> smp.ServiceMesh
	5,  // 6: smi_conformance.Response.details:type_name -> smi_conformance.Detail
	9,  // 7: smi_conformance.conformanceTesting.Info:input_type -> google.protobuf.Empty
	9,  // 8: smi_conformance.conformanceTesting.Health:input_type -> google.protobuf.Empty
	3,  // 9: smi_conformance.conformanceTesting.RunTest:input_type -> smi_conformance.Request
	10, // 10: smi_conformance.conformanceTesting.Info:output_type -> service.ServiceInfo
	11, // 11: smi_conformance.conformanceTesting.Health:output_type -> service.ServiceHealth
	6,  // 12: smi_conformance.conformanceTesting.RunTest:output_type -> smi_conformance.Response
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_init() }
func file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_init() {
	if File_learn_khulnasoft_smi_conformance_conformance_conformance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detail); i {
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
		file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Result_Message)(nil),
		(*Result_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_goTypes,
		DependencyIndexes: file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_depIdxs,
		EnumInfos:         file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_enumTypes,
		MessageInfos:      file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_msgTypes,
	}.Build()
	File_learn_khulnasoft_smi_conformance_conformance_conformance_proto = out.File
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_rawDesc = nil
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_goTypes = nil
	file_learn_khulnasoft_smi_conformance_conformance_conformance_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConformanceTestingClient is the client API for ConformanceTesting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConformanceTestingClient interface {
	Info(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*service.ServiceInfo, error)
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*service.ServiceHealth, error)
	RunTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type conformanceTestingClient struct {
	cc grpc.ClientConnInterface
}

func NewConformanceTestingClient(cc grpc.ClientConnInterface) ConformanceTestingClient {
	return &conformanceTestingClient{cc}
}

func (c *conformanceTestingClient) Info(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*service.ServiceInfo, error) {
	out := new(service.ServiceInfo)
	err := c.cc.Invoke(ctx, "/smi_conformance.conformanceTesting/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conformanceTestingClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*service.ServiceHealth, error) {
	out := new(service.ServiceHealth)
	err := c.cc.Invoke(ctx, "/smi_conformance.conformanceTesting/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conformanceTestingClient) RunTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/smi_conformance.conformanceTesting/RunTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConformanceTestingServer is the server API for ConformanceTesting service.
type ConformanceTestingServer interface {
	Info(context.Context, *empty.Empty) (*service.ServiceInfo, error)
	Health(context.Context, *empty.Empty) (*service.ServiceHealth, error)
	RunTest(context.Context, *Request) (*Response, error)
}

// UnimplementedConformanceTestingServer can be embedded to have forward compatible implementations.
type UnimplementedConformanceTestingServer struct {
}

func (*UnimplementedConformanceTestingServer) Info(context.Context, *empty.Empty) (*service.ServiceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedConformanceTestingServer) Health(context.Context, *empty.Empty) (*service.ServiceHealth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedConformanceTestingServer) RunTest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTest not implemented")
}

func RegisterConformanceTestingServer(s *grpc.Server, srv ConformanceTestingServer) {
	s.RegisterService(&_ConformanceTesting_serviceDesc, srv)
}

func _ConformanceTesting_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConformanceTestingServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smi_conformance.conformanceTesting/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConformanceTestingServer).Info(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConformanceTesting_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConformanceTestingServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smi_conformance.conformanceTesting/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConformanceTestingServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConformanceTesting_RunTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConformanceTestingServer).RunTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smi_conformance.conformanceTesting/RunTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConformanceTestingServer).RunTest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConformanceTesting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smi_conformance.conformanceTesting",
	HandlerType: (*ConformanceTestingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _ConformanceTesting_Info_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _ConformanceTesting_Health_Handler,
		},
		{
			MethodName: "RunTest",
			Handler:    _ConformanceTesting_RunTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "learn-khulnasoft/smi-conformance/conformance/conformance.proto",
}
