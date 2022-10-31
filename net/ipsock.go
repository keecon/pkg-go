// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"fmt"
	"net"
	"strconv"
)

// SplitHostPortNum splits a network address of the form "host:port",
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
// host%zone and port.
func SplitHostPortNum(addr string) (string, int, error) {
	host, portstr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid net addr [%s]: %w", addr, err)
	}
	port, err := strconv.Atoi(portstr)
	if err != nil {
		return "", 0, fmt.Errorf("malformed port [%s]: %w", addr, err)
	}
	return host, port, nil
}

// JoinHostPortNum combines host and port into a network address of the
// form "host:port". If host contains a colon, as found in literal
// IPv6 addresses, then JoinHostPort returns "[host]:port".
func JoinHostPortNum(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}
