package tkn

import (
	"io"

	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

// matrixG1 represents a matrix of G1 elements. They are stored in row-major order.
type matrixG1 struct {
	rows    int
	cols    int
	entries []pairing.G1
}

func (m *matrixG1) marshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *matrixG1) unmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// We write addition for each of these, multiplication by scalars, and action on both
// sides?

// exp computes the naive matrix exponential of a with respect to the basepoint.
func (m *matrixG1) exp(a *matrixZp) { _ = "STUB: not implemented"; return }

// resize only changes the matrix if we have to
func (m *matrixG1) resize(r int, c int) { _ = "STUB: not implemented"; return }

// clear sets a matrix to the all zero matrix
func (m *matrixG1) clear() { _ = "STUB: not implemented"; return }

// conformal returns true iff m and a have the same dimensions.
func (m *matrixG1) conformal(a *matrixG1) bool { _ = "STUB: not implemented"; return false }

// Equal returns true if m == b.
func (m *matrixG1) Equal(b *matrixG1) bool { _ = "STUB: not implemented"; return false }

// set sets m to a.
func (m *matrixG1) set(a *matrixG1) { _ = "STUB: not implemented"; return }

// add sets m to a+b.
func (m *matrixG1) add(a *matrixG1, b *matrixG1) { _ = "STUB: not implemented"; return }

// sub sets m to a-b.
func (m *matrixG1) sub(a *matrixG1, b *matrixG1) { _ = "STUB: not implemented"; return }

// leftMult multiples a*b with a matrixZp, b matrixG1.
func (m *matrixG1) leftMult(a *matrixZp, b *matrixG1) { _ = "STUB: not implemented"; return }

// rightMult multiplies a*b with a matrixG1, b matrixZp.
func (m *matrixG1) rightMult(a *matrixG1, b *matrixZp) { _ = "STUB: not implemented"; return }

// scalarMult sets m to c*a where c is a Scalar,
func (m *matrixG1) scalarMult(c *pairing.Scalar, a *matrixG1) { _ = "STUB: not implemented"; return }

// copy creates and returns a new matrix that shares no storage with m
func (m *matrixG1) copy() *matrixG1 { _ = "STUB: not implemented"; return nil }

func newMatrixG1(r int, c int) *matrixG1 { _ = "STUB: not implemented"; return nil }

func randomMatrixG1(rand io.Reader, r int, c int) (*matrixG1, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// oracle generates 3x2 matrices via the random oracle
func oracle(input []byte) (*matrixG1, *matrixG1) { _ = "STUB: not implemented"; return nil, nil }

// transpose sets m to the transpose of a.
// Not aliasing safe
func (m *matrixG1) transpose(a *matrixG1) { _ = "STUB: not implemented"; return }
