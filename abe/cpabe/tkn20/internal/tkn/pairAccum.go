package tkn

import (
	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

type pairAccum struct {
	as      []*pairing.G1
	bs      []*pairing.G2
	scalars []int
}

func (pairs *pairAccum) addDuals(m1 *matrixG1, m2 *matrixG2, n int) {
	_ = "STUB: not implemented"
	return
}

func (pairs *pairAccum) eval() *pairing.Gt { _ = "STUB: not implemented"; return nil }
