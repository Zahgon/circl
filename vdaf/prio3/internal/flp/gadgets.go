package flp

import (
	"github.com/cloudflare/circl/vdaf/prio3/arith"
	"github.com/cloudflare/circl/vdaf/prio3/arith/fp128"
	"github.com/cloudflare/circl/vdaf/prio3/arith/fp64"
)

type Gadget[
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] interface {
	// Arity is the number of input wires.
	Arity() uint
	// Degree is the arithmetic degree of the gadget circuit.
	Degree() uint
	// Eval is evaluates the gadget over the given inputs.
	Eval(out *E, in V)
	// EvalPoly is evaluates the circuit over the polynomial ring of the field.
	EvalPoly(out P, in []P)
}

type gadgetMul struct{}

func (gadgetMul) Arity() uint  { _ = "STUB: not implemented"; return 0 }
func (gadgetMul) Degree() uint { _ = "STUB: not implemented"; return 0 }

type GadgetMulFp64 struct{ gadgetMul }

func (GadgetMulFp64) Eval(out *fp64.Fp, in fp64.Vec) { _ = "STUB: not implemented"; return }

func (GadgetMulFp64) EvalPoly(out fp64.Poly, in []fp64.Poly) { _ = "STUB: not implemented"; return }

type gadgetMulFp128 struct{ gadgetMul }

func (gadgetMulFp128) Eval(out *fp128.Fp, in fp128.Vec) { _ = "STUB: not implemented"; return }

func (gadgetMulFp128) EvalPoly(out fp128.Poly, in []fp128.Poly) { _ = "STUB: not implemented"; return }

// PolyEval gadget for p(x) = x^2-x.
type GadgetPolyEvalx2x struct{}

func (GadgetPolyEvalx2x) Arity() uint                    { _ = "STUB: not implemented"; return 0 }
func (GadgetPolyEvalx2x) Degree() uint                   { _ = "STUB: not implemented"; return 0 }
func (GadgetPolyEvalx2x) Eval(out *fp64.Fp, in fp64.Vec) { _ = "STUB: not implemented"; return }

func (GadgetPolyEvalx2x) EvalPoly(out fp64.Poly, in []fp64.Poly) { _ = "STUB: not implemented"; return }

type GadgetParallelSumInnerMul struct {
	inner gadgetMulFp128
	Count uint
}

func (g GadgetParallelSumInnerMul) Arity() uint  { _ = "STUB: not implemented"; return 0 }
func (g GadgetParallelSumInnerMul) Degree() uint { _ = "STUB: not implemented"; return 0 }
func (g GadgetParallelSumInnerMul) Eval(out *fp128.Fp, in fp128.Vec) {
	_ = "STUB: not implemented"
	return
}

func (g GadgetParallelSumInnerMul) EvalPoly(out fp128.Poly, in []fp128.Poly) {
	_ = "STUB: not implemented"
	return
}

type wrapperGadget[
	G Gadget[P, V, E, F],
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] struct {
	inner    G
	wires    V
	p, log2p uint
	k        uint
}

func (g *wrapperGadget[G, P, V, E, F]) Arity() uint            { _ = "STUB: not implemented"; return 0 }
func (g *wrapperGadget[G, P, V, E, F]) Degree() uint           { _ = "STUB: not implemented"; return 0 }
func (g *wrapperGadget[G, P, V, E, F]) EvalPoly(out P, in []P) { _ = "STUB: not implemented"; return }
func (g *wrapperGadget[G, P, V, E, F]) eval(input V)           { _ = "STUB: not implemented"; return }

type ProveGadget[
	G Gadget[P, V, E, F],
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] struct {
	wrapperGadget[G, P, V, E, F]
}

func (g *ProveGadget[G, P, V, E, F]) Eval(out *E, input V) { _ = "STUB: not implemented"; return }

type QueryGadget[
	G Gadget[P, V, E, F],
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] struct {
	poly          P
	alpha, alphaK E
	wrapperGadget[G, P, V, E, F]
}

func (g *QueryGadget[G, P, V, E, F]) Eval(out *E, input V) { _ = "STUB: not implemented"; return }
