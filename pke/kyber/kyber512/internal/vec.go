package internal

import (
	"github.com/cloudflare/circl/pke/kyber/internal/common"
)

// A vector of K polynomials
type Vec [K]common.Poly

// Samples v[i] from a centered binomial distribution with given η,
// seed and nonce+i.
//
// Essentially CBD_η(PRF(seed, nonce+i)) from the specification.
func (v *Vec) DeriveNoise(seed []byte, nonce uint8, eta int) { _ = "STUB: not implemented"; return }

// Sets p to the inner product of a and b using "pointwise" multiplication.
//
// See MulHat() and NTT() for a description of the multiplication.
// Assumes a and b are in Montgomery form.  p will be in Montgomery form,
// and its coefficients will be bounded in absolute value by 2kq.
// If a and b are not in Montgomery form, then the action is the same
// as "pointwise" multiplication followed by multiplying by R⁻¹, the inverse
// of the Montgomery factor.
func PolyDotHat(p *common.Poly, a, b *Vec) { _ = "STUB: not implemented"; return }

// set p to zero

// Almost normalizes coefficients in-place.
//
// Ensures each coefficient is in {0, …, q}.
func (v *Vec) BarrettReduce() { _ = "STUB: not implemented"; return }

// Normalizes coefficients in-place.
//
// Ensures each coefficient is in {0, …, q-1}.
func (v *Vec) Normalize() { _ = "STUB: not implemented"; return }

// Applies in-place inverse NTT().  See Poly.InvNTT() for assumptions.
func (v *Vec) InvNTT() { _ = "STUB: not implemented"; return }

// Applies in-place forward NTT().  See Poly.NTT() for assumptions.
func (v *Vec) NTT() { _ = "STUB: not implemented"; return }

// Sets v to a + b.
func (v *Vec) Add(a, b *Vec) { _ = "STUB: not implemented"; return }

// Packs v into buf, which must be of length K*PolySize.
func (v *Vec) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks v from buf which must be of length K*PolySize.
func (v *Vec) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Writes Compress_q(v, d) to m.
//
// Assumes v is normalized and d is in {3, 4, 5, 10, 11}.
func (v *Vec) CompressTo(m []byte, d int) { _ = "STUB: not implemented"; return }

// Set v to Decompress_q(m, 1).
//
// Assumes d is in {3, 4, 5, 10, 11}.  v will be normalized.
func (v *Vec) Decompress(m []byte, d int) { _ = "STUB: not implemented"; return }

// ⌈(256 d)/8⌉
func compressedPolySize(d int) int { _ = "STUB: not implemented"; return 0 }
