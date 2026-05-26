package dilithium

// Sets p to the polynomial whose coefficients are less than 1024 encoded
// into buf (which must be of size PolyT1Size).
//
// p will be normalized.
func (p *Poly) UnpackT1(buf []byte) { _ = "STUB: not implemented"; return }

// Writes p whose coefficients are in (-2ᵈ⁻¹, 2ᵈ⁻¹] into buf which
// has to be of length at least PolyT0Size.
//
// Assumes that the coefficients are not normalized, but lie in the
// range (q-2ᵈ⁻¹, q+2ᵈ⁻¹].
func (p *Poly) PackT0(buf []byte) { _ = "STUB: not implemented"; return }

// Sets p to the polynomial packed into buf by PackT0.
//
// The coefficients of p will not be normalized, but will lie
// in (-2ᵈ⁻¹, 2ᵈ⁻¹].
func (p *Poly) UnpackT0(buf []byte) { _ = "STUB: not implemented"; return }

// Writes p whose coefficients are less than 1024 into buf, which must be
// of size at least PolyT1Size .
//
// Assumes coefficients of p are normalized.
func (p *Poly) PackT1(buf []byte) { _ = "STUB: not implemented"; return }

// Writes p whose coefficients are in [0, 16) to buf, which must be of
// length N/2.
func (p *Poly) packLe16Generic(buf []byte) { _ = "STUB: not implemented"; return }

// Writes p with 60 non-zero coefficients {-1,1} to buf, which must have
// length 40.
func (p *Poly) PackB60(buf []byte) {
	_ = "STUB: not implemented"
	// We start with a mask of the non-zero positions of p (which is 32 bytes)
	// and then append 60 packed bits, where a one indicates a negative
	// coefficients.
	return
}

// UnpackB60 sets p to the polynomial packed into buf with Poly.PackB60().
//
// Returns whether unpacking was successful.
func (p *Poly) UnpackB60(buf []byte) bool {
	_ = "STUB: not implemented"
	// zero p
	return false
}

// ensure unused bits are zero for strong unforgeability

// Note 1 ^ (1 | (Q-1)) = Q-1 and (-1)&x = x
