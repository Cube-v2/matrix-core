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

func IsGetLeaderBoardFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_LEADER_BOARD_FAILED.String() && e.Code == 500
}

func ErrorGetLeaderBoardFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_LEADER_BOARD_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetCollectArticleFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLLECT_ARTICLE_FAILED.String() && e.Code == 500
}

func ErrorGetCollectArticleFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLLECT_ARTICLE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetCollectionFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLLECTION_FAILED.String() && e.Code == 500
}

func ErrorGetCollectionFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLLECTION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetCollectionsListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLLECTIONS_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetCollectionsListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLLECTIONS_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTimelineListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TIMELINE_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetTimelineListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TIMELINE_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateCollectionsFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_COLLECTIONS_FAILED.String() && e.Code == 500
}

func ErrorCreateCollectionsFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_COLLECTIONS_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateTimelineFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_TIMELINE_FAILED.String() && e.Code == 500
}

func ErrorCreateTimelineFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_TIMELINE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsEditCollectionsFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_EDIT_COLLECTIONS_FAILED.String() && e.Code == 500
}

func ErrorEditCollectionsFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_EDIT_COLLECTIONS_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsDeleteCollectionsFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_DELETE_COLLECTIONS_FAILED.String() && e.Code == 500
}

func ErrorDeleteCollectionsFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_DELETE_COLLECTIONS_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetArticleFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_ARTICLE_FAILED.String() && e.Code == 500
}

func ErrorGetArticleFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_ARTICLE_FAILED.String(), fmt.Sprintf(format, args...))
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

func IsGetArticleListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_ARTICLE_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetArticleListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_ARTICLE_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetArticleSearchFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_ARTICLE_SEARCH_FAILED.String() && e.Code == 500
}

func ErrorGetArticleSearchFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_ARTICLE_SEARCH_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetArticleAgreeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_ARTICLE_AGREE_FAILED.String() && e.Code == 500
}

func ErrorGetArticleAgreeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_ARTICLE_AGREE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetArticleCollectFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_ARTICLE_COLLECT_FAILED.String() && e.Code == 500
}

func ErrorGetArticleCollectFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_ARTICLE_COLLECT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateArticleFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_ARTICLE_FAILED.String() && e.Code == 500
}

func ErrorCreateArticleFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_ARTICLE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsEditArticleFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_EDIT_ARTICLE_FAILED.String() && e.Code == 500
}

func ErrorEditArticleFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_EDIT_ARTICLE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsDeleteArticleFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_DELETE_ARTICLE_FAILED.String() && e.Code == 500
}

func ErrorDeleteArticleFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_DELETE_ARTICLE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTalkFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TALK_FAILED.String() && e.Code == 500
}

func ErrorGetTalkFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TALK_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTalkAgreeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TALK_AGREE_FAILED.String() && e.Code == 500
}

func ErrorGetTalkAgreeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TALK_AGREE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTalkCollectFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TALK_COLLECT_FAILED.String() && e.Code == 500
}

func ErrorGetTalkCollectFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TALK_COLLECT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTalkDraftFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TALK_DRAFT_FAILED.String() && e.Code == 500
}

func ErrorGetTalkDraftFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TALK_DRAFT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTalkListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TALK_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetTalkListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TALK_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetTalkSearchFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_TALK_SEARCH_FAILED.String() && e.Code == 500
}

func ErrorGetTalkSearchFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_TALK_SEARCH_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateTalkFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_TALK_FAILED.String() && e.Code == 500
}

func ErrorCreateTalkFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_TALK_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsEditTalkFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_EDIT_TALK_FAILED.String() && e.Code == 500
}

func ErrorEditTalkFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_EDIT_TALK_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsDeleteTalkFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_DELETE_TALK_FAILED.String() && e.Code == 500
}

func ErrorDeleteTalkFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_DELETE_TALK_FAILED.String(), fmt.Sprintf(format, args...))
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

func IsCreateDraftFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_DRAFT_FAILED.String() && e.Code == 500
}

func ErrorCreateDraftFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_DRAFT_FAILED.String(), fmt.Sprintf(format, args...))
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

func IsGetColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorGetColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetColumnAgreeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_AGREE_FAILED.String() && e.Code == 500
}

func ErrorGetColumnAgreeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_AGREE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetColumnCollectFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_COLLECT_FAILED.String() && e.Code == 500
}

func ErrorGetColumnCollectFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_COLLECT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetColumnDraftFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_DRAFT_FAILED.String() && e.Code == 500
}

func ErrorGetColumnDraftFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_DRAFT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetColumnSearchFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_SEARCH_FAILED.String() && e.Code == 500
}

func ErrorGetColumnSearchFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_SEARCH_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorCreateColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsEditColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_EDIT_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorEditColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_EDIT_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsDeleteColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_DELETE_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorDeleteColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_DELETE_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetColumnListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetColumnListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetSubscribeColumnListFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_SUBSCRIBE_COLUMN_LIST_FAILED.String() && e.Code == 500
}

func ErrorGetSubscribeColumnListFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_SUBSCRIBE_COLUMN_LIST_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetColumnSubscribesFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COLUMN_SUBSCRIBES_FAILED.String() && e.Code == 500
}

func ErrorGetColumnSubscribesFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COLUMN_SUBSCRIBES_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetSubscribeColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_SUBSCRIBE_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorGetSubscribeColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_SUBSCRIBE_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsAddColumnIncludesFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_ADD_COLUMN_INCLUDES_FAILED.String() && e.Code == 500
}

func ErrorAddColumnIncludesFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_ADD_COLUMN_INCLUDES_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsDeleteColumnIncludesFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_DELETE_COLUMN_INCLUDES_FAILED.String() && e.Code == 500
}

func ErrorDeleteColumnIncludesFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_DELETE_COLUMN_INCLUDES_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSubscribeColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SUBSCRIBE_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorSubscribeColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SUBSCRIBE_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSubscribeColumnJudgeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SUBSCRIBE_COLUMN_JUDGE_FAILED.String() && e.Code == 500
}

func ErrorSubscribeColumnJudgeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SUBSCRIBE_COLUMN_JUDGE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCancelSubscribeColumnFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CANCEL_SUBSCRIBE_COLUMN_FAILED.String() && e.Code == 500
}

func ErrorCancelSubscribeColumnFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CANCEL_SUBSCRIBE_COLUMN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetNewsFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_NEWS_FAILED.String() && e.Code == 500
}

func ErrorGetNewsFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_NEWS_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetAgreeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SET_AGREE_FAILED.String() && e.Code == 500
}

func ErrorSetAgreeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SET_AGREE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetViewFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SET_VIEW_FAILED.String() && e.Code == 500
}

func ErrorSetViewFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SET_VIEW_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetCollectFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SET_COLLECT_FAILED.String() && e.Code == 500
}

func ErrorSetCollectFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SET_COLLECT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetImageIrregularFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SET_IMAGE_IRREGULAR_FAILED.String() && e.Code == 500
}

func ErrorSetImageIrregularFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SET_IMAGE_IRREGULAR_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetContentIrregularFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SET_CONTENT_IRREGULAR_FAILED.String() && e.Code == 500
}

func ErrorSetContentIrregularFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SET_CONTENT_IRREGULAR_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCancelAgreeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CANCEL_AGREE_FAILED.String() && e.Code == 500
}

func ErrorCancelAgreeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CANCEL_AGREE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCancelViewFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CANCEL_VIEW_FAILED.String() && e.Code == 500
}

func ErrorCancelViewFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CANCEL_VIEW_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCancelCollectFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CANCEL_COLLECT_FAILED.String() && e.Code == 500
}

func ErrorCancelCollectFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CANCEL_COLLECT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetStatisticFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_STATISTIC_FAILED.String() && e.Code == 500
}

func ErrorGetStatisticFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_STATISTIC_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetStatisticJudgeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_STATISTIC_JUDGE_FAILED.String() && e.Code == 500
}

func ErrorGetStatisticJudgeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_STATISTIC_JUDGE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetCountFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_COUNT_FAILED.String() && e.Code == 500
}

func ErrorGetCountFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_COUNT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetCreationUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_CREATION_USER_FAILED.String() && e.Code == 500
}

func ErrorGetCreationUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_CREATION_USER_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetImageReviewFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_IMAGE_REVIEW_FAILED.String() && e.Code == 500
}

func ErrorGetImageReviewFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_IMAGE_REVIEW_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetContentReviewFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_GET_CONTENT_REVIEW_FAILED.String() && e.Code == 500
}

func ErrorGetContentReviewFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_GET_CONTENT_REVIEW_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetRecordFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_SET_RECORD_FAILED.String() && e.Code == 500
}

func ErrorSetRecordFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_SET_RECORD_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsNotEmpty(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_NOT_EMPTY.String() && e.Code == 500
}

func ErrorNotEmpty(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_NOT_EMPTY.String(), fmt.Sprintf(format, args...))
}

func IsAddCommentFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_ADD_COMMENT_FAILED.String() && e.Code == 500
}

func ErrorAddCommentFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_ADD_COMMENT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsReduceCommentFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_REDUCE_COMMENT_FAILED.String() && e.Code == 500
}

func ErrorReduceCommentFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_REDUCE_COMMENT_FAILED.String(), fmt.Sprintf(format, args...))
}
