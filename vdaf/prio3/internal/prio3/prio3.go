// Package prio3 supports a variety of verifiable distributed aggregation functions.
//
// Clients protect the privacy of their measurements by secret sharing them
// and distributing the shares among the Aggregators.
// To ensure each measurement is valid, the Aggregators run a multi-party
// computation on their shares, the result of which is the output of the
// arithmetic circuit.
// This involves verification of a Fully Linear Proof (FLP) that specifies
// the types of measurements.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7
package prio3

import (
	"errors"

	"github.com/cloudflare/circl/vdaf/prio3/arith"
)

// Prio3 supports a variety of verifiable distributed aggregation functions.
// An instance is parametrized by the type of measurement and aggregated data,
// as well as the field to perform arithmetic operations.
type Prio3[
	Measurement, Aggregate any,
	T flp[Measurement, Aggregate, V, E, F],
	V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] struct {
	flp      T
	xof      xofTS[V, E]
	randSize uint
	shares   uint8
}

func New[
	T flp[Measurement, Aggregate, V, E, F],
	Measurement, Aggregate any,
	V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
](f T, algorithmID uint32, numShares uint8, context []byte,
) (v Prio3[Measurement, Aggregate, T, V, E, F], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shard takes a measurement and return a set of shares and a public share
// used for verification.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.1
func (v *Prio3[M, A, T, V, E, F]) Shard(
	measurement M, nonce *Nonce, rand []byte,
) (PublicShare, []InputShare[V, E], error) {
	_ = "STUB: not implemented"
	return *new(PublicShare), nil, nil
}

// FLPs without joint randomness.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.1.1
func (v *Prio3[M, A, T, V, E, F]) shardNoJointRand(
	meas V, seeds []byte,
) ([]InputShare[V, E], error) {
	_ = "STUB: not implemented"
	// Each Aggregator's input share contains its measurement share
	// and its share of the proof.
	return nil, nil
}

// Generate proof of valid measurement.

// Shard the encoded measurement and proof into shares.

// FLPs with joint randomness.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.1.2
func (v *Prio3[M, A, T, V, E, F]) shardWithJointRand(
	meas V, nonce *Nonce, seeds []byte,
) (PublicShare, []InputShare[V, E], error) {
	_ = "STUB: not implemented"
	// Each Aggregator's input share contains its measurement share,
	// share of proof, and blind. The public share contains the
	// Aggregators' joint randomness parts.
	return *new(PublicShare), nil, nil
}

// Shard the encoded measurement into shares and compute the
// joint randomness parts.

// Calculate leader's jointRandPart after leader's measShare
// has been calculated.

// Generate proof of valid measurement.

// Shard the proof into shares.

// PrepInit is used by each aggregator to begin the preparation phase.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.2
func (v *Prio3[M, A, T, V, E, F]) PrepInit(
	verifyKey *VerifyKey, nonce *Nonce, aggID uint8,
	ps PublicShare, inputShare InputShare[V, E],
) (*PrepState[V, E], *PrepShare[V, E], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Compute the joint randomness.

// Query the measurement and proof share.

// PrepSharesToPrep is the deterministic preparation message pre-processing
// algorithm. It combines the prep shares produced by the Aggregators in the
// previous round into the prep message consumed by each Aggregator to start
// the next round.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.2
func (v *Prio3[M, A, T, V, E, F]) PrepSharesToPrep(
	prepShares []PrepShare[V, E],
) (*PrepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unshard the verifier shares into the verifier message.

// Verify that each proof is well-formed and input is valid.

// Combine the joint randomness parts computed by the
// Aggregators into the true joint randomness seed. This is
// used in the last step.

// PrepNext is used by each aggregator to produce its output share.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.2
func (v *Prio3[M, A, T, V, E, F]) PrepNext(
	state *PrepState[V, E], msg *PrepMessage,
) (*OutShare[V, E], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateInit is used to start aggregation.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.4
func (v *Prio3[M, A, T, V, E, F]) AggregateInit() (s AggShare[V, E]) {
	_ = "STUB: not implemented"
	return nil
}

// AggregateUpdate aggregates an output share into an aggregation share.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.4
func (v *Prio3[M, A, T, V, E, F]) AggregateUpdate(
	aggShare *AggShare[V, E], outShare *OutShare[V, E],
) {
	_ = "STUB: not implemented"
	return
}

// aggregateMerge merges several aggregation shares.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.4
func (v *Prio3[M, A, T, V, E, F]) aggregateMerge(
	aggShares []AggShare[V, E],
) (s AggShare[V, E]) {
	_ = "STUB: not implemented"
	return nil
}

// Unshard is used by the Collector to recover the aggregate result from a set
// of aggregation shares.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.2.5
func (v *Prio3[M, A, T, V, E, F]) Unshard(
	aggShares []AggShare[V, E], numMeas uint,
) (*A, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Prio3[M, A, T, V, E, F]) getInputShareContent(
	aggID uint8, inShare InputShare[V, E], params *Params,
) (s inputShareContent[V, E], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Params struct {
	jointRandLen   uint
	measurementLen uint
	outputLen      uint
	evalOutputLen  uint
	queryRandLen   uint
	proveRandLen   uint
	proofLen       uint
	verifierLen    uint
	randSize       uint
	shares         uint8
}

func (p *Params) JointRandLength() uint   { _ = "STUB: not implemented"; return 0 }
func (p *Params) MeasurementLength() uint { _ = "STUB: not implemented"; return 0 }
func (p *Params) OutputLength() uint      { _ = "STUB: not implemented"; return 0 }
func (p *Params) EvalOutputLength() uint  { _ = "STUB: not implemented"; return 0 }
func (p *Params) QueryRandLength() uint   { _ = "STUB: not implemented"; return 0 }
func (p *Params) ProveRandLength() uint   { _ = "STUB: not implemented"; return 0 }
func (p *Params) ProofLength() uint       { _ = "STUB: not implemented"; return 0 }
func (p *Params) VerifierLength() uint    { _ = "STUB: not implemented"; return 0 }
func (p *Params) RandSize() uint          { _ = "STUB: not implemented"; return 0 }
func (p *Params) Shares() uint8           { _ = "STUB: not implemented"; return 0 }

func (v *Prio3[M, A, T, V, E, F]) Params() Params { _ = "STUB: not implemented"; return *new(Params) }

type flp[
	Measurement, AggResult any,
	V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] interface {
	MeasurementLength() uint
	JointRandLength() uint
	OutputLength() uint
	EvalOutputLength() uint
	ProveRandLength() uint
	ProofLength() uint
	VerifierLength() uint
	QueryRandLength() uint
	// Prove returns a proof attesting to the validity of the given measurement.
	Prove(meas, proveRand, jointRand V) V
	// Query is the linear Query algorithm run by each verifier on its share of
	// the measurement and proof.
	Query(measShare, proofShare, queryRnd, jointRnd V, shares uint8) (V, error)
	// Decide returns true if the measurement from which it was generated is
	// valid.
	Decide(V) bool
	// Encode returns a vector of MeasurementLength() elements representing
	// a measurement of type [Measurement].
	Encode(Measurement) (V, error)
	// Truncate returns a vector of OutputLength() elements representing
	// (a share of) an aggregatable output.
	Truncate(V) V
	// Decode returns an aggregate result of type [AggResult].
	// This computation may depend on the number of outputs aggregated.
	Decode(V, uint) (*AggResult, error)
}

var (
	ErrNumShares     = errors.New("invalid numshares, must be greater than 1")
	ErrContextSize   = errors.New("invalid context length, (0, MaxContextSize)")
	ErrNonceSize     = errors.New("invalid nonce length, (NonceSize)")
	ErrVerifyKeySize = errors.New("invalid verify key length, (VerifyKeySize)")
	ErrRandSize      = errors.New("invalid randomness length")
	ErrAggShareSize  = errors.New("invalid aggregate shares length")
	ErrAggID         = errors.New("invalid aggregation ID")
	ErrJointRand     = errors.New("invalid joint randomness")
	ErrShare         = errors.New("share was not provided")
	ErrProofVerify   = errors.New("proof verifier check failed")
)
