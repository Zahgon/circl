package math

import (
	"io"
	"math/big"
)

// IsSafePrime reports whether p is (probably) a safe prime.
// The prime p=2*q+1 is safe prime if both p and q are primes.
// Note that ProbablyPrime is not suitable for judging primes
// that an adversary may have crafted to fool the test.
func IsSafePrime(p *big.Int) bool { _ = "STUB: not implemented"; return false }

// SafePrime returns a number of the given bit length that is a safe prime with high probability.
// The number returned p=2*q+1 is a safe prime if both p and q are primes.
// SafePrime will return error for any error returned by rand.Read or if bits < 2.
func SafePrime(random io.Reader, bits int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
