// Package eddilithium2 implements the hybrid signature scheme Ed25519-Dilithium2.
package eddilithium2

import (
	"crypto"
	"io"

	"github.com/cloudflare/circl/sign"
	"github.com/cloudflare/circl/sign/dilithium/mode2"
	"github.com/cloudflare/circl/sign/ed25519"
)

const (
	// SeedSize is the length of the seed for NewKeyFromSeed
	SeedSize = mode2.SeedSize // = ed25519.SeedSize = 32

	// PublicKeySize is the length in bytes of the packed public key.
	PublicKeySize = mode2.PublicKeySize + ed25519.PublicKeySize

	// PrivateKeySize is the length in bytes of the packed public key.
	PrivateKeySize = mode2.PrivateKeySize + ed25519.SeedSize

	// SignatureSize is the length in bytes of the signatures.
	SignatureSize = mode2.SignatureSize + ed25519.SignatureSize
)

// PublicKey is the type of an EdDilithium2 public key.
type PublicKey struct {
	e ed25519.PublicKey
	d mode2.PublicKey
}

// PrivateKey is the type of an EdDilithium2 private key.
type PrivateKey struct {
	e ed25519.PrivateKey
	d mode2.PrivateKey
}

// GenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func GenerateKey(rand io.Reader) (*PublicKey, *PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NewKeyFromSeed derives a public/private key pair using the given seed.
func NewKeyFromSeed(seed *[SeedSize]byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Internally, Ed25519 and Dilithium hash the seeds they are passed again
// with different hash functions, so it would be safe to use exactly the
// same seed for Ed25519 and Dilithium here.  However, in general, when
// combining any two signature schemes it might not be the case that this
// is safe.  Setting a bad example here isn't worth the tiny gain in
// performance.

// SignTo signs the given message and writes the signature into signature.
// It will panic if signature is not of length at least SignatureSize.
func SignTo(sk *PrivateKey, msg []byte, signature []byte) { _ = "STUB: not implemented"; return }

// Verify checks whether the given signature by pk on msg is valid.
func Verify(pk *PublicKey, msg []byte, signature []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Unpack unpacks pk to the public key encoded in buf.
func (pk *PublicKey) Unpack(buf *[PublicKeySize]byte) { _ = "STUB: not implemented"; return }

// Unpack sets sk to the private key encoded in buf.
func (sk *PrivateKey) Unpack(buf *[PrivateKeySize]byte) { _ = "STUB: not implemented"; return }

// Pack packs the public key into buf.
func (pk *PublicKey) Pack(buf *[PublicKeySize]byte) { _ = "STUB: not implemented"; return }

// Pack packs the private key into buf.
func (sk *PrivateKey) Pack(buf *[PrivateKeySize]byte) { _ = "STUB: not implemented"; return }

// Bytes packs the public key.
func (pk *PublicKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Bytes packs the private key.
func (sk *PrivateKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// MarshalBinary packs the public key.
func (pk *PublicKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalBinary packs the private key.
		nil
}

func (sk *PrivateKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary the public key from data.
		nil
}

func (pk *PublicKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBinary unpacks the private key from data.
func (sk *PrivateKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (sk *PrivateKey) Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }
func (pk *PublicKey) Scheme() sign.Scheme  { _ = "STUB: not implemented"; return *new(sign.Scheme) }

func (sk *PrivateKey) Equal(other crypto.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (pk *PublicKey) Equal(other crypto.PublicKey) bool { _ = "STUB: not implemented"; return false }

// Sign signs the given message.
//
// opts.HashFunc() must return zero, which can be achieved by passing
// crypto.Hash(0) for opts.  rand is ignored.  Will only return an error
// if opts.HashFunc() is non-zero.
//
// This function is used to make PrivateKey implement the crypto.Signer
// interface.  The package-level SignTo function might be more convenient
// to use.
func (sk *PrivateKey) Sign(
	rand io.Reader, msg []byte, opts crypto.SignerOpts,
) (signature []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Public computes the public key corresponding to this private key.
//
// Returns a *PublicKey.  The type crypto.PublicKey is used to make
// PrivateKey implement the crypto.Signer interface.
func (sk *PrivateKey) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}
