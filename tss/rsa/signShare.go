package rsa

import (
	"math/big"
)

// SignShare represents a portion of a signature. It is generated when a message is signed by a KeyShare. t SignShare's are then combined by calling CombineSignShares, where t is the Threshold.
type SignShare struct {
	xi *big.Int

	Index uint

	Players   uint
	Threshold uint
}

func (s SignShare) String() string { _ = "STUB: not implemented"; return "" }

// MarshalBinary encodes SignShare into a byte array in a format readable by UnmarshalBinary.
// Note: Only Index's up to math.MaxUint16 are supported
func (s *SignShare) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// | Players: uint16 | Threshold: uint16 | Index: uint16 | xiLen: uint16 | xi: []byte |
	return nil, nil
}

// UnmarshalBinary converts a byte array outputted from Marshall into a SignShare or returns an error if the value is invalid
func (s *SignShare) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	// | Players: uint16 | Threshold: uint16 | Index: uint16 | xiLen: uint16 | xi: []byte |
	return nil
}
