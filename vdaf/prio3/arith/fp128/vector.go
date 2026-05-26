// Code generated from ./templates/vector.go.tmpl. DO NOT EDIT.

package fp128

import (
	"io"

	"github.com/cloudflare/circl/internal/sha3"
	"golang.org/x/crypto/cryptobyte"
)

type Vec []Fp

func (v Vec) Size() uint { _ = "STUB: not implemented"; return 0 }

func (v Vec) AddAssign(x Vec) { _ = "STUB: not implemented"; return }

func (v Vec) SubAssign(x Vec) { _ = "STUB: not implemented"; return }

func (v Vec) ScalarMul(x *Fp) { _ = "STUB: not implemented"; return }

func (v Vec) DotProduct(x Vec) (out Fp) { _ = "STUB: not implemented"; return *new(Fp) }

func bitRev(x uint, numBits uint) uint { _ = "STUB: not implemented"; return 0 }

func (v Vec) InvNTT(values Vec, n uint) { _ = "STUB: not implemented"; return }

func (v Vec) NTT(values Vec, n uint) { _ = "STUB: not implemented"; return }

func (v Vec) SplitBits(n uint64) error { _ = "STUB: not implemented"; return nil }

func (v Vec) JoinBits() Fp { _ = "STUB: not implemented"; return *new(Fp) }

func (v Vec) Random(rnd io.Reader) error { _ = "STUB: not implemented"; return nil }

func (v Vec) RandomSHA3(s *sha3.State) error { _ = "STUB: not implemented"; return nil }

func (v Vec) RandomSHA3Bytes(out []byte, s *sha3.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (v Vec) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v Vec) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (v Vec) Marshal(b *cryptobyte.Builder) error { _ = "STUB: not implemented"; return nil }

func (v Vec) Unmarshal(s *cryptobyte.String) bool { _ = "STUB: not implemented"; return false }

func isInRange(b *[Size]byte) (out [2]uint64, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func mustSameLen[T ~[]E, E any](x, y T) { _ = "STUB: not implemented"; return }

func mustSumLen[T ~[]E, E any](z, x, y T) { _ = "STUB: not implemented"; return }
