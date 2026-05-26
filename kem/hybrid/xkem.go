package hybrid

import (
	"github.com/cloudflare/circl/dh/x25519"
	"github.com/cloudflare/circl/dh/x448"
	"github.com/cloudflare/circl/kem"
)

type xPublicKey struct {
	scheme *xScheme
	key    []byte
}
type xPrivateKey struct {
	scheme *xScheme
	key    []byte
}
type xScheme struct {
	size int
}

var (
	x25519Kem = &xScheme{x25519.Size}
	x448Kem   = &xScheme{x448.Size}
)

func (sch *xScheme) Name() string { _ = "STUB: not implemented"; return "" }

func (sch *xScheme) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (sch *xScheme) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (sch *xScheme) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (sch *xScheme) SharedKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (sch *xScheme) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (sch *xScheme) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (sk *xPrivateKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (pk *xPublicKey) Scheme() kem.Scheme  { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (sk *xPrivateKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sk *xPrivateKey) Equal(other kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (sk *xPrivateKey) Public() kem.PublicKey {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey)
}

func (pk *xPublicKey) Equal(other kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (pk *xPublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sch *xScheme) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (sch *xScheme) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (sch *xScheme) Encapsulate(pk kem.PublicKey) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (pk *xPublicKey) X(sk *xPrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sch *xScheme) EncapsulateDeterministically(
	pk kem.PublicKey, seed []byte,
) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (sch *xScheme) Decapsulate(sk kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sch *xScheme) UnmarshalBinaryPublicKey(buf []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

func (sch *xScheme) UnmarshalBinaryPrivateKey(buf []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}
