// Package sum is a VDAF for aggregating integers in a pre-determined range.
package sum

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

// Sum is a verifiable distributed aggregation function in which each
// measurement is an integer in the range [0, maxMeasurement], where
// maxMeasurement defines the largest valid measurement, the aggregated result
// is the sum of all the measurements.
type Sum struct {
	p prio3.Prio3[uint64, uint64, *flpSum, Vec, Fp, *Fp]
}

func New(numShares uint8, maxMeasurement uint64, context []byte) (s *Sum, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sum) Params() prio3.Params { _ = "STUB: not implemented"; return *new(prio3.Params) }

func (s *Sum) Shard(measurement uint64, nonce *Nonce, rand []byte,
) (PublicShare, []InputShare, error) {
	_ = "STUB: not implemented"
	return *new(PublicShare), nil, nil
}

func (s *Sum) PrepInit(
	verifyKey *VerifyKey,
	nonce *Nonce,
	aggID uint8,
	publicShare PublicShare,
	inputShare InputShare,
) (*PrepState, *PrepShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *Sum) PrepSharesToPrep(prepShares []PrepShare) (*PrepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sum) PrepNext(state *PrepState, msg *PrepMessage) (*OutShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sum) AggregateInit() AggShare { _ = "STUB: not implemented"; return *new(AggShare) }

func (s *Sum) AggregateUpdate(aggShare *AggShare, outShare *OutShare) {
	_ = "STUB: not implemented"
	return
}

func (s *Sum) Unshard(aggShares []AggShare, numMeas uint) (aggregate *uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type flpSum struct {
	flp.FLP[flp.GadgetPolyEvalx2x, poly, Vec, Fp, *Fp]
	bits   uint
	offset Fp
}

func newFlpSum(maxMeasurement uint64) (*flpSum, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *flpSum) Eval(
	out Vec, g flp.Gadget[poly, Vec, Fp, *Fp], numCalls uint,
	meas, jointRand Vec, numShares uint8,
) {
	_ = "STUB: not implemented"
	return
}

func (s *flpSum) Encode(measurement uint64) (Vec, error) {
	_ = "STUB: not implemented"
	return *new(Vec), nil
}

func (s *flpSum) Truncate(meas Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (s *flpSum) Decode(output Vec, numMeas uint) (*uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
