package hpke

import (
	"crypto"

	"github.com/cloudflare/circl/kem"
)

type dhKEM interface {
	sizeDH() int
	calcDH(dh []byte, sk kem.PrivateKey, pk kem.PublicKey) error
	SeedSize() int
	DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey)
	UnmarshalBinaryPrivateKey(data []byte) (kem.PrivateKey, error)
	UnmarshalBinaryPublicKey(data []byte) (kem.PublicKey, error)
}

type kemBase struct {
	id   KEM
	name string
	crypto.Hash
}

type dhKemBase struct {
	kemBase
	dhKEM
}

func (k kemBase) Name() string       { _ = "STUB: not implemented"; return "" }
func (k kemBase) SharedKeySize() int { _ = "STUB: not implemented"; return 0 }

func (k kemBase) getSuiteID() (sid [5]byte) { _ = "STUB: not implemented"; return nil }

func (k kemBase) extractExpand(dh, kemCtx []byte) []byte { _ = "STUB: not implemented"; return nil }

func (k kemBase) labeledExtract(salt, label, info []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (k kemBase) labeledExpand(prk, label, info []byte, l uint16) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (k dhKemBase) AuthEncapsulate(pkr kem.PublicKey, sks kem.PrivateKey) (
	ct []byte, ss []byte, err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) Encapsulate(pkr kem.PublicKey) (
	ct []byte, ss []byte, err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) AuthEncapsulateDeterministically(
	pkr kem.PublicKey, sks kem.PrivateKey, seed []byte,
) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) EncapsulateDeterministically(
	pkr kem.PublicKey, seed []byte,
) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) encap(
	pkR kem.PublicKey,
	seed []byte,
) (ct []byte, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) authEncap(
	pkR kem.PublicKey,
	skS kem.PrivateKey,
	seed []byte,
) (ct []byte, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) coreEncap(
	dh []byte,
	pkR kem.PublicKey,
	seed []byte,
) (enc []byte, kemCtx []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (k dhKemBase) Decapsulate(skr kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k dhKemBase) AuthDecapsulate(
	skR kem.PrivateKey,
	ct []byte,
	pkS kem.PublicKey,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k dhKemBase) coreDecap(
	dh []byte,
	skR kem.PrivateKey,
	ct []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
