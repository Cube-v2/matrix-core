// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsGetArticleDraftFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_ARTICLE_DRAFT_FAILED.String() && e.Code == 500
}

func ErrorGetArticleDraftFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_ARTICLE_DRAFT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateArticleDraftFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_ARTICLE_DRAFT_FAILED.String() && e.Code == 500
}

func ErrorCreateArticleDraftFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_ARTICLE_DRAFT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateArticleFolderFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_ARTICLE_FOLDER_FAILED.String() && e.Code == 500
}

func ErrorCreateArticleFolderFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_ARTICLE_FOLDER_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_RECORD_NOT_FOUND.String() && e.Code == 500
}

func ErrorRecordNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_RECORD_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsDraftMarkFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_DRAFT_MARK_FAILED.String() && e.Code == 500
}

func ErrorDraftMarkFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_DRAFT_MARK_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetDraftListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_DRAFT_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetDraftListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_DRAFT_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSendArticleFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SEND_ARTICLE_FAILED.String() && e.Code == 500
}

func ErrorSendArticleFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SEND_ARTICLE_FAILED.String(), fmt.Sprintf(format, args...))
}
