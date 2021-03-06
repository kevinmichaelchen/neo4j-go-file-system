// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventRequest struct {
	// Types that are valid to be assigned to Event:
	//	*EventRequest_CreateFileEvent
	//	*EventRequest_UpdateFileEvent
	//	*EventRequest_DeleteFileEvent
	//	*EventRequest_CreateFolderEvent
	//	*EventRequest_UpdateFolderEvent
	//	*EventRequest_DeleteFolderEvent
	Event                isEventRequest_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_74678b33cf7b6863, []int{0}
}
func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (dst *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(dst, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

type isEventRequest_Event interface {
	isEventRequest_Event()
}

type EventRequest_CreateFileEvent struct {
	CreateFileEvent *CreateFileRequest `protobuf:"bytes,1,opt,name=createFileEvent,proto3,oneof"`
}

type EventRequest_UpdateFileEvent struct {
	UpdateFileEvent *UpdateFileRequest `protobuf:"bytes,2,opt,name=updateFileEvent,proto3,oneof"`
}

type EventRequest_DeleteFileEvent struct {
	DeleteFileEvent *DeleteFileRequest `protobuf:"bytes,3,opt,name=deleteFileEvent,proto3,oneof"`
}

type EventRequest_CreateFolderEvent struct {
	CreateFolderEvent *CreateFolderRequest `protobuf:"bytes,4,opt,name=createFolderEvent,proto3,oneof"`
}

type EventRequest_UpdateFolderEvent struct {
	UpdateFolderEvent *UpdateFolderRequest `protobuf:"bytes,5,opt,name=updateFolderEvent,proto3,oneof"`
}

type EventRequest_DeleteFolderEvent struct {
	DeleteFolderEvent *DeleteFolderRequest `protobuf:"bytes,6,opt,name=deleteFolderEvent,proto3,oneof"`
}

func (*EventRequest_CreateFileEvent) isEventRequest_Event() {}

func (*EventRequest_UpdateFileEvent) isEventRequest_Event() {}

func (*EventRequest_DeleteFileEvent) isEventRequest_Event() {}

func (*EventRequest_CreateFolderEvent) isEventRequest_Event() {}

func (*EventRequest_UpdateFolderEvent) isEventRequest_Event() {}

func (*EventRequest_DeleteFolderEvent) isEventRequest_Event() {}

func (m *EventRequest) GetEvent() isEventRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventRequest) GetCreateFileEvent() *CreateFileRequest {
	if x, ok := m.GetEvent().(*EventRequest_CreateFileEvent); ok {
		return x.CreateFileEvent
	}
	return nil
}

func (m *EventRequest) GetUpdateFileEvent() *UpdateFileRequest {
	if x, ok := m.GetEvent().(*EventRequest_UpdateFileEvent); ok {
		return x.UpdateFileEvent
	}
	return nil
}

func (m *EventRequest) GetDeleteFileEvent() *DeleteFileRequest {
	if x, ok := m.GetEvent().(*EventRequest_DeleteFileEvent); ok {
		return x.DeleteFileEvent
	}
	return nil
}

func (m *EventRequest) GetCreateFolderEvent() *CreateFolderRequest {
	if x, ok := m.GetEvent().(*EventRequest_CreateFolderEvent); ok {
		return x.CreateFolderEvent
	}
	return nil
}

func (m *EventRequest) GetUpdateFolderEvent() *UpdateFolderRequest {
	if x, ok := m.GetEvent().(*EventRequest_UpdateFolderEvent); ok {
		return x.UpdateFolderEvent
	}
	return nil
}

func (m *EventRequest) GetDeleteFolderEvent() *DeleteFolderRequest {
	if x, ok := m.GetEvent().(*EventRequest_DeleteFolderEvent); ok {
		return x.DeleteFolderEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EventRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EventRequest_OneofMarshaler, _EventRequest_OneofUnmarshaler, _EventRequest_OneofSizer, []interface{}{
		(*EventRequest_CreateFileEvent)(nil),
		(*EventRequest_UpdateFileEvent)(nil),
		(*EventRequest_DeleteFileEvent)(nil),
		(*EventRequest_CreateFolderEvent)(nil),
		(*EventRequest_UpdateFolderEvent)(nil),
		(*EventRequest_DeleteFolderEvent)(nil),
	}
}

func _EventRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EventRequest)
	// event
	switch x := m.Event.(type) {
	case *EventRequest_CreateFileEvent:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateFileEvent); err != nil {
			return err
		}
	case *EventRequest_UpdateFileEvent:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateFileEvent); err != nil {
			return err
		}
	case *EventRequest_DeleteFileEvent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeleteFileEvent); err != nil {
			return err
		}
	case *EventRequest_CreateFolderEvent:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateFolderEvent); err != nil {
			return err
		}
	case *EventRequest_UpdateFolderEvent:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateFolderEvent); err != nil {
			return err
		}
	case *EventRequest_DeleteFolderEvent:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeleteFolderEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EventRequest.Event has unexpected type %T", x)
	}
	return nil
}

func _EventRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EventRequest)
	switch tag {
	case 1: // event.createFileEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateFileRequest)
		err := b.DecodeMessage(msg)
		m.Event = &EventRequest_CreateFileEvent{msg}
		return true, err
	case 2: // event.updateFileEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateFileRequest)
		err := b.DecodeMessage(msg)
		m.Event = &EventRequest_UpdateFileEvent{msg}
		return true, err
	case 3: // event.deleteFileEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteFileRequest)
		err := b.DecodeMessage(msg)
		m.Event = &EventRequest_DeleteFileEvent{msg}
		return true, err
	case 4: // event.createFolderEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateFolderRequest)
		err := b.DecodeMessage(msg)
		m.Event = &EventRequest_CreateFolderEvent{msg}
		return true, err
	case 5: // event.updateFolderEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateFolderRequest)
		err := b.DecodeMessage(msg)
		m.Event = &EventRequest_UpdateFolderEvent{msg}
		return true, err
	case 6: // event.deleteFolderEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteFolderRequest)
		err := b.DecodeMessage(msg)
		m.Event = &EventRequest_DeleteFolderEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EventRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EventRequest)
	// event
	switch x := m.Event.(type) {
	case *EventRequest_CreateFileEvent:
		s := proto.Size(x.CreateFileEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventRequest_UpdateFileEvent:
		s := proto.Size(x.UpdateFileEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventRequest_DeleteFileEvent:
		s := proto.Size(x.DeleteFileEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventRequest_CreateFolderEvent:
		s := proto.Size(x.CreateFolderEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventRequest_UpdateFolderEvent:
		s := proto.Size(x.UpdateFolderEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventRequest_DeleteFolderEvent:
		s := proto.Size(x.DeleteFolderEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type EventResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_74678b33cf7b6863, []int{1}
}
func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (dst *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(dst, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*EventRequest)(nil), "pb.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "pb.EventResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	EmitEvent(ctx context.Context, opts ...grpc.CallOption) (EventService_EmitEventClient, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) EmitEvent(ctx context.Context, opts ...grpc.CallOption) (EventService_EmitEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventService_serviceDesc.Streams[0], "/pb.EventService/EmitEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceEmitEventClient{stream}
	return x, nil
}

type EventService_EmitEventClient interface {
	Send(*EventRequest) error
	Recv() (*EventResponse, error)
	grpc.ClientStream
}

type eventServiceEmitEventClient struct {
	grpc.ClientStream
}

func (x *eventServiceEmitEventClient) Send(m *EventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServiceEmitEventClient) Recv() (*EventResponse, error) {
	m := new(EventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	EmitEvent(EventService_EmitEventServer) error
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_EmitEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).EmitEvent(&eventServiceEmitEventServer{stream})
}

type EventService_EmitEventServer interface {
	Send(*EventResponse) error
	Recv() (*EventRequest, error)
	grpc.ServerStream
}

type eventServiceEmitEventServer struct {
	grpc.ServerStream
}

func (x *eventServiceEmitEventServer) Send(m *EventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServiceEmitEventServer) Recv() (*EventRequest, error) {
	m := new(EventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EmitEvent",
			Handler:       _EventService_EmitEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_event_74678b33cf7b6863) }

var fileDescriptor_event_74678b33cf7b6863 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd2, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x07, 0x70, 0x5a, 0x04, 0x75, 0xc0, 0x0f, 0x36, 0x31, 0x9a, 0x5e, 0x34, 0x3d, 0x71, 0x6a,
	0x0c, 0x26, 0xde, 0x55, 0x40, 0xcf, 0x35, 0x3c, 0x80, 0xa5, 0x43, 0xd2, 0xb0, 0xb2, 0xeb, 0x76,
	0xe9, 0xcb, 0xf9, 0x72, 0x66, 0x3f, 0xd2, 0x1d, 0xdb, 0xe3, 0xf6, 0x3f, 0xf3, 0xeb, 0x6c, 0x66,
	0x61, 0x82, 0x0d, 0x1e, 0x74, 0x26, 0x95, 0xd0, 0x82, 0xc5, 0xb2, 0x48, 0x60, 0x57, 0x71, 0x74,
	0xe7, 0x64, 0xba, 0x13, 0xbc, 0x44, 0xe5, 0x4e, 0xe9, 0xef, 0x10, 0xa6, 0x2b, 0x53, 0x9d, 0xe3,
	0xcf, 0x11, 0x6b, 0xcd, 0x5e, 0xe0, 0x6a, 0xab, 0xf0, 0x4b, 0xe3, 0xba, 0xe2, 0x68, 0x93, 0xbb,
	0xe8, 0x21, 0x9a, 0x4f, 0x16, 0x37, 0x99, 0x2c, 0xb2, 0xb7, 0x36, 0xf2, 0xf5, 0x1f, 0x83, 0xbc,
	0x5b, 0x6f, 0x88, 0xa3, 0x2c, 0xff, 0x11, 0x71, 0x20, 0x36, 0x6d, 0x44, 0x88, 0x4e, 0xbd, 0x21,
	0x4a, 0xe4, 0x48, 0x89, 0x61, 0x20, 0x96, 0x6d, 0x44, 0x88, 0x4e, 0x3d, 0x7b, 0x87, 0x99, 0x1f,
	0xcc, 0xde, 0xd7, 0x21, 0x27, 0x16, 0xb9, 0x25, 0x57, 0xb1, 0x61, 0x60, 0xfa, 0x3d, 0x06, 0xf2,
	0xe3, 0x11, 0x68, 0x14, 0xa0, 0x0d, 0x09, 0x09, 0xd4, 0xeb, 0x31, 0x90, 0x1f, 0x92, 0x40, 0xe3,
	0x00, 0x2d, 0x49, 0x48, 0xa0, 0x5e, 0xcf, 0xeb, 0x29, 0x8c, 0xec, 0x86, 0xd3, 0x7b, 0xb8, 0xf0,
	0xcb, 0xab, 0xa5, 0x38, 0xd4, 0xc8, 0x2e, 0x21, 0x16, 0x7b, 0xbb, 0xb0, 0xb3, 0x3c, 0x16, 0xfb,
	0xc5, 0xda, 0x6f, 0xf7, 0x13, 0x55, 0x53, 0x6d, 0x91, 0x3d, 0xc3, 0xf9, 0xea, 0xbb, 0xd2, 0x6e,
	0x9e, 0x6b, 0xf3, 0x53, 0xba, 0xfc, 0x64, 0x46, 0xbe, 0x38, 0x31, 0x1d, 0xcc, 0xa3, 0xc7, 0xa8,
	0x18, 0xdb, 0xd7, 0xf2, 0xf4, 0x17, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x70, 0xad, 0xbc, 0x5a, 0x02,
	0x00, 0x00,
}
