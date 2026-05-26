//go:build (!purego && arm64) || (!purego && amd64)
// +build !purego,arm64 !purego,amd64

package p384

import (
	"math/big"
)

// affinePoint represents an affine point of the curve. The point at
// infinity is (0,0) leveraging that it is not an affine point.
type affinePoint struct{ x, y fp384 }

func newAffinePoint(x, y *big.Int) *affinePoint { _ = "STUB: not implemented"; return nil }

func zeroPoint() *affinePoint { _ = "STUB: not implemented"; return nil }

func (ap affinePoint) String() string { _ = "STUB: not implemented"; return "" }

func (ap *affinePoint) isZero() bool { _ = "STUB: not implemented"; return false }

func (ap *affinePoint) neg() { _ = "STUB: not implemented"; return }

func (ap *affinePoint) toInt() (x, y *big.Int) { _ = "STUB: not implemented"; return nil, nil }

func (ap *affinePoint) toJacobian() *jacobianPoint { _ = "STUB: not implemented"; return nil }

func (ap *affinePoint) toProjective() *projectivePoint { _ = "STUB: not implemented"; return nil }

// OddMultiples calculates the points iP for i={1,3,5,7,..., 2^(n-1)-1}
// Ensure that 1 < n < 31, otherwise it returns an empty slice.
func (ap affinePoint) oddMultiples(n uint) []jacobianPoint { _ = "STUB: not implemented"; return nil }

// oddMultiplesProjective calculates the points iP for i={1,3,5,7,..., 2^(n-1)-1}
// Ensure that 1 < n < 31, otherwise it returns an empty slice.
func (ap affinePoint) oddMultiplesProjective(n uint) []projectivePoint {
	_ = "STUB: not implemented"
	return nil
}

// p2Point is a point in P^2
type p2Point struct{ x, y, z fp384 }

func (P *p2Point) String() string { _ = "STUB: not implemented"; return "" }

func (P *p2Point) neg() { _ = "STUB: not implemented"; return }

// condNeg if P is negated if b=1.
func (P *p2Point) cneg(b int) { _ = "STUB: not implemented"; return }

// cmov sets P to Q if b=1.
func (P *p2Point) cmov(Q *p2Point, b int) { _ = "STUB: not implemented"; return }

func (P *p2Point) toInt() (x, y, z *big.Int) { _ = "STUB: not implemented"; return nil, nil, nil }

// jacobianPoint represents a point in Jacobian coordinates. The point at
// infinity is any point (x,y,0) such that x and y are different from 0.
type jacobianPoint struct{ p2Point }

func (P *jacobianPoint) isZero() bool { _ = "STUB: not implemented"; return false }

func (P *jacobianPoint) toAffine() *affinePoint { _ = "STUB: not implemented"; return nil }

// add calculates P=Q+R such that Q and R are different than the identity point.
// This function can be used for doublings.
func (P *jacobianPoint) add(Q, R *jacobianPoint) { _ = "STUB: not implemented"; return }

// Cohen-Miyagi-Ono (1998)
// https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#addition-add-1998-cmo-2

// Z1Z1 = Z1 ^ 2
// Z2Z2 = Z2 ^ 2
// U1 = X1 * Z2Z2
// U2 = X2 * Z1Z1
// t0 = Z2 * Z2Z2
// S1 = Y1 * t0
// t1 = Z1 * Z1Z1
// S2 = Y2 * t1
// H = U2 - U1
// HH = H ^ 2
// HHH = H * HH
// r = S2 - S1

// Detect point doubling

// V = U1 * HH
// t2 = r ^ 2
// t3 = V + V
// t4 = t2 - HHH
// X3 = t4 - t3
// t5 = V - X3
// t6 = S1 * HHH
// t7 = r * t5
// Y3 = t7 - t6
// t8 = Z2 * H
// Z3 = Z1 * t8

// mixadd calculates P=Q+R such that P and Q different than the identity point,
// and Q not in {P,-P, O}.
func (P *jacobianPoint) mixadd(Q *jacobianPoint, R *affinePoint) {
	_ = "STUB: not implemented"
	// See https://www.hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#addition-madd-2007-bl
	return
}

func (P *jacobianPoint) double() {
	_ = "STUB: not implemented"
	// See https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#doubling-dbl-2001-b
	return
}

func (P *jacobianPoint) toProjective() *projectivePoint { _ = "STUB: not implemented"; return nil }

// projectivePoint represents a point in projective homogeneous coordinates.
// The point at infinity is (0,y,0) such that y is different from 0.
type projectivePoint struct{ p2Point }

func (P *projectivePoint) isZero() bool { _ = "STUB: not implemented"; return false }

func (P *projectivePoint) cmov(Q *projectivePoint, b int) { _ = "STUB: not implemented"; return }

func (P *projectivePoint) toAffine() *affinePoint { _ = "STUB: not implemented"; return nil }

// add calculates P=Q+R using complete addition formula for prime groups.
func (P *projectivePoint) completeAdd(Q, R *projectivePoint) {
	_ = "STUB: not implemented"
	// Reference:
	//
	//	"Complete addition formulas for prime order elliptic curves" by
	//	Costello-Renes-Batina. [Alg.4] (eprint.iacr.org/2015/1060).
	return
}

// 1.  t0 ← X1 · X2
// 2.  t1 ← Y1 · Y2
// 3.  t2 ← Z1 · Z2
// 4.  t3 ← X1 + Y1
// 5.  t4 ← X2 + Y2
// 6.  t3 ← t3 · t4
// 7.  t4 ← t0 + t1
// 8.  t3 ← t3 − t4
// 9.  t4 ← Y1 + Z1
// 10. X3 ← Y2 + Z2
// 11. t4 ← t4 · X3
// 12. X3 ← t1 + t2
// 13. t4 ← t4 − X3
// 14. X3 ← X1 + Z1
// 15. Y3 ← X2 + Z2
// 16. X3 ← X3 · Y3
// 17. Y3 ← t0 + t2
// 18. Y3 ← X3 − Y3
// 19. Z3 ←  b · t2
// 20. X3 ← Y3 − Z3
// 21. Z3 ← X3 + X3
// 22. X3 ← X3 + Z3
// 23. Z3 ← t1 − X3
// 24. X3 ← t1 + X3
// 25. Y3 ←  b · Y3
// 26. t1 ← t2 + t2
// 27. t2 ← t1 + t2
// 28. Y3 ← Y3 − t2
// 29. Y3 ← Y3 − t0
// 30. t1 ← Y3 + Y3
// 31. Y3 ← t1 + Y3
// 32. t1 ← t0 + t0
// 33. t0 ← t1 + t0
// 34. t0 ← t0 − t2
// 35. t1 ← t4 · Y3
// 36. t2 ← t0 · Y3
// 37. Y3 ← X3 · Z3
// 38. Y3 ← Y3 + t2
// 39. X3 ← t3 · X3
// 40. X3 ← X3 − t1
// 41. Z3 ← t4 · Z3
// 42. t1 ← t3 · t0
// 43. Z3 ← Z3 + t1
