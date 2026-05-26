// Code generated from pkg.templ.go. DO NOT EDIT.

// Package sikep751 is deprecated, it implements the key encapsulation mechanism SIKEp751.
//
// # DEPRECATION NOTICE
//
// SIDH and SIKE are deprecated as were shown vulnerable to a key recovery
// attack by Castryck-Decru's paper (https://eprint.iacr.org/2022/975). New
// systems should not rely on this package. This package is frozen.
package sikep751

import (
	"io"

	"github.com/cloudflare/circl/dh/sidh"
	"github.com/cloudflare/circl/kem"
)

// Deprecated: not cryptographically secure.
type PrivateKey struct {
	sk *sidh.PrivateKey
	pk *sidh.PublicKey
}

// Deprecated: not cryptographically secure.
type PublicKey sidh.PublicKey

const (
	SeedSize              = 32
	EncapsulationSeedSize = 32
)

type scheme struct{}

var sch kem.Scheme = &scheme{}

// Scheme returns a KEM interface.
//
// Deprecated: not cryptographically secure.
func Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }

var params *sidh.KEM

func (*scheme) Name() string               { _ = "STUB: not implemented"; return "" }
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

func (sk *PrivateKey) Public() kem.PublicKey { _ = "STUB: not implemented"; return *new(kem.PublicKey) }

func (pk *PublicKey) Equal(other kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (pk *PublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Deprecated: not cryptographically secure.
func GenerateKeyPair(rand io.Reader) (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (*scheme) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (*scheme) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (sch *scheme) Encapsulate(pk kem.PublicKey) (ct []byte, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (sch *scheme) EncapsulateDeterministically(pk kem.PublicKey, seed []byte) (ct []byte, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (sch *scheme) Decapsulate(sk kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sch *scheme) UnmarshalBinaryPublicKey(buf []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

func (sch *scheme) UnmarshalBinaryPrivateKey(buf []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}

func init() {
	params = sidh.NewSike751(nil)
}
