// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routeguide.proto

package routeguide

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude             int32    `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            int32    `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a1fa0460de48b1, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	Lo                   *Point   `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	Hi                   *Point   `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a1fa0460de48b1, []int{1}
}

func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (m *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(m, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a given point.
// If a feature could not be named, the name is empty.
type Feature struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location             *Point   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a1fa0460de48b1, []int{2}
}

func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (m *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(m, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

//A RouteNote is a message sent while at a given point..
type RouteNote struct {
	Location             *Point   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteNote) Reset()         { *m = RouteNote{} }
func (m *RouteNote) String() string { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()    {}
func (*RouteNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a1fa0460de48b1, []int{3}
}

func (m *RouteNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteNote.Unmarshal(m, b)
}
func (m *RouteNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteNote.Marshal(b, m, deterministic)
}
func (m *RouteNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteNote.Merge(m, src)
}
func (m *RouteNote) XXX_Size() int {
	return xxx_messageInfo_RouteNote.Size(m)
}
func (m *RouteNote) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteNote.DiscardUnknown(m)
}

var xxx_messageInfo_RouteNote proto.InternalMessageInfo

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	PointCount           int32    `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	FeatureCount         int32    `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	Distance             int32    `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	ElapsedTime          int32    `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteSummary) Reset()         { *m = RouteSummary{} }
func (m *RouteSummary) String() string { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()    {}
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a1fa0460de48b1, []int{4}
}

func (m *RouteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSummary.Unmarshal(m, b)
}
func (m *RouteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSummary.Marshal(b, m, deterministic)
}
func (m *RouteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSummary.Merge(m, src)
}
func (m *RouteSummary) XXX_Size() int {
	return xxx_messageInfo_RouteSummary.Size(m)
}
func (m *RouteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSummary proto.InternalMessageInfo

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "routeguide.Point")
	proto.RegisterType((*Rectangle)(nil), "routeguide.Rectangle")
	proto.RegisterType((*Feature)(nil), "routeguide.Feature")
	proto.RegisterType((*RouteNote)(nil), "routeguide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "routeguide.RouteSummary")
}

func init() { proto.RegisterFile("routeguide.proto", fileDescriptor_f4a1fa0460de48b1) }

var fileDescriptor_f4a1fa0460de48b1 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x62, 0x6b, 0x9b, 0x49, 0x04, 0x1d, 0x11, 0x42, 0x10, 0xb4, 0xf1, 0xd2, 0x8b, 0xa5,
	0x54, 0xf0, 0x58, 0x91, 0x82, 0xbd, 0x14, 0xd1, 0xd8, 0x7b, 0x59, 0x93, 0x35, 0x5d, 0x48, 0xb2,
	0x25, 0x99, 0x1c, 0xfc, 0x01, 0xfe, 0x02, 0xff, 0xb0, 0x64, 0x93, 0xb4, 0xa9, 0xb6, 0x78, 0xdb,
	0x79, 0xf3, 0xde, 0x9b, 0xaf, 0x85, 0xd3, 0x54, 0xe6, 0xc4, 0xc3, 0x5c, 0x04, 0x7c, 0xb8, 0x4e,
	0x25, 0x49, 0x84, 0x2d, 0xe2, 0x3e, 0x42, 0xe7, 0x45, 0x8a, 0x84, 0xd0, 0x81, 0x5e, 0xc4, 0x48,
	0x50, 0x1e, 0x70, 0x5b, 0xbb, 0xd6, 0x06, 0x1d, 0x6f, 0x13, 0xe3, 0x25, 0x18, 0x91, 0x4c, 0xc2,
	0x32, 0xa9, 0xab, 0xe4, 0x16, 0x70, 0x5f, 0xc1, 0xf0, 0xb8, 0x4f, 0x2c, 0x09, 0x23, 0x8e, 0x7d,
	0xd0, 0x23, 0xa9, 0x0c, 0xcc, 0xf1, 0xd9, 0xb0, 0x51, 0x5a, 0x55, 0xf1, 0xf4, 0x48, 0x16, 0x94,
	0x95, 0x50, 0x36, 0xfb, 0x29, 0x2b, 0xe1, 0xce, 0xa1, 0xfb, 0xc4, 0x19, 0xe5, 0x29, 0x47, 0x84,
	0x76, 0xc2, 0xe2, 0xb2, 0x27, 0xc3, 0x53, 0x6f, 0xbc, 0x85, 0x5e, 0x24, 0x7d, 0x46, 0x42, 0x26,
	0x87, 0x7d, 0x36, 0x14, 0x77, 0x01, 0x86, 0x57, 0x64, 0x9f, 0x25, 0xed, 0x6a, 0xb5, 0x7f, 0xb5,
	0x68, 0x43, 0x37, 0xe6, 0x59, 0xc6, 0xc2, 0x72, 0x70, 0xc3, 0xab, 0x43, 0xf7, 0x5b, 0x03, 0x4b,
	0xd9, 0xbe, 0xe5, 0x71, 0xcc, 0xd2, 0x4f, 0xbc, 0x02, 0x73, 0x5d, 0xa8, 0x97, 0xbe, 0xcc, 0x13,
	0xaa, 0x96, 0x08, 0x0a, 0x9a, 0x16, 0x08, 0xde, 0xc0, 0xc9, 0x47, 0x39, 0x55, 0x45, 0x29, 0x57,
	0x69, 0x55, 0x60, 0x49, 0x72, 0xa0, 0x17, 0x88, 0x8c, 0x58, 0xe2, 0x73, 0xfb, 0xa8, 0xbc, 0x43,
	0x1d, 0x63, 0x1f, 0x2c, 0x1e, 0xb1, 0x75, 0xc6, 0x83, 0x25, 0x89, 0x98, 0xdb, 0x6d, 0x95, 0x37,
	0x2b, 0x6c, 0x21, 0x62, 0x3e, 0xfe, 0xd2, 0x01, 0x54, 0x57, 0xb3, 0x62, 0x1c, 0xbc, 0x07, 0x98,
	0x71, 0xaa, 0x77, 0xf9, 0x77, 0x52, 0xe7, 0xbc, 0x09, 0x55, 0x3c, 0xb7, 0x85, 0x13, 0xb0, 0xe6,
	0x22, 0xa3, 0xaa, 0xb3, 0x0c, 0x2f, 0x9a, 0xb4, 0xcd, 0xb5, 0x0f, 0xa8, 0x47, 0x1a, 0x4e, 0xc0,
	0xf4, 0xb8, 0x2f, 0xd3, 0x40, 0xf5, 0xb2, 0xaf, 0xb0, 0xbd, 0xe3, 0xd8, 0xd8, 0xa3, 0xdb, 0x1a,
	0x68, 0xf8, 0x50, 0x9d, 0x6c, 0xba, 0x62, 0xf4, 0xab, 0x78, 0x7d, 0x49, 0x67, 0x3f, 0x5c, 0xc8,
	0x47, 0xda, 0xfb, 0xb1, 0xfa, 0xea, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x91, 0xa6,
	0x43, 0xfe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteGuideClient interface {
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features
	Listfeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListfeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) Listfeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListfeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[0], "/routeguide.RouteGuide/Listfeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListfeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListfeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListfeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListfeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[1], "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[2], "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
type RouteGuideServer interface {
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features
	Listfeatures(*Rectangle, RouteGuide_ListfeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuide_RouteChatServer) error
}

// UnimplementedRouteGuideServer can be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (*UnimplementedRouteGuideServer) GetFeature(ctx context.Context, req *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (*UnimplementedRouteGuideServer) Listfeatures(req *Rectangle, srv RouteGuide_ListfeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method Listfeatures not implemented")
}
func (*UnimplementedRouteGuideServer) RecordRoute(srv RouteGuide_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (*UnimplementedRouteGuideServer) RouteChat(srv RouteGuide_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_Listfeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).Listfeatures(m, &routeGuideListfeaturesServer{stream})
}

type RouteGuide_ListfeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListfeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListfeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listfeatures",
			Handler:       _RouteGuide_Listfeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routeguide.proto",
}