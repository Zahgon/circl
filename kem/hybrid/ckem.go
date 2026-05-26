package hybrid

import (
	"crypto/ecdh"

	"github.com/cloudflare/circl/kem"
)

type cPublicKey struct {
	scheme cScheme
	key    *ecdh.PublicKey
}
type cPrivateKey struct {
	scheme cScheme
	key    *ecdh.PrivateKey
}
type cScheme struct {
	curve ecdh.Curve
}

var p256Kem = &cScheme{ecdh.P256()}

func (sch cScheme) Name() string { _ = "STUB: not implemented"; return "" }

func (sch cScheme) PublicKeySize() int { _ = "STUB: not implemented"; return 0 }

func (sch cScheme) PrivateKeySize() int { _ = "STUB: not implemented"; return 0 }

func (sch cScheme) SeedSize() int { _ = "STUB: not implemented"; return 0 }

func (sch cScheme) SharedKeySize() int { _ = "STUB: not implemented"; return 0 }

func (sch cScheme) CiphertextSize() int { _ = "STUB: not implemented"; return 0 }

func (sch cScheme) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (sk *cPrivateKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (pk *cPublicKey) Scheme() kem.Scheme  { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (sk *cPrivateKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sk *cPrivateKey) Equal(other kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (sk *cPrivateKey) Public() kem.PublicKey {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey)
}

func (pk *cPublicKey) Equal(other kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (pk *cPublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sch cScheme) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (sch cScheme) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (sch cScheme) Encapsulate(pk kem.PublicKey) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (pk *cPublicKey) X(sk *cPrivateKey) []byte { _ = "STUB: not implemented"; return nil }

// ECDH cannot fail for NIST curves as NewPublicKey rejects
// invalid points and the point in infinity, and NewPrivateKey
// rejects invalid scalars and the zero value.

func (sch cScheme) EncapsulateDeterministically(
	pk kem.PublicKey, seed []byte,
) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (sch cScheme) Decapsulate(sk kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sch cScheme) UnmarshalBinaryPublicKey(buf []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

func (sch cScheme) UnmarshalBinaryPrivateKey(buf []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}
