// Copyright 2022 KEECON CO.,LTD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAES256(t *testing.T) {
	// dataset
	dataset := []struct {
		name      string
		secretLen int
	}{
		{
			name:      "SecretLen16",
			secretLen: 16,
		},
		{
			name:      "SecretLen32",
			secretLen: 32,
		},
	}

	// table driven tests
	for _, v := range dataset {
		t.Run(v.name, func(t *testing.T) {
			// given
			secret := newRandHex(t, v.secretLen/2)
			plaintext := newRandBytes(t, 128)
			c := NewAES(secret, WithAES256())

			// when
			salt := c.NewInt64Salt(time.Now().Unix())
			ciphertext, err := c.Encrypt(plaintext, salt)
			assert.NoError(t, err)

			rettext, err := c.Decrypt(ciphertext, salt)
			assert.NoError(t, err)

			// then
			assert.Equal(t, plaintext, rettext)
		})
	}
}

func TestAES256Options(t *testing.T) {
	// dataset
	dataset := []struct {
		name      string
		secretLen int
		nonceLen  int
		hkdfHash  func() hash.Hash
		hkdfInfo  []byte
	}{
		{
			name:      "NonceLen16",
			secretLen: 32,
			nonceLen:  16,
			hkdfHash:  sha256.New,
		},
		{
			name:      "HashSHA512",
			secretLen: 32,
			nonceLen:  12,
			hkdfHash:  sha512.New,
		},
		{
			name:      "HashSHA512",
			secretLen: 32,
			nonceLen:  12,
			hkdfHash:  sha256.New,
			hkdfInfo:  []byte("test-info"),
		},
	}

	// table driven tests
	for _, v := range dataset {
		t.Run(v.name, func(t *testing.T) {
			// given
			secret := newRandHex(t, v.secretLen/2)
			plaintext := newRandBytes(t, 128)
			c := NewAES(secret,
				WithAES256(),
				WithNonceLength(v.nonceLen),
				WithHKDFHash(v.hkdfHash),
				WithHKDFInfo(v.hkdfInfo),
			)

			// when
			salt := c.NewInt64Salt(time.Now().Unix())
			ciphertext, err := c.Encrypt(plaintext, salt)
			assert.NoError(t, err)

			rettext, err := c.Decrypt(ciphertext, salt)
			assert.NoError(t, err)

			// then
			assert.Equal(t, plaintext, rettext)
		})
	}
}

func TestAES192(t *testing.T) {
	// dataset
	dataset := []struct {
		name      string
		secretLen int
	}{
		{
			name:      "SecretLen16",
			secretLen: 16,
		},
		{
			name:      "SecretLen32",
			secretLen: 32,
		},
	}

	// table driven tests
	for _, v := range dataset {
		t.Run(v.name, func(t *testing.T) {
			// given
			secret := newRandHex(t, v.secretLen/2)
			plaintext := newRandBytes(t, 128)
			c := NewAES(secret, WithAES192())

			// when
			salt := c.NewInt64Salt(time.Now().Unix())
			ciphertext, err := c.Encrypt(plaintext, salt)
			assert.NoError(t, err)

			rettext, err := c.Decrypt(ciphertext, salt)
			assert.NoError(t, err)

			// then
			assert.Equal(t, plaintext, rettext)
		})
	}
}

func TestAES192Options(t *testing.T) {
	// dataset
	dataset := []struct {
		name      string
		secretLen int
		nonceLen  int
		hkdfHash  func() hash.Hash
		hkdfInfo  []byte
	}{
		{
			name:      "NonceLen16",
			secretLen: 32,
			nonceLen:  16,
			hkdfHash:  sha256.New,
		},
		{
			name:      "HashSHA512",
			secretLen: 32,
			nonceLen:  12,
			hkdfHash:  sha512.New,
		},
		{
			name:      "HashSHA512",
			secretLen: 32,
			nonceLen:  12,
			hkdfHash:  sha256.New,
			hkdfInfo:  []byte("test-info"),
		},
	}

	// table driven tests
	for _, v := range dataset {
		t.Run(v.name, func(t *testing.T) {
			// given
			secret := newRandHex(t, v.secretLen/2)
			plaintext := newRandBytes(t, 128)
			c := NewAES(secret,
				WithAES192(),
				WithNonceLength(v.nonceLen),
				WithHKDFHash(v.hkdfHash),
				WithHKDFInfo(v.hkdfInfo),
			)

			// when
			salt := c.NewInt64Salt(time.Now().Unix())
			ciphertext, err := c.Encrypt(plaintext, salt)
			assert.NoError(t, err)

			rettext, err := c.Decrypt(ciphertext, salt)
			assert.NoError(t, err)

			// then
			assert.Equal(t, plaintext, rettext)
		})
	}
}

func TestAES128(t *testing.T) {
	// dataset
	dataset := []struct {
		name      string
		secretLen int
	}{
		{
			name:      "SecretLen16",
			secretLen: 16,
		},
		{
			name:      "SecretLen32",
			secretLen: 32,
		},
	}

	// table driven tests
	for _, v := range dataset {
		t.Run(v.name, func(t *testing.T) {
			// given
			secret := newRandHex(t, v.secretLen/2)
			plaintext := newRandBytes(t, 256)
			c := NewAES(secret, WithAES128())

			// when
			salt := c.NewInt64Salt(time.Now().Unix())
			ciphertext, err := c.Encrypt(plaintext, salt)
			assert.NoError(t, err)

			rettext, err := c.Decrypt(ciphertext, salt)
			assert.NoError(t, err)

			// then
			assert.Equal(t, plaintext, rettext)
		})
	}
}

func TestAES128Options(t *testing.T) {
	// dataset
	dataset := []struct {
		name      string
		secretLen int
		nonceLen  int
		hkdfHash  func() hash.Hash
		hkdfInfo  []byte
	}{
		{
			name:      "NonceLen16",
			secretLen: 32,
			nonceLen:  16,
			hkdfHash:  sha256.New,
		},
		{
			name:      "HashSHA512",
			secretLen: 32,
			nonceLen:  12,
			hkdfHash:  sha512.New,
		},
		{
			name:      "HashSHA512",
			secretLen: 32,
			nonceLen:  12,
			hkdfHash:  sha256.New,
			hkdfInfo:  []byte("test-info"),
		},
	}

	// table driven tests
	for _, v := range dataset {
		t.Run(v.name, func(t *testing.T) {
			// given
			secret := newRandHex(t, v.secretLen/2)
			plaintext := newRandBytes(t, 256)
			c := NewAES(secret,
				WithAES128(),
				WithNonceLength(v.nonceLen),
				WithHKDFHash(v.hkdfHash),
				WithHKDFInfo(v.hkdfInfo),
			)

			// when
			salt := c.NewInt64Salt(time.Now().Unix())
			ciphertext, err := c.Encrypt(plaintext, salt)
			assert.NoError(t, err)

			rettext, err := c.Decrypt(ciphertext, salt)
			assert.NoError(t, err)

			// then
			assert.Equal(t, plaintext, rettext)
		})
	}
}

func newRandBytes(t *testing.T, length int) []byte {
	bytes := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, bytes[:])
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	return bytes
}

func newRandHex(t *testing.T, length int) string {
	return hex.EncodeToString(newRandBytes(t, length))
}
