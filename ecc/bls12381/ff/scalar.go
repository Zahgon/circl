package ff

import (
	"io"
)

// ScalarSize is the length in bytes of a Scalar.
const ScalarSize = 32

// scMont represents an element in the Montgomery domain (little-endian).
type scMont = [ScalarSize / 8]uint64

// scRaw represents a scalar in the integers domain (little-endian).
type scRaw = [ScalarSize / 8]uint64

// Scalar represents positive integers less than ScalarOrder.
type Scalar struct{ i scMont }

func (z Scalar) String() string            { _ = "STUB: not implemented"; return "" }
func (z *Scalar) Set(x *Scalar)            { _ = "STUB: not implemented"; return }
func (z *Scalar) SetUint64(n uint64)       { _ = "STUB: not implemented"; return }
func (z *Scalar) SetOne()                  { _ = "STUB: not implemented"; return }
func (z *Scalar) Random(r io.Reader) error { _ = "STUB: not implemented"; return nil }
func (z Scalar) IsZero() int               { _ = "STUB: not implemented"; return 0 }
func (z Scalar) IsEqual(x *Scalar) int     { _ = "STUB: not implemented"; return 0 }
func (z *Scalar) Neg()                     { _ = "STUB: not implemented"; return }
func (z *Scalar) Add(x, y *Scalar)         { _ = "STUB: not implemented"; return }
func (z *Scalar) Sub(x, y *Scalar)         { _ = "STUB: not implemented"; return }
func (z *Scalar) Mul(x, y *Scalar)         { _ = "STUB: not implemented"; return }
func (z *Scalar) Sqr(x *Scalar)            { _ = "STUB: not implemented"; return }
func (z *Scalar) Inv(x *Scalar)            { _ = "STUB: not implemented"; return }
func (z *Scalar) toMont(in *scRaw)         { _ = "STUB: not implemented"; return }
func (z Scalar) fromMont() (out scRaw)     { _ = "STUB: not implemented"; return *new(scRaw) }

// ScalarOrder is the order of the scalar field of the pairing groups, order is
// returned as a big-endian slice.
//
//	ScalarOrder = 0x73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001
func ScalarOrder() []byte { _ = "STUB: not implemented"; return nil }

// exp calculates z=x^n, where n is in big-endian order.
func (z *Scalar) expVarTime(x *Scalar, n []byte) { _ = "STUB: not implemented"; return }

// SetBytes assigns to z the number modulo ScalarOrder stored in the slice
// (in big-endian order).
func (z *Scalar) SetBytes(data []byte) { _ = "STUB: not implemented"; return }

// MarshalBinary returns a slice of ScalarSize bytes that contains the minimal
// residue of z such that 0 <= z < ScalarOrder (in big-endian order).
func (z *Scalar) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary reconstructs a Scalar from a slice that must have at least
// ScalarSize bytes and contain a number (in big-endian order) from 0
// to ScalarOrder-1.
func (z *Scalar) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// SetString reconstructs a Fp from a numeric string from 0 to ScalarOrder-1.
func (z *Scalar) SetString(s string) error { _ = "STUB: not implemented"; return nil }

func fiatScMontCmovznzU64(z *uint64, b, x, y uint64) { _ = "STUB: not implemented"; return }

var (
	// scOrder is the order of the Scalar field (big-endian).
	scOrder = [ScalarSize]byte{
		0x73, 0xed, 0xa7, 0x53, 0x29, 0x9d, 0x7d, 0x48,
		0x33, 0x39, 0xd8, 0x08, 0x09, 0xa1, 0xd8, 0x05,
		0x53, 0xbd, 0xa4, 0x02, 0xff, 0xfe, 0x5b, 0xfe,
		0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x01,
	}
	// scOrderMinus2 is the scOrder minus two used for inversion (big-endian).
	scOrderMinus2 = [ScalarSize]byte{
		0x73, 0xed, 0xa7, 0x53, 0x29, 0x9d, 0x7d, 0x48,
		0x33, 0x39, 0xd8, 0x08, 0x09, 0xa1, 0xd8, 0x05,
		0x53, 0xbd, 0xa4, 0x02, 0xff, 0xfe, 0x5b, 0xfe,
		0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff,
	}
	// scRSquare is R^2 mod scOrder, where R=2^256 (little-endian).
	scRSquare = scMont{
		0xc999e990f3f29c6d, 0x2b6cedcb87925c23,
		0x05d314967254398f, 0x0748d9d99f59ff11,
	}
)
