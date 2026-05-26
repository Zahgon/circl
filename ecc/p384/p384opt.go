//go:build (!purego && arm64) || (!purego && amd64)
// +build !purego,arm64 !purego,amd64

package p384

import (
	"math/big"
)

type curve struct{}

// P384 returns a Curve which implements P-384 (see FIPS 186-3, section D.2.4).
func P384() Curve {
	_ = "STUB: not implemented"

	// IsOnCurve reports whether the given (x,y) lies on the curve.
	return *new(Curve)
}

func (c curve) IsOnCurve(x, y *big.Int) bool { _ = "STUB: not implemented"; return false }

// Add returns the sum of (x1,y1) and (x2,y2).
func (c curve) Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Double returns 2*(x,y).
func (c curve) Double(x1, y1 *big.Int) (x, y *big.Int) { _ = "STUB: not implemented"; return nil, nil }

// reduceScalar shorten a scalar modulo the order of the curve.
func (c curve) reduceScalar(k []byte) []byte { _ = "STUB: not implemented"; return nil }

// toOdd performs k = (-k mod N) if k is even.
func (c curve) toOdd(k []byte) ([]byte, int) { _ = "STUB: not implemented"; return nil, 0 }

// ScalarMult returns (Qx,Qy)=k*(Px,Py) where k is a number in big-endian form.
func (c curve) ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c curve) scalarMultOmega(x1, y1 *big.Int, k []byte, omega uint) (x, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate the last iteration using complete addition formula.

// ScalarBaseMult returns k*G, where G is the base point of the group
// and k is an integer in big-endian form.
func (c curve) ScalarBaseMult(k []byte) (x, y *big.Int) { _ = "STUB: not implemented"; return nil, nil }

// CombinedMult calculates P=mG+nQ, where G is the generator and Q=(x,y,z).
// The scalars m and n are integers in big-endian form. Non-constant time.
func (c curve) CombinedMult(xQ, yQ *big.Int, m, n []byte) (xP, yP *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generator point

// Input point

// absolute returns always a positive value.
func absolute(x int32) int32 { _ = "STUB: not implemented"; return 0 }
