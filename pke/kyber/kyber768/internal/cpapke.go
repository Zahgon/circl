// Code generated from kyber512/internal/cpapke.go by gen.go

package internal

// A Kyber.CPAPKE private key.
type PrivateKey struct {
	sh Vec // NTT(s), normalized
}

// A Kyber.CPAPKE public key.
type PublicKey struct {
	rho [32]byte // ρ, the seed for the matrix A
	th  Vec      // NTT(t), normalized

	// cached values
	aT Mat // the matrix Aᵀ
}

// Packs the private key to buf.
func (sk *PrivateKey) Pack(buf []byte) {
	_ = "STUB: not implemented"

	// Unpacks the private key from buf.
	return
}

func (sk *PrivateKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Packs the public key to buf.
func (pk *PublicKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks the public key from buf. Checks if the public key is normalized.
func (pk *PublicKey) UnpackMLKEM(buf []byte) error {
	_ = "STUB: not implemented"

	// FIPS 203 §7.2 "encapsulation key check" (2).
	return nil
}

// Unpacks the public key from buf.
func (pk *PublicKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Derives a new Kyber.CPAPKE keypair from the given seed.
func NewKeyFromSeed(seed []byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This writes hash into expandedSeed.  Yes, this is idiomatic Go.

// σ, the noise seed

// Expand ρ to matrix A; we'll transpose later

// Sample secret vector s

// Sample blind e

// Next, we compute t = A s + e.

// Note that coefficients of s are bounded by q and those of A
// are bounded by 4.5q and so their product is bounded by 2¹⁵q
// as required for multiplication.

// A and s were not in Montgomery form, so the Montgomery
// multiplications in the inner product added a factor R⁻¹ which
// we'll cancel out now.  This will also ensure the coefficients of
// t are bounded in absolute value by q.

// bounded by 8q.

// Decrypts ciphertext ct meant for private key sk to plaintext pt.
func (sk *PrivateKey) DecryptTo(pt, ct []byte) { _ = "STUB: not implemented"; return }

// Compute m = v - <s, u>

// Compress polynomial m to original message

// Encrypts message pt for the public key to ciphertext ct using randomness
// from seed.
//
// seed has to be of length SeedSize, pt of PlaintextSize and ct of
// CiphertextSize.
func (pk *PublicKey) EncryptTo(ct, pt, seed []byte) { _ = "STUB: not implemented"; return }

// Sample r, e₁ and e₂ from B_η

// Next we compute u = Aᵀ r + e₁.  First Aᵀ.

// Note that coefficients of r are bounded by q and those of Aᵀ
// are bounded by 4.5q and so their product is bounded by 2¹⁵q
// as required for multiplication.

// Aᵀ and r were not in Montgomery form, so the Montgomery
// multiplications in the inner product added a factor R⁻¹ which
// the InvNTT cancels out.

// u = Aᵀ r + e₁

// Next compute v = <t, r> + e₂ + Decompress_q(m, 1).

// v = <t, r> + e₂ + Decompress_q(m, 1)

// Pack ciphertext

// Returns whether sk equals other.
func (sk *PrivateKey) Equal(other *PrivateKey) bool { _ = "STUB: not implemented"; return false }
