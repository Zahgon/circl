package goldilocks

// ScalarSize is the size (in bytes) of scalars.
const ScalarSize = 56 // 448 / 8

// _N is the number of 64-bit words to store scalars.
const _N = 7 // 448 / 64

// Scalar represents a positive integer stored in little-endian order.
type Scalar [ScalarSize]byte

type scalar64 [_N]uint64

func (z *scalar64) fromScalar(x *Scalar) { _ = "STUB: not implemented"; return }

func (z *scalar64) toScalar(x *Scalar) { _ = "STUB: not implemented"; return }

// add calculates z = x + y. Assumes len(z) > max(len(x),len(y)).
func add(z, x, y []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// sub calculates z = x - y. Assumes len(z) > max(len(x),len(y)).
func sub(z, x, y []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// mulWord calculates z = x * y. Assumes len(z) >= len(x)+1.
func mulWord(z, x []uint64, y uint64) { _ = "STUB: not implemented"; return }

// Cmov moves x into z if b=1.
func (z *scalar64) Cmov(b uint64, x *scalar64) { _ = "STUB: not implemented"; return }

// leftShift shifts to the left the words of z returning the more significant word.
func (z *scalar64) leftShift(low uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// reduceOneWord calculates z = z + 2^448*x such that the result fits in a Scalar.
func (z *scalar64) reduceOneWord(x uint64) { _ = "STUB: not implemented"; return }

// modOrder reduces z mod order.
func (z *scalar64) modOrder() { _ = "STUB: not implemented"; return }

// Performs: while (z >= order) { z = z-order }
// At most 8 (eight) iterations reduce 3 bits by subtracting.

// (c || x) = z-order
// if c != 0 { z = x }

// FromBytes stores z = x mod order, where x is a number stored in little-endian order.
func (z *Scalar) FromBytes(x []byte) { _ = "STUB: not implemented"; return }

// divBy4 calculates z = x/4 mod order.
func (z *Scalar) divBy4(x *Scalar) {
	_ = "STUB: not implemented"

	// Red reduces z mod order.
	return
}

func (z *Scalar) Red() { _ = "STUB: not implemented"; return }

// Neg calculates z = -z mod order.
func (z *Scalar) Neg() {
	_ = "STUB: not implemented"

	// Add calculates z = x+y mod order.
	return
}

func (z *Scalar) Add(x, y *Scalar) { _ = "STUB: not implemented"; return }

// Sub calculates z = x-y mod order.
func (z *Scalar) Sub(x, y *Scalar) { _ = "STUB: not implemented"; return }

// Mul calculates z = x*y mod order.
func (z *Scalar) Mul(x, y *Scalar) { _ = "STUB: not implemented"; return }

// IsZero returns true if z=0.
func (z *Scalar) IsZero() bool { _ = "STUB: not implemented"; return false }
