// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package status

import (
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type (
	// RetryInfo describes when the clients can retry a failed request.
	RetryInfo = errdetails.RetryInfo

	// DebugInfo describes additional debugging info.
	DebugInfo = errdetails.DebugInfo

	// QuotaFailure describes how a quota check failed.
	QuotaFailure = errdetails.QuotaFailure

	// ErrorInfo describes the cause of the error with structured details.
	ErrorInfo = errdetails.ErrorInfo

	// PreconditionFailure describes what preconditions have failed.
	PreconditionFailure = errdetails.PreconditionFailure

	// BadRequest describes violations in a client request.
	BadRequest = errdetails.BadRequest

	// RequestInfo contains metadata about the request that clients can attach when filing a bug
	// or providing other forms of feedback.
	RequestInfo = errdetails.RequestInfo

	// ResourceInfo describes the resource that is being accessed.
	ResourceInfo = errdetails.ResourceInfo

	// Help provides links to documentation or for performing an out of band action.
	Help = errdetails.Help

	// LocalizedMessage provides a localized error message that is safe to return to the user
	// which can be attached to an RPC error.
	LocalizedMessage = errdetails.LocalizedMessage
)

// ErrorDetailString returns err.Error() with details message if err is Status
func ErrorDetailString(err error) string {
	if err == nil {
		return ""
	}

	if s, ok := FromError(err); ok {
		var details []string
		for _, v := range s.Details() {
			details = append(details, fmt.Sprintf("{%v}", v))
		}

		if 0 < len(details) {
			return fmt.Sprintf("%s [%s]", err, strings.Join(details, ","))
		}
	}
	return err.Error()
}
