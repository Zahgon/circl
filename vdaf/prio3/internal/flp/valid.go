package flp

import (
	"github.com/cloudflare/circl/vdaf/prio3/arith"
	"github.com/cloudflare/circl/vdaf/prio3/arith/fp128"
)

type Params struct {
	MeasurementLen uint
	JointRandLen   uint
	EvalOutputLen  uint
	OutputLen      uint
}

func (p Params) MeasurementLength() uint { _ = "STUB: not implemented"; return 0 }
func (p Params) JointRandLength() uint   { _ = "STUB: not implemented"; return 0 }
func (p Params) OutputLength() uint      { _ = "STUB: not implemented"; return 0 }
func (p Params) EvalOutputLength() uint  { _ = "STUB: not implemented"; return 0 }
func (p Params) QueryRandLength() uint   { _ = "STUB: not implemented"; return 0 }

type Valid[
	G Gadget[P, V, E, F],
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] struct {
	Gadget         G    // Gadget to be evaluated
	NumGadgetCalls uint // Number of times each gadget is called.
	Params
}

func (v *Valid[G, P, V, E, F]) ProveRandLength() uint { _ = "STUB: not implemented"; return 0 }
func (v *Valid[G, P, V, E, F]) ProofLength() uint     { _ = "STUB: not implemented"; return 0 }
func (v *Valid[G, P, V, E, F]) VerifierLength() uint  { _ = "STUB: not implemented"; return 0 }
func (v *Valid[G, P, V, E, F]) gadgetPolyLen() uint   { _ = "STUB: not implemented"; return 0 }

func (v *Valid[G, P, V, E, F]) wrapProve(proveRand V) *ProveGadget[G, P, V, E, F] {
	_ = "STUB: not implemented"
	return nil
}

func (v *Valid[G, P, V, E, F]) wrapQuery(proof V) *QueryGadget[G, P, V, E, F] {
	_ = "STUB: not implemented"
	return nil
}

func (v *Valid[G, P, V, E, F]) wrap(wireSeeds V) (g wrapperGadget[G, P, V, E, F]) {
	_ = "STUB: not implemented"
	return nil
}

func RangeCheck(
	g Gadget[fp128.Poly, fp128.Vec, fp128.Fp, *fp128.Fp],
	numCalls uint,
	chunkLen uint, sharesInv *fp128.Fp,
	meas, jointRand fp128.Vec,
) (out fp128.Fp) {
	_ = "STUB: not implemented"
	return *new(fp128.Fp)
}
