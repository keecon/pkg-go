// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package net extends standard net package
package net

import "net"

// An IP is a single IP address, a slice of bytes.
type IP = net.IP

// IPv4 returns a list of unicast interface addresses.
func IPv4(opts ...Option) (ret []net.IP, err error) {
	ifaddrs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, ifaddr := range ifaddrs {
		addrs, err := ifaddr.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}

			ipv4 := ip.To4()
			if ipv4 != nil && isAllOk(ip, opts...) {
				ret = append(ret, ipv4)
			}
		}
	}
	return ret, nil
}

func isAllOk(ip net.IP, opts ...Option) bool {
	for _, fn := range opts {
		if !fn(ip) {
			return false
		}
	}
	return true
}

// Option predicates net.IP choice.
type Option func(ip net.IP) bool

// WithGlobalUnicast selects global unicast address.
func WithGlobalUnicast() Option {
	return func(ip net.IP) bool {
		return ip.IsGlobalUnicast()
	}
}

// WithPrivate selects private address.
func WithPrivate() Option {
	return func(ip net.IP) bool {
		return ip.IsPrivate()
	}
}

// WithLoopback selects loopback address.
func WithLoopback() Option {
	return func(ip net.IP) bool {
		return ip.IsLoopback()
	}
}

// WithMulticast selects multicast address.
func WithMulticast() Option {
	return func(ip net.IP) bool {
		return ip.IsMulticast()
	}
}

// WithCIDR selects an address included in CIDR.
func WithCIDR(cidr string) Option {
	_, ipnet, _ := net.ParseCIDR(cidr)
	return func(ip net.IP) bool {
		if ipnet != nil {
			return ipnet.Contains(ip)
		}
		return false
	}
}
