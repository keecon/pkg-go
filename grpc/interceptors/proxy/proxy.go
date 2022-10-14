// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package proxy applies reverse proxy info to peer info
package proxy

import (
	"context"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// UnaryServerInterceptor returns a new unary server interceptor that sets the values for request tags.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return handler(wrapPeerContext(ctx), req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that sets the values for request tags.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		nss := middleware.WrapServerStream(stream)
		nss.WrappedContext = wrapPeerContext(stream.Context())
		return handler(srv, nss)
	}
}

func wrapPeerContext(ctx context.Context) context.Context {
	md := metadata.ExtractIncoming(ctx)
	if addr := md.Get("x-forwarded-for"); addr != "" {
		peerInfo := &peer.Peer{}
		if p, ok := peer.FromContext(ctx); ok {
			peerInfo = &peer.Peer{
				AuthInfo: p.AuthInfo,
			}
		}

		peerInfo.Addr = &netAddr{
			addr: addr,
			port: md.Get("x-forwarded-for-port"),
		}
		return peer.NewContext(ctx, peerInfo)
	}
	return ctx
}

type netAddr struct {
	addr string
	port string
}

func (a *netAddr) Network() string {
	return "tcp"
}

func (a *netAddr) String() string {
	if a.port != "" {
		return a.addr + ":" + a.port
	}
	return a.addr
}
