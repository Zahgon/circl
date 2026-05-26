// Package count is a VDAF for counting Boolean measurements.
package count

import (
	"github.com/cloudflare/circl/vdaf/prio3/arith/fp64"
	"github.com/cloudflare/circl/vdaf/prio3/internal/flp"
	"github.com/cloudflare/circl/vdaf/prio3/internal/prio3"
)

type (
	poly        = fp64.Poly
	Vec         = fp64.Vec
	Fp          = fp64.Fp
	AggShare    = prio3.AggShare[Vec, Fp]
	InputShare  = prio3.InputShare[Vec, Fp]
	Nonce       = prio3.Nonce
	OutShare    = prio3.OutShare[Vec, Fp]
	PrepMessage = prio3.PrepMessage
	PrepShare   = prio3.PrepShare[Vec, Fp]
	PrepState   = prio3.PrepState[Vec, Fp]
	PublicShare = prio3.PublicShare
	VerifyKey   = prio3.VerifyKey
)

// Count is a verifiable distributed aggregation function in which each
// measurement is either one or zero and the aggregate result is the sum of
// the measurements.
type Count struct {
	p prio3.Prio3[bool, uint64, *flpCount, Vec, Fp, *Fp]
}

func New(numShares uint8, context []byte) (c *Count, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Count) Params() prio3.Params { _ = "STUB: not implemented"; return *new(prio3.Params) }

func (c *Count) Shard(measurement bool, nonce *Nonce, rand []byte,
) (PublicShare, []InputShare, error) {
	_ = "STUB: not implemented"
	return *new(PublicShare), nil, nil
}

func (c *Count) PrepInit(
	verifyKey *VerifyKey,
	nonce *Nonce,
	aggID uint8,
	publicShare PublicShare,
	inputShare InputShare,
) (*PrepState, *PrepShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Count) PrepSharesToPrep(prepShares []PrepShare) (*PrepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Count) PrepNext(state *PrepState, msg *PrepMessage) (*OutShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Count) AggregateInit() AggShare { _ = "STUB: not implemented"; return *new(AggShare) }

func (c *Count) AggregateUpdate(aggShare *AggShare, outShare *OutShare) {
	_ = "STUB: not implemented"
	return
}

func (c *Count) Unshard(aggShares []AggShare, numMeas uint) (aggregate *uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type flpCount struct {
	flp.FLP[flp.GadgetMulFp64, poly, Vec, Fp, *Fp]
}

func newFlpCount() *flpCount { _ = "STUB: not implemented"; return nil }

func (c *flpCount) Eval(
	out Vec, g flp.Gadget[poly, Vec, Fp, *Fp], numCalls uint,
	meas, jointRand Vec, numShares uint8,
) {
	_ = "STUB: not implemented"
	return
}

func (c *flpCount) Encode(measurement bool) (Vec, error) {
	_ = "STUB: not implemented"
	return *new(Vec), nil
}

func (c *flpCount) Truncate(meas Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (c *flpCount) Decode(output Vec, numMeas uint) (*uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
