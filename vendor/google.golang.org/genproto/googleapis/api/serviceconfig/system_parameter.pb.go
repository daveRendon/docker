// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/system_parameter.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ### System parameter configuration
//
// A system parameter is a special kind of parameter defined by the API
// system, not by an individual API. It is typically mapped to an HTTP header
// and/or a URL query parameter. This configuration specifies which methods
// change the names of the system parameters.
type SystemParameters struct {
	// Define system parameters.
	//
	// The parameters defined here will override the default parameters
	// implemented by the system. If this field is missing from the service
	// config, default system parameters will be used. Default system parameters
	// and names is implementation-dependent.
	//
	// Example: define api key and alt name for all methods
	//
	// system_parameters
	//   rules:
	//     - selector: "*"
	//       parameters:
	//         - name: api_key
	//           url_query_parameter: api_key
	//         - name: alt
	//           http_header: Response-Content-Type
	//
	// Example: define 2 api key names for a specific method.
	//
	// system_parameters
	//   rules:
	//     - selector: "/ListShelves"
	//       parameters:
	//         - name: api_key
	//           http_header: Api-Key1
	//         - name: api_key
	//           http_header: Api-Key2
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*SystemParameterRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *SystemParameters) Reset()                    { *m = SystemParameters{} }
func (m *SystemParameters) String() string            { return proto.CompactTextString(m) }
func (*SystemParameters) ProtoMessage()               {}
func (*SystemParameters) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *SystemParameters) GetRules() []*SystemParameterRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Define a system parameter rule mapping system parameter definitions to
// methods.
type SystemParameterRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Define parameters. Multiple names may be defined for a parameter.
	// For a given method call, only one of them should be used. If multiple
	// names are used the behavior is implementation-dependent.
	// If none of the specified names are present the behavior is
	// parameter-dependent.
	Parameters []*SystemParameter `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
}

func (m *SystemParameterRule) Reset()                    { *m = SystemParameterRule{} }
func (m *SystemParameterRule) String() string            { return proto.CompactTextString(m) }
func (*SystemParameterRule) ProtoMessage()               {}
func (*SystemParameterRule) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *SystemParameterRule) GetParameters() []*SystemParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Define a parameter's name and location. The parameter may be passed as either
// an HTTP header or a URL query parameter, and if both are passed the behavior
// is implementation-dependent.
type SystemParameter struct {
	// Define the name of the parameter, such as "api_key", "alt", "callback",
	// and etc. It is case sensitive.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Define the HTTP header name to use for the parameter. It is case
	// insensitive.
	HttpHeader string `protobuf:"bytes,2,opt,name=http_header,json=httpHeader" json:"http_header,omitempty"`
	// Define the URL query parameter name to use for the parameter. It is case
	// sensitive.
	UrlQueryParameter string `protobuf:"bytes,3,opt,name=url_query_parameter,json=urlQueryParameter" json:"url_query_parameter,omitempty"`
}

func (m *SystemParameter) Reset()                    { *m = SystemParameter{} }
func (m *SystemParameter) String() string            { return proto.CompactTextString(m) }
func (*SystemParameter) ProtoMessage()               {}
func (*SystemParameter) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func init() {
	proto.RegisterType((*SystemParameters)(nil), "google.api.SystemParameters")
	proto.RegisterType((*SystemParameterRule)(nil), "google.api.SystemParameterRule")
	proto.RegisterType((*SystemParameter)(nil), "google.api.SystemParameter")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/system_parameter.proto", fileDescriptor14)
}

var fileDescriptor14 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0xb6, 0x20, 0xb8, 0x4a, 0xfc, 0x71, 0x19, 0x22, 0x18, 0x8a, 0x32, 0x75, 0xb2,
	0x25, 0x10, 0x13, 0x13, 0x5d, 0xa0, 0x0b, 0x0a, 0xe1, 0x01, 0xa2, 0x10, 0x0e, 0x37, 0x52, 0x62,
	0x87, 0xb3, 0x53, 0xa9, 0xaf, 0xc3, 0x93, 0xe2, 0xb8, 0x25, 0xad, 0x22, 0xd4, 0xc5, 0x3a, 0xdf,
	0xf7, 0xbb, 0xfb, 0x4e, 0x1f, 0xbc, 0x4a, 0xad, 0x65, 0x89, 0x5c, 0xea, 0x32, 0x53, 0x92, 0x6b,
	0x92, 0x42, 0xa2, 0xaa, 0x49, 0x5b, 0x2d, 0x36, 0x52, 0x56, 0x17, 0x46, 0xb8, 0x47, 0x18, 0xa4,
	0x55, 0x91, 0x63, 0xae, 0xd5, 0x57, 0x21, 0x85, 0x59, 0x1b, 0x8b, 0x55, 0x5a, 0x67, 0x94, 0x55,
	0x68, 0x91, 0xb8, 0x9f, 0x61, 0xb0, 0xdd, 0xe7, 0x06, 0xa2, 0x05, 0x5c, 0xbc, 0x7b, 0x2a, 0xfe,
	0x83, 0x0c, 0x7b, 0x80, 0x23, 0x6a, 0x4a, 0x34, 0x61, 0x70, 0x3b, 0x9c, 0x8d, 0xef, 0xa6, 0x7c,
	0xc7, 0xf3, 0x1e, 0x9c, 0x38, 0x2e, 0xd9, 0xd0, 0x91, 0x82, 0xc9, 0x3f, 0x2a, 0xbb, 0x86, 0x13,
	0x83, 0x25, 0xe6, 0x56, 0x93, 0x5b, 0x18, 0xcc, 0x4e, 0x93, 0xee, 0xcf, 0x1e, 0x01, 0xba, 0xe3,
	0x4c, 0x38, 0xf0, 0x76, 0x37, 0x87, 0xec, 0xf6, 0xf0, 0x68, 0x05, 0xe7, 0x3d, 0x99, 0x31, 0x18,
	0x29, 0x57, 0x6e, 0x7d, 0x7c, 0xcd, 0xa6, 0x30, 0x5e, 0x5a, 0x5b, 0xa7, 0x4b, 0xcc, 0x3e, 0x91,
	0x9c, 0x49, 0x2b, 0x41, 0xdb, 0x7a, 0xf1, 0x1d, 0xc6, 0x61, 0xd2, 0x50, 0x99, 0x7e, 0x37, 0x48,
	0xeb, 0x5d, 0x56, 0xe1, 0xd0, 0x83, 0x97, 0x4e, 0x7a, 0x6b, 0x95, 0xce, 0x64, 0x2e, 0xe0, 0x2c,
	0xd7, 0xd5, 0xde, 0x95, 0xf3, 0xab, 0xde, 0x1d, 0x71, 0x1b, 0x73, 0x1c, 0xfc, 0x0c, 0x46, 0xcf,
	0x4f, 0xf1, 0xe2, 0xe3, 0xd8, 0xc7, 0x7e, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x56, 0xd1, 0x77,
	0xac, 0xc8, 0x01, 0x00, 0x00,
}
