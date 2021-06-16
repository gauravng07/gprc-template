// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shop.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SearchRequest struct {
	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=resultPerPage,proto3" json:"resultPerPage,omitempty"`
}

func (m *SearchRequest) Reset()      { *m = SearchRequest{} }
func (*SearchRequest) ProtoMessage() {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{0}
}
func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return m.Size()
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

type SearchResponse struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (m *SearchResponse) Reset()      { *m = SearchResponse{} }
func (*SearchResponse) ProtoMessage() {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{1}
}
func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return m.Size()
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SearchResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
}

func init() { proto.RegisterFile("shop.proto", fileDescriptor_0f3030369b20fd61) }

var fileDescriptor_0f3030369b20fd61 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0x86, 0x67, 0x50, 0x50, 0x5e, 0x02, 0x26, 0x13, 0x17, 0x84, 0xc5, 0x93, 0x10, 0x17, 0x2c,
	0x0c, 0x24, 0x6a, 0x3c, 0x00, 0x07, 0x30, 0xb5, 0x9e, 0x60, 0x0a, 0xcf, 0xd2, 0x48, 0x99, 0x32,
	0x33, 0x5d, 0x74, 0xe7, 0x11, 0x3c, 0x86, 0x47, 0x71, 0xd9, 0x25, 0x4b, 0x3b, 0xdd, 0xb8, 0xe4,
	0x08, 0x26, 0x03, 0x1a, 0xbb, 0x7b, 0xff, 0x9f, 0x3f, 0xff, 0xf7, 0xde, 0x03, 0x30, 0x2b, 0x95,
	0x4d, 0x33, 0xad, 0xac, 0x1a, 0x5e, 0xc5, 0x4a, 0xc5, 0x6b, 0x9a, 0x79, 0x15, 0xe5, 0x2f, 0x33,
	0x9b, 0xa4, 0x64, 0xac, 0x4c, 0x8f, 0x81, 0xf1, 0x2b, 0xf4, 0x9e, 0x49, 0xea, 0xc5, 0x2a, 0xa4,
	0x6d, 0x4e, 0xc6, 0x8a, 0x4b, 0x68, 0x6f, 0x73, 0xd2, 0xc5, 0x80, 0x8f, 0xf8, 0xa4, 0x1b, 0x1e,
	0x84, 0x40, 0x80, 0x4c, 0xc6, 0xf4, 0x98, 0xa7, 0x11, 0xe9, 0x41, 0x6b, 0xc4, 0x27, 0xed, 0xf0,
	0x9f, 0x23, 0xae, 0xa1, 0xa7, 0xc9, 0xe4, 0x6b, 0x1b, 0x90, 0x0e, 0x64, 0x4c, 0x83, 0x13, 0x1f,
	0x69, 0x9a, 0xe3, 0x00, 0xfa, 0xbf, 0x30, 0x93, 0xa9, 0x8d, 0x21, 0xd1, 0x87, 0x56, 0xb2, 0x3c,
	0xa2, 0x5a, 0xc9, 0x52, 0x08, 0x38, 0xdd, 0xc8, 0x94, 0x3c, 0xa1, 0x1b, 0xfa, 0x59, 0x0c, 0xe1,
	0x7c, 0x21, 0x2d, 0xc5, 0x4a, 0x17, 0xbe, 0xb6, 0x1b, 0xfe, 0xe9, 0xdb, 0x07, 0xe8, 0x1c, 0x1a,
	0xc5, 0x0d, 0x9c, 0xcd, 0x8b, 0x27, 0xbf, 0x6c, 0x7f, 0xda, 0x38, 0x69, 0x78, 0x31, 0x6d, 0x52,
	0xc7, 0x6c, 0x7e, 0x5f, 0x56, 0xc8, 0x76, 0x15, 0xb2, 0x7d, 0x85, 0xfc, 0xcd, 0x21, 0xff, 0x70,
	0xc8, 0x3f, 0x1d, 0xf2, 0xd2, 0x21, 0xff, 0x72, 0xc8, 0xbf, 0x1d, 0xb2, 0xbd, 0x43, 0xfe, 0x5e,
	0x23, 0x2b, 0x6b, 0x64, 0xbb, 0x1a, 0x59, 0xd4, 0xf1, 0x3f, 0xbb, 0xfb, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x01, 0x17, 0xe3, 0xe1, 0x62, 0x01, 0x00, 0x00,
}

func (this *SearchRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SearchRequest)
	if !ok {
		that2, ok := that.(SearchRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	if this.PageNumber != that1.PageNumber {
		return false
	}
	if this.ResultPerPage != that1.ResultPerPage {
		return false
	}
	return true
}
func (this *SearchResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SearchResponse)
	if !ok {
		that2, ok := that.(SearchResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Category != that1.Category {
		return false
	}
	return true
}
func (this *SearchRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&api.SearchRequest{")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "PageNumber: "+fmt.Sprintf("%#v", this.PageNumber)+",\n")
	s = append(s, "ResultPerPage: "+fmt.Sprintf("%#v", this.ResultPerPage)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SearchResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&api.SearchResponse{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Category: "+fmt.Sprintf("%#v", this.Category)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringShop(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchClient interface {
	ByQuery(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchClient struct {
	cc *grpc.ClientConn
}

func NewSearchClient(cc *grpc.ClientConn) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) ByQuery(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/Search/ByQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
type SearchServer interface {
	ByQuery(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedSearchServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (*UnimplementedSearchServer) ByQuery(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByQuery not implemented")
}

func RegisterSearchServer(s *grpc.Server, srv SearchServer) {
	s.RegisterService(&_Search_serviceDesc, srv)
}

func _Search_ByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).ByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Search/ByQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).ByQuery(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Search_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByQuery",
			Handler:    _Search_ByQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

func (m *SearchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SearchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SearchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResultPerPage != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.ResultPerPage))
		i--
		dAtA[i] = 0x18
	}
	if m.PageNumber != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.PageNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SearchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SearchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SearchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintShop(dAtA []byte, offset int, v uint64) int {
	offset -= sovShop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SearchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.PageNumber != 0 {
		n += 1 + sovShop(uint64(m.PageNumber))
	}
	if m.ResultPerPage != 0 {
		n += 1 + sovShop(uint64(m.ResultPerPage))
	}
	return n
}

func (m *SearchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	return n
}

func sovShop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShop(x uint64) (n int) {
	return sovShop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SearchRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SearchRequest{`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`PageNumber:` + fmt.Sprintf("%v", this.PageNumber) + `,`,
		`ResultPerPage:` + fmt.Sprintf("%v", this.ResultPerPage) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SearchResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SearchResponse{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Category:` + fmt.Sprintf("%v", this.Category) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringShop(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SearchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SearchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SearchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageNumber", wireType)
			}
			m.PageNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultPerPage", wireType)
			}
			m.ResultPerPage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResultPerPage |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SearchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SearchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SearchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipShop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShop
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthShop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShop = fmt.Errorf("proto: unexpected end of group")
)