// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: github.com/xdu31/exercise-service/pkg/server/pb/service.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on GetExercisesForMuscleReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExercisesForMuscleReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MuscleName

	return nil
}

// GetExercisesForMuscleReqValidationError is the validation error returned by
// GetExercisesForMuscleReq.Validate if the designated constraints aren't met.
type GetExercisesForMuscleReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExercisesForMuscleReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExercisesForMuscleReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExercisesForMuscleReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExercisesForMuscleReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExercisesForMuscleReqValidationError) ErrorName() string {
	return "GetExercisesForMuscleReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetExercisesForMuscleReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExercisesForMuscleReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExercisesForMuscleReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExercisesForMuscleReqValidationError{}

// Validate checks the field values on GetExercisesForMuscleResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExercisesForMuscleResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetExercisesForMuscleRespValidationError is the validation error returned by
// GetExercisesForMuscleResp.Validate if the designated constraints aren't met.
type GetExercisesForMuscleRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExercisesForMuscleRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExercisesForMuscleRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExercisesForMuscleRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExercisesForMuscleRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExercisesForMuscleRespValidationError) ErrorName() string {
	return "GetExercisesForMuscleRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetExercisesForMuscleRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExercisesForMuscleResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExercisesForMuscleRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExercisesForMuscleRespValidationError{}

// Validate checks the field values on GetMusclesForExerciseReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMusclesForExerciseReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ExerciseName

	return nil
}

// GetMusclesForExerciseReqValidationError is the validation error returned by
// GetMusclesForExerciseReq.Validate if the designated constraints aren't met.
type GetMusclesForExerciseReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMusclesForExerciseReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMusclesForExerciseReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMusclesForExerciseReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMusclesForExerciseReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMusclesForExerciseReqValidationError) ErrorName() string {
	return "GetMusclesForExerciseReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetMusclesForExerciseReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMusclesForExerciseReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMusclesForExerciseReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMusclesForExerciseReqValidationError{}

// Validate checks the field values on GetMusclesForExerciseResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMusclesForExerciseResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetMusclesForExerciseRespValidationError is the validation error returned by
// GetMusclesForExerciseResp.Validate if the designated constraints aren't met.
type GetMusclesForExerciseRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMusclesForExerciseRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMusclesForExerciseRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMusclesForExerciseRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMusclesForExerciseRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMusclesForExerciseRespValidationError) ErrorName() string {
	return "GetMusclesForExerciseRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetMusclesForExerciseRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMusclesForExerciseResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMusclesForExerciseRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMusclesForExerciseRespValidationError{}
