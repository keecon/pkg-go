// Copyright 2022 Keecon Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lazy implements lazy initialize value
package lazy

import "sync"

// New implements lazy initialize value
func New[T any](fn func() T) func() T {
	var (
		once  sync.Once
		value T
	)
	return func() T {
		once.Do(func() {
			value = fn()
		})
		return value
	}
}
