// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

// KeccakF1600 applies the Keccak permutation to a 1600b-wide
// state represented as a slice of 25 uint64s.
// If turbo is true, applies the 12-round variant instead of the
// regular 24-round variant.
// nolint:funlen
func KeccakF1600(a *[25]uint64, turbo bool) {
	_ = "STUB: not implemented"
	// Implementation translated from Keccak-inplace.c
	// in the keccak reference code.
	return
}

// Combines the 5 steps in each round into 2 steps.
// Unrolls 4 rounds per loop and spreads some steps across rounds.

// Round 1

// Round 2

// Round 3

// Round 4
