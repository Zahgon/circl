// Package eddilithium3 implements the hybrid signature scheme Ed448-Dilithium3.
package eddilithium3

import (
	"crypto"
	"io"

	"github.com/cloudflare/circl/sign"
	"github.com/cloudflare/circl/sign/dilithium/mode3"
	"github.com/cloudflare/circl/sign/ed448"
)

const (
	// SeedSize is the length of the seed for NewKeyFromSeed
	SeedSize = ed448.SeedSize // > mode3.SeedSize

	// PublicKeySize is the length in bytes of the packed public key.
	PublicKeySize = mode3.PublicKeySize + ed448.PublicKeySize

	// PrivateKeySize is the length in bytes of the packed public key.
	PrivateKeySize = mode3.PrivateKeySize + ed448.SeedSize

	// SignatureSize is the length in bytes of the signatures.
	SignatureSize = mode3.SignatureSize + ed448.SignatureSize
)

// PublicKey is the type of an EdDilithium3 public key.
type PublicKey struct {
	e ed448.PublicKey
	d mode3.PublicKey
}

// PrivateKey is the type of an EdDilithium3 private key.
type PrivateKey struct {
	e ed448.PrivateKey
	d mode3.PrivateKey
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
