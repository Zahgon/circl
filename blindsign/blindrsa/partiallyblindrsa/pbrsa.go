// Package partiallyblindrsa implements a partially blind RSA protocol.
package partiallyblindrsa

import (
	"crypto"
	"crypto/rsa"
	"errors"
	"hash"
	"io"
	"math/big"

	"github.com/cloudflare/circl/blindsign/blindrsa/internal/common"
	"github.com/cloudflare/circl/blindsign/blindrsa/internal/keys"
)

func encodeMessageMetadata(message, metadata []byte) []byte { _ = "STUB: not implemented"; return nil }

// A randomizedVerifier represents a Verifier in the partially blind RSA signature protocol.
// It carries state needed to produce and validate an RSA signature produced
// using the blind RSA protocol.
type randomizedVerifier struct {
	// Public key of the Signer
	pk *keys.BigPublicKey

	// Identifier of the cryptographic hash function used in producing the message signature
	cryptoHash crypto.Hash

	// Hash function used in producing the message signature
	hash hash.Hash
}

// NewVerifier creates a new PBRSAVerifier using the corresponding Signer parameters.
// This corresponds to the RSAPBSSA-SHA384-PSS-Deterministic variant. See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa#name-rsapbssa-variants
func NewVerifier(pk *rsa.PublicKey, hash crypto.Hash) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

// derivePublicKey tweaks the public key based on the input metadata.
//
// See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00#name-public-key-augmentation
//
// See the following issue for more discussion on HKDF vs hash-to-field:
// https://github.com/cfrg/draft-irtf-cfrg-hash-to-curve/issues/202
func derivePublicKey(h crypto.Hash, pk *keys.BigPublicKey, metadata []byte) *keys.BigPublicKey {
	_ = "STUB: not implemented"
	// expandLen = ceil((ceil(log2(\lambda)/2) + k) / 8), where k is the security parameter of the suite (e.g., k = 128).
	// We stretch the input metadata beyond \lambda bits s.t. the output bytes are indifferentiable from truly random bytes
	return nil
}

// H_MD(D) = 1 || G(x), where G(x) is output of length \lambda-2 bits
// We do this by sampling \lambda bits, clearing the top two bits (so the output is \lambda-2 bits)
// and setting the bottom bit (so the result is odd).

// Compute e_MD = e * H_MD(D)

// deriveKeyPair tweaks the private key using the metadata as input.
//
// See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00#name-private-key-augmentation
func deriveKeyPair(h crypto.Hash, sk *keys.BigPrivateKey, metadata []byte) *keys.BigPrivateKey {
	_ = "STUB: not implemented"
	// pih(N) = (p-1)(q-1)
	return nil
}

// d = e^-1 mod phi(N)

func fixedPartiallyBlind(message, salt []byte, r, rInv *big.Int, pk *keys.BigPublicKey, hash hash.Hash) ([]byte, VerifierState, error) {
	_ = "STUB: not implemented"
	return nil, *new(VerifierState), nil
}

// Verifier is a type that implements the client side of the partially blind RSA
// protocol, described in https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00
type Verifier interface {
	// Blind initializes the partially blind RSA protocol using an input message and source of
	// randomness. The signature includes a randomly generated PSS salt whose length equals the
	// size of the underlying hash function. This function fails if randomness was not provided.
	Blind(random io.Reader, message, metadata []byte) ([]byte, VerifierState, error)

	// FixedBlind initializes the partially blind RSA protocol using an input message, metadata, and randomness values.
	FixedBlind(message, metadata, salt, blind, blindInv []byte) ([]byte, VerifierState, error)

	// Verify verifies the input (message, signature) pair using the augmented public key
	// and produces an error upon failure.
	Verify(message, signature, metadata []byte) error

	// Hash returns the hash function associated with the Verifier.
	Hash() hash.Hash
}

// Blind initializes the partially blind RSA protocol using an input message and source of randomness. The
// signature includes a randomly generated PSS salt whose length equals the size of the underlying
// hash function. This function fails if randomness was not provided.
//
// See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00#name-blind
func (v randomizedVerifier) Blind(random io.Reader, message, metadata []byte) ([]byte, VerifierState, error) {
	_ = "STUB: not implemented"
	return nil, *new(VerifierState), nil
}

// FixedBlind initializes the partially blind RSA using fixed randomness as input.
func (v randomizedVerifier) FixedBlind(message, metadata, salt, blind, blindInv []byte) ([]byte, VerifierState, error) {
	_ = "STUB: not implemented"
	return nil, *new(VerifierState), nil
}

// Verify verifies the input (message, signature) pair using the augmented public key
// and produces an error upon failure.
//
// See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00#name-verification-2
func (v randomizedVerifier) Verify(message, metadata, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash returns the hash function associated with the Verifier.
func (v randomizedVerifier) Hash() hash.Hash {
	_ = "STUB: not implemented"

	// A VerifierState carries state needed to complete the blind signature protocol
	// as a verifier.
	return *new(hash.Hash)
}

type VerifierState struct {
	// Public key of the Signer
	pk *keys.BigPublicKey

	// Hash function used in producing the message signature
	hash hash.Hash

	// The hashed and encoded message being signed
	encodedMsg []byte

	// The salt used when encoding the message
	salt []byte

	// Inverse of the blinding factor produced by the Verifier
	rInv *big.Int
}

// Finalize computes and outputs the final signature, if it's valid. Otherwise, it returns an error.
//
// See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00#name-finalize
func (state VerifierState) Finalize(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CopyBlind returns an encoding of the blind value used in the protocol.
func (state VerifierState) CopyBlind() []byte { _ = "STUB: not implemented"; return nil }

// CopySalt returns an encoding of the per-message salt used in the protocol.
func (state VerifierState) CopySalt() []byte { _ = "STUB: not implemented"; return nil }

// An Signer represents the Signer in the blind RSA protocol.
// It carries the raw RSA private key used for signing blinded messages.
type Signer struct {
	// An RSA private key
	sk *keys.BigPrivateKey
	h  crypto.Hash
}

// isSafePrime returns true if the input prime p is safe, i.e., p = (2 * q) + 1 for some prime q
func isSafePrime(p *big.Int) bool { _ = "STUB: not implemented"; return false }

// NewSigner creates a new Signer for the blind RSA protocol using an RSA private key.
func NewSigner(sk *rsa.PrivateKey, h crypto.Hash) (Signer, error) {
	_ = "STUB: not implemented"
	return *new(Signer), nil
}

// BlindSign blindly computes the RSA operation using the Signer's private key on the blinded
// message input, if it's of valid length, and returns an error should the function fail.
//
// See the specification for more details:
// https://datatracker.ietf.org/doc/html/draft-amjad-cfrg-partially-blind-rsa-00#name-blindsign
func (signer Signer) BlindSign(data, metadata []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	// ErrInvalidPrivateKey is the error used if a private key is invalid
	ErrInvalidPrivateKey    = errors.New("blindsign/blindrsa/partiallyblindrsa: invalid private key")
	ErrUnexpectedSize       = common.ErrUnexpectedSize
	ErrInvalidMessageLength = common.ErrInvalidMessageLength
	ErrInvalidRandomness    = common.ErrInvalidRandomness
)
