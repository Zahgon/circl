package keys

import (
	"crypto/rsa"
	"math/big"
)

// BigPublicKey is the same as an rsa.PublicKey struct, except the public
// key is represented as a big integer as opposed to an int. For the partially
// blind scheme, this is required since the public key will typically be
// any value in the RSA group.
type BigPublicKey struct {
	N *big.Int
	E *big.Int
}

// Size returns the size of the public key.
func (pub *BigPublicKey) Size() int { _ = "STUB: not implemented"; return 0 }

// Marshal encodes the public key exponent (e).
func (pub *BigPublicKey) Marshal() []byte { _ = "STUB: not implemented"; return nil }

// NewBigPublicKey creates a BigPublicKey from a rsa.PublicKey.
func NewBigPublicKey(pk *rsa.PublicKey) *BigPublicKey { _ = "STUB: not implemented"; return nil }

// CustomPublicKey is similar to rsa.PrivateKey, containing information needed
// for a private key used in the partially blind signature protocol.
type BigPrivateKey struct {
	Pk *BigPublicKey
	D  *big.Int
	P  *big.Int
	Q  *big.Int
}

// NewBigPrivateKey creates a BigPrivateKey from a rsa.PrivateKey.
func NewBigPrivateKey(sk *rsa.PrivateKey) *BigPrivateKey { _ = "STUB: not implemented"; return nil }
