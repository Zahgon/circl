package internal

// A k by k matrix of polynomials.
type Mat [K]Vec

// Expands the given seed to the corresponding matrix A or its transpose Aᵀ.
func (m *Mat) Derive(seed *[32]byte, transpose bool) { _ = "STUB: not implemented"; return }

// If there is just one left, then a plain DeriveUniform
// is quicker than the X4 variant.

// Transposes A in place.
func (m *Mat) Transpose() { _ = "STUB: not implemented"; return }
