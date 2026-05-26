// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

// This file provides functions for creating instances of the SHA-3
// and SHAKE hash functions, as well as utility functions for hashing
// bytes.

// New224 creates a new SHA3-224 hash.
// Its generic security strength is 224 bits against preimage attacks,
// and 112 bits against collision attacks.
func New224() State { _ = "STUB: not implemented"; return *new(State) }

// New256 creates a new SHA3-256 hash.
// Its generic security strength is 256 bits against preimage attacks,
// and 128 bits against collision attacks.
func New256() State { _ = "STUB: not implemented"; return *new(State) }

// New384 creates a new SHA3-384 hash.
// Its generic security strength is 384 bits against preimage attacks,
// and 192 bits against collision attacks.
func New384() State { _ = "STUB: not implemented"; return *new(State) }

// New512 creates a new SHA3-512 hash.
// Its generic security strength is 512 bits against preimage attacks,
// and 256 bits against collision attacks.
func New512() State { _ = "STUB: not implemented"; return *new(State) }

// Sum224 returns the SHA3-224 digest of the data.
func Sum224(data []byte) (digest [28]byte) { _ = "STUB: not implemented"; return nil }

// Sum256 returns the SHA3-256 digest of the data.
func Sum256(data []byte) (digest [32]byte) { _ = "STUB: not implemented"; return nil }

// Sum384 returns the SHA3-384 digest of the data.
func Sum384(data []byte) (digest [48]byte) { _ = "STUB: not implemented"; return nil }

// Sum512 returns the SHA3-512 digest of the data.
func Sum512(data []byte) (digest [64]byte) { _ = "STUB: not implemented"; return nil }
