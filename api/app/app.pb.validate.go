// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: app/app.proto

package app

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

// Validate checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Application) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ApplicationMultiError, or
// nil if none found.
func (m *Application) ValidateAll() error {
	return m.validate(true)
}

func (m *Application) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetPid()) != 24 {
		err := ApplicationValidationError{
			field:  "Pid",
			reason: "value length must be 24 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if l := utf8.RuneCountInString(m.GetEnv()); l < 3 || l > 16 {
		err := ApplicationValidationError{
			field:  "Env",
			reason: "value length must be between 3 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	if l := utf8.RuneCountInString(m.GetImage()); l < 3 || l > 128 {
		err := ApplicationValidationError{
			field:  "Image",
			reason: "value length must be between 3 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetPorts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("Ports[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("Ports[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationValidationError{
					field:  fmt.Sprintf("Ports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Status

	// no validation rules for Size

	for idx, item := range m.GetConfigs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("Configs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("Configs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationValidationError{
					field:  fmt.Sprintf("Configs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "CreatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "UpdatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ApplicationMultiError(errors)
	}

	return nil
}

// ApplicationMultiError is an error wrapping multiple validation errors
// returned by Application.ValidateAll() if the designated constraints aren't met.
type ApplicationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplicationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplicationMultiError) AllErrors() []error { return m }

// ApplicationValidationError is the validation error returned by
// Application.Validate if the designated constraints aren't met.
type ApplicationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationValidationError) ErrorName() string { return "ApplicationValidationError" }

// Error satisfies the builtin error interface
func (e ApplicationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pid

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListResponseMultiError, or
// nil if none found.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListResponseMultiError(errors)
	}

	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on AppOpRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AppOpRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppOpRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppOpRequestMultiError, or
// nil if none found.
func (m *AppOpRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppOpRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Op

	if len(errors) > 0 {
		return AppOpRequestMultiError(errors)
	}

	return nil
}

// AppOpRequestMultiError is an error wrapping multiple validation errors
// returned by AppOpRequest.ValidateAll() if the designated constraints aren't met.
type AppOpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppOpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppOpRequestMultiError) AllErrors() []error { return m }

// AppOpRequestValidationError is the validation error returned by
// AppOpRequest.Validate if the designated constraints aren't met.
type AppOpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppOpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppOpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppOpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppOpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppOpRequestValidationError) ErrorName() string { return "AppOpRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppOpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppOpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppOpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppOpRequestValidationError{}

// Validate checks the field values on AppOpReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AppOpReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppOpReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppOpReplyMultiError, or
// nil if none found.
func (m *AppOpReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AppOpReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AppOpReplyMultiError(errors)
	}

	return nil
}

// AppOpReplyMultiError is an error wrapping multiple validation errors
// returned by AppOpReply.ValidateAll() if the designated constraints aren't met.
type AppOpReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppOpReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppOpReplyMultiError) AllErrors() []error { return m }

// AppOpReplyValidationError is the validation error returned by
// AppOpReply.Validate if the designated constraints aren't met.
type AppOpReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppOpReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppOpReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppOpReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppOpReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppOpReplyValidationError) ErrorName() string { return "AppOpReplyValidationError" }

// Error satisfies the builtin error interface
func (e AppOpReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppOpReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppOpReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppOpReplyValidationError{}

// Validate checks the field values on AppStatusRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AppStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppStatusRequestMultiError, or nil if none found.
func (m *AppStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	if len(errors) > 0 {
		return AppStatusRequestMultiError(errors)
	}

	return nil
}

// AppStatusRequestMultiError is an error wrapping multiple validation errors
// returned by AppStatusRequest.ValidateAll() if the designated constraints
// aren't met.
type AppStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppStatusRequestMultiError) AllErrors() []error { return m }

// AppStatusRequestValidationError is the validation error returned by
// AppStatusRequest.Validate if the designated constraints aren't met.
type AppStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppStatusRequestValidationError) ErrorName() string { return "AppStatusRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppStatusRequestValidationError{}

// Validate checks the field values on AppStatusReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AppStatusReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppStatusReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppStatusReplyMultiError,
// or nil if none found.
func (m *AppStatusReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AppStatusReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AppStatusReplyMultiError(errors)
	}

	return nil
}

// AppStatusReplyMultiError is an error wrapping multiple validation errors
// returned by AppStatusReply.ValidateAll() if the designated constraints
// aren't met.
type AppStatusReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppStatusReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppStatusReplyMultiError) AllErrors() []error { return m }

// AppStatusReplyValidationError is the validation error returned by
// AppStatusReply.Validate if the designated constraints aren't met.
type AppStatusReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppStatusReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppStatusReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppStatusReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppStatusReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppStatusReplyValidationError) ErrorName() string { return "AppStatusReplyValidationError" }

// Error satisfies the builtin error interface
func (e AppStatusReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppStatusReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppStatusReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppStatusReplyValidationError{}
