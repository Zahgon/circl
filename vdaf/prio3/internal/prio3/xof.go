package prio3

import (
	"github.com/cloudflare/circl/internal/sha3"
	"github.com/cloudflare/circl/vdaf/prio3/arith"
)

// xofTS allows to derive seeds and vector of elements from TurboSHAKE.
type xofTS[V arith.Vec[V, E], E arith.Elt] struct {
	usage  *[2]byte
	Header []byte
	sha3.State
}

// NewXof returns an xof based on TurboSHAKE given a VDAF ID and an application
// context string.
func NewXof[V arith.Vec[V, E], E arith.Elt](
	algorithmID uint32, context []byte,
) (x xofTS[V, E], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// | 2 | uint16(len(dst)) | little-endian |
// | 1 | Version          |
// | 1 | AlgoClass        |
// | 4 | ID               | big-endian    |
// | 2 | Usage            | big-endian    |
// | * | context          |
// | 1 | SeedSize         |

func (x *xofTS[V, E]) Init(usage uint16, s *Seed) error { _ = "STUB: not implemented"; return nil }

func (x *xofTS[V, E]) SetBinderByte(binder ...byte) error { _ = "STUB: not implemented"; return nil }

func (x *xofTS[V, E]) SetBinderBytes(binder ...[]byte) error { _ = "STUB: not implemented"; return nil }

func (x *xofTS[V, E]) helperMeasShareEnc(
	encOut []byte, out V, aggID uint8, s *Seed,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xofTS[V, E]) helperMeasShare(out V, aggID uint8, s *Seed) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xofTS[V, E]) helperProofsShare(out V, aggID uint8, s *Seed) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xofTS[V, E]) proveRands(out V, proveSeed *Seed) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xofTS[V, E]) queryRands(out V, k *VerifyKey, nonce *Nonce) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xofTS[V, E]) jointRandPart(
	out []byte, blind *Seed, aggID uint8, nonce *Nonce, measShareEnc []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xofTS[V, E]) jointRandSeed(jointRandParts []byte) (s Seed, err error) {
	_ = "STUB: not implemented"
	return *new(Seed), nil
}

func (x *xofTS[V, E]) jointRands(out V, jointRandSeed *Seed) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	usageMeasuShare uint16 = iota + 1
	usageProofShare
	usageJointRandomness
	usageProveRandomness
	usageQueryRandomness
	usageJointRandSeed
	usageJointRandPart
)
