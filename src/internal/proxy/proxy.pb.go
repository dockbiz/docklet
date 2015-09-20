// Code generated by protoc-gen-go.
// source: internal/proxy/proxy.proto
// DO NOT EDIT!

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	internal/proxy/proxy.proto

It has these top-level messages:
	Request
	ServiceError
	RPCError
	Response
*/
package proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RPCError_ErrorCode int32

const (
	RPCError_UNKNOWN            RPCError_ErrorCode = 0
	RPCError_CALL_NOT_FOUND     RPCError_ErrorCode = 1
	RPCError_SECURITY_VIOLATION RPCError_ErrorCode = 2
	RPCError_REQUEST_TOO_LARGE  RPCError_ErrorCode = 3
	RPCError_BAD_REQUEST        RPCError_ErrorCode = 4
	RPCError_CANCELLED          RPCError_ErrorCode = 5
	RPCError_REPLAY_ERROR       RPCError_ErrorCode = 6
	RPCError_DEADLINE_EXCEEDED  RPCError_ErrorCode = 7
)

var RPCError_ErrorCode_name = map[int32]string{
	0: "UNKNOWN",
	1: "CALL_NOT_FOUND",
	2: "SECURITY_VIOLATION",
	3: "REQUEST_TOO_LARGE",
	4: "BAD_REQUEST",
	5: "CANCELLED",
	6: "REPLAY_ERROR",
	7: "DEADLINE_EXCEEDED",
}
var RPCError_ErrorCode_value = map[string]int32{
	"UNKNOWN":            0,
	"CALL_NOT_FOUND":     1,
	"SECURITY_VIOLATION": 2,
	"REQUEST_TOO_LARGE":  3,
	"BAD_REQUEST":        4,
	"CANCELLED":          5,
	"REPLAY_ERROR":       6,
	"DEADLINE_EXCEEDED":  7,
}

func (x RPCError_ErrorCode) Enum() *RPCError_ErrorCode {
	p := new(RPCError_ErrorCode)
	*p = x
	return p
}
func (x RPCError_ErrorCode) String() string {
	return proto.EnumName(RPCError_ErrorCode_name, int32(x))
}
func (x *RPCError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RPCError_ErrorCode_value, data, "RPCError_ErrorCode")
	if err != nil {
		return err
	}
	*x = RPCError_ErrorCode(value)
	return nil
}

type Request struct {
	ServiceName      *string `protobuf:"bytes,1,req,name=service_name" json:"service_name,omitempty"`
	Method           *string `protobuf:"bytes,2,req,name=method" json:"method,omitempty"`
	Request          []byte  `protobuf:"bytes,3,req,name=request" json:"request,omitempty"`
	RequestId        *string `protobuf:"bytes,4,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Request) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Request) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

type ServiceError struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail           *string `protobuf:"bytes,2,req,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServiceError) Reset()         { *m = ServiceError{} }
func (m *ServiceError) String() string { return proto.CompactTextString(m) }
func (*ServiceError) ProtoMessage()    {}

func (m *ServiceError) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ServiceError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type RPCError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RPCError) Reset()         { *m = RPCError{} }
func (m *RPCError) String() string { return proto.CompactTextString(m) }
func (*RPCError) ProtoMessage()    {}

type Response struct {
	Response         []byte        `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Exception        []byte        `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	ServiceError     *ServiceError `protobuf:"bytes,3,opt,name=service_error" json:"service_error,omitempty"`
	CodeId           []byte        `protobuf:"bytes,4,opt,name=code_id" json:"code_id,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetException() []byte {
	if m != nil {
		return m.Exception
	}
	return nil
}

func (m *Response) GetServiceError() *ServiceError {
	if m != nil {
		return m.ServiceError
	}
	return nil
}

func (m *Response) GetCodeId() []byte {
	if m != nil {
		return m.CodeId
	}
	return nil
}

func init() {
	proto.RegisterEnum("proxy.RPCError_ErrorCode", RPCError_ErrorCode_name, RPCError_ErrorCode_value)
}