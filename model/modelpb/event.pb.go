// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.3
// source: event.proto

package modelpb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	Outcome      string                 `protobuf:"bytes,1,opt,name=outcome,proto3" json:"outcome,omitempty"`
	Action       string                 `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Dataset      string                 `protobuf:"bytes,3,opt,name=dataset,proto3" json:"dataset,omitempty"`
	Kind         string                 `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Category     string                 `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Type         string                 `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	SuccessCount *SummaryMetric         `protobuf:"bytes,7,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	// nanoseconds
	Duration uint64 `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	Severity uint64 `protobuf:"varint,9,opt,name=severity,proto3" json:"severity,omitempty"`
	// nanoseconds since epoch
	Received      uint64 `protobuf:"varint,10,opt,name=received,proto3" json:"received,omitempty"`
	Module        string `protobuf:"bytes,11,opt,name=module,proto3" json:"module,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if x != nil {
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
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetOutcome() string {
	if x != nil {
		return x.Outcome
	}
	return ""
}

func (x *Event) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Event) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *Event) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Event) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetSuccessCount() *SummaryMetric {
	if x != nil {
		return x.SuccessCount
	}
	return nil
}

func (x *Event) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Event) GetSeverity() uint64 {
	if x != nil {
		return x.Severity
	}
	return 0
}

func (x *Event) GetReceived() uint64 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *Event) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

var File_event_proto protoreflect.FileDescriptor

const file_event_proto_rawDesc = "" +
	"\n" +
	"\vevent.proto\x12\x0eelastic.apm.v1\x1a\x0fmetricset.proto\"\xc7\x02\n" +
	"\x05Event\x12\x18\n" +
	"\aoutcome\x18\x01 \x01(\tR\aoutcome\x12\x16\n" +
	"\x06action\x18\x02 \x01(\tR\x06action\x12\x18\n" +
	"\adataset\x18\x03 \x01(\tR\adataset\x12\x12\n" +
	"\x04kind\x18\x04 \x01(\tR\x04kind\x12\x1a\n" +
	"\bcategory\x18\x05 \x01(\tR\bcategory\x12\x12\n" +
	"\x04type\x18\x06 \x01(\tR\x04type\x12B\n" +
	"\rsuccess_count\x18\a \x01(\v2\x1d.elastic.apm.v1.SummaryMetricR\fsuccessCount\x12\x1a\n" +
	"\bduration\x18\b \x01(\x04R\bduration\x12\x1a\n" +
	"\bseverity\x18\t \x01(\x04R\bseverity\x12\x1a\n" +
	"\breceived\x18\n" +
	" \x01(\x04R\breceived\x12\x16\n" +
	"\x06module\x18\v \x01(\tR\x06moduleB+Z)github.com/elastic/apm-data/model/modelpbb\x06proto3"

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData []byte
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_event_proto_rawDesc), len(file_event_proto_rawDesc)))
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_proto_goTypes = []any{
	(*Event)(nil),         // 0: elastic.apm.v1.Event
	(*SummaryMetric)(nil), // 1: elastic.apm.v1.SummaryMetric
}
var file_event_proto_depIdxs = []int32{
	1, // 0: elastic.apm.v1.Event.success_count:type_name -> elastic.apm.v1.SummaryMetric
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_metricset_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_event_proto_rawDesc), len(file_event_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
