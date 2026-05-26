// Package ed25519 implements Ed25519 signature scheme as described in RFC-8032.
//
// This package provides optimized implementations of the three signature
// variants and maintaining closer compatibility with crypto/ed25519.
//
//	| Scheme Name | Sign Function     | Verification  | Context           |
//	|-------------|-------------------|---------------|-------------------|
//	| Ed25519     | Sign              | Verify        | None              |
//	| Ed25519Ph   | SignPh            | VerifyPh      | Yes, can be empty |
//	| Ed25519Ctx  | SignWithCtx       | VerifyWithCtx | Yes, non-empty    |
//	| All above   | (PrivateKey).Sign | VerifyAny     | As above          |
//
// Specific functions for sign and verify are defined. A generic signing
// function for all schemes is available through the crypto.Signer interface,
// which is implemented by the PrivateKey type. A correspond all-in-one
// verification method is provided by the VerifyAny function.
//
// Signing with Ed25519Ph or Ed25519Ctx requires a context string for domain
// separation. This parameter is passed using a SignerOptions struct defined
// in this package. While Ed25519Ph accepts an empty context, Ed25519Ctx
// enforces non-empty context strings.
//
// # Compatibility with crypto.ed25519
//
// These functions are compatible with the “Ed25519” function defined in
// RFC-8032. However, unlike RFC 8032's formulation, this package's private
// key representation includes a public key suffix to make multiple signing
// operations with the same key more efficient. This package refers to the
// RFC-8032 private key as the “seed”.
//
// References
//
//   - RFC-8032: https://rfc-editor.org/rfc/rfc8032.txt
//   - Ed25519: https://ed25519.cr.yp.to/
//   - EdDSA: High-speed high-security signatures. https://doi.org/10.1007/s13389-012-0027-1
package ed25519

import (
	"crypto"
	"io"

	"github.com/cloudflare/circl/sign"
)

const (
	// ContextMaxSize is the maximum length (in bytes) allowed for context.
	ContextMaxSize = 255
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize = 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize = 64
	// SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.
	SeedSize = 32
)

const (
	paramB = 256 / 8 // Size of keys in bytes.
)

// SignerOptions implements crypto.SignerOpts and augments with parameters
// that are specific to the Ed25519 signature schemes.
type SignerOptions struct {
	// Hash must be crypto.Hash(0) for Ed25519/Ed25519ctx, or crypto.SHA512
	// for Ed25519ph.
	crypto.Hash

	// Context is an optional domain separation string for Ed25519ph and a
	// must for Ed25519ctx. Its length must be less or equal than 255 bytes.
	Context string

	// Scheme is an identifier for choosing a signature scheme. The zero value
	// is ED25519.
	Scheme SchemeID
}

// SchemeID is an identifier for each signature scheme.
type SchemeID uint

const (
	ED25519 SchemeID = iota
	ED25519Ph
	ED25519Ctx
)

// PrivateKey is the type of Ed25519 private keys. It implements crypto.Signer.
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

// Equal reports whether pub and x have the same value.
func (pub PublicKey) Equal(x crypto.PublicKey) bool { _ = "STUB: not implemented"; return false }

// Sign creates a signature of a message with priv key.
// This function is compatible with crypto.ed25519 and also supports the
// three signature variants defined in RFC-8032, namely Ed25519 (or pure
// EdDSA), Ed25519Ph, and Ed25519Ctx.
// The opts.HashFunc() must return zero to specify either Ed25519 or Ed25519Ctx
// variant. This can be achieved by passing crypto.Hash(0) as the value for
// opts.
// The opts.HashFunc() must return SHA512 to specify the Ed25519Ph variant.
// This can be achieved by passing crypto.SHA512 as the value for opts.
// Use a SignerOptions struct (defined in this package) to pass a context
// string for signing.
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

// 1.  Hash the 32-byte private key using SHA-512.

// 2.  Compute SHA-512(dom2(F, C) || prefix || PH(M))

// 3.  Compute the point [r]B.

// 4.  Compute SHA512(dom2(F, C) || R || A || PH(M)).

// 5.  Compute S = (r + k * s) mod order.

// 6.  The signature is the concatenation of R and S.

// Sign signs the message with privateKey and returns a signature.
// This function supports the signature variant defined in RFC-8032: Ed25519,
// also known as the pure version of EdDSA.
// It will panic if len(privateKey) is not PrivateKeySize.
func Sign(privateKey PrivateKey, message []byte) []byte { _ = "STUB: not implemented"; return nil }

// SignPh creates a signature of a message with private key and context.
// This function supports the signature variant defined in RFC-8032: Ed25519ph,
// meaning it internally hashes the message using SHA-512, and optionally
// accepts a context string.
// It will panic if len(privateKey) is not PrivateKeySize.
// Context could be passed to this function, which length should be no more than
// ContextMaxSize=255. It can be empty.
func SignPh(privateKey PrivateKey, message []byte, ctx string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// SignWithCtx creates a signature of a message with private key and context.
// This function supports the signature variant defined in RFC-8032: Ed25519ctx,
// meaning it accepts a non-empty context string.
// It will panic if len(privateKey) is not PrivateKeySize.
// Context must be passed to this function, which length should be no more than
// ContextMaxSize=255 and cannot be empty.
func SignWithCtx(privateKey PrivateKey, message []byte, ctx string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func verify(public PublicKey, message, signature, ctx []byte, preHash bool) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyAny returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded.
// This function supports all the three signature variants defined in RFC-8032,
// namely Ed25519 (or pure EdDSA), Ed25519Ph, and Ed25519Ctx.
// The opts.HashFunc() must return zero to specify either Ed25519 or Ed25519Ctx
// variant. This can be achieved by passing crypto.Hash(0) as the value for opts.
// The opts.HashFunc() must return SHA512 to specify the Ed25519Ph variant.
// This can be achieved by passing crypto.SHA512 as the value for opts.
// Use a SignerOptions struct to pass a context string for signing.
func VerifyAny(public PublicKey, message, signature []byte, opts crypto.SignerOpts) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded.
// This function supports the signature variant defined in RFC-8032: Ed25519,
// also known as the pure version of EdDSA.
func Verify(public PublicKey, message, signature []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyPh returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded.
// This function supports the signature variant defined in RFC-8032: Ed25519ph,
// meaning it internally hashes the message using SHA-512.
// Context could be passed to this function, which length should be no more than
// 255. It can be empty.
func VerifyPh(public PublicKey, message, signature []byte, ctx string) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyWithCtx returns true if the signature is valid. Failure cases are invalid
// signature, or when the public key cannot be decoded, or when context is
// not provided.
// This function supports the signature variant defined in RFC-8032: Ed25519ctx,
// meaning it does not handle prehashed messages. Non-empty context string must be
// provided, and must not be more than 255 of length.
func VerifyWithCtx(public PublicKey, message, signature []byte, ctx string) bool {
	_ = "STUB: not implemented"
	return false
}

func clamp(k []byte) { _ = "STUB: not implemented"; return }

// isLessThanOrder returns true if 0 <= x < order.
func isLessThanOrder(x []byte) bool { _ = "STUB: not implemented"; return false }

func writeDom(h io.Writer, ctx []byte, preHash bool) { _ = "STUB: not implemented"; return }
