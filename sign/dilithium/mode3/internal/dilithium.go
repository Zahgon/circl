package internal

import (
	"io"

	common "github.com/cloudflare/circl/sign/internal/dilithium"
)

const (
	// Size of a packed polynomial of norm ≤η.
	// (Note that the  formula is not valid in general.)
	PolyLeqEtaSize = (common.N * DoubleEtaBits) / 8

	// β = τη, the maximum size of c s₂.
	Beta = Tau * Eta

	// γ₁ range of y
	Gamma1 = 1 << Gamma1Bits

	// Size of packed polynomial of norm <γ₁ such as z
	PolyLeGamma1Size = (Gamma1Bits + 1) * common.N / 8

	// α = 2γ₂ parameter for decompose
	Alpha = 2 * Gamma2

	// Size of a packed private key
	PrivateKeySize = 32 + 32 + TRSize + PolyLeqEtaSize*(L+K) + common.PolyT0Size*K

	// Size of a packed public key
	PublicKeySize = 32 + common.PolyT1Size*K

	// Size of a packed signature
	SignatureSize = L*PolyLeGamma1Size + Omega + K + CTildeSize

	// Size of packed w₁
	PolyW1Size = (common.N * (common.QBits - Gamma1Bits)) / 8
)

// PublicKey is the type of Dilithium public keys.
type PublicKey struct {
	rho [32]byte
	t1  VecK

	// Cached values
	t1p [common.PolyT1Size * K]byte
	A   *Mat
	tr  [TRSize]byte
}

// PrivateKey is the type of Dilithium private keys.
type PrivateKey struct {
	rho [32]byte
	key [32]byte
	s1  VecL
	s2  VecK
	t0  VecK
	tr  [TRSize]byte

	// Cached values
	A   Mat  // ExpandA(ρ)
	s1h VecL // NTT(s₁)
	s2h VecK // NTT(s₂)
	t0h VecK // NTT(t₀)

	seed    [common.SeedSize]byte
	seedSet bool
}

type unpackedSignature struct {
	z    VecL
	hint VecK
	c    [CTildeSize]byte
}

// Packs the signature into buf.
func (sig *unpackedSignature) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Sets sig to the signature encoded in the buffer.
//
// Returns whether buf contains a properly packed signature.
func (sig *unpackedSignature) Unpack(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Packs the public key into buf.
func (pk *PublicKey) Pack(buf *[PublicKeySize]byte) { _ = "STUB: not implemented"; return }

// Sets pk to the public key encoded in buf.
func (pk *PublicKey) Unpack(buf *[PublicKeySize]byte) { _ = "STUB: not implemented"; return }

// tr = CRH(ρ ‖ t1) = CRH(pk)

// Packs the private key into buf.
func (sk *PrivateKey) Pack(buf *[PrivateKeySize]byte) { _ = "STUB: not implemented"; return }

// Sets sk to the private key encoded in buf.
func (sk *PrivateKey) Unpack(buf *[PrivateKeySize]byte) { _ = "STUB: not implemented"; return }

// Cached values

// GenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func GenerateKey(rand io.Reader) (*PublicKey, *PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NewKeyFromSeed derives a public/private key pair using the given seed.
func NewKeyFromSeed(seed *[common.SeedSize]byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	// expanded seed
	return nil, nil
}

// Complete public key far enough to be packed

// Finish private key

// tr = CRH(ρ ‖ t1) = CRH(pk)

// Finish cache of public key

func (sk *PrivateKey) Seed() []byte { _ = "STUB: not implemented"; return nil }

// Computes t0 and t1 from sk.s1h, sk.s2 and sk.A.
func (sk *PrivateKey) computeT0andT1(t0, t1 *VecK) {
	_ = "STUB: not implemented"

	// Set t to A s₁ + s₂
	return
}

// Compute t₀, t₁ = Power2Round(t)

// Verify checks whether the given signature by pk on msg is valid.
//
// For Dilithium this is the top-level verification function.
// In ML-DSA, this is ML-DSA.Verify_internal.
func Verify(pk *PublicKey, msg func(io.Writer), signature []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Note that Unpack() checked whether ‖z‖_∞ < γ₁ - β
// and ensured that there at most ω ones in pk.hint.

// μ = CRH(tr ‖ msg)

// Compute Az

// Next, we compute Az - 2ᵈ·c·t₁.
// Note that the coefficients of t₁ are bounded by 256 = 2⁹,
// so the coefficients of Az2dct1 will bounded by 2⁹⁺ᵈ = 2²³ < 2q,
// which is small enough for NTT().

// UseHint(pk.hint, Az - 2ᵈ·c·t₁)
//    = UseHint(pk.hint, w - c·s₂ + c·t₀)
//    = UseHint(pk.hint, r + c·t₀)
//    = r₁ = w₁.

// c' = H(μ, w₁)

// SignTo signs the given message and writes the signature into signature.
//
// For Dilithium this is the top-level signing function. For ML-DSA
// this is ML-DSA.Sign_internal.
//
//nolint:funlen
func SignTo(sk *PrivateKey, msg func(io.Writer), rnd [32]byte, signature []byte) {
	_ = "STUB: not implemented"
	return
}

//  μ = CRH(tr ‖ msg)

// ρ' = CRH(key ‖ μ)

// Main rejection loop

// Depending on the mode, one try has a chance between 1/7 and 1/4
// of succeeding.  Thus it is safe to say that 576 iterations
// are enough as (6/7)⁵⁷⁶ < 2⁻¹²⁸.

// y = ExpandMask(ρ', key)

// Set w to A y

// Decompose w into w₀ and w₁

// c~ = H(μ ‖ w₁)

// Ensure ‖ w₀ - c·s2 ‖_∞ < γ₂ - β.
//
// By Lemma 3 of the specification this is equivalent to checking that
// both ‖ r₀ ‖_∞ < γ₂ - β and r₁ = w₁, for the decomposition
// w - c·s₂	 = r₁ α + r₀ as computed by decompose().
// See also §4.1 of the specification.

// z = y + c·s₁

// Ensure  ‖z‖_∞ < γ₁ - β

// Compute c·t₀

// Ensure ‖c·t₀‖_∞ < γ₂.

// Create the hint to be able to reconstruct w₁ from w - c·s₂ + c·t0.
// Note that we're not using makeHint() in the obvious way as we
// do not know whether ‖ sc·s₂ - c·t₀ ‖_∞ < γ₂.  Instead we note
// that our makeHint() is actually the same as a makeHint for a
// different decomposition:
//
// Earlier we ensured indirectly with a check that r₁ = w₁ where
// r = w - c·s₂.  Hence r₀ = r - r₁ α = w - c·s₂ - w₁ α = w₀ - c·s₂.
// Thus  MakeHint(w₀ - c·s₂ + c·t₀, w₁) = MakeHint(r0 + c·t₀, r₁)
// and UseHint(w - c·s₂ + c·t₀, w₁) = UseHint(r + c·t₀, r₁).
// As we just ensured that ‖ c·t₀ ‖_∞ < γ₂ our usage is correct.

// Computes the public key corresponding to this private key.
func (sk *PrivateKey) Public() *PublicKey { _ = "STUB: not implemented"; return nil }

// Equal returns whether the two public keys are equal
func (pk *PublicKey) Equal(other *PublicKey) bool { _ = "STUB: not implemented"; return false }

// Equal returns whether the two private keys are equal
func (sk *PrivateKey) Equal(other *PrivateKey) bool { _ = "STUB: not implemented"; return false }
