// Code generated from mode3/internal/pack.go by gen.go

package internal

import (
	common "github.com/cloudflare/circl/sign/internal/dilithium"
)

// Writes p with norm less than or equal η into buf, which must be of
// size PolyLeqEtaSize.
//
// Assumes coefficients of p are not normalized, but in [q-η,q+η].
func PolyPackLeqEta(p *common.Poly, buf []byte) { _ = "STUB: not implemented"; return }

// compiler eliminates branch

// Sets p to the polynomial of norm less than or equal η encoded in the
// given buffer of size PolyLeqEtaSize.
//
// Output coefficients of p are not normalized, but in [q-η,q+η] provided
// buf was created using PackLeqEta.
//
// Beware, for arbitrary buf the coefficients of p might end up in
// the interval [q-2^b,q+2^b] where b is the least b with η≤2^b.
func PolyUnpackLeqEta(p *common.Poly, buf []byte) { _ = "STUB: not implemented"; return }

// compiler eliminates branch

// Writes v with coefficients in {0, 1} of which at most ω non-zero
// to buf, which must have length ω+k.
func (v *VecK) PackHint(buf []byte) {
	_ = "STUB: not implemented"
	// The packed hint starts with the indices of the non-zero coefficients
	// For instance:
	//
	//	(x⁵⁶ + x¹⁰⁰, x²⁵⁵, 0, x² + x²³, x¹)
	//
	// Yields
	//
	//	56, 100, 255, 2, 23, 1
	//
	// Then we pad with zeroes until we have a list of ω items:
	// //  56, 100, 255, 2, 23, 1, 0, 0, ..., 0
	//
	// Then we finish with a list of the switch-over-indices in this
	// list between polynomials, so:
	//
	//	56, 100, 255, 2, 23, 1, 0, 0, ..., 0, 2, 3, 3, 5, 6
	return
}

// Sets v to the vector encoded using VecK.PackHint()
//
// Returns whether unpacking was successful.
func (v *VecK) UnpackHint(buf []byte) bool {
	_ = "STUB: not implemented"
	// A priori, there would be several reasonable ways to encode the same
	// hint vector.  We take care to only allow only one encoding, to ensure
	// "strong unforgeability".
	//
	// See PackHint() source for description of the encoding.
	return false
}

// zero v
// previous switch-over-point

// ensures switch-over-points are increasing

// ensures indices are increasing (within a poly)

// ensures padding indices are zero

// Sets p to the polynomial packed into buf by PolyPackLeGamma1.
//
// p will be normalized.
func PolyUnpackLeGamma1(p *common.Poly, buf []byte) { _ = "STUB: not implemented"; return }

// coefficients in [0,…,2γ₁)
// (-γ₁,…,γ₁]

// normalize

// Writes p whose coefficients are in (-γ₁,γ₁] into buf
// which has to be of length PolyLeGamma1Size.
//
// Assumes p is normalized.
func PolyPackLeGamma1(p *common.Poly, buf []byte) { _ = "STUB: not implemented"; return }

// coefficients in [0,…,γ₁] ∪ (q-γ₁,…,q)

// [0,…,γ₁] ∪ (γ₁-q,…,2γ₁-q)
// [0,…,2γ₁)

// Coefficients are in [0, γ₁] ∪ (Q-γ₁, Q)

// Pack w₁ into buf, which must be of length PolyW1Size.
//
// Assumes w₁ is normalized.
func PolyPackW1(p *common.Poly, buf []byte) { _ = "STUB: not implemented"; return }
