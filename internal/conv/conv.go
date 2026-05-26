package conv

import (
	"math/big"

	"golang.org/x/crypto/cryptobyte"
)

// BytesLe2Hex returns an hexadecimal string of a number stored in a
// little-endian order slice x.
func BytesLe2Hex(x []byte) string { _ = "STUB: not implemented"; return "" }

// BytesLe2BigInt converts a little-endian slice x into a big-endian
// math/big.Int.
func BytesLe2BigInt(x []byte) *big.Int { _ = "STUB: not implemented"; return nil }

// BytesBe2Uint64Le converts a big-endian slice x to a little-endian slice of uint64.
func BytesBe2Uint64Le(x []byte) []uint64 { _ = "STUB: not implemented"; return nil }

// BigInt2BytesLe stores a positive big.Int number x into a little-endian slice z.
// The slice is modified if the bitlength of x <= 8*len(z) (padding with zeros).
// If x does not fit in the slice or is negative, z is not modified.
func BigInt2BytesLe(z []byte, x *big.Int) { _ = "STUB: not implemented"; return }

// Uint64Le2BigInt converts a little-endian slice x into a big number.
func Uint64Le2BigInt(x []uint64) *big.Int { _ = "STUB: not implemented"; return nil }

// Uint64Le2BytesLe converts a little-endian slice x to a little-endian slice of bytes.
func Uint64Le2BytesLe(x []uint64) []byte { _ = "STUB: not implemented"; return nil }

// Uint64Le2BytesBe converts a little-endian slice x to a big-endian slice of bytes.
func Uint64Le2BytesBe(x []uint64) []byte { _ = "STUB: not implemented"; return nil }

// Uint64Le2Hex returns an hexadecimal string of a number stored in a
// little-endian order slice x.
func Uint64Le2Hex(x []uint64) string { _ = "STUB: not implemented"; return "" }

// BigInt2Uint64Le stores a positive big.Int number x into a little-endian slice z.
// The slice is modified if the bitlength of x <= 8*len(z) (padding with zeros).
// If x does not fit in the slice or is negative, z is not modified.
func BigInt2Uint64Le(z []uint64, x *big.Int) { _ = "STUB: not implemented"; return }

// number of 64-bit words

// MarshalBinary encodes a value into a byte array in a format readable by UnmarshalBinary.
func MarshalBinary(v cryptobyte.MarshalingValue) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalBinaryLen encodes a value into an array of n bytes in a format readable by UnmarshalBinary.
func MarshalBinaryLen(v cryptobyte.MarshalingValue, length uint) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A UnmarshalingValue decodes itself from a cryptobyte.String and advances the pointer.
// It reports whether the read was successful.
type UnmarshalingValue interface {
	Unmarshal(*cryptobyte.String) bool
}

// UnmarshalBinary recovers a value from a byte array.
// It returns an error if the read was unsuccessful.
func UnmarshalBinary(v UnmarshalingValue, data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
