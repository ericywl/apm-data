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
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: apmevent.proto

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

type APMEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// nanoseconds since epoch
	Timestamp     uint64                        `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Span          *Span                         `protobuf:"bytes,2,opt,name=span,proto3" json:"span,omitempty"`
	NumericLabels map[string]*NumericLabelValue `protobuf:"bytes,3,rep,name=numeric_labels,json=numericLabels,proto3" json:"numeric_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Labels        map[string]*LabelValue        `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Transaction   *Transaction                  `protobuf:"bytes,5,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Metricset     *Metricset                    `protobuf:"bytes,6,opt,name=metricset,proto3" json:"metricset,omitempty"`
	Error         *Error                        `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	Cloud         *Cloud                        `protobuf:"bytes,8,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Service       *Service                      `protobuf:"bytes,9,opt,name=service,proto3" json:"service,omitempty"`
	Faas          *Faas                         `protobuf:"bytes,10,opt,name=faas,proto3" json:"faas,omitempty"`
	Network       *Network                      `protobuf:"bytes,11,opt,name=network,proto3" json:"network,omitempty"`
	Container     *Container                    `protobuf:"bytes,12,opt,name=container,proto3" json:"container,omitempty"`
	User          *User                         `protobuf:"bytes,13,opt,name=user,proto3" json:"user,omitempty"`
	Device        *Device                       `protobuf:"bytes,14,opt,name=device,proto3" json:"device,omitempty"`
	Kubernetes    *Kubernetes                   `protobuf:"bytes,15,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Observer      *Observer                     `protobuf:"bytes,16,opt,name=observer,proto3" json:"observer,omitempty"`
	DataStream    *DataStream                   `protobuf:"bytes,17,opt,name=data_stream,json=dataStream,proto3" json:"data_stream,omitempty"`
	Agent         *Agent                        `protobuf:"bytes,18,opt,name=agent,proto3" json:"agent,omitempty"`
	Http          *HTTP                         `protobuf:"bytes,19,opt,name=http,proto3" json:"http,omitempty"`
	UserAgent     *UserAgent                    `protobuf:"bytes,20,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// parent_id holds an optional parent span/transaction ID.
	ParentId string  `protobuf:"bytes,21,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Message  string  `protobuf:"bytes,22,opt,name=message,proto3" json:"message,omitempty"`
	Trace    *Trace  `protobuf:"bytes,23,opt,name=trace,proto3" json:"trace,omitempty"`
	Host     *Host   `protobuf:"bytes,24,opt,name=host,proto3" json:"host,omitempty"`
	Url      *URL    `protobuf:"bytes,25,opt,name=url,proto3" json:"url,omitempty"`
	Log      *Log    `protobuf:"bytes,26,opt,name=log,proto3" json:"log,omitempty"`
	Source   *Source `protobuf:"bytes,27,opt,name=source,proto3" json:"source,omitempty"`
	Client   *Client `protobuf:"bytes,28,opt,name=client,proto3" json:"client,omitempty"`
	// child_ids holds an optional set of child span IDs. This is used for exotic
	// use cases where the parent knows the child ID, but the child does not know
	// the parent ID; namely for profiler-inferred spans.
	ChildIds      []string     `protobuf:"bytes,29,rep,name=child_ids,json=childIds,proto3" json:"child_ids,omitempty"`
	Destination   *Destination `protobuf:"bytes,30,opt,name=destination,proto3" json:"destination,omitempty"`
	Session       *Session     `protobuf:"bytes,31,opt,name=session,proto3" json:"session,omitempty"`
	Process       *Process     `protobuf:"bytes,32,opt,name=process,proto3" json:"process,omitempty"`
	Event         *Event       `protobuf:"bytes,33,opt,name=event,proto3" json:"event,omitempty"`
	Code          *Code        `protobuf:"bytes,34,opt,name=code,proto3" json:"code,omitempty"`
	System        *System      `protobuf:"bytes,35,opt,name=system,proto3" json:"system,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *APMEvent) Reset() {
	*x = APMEvent{}
	mi := &file_apmevent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APMEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APMEvent) ProtoMessage() {}

func (x *APMEvent) ProtoReflect() protoreflect.Message {
	mi := &file_apmevent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APMEvent.ProtoReflect.Descriptor instead.
func (*APMEvent) Descriptor() ([]byte, []int) {
	return file_apmevent_proto_rawDescGZIP(), []int{0}
}

func (x *APMEvent) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *APMEvent) GetSpan() *Span {
	if x != nil {
		return x.Span
	}
	return nil
}

func (x *APMEvent) GetNumericLabels() map[string]*NumericLabelValue {
	if x != nil {
		return x.NumericLabels
	}
	return nil
}

func (x *APMEvent) GetLabels() map[string]*LabelValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *APMEvent) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *APMEvent) GetMetricset() *Metricset {
	if x != nil {
		return x.Metricset
	}
	return nil
}

func (x *APMEvent) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *APMEvent) GetCloud() *Cloud {
	if x != nil {
		return x.Cloud
	}
	return nil
}

func (x *APMEvent) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *APMEvent) GetFaas() *Faas {
	if x != nil {
		return x.Faas
	}
	return nil
}

func (x *APMEvent) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *APMEvent) GetContainer() *Container {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *APMEvent) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *APMEvent) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *APMEvent) GetKubernetes() *Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

func (x *APMEvent) GetObserver() *Observer {
	if x != nil {
		return x.Observer
	}
	return nil
}

func (x *APMEvent) GetDataStream() *DataStream {
	if x != nil {
		return x.DataStream
	}
	return nil
}

func (x *APMEvent) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *APMEvent) GetHttp() *HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *APMEvent) GetUserAgent() *UserAgent {
	if x != nil {
		return x.UserAgent
	}
	return nil
}

func (x *APMEvent) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *APMEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *APMEvent) GetTrace() *Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *APMEvent) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *APMEvent) GetUrl() *URL {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *APMEvent) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *APMEvent) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *APMEvent) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *APMEvent) GetChildIds() []string {
	if x != nil {
		return x.ChildIds
	}
	return nil
}

func (x *APMEvent) GetDestination() *Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *APMEvent) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *APMEvent) GetProcess() *Process {
	if x != nil {
		return x.Process
	}
	return nil
}

func (x *APMEvent) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *APMEvent) GetCode() *Code {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *APMEvent) GetSystem() *System {
	if x != nil {
		return x.System
	}
	return nil
}

var File_apmevent_proto protoreflect.FileDescriptor

var file_apmevent_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x61, 0x70, 0x6d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70,
	0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda,
	0x0e, 0x0a, 0x08, 0x41, 0x50, 0x4d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x70, 0x61,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x04, 0x73,
	0x70, 0x61, 0x6e, 0x12, 0x52, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6c,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x4d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x4d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6c, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x65,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x65, 0x74, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x65, 0x74, 0x12, 0x2b, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x66, 0x61,
	0x61, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x61, 0x73, 0x52, 0x04,
	0x66, 0x61, 0x61, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e,
	0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6c, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x08, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x12, 0x38, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e,
	0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x25, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x49, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a,
	0x63, 0x0a, 0x12, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x55, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_apmevent_proto_rawDescOnce sync.Once
	file_apmevent_proto_rawDescData []byte
)

func file_apmevent_proto_rawDescGZIP() []byte {
	file_apmevent_proto_rawDescOnce.Do(func() {
		file_apmevent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apmevent_proto_rawDesc), len(file_apmevent_proto_rawDesc)))
	})
	return file_apmevent_proto_rawDescData
}

var file_apmevent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apmevent_proto_goTypes = []any{
	(*APMEvent)(nil),          // 0: elastic.apm.v1.APMEvent
	nil,                       // 1: elastic.apm.v1.APMEvent.NumericLabelsEntry
	nil,                       // 2: elastic.apm.v1.APMEvent.LabelsEntry
	(*Span)(nil),              // 3: elastic.apm.v1.Span
	(*Transaction)(nil),       // 4: elastic.apm.v1.Transaction
	(*Metricset)(nil),         // 5: elastic.apm.v1.Metricset
	(*Error)(nil),             // 6: elastic.apm.v1.Error
	(*Cloud)(nil),             // 7: elastic.apm.v1.Cloud
	(*Service)(nil),           // 8: elastic.apm.v1.Service
	(*Faas)(nil),              // 9: elastic.apm.v1.Faas
	(*Network)(nil),           // 10: elastic.apm.v1.Network
	(*Container)(nil),         // 11: elastic.apm.v1.Container
	(*User)(nil),              // 12: elastic.apm.v1.User
	(*Device)(nil),            // 13: elastic.apm.v1.Device
	(*Kubernetes)(nil),        // 14: elastic.apm.v1.Kubernetes
	(*Observer)(nil),          // 15: elastic.apm.v1.Observer
	(*DataStream)(nil),        // 16: elastic.apm.v1.DataStream
	(*Agent)(nil),             // 17: elastic.apm.v1.Agent
	(*HTTP)(nil),              // 18: elastic.apm.v1.HTTP
	(*UserAgent)(nil),         // 19: elastic.apm.v1.UserAgent
	(*Trace)(nil),             // 20: elastic.apm.v1.Trace
	(*Host)(nil),              // 21: elastic.apm.v1.Host
	(*URL)(nil),               // 22: elastic.apm.v1.URL
	(*Log)(nil),               // 23: elastic.apm.v1.Log
	(*Source)(nil),            // 24: elastic.apm.v1.Source
	(*Client)(nil),            // 25: elastic.apm.v1.Client
	(*Destination)(nil),       // 26: elastic.apm.v1.Destination
	(*Session)(nil),           // 27: elastic.apm.v1.Session
	(*Process)(nil),           // 28: elastic.apm.v1.Process
	(*Event)(nil),             // 29: elastic.apm.v1.Event
	(*Code)(nil),              // 30: elastic.apm.v1.Code
	(*System)(nil),            // 31: elastic.apm.v1.System
	(*NumericLabelValue)(nil), // 32: elastic.apm.v1.NumericLabelValue
	(*LabelValue)(nil),        // 33: elastic.apm.v1.LabelValue
}
var file_apmevent_proto_depIdxs = []int32{
	3,  // 0: elastic.apm.v1.APMEvent.span:type_name -> elastic.apm.v1.Span
	1,  // 1: elastic.apm.v1.APMEvent.numeric_labels:type_name -> elastic.apm.v1.APMEvent.NumericLabelsEntry
	2,  // 2: elastic.apm.v1.APMEvent.labels:type_name -> elastic.apm.v1.APMEvent.LabelsEntry
	4,  // 3: elastic.apm.v1.APMEvent.transaction:type_name -> elastic.apm.v1.Transaction
	5,  // 4: elastic.apm.v1.APMEvent.metricset:type_name -> elastic.apm.v1.Metricset
	6,  // 5: elastic.apm.v1.APMEvent.error:type_name -> elastic.apm.v1.Error
	7,  // 6: elastic.apm.v1.APMEvent.cloud:type_name -> elastic.apm.v1.Cloud
	8,  // 7: elastic.apm.v1.APMEvent.service:type_name -> elastic.apm.v1.Service
	9,  // 8: elastic.apm.v1.APMEvent.faas:type_name -> elastic.apm.v1.Faas
	10, // 9: elastic.apm.v1.APMEvent.network:type_name -> elastic.apm.v1.Network
	11, // 10: elastic.apm.v1.APMEvent.container:type_name -> elastic.apm.v1.Container
	12, // 11: elastic.apm.v1.APMEvent.user:type_name -> elastic.apm.v1.User
	13, // 12: elastic.apm.v1.APMEvent.device:type_name -> elastic.apm.v1.Device
	14, // 13: elastic.apm.v1.APMEvent.kubernetes:type_name -> elastic.apm.v1.Kubernetes
	15, // 14: elastic.apm.v1.APMEvent.observer:type_name -> elastic.apm.v1.Observer
	16, // 15: elastic.apm.v1.APMEvent.data_stream:type_name -> elastic.apm.v1.DataStream
	17, // 16: elastic.apm.v1.APMEvent.agent:type_name -> elastic.apm.v1.Agent
	18, // 17: elastic.apm.v1.APMEvent.http:type_name -> elastic.apm.v1.HTTP
	19, // 18: elastic.apm.v1.APMEvent.user_agent:type_name -> elastic.apm.v1.UserAgent
	20, // 19: elastic.apm.v1.APMEvent.trace:type_name -> elastic.apm.v1.Trace
	21, // 20: elastic.apm.v1.APMEvent.host:type_name -> elastic.apm.v1.Host
	22, // 21: elastic.apm.v1.APMEvent.url:type_name -> elastic.apm.v1.URL
	23, // 22: elastic.apm.v1.APMEvent.log:type_name -> elastic.apm.v1.Log
	24, // 23: elastic.apm.v1.APMEvent.source:type_name -> elastic.apm.v1.Source
	25, // 24: elastic.apm.v1.APMEvent.client:type_name -> elastic.apm.v1.Client
	26, // 25: elastic.apm.v1.APMEvent.destination:type_name -> elastic.apm.v1.Destination
	27, // 26: elastic.apm.v1.APMEvent.session:type_name -> elastic.apm.v1.Session
	28, // 27: elastic.apm.v1.APMEvent.process:type_name -> elastic.apm.v1.Process
	29, // 28: elastic.apm.v1.APMEvent.event:type_name -> elastic.apm.v1.Event
	30, // 29: elastic.apm.v1.APMEvent.code:type_name -> elastic.apm.v1.Code
	31, // 30: elastic.apm.v1.APMEvent.system:type_name -> elastic.apm.v1.System
	32, // 31: elastic.apm.v1.APMEvent.NumericLabelsEntry.value:type_name -> elastic.apm.v1.NumericLabelValue
	33, // 32: elastic.apm.v1.APMEvent.LabelsEntry.value:type_name -> elastic.apm.v1.LabelValue
	33, // [33:33] is the sub-list for method output_type
	33, // [33:33] is the sub-list for method input_type
	33, // [33:33] is the sub-list for extension type_name
	33, // [33:33] is the sub-list for extension extendee
	0,  // [0:33] is the sub-list for field type_name
}

func init() { file_apmevent_proto_init() }
func file_apmevent_proto_init() {
	if File_apmevent_proto != nil {
		return
	}
	file_agent_proto_init()
	file_client_proto_init()
	file_cloud_proto_init()
	file_code_proto_init()
	file_container_proto_init()
	file_datastream_proto_init()
	file_destination_proto_init()
	file_device_proto_init()
	file_error_proto_init()
	file_event_proto_init()
	file_faas_proto_init()
	file_host_proto_init()
	file_http_proto_init()
	file_kubernetes_proto_init()
	file_labels_proto_init()
	file_log_proto_init()
	file_metricset_proto_init()
	file_network_proto_init()
	file_observer_proto_init()
	file_process_proto_init()
	file_service_proto_init()
	file_session_proto_init()
	file_source_proto_init()
	file_span_proto_init()
	file_system_proto_init()
	file_trace_proto_init()
	file_transaction_proto_init()
	file_url_proto_init()
	file_user_proto_init()
	file_useragent_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apmevent_proto_rawDesc), len(file_apmevent_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apmevent_proto_goTypes,
		DependencyIndexes: file_apmevent_proto_depIdxs,
		MessageInfos:      file_apmevent_proto_msgTypes,
	}.Build()
	File_apmevent_proto = out.File
	file_apmevent_proto_goTypes = nil
	file_apmevent_proto_depIdxs = nil
}
