//go:build purego || (!amd64 && !arm64)
// +build purego !amd64,!arm64

package p384

import (
	"crypto/elliptic"
	"math/big"
)

type curve struct{ elliptic.Curve }

func P384() Curve { _ = "STUB: not implemented"; return *new(Curve) }

// CombinedMult calculates P=mG+nQ, where G is the generator and Q=(x,y,z).
// The scalars m and n are integers in big-endian form. Non-constant time.
func (c curve) CombinedMult(xQ, yQ *big.Int, m, n []byte) (xP, yP *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}
