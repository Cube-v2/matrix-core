// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserHTTPServer interface {
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	SetUserEmail(context.Context, *SetUserEmailReq) (*SetUserEmailReply, error)
	SetUserPhone(context.Context, *SetUserPhoneReq) (*SetUserPhoneReply, error)
	SetUserProfile(context.Context, *SetUserProfileReq) (*emptypb.Empty, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user/get", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/v1/user/set/phone", _User_SetUserPhone0_HTTP_Handler(srv))
	r.POST("/v1/user/set/email", _User_SetUserEmail0_HTTP_Handler(srv))
	r.POST("/v1/set/user/profile", _User_SetUserProfile0_HTTP_Handler(srv))
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserPhone0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserPhoneReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/SetUserPhone")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserPhone(ctx, req.(*SetUserPhoneReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetUserPhoneReply)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserEmailReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/SetUserEmail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserEmail(ctx, req.(*SetUserEmailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetUserEmailReply)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserProfile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserProfileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/SetUserProfile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserProfile(ctx, req.(*SetUserProfileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserReply, err error)
	SetUserEmail(ctx context.Context, req *SetUserEmailReq, opts ...http.CallOption) (rsp *SetUserEmailReply, err error)
	SetUserPhone(ctx context.Context, req *SetUserPhoneReq, opts ...http.CallOption) (rsp *SetUserPhoneReply, err error)
	SetUserProfile(ctx context.Context, req *SetUserProfileReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/v1/user/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserEmail(ctx context.Context, in *SetUserEmailReq, opts ...http.CallOption) (*SetUserEmailReply, error) {
	var out SetUserEmailReply
	pattern := "/v1/user/set/email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/SetUserEmail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserPhone(ctx context.Context, in *SetUserPhoneReq, opts ...http.CallOption) (*SetUserPhoneReply, error) {
	var out SetUserPhoneReply
	pattern := "/v1/user/set/phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/SetUserPhone"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserProfile(ctx context.Context, in *SetUserProfileReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/set/user/profile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/SetUserProfile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
