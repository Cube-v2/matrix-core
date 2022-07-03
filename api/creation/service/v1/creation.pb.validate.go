// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: creation/service/v1/creation.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _creation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetLastArticleDraftReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLastArticleDraftReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLastArticleDraftReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLastArticleDraftReqMultiError, or nil if none found.
func (m *GetLastArticleDraftReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLastArticleDraftReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = GetLastArticleDraftReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetLastArticleDraftReqMultiError(errors)
	}

	return nil
}

func (m *GetLastArticleDraftReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetLastArticleDraftReqMultiError is an error wrapping multiple validation
// errors returned by GetLastArticleDraftReq.ValidateAll() if the designated
// constraints aren't met.
type GetLastArticleDraftReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLastArticleDraftReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLastArticleDraftReqMultiError) AllErrors() []error { return m }

// GetLastArticleDraftReqValidationError is the validation error returned by
// GetLastArticleDraftReq.Validate if the designated constraints aren't met.
type GetLastArticleDraftReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLastArticleDraftReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLastArticleDraftReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLastArticleDraftReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLastArticleDraftReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLastArticleDraftReqValidationError) ErrorName() string {
	return "GetLastArticleDraftReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetLastArticleDraftReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLastArticleDraftReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLastArticleDraftReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLastArticleDraftReqValidationError{}

// Validate checks the field values on GetLastArticleDraftReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLastArticleDraftReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLastArticleDraftReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLastArticleDraftReplyMultiError, or nil if none found.
func (m *GetLastArticleDraftReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLastArticleDraftReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	if len(errors) > 0 {
		return GetLastArticleDraftReplyMultiError(errors)
	}

	return nil
}

// GetLastArticleDraftReplyMultiError is an error wrapping multiple validation
// errors returned by GetLastArticleDraftReply.ValidateAll() if the designated
// constraints aren't met.
type GetLastArticleDraftReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLastArticleDraftReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLastArticleDraftReplyMultiError) AllErrors() []error { return m }

// GetLastArticleDraftReplyValidationError is the validation error returned by
// GetLastArticleDraftReply.Validate if the designated constraints aren't met.
type GetLastArticleDraftReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLastArticleDraftReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLastArticleDraftReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLastArticleDraftReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLastArticleDraftReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLastArticleDraftReplyValidationError) ErrorName() string {
	return "GetLastArticleDraftReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetLastArticleDraftReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLastArticleDraftReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLastArticleDraftReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLastArticleDraftReplyValidationError{}

// Validate checks the field values on CreateArticleReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleReqMultiError, or nil if none found.
func (m *CreateArticleReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = CreateArticleReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateArticleReqMultiError(errors)
	}

	return nil
}

func (m *CreateArticleReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateArticleReqMultiError is an error wrapping multiple validation errors
// returned by CreateArticleReq.ValidateAll() if the designated constraints
// aren't met.
type CreateArticleReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleReqMultiError) AllErrors() []error { return m }

// CreateArticleReqValidationError is the validation error returned by
// CreateArticleReq.Validate if the designated constraints aren't met.
type CreateArticleReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleReqValidationError) ErrorName() string { return "CreateArticleReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateArticleReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleReqValidationError{}

// Validate checks the field values on CreateArticleCacheAndSearchReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleCacheAndSearchReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleCacheAndSearchReq with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateArticleCacheAndSearchReqMultiError, or nil if none found.
func (m *CreateArticleCacheAndSearchReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleCacheAndSearchReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = CreateArticleCacheAndSearchReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateArticleCacheAndSearchReqMultiError(errors)
	}

	return nil
}

func (m *CreateArticleCacheAndSearchReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateArticleCacheAndSearchReqMultiError is an error wrapping multiple
// validation errors returned by CreateArticleCacheAndSearchReq.ValidateAll()
// if the designated constraints aren't met.
type CreateArticleCacheAndSearchReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleCacheAndSearchReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleCacheAndSearchReqMultiError) AllErrors() []error { return m }

// CreateArticleCacheAndSearchReqValidationError is the validation error
// returned by CreateArticleCacheAndSearchReq.Validate if the designated
// constraints aren't met.
type CreateArticleCacheAndSearchReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleCacheAndSearchReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleCacheAndSearchReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleCacheAndSearchReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleCacheAndSearchReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleCacheAndSearchReqValidationError) ErrorName() string {
	return "CreateArticleCacheAndSearchReqValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleCacheAndSearchReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleCacheAndSearchReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleCacheAndSearchReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleCacheAndSearchReqValidationError{}

// Validate checks the field values on CreateArticleDraftReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleDraftReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleDraftReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleDraftReqMultiError, or nil if none found.
func (m *CreateArticleDraftReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleDraftReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = CreateArticleDraftReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateArticleDraftReqMultiError(errors)
	}

	return nil
}

func (m *CreateArticleDraftReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateArticleDraftReqMultiError is an error wrapping multiple validation
// errors returned by CreateArticleDraftReq.ValidateAll() if the designated
// constraints aren't met.
type CreateArticleDraftReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleDraftReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleDraftReqMultiError) AllErrors() []error { return m }

// CreateArticleDraftReqValidationError is the validation error returned by
// CreateArticleDraftReq.Validate if the designated constraints aren't met.
type CreateArticleDraftReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleDraftReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleDraftReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleDraftReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleDraftReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleDraftReqValidationError) ErrorName() string {
	return "CreateArticleDraftReqValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleDraftReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleDraftReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleDraftReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleDraftReqValidationError{}

// Validate checks the field values on CreateArticleDraftReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleDraftReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleDraftReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleDraftReplyMultiError, or nil if none found.
func (m *CreateArticleDraftReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleDraftReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateArticleDraftReplyMultiError(errors)
	}

	return nil
}

// CreateArticleDraftReplyMultiError is an error wrapping multiple validation
// errors returned by CreateArticleDraftReply.ValidateAll() if the designated
// constraints aren't met.
type CreateArticleDraftReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleDraftReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleDraftReplyMultiError) AllErrors() []error { return m }

// CreateArticleDraftReplyValidationError is the validation error returned by
// CreateArticleDraftReply.Validate if the designated constraints aren't met.
type CreateArticleDraftReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleDraftReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleDraftReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleDraftReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleDraftReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleDraftReplyValidationError) ErrorName() string {
	return "CreateArticleDraftReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleDraftReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleDraftReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleDraftReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleDraftReplyValidationError{}

// Validate checks the field values on ArticleDraftMarkReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ArticleDraftMarkReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleDraftMarkReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ArticleDraftMarkReqMultiError, or nil if none found.
func (m *ArticleDraftMarkReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleDraftMarkReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = ArticleDraftMarkReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ArticleDraftMarkReqMultiError(errors)
	}

	return nil
}

func (m *ArticleDraftMarkReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ArticleDraftMarkReqMultiError is an error wrapping multiple validation
// errors returned by ArticleDraftMarkReq.ValidateAll() if the designated
// constraints aren't met.
type ArticleDraftMarkReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleDraftMarkReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleDraftMarkReqMultiError) AllErrors() []error { return m }

// ArticleDraftMarkReqValidationError is the validation error returned by
// ArticleDraftMarkReq.Validate if the designated constraints aren't met.
type ArticleDraftMarkReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleDraftMarkReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleDraftMarkReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleDraftMarkReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleDraftMarkReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleDraftMarkReqValidationError) ErrorName() string {
	return "ArticleDraftMarkReqValidationError"
}

// Error satisfies the builtin error interface
func (e ArticleDraftMarkReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleDraftMarkReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleDraftMarkReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleDraftMarkReqValidationError{}

// Validate checks the field values on GetArticleDraftListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetArticleDraftListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetArticleDraftListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetArticleDraftListReqMultiError, or nil if none found.
func (m *GetArticleDraftListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetArticleDraftListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = GetArticleDraftListReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetArticleDraftListReqMultiError(errors)
	}

	return nil
}

func (m *GetArticleDraftListReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetArticleDraftListReqMultiError is an error wrapping multiple validation
// errors returned by GetArticleDraftListReq.ValidateAll() if the designated
// constraints aren't met.
type GetArticleDraftListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetArticleDraftListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetArticleDraftListReqMultiError) AllErrors() []error { return m }

// GetArticleDraftListReqValidationError is the validation error returned by
// GetArticleDraftListReq.Validate if the designated constraints aren't met.
type GetArticleDraftListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleDraftListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleDraftListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleDraftListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleDraftListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleDraftListReqValidationError) ErrorName() string {
	return "GetArticleDraftListReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleDraftListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleDraftListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleDraftListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleDraftListReqValidationError{}

// Validate checks the field values on GetArticleDraftListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetArticleDraftListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetArticleDraftListReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetArticleDraftListReplyMultiError, or nil if none found.
func (m *GetArticleDraftListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetArticleDraftListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDraft() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetArticleDraftListReplyValidationError{
						field:  fmt.Sprintf("Draft[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetArticleDraftListReplyValidationError{
						field:  fmt.Sprintf("Draft[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetArticleDraftListReplyValidationError{
					field:  fmt.Sprintf("Draft[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetArticleDraftListReplyMultiError(errors)
	}

	return nil
}

// GetArticleDraftListReplyMultiError is an error wrapping multiple validation
// errors returned by GetArticleDraftListReply.ValidateAll() if the designated
// constraints aren't met.
type GetArticleDraftListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetArticleDraftListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetArticleDraftListReplyMultiError) AllErrors() []error { return m }

// GetArticleDraftListReplyValidationError is the validation error returned by
// GetArticleDraftListReply.Validate if the designated constraints aren't met.
type GetArticleDraftListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleDraftListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleDraftListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleDraftListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleDraftListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleDraftListReplyValidationError) ErrorName() string {
	return "GetArticleDraftListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleDraftListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleDraftListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleDraftListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleDraftListReplyValidationError{}

// Validate checks the field values on SendArticleReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendArticleReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendArticleReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendArticleReqMultiError,
// or nil if none found.
func (m *SendArticleReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendArticleReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = SendArticleReqValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendArticleReqMultiError(errors)
	}

	return nil
}

func (m *SendArticleReq) _validateUuid(uuid string) error {
	if matched := _creation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// SendArticleReqMultiError is an error wrapping multiple validation errors
// returned by SendArticleReq.ValidateAll() if the designated constraints
// aren't met.
type SendArticleReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendArticleReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendArticleReqMultiError) AllErrors() []error { return m }

// SendArticleReqValidationError is the validation error returned by
// SendArticleReq.Validate if the designated constraints aren't met.
type SendArticleReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendArticleReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendArticleReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendArticleReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendArticleReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendArticleReqValidationError) ErrorName() string { return "SendArticleReqValidationError" }

// Error satisfies the builtin error interface
func (e SendArticleReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendArticleReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendArticleReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendArticleReqValidationError{}

// Validate checks the field values on GetArticleDraftListReply_Draft with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetArticleDraftListReply_Draft) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetArticleDraftListReply_Draft with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetArticleDraftListReply_DraftMultiError, or nil if none found.
func (m *GetArticleDraftListReply_Draft) ValidateAll() error {
	return m.validate(true)
}

func (m *GetArticleDraftListReply_Draft) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetArticleDraftListReply_DraftMultiError(errors)
	}

	return nil
}

// GetArticleDraftListReply_DraftMultiError is an error wrapping multiple
// validation errors returned by GetArticleDraftListReply_Draft.ValidateAll()
// if the designated constraints aren't met.
type GetArticleDraftListReply_DraftMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetArticleDraftListReply_DraftMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetArticleDraftListReply_DraftMultiError) AllErrors() []error { return m }

// GetArticleDraftListReply_DraftValidationError is the validation error
// returned by GetArticleDraftListReply_Draft.Validate if the designated
// constraints aren't met.
type GetArticleDraftListReply_DraftValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleDraftListReply_DraftValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleDraftListReply_DraftValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleDraftListReply_DraftValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleDraftListReply_DraftValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleDraftListReply_DraftValidationError) ErrorName() string {
	return "GetArticleDraftListReply_DraftValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleDraftListReply_DraftValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleDraftListReply_Draft.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleDraftListReply_DraftValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleDraftListReply_DraftValidationError{}
