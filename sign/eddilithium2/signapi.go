package eddilithium2

import (
	"encoding/asn1"

	"github.com/cloudflare/circl/sign"
)

var sch sign.Scheme = &scheme{}

// Scheme returns a signature interface.
func Scheme() sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

type scheme struct{}

func (*scheme) Name() string          { _ = "STUB: not implemented"; return "" }
func (*scheme) PublicKeySize() int    { _ = "STUB: not implemented"; return 0 }
func (*scheme) PrivateKeySize() int   { _ = "STUB: not implemented"; return 0 }
func (*scheme) SignatureSize() int    { _ = "STUB: not implemented"; return 0 }
func (*scheme) SeedSize() int         { _ = "STUB: not implemented"; return 0 }
func (*scheme) TLSIdentifier() uint   { _ = "STUB: not implemented"; /* temp*/ return 0 }
func (*scheme) SupportsContext() bool { _ = "STUB: not implemented"; return false }
func (*scheme) Oid() asn1.ObjectIdentifier {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier)
}

func (*scheme) GenerateKey() (sign.PublicKey, sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), *new(sign.PrivateKey), nil
}

func (*scheme) Sign(
	sk sign.PrivateKey,
	message []byte,
	opts *sign.SignatureOpts,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (*scheme) Verify(
	pk sign.PublicKey,
	message, signature []byte,
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
