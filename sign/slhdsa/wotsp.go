package slhdsa

// See FIPS 205 -- Section 5
// Winternitz One-Time Signature Plus Scheme

const (
	wotsW    uint32 = 16 // wotsW is w = 2^lg_w, where lg_w = 4.
	wotsLen2 uint32 = 3  // wotsLen2 is len_2 fixed to 3.
)

type (
	wotsPublicKey []byte // n bytes
	wotsSignature []byte // wotsLen()*n bytes
)

func (p *params) wotsSigSize() uint32 { _ = "STUB: not implemented"; return 0 }
func (p *params) wotsLen() uint32     { _ = "STUB: not implemented"; return 0 }
func (p *params) wotsLen1() uint32    { _ = "STUB: not implemented"; return 0 }

func (ws *wotsSignature) fromBytes(p *params, c *cursor) { _ = "STUB: not implemented"; return }

// See FIPS 205 -- Section 5 -- Algorithm 5.
func (s *state) chain(
	x []byte, index, steps uint32, addr address,
) (out []byte) {
	_ = "STUB: not implemented"
	return nil
}

// See FIPS 205 -- Section 5.1 -- Algorithm 6.
func (s *statePriv) wotsPkGen(addr address) wotsPublicKey {
	_ = "STUB: not implemented"
	return *new(wotsPublicKey)
}

// See FIPS 205 -- Section 5.2 -- Algorithm 7.
func (s *statePriv) wotsSign(sig wotsSignature, msg []byte, addr address) {
	_ = "STUB: not implemented"
	return
}

// Signs every nibble of the message and computes the checksum.

// Lastly, every nibble of the checksum is also signed.

// See FIPS 205 -- Section 5.3 -- Algorithm 8.
func (s *state) wotsPkFromSig(
	sig wotsSignature, msg []byte, addr address,
) wotsPublicKey {
	_ = "STUB: not implemented"
	return *new(wotsPublicKey)
}

// Signs every nibble of the message, computes the checksum, and
// feeds each signature to the T function.

// Every nibble of the checksum is also signed feeding the signature
// to the T function.

// Generates the public key as the output of the T function.
