// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (!amd64 || appengine) && (!386 || appengine) && (!ppc64le || appengine)
// +build !amd64 appengine
// +build !386 appengine
// +build !ppc64le appengine

package sha3

// xorIn xors the bytes in buf into the state; it
// makes no non-portable assumptions about memory layout
// or alignment.
func xorIn(d *State, buf []byte) { _ = "STUB: not implemented"; return }

// copyOut copies ulint64s to a byte buffer.
func copyOut(d *State, b []byte) { _ = "STUB: not implemented"; return }
