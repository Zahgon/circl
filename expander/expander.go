// Package expander generates arbitrary bytes from an XOF or Hash function.
package expander

import (
	"crypto"
	"errors"
	"io"

	"github.com/cloudflare/circl/xof"
)

type Expander interface {
	// Expand generates a pseudo-random byte string of a determined length by
	// expanding an input string.
	Expand(in []byte, length uint) (pseudo []byte)
}

type expanderMD struct {
	h   crypto.Hash
	dst []byte
}

// NewExpanderMD returns a hash function based on a Merkle-Damgård hash function.
func NewExpanderMD(h crypto.Hash, dst []byte) *expanderMD { _ = "STUB: not implemented"; return nil }

func (e *expanderMD) calcDSTPrime() []byte { _ = "STUB: not implemented"; return nil }

func (e *expanderMD) Expand(in []byte, n uint) []byte { _ = "STUB: not implemented"; return nil }

// expanderXOF is based on an extendable output function.
type expanderXOF struct {
	id        xof.ID
	kSecLevel uint
	dst       []byte
}

// NewExpanderXOF returns an Expander based on an extendable output function.
// The kSecLevel parameter is the target security level in bits, and dst is
// a domain separation string.
func NewExpanderXOF(id xof.ID, kSecLevel uint, dst []byte) *expanderXOF {
	_ = "STUB: not implemented"
	return nil
}

// Expand panics if output's length is longer than 2^16 bytes.
func (e *expanderXOF) Expand(in []byte, n uint) []byte { _ = "STUB: not implemented"; return nil }

func (e *expanderXOF) calcDSTPrime() []byte { _ = "STUB: not implemented"; return nil }

func mustWrite(w io.Writer, b []byte) { _ = "STUB: not implemented"; return }

func mustReadFull(r io.Reader, b []byte) { _ = "STUB: not implemented"; return }

const maxDSTLength = 255

var (
	longDSTPrefix = [17]byte{'H', '2', 'C', '-', 'O', 'V', 'E', 'R', 'S', 'I', 'Z', 'E', '-', 'D', 'S', 'T', '-'}

	errorLongOutput = errors.New("requested too many bytes")
)
