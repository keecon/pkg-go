// Copyright 2022 Keecon Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cipher implements AES encryption with GCM, KDF Hash
package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"

	"golang.org/x/crypto/hkdf"
)

// AES implements encrypt/decrypt AES algorithm (GCM)
type AES struct {
	secret    string
	alg       string
	algKeyLen int
	nonceLen  int
	hkdfHash  func() hash.Hash
	hkdfInfo  []byte
}

// Option defines configure AES settings
type Option func(*AES)

// NewAES creates AES
func NewAES(secret string, opts ...Option) *AES {
	ret := &AES{
		secret:    secret,
		alg:       "AES256",   // recommends
		algKeyLen: 32,         // recommends
		nonceLen:  12,         // strongly recommends
		hkdfHash:  sha256.New, // recommends
	}

	for _, o := range opts {
		o(ret)
	}
	return ret
}

// WithAES256 configures AES256 algorithm
func WithAES256() Option {
	return func(c *AES) {
		c.alg = "AES256"
		c.algKeyLen = 32
	}
}

// WithAES192 configures AES192 algorithm
func WithAES192() Option {
	return func(c *AES) {
		c.alg = "AES192"
		c.algKeyLen = 24
	}
}

// WithAES128 configures AES128 algorithm
func WithAES128() Option {
	return func(c *AES) {
		c.alg = "AES128"
		c.algKeyLen = 16
	}
}

// WithNonceLength configures nonce length
func WithNonceLength(n int) Option {
	return func(c *AES) {
		c.nonceLen = n
	}
}

// WithHKDFHash configures Key Derivation Function (HKDF)
func WithHKDFHash(fn func() hash.Hash) Option {
	return func(c *AES) {
		c.hkdfHash = fn
	}
}

// WithHKDFInfo configures Key Derivation Function (HKDF) info
func WithHKDFInfo(info []byte) Option {
	return func(c *AES) {
		c.hkdfInfo = info
	}
}

// NewInt64Salt returns bytes for using salt
func (c *AES) NewInt64Salt(data int64) []byte {
	salt := make([]byte, 8)
	binary.LittleEndian.PutUint64(salt, uint64(data))
	return salt
}

// NewInt32Salt returns bytes for using salt
func (c *AES) NewInt32Salt(data int32) []byte {
	salt := make([]byte, 8)
	binary.LittleEndian.PutUint32(salt, uint32(data))
	return salt
}

// Encrypt implements encrypt and authenticates plaintext
func (c *AES) Encrypt(plaintext, salt []byte) ([]byte, error) {
	key, nonce, err := c.newKeyNonce(salt)
	if err != nil {
		return nil, err
	}

	aead, err := c.newCipherAEAD(key)
	if err != nil {
		return nil, err
	}

	return aead.Seal(nil, nonce, plaintext, nil), nil
}

// Decrypt implements decrypt and authenticates ciphertext
func (c *AES) Decrypt(ciphertext, salt []byte) ([]byte, error) {
	key, nonce, err := c.newKeyNonce(salt)
	if err != nil {
		return nil, err
	}

	aead, err := c.newCipherAEAD(key)
	if err != nil {
		return nil, err
	}

	return aead.Open(nil, nonce, ciphertext, nil)
}

func (c *AES) newKeyNonce(salt []byte) (key []byte, nonce []byte, err error) {
	kdf := hkdf.New(c.hkdfHash, []byte(c.secret), salt, c.hkdfInfo)
	key = make([]byte, c.algKeyLen)
	if _, err := kdf.Read(key); err != nil {
		return nil, nil, fmt.Errorf("hkdf expand key: %w", err)
	}

	nonce = make([]byte, c.nonceLen)
	if _, err := kdf.Read(nonce); err != nil {
		return nil, nil, fmt.Errorf("hkdf expand nonce: %w", err)
	}
	return key, nonce, nil
}

func (c *AES) newCipherAEAD(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("new aes cipher: %w", err)
	}

	aead, err := cipher.NewGCMWithNonceSize(block, c.nonceLen)
	if err != nil {
		return nil, fmt.Errorf("new aead gcm: %w", err)
	}
	return aead, nil
}
