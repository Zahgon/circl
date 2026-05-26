// Package frodo640shake implements the variant FrodoKEM-640 with SHAKE.
package frodo640shake

import (
	"io"

	"github.com/cloudflare/circl/kem"
)

const (
	paramN = 640

	// Denoted by 'mbar' in the FrodoKEM spec.
	paramNbar = 8

	logQ       = 15
	logQMask   = ((1 << logQ) - 1)
	seedASize  = 16
	pkHashSize = 16

	// Denoted by 'B' in the FrodoKEM spec.
	extractedBits = 2

	messageSize        = 16
	matrixBpPackedSize = (logQ * (paramN * paramNbar)) / 8
)

const (
	// Size of seed for NewKeyFromSeed.
	// = len(s) + len(seedSE) + len(z).
	KeySeedSize = SharedKeySize + SharedKeySize + 16

	// Size of seed for EncapsulateTo.
	EncapsulationSeedSize = 16

	// Size of the established shared key.
	SharedKeySize = 16

	// Size of the encapsulated shared key.
	CiphertextSize = 9720

	// Size of a packed public key.
	PublicKeySize = 9616

	// Size of a packed private key.
	PrivateKeySize = 19888
)

// Multi-dimensional arrays are stored in 1-dimensional arrays in
// row-major order.
type (
	nByNU16       [paramN * paramN]uint16
	nByNbarU16    [paramN * paramNbar]uint16
	nbarByNU16    [paramNbar * paramN]uint16
	nbarByNbarU16 [paramNbar * paramNbar]uint16
)

// Type of a FrodoKEM-640-SHAKE public key
type PublicKey struct {
	seedA   [seedASize]byte
	matrixB nByNbarU16
}

// Type of a FrodoKEM-640-SHAKE private key
type PrivateKey struct {
	hashInputIfDecapsFail [SharedKeySize]byte
	pk                    *PublicKey

	// matrixS stores transpose(S)
	matrixS nByNbarU16

	// H(packed(pk))
	hpk [pkHashSize]byte
}

// NewKeyFromSeed derives a public/private keypair deterministically
// from the given seed.
//
// Panics if seed is not of length KeySeedSize.
func newKeyFromSeed(seed []byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate the secret value s, and the seed for S, E, and A. Add seedA to the public key

// Populate the private key

// Add H(pk) to the private key

// GenerateKeyPair generates public and private keys using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func generateKeyPair(rand io.Reader) (*PublicKey, *PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EncapsulateTo generates a shared key and a ciphertext containing said key
// from the public key and the randomness from seed and writes the shared key
// to ss and ciphertext to ct.
//
// Panics if ss, ct, or seed are not of length SharedKeySize, CiphertextSize
// and EncapsulationSeedSize respectively.
//
// seed may be nil, in which case crypto/rand.Reader is used to generate one.
func (pk *PublicKey) EncapsulateTo(ct []byte, ss []byte, seed []byte) {
	_ = "STUB: not implemented"
	return
}

// compute hpk = G_1(packed(pk))

// compute (seedSE || k) = G_2(hpk || mu)

// Generate Sp, Ep, Epp, and A, and compute:
// Bp = Sp*A + Ep
// V = Sp*B + Epp

// Encode mu, and compute C = V + enc(mu) (mod q)

// Prepare the ciphertext

// Compute ss = F(ct||k)

// DecapsulateTo computes the shared key that is encapsulated in ct
// from the private key.
//
// Panics if ct or ss are not of length CiphertextSize and SharedKeySize
// respectively.
func (sk *PrivateKey) DecapsulateTo(ss, ct []byte) { _ = "STUB: not implemented"; return }

// Compute W = C - Bp*S (mod q), and decode the randomness mu

// Generate (seedSE' || k') = G_2(hpk || mu')

// Generate Sp, Ep, Epp, A, and compute BBp = Sp*A + Ep.

// Reduce BBp modulo q

// compute W = Sp*B + Epp

// Encode mu, and compute CC = W + enc(mu') (mod q)

// Prepare input to F

// If (Bp == BBp & C == CC) then ss = F(ct || k'), else ss = F(ct || s)
// Needs to avoid branching on secret data as per:
//     Qian Guo, Thomas Johansson, Alexander Nilsson. A key-recovery timing attack on post-quantum
//     primitives using the Fujisaki-Okamoto transformation and its application on FrodoKEM. In CRYPTO 2020.

// If (selector == 0) then load k' to do ss = F(ct || k'), else if (selector == 1) load s to do ss = F(ct || s)

// Packs sk to buf.
//
// Panics if buf is not of size PrivateKeySize.
func (sk *PrivateKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks sk from buf.
//
// Panics if buf is not of size PrivateKeySize.
func (sk *PrivateKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Packs pk to buf.
//
// Panics if buf is not of size PublicKeySize.
func (pk *PublicKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// TODO: Unpacks pk from buf.
//
// Panics if buf is not of size PublicKeySize.
func (pk *PublicKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Boilerplate down below for the KEM scheme API.

type scheme struct{}

var sch kem.Scheme = &scheme{}

// Scheme returns a KEM interface.
func Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (scheme) Name() string                { _ = "STUB: not implemented"; return "" }
func (*scheme) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (*scheme) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (*scheme) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (*scheme) SharedKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (*scheme) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (*scheme) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (sk *PrivateKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (pk *PublicKey) Scheme() kem.Scheme  { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (sk *PrivateKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sk *PrivateKey) Equal(other kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (pk *PublicKey) Equal(other kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (sk *PrivateKey) Public() kem.PublicKey { _ = "STUB: not implemented"; return *new(kem.PublicKey) }

func (pk *PublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (*scheme) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (*scheme) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (*scheme) Encapsulate(pk kem.PublicKey) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (*scheme) EncapsulateDeterministically(
	pk kem.PublicKey, seed []byte,
) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (*scheme) Decapsulate(sk kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*scheme) UnmarshalBinaryPublicKey(buf []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

func (*scheme) UnmarshalBinaryPrivateKey(buf []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}
