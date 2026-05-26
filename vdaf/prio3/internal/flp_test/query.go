package flp_test

import (
	"testing"

	"github.com/cloudflare/circl/vdaf/prio3/arith"
	"github.com/cloudflare/circl/vdaf/prio3/internal/flp"
)

func TestInvalidQuery[
	G flp.Gadget[P, V, E, F],
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
](t *testing.T, f *flp.FLP[G, P, V, E, F]) {
	_ = "STUB: not implemented"
	return
}

// Check all subgroups of order 2^logN <= 2^logP.

// Check every element in the subgroup of order 2^logN.
