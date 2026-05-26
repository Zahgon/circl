// Package rsa provides RSA threshold signature scheme.
//
// This package implements the Protocol 1 of "Practical Threshold Signatures"
// by Victor Shoup [1].
//
// # References
//
// [1] https://www.iacr.org/archive/eurocrypt2000/1807/18070209-new.pdf
package rsa

import (
	"crypto"
	"crypto/rsa"
	"errors"
	"io"
	"math/big"
)

// GenerateKey generates a RSA keypair for its use in RSA threshold signatures.
// Internally, the modulus is the product of two safe primes. The time
// consumed by this function is relatively longer than the regular
// GenerateKey function from the crypto/rsa package.
func GenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check for different primes.

// check n has the desired bitlength.

// l or `Players`, the total number of Players.
// t, the number of corrupted Players.
// k=t+1 or `Threshold`, the number of signature shares needed to obtain a signature.

func validateParams(players, threshold uint) error { _ = "STUB: not implemented"; return nil }

// Deal takes in an existing RSA private key generated elsewhere. If cache is true, cached values are stored in KeyShare taking up more memory by reducing Sign time.
// See KeyShare documentation. Multi-prime RSA keys are unsupported.
func Deal(randSource io.Reader, players, threshold uint, key *rsa.PrivateKey, cache bool) ([]KeyShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// p = 2p' + 1
// q = 2q' + 1
// p' = (p - 1)/2
// q' = (q - 1)/2
// m = p'q' = (p - 1)(q - 1)/4

// p - 1

// q - 1

// (p - 1)(q - 1)

// >> 2 == / 4

// de ≡ 1

// a_0...a_{k-1}

// a_0 = d

// a_0...a_{k-1} = rand from {0, ..., m - 1}

// 1 <= i <= l

// Σ^{k-1}_{i=0} | a_i * X^i (mod m)

func calcN(p, q *big.Int) big.Int {
	_ = "STUB: not implemented"
	// n = pq
	return *new(big.Int)
}

// f(X) = Σ^{k-1}_{i=0} | a_i * X^i (mod m), where k = len(a).
func computePolynomial(a []*big.Int, x uint, m *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// PadHash MUST be called before signing a message
func PadHash(padder Padder, hash crypto.Hash, pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Sign(Pad(Hash(M)))
	return nil, nil
}

type Signature = []byte

// CombineSignShares combines t SignShare's to produce a valid signature
func CombineSignShares(pub *rsa.PublicKey, shares []SignShare, msg []byte) (Signature, error) {
	_ = "STUB: not implemented"
	return *new(Signature), nil
}

// i_1 ... i_k

// λ(S, 0, i)

// 2λ

// faster than TWO * lambda

// we need to handle negative λ's (aka inverse), so abs it, compare, and if necessary modinverse

// x_i^{|2λ|}

// TODO  first compute all the powers for the negative exponents (but don't invert yet); multiply these together and then invert all at once. This is ok since (ab)^-1 = a^-1 b^-1

// e′ = 4∆^2

// faster than delta^TWO
// faster than FOUR * eprime

// e′a + eb = 1

// TODO You can compute a earlier and multiply a into the exponents used when computing w.
// w^a

// TODO justification
// x^b

// TODO justification
// y = w^a * x^b

// verify that signature is valid by checking x == y^e.

// ensure signature has the right size.

func checkIndices(i, j, l int64, S []SignShare) bool { _ = "STUB: not implemented"; return false }

// j must be in S
// i must be in {0..l} but not in S

// computes Lagrange Interpolation for the shares
// i must be in {0..l} but not in S
// j must be in S
func computeLambda(delta *big.Int, S []SignShare, i, j, l int64) (*big.Int, error) {
	_ = "STUB: not implemented"
	// Equation (2) of https://www.iacr.org/archive/eurocrypt2000/1807/18070209-new.pdf
	return nil, nil
}

// λ(s, i, j) = ∆( (  π{j'∈S\{j}} (i - j')  ) /  (  π{j'∈S\{j}} (j - j') ) )

// ∈ S

// j'

// (i - j')  for j' ∈ S \ {j}
// (j - j')  for j' ∈ S \ {j}

// den must be different of zero.

// lambda = (delta * num) / den, and lambda must an integer.

var ErrInvalidCalc = errors.New("tss/rsa: invalid calculation")
