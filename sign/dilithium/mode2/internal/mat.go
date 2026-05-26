// Code generated from mode3/internal/mat.go by gen.go

package internal

import (
	common "github.com/cloudflare/circl/sign/internal/dilithium"
)

// A k by l matrix of polynomials.
type Mat [K]VecL

// Expands the given seed to a complete matrix.
//
// This function is called ExpandA in the specification.
func (m *Mat) Derive(seed *[32]byte) { _ = "STUB: not implemented"; return }

// Set p to the inner product of a and b using pointwise multiplication.
//
// Assumes a and b are in Montgomery form and their coefficients are
// pairwise sufficiently small to multiply, see Poly.MulHat().  Resulting
// coefficients are bounded by 2Lq.
func PolyDotHat(p *common.Poly, a, b *VecL) { _ = "STUB: not implemented"; return }

// zero p
