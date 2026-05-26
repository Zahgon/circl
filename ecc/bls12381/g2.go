package bls12381

import (
	"github.com/cloudflare/circl/ecc/bls12381/ff"
)

// G2Size is the length in bytes of an element in G2 in uncompressed form..
const G2Size = 2 * ff.Fp2Size

// G2SizeCompressed is the length in bytes of an element in G2 in compressed form.
const G2SizeCompressed = ff.Fp2Size

// G2 is a point in the twist of the BLS12 curve over Fp2.
type G2 struct{ x, y, z ff.Fp2 }

func (g G2) String() string { _ = "STUB: not implemented"; return "" }

// Bytes serializes a G2 element in uncompressed form.
func (g G2) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Bytes serializes a G2 element in compressed form.
func (g G2) BytesCompressed() []byte { _ = "STUB: not implemented"; return nil }

// SetBytes sets g to the value in bytes, and returns a non-nil error if not in G2.
func (g *G2) SetBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

// Check for invalid prefixes

func (g G2) encodeBytes(compressed bool) []byte { _ = "STUB: not implemented"; return nil }

// Neg inverts g.
func (g *G2) Neg() {
	_ = "STUB: not implemented"

	// SetIdentity assigns g to the identity element.
	return
}

func (g *G2) SetIdentity() { _ = "STUB: not implemented"; return }

// isValidProjective returns true if the point is not a projective point.
func (g *G2) isValidProjective() bool { _ = "STUB: not implemented"; return false }

// IsOnG2 returns true if the point is in the group G2.
func (g *G2) IsOnG2() bool { _ = "STUB: not implemented"; return false }

// IsIdentity return true if the point is the identity of G2.
func (g *G2) IsIdentity() bool { _ = "STUB: not implemented"; return false }

// cmov sets g to P if b == 1
func (g *G2) cmov(P *G2, b int) { _ = "STUB: not implemented"; return }

// isRTorsion returns true if point is in the r-torsion subgroup.
func (g *G2) isRTorsion() bool {
	_ = "STUB: not implemented"
	// Bowe, "Faster Subgroup Checks for BLS12-381" (https://eprint.iacr.org/2019/814)
	return false
}

// Q = \psi(g)
// Q = -[z]\psi(g)
// Q = -[z]\psi(g)+g
// Q = -[z]\psi^2(g)+\psi(g)
// Q = -[z]\psi^3(g)+\psi^2(g)

// Equivalent to verification equation in paper

// psi is the Galbraith-Scott endomorphism. See https://eprint.iacr.org/2008/117.
func (g *G2) psi() { _ = "STUB: not implemented"; return }

// clearCofactor maps g to a point in the r-torsion subgroup.
//
// This method multiplies g times a multiple of the cofactor as proposed by
// Fuentes-Knapp-Rodríguez at https://doi.org/10.1007/978-3-642-28496-0_25.
//
// The explicit formulas for BLS curves are in Section 4.1 of Budroni-Pintore
// "Efficient hash maps to G2 on BLS curves" at https://eprint.iacr.org/2017/419
//
//	h(a)P = [x^2-x-1]P + [x-1]ψ(P) + ψ^2(2P)
func (g *G2) clearCofactor() { _ = "STUB: not implemented"; return }

// 2P
// ψ(2P)
// ψ^2(2P)
// -xP
// -xP + P = [-x+1]P
//
// ψ(-xP + P) = [-x+1]ψ(P)
// x^2P - xP = [x^2-x]P
// P + [-x+1]ψ(P)
// -P + [x-1]ψ(P)
// [x^2-x-1]P + [x-1]ψ(P)
// [x^2-x-1]P + [x-1]ψ(P) + 2ψ^2(P)

// Double updates g = 2g.
func (g *G2) Double() { _ = "STUB: not implemented"; return }

// Add updates g=P+Q.
func (g *G2) Add(P, Q *G2) { _ = "STUB: not implemented"; return }

// ScalarMult calculates g = kP.
func (g *G2) ScalarMult(k *Scalar, P *G2) { _ = "STUB: not implemented"; return }

// scalarMult calculates g = kP, where k is the scalar in big-endian order.
func (g *G2) scalarMult(k []byte, P *G2) { _ = "STUB: not implemented"; return }

// scalarMultShort multiplies by a short, constant scalar k, where k is the
// scalar in big-endian order. Runtime depends on the scalar.
func (g *G2) scalarMultShort(k []byte, P *G2) {
	_ = "STUB: not implemented"
	// Since the scalar is short and low Hamming weight not much helps.
	return
}

// IsEqual returns true if g and p are equivalent.
func (g *G2) IsEqual(p *G2) bool { _ = "STUB: not implemented"; return false }

// lx = x1*z2
// rx = x2*z1
// lx = lx-rx
// ly = y1*z2
// ry = y2*z1
// ly = ly-ry

// EncodeToCurve is a non-uniform encoding from an input byte string (and
// an optional domain separation tag) to elements in G2. This function must not
// be used as a hash function, otherwise use G2.Hash instead.
func (g *G2) Encode(input, dst []byte) { _ = "STUB: not implemented"; return }

// Hash produces an element of G2 from the hash of an input byte string and
// an optional domain separation tag. This function is safe to use when a
// random oracle returning points in G2 be required.
func (g *G2) Hash(input, dst []byte) { _ = "STUB: not implemented"; return }

// isOnCurve returns true if g is a valid point on the curve.
func (g *G2) isOnCurve() bool { _ = "STUB: not implemented"; return false }

// y2 = y^2
// y2 = y^2*z
// x3 = x^2
// x3 = x^3
// z3 = z^2
// z3 = z^3
// z3 = (4+4i)*z^3
// x3 = x^3 + (4+4i)*z^3
// y2 = y^2*z - (x^3 + (4+4i)*z^3)

// toAffine updates g with its affine representation.
func (g *G2) toAffine() { _ = "STUB: not implemented"; return }

// G2Generator returns the generator point of G2.
func G2Generator() *G2 { _ = "STUB: not implemented"; return nil }
