// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
)

type sessionErrCtxKey struct{}

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
