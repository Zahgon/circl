package slhdsa

// See FIPS 205 -- Section 6
// eXtended Merkle Signature Scheme (XMSS) extends the WOTS+ signature
// scheme into one that can sign multiple messages.

type (
	xmssPublicKey []byte // n bytes
	xmssSignature struct {
		wotsSig  wotsSignature // wotsSigSize() bytes
		authPath []byte        // hPrime*n bytes
	} // wotsSigSize() + hPrime*n bytes
)

func (p *params) xmssPkSize() uint32       { _ = "STUB: not implemented"; return 0 }
func (p *params) xmssAuthPathSize() uint32 { _ = "STUB: not implemented"; return 0 }
func (p *params) xmssSigSize() uint32      { _ = "STUB: not implemented"; return 0 }

func (xs *xmssSignature) fromBytes(p *params, c *cursor) { _ = "STUB: not implemented"; return }

// See FIPS 205 -- Section 6.1 -- Algorithm 9 -- Iterative version.
//
// This is a stack-based implementation that computes the tree leaves
// in order (from the left to the right).
// Its recursive version can be found at xmss_test.go file.
func (s *statePriv) xmssNodeIter(
	stack stackNode, root []byte, i, z uint32, addr address,
) {
	_ = "STUB: not implemented"
	return
}

// See FIPS 205 -- Section 6.2 -- Algorithm 10.
func (s *statePriv) xmssSign(
	stack stackNode, sig xmssSignature, msg []byte, idx uint32, addr address,
) {
	_ = "STUB: not implemented"
	return
}

// See FIPS 205 -- Section 6.3 -- Algorithm 11.
func (s *state) xmssPkFromSig(
	out xmssPublicKey, msg []byte, sig xmssSignature, idx uint32, addr address,
) {
	_ = "STUB: not implemented"
	return
}
