// Copyright 2024 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// container package provides a generic container interface
package container

// Container is a generic container interface
type Container[T comparable] interface {
	// Size returns the number of elements in the container
	Size() int

	// IsEmpty returns true if the container is empty, false otherwise
	IsEmpty() bool

	// Values returns all elements in the container
	Values() []T

	// Add adds a value to the container
	Add(value T) error

	// Remove removes a value from the container
	Remove(value T) error

	// Clear removes all elements from the container
	Clear()
}
