package common

import (
	"github.com/cloudflare/circl/simd/keccakf1600"
)

// DeriveX4Available indicates whether the system supports the quick fourway
// sampling variants like PolyDeriveUniformX4.
var DeriveX4Available = keccakf1600.IsEnabledX4()

// Samples p from a centered binomial distribution with given η.
//
// Essentially CBD_η(PRF(seed, nonce)) from the specification.
func (p *Poly) DeriveNoise(seed []byte, nonce uint8, eta int) { _ = "STUB: not implemented"; return }

// Sample p from a centered binomial distribution with n=6 and p=½ - that is:
// coefficients are in {-3, -2, -1, 0, 1, 2, 3} with probabilities {1/64, 3/32,
// 15/64, 5/16, 16/64, 3/32, 1/64}.
func (p *Poly) DeriveNoise3(seed []byte, nonce uint8) { _ = "STUB: not implemented"; return }

// The distribution at hand is exactly the same as that
// of (a₁ + a₂ + a₃) - (b₁ + b₂+b₃) where a_i,b_i~U(1).  Thus we need
// 6 bits per coefficients, thus 192 bytes of input entropy.

// We add two extra zero bytes in the buffer to be able to read 8 bytes
// at the same time (while using only 6.)

// t is interpreted as a₁ + 2a₂ + 4a₃ + 8b₁ + 16b₂ + ….

// a₁ + 8b₁ + …
// a₁ + a₂ + 8(b₁ + b₂) + …
// a₁ + a₂ + a₃ + 4(b₁ + b₂ + b₃) + …

// a₁ + a₂ + a₃

// b₁ + b₂ + b₃

// Sample p from a centered binomial distribution with n=4 and p=½ - that is:
// coefficients are in {-2, -1, 0, 1, 2} with probabilities {1/16, 1/4,
// 3/8, 1/4, 1/16}.
func (p *Poly) DeriveNoise2(seed []byte, nonce uint8) { _ = "STUB: not implemented"; return }

// The distribution at hand is exactly the same as that
// of (a + a') - (b + b') where a,a',b,b'~U(1).  Thus we need 4 bits per
// coefficients, thus 128 bytes of input entropy.

// t is interpreted as a + 2a' + 4b + 8b' + ….

// a + 4b + …
// a+a' + 4(b + b') + …

// For each i, sample ps[i] uniformly from the given seed for coordinates
// xs[i] and ys[i]. ps[i] may be nil and is ignored in that case.
//
// Can only be called when DeriveX4Available is true.
func PolyDeriveUniformX4(ps [4]*Poly, seed *[32]byte, xs, ys [4]uint8) {
	_ = "STUB: not implemented"
	return
}

// Absorb the seed in the four states

// Absorb the coordinates, the SHAKE128 domain separator (0b1111), the
// start of the padding (0b…001) and the end of the padding 0b100….
// Recall that the rate of SHAKE128 is 168; ie. 21 uint64s.

// indices into ps

// mark nil polynomials as completed

// Applies KeccaK-f[1600] to state to get the next 21 uint64s of each of
// the four SHAKE128 streams.

// Sample p uniformly from the given seed and x and y coordinates.
//
// Coefficients are reduced and will be in "tangled" order.  See Tangle().
func (p *Poly) DeriveUniform(seed *[32]byte, x, y uint8) { _ = "STUB: not implemented"; return }

// rate of SHAKE-128
