package tkn

import (
	"io"

	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

// matrixGT represents a matrix of GT elements. They are stored in row-major order.
type matrixGT struct {
	rows    int
	cols    int
	entries []pairing.Gt
}

func (m *matrixGT) marshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *matrixGT) unmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// exp computes the naive matrix exponential of a with respect to the basepoint.
func (m *matrixGT) exp(a *matrixZp) { _ = "STUB: not implemented"; return }

// resize sets up m to be r x c
func (m *matrixGT) resize(r int, c int) { _ = "STUB: not implemented"; return }

// clear sets m to be the "zero" matrix
func (m *matrixGT) clear() { _ = "STUB: not implemented"; return }

// conformal returns true iff m and a have the same dimensions.
func (m *matrixGT) conformal(a *matrixGT) bool { _ = "STUB: not implemented"; return false }

// Equal returns true if m == b.
func (m *matrixGT) Equal(b *matrixGT) bool { _ = "STUB: not implemented"; return false }

// set sets m to b.
func (m *matrixGT) set(b *matrixGT) { _ = "STUB: not implemented"; return }

// add sets m to a+b.
func (m *matrixGT) add(a *matrixGT, b *matrixGT) { _ = "STUB: not implemented"; return }

// leftMult multiples a*b with a matrixZp, b matrixGT.
func (m *matrixGT) leftMult(a *matrixZp, b *matrixGT) { _ = "STUB: not implemented"; return }

// rightMult multiplies a*b with a matrixG1, b matrixZp.
func (m *matrixGT) rightMult(a *matrixGT, b *matrixZp) { _ = "STUB: not implemented"; return }

// to transpose can index bt[i,j] as b.entries[j*b.rows+i]

func newMatrixGT(r int, c int) *matrixGT { _ = "STUB: not implemented"; return nil }

func randomMatrixGT(rand io.Reader, r int, c int) (*matrixGT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
