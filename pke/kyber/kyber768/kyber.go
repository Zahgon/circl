// Code generated from pkg.templ.go. DO NOT EDIT.

// kyber768 implements the IND-CPA-secure Public Key Encryption
// scheme Kyber768.CPAPKE as submitted to round 3 of the NIST PQC competition
// and described in
//
// https://pq-crystals.org/kyber/data/kyber-specification-round3.pdf
package kyber768

import (
	"io"

	"github.com/cloudflare/circl/pke/kyber/kyber768/internal"
)

const (
	// Size of seed for NewKeyFromSeed
	KeySeedSize = internal.SeedSize

	// Size of seed for EncryptTo
	EncryptionSeedSize = internal.SeedSize

	// Size of a packed PublicKey
	PublicKeySize = internal.PublicKeySize

	// Size of a packed PrivateKey
	PrivateKeySize = internal.PrivateKeySize

	// Size of a ciphertext
	CiphertextSize = internal.CiphertextSize

	// Size of a plaintext
	PlaintextSize = internal.PlaintextSize
)

// PublicKey is the type of Kyber768.CPAPKE public key
type PublicKey internal.PublicKey

// PrivateKey is the type of Kyber768.CPAPKE private key
type PrivateKey internal.PrivateKey

// GenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func GenerateKey(rand io.Reader) (*PublicKey, *PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NewKeyFromSeed derives a public/private key pair using the given seed.
//
// Note: does not include the domain separation of ML-KEM (line 1, algorithm 13
// of FIPS 203). For that use NewKeyFromSeedMLKEM().
//
// Panics if seed is not of length KeySeedSize.
func NewKeyFromSeed(seed []byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewKeyFromSeedMLKEM derives a public/private key pair using the given seed
// using the domain separation of ML-KEM.
//
// Panics if seed is not of length KeySeedSize.
func NewKeyFromSeedMLKEM(seed []byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncryptTo encrypts message pt for the public key and writes the ciphertext
// to ct using randomness from seed.
//
// This function panics if the lengths of pt, seed, and ct are not
// PlaintextSize, EncryptionSeedSize, and CiphertextSize respectively.
func (pk *PublicKey) EncryptTo(ct []byte, pt []byte, seed []byte) {
	_ = "STUB: not implemented"
	return
}

// DecryptTo decrypts message ct for the private key and writes the
// plaintext to pt.
//
// This function panics if the lengths of ct and pt are not
// CiphertextSize and PlaintextSize respectively.
func (sk *PrivateKey) DecryptTo(pt []byte, ct []byte) { _ = "STUB: not implemented"; return }

// Packs pk into the given buffer.
//
// Panics if buf is not of length PublicKeySize.
func (pk *PublicKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Packs sk into the given buffer.
//
// Panics if buf is not of length PrivateKeySize.
func (sk *PrivateKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks pk from the given buffer.
//
// Panics if buf is not of length PublicKeySize.
func (pk *PublicKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks pk from the given buffer.
//
// Returns an error if the buffer is not of the right size, or the public
// key is not normalized.
func (pk *PublicKey) UnpackMLKEM(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Unpacks sk from the given buffer.
//
// Panics if buf is not of length PrivateKeySize.
func (sk *PrivateKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }

// Returns whether the two private keys are equal.
func (sk *PrivateKey) Equal(other *PrivateKey) bool { _ = "STUB: not implemented"; return false }
