package tkn

import (
	"io"
)

// This file is based on the techniques in
// https://www.iacr.org/archive/pkc2011/65710074/65710074.pdf that
// apply the Boneh-Katz transform to Attribute based encryption.

// Seed size is chosen based on the proof for BK transform
// (https://eprint.iacr.org/2004/261.pdf - page 12, theorem 2) to maintain the
// statistical hiding property. Their input is 448 bits -> 128 bits,
// whereas we require a seed size of 576 bits to ensure a 2^(-65) statistical difference
// for our output size of 256 bits.
const macKeySeedSize = 72

// As of v1.3.8, ciphertexts are prefixed with this string.
const CiphertextVersion = "v1.3.8"

func blakeEncrypt(key []byte, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blakeDecrypt(key []byte, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blakeMac(key []byte, msg []byte) (tag []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandSeed(seed []byte) (id []byte, macKey []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func DeriveAttributeKeysCCA(rand io.Reader, sp *SecretParams, attrs *Attributes) (*AttributesKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncryptCCA(rand io.Reader, public *PublicParams, policy *Policy, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send the policy that was not enhanced. The receiver will recover with the ID.
// This avoids a bug where we omit the check that the ID is correct

type rmLenPref = func([]byte) ([]byte, []byte, error)

func checkCiphertextFormat(ciphertext []byte) (ct []byte, fn rmLenPref) {
	_ = "STUB: not implemented"
	return nil, *new(rmLenPref)
}

func DecryptCCA(ciphertext []byte, key *AttributesKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt the envelope

// Now check that compTag = tag and compID = id
// We don't want to distinguish which fails.

func CouldDecrypt(ciphertext []byte, a *Attributes) bool { _ = "STUB: not implemented"; return false }

func (p *Policy) ExtractFromCiphertext(ct []byte) error { _ = "STUB: not implemented"; return nil }
