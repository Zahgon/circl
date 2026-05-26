// Package histogram is a VDAF for aggregating integer measurements into buckets.
package histogram

import (
	"github.com/cloudflare/circl/vdaf/prio3/arith/fp128"
	"github.com/cloudflare/circl/vdaf/prio3/internal/flp"
	"github.com/cloudflare/circl/vdaf/prio3/internal/prio3"
)

type (
	poly        = fp128.Poly
	Vec         = fp128.Vec
	Fp          = fp128.Fp
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

// Histogram is a verifiable distributed aggregation function in which each
// measurement increments by one the histogram bucket, out of a set of fixed
// buckets, and the aggregate result counts the number of measurements in each
// bucket.
type Histogram struct {
	p prio3.Prio3[uint64, []uint64, *flpHistogram, Vec, Fp, *Fp]
}

func New(numShares uint8, length, chunkLen uint, context []byte) (h *Histogram, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Histogram) Params() prio3.Params { _ = "STUB: not implemented"; return *new(prio3.Params) }

func (h *Histogram) Shard(measurement uint64, nonce *Nonce, rand []byte,
) (PublicShare, []InputShare, error) {
	_ = "STUB: not implemented"
	return *new(PublicShare), nil, nil
}

func (h *Histogram) PrepInit(
	verifyKey *VerifyKey,
	nonce *Nonce,
	aggID uint8,
	publicShare PublicShare,
	inputShare InputShare,
) (*PrepState, *PrepShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (h *Histogram) PrepSharesToPrep(prepShares []PrepShare) (*PrepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Histogram) PrepNext(state *PrepState, msg *PrepMessage) (*OutShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Histogram) AggregateInit() AggShare { _ = "STUB: not implemented"; return *new(AggShare) }

func (h *Histogram) AggregateUpdate(aggShare *AggShare, outShare *OutShare) {
	_ = "STUB: not implemented"
	return
}

func (h *Histogram) Unshard(aggShares []AggShare, numMeas uint) (aggregate *[]uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type flpHistogram struct {
	flp.FLP[flp.GadgetParallelSumInnerMul, poly, Vec, Fp, *Fp]
	length   uint
	chunkLen uint
}

func newFlpHistogram(length, chunkLen uint) *flpHistogram { _ = "STUB: not implemented"; return nil }

func (h *flpHistogram) Eval(
	out Vec, g flp.Gadget[poly, Vec, Fp, *Fp], numCalls uint,
	meas, jointRand Vec, numShares uint8,
) {
	_ = "STUB: not implemented"
	return
}

func (h *flpHistogram) Encode(measurement uint64) (out Vec, err error) {
	_ = "STUB: not implemented"
	return *new(Vec), nil
}

func (h *flpHistogram) Truncate(meas Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (h *flpHistogram) Decode(output Vec, numMeas uint) (*[]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
