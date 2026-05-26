package tkn

import (
	"io"

	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

// matrixG2 represents a matrix of G2 elements. They are stored in row-major order.
type matrixG2 struct {
	rows    int
	cols    int
	entries []pairing.G2
}

func (m *matrixG2) marshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *matrixG2) unmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// exp computes the naive matrix exponential of a with respect to the basepoint.
func (m *matrixG2) exp(a *matrixZp) { _ = "STUB: not implemented"; return }

// clear sets m to the zero matrix
func (m *matrixG2) clear() { _ = "STUB: not implemented"; return }

// resize only changes the matrix if we have to
func (m *matrixG2) resize(r int, c int) { _ = "STUB: not implemented"; return }

// conformal returns true iff m and a have the same dimensions.
func (m *matrixG2) conformal(a *matrixG2) bool { _ = "STUB: not implemented"; return false }

// Equal returns true if m == b.
func (m *matrixG2) Equal(b *matrixG2) bool { _ = "STUB: not implemented"; return false }

// set sets m to a.
func (m *matrixG2) set(a *matrixG2) { _ = "STUB: not implemented"; return }

// add sets m to a+b.
func (m *matrixG2) add(a *matrixG2, b *matrixG2) { _ = "STUB: not implemented"; return }

// leftMult multiples a*b with a matrixZp, b matrixG2.
func (m *matrixG2) leftMult(a *matrixZp, b *matrixG2) { _ = "STUB: not implemented"; return }

// rightMult multiplies a*b with a matrixG1, b matrixZp.
func (m *matrixG2) rightMult(a *matrixG2, b *matrixZp) { _ = "STUB: not implemented"; return }

func newMatrixG2(r int, c int) *matrixG2 { _ = "STUB: not implemented"; return nil }

func randomMatrixG2(rand io.Reader, r int, c int) (*matrixG2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
