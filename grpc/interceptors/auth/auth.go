// Copyright 2022 Keecon Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package auth extends github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth
package auth

import (
	"context"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/keecon/pkg-go/grpc/status"
	"google.golang.org/grpc"
)

// Session defines authorized session identifier
type Session interface {
	ID() string
}

// Decode decoding session.
type Decode func(token string) (Session, error)

// UnaryServerInterceptor returns a new unary server interceptor that sets the values for request tags.
func UnaryServerInterceptor(expectedScheme string, decode Decode) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		newCtx, err := wrapSessionContext(ctx, expectedScheme, decode)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that sets the values for request tags.
func StreamServerInterceptor(expectedScheme string, decode Decode) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		nss := middleware.WrapServerStream(stream)
		nss.WrappedContext, err = wrapSessionContext(stream.Context(), expectedScheme, decode)
		if err != nil {
			return err
		}
		return handler(srv, nss)
	}
}

func wrapSessionContext(ctx context.Context, expectedScheme string, decode Decode) (context.Context, error) {
	if token, err := auth.AuthFromMD(ctx, expectedScheme); err == nil {
		session, err := decode(token)
		if err != nil {
			return nil, status.Unauthenticated("authenticated decode error: %v", err).Err()
		}

		logCtx := logging.InjectFields(ctx, logging.Fields{
			"session.id", session.ID(),
		})
		ctx = contextWithSession(logCtx, session)
	}
	return ctx, nil
}

type sessionCtxKey struct{}

// SessionFromContext returns session info from context
func SessionFromContext(ctx context.Context) Session {
	v := ctx.Value(sessionCtxKey{})
	if v == nil {
		return nil
	}

	return v.(Session)
}

func contextWithSession(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionCtxKey{}, session)
}
