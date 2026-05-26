package slhdsa

// See FIPS 205 -- Section 7
// SLH-DSA uses a hypertree to sign the FORS keys.

type hyperTreeSignature []xmssSignature // d*xmssSigSize() bytes

func (p *params) hyperTreeSigSize() uint32 { _ = "STUB: not implemented"; return 0 }

func (hts *hyperTreeSignature) fromBytes(p *params, c *cursor) { _ = "STUB: not implemented"; return }

func nextIndex(idxTree *[3]uint32, n uint32) (idxLeaf uint32) { _ = "STUB: not implemented"; return 0 }

// See FIPS 205 -- Section 7.1 -- Algorithm 12.
func (s *statePriv) htSign(
	sig hyperTreeSignature, msg []byte, idxTree [3]uint32, idxLeaf uint32,
) {
	_ = "STUB: not implemented"
	return
}

// See FIPS 205 -- Section 7.2 -- Algorithm 13.
func (s *state) htVerify(
	msg, root []byte, idxTree [3]uint32, idxLeaf uint32, sig hyperTreeSignature,
) bool {
	_ = "STUB: not implemented"
	return false
}
