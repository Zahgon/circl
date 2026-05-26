package slhdsa

import (
	"crypto"

	"golang.org/x/crypto/cryptobyte"
)

// [PrivateKey] stores a private key of the SLH-DSA scheme.
// It implements the [crypto.Signer] and [crypto.PrivateKey] interfaces.
// For serialization, it also implements [cryptobyte.MarshalingValue],
// [encoding.BinaryMarshaler], and [encoding.BinaryUnmarshaler].
type PrivateKey struct {
	seed, prfKey []byte
	publicKey    PublicKey
	ID
}

func (p *params) PrivateKeySize() int { _ = "STUB: not implemented"; return 0 }

// Marshal serializes the key using a [cryptobyte.Builder].
func (k PrivateKey) Marshal(b *cryptobyte.Builder) error { _ = "STUB: not implemented"; return nil }

// Unmarshal recovers a [PrivateKey] from a [cryptobyte.String].
// Caller must specify the private key's [ID] in advance.
// Example:
//
//	key := PrivateKey{ID: SHA2Small192}
//	key.Unmarshal(str) // returns true
func (k *PrivateKey) Unmarshal(s *cryptobyte.String) bool { _ = "STUB: not implemented"; return false }

func (k *PrivateKey) fromBytes(p *params, c *cursor) bool { _ = "STUB: not implemented"; return false }

// UnmarshalBinary recovers a [PrivateKey] from a slice of bytes.
// Caller must specify the private key's [ID] in advance.
// Example:
//
//	key := PrivateKey{ID: SHA2Small192}
//	key.UnmarshalBinary(bytes) // returns nil
func (k *PrivateKey) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }
func (k PrivateKey) MarshalBinary() ([]byte, error)  { _ = "STUB: not implemented"; return nil, nil }
func (k PrivateKey) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}
func (k PrivateKey) PublicKey() (pub PublicKey) { _ = "STUB: not implemented"; return *new(PublicKey) }

func (k PrivateKey) Equal(x crypto.PrivateKey) bool { _ = "STUB: not implemented"; return false }

// [PublicKey] stores a public key of the SLH-DSA scheme.
// It implements the [crypto.PublicKey] interface.
// For serialization, it also implements [cryptobyte.MarshalingValue],
// [encoding.BinaryMarshaler], and [encoding.BinaryUnmarshaler].
type PublicKey struct {
	seed, root []byte
	ID
}

func (p *params) PublicKeySize() int {
	_ = "STUB: not implemented"

	// Marshal serializes the key using a [cryptobyte.Builder].
	return 0
}

func (k PublicKey) Marshal(b *cryptobyte.Builder) error { _ = "STUB: not implemented"; return nil }

// Unmarshal recovers a [PublicKey] from a [cryptobyte.String].
// Caller must specify the public key's [ID] in advance.
// Example:
//
//	key := PublicKey{ID: SHA2Small192}
//	key.Unmarshal(str) // returns true
func (k *PublicKey) Unmarshal(s *cryptobyte.String) bool { _ = "STUB: not implemented"; return false }

func (k *PublicKey) fromBytes(p *params, c *cursor) bool { _ = "STUB: not implemented"; return false }

// UnmarshalBinary recovers a [PublicKey] from a slice of bytes.
// Caller must specify the public key's [ID] in advance.
// Example:
//
//	key := PublicKey{ID: SHA2Small192}
//	key.UnmarshalBinary(bytes) // returns nil
func (k *PublicKey) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }
func (k PublicKey) MarshalBinary() ([]byte, error)  { _ = "STUB: not implemented"; return nil, nil }
func (k PublicKey) Equal(x crypto.PublicKey) bool   { _ = "STUB: not implemented"; return false }
