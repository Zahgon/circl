package slhdsa

// See FIPS 205 -- Section 8
// Forest of Random Subsets (FORS) is a few-time signature scheme that is
// used to sign the digests of the actual messages.

type (
	forsPublicKey  []byte     // n bytes
	forsPrivateKey []byte     // n bytes
	forsSignature  []forsPair // k*forsPairSize() bytes
	forsPair       struct {
		sk   forsPrivateKey // forsSkSize() bytes
		auth [][]byte       // a*n bytes
	} // forsSkSize() + a*n bytes
)

func (p *params) forsMsgSize() uint32  { _ = "STUB: not implemented"; return 0 }
func (p *params) forsPkSize() uint32   { _ = "STUB: not implemented"; return 0 }
func (p *params) forsSkSize() uint32   { _ = "STUB: not implemented"; return 0 }
func (p *params) forsSigSize() uint32  { _ = "STUB: not implemented"; return 0 }
func (p *params) forsPairSize() uint32 { _ = "STUB: not implemented"; return 0 }

func (fs *forsSignature) fromBytes(p *params, c *cursor) { _ = "STUB: not implemented"; return }

func (fp *forsPair) fromBytes(p *params, c *cursor) { _ = "STUB: not implemented"; return }

// See FIPS 205 -- Section 8.1 -- Algorithm 14.
func (s *statePriv) forsSkGen(addr address, idx uint32) forsPrivateKey {
	_ = "STUB: not implemented"
	return *new(forsPrivateKey)
}

// See FIPS 205 -- Section 8.2 -- Algorithm 15 -- Iterative version.
//
// This is a stack-based implementation that computes the tree leaves
// in order (from the left to the right).
// Its recursive version can be found at fors_test.go file.
func (s *statePriv) forsNodeIter(
	stack stackNode, root []byte, i, z uint32, addr address,
) {
	_ = "STUB: not implemented"
	return
}

// See FIPS 205 -- Section 8.3 -- Algorithm 16.
func (s *statePriv) forsSign(sig forsSignature, digest []byte, addr address) {
	_ = "STUB: not implemented"
	return
}

// See FIPS 205 -- Section 8.4 -- Algorithm 17.
func (s *state) forsPkFromSig(
	sig forsSignature, digest []byte, addr address,
) (pk forsPublicKey) {
	_ = "STUB: not implemented"
	return *new(forsPublicKey)
}
