// Copyright 2022 Keecon Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ctxerr implements various error handling
package ctxerr

import (
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor returns a new unary server interceptor that returns a grpc error from context error.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		return resp, fromContextError(err)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that returns a grpc error from context error.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, stream)
		return fromContextError(err)
	}
}

func fromContextError(err error) error {
	switch {
	case err == nil:
		return nil

	case errors.Is(err, context.DeadlineExceeded):
		return status.New(codes.DeadlineExceeded, err.Error()).Err()

	case errors.Is(err, context.Canceled):
		return status.New(codes.Canceled, err.Error()).Err()
	}

	if _, ok := err.(interface{ GRPCStatus() *status.Status }); ok {
		return err
	}
	if etype, ok := err.(net.Error); ok && etype.Timeout() {
		return status.New(codes.DeadlineExceeded, err.Error()).Err()
	}
	return status.New(codes.Unknown, err.Error()).Err()
}
