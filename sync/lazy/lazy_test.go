// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazyCallOnce(t *testing.T) {
	// given
	var callCount int
	fn := func() int {
		callCount++
		return 1
	}
	lazy := New(fn)

	// when
	ret1 := lazy()
	ret2 := lazy()

	// then
	assert.Equal(t, callCount, 1)
	assert.Equal(t, ret1, ret2)
}
