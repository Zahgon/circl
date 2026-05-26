// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

// spongeDirection indicates the direction bytes are flowing through the sponge.
type spongeDirection int

const (
	// spongeAbsorbing indicates that the sponge is absorbing input.
	spongeAbsorbing spongeDirection = iota
	// spongeSqueezing indicates that the sponge is being squeezed.
	spongeSqueezing
)

const (
	// maxRate is the maximum size of the internal buffer. SHAKE-256
	// currently needs the largest buffer.
	maxRate = 168
)

func (d *State) buf() []byte { _ = "STUB: not implemented"; return nil }

type State struct {
	// Generic sponge components.
	a    [25]uint64 // main state of the hash
	rate int        // the number of bytes of state to use

	bufo int // offset of buffer in storage
	bufe int // end of buffer in storage

	// dsbyte contains the "domain separation" bits and the first bit of
	// the padding. Sections 6.1 and 6.2 of [1] separate the outputs of the
	// SHA-3 and SHAKE functions by appending bitstrings to the message.
	// Using a little-endian bit-ordering convention, these are "01" for SHA-3
	// and "1111" for SHAKE, or 00000010b and 00001111b, respectively. Then the
	// padding rule from section 5.1 is applied to pad the message to a multiple
	// of the rate, which involves adding a "1" bit, zero or more "0" bits, and
	// a final "1" bit. We merge the first "1" bit from the padding into dsbyte,
	// giving 00000110b (0x06) and 00011111b (0x1f).
	// [1] http://csrc.nist.gov/publications/drafts/fips-202/fips_202_draft.pdf
	//     "Draft FIPS 202: SHA-3 Standard: Permutation-Based Hash and
	//      Extendable-Output Functions (May 2014)"
	dsbyte byte

	storage storageBuf

	// Specific to SHA-3 and SHAKE.
	outputLen int             // the default output size in bytes
	state     spongeDirection // whether the sponge is absorbing or squeezing
	turbo     bool            // Whether we're using 12 rounds instead of 24
}

// BlockSize returns the rate of sponge underlying this hash function.
func (d *State) BlockSize() int {
	_ = "STUB: not implemented"

	// Size returns the output size of the hash function in bytes.
	return 0
}

func (d *State) Size() int {
	_ = "STUB: not implemented"

	// Reset clears the internal state by zeroing the sponge state and
	// the byte buffer, and setting Sponge.state to absorbing.
	return 0
}

func (d *State) Reset() {
	_ = "STUB: not implemented"
	// Zero the permutation's state.
	return
}

func (d *State) clone() *State { _ = "STUB: not implemented"; return nil }

// permute applies the KeccakF-1600 permutation. It handles
// any input-output buffering.
func (d *State) permute() { _ = "STUB: not implemented"; return }

// If we're absorbing, we need to xor the input into the state
// before applying the permutation.

// If we're squeezing, we need to apply the permutation before
// copying more output.

// pads appends the domain separation bits in dsbyte, applies
// the multi-bitrate 10..1 padding rule, and permutes the state.
func (d *State) padAndPermute(dsbyte byte) {
	_ = "STUB: not implemented"
	// Pad with this instance's domain-separator bits. We know that there's
	// at least one byte of space in d.buf() because, if it were full,
	// permute would have been called to empty it. dsbyte also contains the
	// first one bit for the padding. See the comment in the state struct.
	return
}

// This adds the final one bit for the padding. Because of the way that
// bits are numbered from the LSB upwards, the final bit is the MSB of
// the last byte.

// Apply the permutation

// Write absorbs more data into the hash's state. It produces an error
// if more data is written to the ShakeHash after writing
func (d *State) Write(p []byte) (written int, err error) { _ = "STUB: not implemented"; return 0, nil }

// The fast path; absorb a full "rate" bytes of input and apply the permutation.

// The slow path; buffer the input until we can fill the sponge, and then xor it in.

// If the sponge is full, apply the permutation.

// Read squeezes an arbitrary number of bytes from the sponge.
func (d *State) Read(out []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// If we're still absorbing, pad and apply the permutation.
	return 0, nil
}

// Now, do the squeezing.

// Apply the permutation if we've squeezed the sponge dry.

// Sum applies padding to the hash state and then squeezes out the desired
// number of output bytes.
func (d *State) Sum(in []byte) []byte {
	_ = "STUB: not implemented"
	// Make a copy of the original hash so that caller can keep writing
	// and summing.
	return nil
}

func (d *State) IsAbsorbing() bool { _ = "STUB: not implemented"; return false }
