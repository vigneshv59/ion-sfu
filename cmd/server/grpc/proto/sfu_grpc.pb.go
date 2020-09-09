// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SFUClient is the client API for SFU service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SFUClient interface {
	Signal(ctx context.Context, opts ...grpc.CallOption) (SFU_SignalClient, error)
	Relay(ctx context.Context, in *RelayRequest, opts ...grpc.CallOption) (*RelayResponse, error)
}

type sFUClient struct {
	cc grpc.ClientConnInterface
}

func NewSFUClient(cc grpc.ClientConnInterface) SFUClient {
	return &sFUClient{cc}
}

var sFUSignalStreamDesc = &grpc.StreamDesc{
	StreamName:    "Signal",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *sFUClient) Signal(ctx context.Context, opts ...grpc.CallOption) (SFU_SignalClient, error) {
	stream, err := c.cc.NewStream(ctx, sFUSignalStreamDesc, "/sfu.SFU/Signal", opts...)
	if err != nil {
		return nil, err
	}
	x := &sFUSignalClient{stream}
	return x, nil
}

type SFU_SignalClient interface {
	Send(*SignalRequest) error
	Recv() (*SignalReply, error)
	grpc.ClientStream
}

type sFUSignalClient struct {
	grpc.ClientStream
}

func (x *sFUSignalClient) Send(m *SignalRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sFUSignalClient) Recv() (*SignalReply, error) {
	m := new(SignalReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var sFURelayStreamDesc = &grpc.StreamDesc{
	StreamName: "Relay",
}

func (c *sFUClient) Relay(ctx context.Context, in *RelayRequest, opts ...grpc.CallOption) (*RelayResponse, error) {
	out := new(RelayResponse)
	err := c.cc.Invoke(ctx, "/sfu.SFU/Relay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SFUService is the service API for SFU service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSFUService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SFUService struct {
	Signal func(SFU_SignalServer) error
	Relay  func(context.Context, *RelayRequest) (*RelayResponse, error)
}

func (s *SFUService) signal(_ interface{}, stream grpc.ServerStream) error {
	return s.Signal(&sFUSignalServer{stream})
}
func (s *SFUService) relay(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Relay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sfu.SFU/Relay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Relay(ctx, req.(*RelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type SFU_SignalServer interface {
	Send(*SignalReply) error
	Recv() (*SignalRequest, error)
	grpc.ServerStream
}

type sFUSignalServer struct {
	grpc.ServerStream
}

func (x *sFUSignalServer) Send(m *SignalReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sFUSignalServer) Recv() (*SignalRequest, error) {
	m := new(SignalRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterSFUService registers a service implementation with a gRPC server.
func RegisterSFUService(s grpc.ServiceRegistrar, srv *SFUService) {
	srvCopy := *srv
	if srvCopy.Signal == nil {
		srvCopy.Signal = func(SFU_SignalServer) error {
			return status.Errorf(codes.Unimplemented, "method Signal not implemented")
		}
	}
	if srvCopy.Relay == nil {
		srvCopy.Relay = func(context.Context, *RelayRequest) (*RelayResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Relay not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "sfu.SFU",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Relay",
				Handler:    srvCopy.relay,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Signal",
				Handler:       srvCopy.signal,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "cmd/server/grpc/proto/sfu.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSFUService creates a new SFUService containing the
// implemented methods of the SFU service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSFUService(s interface{}) *SFUService {
	ns := &SFUService{}
	if h, ok := s.(interface{ Signal(SFU_SignalServer) error }); ok {
		ns.Signal = h.Signal
	}
	if h, ok := s.(interface {
		Relay(context.Context, *RelayRequest) (*RelayResponse, error)
	}); ok {
		ns.Relay = h.Relay
	}
	return ns
}

// UnstableSFUService is the service API for SFU service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSFUService interface {
	Signal(SFU_SignalServer) error
	Relay(context.Context, *RelayRequest) (*RelayResponse, error)
}
