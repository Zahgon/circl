// Package xof provides an interface for eXtendable-Output Functions.
//
// # Available Functions
//
// SHAKE functions are defined in FIPS-202, see https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.202.pdf.
// BLAKE2Xb and BLAKE2Xs are defined in https://www.blake2.net/blake2x.pdf.
package xof

import (
	"io"

	"github.com/cloudflare/circl/internal/sha3"
	"github.com/cloudflare/circl/xof/k12"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

// XOF defines the interface to hash functions that support arbitrary-length output.
type XOF interface {
	// Write absorbs more data into the XOF's state. It panics if called
	// after Read.
	io.Writer

	// Read reads more output from the XOF. It returns io.EOF if the limit
	// has been reached.
	io.Reader

	// Clone returns a copy of the XOF in its current state.
	Clone() XOF

	// Reset restores the XOF to its initial state and discards all data appended by Write.
	Reset()
}

type ID uint

const (
	SHAKE128 ID = iota + 1
	SHAKE256
	BLAKE2XB
	BLAKE2XS
	K12D10
)

func (x ID) New() XOF { _ = "STUB: not implemented"; return *new(XOF) }

type shakeBody struct{ sha3.ShakeHash }

func (s shakeBody) Clone() XOF { _ = "STUB: not implemented"; return *new(XOF) }

type blake2xb struct{ blake2b.XOF }

func (s blake2xb) Clone() XOF { _ = "STUB: not implemented"; return *new(XOF) }

type blake2xs struct{ blake2s.XOF }

func (s blake2xs) Clone() XOF { _ = "STUB: not implemented"; return *new(XOF) }

type k12d10 struct{ *k12.State }

func (s k12d10) Clone() XOF { _ = "STUB: not implemented"; return *new(XOF) }
