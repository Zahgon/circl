package bls12381

import (
	_ "crypto/sha256"

	"github.com/cloudflare/circl/ecc/bls12381/ff"
)

// G1Size is the length in bytes of an element in G1 in uncompressed form..
const G1Size = 2 * ff.FpSize

// G1SizeCompressed is the length in bytes of an element in G1 in compressed form.
const G1SizeCompressed = ff.FpSize

// G1 is a point in the BLS12 curve over Fp.
type G1 struct{ x, y, z ff.Fp }

func (g G1) String() string { _ = "STUB: not implemented"; return "" }

// Bytes serializes a G1 element in uncompressed form.
func (g G1) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Bytes serializes a G1 element in compressed form.
func (g G1) BytesCompressed() []byte { _ = "STUB: not implemented"; return nil }

// SetBytes sets g to the value in bytes, and returns a non-nil error if not in G1.
func (g *G1) SetBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

// Check for invalid prefixes

func (g G1) encodeBytes(compressed bool) []byte { _ = "STUB: not implemented"; return nil }

// Neg inverts g.
func (g *G1) Neg() {
	_ = "STUB: not implemented"

	// SetIdentity assigns g to the identity element.
	return
}

func (g *G1) SetIdentity() { _ = "STUB: not implemented"; return }

// isValidProjective returns true if the point is not a projective point.
func (g *G1) isValidProjective() bool { _ = "STUB: not implemented"; return false }

// IsOnG1 returns true if the point is in the group G1.
func (g *G1) IsOnG1() bool { _ = "STUB: not implemented"; return false }

// IsIdentity return true if the point is the identity of G1.
func (g *G1) IsIdentity() bool { _ = "STUB: not implemented"; return false }

// cmov sets g to P if b == 1
func (g *G1) cmov(P *G1, b int) { _ = "STUB: not implemented"; return }

// sigma is an edomorphism defined by (x, y) → (βx, y) for some β ∈ Fp of
// multiplicative order 3.
func (g *G1) sigma(P *G1) { _ = "STUB: not implemented"; return }

// sigma2 is sigma(sigma(P)).
func (g *G1) sigma2(P *G1) { _ = "STUB: not implemented"; return }

// isRTorsion returns true if point is in the r-torsion subgroup.
func (g *G1) isRTorsion() bool {
	_ = "STUB: not implemented"
	// Bowe, "Faster Subgroup Checks for BLS12-381" (https://eprint.iacr.org/2019/814)
	return false
}

// s(P)
// 2*s(P)
// s(s(P))
// P + s(s(P))
// -P - s(s(P))
// 2*s(P) - P - s(s(P))
// coef * [2*s(P) - P - s(s(P))]
// -s(s(P))
// coef * [2*s(P) - P - s(s(P))] - s(s(P))

// clearCofactor maps g to a point in the r-torsion subgroup.
//
// This method multiplies g times (1-z) rather than (z-1)^2/3, where z is the
// BLS12 parameter. This is enough to remove points of order
//
//	h \in {3, 11, 10177, 859267, 52437899},
//
// and because there are no points of order h^2. See Section 5 of Wahby-Boneh
// "Fast and simple constant-time hashing to the BLS12-381 elliptic curve" at
// https://eprint.iacr.org/2019/403
func (g *G1) clearCofactor() { _ = "STUB: not implemented"; return }

// Double updates g = 2g.
func (g *G1) Double() {
	_ = "STUB: not implemented"
	// Reference:
	//
	//	"Complete addition formulas for prime order elliptic curves" by
	//	Costello-Renes-Batina. [Alg.9] (eprint.iacr.org/2015/1060).
	return
}

// 1.  t0 =  Y * Y
// 2.  Z3 = t0 + t0
// 3.  Z3 = Z3 + Z3
// 4.  Z3 = Z3 + Z3
// 5.  t1 =  Y * Z
// 6.  t2 =  Z * Z
// 7.  t2 = b3 * t2
// 8.  X3 = t2 * Z3
// 9.  Y3 = t0 + t2
// 10. Z3 = t1 * Z3
// 11. t1 = t2 + t2
// 12. t2 = t1 + t2
// 13. t0 = t0 - t2
// 14. Y3 = t0 * Y3
// 15. Y3 = X3 + Y3
// 16. t1 =  X * Y
// 17. X3 = t0 * t1
// 18. X3 = X3 + X3

// Add updates g=P+Q.
func (g *G1) Add(P, Q *G1) {
	_ = "STUB: not implemented"
	// Reference:
	//
	//	"Complete addition formulas for prime order elliptic curves" by
	//	Costello-Renes-Batina. [Alg.7] (eprint.iacr.org/2015/1060).
	return
}

// 1.  t0 = X1 * X2
// 2.  t1 = Y1 * Y2
// 3.  t2 = Z1 * Z2
// 4.  t3 = X1 + Y1
// 5.  t4 = X2 + Y2
// 6.  t3 = t3 * t4
// 7.  t4 = t0 + t1
// 8.  t3 = t3 - t4
// 9.  t4 = Y1 + Z1
// 10. X3 = Y2 + Z2
// 11. t4 = t4 * X3
// 12. X3 = t1 + t2
// 13. t4 = t4 - X3
// 14. X3 = X1 + Z1
// 15. Y3 = X2 + Z2
// 16. X3 = X3 * Y3
// 17. Y3 = t0 + t2
// 18. Y3 = X3 - Y3
// 19. X3 = t0 + t0
// 20. t0 = X3 + t0
// 21. t2 = b3 * t2
// 22. Z3 = t1 + t2
// 23. t1 = t1 - t2
// 24. Y3 = b3 * Y3
// 25. X3 = t4 * Y3
// 26. t2 = t3 * t1
// 27. X3 = t2 - X3
// 28. Y3 = Y3 * t0
// 29. t1 = t1 * Z3
// 30. Y3 = t1 + Y3
// 31. t0 = t0 * t3
// 32. Z3 = Z3 * t4
// 33. Z3 = Z3 + t0

// ScalarMult calculates g = kP.
func (g *G1) ScalarMult(k *Scalar, P *G1) { _ = "STUB: not implemented"; return }

// scalarMult calculates g = kP, where k is the scalar in big-endian order.
func (g *G1) scalarMult(k []byte, P *G1) { _ = "STUB: not implemented"; return }

// scalarMultShort multiplies by a short, constant scalar k, where k is the
// scalar in big-endian order. Runtime depends on the scalar.
func (g *G1) scalarMultShort(k []byte, P *G1) {
	_ = "STUB: not implemented"
	// Since the scalar is short and low Hamming weight not much helps.
	return
}

// IsEqual returns true if g and p are equivalent.
func (g *G1) IsEqual(p *G1) bool { _ = "STUB: not implemented"; return false }

// lx = x1*z2
// rx = x2*z1
// lx = lx-rx
// ly = y1*z2
// ry = y2*z1
// ly = ly-ry

// isOnCurve returns true if g is a valid point on the curve.
func (g *G1) isOnCurve() bool { _ = "STUB: not implemented"; return false }

// y2 = y^2
// y2 = y^2*z
// x3 = x^2
// x3 = x^3
// z3 = z^2
// z3 = z^3
// z3 = 4*z^3
// x3 = x^3 + 4*z^3
// y2 = y^2*z - (x^3 + 4*z^3)

// toAffine updates g with its affine representation.
func (g *G1) toAffine() { _ = "STUB: not implemented"; return }

// EncodeToCurve is a non-uniform encoding from an input byte string (and
// an optional domain separation tag) to elements in G1. This function must not
// be used as a hash function, otherwise use G1.Hash instead.
func (g *G1) Encode(input, dst []byte) { _ = "STUB: not implemented"; return }

// Hash produces an element of G1 from the hash of an input byte string and
// an optional domain separation tag. This function is safe to use when a
// random oracle returning points in G1 be required.
func (g *G1) Hash(input, dst []byte) { _ = "STUB: not implemented"; return }

// G1Generator returns the generator point of G1.
func G1Generator() *G1 { _ = "STUB: not implemented"; return nil }

// affinize converts an entire slice to affine at once
// handling the points at the infinity.
func affinize(points []*G1) (out []G1) { _ = "STUB: not implemented"; return nil }
