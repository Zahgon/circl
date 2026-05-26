// Package nist implements helpers to generate NIST's Known Answer Tests (KATs).
package nist

// See NIST's PQCgenKAT.c.
type DRBG struct {
	key [32]byte
	v   [16]byte
}

func (g *DRBG) incV() { _ = "STUB: not implemented"; return }

// AES256_CTR_DRBG_Update(pd, &g.key, &g.v).
func (g *DRBG) update(pd *[48]byte) { _ = "STUB: not implemented"; return }

// randombyte_init(seed, NULL, 256).
func NewDRBG(seed *[48]byte) (g DRBG) {
	_ = "STUB: not implemented"
	return *

	// randombytes.
	new(DRBG)
}

func (g *DRBG) Fill(x []byte) { _ = "STUB: not implemented"; return }
