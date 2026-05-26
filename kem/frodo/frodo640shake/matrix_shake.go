package frodo640shake

import (
	"github.com/cloudflare/circl/internal/sha3"
)

func expandSeedIntoA(A *nByNU16, seed *[seedASize]byte, xof *sha3.State) {
	_ = "STUB: not implemented"
	return
}

// No need to reduce modulo 2^15, extra bits are removed
// later on via packing or explicit reduction.

func mulAddASPlusE(out *nByNbarU16, A *nByNU16, s *nByNbarU16, e *nByNbarU16) {
	_ = "STUB: not implemented"
	return
}

// No need to reduce modulo 2^15, extra bits are removed
// later on via packing or explicit reduction.

func mulAddSAPlusE(out *nbarByNU16, s []uint16, A *nByNU16, e []uint16) {
	_ = "STUB: not implemented"
	return
}

// No need to reduce modulo 2^15, extra bits are removed
// later on via packing or explicit reduction.
