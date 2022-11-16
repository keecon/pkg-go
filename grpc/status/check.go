// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package status

import (
	"google.golang.org/grpc/codes"
)

// IsOK return true if Status.Code equals codes.OK, else false
func IsOK(err error) bool {
	return check(err, codes.OK)
}

// IsCanceled return true if Status.Code equals codes.Canceled, else false
func IsCanceled(err error) bool {
	return check(err, codes.Canceled)
}

// IsInvalidArgument return true if Status.Code equals codes.InvalidArgument, else false
func IsInvalidArgument(err error) bool {
	return check(err, codes.InvalidArgument)
}

// IsDeadlineExceeded return true if Status.Code equals codes.DeadlineExceeded, else false
func IsDeadlineExceeded(err error) bool {
	return check(err, codes.DeadlineExceeded)
}

// IsNotFound return true if Status.Code equals codes.NotFound, else false
func IsNotFound(err error) bool {
	return check(err, codes.NotFound)
}

// IsAlreadyExists return true if Status.Code equals codes.AlreadyExists, else false
func IsAlreadyExists(err error) bool {
	return check(err, codes.AlreadyExists)
}

// IsPermissionDenied return true if Status.Code equals codes.PermissionDenied, else false
func IsPermissionDenied(err error) bool {
	return check(err, codes.PermissionDenied)
}

// IsFailedPrecondition return true if Status.Code equals codes.FailedPrecondition, else false
func IsFailedPrecondition(err error) bool {
	return check(err, codes.FailedPrecondition)
}

// IsAborted return true if Status.Code equals codes.Aborted, else false
func IsAborted(err error) bool {
	return check(err, codes.Aborted)
}

// IsResourceExhausted return true if Status.Code equals codes.ResourceExhausted, else false
func IsResourceExhausted(err error) bool {
	return check(err, codes.ResourceExhausted)
}

// IsUnauthenticated return true if Status.Code equals codes.Unauthenticated, else false
func IsUnauthenticated(err error) bool {
	return check(err, codes.Unauthenticated)
}

// IsOutOfRange return true if Status.Code equals codes.OutOfRange, else false
func IsOutOfRange(err error) bool {
	return check(err, codes.OutOfRange)
}

// IsInternal return true if Status.Code equals codes.Internal, else false
func IsInternal(err error) bool {
	return check(err, codes.Internal)
}

// IsUnknown return true if Status.Code equals codes.Unknown, else false
func IsUnknown(err error) bool {
	return check(err, codes.Unknown)
}

// IsUnimplemented return true if Status.Code equals codes.Unimplemented, else false
func IsUnimplemented(err error) bool {
	return check(err, codes.Unimplemented)
}

// IsUnavailable return true if Status.Code equals codes.Unavailable, else false
func IsUnavailable(err error) bool {
	return check(err, codes.Unavailable)
}

// IsDataLoss return true if Status.Code equals codes.DataLoss, else false
func IsDataLoss(err error) bool {
	return check(err, codes.DataLoss)
}

func check(err error, code codes.Code) bool {
	return Code(err) == code
}
