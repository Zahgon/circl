// Package ff provides finite fields of characteristic P381.
package ff

import (
	"errors"
	"io"
)

var (
	errInputLength = errors.New("incorrect input length")
	errInputRange  = errors.New("value out of range [0,order)")
	errInputString = errors.New("invalid string")
)

func errFirst(e ...error) (err error) { _ = "STUB: not implemented"; return nil }

func setString(in string, order []byte) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setBytesBounded(in []byte, order []byte) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setBytesUnbounded(in []byte, order []byte) []uint64 { _ = "STUB: not implemented"; return nil }

// isLessThan returns 1 if 0 <= x < y, otherwise 0. Assumes that slices have the same length.
func isLessThan(x, y []byte) int { _ = "STUB: not implemented"; return 0 }

func randomInt(out []uint64, rnd io.Reader, order []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ctUint64Eq returns 1 if the two slices have equal contents and 0 otherwise.
func ctUint64Eq(x, y []uint64) (b int) { _ = "STUB: not implemented"; return 0 }

func cselectU64(z *uint64, b, x, y uint64) { _ = "STUB: not implemented"; return }
