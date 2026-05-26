package fourq

import (
	"math/big"
)

// Size of scalars used for point multiplication.
const Size = 32

// Point represents an affine point of the curve. The identity is (0,1).
type Point struct{ X, Y Fq }

func (P Point) String() string { _ = "STUB: not implemented"; return "" }

// CurveParams contains the parameters of the elliptic curve.
type CurveParams struct {
	Name string   // The canonical name of the curve.
	P    *big.Int // The order of the underlying field Fp.
	N    *big.Int // The order of the generator point.
	G    Point    // This is the generator point.
}

// Params returns the parameters for the curve.
func Params() *CurveParams { _ = "STUB: not implemented"; return nil }

// IsOnCurve reports whether the given P=(x,y) lies on the curve.
func (P *Point) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

// SetGenerator assigns to P the generator point G.
func (P *Point) SetGenerator() { _ = "STUB: not implemented"; return }

// SetIdentity assigns to P the identity element.
func (P *Point) SetIdentity() { _ = "STUB: not implemented"; return }

// IsIdentity returns true if P is the identity element.
func (P *Point) IsIdentity() bool { _ = "STUB: not implemented"; return false }

// Add calculates a point addition P = Q + R.
func (P *Point) Add(Q, R *Point) { _ = "STUB: not implemented"; return }

// ScalarMult calculates P = k*Q, where Q is an N-torsion point.
func (P *Point) ScalarMult(k *[Size]byte, Q *Point) { _ = "STUB: not implemented"; return }

// ScalarBaseMult calculates P = k*G, where G is the generator point.
func (P *Point) ScalarBaseMult(k *[Size]byte) { _ = "STUB: not implemented"; return }

func (P *Point) fromR1(Q *pointR1) { _ = "STUB: not implemented"; return }

func (P *Point) toR1(projP *pointR1) { _ = "STUB: not implemented"; return }
