// Package ed448 implements Ed448 signature scheme as described in RFC-8032.
//
// This package implements two signature variants.
//
//	| Scheme Name | Sign Function     | Verification  | Context           |
//	|-------------|-------------------|---------------|-------------------|
//	| Ed448       | Sign              | Verify        | Yes, can be empty |
//	| Ed448Ph     | SignPh            | VerifyPh      | Yes, can be empty |
//	| All above   | (PrivateKey).Sign | VerifyAny     | As above          |
//
// Specific functions for sign and verify are defined. A generic signing
// function for all schemes is available through the crypto.Signer interface,
// which is implemented by the PrivateKey type. A correspond all-in-one
// verification method is provided by the VerifyAny function.
//
// Both schemes require a context string for domain separation. This parameter
// is passed using a SignerOptions struct defined in this package.
//
// References:
//
//   - RFC8032: https://rfc-editor.org/rfc/rfc8032.txt
//   - EdDSA for more curves: https://eprint.iacr.org/2015/677
//   - High-speed high-security signatures: https://doi.org/10.1007/s13389-012-0027-1
package ed448

import (
	"crypto"
	"io"

	"github.com/cloudflare/circl/ecc/goldilocks"
	"github.com/cloudflare/circl/sign"
)

const (
	// ContextMaxSize is the maximum length (in bytes) allowed for context.
	ContextMaxSize = 255
	// PublicKeySize is the length in bytes of Ed448 public keys.
	PublicKeySize = 57
	// PrivateKeySize is the length in bytes of Ed448 private keys.
	PrivateKeySize = 114
	// SignatureSize is the length in bytes of signatures.
	SignatureSize = 114
	// SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.
	SeedSize = 57
)

const (
	paramB   = 456 / 8    // Size of keys in bytes.
	hashSize = 2 * paramB // Size of the hash function's output.
)

// SignerOptions implements crypto.SignerOpts and augments with parameters
// that are specific to the Ed448 signature schemes.
type SignerOptions struct {
	// Hash must be crypto.Hash(0) for both Ed448 and Ed448Ph.
	crypto.Hash

	// Context is an optional domain separation string for signing.
	// Its length must be less or equal than 255 bytes.
	Context string

	// Scheme is an identifier for choosing a signature scheme.
	Scheme SchemeID
}

// SchemeID is an identifier for each signature scheme.
type SchemeID uint

const (
	ED448 SchemeID = iota
	ED448Ph
)

// PublicKey is the type of Ed448 public keys.
type PublicKey []byte

// Equal reports whether pub and x have the same value.
func (pub PublicKey) Equal(x crypto.PublicKey) bool { _ = "STUB: not implemented"; return false }

// PrivateKey is the type of Ed448 private keys. It implements crypto.Signer.
type PrivateKey []byte

// Equal reports whether priv and x have the same value.
func (priv PrivateKey) Equal(x crypto.PrivateKey) bool { _ = "STUB: not implemented"; return false }

// Public returns the PublicKey corresponding to priv.
func (priv PrivateKey) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

// Seed returns the private key seed corresponding to priv. It is provided for
// interoperability with RFC 8032. RFC 8032's private keys correspond to seeds
// in this package.
func (priv PrivateKey) Seed() []byte { _ = "STUB: not implemented"; return nil }

func (priv PrivateKey) Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

func (pub PublicKey) Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

func (priv PrivateKey) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pub PublicKey) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sign creates a signature of a message given a key pair.
// This function supports all the two signature variants defined in RFC-8032,
// namely Ed448 (or pure EdDSA) and Ed448Ph.
// The opts.HashFunc() must return zero to the specify Ed448 variant. This can
// be achieved by passing crypto.Hash(0) as the value for opts.
// Use an Options struct to pass a bool indicating that the ed448Ph variant
// should be used.
// The struct can also be optionally used to pass a context string for signing.
func (priv PrivateKey) Sign(
	rand io.Reader,
	message []byte,
	opts crypto.SignerOpts,
) (signature []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(PublicKey), *new(PrivateKey), nil
}

// NewKeyFromSeed calculates a private key from a seed. It will panic if
// len(seed) is not SeedSize. This function is provided for interoperability
// with RFC 8032. RFC 8032's private keys correspond to seeds in this
// package.
func NewKeyFromSeed(seed []byte) PrivateKey { _ = "STUB: not implemented"; return *new(PrivateKey) }

func newKeyFromSeed(privateKey, seed []byte) { _ = "STUB: not implemented"; return }

func signAll(signature []byte, privateKey PrivateKey, message, ctx []byte, preHash bool) {
	_ = "STUB: not implemented"
	return
}

// 1.  Hash the 57-byte private key using SHAKE256(x, 114).

// 2.  Compute SHAKE256(dom4(F, C) || prefix || PH(M), 114).

// 3.  Compute the point [r]B.

// 4.  Compute SHAKE256(dom4(F, C) || R || A || PH(M), 114)

// 5.  Compute S = (r + k * s) mod order.

// 6.  The signature is the concatenation of R and S.

// Sign signs the message with privateKey and returns a signature.
// This function supports the signature variant defined in RFC-8032: Ed448,
// also known as the pure version of EdDSA.
// It will panic if len(privateKey) is not PrivateKeySize.
func Sign(priv PrivateKey, message []byte, ctx string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// SignPh creates a signature of a message given a keypair.
// This function supports the signature variant defined in RFC-8032: Ed448ph,
// meaning it internally hashes the message using SHAKE-256.
// Context could be passed to this function, which length should be no more than
// 255. It can be empty.
func SignPh(priv PrivateKey, message []byte, ctx string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func verify(public PublicKey, message, signature, ctx []byte, preHash bool) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyAny returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded.
// This function supports all the two signature variants defined in RFC-8032,
// namely Ed448 (or pure EdDSA) and Ed448Ph.
// The opts.HashFunc() must return zero, this can be achieved by passing
// crypto.Hash(0) as the value for opts.
// Use a SignerOptions struct to pass a context string for signing.
func VerifyAny(public PublicKey, message, signature []byte, opts crypto.SignerOpts) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded.
// This function supports the signature variant defined in RFC-8032: Ed448,
// also known as the pure version of EdDSA.
func Verify(public PublicKey, message, signature []byte, ctx string) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyPh returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded.
// This function supports the signature variant defined in RFC-8032: Ed448ph,
// meaning it internally hashes the message using SHAKE-256.
// Context could be passed to this function, which length should be no more than
// 255. It can be empty.
func VerifyPh(public PublicKey, message, signature []byte, ctx string) bool {
	_ = "STUB: not implemented"
	return false
}

func deriveSecretScalar(s *goldilocks.Scalar, h []byte) {
	_ = "STUB: not implemented"
	// The two least significant bits of the first octet are cleared,
	return
}

// all eight bits the last octet are cleared, and
// the highest bit of the second to last octet is set.

// isLessThanOrder returns true if 0 <= x < order and if the last byte of x is zero.
func isLessThanOrder(x []byte) bool { _ = "STUB: not implemented"; return false }

func writeDom(h io.Writer, ctx []byte, preHash bool) { _ = "STUB: not implemented"; return }
