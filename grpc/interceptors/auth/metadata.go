// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
)

// AuthFromMD is a helper function for extracting the :authorization header from the gRPC metadata of the request.
var AuthFromMD = auth.AuthFromMD
