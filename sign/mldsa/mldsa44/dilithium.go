// Code generated from pkg.templ.go. DO NOT EDIT.

// mldsa44 implements NIST signature scheme ML-DSA-44 as defined in FIPS204.
package mldsa44

import (
	"crypto"
	"encoding/asn1"
	"io"

	"github.com/cloudflare/circl/sign"
	common "github.com/cloudflare/circl/sign/internal/dilithium"
	"github.com/cloudflare/circl/sign/mldsa/mldsa44/internal"
)

const (
	// Size of seed for NewKeyFromSeed
	SeedSize = common.SeedSize

	// Size of a packed PublicKey
	PublicKeySize = internal.PublicKeySize

	// Size of a packed PrivateKey
	PrivateKeySize = internal.PrivateKeySize

	// Size of a signature
	SignatureSize = internal.SignatureSize
)

// PublicKey is the type of ML-DSA-44 public key
type PublicKey internal.PublicKey

// PrivateKey is the type of ML-DSA-44 private key
type PrivateKey internal.PrivateKey

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
//
// ctx is the optional context string. Errors if ctx is larger than 255 bytes.
// A nil context string is equivalent to an empty context string.
func SignTo(sk *PrivateKey, msg, ctx []byte, randomized bool, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not use. Implements ML-DSA.Sign_internal used for compatibility tests.
func (sk *PrivateKey) unsafeSignInternal(msg []byte, rnd [32]byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Do not use. Implements ML-DSA.Verify_internal used for compatibility tests.
func unsafeVerifyInternal(pk *PublicKey, msg, sig []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify checks whether the given signature by pk on msg is valid.
//
// ctx is the optional context string. Fails if ctx is larger than 255 bytes.
// A nil context string is equivalent to an empty context string.
func Verify(pk *PublicKey, msg, ctx, sig []byte) bool { _ = "STUB: not implemented"; return false }

// Sets pk to the public key encoded in buf.
func (pk *PublicKey) Unpack(buf *[PublicKeySize]byte) { _ = "STUB: not implemented"; return }

// Sets sk to the private key encoded in buf.
func (sk *PrivateKey) Unpack(buf *[PrivateKeySize]byte) { _ = "STUB: not implemented"; return }

// Packs the public key into buf.
func (pk *PublicKey) Pack(buf *[PublicKeySize]byte) { _ = "STUB: not implemented"; return }

// Packs the private key into buf.
func (sk *PrivateKey) Pack(buf *[PrivateKeySize]byte) { _ = "STUB: not implemented"; return }

// Packs the public key.
func (pk *PublicKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Packs the private key.
func (sk *PrivateKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Packs the public key.
func (pk *PublicKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Packs the private key.
		nil
}

func (sk *PrivateKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unpacks the public key from data.
		nil
}

func (pk *PublicKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// Unpacks the private key from data.
func (sk *PrivateKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// Returns seed used to generate PrivateKey, and nil if not retained.
func (sk *PrivateKey) Seed() []byte { _ = "STUB: not implemented"; return nil }

// Sign signs the given message.
//
// opts.HashFunc() must return zero, which can be achieved by passing
// crypto.Hash(0) or nil for opts.  rand is ignored.  Will only return an error
// if opts.HashFunc() is non-zero.
//
// This function is used to make PrivateKey implement the crypto.Signer
// interface.  The package-level SignTo function might be more convenient
// to use.
func (sk *PrivateKey) Sign(rand io.Reader, msg []byte, opts crypto.SignerOpts) (
	sig []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Computes the public key corresponding to this private key.
//
// Returns a *PublicKey.  The type crypto.PublicKey is used to make
// PrivateKey implement the crypto.Signer interface.
func (sk *PrivateKey) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

// Equal returns whether the two private keys equal.
func (sk *PrivateKey) Equal(other crypto.PrivateKey) bool { _ = "STUB: not implemented"; return false }

// Equal returns whether the two public keys equal.
func (pk *PublicKey) Equal(other crypto.PublicKey) bool { _ = "STUB: not implemented"; return false }

// Boilerplate for generic signatures API

type scheme struct{}

var sch sign.Scheme = &scheme{}

// Scheme returns a generic signature interface for ML-DSA-44.
func Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

func (*scheme) Name() string        { _ = "STUB: not implemented"; return "" }
func (*scheme) PublicKeySize() int  { _ = "STUB: not implemented"; return 0 }
func (*scheme) PrivateKeySize() int { _ = "STUB: not implemented"; return 0 }
func (*scheme) SignatureSize() int  { _ = "STUB: not implemented"; return 0 }
func (*scheme) SeedSize() int       { _ = "STUB: not implemented"; return 0 }

// TODO TLSIdentifier()
func (*scheme) Oid() asn1.ObjectIdentifier {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier)
}

func (*scheme) SupportsContext() bool { _ = "STUB: not implemented"; return false }

func (*scheme) GenerateKey() (sign.PublicKey, sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), *new(sign.PrivateKey), nil
}

func (*scheme) Sign(
	sk sign.PrivateKey,
	msg []byte,
	opts *sign.SignatureOpts,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (*scheme) Verify(
	pk sign.PublicKey,
	msg, sig []byte,
	opts *sign.SignatureOpts,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (*scheme) DeriveKey(seed []byte) (sign.PublicKey, sign.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), *new(sign.PrivateKey)
}

func (*scheme) UnmarshalBinaryPublicKey(buf []byte) (sign.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), nil
}

func (*scheme) UnmarshalBinaryPrivateKey(buf []byte) (sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PrivateKey), nil
}

func (sk *PrivateKey) Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

func (sk *PublicKey) Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }
