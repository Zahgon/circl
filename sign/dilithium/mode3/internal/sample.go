package internal

import (
	common "github.com/cloudflare/circl/sign/internal/dilithium"
	"github.com/cloudflare/circl/simd/keccakf1600"
)

// DeriveX4Available indicates whether the system supports the quick fourway
// sampling variants like PolyDeriveUniformX4.
var DeriveX4Available = keccakf1600.IsEnabledX4()

// For each i, sample ps[i] uniformly from the given seed and nonces[i].
// ps[i] may be nil and is ignored in that case.
//
// Can only be called when DeriveX4Available is true.
func PolyDeriveUniformX4(ps [4]*common.Poly, seed *[32]byte, nonces [4]uint16) {
	_ = "STUB: not implemented"
	return
}

// Absorb the seed in the four states

// Absorb the nonces, the SHAKE128 domain separator (0b1111), the
// start of the padding (0b...001) and the end of the padding 0b100...
// Recall that the rate of SHAKE128 is 168 --- i.e. 21 uint64s.

// indices into ps

// mark nil polynomial as completed

// Applies KeccaK-f[1600] to state to get the next 21 uint64s of each
// of the four SHAKE128 streams.

// Sample p uniformly from the given seed and nonce.
//
// p will be normalized.
func PolyDeriveUniform(p *common.Poly, seed *[32]byte, nonce uint16) {
	_ = "STUB: not implemented"
	return
}

// fits 168B SHAKE-128 rate

// Note that 3 divides into 168 and 12*16, so we use up buf completely.

// We use rejection sampling

// 32 byte seed + uint16 nonce

// Sample p uniformly with coefficients of norm less than or equal η,
// using the given seed and nonce.
//
// p will not be normalized, but will have coefficients in [q-η,q+η].
func PolyDeriveUniformLeqEta(p *common.Poly, seed *[64]byte, nonce uint16) {
	_ = "STUB: not implemented"
	// Assumes 2 < η < 8.
	return
}

// fits 136B SHAKE-256 rate

// We use rejection sampling

// branch is eliminated by compiler

// reduce mod  5

// reduce mod 5

// 64 byte seed + uint16 nonce

// 136 is SHAKE-256 rate

// Sample v[i] uniformly with coefficients in (-γ₁,…,γ₁]  using the
// given seed and nonce+i
//
// p will be normalized.
func VecLDeriveUniformLeGamma1(v *VecL, seed *[64]byte, nonce uint16) {
	_ = "STUB: not implemented"
	return
}

// Sample p uniformly with coefficients in (-γ₁,…,γK1s] using the
// given seed and nonce.
//
// p will be normalized.
func PolyDeriveUniformLeGamma1(p *common.Poly, seed *[64]byte, nonce uint16) {
	_ = "STUB: not implemented"
	return
}

// For each i, sample ps[i] uniformly with τ non-zero coefficients in {q-1,1}
// using the given seed and w1[i].  ps[i] may be nil and is ignored
// in that case.  ps[i] will be normalized.
//
// Can only be called when DeriveX4Available is true.
//
// This function is currently not used (yet).
func PolyDeriveUniformBallX4(ps [4]*common.Poly, seed []byte) { _ = "STUB: not implemented"; return }

// Absorb the seed in the four states

// SHAKE256 domain separator and padding

// indices into ps

// zero ps[j]

// mark as completed

// Takes least significant bit of signs and uses it for the sign.
// Note 1 ^ (1 | (Q-1)) = Q-1.

// Samples p uniformly with τ non-zero coefficients in {q-1,1}.
//
// The polynomial p will be normalized.
func PolyDeriveUniformBall(p *common.Poly, seed []byte) {
	_ = "STUB: not implemented"
	// SHAKE-256 rate is 136
	return
}

// Essentially we generate a sequence of τ ones or minus ones,
// prepend 196 zeroes and shuffle the concatenation using the
// usual algorithm (Fisher--Yates.)

// offset into buf

// zero p

// Find location of where to move the new coefficient to using
// rejection sampling.

// Takes least significant bit of signs and uses it for the sign.
// Note 1 ^ (1 | (Q-1)) = Q-1.
