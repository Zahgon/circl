package group

import (
	"math/big"

	"github.com/cloudflare/circl/expander"
)

// HashToField generates a set of elements {u1,..., uN} = Hash(b) where each
// u in GF(p) and L is the security parameter.
func HashToField(u []big.Int, b []byte, e expander.Expander, p *big.Int, L uint) {
	_ = "STUB: not implemented"
	return
}
