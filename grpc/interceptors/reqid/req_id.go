// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reqid sets unique id info into the request.
package reqid

import (
	"context"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/xid"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptor that sets the values for request unique id.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return handler(wrapPeerContext(ctx), req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that sets the values for request unique id.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		nss := middleware.WrapServerStream(stream)
		nss.WrappedContext = wrapPeerContext(stream.Context())
		return handler(srv, nss)
	}
}

func wrapPeerContext(ctx context.Context) context.Context {
	reqID := xid.New()
	return logging.InjectFields(ctx, logging.Fields{"req.id", reqID.String()})
}
