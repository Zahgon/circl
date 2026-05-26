package slhdsa

import (
	"github.com/cloudflare/circl/sign"
)

func (id ID) Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

type scheme struct{ *params }

func (s scheme) Name() string  { _ = "STUB: not implemented"; return "" }
func (s scheme) SeedSize() int { _ = "STUB: not implemented"; return 0 }
func (s scheme) SupportsContext() bool {
	_ = "STUB: not implemented"

	// GenerateKey is similar to [GenerateKey] function, except it always reads
	// random bytes from [rand.Reader].
	return false
}

func (s scheme) GenerateKey() (sign.PublicKey, sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), *new(sign.PrivateKey), nil
}

// Sign returns a randomized pure signature of the message with the context
// given.
// If options is nil, an empty context is used.
// It returns an empty slice if the signature generation fails.
//
// Panics if the key is not a [PrivateKey] or when the [ID] mismatches.
func (s scheme) Sign(
	priv sign.PrivateKey, message []byte, options *sign.SignatureOpts,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Verify returns true if the signature of the message with the specified
// context is valid.
// If options is nil, an empty context is used.
//
// Panics if the key is not a [PublicKey] or when the [ID] mismatches.
func (s scheme) Verify(
	pub sign.PublicKey, message, signature []byte, options *sign.SignatureOpts,
) bool {
	_ = "STUB: not implemented"
	return false
}

// DeriveKey deterministically generates a pair of keys from a seed.
//
// Panics if seed is not of length [sign.Scheme.SeedSize].
func (s scheme) DeriveKey(seed []byte) (sign.PublicKey, sign.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), *new(sign.PrivateKey)
}

func (s scheme) UnmarshalBinaryPublicKey(b []byte) (sign.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), nil
}

func (s scheme) UnmarshalBinaryPrivateKey(b []byte) (sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PrivateKey), nil
}
