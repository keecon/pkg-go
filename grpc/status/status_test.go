// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func TestCheckNotFound(t *testing.T) {
	// dataset
	allErrors := []*Status{
		OK(),
		Canceled(codes.Canceled.String()),
		InvalidArgument(codes.InvalidArgument.String()),
		DeadlineExceeded(codes.DeadlineExceeded.String()),
		NotFound(codes.NotFound.String()),
		AlreadyExists(codes.AlreadyExists.String()),
		PermissionDenied(codes.PermissionDenied.String()),
		FailedPrecondition(codes.FailedPrecondition.String()),
		Aborted(codes.Aborted.String()),
		ResourceExhausted(codes.ResourceExhausted.String()),
		Unauthenticated(codes.Unauthenticated.String()),
		OutOfRange(codes.OutOfRange.String()),
		Internal(codes.Internal.String()),
		Unknown(codes.Unknown.String()),
		Unimplemented(codes.Unimplemented.String()),
		Unavailable(codes.Unavailable.String()),
		DataLoss(codes.DataLoss.String()),
	}
	allCheckers := []func(error) bool{
		IsOK,
		IsCanceled,
		IsInvalidArgument,
		IsDeadlineExceeded,
		IsNotFound,
		IsAlreadyExists,
		IsPermissionDenied,
		IsFailedPrecondition,
		IsAborted,
		IsResourceExhausted,
		IsUnauthenticated,
		IsOutOfRange,
		IsInternal,
		IsUnknown,
		IsUnimplemented,
		IsUnavailable,
		IsDataLoss,
	}

	// table driven tests
	for _, checker := range allCheckers {
		var checkCount int
		for _, err := range allErrors {
			if checker(err.Err()) {
				checkCount++
			}
		}

		assert.Equal(t, 1, checkCount)
	}
}

func TestCheckErrorDetailString(t *testing.T) {
	// dataset
	dataset := []struct {
		code codes.Code
		msg  string
		fn   func(string, ...any) *Status
	}{
		{
			code: codes.InvalidArgument,
			fn:   InvalidArgument,
		},
		{
			code: codes.NotFound,
			fn:   NotFound,
		},
		{
			code: codes.AlreadyExists,
			fn:   AlreadyExists,
		},
		{
			code: codes.Aborted,
			fn:   Aborted,
		},
		{
			code: codes.FailedPrecondition,
			fn:   FailedPrecondition,
		},
		{
			code: codes.Internal,
			fn:   Internal,
		},
		{
			code: codes.Unimplemented,
			fn:   Unimplemented,
		},
	}

	// table driven tests
	for _, v := range dataset {
		// given
		errDetail := &errdetails.ErrorInfo{Reason: "test"}
		errstr := fmt.Sprintf("rpc error: code = %s desc = %s", v.code, v.code)

		// when
		err := ErrorDetailString(v.fn(v.code.String()).Err())
		serr, _ := v.fn(v.code.String()).WithDetails(errDetail)
		errWithDetail := ErrorDetailString(serr.Err())

		// then
		assert.Equal(t, errstr, err)
		assert.NotEqual(t, errstr, errWithDetail)
	}
}
