// Copyright 2024 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"github.com/keecon/pkg-go/container"
)

// List is a generic list interface
type List[T comparable] interface {
	container.Container[T]

	// Get returns the value at the specified index
	Get(index int) (T, error)

	// Set sets the value at the specified index
	Set(index int, value T) error

	// InsertAt inserts a value at the specified index
	InsertAt(index int, value T) error

	// RemoveAt removes the value at the specified index
	RemoveAt(index int) (T, error)
}
