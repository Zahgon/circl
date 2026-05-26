package dilithium

// An element of our base ring R which are polynomials over Z_q modulo
// the equation Xᴺ = -1, where q=2²³ - 2¹³ + 1 and N=256.
//
// Coefficients aren't always reduced.  See Normalize().
type Poly [N]uint32

// Reduces each of the coefficients to <2q.
func (p *Poly) reduceLe2QGeneric() { _ = "STUB: not implemented"; return }

// Reduce each of the coefficients to <q.
func (p *Poly) normalizeGeneric() { _ = "STUB: not implemented"; return }

// Normalize the coefficients in this polynomial assuming they are already
// bounded by 2q.
func (p *Poly) normalizeAssumingLe2QGeneric() { _ = "STUB: not implemented"; return }

// Sets p to a + b.  Does not normalize polynomials.
func (p *Poly) addGeneric(a, b *Poly) { _ = "STUB: not implemented"; return }

// Sets p to a - b.
//
// Warning: assumes coefficients of b are less than 2q.
func (p *Poly) subGeneric(a, b *Poly) { _ = "STUB: not implemented"; return }

// Checks whether the "supnorm" (see sec 2.1 of the spec) of p is equal
// or greater than the given bound.
//
// Requires the coefficients of p to be normalized.
func (p *Poly) exceedsGeneric(bound uint32) bool {
	_ = "STUB: not implemented"
	// Note that we are allowed to leak which coefficients break the bound,
	// but not their sign.
	return false
}

// The central. reps. of {0,       1, ..., (Q-1)/2,  (Q+1)/2, ..., Q-1}
// are given by          {0,       1, ..., (Q-1)/2, -(Q-1)/2, ...,  -1}
// so their norms are    {0,       1, ..., (Q-1)/2,  (Q-1)/2, ...,   1}.
// We'll compute them in a different way though.

// Sets x to             {(Q-1)/2, (Q-3)/2, ..., 0, -1, ..., -(Q-1)/2}

// Sets x to             {(Q-1)/2, (Q-3)/2, ..., 0, 0, ...,  (Q-3)/2}

// Sets x to             {0,       1, ...,  (Q-1)/2, (Q-1)/2, ..., 1}

// Splits p into p1 and p0 such that [i]p1 * 2ᴰ + [i]p0 = [i]p
// with -2ᴰ⁻¹ < [i]p0 ≤ 2ᴰ⁻¹.  Returns p0 + Q and p1.
//
// Requires the coefficients of p to be normalized.
func (p *Poly) power2RoundGeneric(p0PlusQ, p1 *Poly) { _ = "STUB: not implemented"; return }

// Sets p to the polynomial whose coefficients are the pointwise multiplication
// of those of a and b.  The coefficients of p are bounded by 2q.
//
// Assumes a and b are in Montgomery form and that the pointwise product
// of each coefficient is below 2³² q.
func (p *Poly) mulHatGeneric(a, b *Poly) { _ = "STUB: not implemented"; return }

// Sets p to 2ᵈ q without reducing.
//
// So it requires the coefficients of p  to be less than 2³²⁻ᴰ.
func (p *Poly) mulBy2toDGeneric(q *Poly) { _ = "STUB: not implemented"; return }
