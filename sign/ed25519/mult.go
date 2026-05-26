package ed25519

import (
	fp "github.com/cloudflare/circl/math/fp25519"
)

var paramD = fp.Elt{
	0xa3, 0x78, 0x59, 0x13, 0xca, 0x4d, 0xeb, 0x75,
	0xab, 0xd8, 0x41, 0x41, 0x4d, 0x0a, 0x70, 0x00,
	0x98, 0xe8, 0x79, 0x77, 0x79, 0x40, 0xc7, 0x8c,
	0x73, 0xfe, 0x6f, 0x2b, 0xee, 0x6c, 0x03, 0x52,
}

// mLSBRecoding parameters.
const (
	fxT        = 257
	fxV        = 2
	fxW        = 3
	fx2w1      = 1 << (uint(fxW) - 1)
	numWords64 = (paramB * 8 / 64)
)

// mLSBRecoding is the odd-only modified LSB-set.
//
// Reference:
//
//	"Efficient and secure algorithms for GLV-based scalar multiplication and
//	 their implementation on GLV–GLS curves" by (Faz-Hernandez et al.)
//	 http://doi.org/10.1007/s13389-014-0085-7.
func mLSBRecoding(L []int8, k []byte) { _ = "STUB: not implemented"; return }

// right-shift by d

// absolute returns always a positive value.
func absolute(x int32) int32 { _ = "STUB: not implemented"; return 0 }

// condAddOrderN updates x = x+order if x is even, otherwise x remains unchanged.
func condAddOrderN(x *[numWords64 + 1]uint64) { _ = "STUB: not implemented"; return }

// div2subY update x = (x/2) - y.
func div2subY(x []uint64, y int64, l int) { _ = "STUB: not implemented"; return }

func (P *pointR1) fixedMult(scalar []byte) { _ = "STUB: not implemented"; return }

const (
	omegaFix = 7
	omegaVar = 5
)

// doubleMult returns P=mG+nQ.
func (P *pointR1) doubleMult(Q *pointR1, m, n []byte) { _ = "STUB: not implemented"; return }

// Generator point

// Variable input point
