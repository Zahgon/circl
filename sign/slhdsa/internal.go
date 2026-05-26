package slhdsa

// See FIPS 205 -- Section 9
// SLH-DSA Internal Functions

// See FIPS 205 -- Section 9.1 -- Algorithm 18.
func slhKeyGenInternal(
	p *params, skSeed, skPrf, pkSeed []byte,
) (pub PublicKey, priv PrivateKey) {
	_ = "STUB: not implemented"
	return *new(PublicKey), *new(PrivateKey)
}

func (p *params) parseMsg(
	digest []byte,
) (md []byte, idxTree [3]uint32, idxLeaf uint32) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

// See FIPS 205 -- Section 9.2 -- Algorithm 19.
func slhSignInternal(sk *PrivateKey, message, addRand []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See FIPS 205 -- Section 9.3 -- Algorithm 20.
func slhVerifyInternal(pub *PublicKey, message, sigBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// signature represents a SLH-DSA signature.
type signature struct {
	rnd     []byte             // n bytes
	forsSig forsSignature      // forsSigSize() bytes
	htSig   hyperTreeSignature // hyperTreeSigSize() bytes
}

func (p *params) SignatureSize() int { _ = "STUB: not implemented"; return 0 }

func (s *signature) fromBytes(p *params, c *cursor) bool { _ = "STUB: not implemented"; return false }
