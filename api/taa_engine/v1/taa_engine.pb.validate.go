// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/taa_engine/v1/taa_engine.proto

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

// Validate checks the field values on TaaEngineInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TaaEngineInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TaaEngineInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TaaEngineInfoMultiError, or
// nil if none found.
func (m *TaaEngineInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *TaaEngineInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return TaaEngineInfoMultiError(errors)
	}

	return nil
}

// TaaEngineInfoMultiError is an error wrapping multiple validation errors
// returned by TaaEngineInfo.ValidateAll() if the designated constraints
// aren't met.
type TaaEngineInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaaEngineInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaaEngineInfoMultiError) AllErrors() []error { return m }

// TaaEngineInfoValidationError is the validation error returned by
// TaaEngineInfo.Validate if the designated constraints aren't met.
type TaaEngineInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaaEngineInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaaEngineInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaaEngineInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaaEngineInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaaEngineInfoValidationError) ErrorName() string { return "TaaEngineInfoValidationError" }

// Error satisfies the builtin error interface
func (e TaaEngineInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaaEngineInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaaEngineInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaaEngineInfoValidationError{}

// Validate checks the field values on CreateTaaEngineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTaaEngineRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaaEngineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaaEngineRequestMultiError, or nil if none found.
func (m *CreateTaaEngineRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaaEngineRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return CreateTaaEngineRequestMultiError(errors)
	}

	return nil
}

// CreateTaaEngineRequestMultiError is an error wrapping multiple validation
// errors returned by CreateTaaEngineRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateTaaEngineRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaaEngineRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaaEngineRequestMultiError) AllErrors() []error { return m }

// CreateTaaEngineRequestValidationError is the validation error returned by
// CreateTaaEngineRequest.Validate if the designated constraints aren't met.
type CreateTaaEngineRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaaEngineRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaaEngineRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaaEngineRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaaEngineRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaaEngineRequestValidationError) ErrorName() string {
	return "CreateTaaEngineRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaaEngineRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaaEngineRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaaEngineRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaaEngineRequestValidationError{}

// Validate checks the field values on CreateTaaEngineResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTaaEngineResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaaEngineResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaaEngineResponseMultiError, or nil if none found.
func (m *CreateTaaEngineResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaaEngineResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTaaEngineResponseValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTaaEngineResponseValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaaEngineResponseValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTaaEngineResponseMultiError(errors)
	}

	return nil
}

// CreateTaaEngineResponseMultiError is an error wrapping multiple validation
// errors returned by CreateTaaEngineResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateTaaEngineResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaaEngineResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaaEngineResponseMultiError) AllErrors() []error { return m }

// CreateTaaEngineResponseValidationError is the validation error returned by
// CreateTaaEngineResponse.Validate if the designated constraints aren't met.
type CreateTaaEngineResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaaEngineResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaaEngineResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaaEngineResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaaEngineResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaaEngineResponseValidationError) ErrorName() string {
	return "CreateTaaEngineResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaaEngineResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaaEngineResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaaEngineResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaaEngineResponseValidationError{}