// Package goldilocks provides elliptic curve operations over the goldilocks curve.
package goldilocks

// Curve is the Goldilocks curve x^2+y^2=z^2-39081x^2y^2.
type Curve struct{}

// Identity returns the identity point.
func (Curve) Identity() *Point { _ = "STUB: not implemented"; return nil }

// IsOnCurve returns true if the point lies on the curve.
func (Curve) IsOnCurve(P *Point) bool { _ = "STUB: not implemented"; return false }

// Check z != 0

// t = ta*tb
// x^2
// y^2
// z^2
// t^2
// x^2 + y^2
// dt^2
// z^2 + dt^2
// x^2 + y^2 - (z^2 + dt^2)

// xy
// tz
// xy - tz

// Generator returns the generator point.
func (Curve) Generator() *Point { _ = "STUB: not implemented"; return nil }

// Order returns the number of points in the prime subgroup.
func (Curve) Order() Scalar {
	_ = "STUB: not implemented"

	// Double returns 2P.
	return *new(Scalar)
}

func (Curve) Double(P *Point) *Point { _ = "STUB: not implemented"; return nil }

// Add returns P+Q.
func (Curve) Add(P, Q *Point) *Point { _ = "STUB: not implemented"; return nil }

// ScalarMult returns kP. This function runs in constant time.
func (e Curve) ScalarMult(k *Scalar, P *Point) *Point { _ = "STUB: not implemented"; return nil }

// ScalarBaseMult returns kG where G is the generator point. This function runs in constant time.
func (e Curve) ScalarBaseMult(k *Scalar) *Point { _ = "STUB: not implemented"; return nil }

// CombinedMult returns mG+nP, where G is the generator point. This function is non-constant time.
func (e Curve) CombinedMult(m, n *Scalar, P *Point) *Point { _ = "STUB: not implemented"; return nil }
