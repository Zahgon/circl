// k12 implements the KangarooTwelve XOF.
//
// KangarooTwelve is being standardised at the CFRG working group
// of the IRTF. This package implements draft 10.
//
// https://datatracker.ietf.org/doc/draft-irtf-cfrg-kangarootwelve/10/
package k12

import (
	"github.com/cloudflare/circl/internal/sha3"
)

const chunkSize = 8192 // aka B

// KangarooTwelve splits the message into chunks of 8192 bytes each.
// The first chunk is absorbed directly in a TurboSHAKE128 instance, which
// we call the stalk. The subsequent chunks aren't absorbed directly, but
// instead their hash is absorbed: they're like leafs on a stalk.
// If we have a fast TurboSHAKE128 available, we buffer chunks until we have
// enough to do the parallel TurboSHAKE128. If not, we absorb directly into
// a separate TurboSHAKE128 state.

type State struct {
	initialTodo int // Bytes left to absorb for the first chunk.

	stalk sha3.State

	context []byte // context string "C" provided by the user

	// buffer of incoming data so we can do parallel TurboSHAKE128:
	// nil when we haven't absorbed the first chunk yet;
	// empty if we have, but we do not have a fast parallel TurboSHAKE128;
	// and chunkSize*lanes in length if we have.
	buf []byte

	offset int // offset in buf or bytes written to leaf

	// Number of chunk hashes ("CV_i") absorbed into the stalk.
	chunk uint

	// TurboSHAKE128 instance to compute the leaf in case we don't have
	// a fast parallel TurboSHAKE128, viz when lanes == 1.
	leaf *sha3.State

	lanes uint8 // number of TurboSHAKE128s to compute in parallel
}

// NewDraft10 creates a new instance of Kangaroo12 draft version -10.
func NewDraft10(c []byte) State { _ = "STUB: not implemented"; return *new(State) }

func newDraft10(c []byte, lanes byte) State { _ = "STUB: not implemented"; return *new(State) }

func (s *State) Reset() { _ = "STUB: not implemented"; return }

func (s *State) Clone() State { _ = "STUB: not implemented"; return *new(State) }

func Draft10Sum(hash []byte, msg []byte, c []byte) {
	_ = "STUB: not implemented"
	// TODO Tweak number of lanes depending on the length of the message
	return
}

func (s *State) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"

	// The first chunk is written directly to the stalk.
	return 0, nil
}

// If this is the first bit of data written after the initial chunk,
// we're out of the fast-path and allocate some buffers.

// We create the buffer to signal we're past the first chunk,
// but do not use it.

// If we're just using one lane, we don't need to cache in a buffer
// for parallel hashing. Instead, we feed directly to TurboSHAKE.

// Write to current leaf.

// Did we fill the chunk?

// If we can't fill all our lanes or the buffer isn't empty, we write the
// data to the buffer.

// Absorb the buffer if we filled it

// Note that at this point we may assume that s.offset = 0 if len(p) != 0

// Absorb a bunch of chunks at the same time.

// Put the remainder in the buffer.

// Absorb a multiple of a multiple of lanes * chunkSize.
// Returns the remainder.
func (s *State) writeX(p []byte) []byte { _ = "STUB: not implemented"; return nil }

func (s *State) writeX4(p []byte) []byte { _ = "STUB: not implemented"; return nil }

func (s *State) writeX2(p []byte) []byte {
	_ = "STUB: not implemented"
	// TODO On M2 Pro, 1/3 of the time is spent on this function
	// and LittleEndian.Uint64 excluding the actual permutation.
	// Rewriting in assembler might be worthwhile.
	return nil
}

func (s *State) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Write context string C
		nil
}

// Write length_encode( |C| )

// Find first non-zero digit in big endian encoding of context length

// number of bytes to represent |C|

// We need to write the chunk number if we're past the first chunk.

// Write last remaining chunk(s)

// Write length_encode( chunk )

// Find first non-zero digit in big endian encoding of number of chunks

// number of bytes to represent number of chunks.
