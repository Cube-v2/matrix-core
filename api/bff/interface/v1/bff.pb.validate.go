// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bff/interface/v1/bff.proto

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

// Validate checks the field values on UserRegisterReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserRegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRegisterReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRegisterReqMultiError, or nil if none found.
func (m *UserRegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	// no validation rules for Code

	if len(errors) > 0 {
		return UserRegisterReqMultiError(errors)
	}

	return nil
}

// UserRegisterReqMultiError is an error wrapping multiple validation errors
// returned by UserRegisterReq.ValidateAll() if the designated constraints
// aren't met.
type UserRegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRegisterReqMultiError) AllErrors() []error { return m }

// UserRegisterReqValidationError is the validation error returned by
// UserRegisterReq.Validate if the designated constraints aren't met.
type UserRegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRegisterReqValidationError) ErrorName() string { return "UserRegisterReqValidationError" }

// Error satisfies the builtin error interface
func (e UserRegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRegisterReqValidationError{}

// Validate checks the field values on LoginByPasswordReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LoginByPasswordReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginByPasswordReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginByPasswordReqMultiError, or nil if none found.
func (m *LoginByPasswordReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginByPasswordReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Account

	// no validation rules for Password

	// no validation rules for Mode

	if len(errors) > 0 {
		return LoginByPasswordReqMultiError(errors)
	}

	return nil
}

// LoginByPasswordReqMultiError is an error wrapping multiple validation errors
// returned by LoginByPasswordReq.ValidateAll() if the designated constraints
// aren't met.
type LoginByPasswordReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginByPasswordReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginByPasswordReqMultiError) AllErrors() []error { return m }

// LoginByPasswordReqValidationError is the validation error returned by
// LoginByPasswordReq.Validate if the designated constraints aren't met.
type LoginByPasswordReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginByPasswordReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginByPasswordReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginByPasswordReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginByPasswordReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginByPasswordReqValidationError) ErrorName() string {
	return "LoginByPasswordReqValidationError"
}

// Error satisfies the builtin error interface
func (e LoginByPasswordReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginByPasswordReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginByPasswordReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginByPasswordReqValidationError{}

// Validate checks the field values on LoginByCodeReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginByCodeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginByCodeReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginByCodeReqMultiError,
// or nil if none found.
func (m *LoginByCodeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginByCodeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Phone

	// no validation rules for Code

	if len(errors) > 0 {
		return LoginByCodeReqMultiError(errors)
	}

	return nil
}

// LoginByCodeReqMultiError is an error wrapping multiple validation errors
// returned by LoginByCodeReq.ValidateAll() if the designated constraints
// aren't met.
type LoginByCodeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginByCodeReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginByCodeReqMultiError) AllErrors() []error { return m }

// LoginByCodeReqValidationError is the validation error returned by
// LoginByCodeReq.Validate if the designated constraints aren't met.
type LoginByCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginByCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginByCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginByCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginByCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginByCodeReqValidationError) ErrorName() string { return "LoginByCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginByCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginByCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginByCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginByCodeReqValidationError{}

// Validate checks the field values on LoginByWeChatReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginByWeChatReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginByWeChatReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginByWeChatReqMultiError, or nil if none found.
func (m *LoginByWeChatReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginByWeChatReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Wechat

	if len(errors) > 0 {
		return LoginByWeChatReqMultiError(errors)
	}

	return nil
}

// LoginByWeChatReqMultiError is an error wrapping multiple validation errors
// returned by LoginByWeChatReq.ValidateAll() if the designated constraints
// aren't met.
type LoginByWeChatReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginByWeChatReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginByWeChatReqMultiError) AllErrors() []error { return m }

// LoginByWeChatReqValidationError is the validation error returned by
// LoginByWeChatReq.Validate if the designated constraints aren't met.
type LoginByWeChatReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginByWeChatReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginByWeChatReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginByWeChatReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginByWeChatReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginByWeChatReqValidationError) ErrorName() string { return "LoginByWeChatReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginByWeChatReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginByWeChatReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginByWeChatReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginByWeChatReqValidationError{}

// Validate checks the field values on LoginByGithubReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginByGithubReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginByGithubReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginByGithubReqMultiError, or nil if none found.
func (m *LoginByGithubReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginByGithubReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Github

	if len(errors) > 0 {
		return LoginByGithubReqMultiError(errors)
	}

	return nil
}

// LoginByGithubReqMultiError is an error wrapping multiple validation errors
// returned by LoginByGithubReq.ValidateAll() if the designated constraints
// aren't met.
type LoginByGithubReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginByGithubReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginByGithubReqMultiError) AllErrors() []error { return m }

// LoginByGithubReqValidationError is the validation error returned by
// LoginByGithubReq.Validate if the designated constraints aren't met.
type LoginByGithubReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginByGithubReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginByGithubReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginByGithubReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginByGithubReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginByGithubReqValidationError) ErrorName() string { return "LoginByGithubReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginByGithubReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginByGithubReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginByGithubReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginByGithubReqValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on LoginPasswordResetReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LoginPasswordResetReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginPasswordResetReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginPasswordResetReqMultiError, or nil if none found.
func (m *LoginPasswordResetReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginPasswordResetReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Account

	// no validation rules for Code

	// no validation rules for Password

	// no validation rules for Mode

	if len(errors) > 0 {
		return LoginPasswordResetReqMultiError(errors)
	}

	return nil
}

// LoginPasswordResetReqMultiError is an error wrapping multiple validation
// errors returned by LoginPasswordResetReq.ValidateAll() if the designated
// constraints aren't met.
type LoginPasswordResetReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginPasswordResetReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginPasswordResetReqMultiError) AllErrors() []error { return m }

// LoginPasswordResetReqValidationError is the validation error returned by
// LoginPasswordResetReq.Validate if the designated constraints aren't met.
type LoginPasswordResetReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginPasswordResetReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginPasswordResetReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginPasswordResetReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginPasswordResetReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginPasswordResetReqValidationError) ErrorName() string {
	return "LoginPasswordResetReqValidationError"
}

// Error satisfies the builtin error interface
func (e LoginPasswordResetReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginPasswordResetReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginPasswordResetReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginPasswordResetReqValidationError{}

// Validate checks the field values on SendPhoneCodeReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendPhoneCodeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendPhoneCodeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendPhoneCodeReqMultiError, or nil if none found.
func (m *SendPhoneCodeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendPhoneCodeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Phone

	// no validation rules for Template

	if len(errors) > 0 {
		return SendPhoneCodeReqMultiError(errors)
	}

	return nil
}

// SendPhoneCodeReqMultiError is an error wrapping multiple validation errors
// returned by SendPhoneCodeReq.ValidateAll() if the designated constraints
// aren't met.
type SendPhoneCodeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendPhoneCodeReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendPhoneCodeReqMultiError) AllErrors() []error { return m }

// SendPhoneCodeReqValidationError is the validation error returned by
// SendPhoneCodeReq.Validate if the designated constraints aren't met.
type SendPhoneCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendPhoneCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendPhoneCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendPhoneCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendPhoneCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendPhoneCodeReqValidationError) ErrorName() string { return "SendPhoneCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e SendPhoneCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendPhoneCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendPhoneCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendPhoneCodeReqValidationError{}

// Validate checks the field values on SendEmailCodeReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendEmailCodeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendEmailCodeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendEmailCodeReqMultiError, or nil if none found.
func (m *SendEmailCodeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendEmailCodeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Template

	if len(errors) > 0 {
		return SendEmailCodeReqMultiError(errors)
	}

	return nil
}

// SendEmailCodeReqMultiError is an error wrapping multiple validation errors
// returned by SendEmailCodeReq.ValidateAll() if the designated constraints
// aren't met.
type SendEmailCodeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendEmailCodeReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendEmailCodeReqMultiError) AllErrors() []error { return m }

// SendEmailCodeReqValidationError is the validation error returned by
// SendEmailCodeReq.Validate if the designated constraints aren't met.
type SendEmailCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendEmailCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendEmailCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendEmailCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendEmailCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendEmailCodeReqValidationError) ErrorName() string { return "SendEmailCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e SendEmailCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendEmailCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendEmailCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendEmailCodeReqValidationError{}

// Validate checks the field values on GetCosSessionKeyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCosSessionKeyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCosSessionKeyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCosSessionKeyReplyMultiError, or nil if none found.
func (m *GetCosSessionKeyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCosSessionKeyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TmpSecretId

	// no validation rules for TmpSecretKey

	// no validation rules for SessionToken

	// no validation rules for StartTime

	// no validation rules for ExpiredTime

	if len(errors) > 0 {
		return GetCosSessionKeyReplyMultiError(errors)
	}

	return nil
}

// GetCosSessionKeyReplyMultiError is an error wrapping multiple validation
// errors returned by GetCosSessionKeyReply.ValidateAll() if the designated
// constraints aren't met.
type GetCosSessionKeyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCosSessionKeyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCosSessionKeyReplyMultiError) AllErrors() []error { return m }

// GetCosSessionKeyReplyValidationError is the validation error returned by
// GetCosSessionKeyReply.Validate if the designated constraints aren't met.
type GetCosSessionKeyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCosSessionKeyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCosSessionKeyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCosSessionKeyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCosSessionKeyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCosSessionKeyReplyValidationError) ErrorName() string {
	return "GetCosSessionKeyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetCosSessionKeyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCosSessionKeyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCosSessionKeyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCosSessionKeyReplyValidationError{}

// Validate checks the field values on GetUserProfileReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserProfileReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserProfileReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserProfileReplyMultiError, or nil if none found.
func (m *GetUserProfileReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserProfileReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Username

	// no validation rules for Avatar

	// no validation rules for School

	// no validation rules for Company

	// no validation rules for Job

	// no validation rules for Homepage

	// no validation rules for Introduce

	if len(errors) > 0 {
		return GetUserProfileReplyMultiError(errors)
	}

	return nil
}

// GetUserProfileReplyMultiError is an error wrapping multiple validation
// errors returned by GetUserProfileReply.ValidateAll() if the designated
// constraints aren't met.
type GetUserProfileReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserProfileReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserProfileReplyMultiError) AllErrors() []error { return m }

// GetUserProfileReplyValidationError is the validation error returned by
// GetUserProfileReply.Validate if the designated constraints aren't met.
type GetUserProfileReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserProfileReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserProfileReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserProfileReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserProfileReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserProfileReplyValidationError) ErrorName() string {
	return "GetUserProfileReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserProfileReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserProfileReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserProfileReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserProfileReplyValidationError{}

// Validate checks the field values on GetProfileUpdateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProfileUpdateReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileUpdateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileUpdateReplyMultiError, or nil if none found.
func (m *GetProfileUpdateReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileUpdateReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Avatar

	// no validation rules for School

	// no validation rules for Company

	// no validation rules for Job

	// no validation rules for Homepage

	// no validation rules for Introduce

	// no validation rules for Status

	if len(errors) > 0 {
		return GetProfileUpdateReplyMultiError(errors)
	}

	return nil
}

// GetProfileUpdateReplyMultiError is an error wrapping multiple validation
// errors returned by GetProfileUpdateReply.ValidateAll() if the designated
// constraints aren't met.
type GetProfileUpdateReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileUpdateReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileUpdateReplyMultiError) AllErrors() []error { return m }

// GetProfileUpdateReplyValidationError is the validation error returned by
// GetProfileUpdateReply.Validate if the designated constraints aren't met.
type GetProfileUpdateReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileUpdateReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileUpdateReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileUpdateReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileUpdateReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileUpdateReplyValidationError) ErrorName() string {
	return "GetProfileUpdateReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileUpdateReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileUpdateReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileUpdateReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileUpdateReplyValidationError{}

// Validate checks the field values on SetProfileUpdateReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetProfileUpdateReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetProfileUpdateReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetProfileUpdateReqMultiError, or nil if none found.
func (m *SetProfileUpdateReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SetProfileUpdateReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for School

	// no validation rules for Company

	// no validation rules for Job

	// no validation rules for Homepage

	// no validation rules for Introduce

	if len(errors) > 0 {
		return SetProfileUpdateReqMultiError(errors)
	}

	return nil
}

// SetProfileUpdateReqMultiError is an error wrapping multiple validation
// errors returned by SetProfileUpdateReq.ValidateAll() if the designated
// constraints aren't met.
type SetProfileUpdateReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetProfileUpdateReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetProfileUpdateReqMultiError) AllErrors() []error { return m }

// SetProfileUpdateReqValidationError is the validation error returned by
// SetProfileUpdateReq.Validate if the designated constraints aren't met.
type SetProfileUpdateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetProfileUpdateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetProfileUpdateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetProfileUpdateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetProfileUpdateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetProfileUpdateReqValidationError) ErrorName() string {
	return "SetProfileUpdateReqValidationError"
}

// Error satisfies the builtin error interface
func (e SetProfileUpdateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetProfileUpdateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetProfileUpdateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetProfileUpdateReqValidationError{}
