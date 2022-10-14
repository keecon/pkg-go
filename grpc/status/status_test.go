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
	dataset := []struct {
		err error
		ok  bool
	}{
		{NotFound("NotFound").Err(), true},
		{OK().Err(), false},
		{InvalidArgument("InvalidArgument").Err(), false},
		{AlreadyExists("AlreadyExists").Err(), false},
		{Aborted("Aborted").Err(), false},
		{FailedPrecondition("FailedPrecondition").Err(), false},
		{ResourceExhausted("ResourceExhausted").Err(), false},
		{Internal("Internal").Err(), false},
		{Unimplemented("unimplemented").Err(), false},
	}

	// table driven tests
	for _, v := range dataset {
		// given
		err := v.err
		expected := v.ok

		// when
		ok := IsNotFound(err)

		// then
		assert.Equal(t, expected, ok)

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
