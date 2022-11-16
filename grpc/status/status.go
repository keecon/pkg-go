// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package status extends google.golang.org/grpc/status
package status

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Status references google.golang.org/grpc/internal/status. It represents an
// RPC status code, message, and details.  It is immutable and should be
// created with New, Newf, or FromProto.
// https://godoc.org/google.golang.org/grpc/internal/status
type Status = status.Status

var (
	// New returns a Status representing c and msg.
	New = status.New

	// Newf returns New(c, fmt.Sprintf(format, a...)).
	Newf = status.Newf

	// Error returns an error representing c and msg.  If c is OK, returns nil.
	Error = status.Error

	// Errorf returns Error(c, fmt.Sprintf(format, a...)).
	Errorf = status.Errorf

	// ErrorProto returns an error representing s.  If s.Code is OK, returns nil.
	ErrorProto = status.ErrorProto

	// FromProto returns a Status representing s.
	FromProto = status.FromProto

	// FromError returns a Status representing err if it was produced from this
	// package or has a method `GRPCStatus() *Status`. Otherwise, ok is false and a
	// Status is returned with codes.Unknown and the original error message.
	FromError = status.FromError

	// Convert is a convenience function which removes the need to handle the
	// boolean return value from FromError.
	Convert = status.Convert

	// Code returns the Code of the error if it is a Status error, codes.OK if err
	// is nil, or codes.Unknown otherwise.
	Code = status.Code

	// FromContextError converts a context error into a Status.  It returns a
	// Status with codes.OK if err is nil, or a Status with codes.Unknown if err is
	// non-nil and not a context error.
	FromContextError = status.FromContextError
)

// OK returns a Status Ok
func OK() *Status {
	return New(codes.OK, "")
}

// Canceled returns a Status Canceled
func Canceled(format string, a ...any) *Status {
	return Newf(codes.Canceled, format, a...)
}

// InvalidArgument returns a Status InvalidArgument
func InvalidArgument(format string, a ...any) *Status {
	return Newf(codes.InvalidArgument, format, a...)
}

// DeadlineExceeded returns a Status DeadlineExceeded
func DeadlineExceeded(format string, a ...any) *Status {
	return Newf(codes.DeadlineExceeded, format, a...)
}

// NotFound returns a Status NotFound
func NotFound(format string, a ...any) *Status {
	return Newf(codes.NotFound, format, a...)
}

// AlreadyExists returns a Status AlreadyExists
func AlreadyExists(format string, a ...any) *Status {
	return Newf(codes.AlreadyExists, format, a...)
}

// PermissionDenied returns a Status PermissionDenied
func PermissionDenied(format string, a ...any) *Status {
	return Newf(codes.PermissionDenied, format, a...)
}

// FailedPrecondition returns a Status FailedPrecondition
func FailedPrecondition(format string, a ...any) *Status {
	return Newf(codes.FailedPrecondition, format, a...)
}

// Aborted returns a Status Aborted
func Aborted(format string, a ...any) *Status {
	return Newf(codes.Aborted, format, a...)
}

// ResourceExhausted returns a Status ResourceExhausted
func ResourceExhausted(format string, a ...any) *Status {
	return Newf(codes.ResourceExhausted, format, a...)
}

// Unauthenticated returns a Status Unauthenticated
func Unauthenticated(format string, a ...any) *Status {
	return Newf(codes.Unauthenticated, format, a...)
}

// OutOfRange returns a Status OutOfRange
func OutOfRange(format string, a ...any) *Status {
	return Newf(codes.OutOfRange, format, a...)
}

// Internal returns a Status Internal
func Internal(format string, a ...any) *Status {
	return Newf(codes.Internal, format, a...)
}

// Unknown returns a Status Unknown
func Unknown(format string, a ...any) *Status {
	return Newf(codes.Unknown, format, a...)
}

// Unimplemented returns a Status Unimplemented
func Unimplemented(format string, a ...any) *Status {
	return Newf(codes.Unimplemented, format, a...)
}

// Unavailable returns a Status Unavailable
func Unavailable(format string, a ...any) *Status {
	return Newf(codes.Unavailable, format, a...)
}

// DataLoss returns a Status DataLoss
func DataLoss(format string, a ...any) *Status {
	return Newf(codes.DataLoss, format, a...)
}
