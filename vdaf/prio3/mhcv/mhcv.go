// Package mhcv is a VDAF for aggregating vectors of Booleans with bounded weight.
package mhcv

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

// MultiHotCountVec is a verifiable distributed aggregation function in which
// each measurement is a vector of Booleans, where the number of True
// values is bounded.
// This provides a functionality similar to Histogram except that more than
// one entry (or none at all) may be non-zero.
type MultiHotCountVec struct {
	p prio3.Prio3[[]bool, []uint64, *flpMultiHotCountVec, Vec, Fp, *Fp]
}

func New(numShares uint8, length, maxWeight, chunkLength uint, context []byte) (m *MultiHotCountVec, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MultiHotCountVec) Params() prio3.Params {
	_ = "STUB: not implemented"
	return *new(prio3.Params)
}

func (m *MultiHotCountVec) Shard(measurement []bool, nonce *Nonce, rand []byte,
) (PublicShare, []InputShare, error) {
	_ = "STUB: not implemented"
	return *new(PublicShare), nil, nil
}

func (m *MultiHotCountVec) PrepInit(
	verifyKey *VerifyKey,
	nonce *Nonce,
	aggID uint8,
	publicShare PublicShare,
	inputShare InputShare,
) (*PrepState, *PrepShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *MultiHotCountVec) PrepSharesToPrep(prepShares []PrepShare) (*PrepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MultiHotCountVec) PrepNext(state *PrepState, msg *PrepMessage) (*OutShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MultiHotCountVec) AggregateInit() AggShare {
	_ = "STUB: not implemented"
	return *new(AggShare)
}

func (m *MultiHotCountVec) AggregateUpdate(aggShare *AggShare, outShare *OutShare) {
	_ = "STUB: not implemented"
	return
}

func (m *MultiHotCountVec) Unshard(aggShares []AggShare, numMeas uint) (aggregate *[]uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type flpMultiHotCountVec struct {
	flp.FLP[flp.GadgetParallelSumInnerMul, poly, Vec, Fp, *Fp]
	bits        uint
	chunkLength uint
	length      uint
	offset      Fp
}

func newFlpMultiCountHotVec(length, maxWeight, chunkLength uint) (*flpMultiHotCountVec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *flpMultiHotCountVec) Eval(
	out Vec, g flp.Gadget[poly, Vec, Fp, *Fp], numCalls uint,
	meas, jointRand Vec, numShares uint8,
) {
	_ = "STUB: not implemented"
	return
}

func (m *flpMultiHotCountVec) Encode(measurement []bool) (out Vec, err error) {
	_ = "STUB: not implemented"
	return *new(Vec), nil
}

func (m *flpMultiHotCountVec) Truncate(meas Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (m *flpMultiHotCountVec) Decode(output Vec, numMeas uint) (*[]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	ErrLength      = errors.New("length cannot be zero")
	ErrMaxWeight   = errors.New("maxWeight cannot be greater than length")
	ErrChunkLength = errors.New("chunkLength cannot be zero")
	ErrFieldSize   = errors.New("length and maxWeight are too large for the current field size")
)
