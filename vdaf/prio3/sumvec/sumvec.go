// Package sumvec is a VDAF for aggregating vectors of integers in a pre-determined range.
package sumvec

import (
	"errors"

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

// SumVec is a verifiable distributed aggregation function in which each
// measurement is a fixed-length vector of integers in the range [0, 2^bits).
// the aggregated result is the sum of all the vectors.
type SumVec struct {
	p prio3.Prio3[[]uint64, []uint64, *flpSumVec, Vec, Fp, *Fp]
}

func New(numShares uint8, length, bits, chunkLength uint, context []byte) (s *SumVec, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SumVec) Params() prio3.Params { _ = "STUB: not implemented"; return *new(prio3.Params) }

func (s *SumVec) Shard(measurement []uint64, nonce *Nonce, rand []byte,
) (PublicShare, []InputShare, error) {
	_ = "STUB: not implemented"
	return *new(PublicShare), nil, nil
}

func (s *SumVec) PrepInit(
	verifyKey *VerifyKey,
	nonce *Nonce,
	aggID uint8,
	publicShare PublicShare,
	inputShare InputShare,
) (*PrepState, *PrepShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *SumVec) PrepSharesToPrep(prepShares []PrepShare) (*PrepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SumVec) PrepNext(state *PrepState, msg *PrepMessage) (*OutShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SumVec) AggregateInit() AggShare { _ = "STUB: not implemented"; return *new(AggShare) }

func (s *SumVec) AggregateUpdate(aggShare *AggShare, outShare *OutShare) {
	_ = "STUB: not implemented"
	return
}

func (s *SumVec) Unshard(aggShares []AggShare, numMeas uint) (aggregate *[]uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type flpSumVec struct {
	flp.FLP[flp.GadgetParallelSumInnerMul, poly, Vec, Fp, *Fp]
	length   uint
	bits     uint
	chunkLen uint
}

func newFlpSumVec(length, bits, chunkLen uint) (*flpSumVec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *flpSumVec) Eval(
	out Vec, g flp.Gadget[poly, Vec, Fp, *Fp], numCalls uint,
	meas, jointRand Vec, numShares uint8,
) {
	_ = "STUB: not implemented"
	return
}

func (s *flpSumVec) Encode(measurement []uint64) (out Vec, err error) {
	_ = "STUB: not implemented"
	return *new(Vec), nil
}

func (s *flpSumVec) Truncate(meas Vec) (out Vec) { _ = "STUB: not implemented"; return *new(Vec) }

func (s *flpSumVec) Decode(output Vec, numMeas uint) (*[]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrBits = errors.New("bits larger than 64 is not supported")
