// Copyright 2022 Keecon Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package timeout implements a timeout error after the given duration.
package timeout

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"google.golang.org/grpc"
)

// UnaryClientInterceptor returns a new unary client interceptor that returns a grpc error from context error.
var UnaryClientInterceptor = timeout.TimeoutUnaryClientInterceptor

// UnaryServerInterceptor returns a new unary server interceptor that returns a grpc error from context error.
func UnaryServerInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		newCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return handler(newCtx, req)
	}
}
