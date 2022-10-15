// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
)

type (
	sessionCtxKey    struct{}
	sessionErrCtxKey struct{}
)

// SessionFromContext returns authenticated session info.
func SessionFromContext(ctx context.Context) any {
	return ctx.Value(sessionCtxKey{})
}

func newSessionContext(ctx context.Context, session any) context.Context {
	return context.WithValue(ctx, sessionCtxKey{}, session)
}

func errFromContext(ctx context.Context) error {
	v := ctx.Value(sessionErrCtxKey{})
	if v == nil {
		return nil
	}

	return v.(error)
}

func newErrContext(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, sessionErrCtxKey{}, err)
}
