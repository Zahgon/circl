// Code generated from mode3/internal/rounding.go by gen.go

package internal

import (
	common "github.com/cloudflare/circl/sign/internal/dilithium"
)

// Splits 0 ≤ a < q into a₀ and a₁ with a = a₁*α + a₀ with -α/2 < a₀ ≤ α/2,
// except for when we would have a₁ = (q-1)/α in which case a₁=0 is taken
// and -α/2 ≤ a₀ < 0.  Returns a₀ + q.  Note 0 ≤ a₁ < (q-1)/α.
// Recall α = 2γ₂.
func decompose(a uint32) (a0plusQ, a1 uint32) {
	_ = "STUB: not implemented"
	// a₁ = ⌈a / 128⌉
	return 0, 0
}

// 1025/2²² is close enough to 1/4092 so that a₁
// becomes a/α rounded down.

// For the corner-case a₁ = (q-1)/α = 16, we have to set a₁=0.

// 1488/2²⁴ is close enough to 1/1488 so that a₁
// becomes a/α rounded down.

// For the corner-case a₁ = (q-1)/α = 44, we have to set a₁=0.

// In the corner-case, when we set a₁=0, we will incorrectly
// have a₀ > (q-1)/2 and we'll need to subtract q.  As we
// return a₀ + q, that comes down to adding q if a₀ < (q-1)/2.

// Assume 0 ≤ r, f < Q with ‖f‖_∞ ≤ α/2.  Decompose r as r = r1*α + r0 as
// computed by decompose().  Write r' := r - f (mod Q).  Now, decompose
// r'=r-f again as  r' = r'1*α + r'0 using decompose().  As f is small, we
// have r'1 = r1 + h, where h ∈ {-1, 0, 1}.  makeHint() computes |h|
// given z0 := r0 - f (mod Q) and r1.  With |h|, which is called the hint,
// we can reconstruct r1 using only r' = r - f, which is done by useHint().
// To wit:
//
//	useHint( r - f, makeHint( r0 - f, r1 ) ) = r1.
//
// Assumes 0 ≤ z0 < Q.
func makeHint(z0, r1 uint32) uint32 {
	_ = "STUB: not implemented"
	// If -α/2 < r0 - f ≤ α/2, then r1*α + r0 - f is a valid decomposition of r'
	// with the restrictions of decompose() and so r'1 = r1.  So the hint
	// should be 0. This is covered by the first two inequalities.
	// There is one other case: if r0 - f = -α/2, then r1*α + r0 - f is also
	// a valid decomposition if r1 = 0.  In the other cases a one is carried
	// and the hint should be 1.
	return 0
}

// Uses the hint created by makeHint() to reconstruct r1 from r'=r-f; see
// documentation of makeHint() for context.
// Assumes 0 ≤ r' < Q.
func useHint(rp uint32, hint uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// Sets p to the hint polynomial for p0 the modified low bits and p1
// the unmodified high bits --- see makeHint().
//
// Returns the number of ones in the hint polynomial.
func PolyMakeHint(p, p0, p1 *common.Poly) (pop uint32) { _ = "STUB: not implemented"; return 0 }

// Computes corrections to the high bits of the polynomial q according
// to the hints in h and sets p to the corrected high bits.  Returns p.
func PolyUseHint(p, q, hint *common.Poly) { _ = "STUB: not implemented"; return }

// See useHint() and makeHint() for an explanation.  We reimplement it
// here so that we can call Poly.Decompose(), which might be way faster
// than calling decompose() in a loop (for instance when having AVX2.)

// Splits each of the coefficients of p using decompose.
func PolyDecompose(p, p0PlusQ, p1 *common.Poly) { _ = "STUB: not implemented"; return }
