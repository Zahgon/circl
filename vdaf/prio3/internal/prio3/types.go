package prio3

import (
	"github.com/cloudflare/circl/vdaf/prio3/arith"
	"golang.org/x/crypto/cryptobyte"
)

const (
	SeedSize      uint = 32 // Size of Seed in bytes.
	NonceSize     uint = 16 // Size of Nonce in bytes.
	VerifyKeySize uint = 32 // Size of VerifyKey in bytes.
)

type (
	// Nonce is a public random value associated with the report.
	Nonce [NonceSize]byte
	// VerifyKey is a secret verification key held by each of the Aggregators.
	// This key is used to verify validity of the output shares they compute.
	VerifyKey [VerifyKeySize]byte
	// Seed is used to feed an extendable output function.
	Seed [SeedSize]byte
)

func (s *Seed) Marshal(b *cryptobyte.Builder) error { _ = "STUB: not implemented"; return nil }

func (s *Seed) Unmarshal(str *cryptobyte.String) bool { _ = "STUB: not implemented"; return false }

// PublicShare must be distributed to each of the Aggregators.
// Its content depends on whether joint randomness is required for the
// underlying FLP.
// If joint randomness is not used, then the public share is an empty slice.
//
//	struct {
//	    Prio3Seed joint_rand_parts[SEED_SIZE * prio3.SHARES];
//	} Prio3PublicShareWithJointRand;
type PublicShare []byte

func (s *PublicShare) New(p *Params) *PublicShare { _ = "STUB: not implemented"; return nil }

func (s *PublicShare) Marshal(b *cryptobyte.Builder) error { _ = "STUB: not implemented"; return nil }

func (s *PublicShare) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *PublicShare) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *PublicShare) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

type inputShareContent[V arith.Vec[V, E], E arith.Elt] struct {
	blind      *Seed
	proofShare V
	measShare  V
}

func (s *inputShareContent[V, E]) New(p *Params) *inputShareContent[V, E] {
	_ = "STUB: not implemented"
	return nil
}

func (s *inputShareContent[V, E]) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *inputShareContent[V, E]) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

// InputShareLeader represents one of these two structures.
//
//	struct {
//	    Prio3Field meas_share[F * prio3.flp.MEAS_LEN];
//	    Prio3Field proofs_share[F * prio3.flp.PROOF_LEN * prio3.PROOFS];
//	} Prio3LeaderShare;
//
//	struct {
//	    Prio3LeaderShare inner;
//	    Prio3Seed blind;
//	} Prio3LeaderShareWithJointRand;
type InputShareLeader[V arith.Vec[V, E], E arith.Elt] struct {
	inputShareContent[V, E]
}

func (s *InputShareLeader[V, E]) New(p *Params) *InputShareLeader[V, E] {
	_ = "STUB: not implemented"
	return nil
}

// InputShareHelper represents one of these two structures.
//
//	struct {
//	    Prio3Seed share;
//	} Prio3HelperShare;
//
//	struct {
//	    Prio3HelperShare inner;
//	    Prio3Seed blind;
//	} Prio3HelperShareWithJointRand;
type InputShareHelper struct {
	blind *Seed
	share Seed
}

func (s *InputShareHelper) New(p *Params) *InputShareHelper { _ = "STUB: not implemented"; return nil }

func (s *InputShareHelper) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InputShareHelper) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

// InputShare is a generic struct that stores shares for the leader or helper.
type InputShare[V arith.Vec[V, E], E arith.Elt] struct {
	leader *InputShareLeader[V, E]
	helper *InputShareHelper
}

func (s *InputShare[V, E]) New(p *Params, aggID uint) *InputShare[V, E] {
	_ = "STUB: not implemented"
	return nil
}

func (s *InputShare[V, E]) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InputShare[V, E]) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *InputShare[V, E]) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InputShare[V, E]) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

// InputShareHelper represents one of these two structures.
//
//	struct {
//	    Prio3Field verifiers_share[F * V];
//	} Prio3PrepShare;
//
//	struct {
//	    Prio3Field verifiers_share[F * V];
//	    Prio3Seed joint_rand_part;
//	} Prio3PrepShareWithJointRand;
type PrepShare[V arith.Vec[V, E], E arith.Elt] struct {
	jointRandPart  *Seed
	verifiersShare V
}

func (s *PrepShare[V, E]) New(p *Params) *PrepShare[V, E] { _ = "STUB: not implemented"; return nil }

func (s *PrepShare[V, E]) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PrepShare[V, E]) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *PrepShare[V, E]) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PrepShare[V, E]) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

type PrepState[V arith.Vec[V, E], E arith.Elt] struct {
	correctedJointRandSeed *Seed
	outShare               V
}

func (s *PrepState[V, E]) New(p *Params) *PrepState[V, E] { _ = "STUB: not implemented"; return nil }

func (s *PrepState[V, E]) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PrepState[V, E]) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *PrepState[V, E]) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PrepState[V, E]) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

// PrepMessage represents the following structure.
//
//	struct {
//	    Prio3Seed joint_rand;
//	} Prio3PrepMessageWithJointRand;
type PrepMessage struct{ joinRand *Seed }

func (s *PrepMessage) New(p *Params) *PrepMessage { _ = "STUB: not implemented"; return nil }

func (s *PrepMessage) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *PrepMessage) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *PrepMessage) Marshal(b *cryptobyte.Builder) error { _ = "STUB: not implemented"; return nil }

func (s *PrepMessage) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

type OutShare[V arith.Vec[V, E], E arith.Elt] struct{ share V }

func (s *OutShare[V, E]) New(p *Params) *OutShare[V, E] { _ = "STUB: not implemented"; return nil }

func (s *OutShare[V, E]) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *OutShare[V, E]) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *OutShare[V, E]) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *OutShare[V, E]) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}

// AggShare represents the following structure.
//
//	struct {
//	    Prio3Field agg_share[F * prio3.flp.OUTPUT_LEN];
//	} Prio3AggShare;
type AggShare[V arith.Vec[V, E], E arith.Elt] struct{ share V }

func (s *AggShare[V, E]) New(p *Params) *AggShare[V, E] { _ = "STUB: not implemented"; return nil }

func (s *AggShare[V, E]) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *AggShare[V, E]) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *AggShare[V, E]) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *AggShare[V, E]) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}
