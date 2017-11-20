// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	metadata.proto

It has these top-level messages:
	AgentListResponse
	AgentGetRequest
	AgentGetResponse
	Agent
	AgentCreateRequest
	AgentCreateResponse
	AgentDeleteRequest
	ServerInfoResponse
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type AgentListResponse struct {
	Agents []*Agent `protobuf:"bytes,1,rep,name=agents" json:"agents,omitempty"`
}

func (m *AgentListResponse) Reset()                    { *m = AgentListResponse{} }
func (m *AgentListResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentListResponse) ProtoMessage()               {}
func (*AgentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentListResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type AgentGetRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *AgentGetRequest) Reset()                    { *m = AgentGetRequest{} }
func (m *AgentGetRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentGetRequest) ProtoMessage()               {}
func (*AgentGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentGetRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type AgentGetResponse struct {
	Agent *Agent `protobuf:"bytes,1,opt,name=agent" json:"agent,omitempty"`
}

func (m *AgentGetResponse) Reset()                    { *m = AgentGetResponse{} }
func (m *AgentGetResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentGetResponse) ProtoMessage()               {}
func (*AgentGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AgentGetResponse) GetAgent() *Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

type Agent struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Role       string `protobuf:"bytes,3,opt,name=role" json:"role,omitempty"`
	ExternalIp string `protobuf:"bytes,4,opt,name=external_ip,json=externalIp" json:"external_ip,omitempty"`
	InternalIp string `protobuf:"bytes,5,opt,name=internal_ip,json=internalIp" json:"internal_ip,omitempty"`
	IsDeleted  bool   `protobuf:"varint,6,opt,name=isDeleted" json:"isDeleted,omitempty"`
	CreatedAt  int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt  int64  `protobuf:"varint,8,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Agent) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Agent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Agent) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Agent) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *Agent) GetInternalIp() string {
	if m != nil {
		return m.InternalIp
	}
	return ""
}

func (m *Agent) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Agent) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Agent) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type AgentCreateRequest struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Role            string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	ExternalIp      string `protobuf:"bytes,3,opt,name=external_ip,json=externalIp" json:"external_ip,omitempty"`
	InternalIp      string `protobuf:"bytes,4,opt,name=internal_ip,json=internalIp" json:"internal_ip,omitempty"`
	SshUser         string `protobuf:"bytes,5,opt,name=ssh_user,json=sshUser" json:"ssh_user,omitempty"`
	SshPort         int32  `protobuf:"varint,6,opt,name=ssh_port,json=sshPort" json:"ssh_port,omitempty"`
	JenkinsJnlpPort int32  `protobuf:"varint,7,opt,name=jenkins_jnlp_port,json=jenkinsJnlpPort" json:"jenkins_jnlp_port,omitempty"`
	GitSshPublicKey string `protobuf:"bytes,8,opt,name=git_ssh_public_key,json=gitSshPublicKey" json:"git_ssh_public_key,omitempty"`
	JenkinsUrl      string `protobuf:"bytes,9,opt,name=jenkins_url,json=jenkinsUrl" json:"jenkins_url,omitempty"`
	CaCert          string `protobuf:"bytes,10,opt,name=ca_cert,json=caCert" json:"ca_cert,omitempty"`
}

func (m *AgentCreateRequest) Reset()                    { *m = AgentCreateRequest{} }
func (m *AgentCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateRequest) ProtoMessage()               {}
func (*AgentCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AgentCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentCreateRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *AgentCreateRequest) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *AgentCreateRequest) GetInternalIp() string {
	if m != nil {
		return m.InternalIp
	}
	return ""
}

func (m *AgentCreateRequest) GetSshUser() string {
	if m != nil {
		return m.SshUser
	}
	return ""
}

func (m *AgentCreateRequest) GetSshPort() int32 {
	if m != nil {
		return m.SshPort
	}
	return 0
}

func (m *AgentCreateRequest) GetJenkinsJnlpPort() int32 {
	if m != nil {
		return m.JenkinsJnlpPort
	}
	return 0
}

func (m *AgentCreateRequest) GetGitSshPublicKey() string {
	if m != nil {
		return m.GitSshPublicKey
	}
	return ""
}

func (m *AgentCreateRequest) GetJenkinsUrl() string {
	if m != nil {
		return m.JenkinsUrl
	}
	return ""
}

func (m *AgentCreateRequest) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

type AgentCreateResponse struct {
	Namespace              string                           `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	SshAuthorizedPublicKey string                           `protobuf:"bytes,2,opt,name=ssh_authorized_public_key,json=sshAuthorizedPublicKey" json:"ssh_authorized_public_key,omitempty"`
	GitHostname            string                           `protobuf:"bytes,3,opt,name=git_hostname,json=gitHostname" json:"git_hostname,omitempty"`
	GitHostPublicKey       string                           `protobuf:"bytes,4,opt,name=git_host_public_key,json=gitHostPublicKey" json:"git_host_public_key,omitempty"`
	GitUser                *AgentCreateResponse_ConduitUser `protobuf:"bytes,5,opt,name=git_user,json=gitUser" json:"git_user,omitempty"`
}

func (m *AgentCreateResponse) Reset()                    { *m = AgentCreateResponse{} }
func (m *AgentCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateResponse) ProtoMessage()               {}
func (*AgentCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AgentCreateResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *AgentCreateResponse) GetSshAuthorizedPublicKey() string {
	if m != nil {
		return m.SshAuthorizedPublicKey
	}
	return ""
}

func (m *AgentCreateResponse) GetGitHostname() string {
	if m != nil {
		return m.GitHostname
	}
	return ""
}

func (m *AgentCreateResponse) GetGitHostPublicKey() string {
	if m != nil {
		return m.GitHostPublicKey
	}
	return ""
}

func (m *AgentCreateResponse) GetGitUser() *AgentCreateResponse_ConduitUser {
	if m != nil {
		return m.GitUser
	}
	return nil
}

type AgentCreateResponse_ConduitUser struct {
	Phid     string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *AgentCreateResponse_ConduitUser) Reset()         { *m = AgentCreateResponse_ConduitUser{} }
func (m *AgentCreateResponse_ConduitUser) String() string { return proto.CompactTextString(m) }
func (*AgentCreateResponse_ConduitUser) ProtoMessage()    {}
func (*AgentCreateResponse_ConduitUser) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *AgentCreateResponse_ConduitUser) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *AgentCreateResponse_ConduitUser) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AgentCreateResponse_ConduitUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AgentCreateResponse_ConduitUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AgentDeleteRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *AgentDeleteRequest) Reset()                    { *m = AgentDeleteRequest{} }
func (m *AgentDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDeleteRequest) ProtoMessage()               {}
func (*AgentDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AgentDeleteRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*AgentListResponse)(nil), "appscode.ci.v1beta1.AgentListResponse")
	proto.RegisterType((*AgentGetRequest)(nil), "appscode.ci.v1beta1.AgentGetRequest")
	proto.RegisterType((*AgentGetResponse)(nil), "appscode.ci.v1beta1.AgentGetResponse")
	proto.RegisterType((*Agent)(nil), "appscode.ci.v1beta1.Agent")
	proto.RegisterType((*AgentCreateRequest)(nil), "appscode.ci.v1beta1.AgentCreateRequest")
	proto.RegisterType((*AgentCreateResponse)(nil), "appscode.ci.v1beta1.AgentCreateResponse")
	proto.RegisterType((*AgentCreateResponse_ConduitUser)(nil), "appscode.ci.v1beta1.AgentCreateResponse.ConduitUser")
	proto.RegisterType((*AgentDeleteRequest)(nil), "appscode.ci.v1beta1.AgentDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agents service

type AgentsClient interface {
	List(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*AgentListResponse, error)
	Get(ctx context.Context, in *AgentGetRequest, opts ...grpc.CallOption) (*AgentGetResponse, error)
	Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*AgentCreateResponse, error)
	Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) List(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*AgentListResponse, error) {
	out := new(AgentListResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Get(ctx context.Context, in *AgentGetRequest, opts ...grpc.CallOption) (*AgentGetResponse, error) {
	out := new(AgentGetResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*AgentCreateResponse, error) {
	out := new(AgentCreateResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	List(context.Context, *appscode_dtypes.VoidRequest) (*AgentListResponse, error)
	Get(context.Context, *AgentGetRequest) (*AgentGetResponse, error)
	Create(context.Context, *AgentCreateRequest) (*AgentCreateResponse, error)
	Delete(context.Context, *AgentDeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).List(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Get(ctx, req.(*AgentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Create(ctx, req.(*AgentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Delete(ctx, req.(*AgentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Agents_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Agents_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Agents_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Agents_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xf3, 0x9f, 0x13, 0xa4, 0x76, 0xa7, 0x08, 0xbc, 0xd9, 0x2c, 0x1b, 0x0c, 0xbb, 0x1b,
	0x65, 0xb5, 0x09, 0x9b, 0xad, 0x90, 0xca, 0x5d, 0xda, 0x4a, 0xa5, 0x80, 0x20, 0x32, 0x2a, 0x17,
	0xdc, 0x58, 0x53, 0x7b, 0x94, 0x4c, 0xea, 0x7a, 0x06, 0xcf, 0x18, 0x51, 0x10, 0x17, 0xf4, 0x16,
	0x71, 0xd5, 0x67, 0xe1, 0x9e, 0x3b, 0x1e, 0x80, 0x27, 0x40, 0xe2, 0x35, 0x90, 0xd0, 0xfc, 0xb8,
	0x76, 0x44, 0xd2, 0x74, 0x6f, 0xa2, 0xc9, 0x77, 0xbe, 0x39, 0xf3, 0x7d, 0xe7, 0xcc, 0x1c, 0x43,
	0x07, 0xcf, 0x49, 0x22, 0x47, 0x3c, 0x65, 0x92, 0xa1, 0x3d, 0xcc, 0xb9, 0x08, 0x59, 0x44, 0x46,
	0x21, 0x1d, 0x7d, 0xff, 0xea, 0x9c, 0x48, 0xfc, 0xaa, 0xdb, 0x9b, 0x33, 0x36, 0x8f, 0xc9, 0x18,
	0x73, 0x3a, 0xc6, 0x49, 0xc2, 0x24, 0x96, 0x94, 0x25, 0xc2, 0x6c, 0xe9, 0xbe, 0x97, 0x6f, 0xd9,
	0x10, 0x7f, 0xb2, 0x12, 0x8f, 0xe4, 0x15, 0x27, 0x62, 0xac, 0x7f, 0x0d, 0xc1, 0x3b, 0x81, 0x07,
	0x53, 0x25, 0xe1, 0x0b, 0x2a, 0xa4, 0x4f, 0x04, 0x67, 0x89, 0x20, 0x68, 0x02, 0x0d, 0xad, 0x4b,
	0xb8, 0x4e, 0xbf, 0x3a, 0xe8, 0x4c, 0xba, 0xa3, 0x35, 0xca, 0x46, 0x7a, 0x9f, 0x6f, 0x99, 0xde,
	0x07, 0xb0, 0xa3, 0x81, 0x13, 0x22, 0x7d, 0xf2, 0x5d, 0x46, 0x84, 0x44, 0xbb, 0x50, 0xcd, 0x68,
	0xe4, 0x3a, 0x7d, 0x67, 0xd0, 0xf6, 0xd5, 0xd2, 0x3b, 0x86, 0xdd, 0x82, 0x64, 0x0f, 0xfb, 0x08,
	0xea, 0x3a, 0x85, 0xe6, 0xdd, 0x7d, 0x96, 0x21, 0x7a, 0x7f, 0x3b, 0x50, 0xd7, 0x00, 0x42, 0x50,
	0xe3, 0x8b, 0xdb, 0x23, 0xf4, 0x5a, 0x61, 0x09, 0xbe, 0x24, 0x6e, 0xc5, 0x60, 0x6a, 0xad, 0xb0,
	0x94, 0xc5, 0xc4, 0xad, 0x1a, 0x4c, 0xad, 0xd1, 0x13, 0xe8, 0x90, 0x1f, 0x24, 0x49, 0x13, 0x1c,
	0x07, 0x94, 0xbb, 0x35, 0x1d, 0x82, 0x1c, 0x3a, 0xe5, 0x8a, 0x40, 0x93, 0x82, 0x50, 0x37, 0x84,
	0x1c, 0x3a, 0xe5, 0xa8, 0x07, 0x6d, 0x2a, 0x8e, 0x49, 0x4c, 0x24, 0x89, 0xdc, 0x46, 0xdf, 0x19,
	0xb4, 0xfc, 0x02, 0x40, 0x8f, 0x01, 0xc2, 0x94, 0x60, 0x49, 0xa2, 0x00, 0x4b, 0xb7, 0xd9, 0x77,
	0x06, 0x55, 0xbf, 0x6d, 0x91, 0xa9, 0x54, 0xe1, 0x8c, 0x47, 0x79, 0xb8, 0x65, 0xc2, 0x16, 0x99,
	0x4a, 0xef, 0xcf, 0x0a, 0x20, 0xed, 0xf1, 0x48, 0xef, 0xc8, 0x4b, 0x9a, 0x9b, 0x73, 0xd6, 0x98,
	0xab, 0x6c, 0x36, 0x57, 0xdd, 0x66, 0xae, 0xf6, 0x3f, 0x73, 0x0f, 0xa1, 0x25, 0xc4, 0x22, 0xc8,
	0x04, 0x49, 0xad, 0xf5, 0xa6, 0x10, 0x8b, 0x33, 0x41, 0xd2, 0x3c, 0xc4, 0x59, 0x2a, 0xb5, 0xed,
	0xba, 0x0e, 0xcd, 0x58, 0x2a, 0xd1, 0x10, 0x1e, 0x2c, 0x49, 0x72, 0x41, 0x13, 0x11, 0x2c, 0x93,
	0x98, 0x1b, 0x4e, 0x53, 0x73, 0x76, 0x6c, 0xe0, 0xb3, 0x24, 0xe6, 0x9a, 0xfb, 0x02, 0xd0, 0x9c,
	0xca, 0x40, 0xa7, 0xca, 0xce, 0x63, 0x1a, 0x06, 0x17, 0xe4, 0x4a, 0x57, 0xa2, 0xed, 0xef, 0xcc,
	0xa9, 0xfc, 0x5a, 0x2c, 0x66, 0x1a, 0xff, 0x9c, 0x5c, 0x29, 0xbd, 0x79, 0xe2, 0x2c, 0x8d, 0xdd,
	0xb6, 0xd1, 0x6b, 0xa1, 0xb3, 0x34, 0x46, 0xef, 0x42, 0x33, 0xc4, 0x41, 0x48, 0x52, 0xe9, 0x82,
	0x0e, 0x36, 0x42, 0x7c, 0x44, 0x52, 0xe9, 0xfd, 0x5b, 0x81, 0xbd, 0x95, 0x4a, 0xda, 0x7b, 0xd7,
	0x83, 0xb6, 0x2a, 0x9f, 0xe0, 0x38, 0xcc, 0xeb, 0x59, 0x00, 0xe8, 0x00, 0x1e, 0x2a, 0x61, 0x38,
	0x93, 0x0b, 0x96, 0xd2, 0x1f, 0x49, 0x54, 0xd6, 0x68, 0x2a, 0xfd, 0x8e, 0x10, 0x8b, 0xe9, 0x6d,
	0xbc, 0x90, 0xfa, 0x3e, 0xbc, 0xa5, 0x7c, 0x2d, 0x98, 0x90, 0xba, 0x57, 0xa6, 0xf8, 0x9d, 0x39,
	0x95, 0x9f, 0x5a, 0x08, 0xbd, 0x84, 0xbd, 0x9c, 0x52, 0xce, 0x6b, 0xba, 0xb0, 0x6b, 0x99, 0x45,
	0xc6, 0xaf, 0xa0, 0xa5, 0xe8, 0xb7, 0xbd, 0xe8, 0x4c, 0xf6, 0x37, 0xbf, 0x92, 0x55, 0x9b, 0xa3,
	0x23, 0x96, 0x44, 0x19, 0x95, 0xaa, 0x71, 0x7e, 0x73, 0x6e, 0x16, 0xdd, 0x25, 0x74, 0x4a, 0xf8,
	0xda, 0x67, 0xf4, 0x08, 0xda, 0xea, 0xbc, 0xa0, 0xf4, 0x96, 0x5a, 0x0a, 0xf8, 0x52, 0xe9, 0x7f,
	0x1b, 0xea, 0x92, 0x5d, 0x90, 0xc4, 0x7a, 0x33, 0x7f, 0x14, 0x4a, 0x2e, 0x31, 0x8d, 0xad, 0x0f,
	0xf3, 0xc7, 0x7b, 0x66, 0x2f, 0xb2, 0x79, 0x17, 0x1b, 0x67, 0xc3, 0xe4, 0x8f, 0x1a, 0x34, 0x34,
	0x51, 0xa0, 0x5f, 0x1c, 0xa8, 0xa9, 0x81, 0x84, 0x7a, 0x85, 0x4d, 0x33, 0xbb, 0x46, 0xdf, 0x30,
	0x1a, 0xd9, 0x1c, 0xdd, 0x67, 0x9b, 0x8b, 0x50, 0x1e, 0x67, 0xde, 0xeb, 0xeb, 0xdf, 0xdd, 0x4a,
	0xcb, 0xb9, 0xfe, 0xeb, 0x9f, 0x9b, 0xca, 0x73, 0xf4, 0x74, 0x1c, 0xac, 0x0c, 0xc5, 0x90, 0x8e,
	0xed, 0xd6, 0xb1, 0x19, 0x65, 0xe3, 0xa5, 0x60, 0x09, 0xfa, 0xcd, 0x81, 0xea, 0x09, 0x91, 0xe8,
	0xc3, 0xcd, 0x87, 0x14, 0xa3, 0xae, 0xfb, 0x74, 0x0b, 0xcb, 0x2a, 0x39, 0x28, 0x29, 0x79, 0x89,
	0x5e, 0x6c, 0x55, 0xf2, 0x53, 0x46, 0xa3, 0x9f, 0x8d, 0x9e, 0x1b, 0x07, 0x1a, 0xa6, 0xb5, 0xe8,
	0xf9, 0xf6, 0xe6, 0x1b, 0x55, 0x83, 0xfb, 0xde, 0x12, 0xef, 0xe3, 0x92, 0xb0, 0xa1, 0x77, 0xbf,
	0x12, 0x7d, 0xe2, 0x0c, 0xd1, 0xaf, 0x0e, 0x34, 0x4c, 0x63, 0xef, 0x52, 0xb5, 0xd2, 0xfa, 0xee,
	0xe3, 0x0d, 0x4d, 0x5d, 0x57, 0xa3, 0xe1, 0x9b, 0xd4, 0xe8, 0x70, 0x1f, 0x1e, 0x85, 0xec, 0xb2,
	0x48, 0x8f, 0x39, 0x2d, 0x69, 0x39, 0x04, 0x2d, 0x66, 0xa6, 0xbe, 0x7b, 0x33, 0xe7, 0xdb, 0xa6,
	0x85, 0xcf, 0x1b, 0xfa, 0x4b, 0xf8, 0xfa, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xca, 0x97,
	0x44, 0x8c, 0x07, 0x00, 0x00,
}
