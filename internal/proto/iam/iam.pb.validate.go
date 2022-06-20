// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iam/iam.proto

package iampb

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

// Validate checks the field values on AuthResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthResponseMultiError, or
// nil if none found.
func (m *AuthResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for ErrorMessage

	// no validation rules for DomainId

	// no validation rules for ProjectId

	// no validation rules for GroupId

	// no validation rules for AccountId

	// no validation rules for AccountType

	// no validation rules for IssuedAt

	// no validation rules for ExpiresAt

	if len(errors) > 0 {
		return AuthResponseMultiError(errors)
	}

	return nil
}

// AuthResponseMultiError is an error wrapping multiple validation errors
// returned by AuthResponse.ValidateAll() if the designated constraints aren't met.
type AuthResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthResponseMultiError) AllErrors() []error { return m }

// AuthResponseValidationError is the validation error returned by
// AuthResponse.Validate if the designated constraints aren't met.
type AuthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthResponseValidationError) ErrorName() string { return "AuthResponseValidationError" }

// Error satisfies the builtin error interface
func (e AuthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthResponseValidationError{}

// Validate checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthRequestMultiError, or
// nil if none found.
func (m *AuthRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return AuthRequestMultiError(errors)
	}

	return nil
}

// AuthRequestMultiError is an error wrapping multiple validation errors
// returned by AuthRequest.ValidateAll() if the designated constraints aren't met.
type AuthRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRequestMultiError) AllErrors() []error { return m }

// AuthRequestValidationError is the validation error returned by
// AuthRequest.Validate if the designated constraints aren't met.
type AuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRequestValidationError) ErrorName() string { return "AuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRequestValidationError{}

// Validate checks the field values on VerifyTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyTokenRequestMultiError, or nil if none found.
func (m *VerifyTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return VerifyTokenRequestMultiError(errors)
	}

	return nil
}

// VerifyTokenRequestMultiError is an error wrapping multiple validation errors
// returned by VerifyTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type VerifyTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyTokenRequestMultiError) AllErrors() []error { return m }

// VerifyTokenRequestValidationError is the validation error returned by
// VerifyTokenRequest.Validate if the designated constraints aren't met.
type VerifyTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyTokenRequestValidationError) ErrorName() string {
	return "VerifyTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyTokenRequestValidationError{}

// Validate checks the field values on VerifyTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyTokenResponseMultiError, or nil if none found.
func (m *VerifyTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for ErrorMessage

	// no validation rules for DomainId

	// no validation rules for ProjectId

	// no validation rules for GroupId

	// no validation rules for AccountId

	// no validation rules for AccountType

	// no validation rules for IssuedAt

	// no validation rules for ExpiresAt

	if len(errors) > 0 {
		return VerifyTokenResponseMultiError(errors)
	}

	return nil
}

// VerifyTokenResponseMultiError is an error wrapping multiple validation
// errors returned by VerifyTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type VerifyTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyTokenResponseMultiError) AllErrors() []error { return m }

// VerifyTokenResponseValidationError is the validation error returned by
// VerifyTokenResponse.Validate if the designated constraints aren't met.
type VerifyTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyTokenResponseValidationError) ErrorName() string {
	return "VerifyTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyTokenResponseValidationError{}

// Validate checks the field values on CreateTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTokenRequestMultiError, or nil if none found.
func (m *CreateTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDomainId() <= 0 {
		err := CreateTokenRequestValidationError{
			field:  "DomainId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProjectId() <= 0 {
		err := CreateTokenRequestValidationError{
			field:  "ProjectId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetGroupId() <= 0 {
		err := CreateTokenRequestValidationError{
			field:  "GroupId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAccountId() <= 0 {
		err := CreateTokenRequestValidationError{
			field:  "AccountId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for AccountType

	if len(errors) > 0 {
		return CreateTokenRequestMultiError(errors)
	}

	return nil
}

// CreateTokenRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenRequestMultiError) AllErrors() []error { return m }

// CreateTokenRequestValidationError is the validation error returned by
// CreateTokenRequest.Validate if the designated constraints aren't met.
type CreateTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenRequestValidationError) ErrorName() string {
	return "CreateTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenRequestValidationError{}

// Validate checks the field values on CreateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTokenResponseMultiError, or nil if none found.
func (m *CreateTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return CreateTokenResponseMultiError(errors)
	}

	return nil
}

// CreateTokenResponseMultiError is an error wrapping multiple validation
// errors returned by CreateTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenResponseMultiError) AllErrors() []error { return m }

// CreateTokenResponseValidationError is the validation error returned by
// CreateTokenResponse.Validate if the designated constraints aren't met.
type CreateTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenResponseValidationError) ErrorName() string {
	return "CreateTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenResponseValidationError{}