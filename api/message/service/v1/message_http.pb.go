// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.0
// source: message/service/v1/message.proto

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

const OperationMessageArticleCreateReview = "/message.v1.Message/ArticleCreateReview"
const OperationMessageArticleEditReview = "/message.v1.Message/ArticleEditReview"
const OperationMessageArticleImageReview = "/message.v1.Message/ArticleImageReview"
const OperationMessageAvatarReview = "/message.v1.Message/AvatarReview"
const OperationMessageCollectionsCreateReview = "/message.v1.Message/CollectionsCreateReview"
const OperationMessageCollectionsEditReview = "/message.v1.Message/CollectionsEditReview"
const OperationMessageColumnCreateReview = "/message.v1.Message/ColumnCreateReview"
const OperationMessageColumnEditReview = "/message.v1.Message/ColumnEditReview"
const OperationMessageColumnImageReview = "/message.v1.Message/ColumnImageReview"
const OperationMessageCommentCreateReview = "/message.v1.Message/CommentCreateReview"
const OperationMessageCoverReview = "/message.v1.Message/CoverReview"
const OperationMessageProfileReview = "/message.v1.Message/ProfileReview"
const OperationMessageSubCommentCreateReview = "/message.v1.Message/SubCommentCreateReview"
const OperationMessageTalkCreateReview = "/message.v1.Message/TalkCreateReview"
const OperationMessageTalkEditReview = "/message.v1.Message/TalkEditReview"
const OperationMessageTalkImageReview = "/message.v1.Message/TalkImageReview"

type MessageHTTPServer interface {
	ArticleCreateReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	ArticleEditReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	ArticleImageReview(context.Context, *ImageReviewReq) (*emptypb.Empty, error)
	AvatarReview(context.Context, *ImageReviewReq) (*emptypb.Empty, error)
	CollectionsCreateReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	CollectionsEditReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	ColumnCreateReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	ColumnEditReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	ColumnImageReview(context.Context, *ImageReviewReq) (*emptypb.Empty, error)
	CommentCreateReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	CoverReview(context.Context, *ImageReviewReq) (*emptypb.Empty, error)
	ProfileReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	SubCommentCreateReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	TalkCreateReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	TalkEditReview(context.Context, *TextReviewReq) (*emptypb.Empty, error)
	TalkImageReview(context.Context, *ImageReviewReq) (*emptypb.Empty, error)
}

func RegisterMessageHTTPServer(s *http.Server, srv MessageHTTPServer) {
	r := s.Route("/")
	r.POST("/tx/message/avatar/review", _Message_AvatarReview0_HTTP_Handler(srv))
	r.POST("/tx/message/cover/review", _Message_CoverReview0_HTTP_Handler(srv))
	r.POST("/tx/message/profile/review", _Message_ProfileReview0_HTTP_Handler(srv))
	r.POST("/tx/message/article/create/review", _Message_ArticleCreateReview0_HTTP_Handler(srv))
	r.POST("/tx/message/article/edit/review", _Message_ArticleEditReview0_HTTP_Handler(srv))
	r.POST("/tx/message/article/image/review", _Message_ArticleImageReview0_HTTP_Handler(srv))
	r.POST("/tx/message/talk/create/review", _Message_TalkCreateReview0_HTTP_Handler(srv))
	r.POST("/tx/message/talk/edit/review", _Message_TalkEditReview0_HTTP_Handler(srv))
	r.POST("/tx/message/talk/image/review", _Message_TalkImageReview0_HTTP_Handler(srv))
	r.POST("/tx/message/column/create/review", _Message_ColumnCreateReview0_HTTP_Handler(srv))
	r.POST("/tx/message/column/edit/review", _Message_ColumnEditReview0_HTTP_Handler(srv))
	r.POST("/tx/message/column/image/review", _Message_ColumnImageReview0_HTTP_Handler(srv))
	r.POST("/tx/message/collections/create/review", _Message_CollectionsCreateReview0_HTTP_Handler(srv))
	r.POST("/tx/message/collections/edit/review", _Message_CollectionsEditReview0_HTTP_Handler(srv))
	r.POST("/tx/message/comment/create/review", _Message_CommentCreateReview0_HTTP_Handler(srv))
	r.POST("/tx/message/subcomment/create/review", _Message_SubCommentCreateReview0_HTTP_Handler(srv))
}

func _Message_AvatarReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImageReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageAvatarReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AvatarReview(ctx, req.(*ImageReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_CoverReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImageReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageCoverReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CoverReview(ctx, req.(*ImageReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ProfileReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageProfileReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProfileReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ArticleCreateReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageArticleCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ArticleCreateReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ArticleEditReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageArticleEditReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ArticleEditReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ArticleImageReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImageReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageArticleImageReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ArticleImageReview(ctx, req.(*ImageReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_TalkCreateReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageTalkCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TalkCreateReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_TalkEditReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageTalkEditReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TalkEditReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_TalkImageReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImageReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageTalkImageReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TalkImageReview(ctx, req.(*ImageReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ColumnCreateReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageColumnCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ColumnCreateReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ColumnEditReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageColumnEditReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ColumnEditReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_ColumnImageReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImageReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageColumnImageReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ColumnImageReview(ctx, req.(*ImageReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_CollectionsCreateReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageCollectionsCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CollectionsCreateReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_CollectionsEditReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageCollectionsEditReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CollectionsEditReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_CommentCreateReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageCommentCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentCreateReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Message_SubCommentCreateReview0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextReviewReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageSubCommentCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubCommentCreateReview(ctx, req.(*TextReviewReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type MessageHTTPClient interface {
	ArticleCreateReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ArticleEditReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ArticleImageReview(ctx context.Context, req *ImageReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	AvatarReview(ctx context.Context, req *ImageReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CollectionsCreateReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CollectionsEditReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ColumnCreateReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ColumnEditReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ColumnImageReview(ctx context.Context, req *ImageReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CommentCreateReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CoverReview(ctx context.Context, req *ImageReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ProfileReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SubCommentCreateReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	TalkCreateReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	TalkEditReview(ctx context.Context, req *TextReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	TalkImageReview(ctx context.Context, req *ImageReviewReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type MessageHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageHTTPClient(client *http.Client) MessageHTTPClient {
	return &MessageHTTPClientImpl{client}
}

func (c *MessageHTTPClientImpl) ArticleCreateReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/article/create/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageArticleCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ArticleEditReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/article/edit/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageArticleEditReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ArticleImageReview(ctx context.Context, in *ImageReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/article/image/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageArticleImageReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) AvatarReview(ctx context.Context, in *ImageReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/avatar/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageAvatarReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) CollectionsCreateReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/collections/create/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageCollectionsCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) CollectionsEditReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/collections/edit/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageCollectionsEditReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ColumnCreateReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/column/create/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageColumnCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ColumnEditReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/column/edit/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageColumnEditReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ColumnImageReview(ctx context.Context, in *ImageReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/column/image/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageColumnImageReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) CommentCreateReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/comment/create/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageCommentCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) CoverReview(ctx context.Context, in *ImageReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/cover/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageCoverReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ProfileReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/profile/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageProfileReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) SubCommentCreateReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/subcomment/create/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageSubCommentCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) TalkCreateReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/talk/create/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageTalkCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) TalkEditReview(ctx context.Context, in *TextReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/talk/edit/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageTalkEditReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) TalkImageReview(ctx context.Context, in *ImageReviewReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/tx/message/talk/image/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageTalkImageReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
