package rsa

import (
	"crypto/rsa"
	"io"
	"math/big"
)

// KeyShare represents a portion of the key. It can only be used to generate SignShare's. During the dealing phase (when Deal is called), one KeyShare is generated per player.
type KeyShare struct {
	si *big.Int

	twoDeltaSi *big.Int // optional cached value, this value is used to marginally speed up SignShare generation in Sign. If nil, it will be generated when needed and then cached.
	Index      uint     // When KeyShare's are generated they are each assigned an index sequentially

	Players   uint
	Threshold uint
}

func (kshare KeyShare) String() string { _ = "STUB: not implemented"; return "" }

// MarshalBinary encodes a KeyShare into a byte array in a format readable by UnmarshalBinary.
// Note: Only Index's up to math.MaxUint16 are supported
func (kshare *KeyShare) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// The encoding format is
	// | Players: uint16 | Threshold: uint16 | Index: uint16 | siLen: uint16 | si: []byte | twoDeltaSiNil: bool | twoDeltaSiLen: uint16 | twoDeltaSi: []byte |
	// with all values in big-endian.
	return nil, nil
}

// okay because of conditions checked above

// twoDeltaSiNil

// UnmarshalBinary recovers a KeyShare from a slice of bytes, or returns an error if the encoding is invalid.
func (kshare *KeyShare) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	// The encoding format is
	// | Players: uint16 | Threshold: uint16 | Index: uint16 | siLen: uint16 | si: []byte | twoDeltaSiNil: bool | twoDeltaSiLen: uint16 | twoDeltaSi: []byte |
	// with all values in big-endian.
	return nil
}

// Returns the cached value in twoDeltaSi or if nil, generates 2∆s_i, stores it in twoDeltaSi, and returns it
func (kshare *KeyShare) get2DeltaSi(players int64) *big.Int {
	_ = "STUB: not implemented"
	// use the cached value if it exists
	return nil
}

// 2∆s_i
// delta << 1 == delta * 2

// Sign msg using a KeyShare. msg MUST be padded and hashed. Call PadHash before this method.
//
// If rand is not nil then blinding will be used to avoid timing
// side-channel attacks.
//
// parallel indicates whether the blinding operations should use go routines to operate in parallel.
// If parallel is false, blinding will take about 2x longer than nonbinding, otherwise it will take about the same time
// (see benchmarks). If randSource is nil, parallel has no effect. parallel should almost always be set to true.
func (kshare *KeyShare) Sign(randSource io.Reader, pub *rsa.PublicKey, digest []byte, parallel bool) (SignShare, error) {
	_ = "STUB: not implemented"
	return *new(SignShare), nil
}

// Let's blind.
// We can't use traditional RSA blinding (as used in rsa.go) because we are exponentiating by exp and not d.
// As such, Euler's theorem doesn't apply ( exp * d != 0 (mod ϕ(n)) ).
// Instead, we will choose a random r and compute x^{exp+r} * x^{-r} = x^{exp}.
// This should (hopefully) prevent revealing information of the true value of exp, since with exp you can derive
// s_i, the secret key share.

// exp + r

// x^{|2∆s_i+r|}

// x^r

// x^{-r}

// extremely unlikely, somehow x^r is p or q

// x^{|2∆s_i+r|} * x^{-r} = x^{2∆s_i}

// x^{2∆s_i}
