package ff

import (
	"io"
)

// FpSize is the length in bytes of an Fp element.
const FpSize = 48

// fpMont represents an element in the Montgomery domain (little-endian).
type fpMont = [FpSize / 8]uint64

// fpRaw represents an element in the integers domain (little-endian).
type fpRaw = [FpSize / 8]uint64

// Fp represents prime field elements as positive integers less than FpOrder.
type Fp struct{ i fpMont }

func (z Fp) String() string            { _ = "STUB: not implemented"; return "" }
func (z *Fp) SetUint64(n uint64)       { _ = "STUB: not implemented"; return }
func (z *Fp) SetOne()                  { _ = "STUB: not implemented"; return }
func (z *Fp) Random(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// IsNegative returns 0 if the least absolute residue for z is in [0,(p-1)/2],
// and 1 otherwise. Equivalently, this function returns 1 if z is
// lexicographically larger than -z.
func (z Fp) IsNegative() int { _ = "STUB: not implemented"; return 0 }

// IsZero returns 1 if z == 0 and 0 otherwise.
func (z Fp) IsZero() int { _ = "STUB: not implemented"; return 0 }

// IsEqual returns 1 if z == x and 0 otherwise.
func (z Fp) IsEqual(x *Fp) int     { _ = "STUB: not implemented"; return 0 }
func (z *Fp) Neg()                 { _ = "STUB: not implemented"; return }
func (z *Fp) Add(x, y *Fp)         { _ = "STUB: not implemented"; return }
func (z *Fp) Sub(x, y *Fp)         { _ = "STUB: not implemented"; return }
func (z *Fp) Mul(x, y *Fp)         { _ = "STUB: not implemented"; return }
func (z *Fp) Sqr(x *Fp)            { _ = "STUB: not implemented"; return }
func (z *Fp) toMont(in *fpRaw)     { _ = "STUB: not implemented"; return }
func (z Fp) fromMont() (out fpRaw) { _ = "STUB: not implemented"; return *new(fpRaw) }
func (z Fp) Sgn0() int             { _ = "STUB: not implemented"; return 0 }

// Sqrt returns 1 and sets z=sqrt(x) only if x is a quadratic-residue; otherwise, returns 0 and z is unmodified.
func (z *Fp) Sqrt(x *Fp) int { _ = "STUB: not implemented"; return 0 }

// CMov sets z=x if b == 0 and z=y if b == 1. Its behavior is undefined if b takes any other value.
func (z *Fp) CMov(x, y *Fp, b int) { _ = "STUB: not implemented"; return }

// FpOrder is the order of the base field for towering returned as a big-endian slice.
//
//	FpOrder = 0x1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab.
func FpOrder() []byte { _ = "STUB: not implemented"; return nil }

// ExpVarTime calculates z=x^n, where n is the exponent in big-endian order.
func (z *Fp) ExpVarTime(x *Fp, n []byte) { _ = "STUB: not implemented"; return }

// SetBytes assigns to z the number modulo FpOrder stored in the slice
// (in big-endian order).
func (z *Fp) SetBytes(data []byte) { _ = "STUB: not implemented"; return }

// MarshalBinary returns a slice of FpSize bytes that contains the minimal
// residue of z such that 0 <= z < FpOrder (in big-endian order).
func (z *Fp) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary reconstructs a Fp from a slice that must have at least
// FpSize bytes and contain a number (in big-endian order) from 0
// to FpOrder-1.
func (z *Fp) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

// SetString reconstructs a Fp from a numeric string from 0 to FpOrder-1.
func (z *Fp) SetString(s string) error { _ = "STUB: not implemented"; return nil }

func fiatFpMontCmovznzU64(z *uint64, b, x, y uint64) { _ = "STUB: not implemented"; return }

func (z *Fp) Inv(x *Fp) {
	_ = "STUB: not implemented"
	// Addition chain found using mmcloughlin/addchain: v0.3.0
	// McLoughlin, Michael Ben. (2021). https://doi.org/10.5281/zenodo.4758226
	return
}

var (
	// fpOrder is the order of the Fp field (big-endian).
	fpOrder = [FpSize]byte{
		0x1a, 0x01, 0x11, 0xea, 0x39, 0x7f, 0xe6, 0x9a,
		0x4b, 0x1b, 0xa7, 0xb6, 0x43, 0x4b, 0xac, 0xd7,
		0x64, 0x77, 0x4b, 0x84, 0xf3, 0x85, 0x12, 0xbf,
		0x67, 0x30, 0xd2, 0xa0, 0xf6, 0xb0, 0xf6, 0x24,
		0x1e, 0xab, 0xff, 0xfe, 0xb1, 0x53, 0xff, 0xff,
		0xb9, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xaa, 0xab,
	}
	// fpOrderPlus1Div2 is the half of (fpOrder plus one) used for lexicographically order (big-endian).
	fpOrderPlus1Div2 = [FpSize]byte{
		0x0d, 0x00, 0x88, 0xf5, 0x1c, 0xbf, 0xf3, 0x4d,
		0x25, 0x8d, 0xd3, 0xdb, 0x21, 0xa5, 0xd6, 0x6b,
		0xb2, 0x3b, 0xa5, 0xc2, 0x79, 0xc2, 0x89, 0x5f,
		0xb3, 0x98, 0x69, 0x50, 0x7b, 0x58, 0x7b, 0x12,
		0x0f, 0x55, 0xff, 0xff, 0x58, 0xa9, 0xff, 0xff,
		0xdc, 0xff, 0x7f, 0xff, 0xff, 0xff, 0xd5, 0x56,
	}
	// fpOrderPlus1Div4 is (fpOrder plus one) divided by four used for square-roots (big-endian).
	fpOrderPlus1Div4 = [FpSize]byte{
		0x06, 0x80, 0x44, 0x7a, 0x8e, 0x5f, 0xf9, 0xa6,
		0x92, 0xc6, 0xe9, 0xed, 0x90, 0xd2, 0xeb, 0x35,
		0xd9, 0x1d, 0xd2, 0xe1, 0x3c, 0xe1, 0x44, 0xaf,
		0xd9, 0xcc, 0x34, 0xa8, 0x3d, 0xac, 0x3d, 0x89,
		0x07, 0xaa, 0xff, 0xff, 0xac, 0x54, 0xff, 0xff,
		0xee, 0x7f, 0xbf, 0xff, 0xff, 0xff, 0xea, 0xab,
	}
	// fpRSquare is R^2 mod fpOrder, where R=2^384 (little-endian).
	fpRSquare = fpMont{
		0xf4df1f341c341746, 0x0a76e6a609d104f1,
		0x8de5476c4c95b6d5, 0x67eb88a9939d83c0,
		0x9a793e85b519952d, 0x11988fe592cae3aa,
	}
)
