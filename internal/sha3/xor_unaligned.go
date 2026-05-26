// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (amd64 || 386 || ppc64le) && !appengine
// +build amd64 386 ppc64le
// +build !appengine

package sha3

// A storageBuf is an aligned array of maxRate bytes.
type storageBuf [maxRate / 8]uint64

func (b *storageBuf) asBytes() *[maxRate]byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// xorInuses unaligned reads and writes to update d.a to contain d.a
// XOR buf.
func xorIn(d *State, buf []byte) { _ = "STUB: not implemented"; return }

//nolint:gosec

func copyOut(d *State, buf []byte) { _ = "STUB: not implemented"; return }

//nolint:gosec
