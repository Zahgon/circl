package goldilocks

import (
	mlsb "github.com/cloudflare/circl/math/mlsbset"
)

const (
	// MLSBRecoding parameters
	fxT   = 448
	fxV   = 2
	fxW   = 3
	fx2w1 = 1 << (uint(fxW) - 1)
)

// ScalarBaseMult returns kG where G is the generator point.
func (e twistCurve) ScalarBaseMult(k *Scalar) *twistPoint { _ = "STUB: not implemented"; return nil }

type groupMLSB struct{}

func (e groupMLSB) ExtendedEltP() mlsb.EltP                { _ = "STUB: not implemented"; return *new(mlsb.EltP) }
func (e groupMLSB) Sqr(x mlsb.EltG)                        { _ = "STUB: not implemented"; return }
func (e groupMLSB) Mul(x mlsb.EltG, y mlsb.EltP)           { _ = "STUB: not implemented"; return }
func (e groupMLSB) Identity() mlsb.EltG                    { _ = "STUB: not implemented"; return *new(mlsb.EltG) }
func (e groupMLSB) NewEltP() mlsb.EltP                     { _ = "STUB: not implemented"; return *new(mlsb.EltP) }
func (e groupMLSB) Lookup(a mlsb.EltP, v uint, s, u int32) { _ = "STUB: not implemented"; return }
