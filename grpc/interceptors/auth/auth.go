// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package auth extends github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth
package auth

import (
	"context"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"
)

// ServiceAuthFunc performs authentication per service.
type ServiceAuthFunc interface {
	AuthFunc(ctx context.Context, fullMethodName string) (context.Context, error)
}

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request auth.
func UnaryServerInterceptor() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
			if srv, ok := info.Server.(ServiceAuthFunc); ok {
				if ctx, err = srv.AuthFunc(ctx, info.FullMethod); err != nil {
					ctx = newErrContext(ctx, err)
				}
			}
			return handler(ctx, req)
		}, func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
			if err := errFromContext(ctx); err != nil {
				return nil, err
			}
			return handler(ctx, req)
		},
	}
}

// StreamServerInterceptor returns a new stream server interceptors that performs per-request auth.
func StreamServerInterceptor() []grpc.StreamServerInterceptor {
	return []grpc.StreamServerInterceptor{
		func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
			ctx := stream.Context()
			if srv, ok := srv.(ServiceAuthFunc); ok {
				if ctx, err = srv.AuthFunc(ctx, info.FullMethod); err != nil {
					ctx = newErrContext(ctx, err)
				}
			}
			wrapped := middleware.WrapServerStream(stream)
			wrapped.WrappedContext = ctx
			return handler(srv, wrapped)
		},
		func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if err := errFromContext(stream.Context()); err != nil {
				return err
			}
			return handler(srv, stream)
		},
	}
}
