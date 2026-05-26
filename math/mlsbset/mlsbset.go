// Package mlsbset provides a constant-time exponentiation method with precomputation.
//
// References: "Efficient and secure algorithms for GLV-based scalar
// multiplication and their implementation on GLV–GLS curves" by (Faz-Hernandez et al.)
//   - https://doi.org/10.1007/s13389-014-0085-7
//   - https://eprint.iacr.org/2013/158
package mlsbset

// EltG is a group element.
type EltG interface{}

// EltP is a precomputed group element.
type EltP interface{}

// Group defines the operations required by MLSBSet exponentiation method.
type Group interface {
	Identity() EltG                    // Returns the identity of the group.
	Sqr(x EltG)                        // Calculates x = x^2.
	Mul(x EltG, y EltP)                // Calculates x = x*y.
	NewEltP() EltP                     // Returns an arbitrary precomputed element.
	ExtendedEltP() EltP                // Returns the precomputed element x^(2^(w*d)).
	Lookup(a EltP, v uint, s, u int32) // Sets a = s*T[v][u].
}

// Params contains the parameters of the encoding.
type Params struct {
	T uint // T is the maximum size (in bits) of exponents.
	V uint // V is the number of tables.
	W uint // W is the window size.
	E uint // E is the number of digits per table.
	D uint // D is the number of digits in total.
	L uint // L is the length of the code.
}

// Encoder allows to convert integers into valid powers.
type Encoder struct{ p Params }

// New produces an encoder of the MLSBSet algorithm.
func New(t, v, w uint) (Encoder, error) { _ = "STUB: not implemented"; return *new(Encoder), nil }

// Encode converts an odd integer k into a valid power for exponentiation.
func (m Encoder) Encode(k []byte) (*Power, error) { _ = "STUB: not implemented"; return nil, nil }

// signs calculates the set of signs.
func (m Encoder) signs(k []byte) []int32 { _ = "STUB: not implemented"; return nil }

// GetParams returns the complementary parameters of the encoding.
func (m Encoder) GetParams() Params {
	_ = "STUB: not implemented"

	// tableSize returns the size of each table.
	return *new(Params)
}

func (m Encoder) tableSize() uint { _ = "STUB: not implemented"; return 0 }

// Elts returns the total number of elements that must be precomputed.
func (m Encoder) Elts() uint { _ = "STUB: not implemented"; return 0 }

// IsExtended returns true if the element x^(2^(wd)) must be calculated.
func (m Encoder) IsExtended() bool { _ = "STUB: not implemented"; return false }

// Ops returns the number of squares and multiplications executed during an exponentiation.
func (m Encoder) Ops() (S uint, M uint) { _ = "STUB: not implemented"; return 0, 0 }

func (m Encoder) String() string { _ = "STUB: not implemented"; return "" }
