package tkn

import (
	"io"

	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

// matrixZp represents a matrix of mod grouporder  elements. They are stored in row-major order.
// The name is a gesture toward the paper.
type matrixZp struct {
	rows    int
	cols    int
	entries []pairing.Scalar
}

func (m *matrixZp) marshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *matrixZp) unmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// sampleDlin samples from the distribution Dk.
// See section 3.2 of the paper for details.
func sampleDlin(rand io.Reader) (*matrixZp, error) { _ = "STUB: not implemented"; return nil, nil }

func randomMatrixZp(rand io.Reader, r int, c int) (*matrixZp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We adopt the interface that math.Big uses
// Receivers get set to the results of operations, and return themselves.
// All aliases are allowed.

// Now recall that G1, G2, GT are acted on by Zp.
// So we don't have all products, just left and right with Zp.
// Zp doesn't need this distinction.

// Errors are signalled by returning nil, which propagates.

// initialize sets up m to be r x c
func (m *matrixZp) resize(r int, c int) { _ = "STUB: not implemented"; return }

// clear makes m an all 0
func (m *matrixZp) clear() { _ = "STUB: not implemented"; return }

func newMatrixZp(r int, c int) *matrixZp { _ = "STUB: not implemented"; return nil }

// eye returns the k by k identity matrix.
func eye(k int) *matrixZp { _ = "STUB: not implemented"; return nil }

// conformal returns true iff m and a have the same dimensions.
func (m *matrixZp) conformal(a *matrixZp) bool { _ = "STUB: not implemented"; return false }

// Equal returns true iff m == a.
func (m *matrixZp) Equal(a *matrixZp) bool { _ = "STUB: not implemented"; return false }

// set sets m to a.
func (m *matrixZp) set(a *matrixZp) { _ = "STUB: not implemented"; return }

// add sets m to a+b.
func (m *matrixZp) add(a *matrixZp, b *matrixZp) { _ = "STUB: not implemented"; return }

// sub sets m to a-b.
func (m *matrixZp) sub(a *matrixZp, b *matrixZp) { _ = "STUB: not implemented"; return }

// mul sets m to a*b.
func (m *matrixZp) mul(a *matrixZp, b *matrixZp) { _ = "STUB: not implemented"; return }

// transpose sets m to the transpose of a.
func (m *matrixZp) transpose(a *matrixZp) { _ = "STUB: not implemented"; return }

// swaprows swaps two rows.
func (m *matrixZp) swapRows(i int, j int) { _ = "STUB: not implemented"; return }

// scalerow scales a row.
func (m *matrixZp) scaleRow(alpha *pairing.Scalar, i int) { _ = "STUB: not implemented"; return }

// addscaledrow takes alpha * row i and adds it to row j.
func (m *matrixZp) addScaledRow(alpha *pairing.Scalar, i int, j int) {
	_ = "STUB: not implemented"
	return
}

// inverse sets m to the inverse of a. If a is not invertible,
// the result is undefined and an error is returned.
// Aliasing safe
func (m *matrixZp) inverse(a *matrixZp) error { _ = "STUB: not implemented"; return nil }

// Any way we slice it we need additional storage.

// Gaussian elimination with pivoting begins here.

// At this point the matrix is in reduced row echelon form.
// The next step is to substitute back.

// prf computes a prf with output in pairs of 3x2 matrices
func prf(key []byte, input []byte) (*matrixZp, *matrixZp, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// scalarmul sets m to a matrix a*B
func (m *matrixZp) scalarmul(a *pairing.Scalar, b *matrixZp) { _ = "STUB: not implemented"; return }

// colsel sets m to a matrix with the selected columns.
func (m *matrixZp) colsel(a *matrixZp, cols []int) { _ = "STUB: not implemented"; return }

func (m *matrixZp) String() string { _ = "STUB: not implemented"; return "" }
