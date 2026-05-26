package sidh

import (
	"io"

	"github.com/cloudflare/circl/dh/sidh/internal/common"
)

// I keep it bool in order to be able to apply logical NOT.
//
// Deprecated: not cryptographically secure.
type KeyVariant uint

// Base type for public and private key. Used mainly to carry domain
// parameters.
type key struct {
	// Domain parameters of the algorithm to be used with a key
	params *common.SidhParams
	// Flag indicates whether corresponds to 2-, 3-torsion group or SIKE
	keyVariant KeyVariant
}

// Defines operations on public key
//
// Deprecated: not cryptographically secure.
type PublicKey struct {
	key
	// x-coordinates of P,Q,P-Q in this exact order
	affine3Pt [3]common.Fp2
}

// Defines operations on private key
//
// Deprecated: not cryptographically secure.
type PrivateKey struct {
	key
	// Secret key
	Scalar []byte
	// Used only by KEM
	S []byte
}

// Identifiers correspond to the bitlength of the prime field characteristic.
const (
	Fp434 = common.Fp434
	Fp503 = common.Fp503
	Fp751 = common.Fp751
)

const (
	// First 2 bits identify SIDH variant third bit indicates
	// whether key is a SIKE variant (set) or SIDH (not set)

	// 001 - SIDH: corresponds to 2-torsion group
	KeyVariantSidhA KeyVariant = 1 << 0
	// 010 - SIDH: corresponds to 3-torsion group
	KeyVariantSidhB = 1 << 1
	// 110 - SIKE
	KeyVariantSike = 1<<2 | KeyVariantSidhB
)

// Accessor to key variant.
func (key *key) Variant() KeyVariant {
	_ = "STUB: not implemented"
	return *

	// NewPublicKey initializes public key.
	// Usage of this function guarantees that the object is correctly initialized.
	//
	// Deprecated: not cryptographically secure.
	new(KeyVariant)
}

func NewPublicKey(id uint8, v KeyVariant) *PublicKey { _ = "STUB: not implemented"; return nil }

// Import clears content of the public key currently stored in the structure
// and imports key stored in the byte string. Returns error in case byte string
// size is wrong. Doesn't perform any validation.
func (pub *PublicKey) Import(input []byte) error { _ = "STUB: not implemented"; return nil }

// Exports currently stored key. In case structure hasn't been filled with key data
// returned byte string is filled with zeros.
func (pub *PublicKey) Export(out []byte) { _ = "STUB: not implemented"; return }

// Size returns size of the public key in bytes.
func (pub *PublicKey) Size() int { _ = "STUB: not implemented"; return 0 }

// NewPrivateKey initializes private key.
// Usage of this function guarantees that the object is correctly initialized.
//
// Deprecated: not cryptographically secure.
func NewPrivateKey(id uint8, v KeyVariant) *PrivateKey { _ = "STUB: not implemented"; return nil }

// Exports currently stored key. In case structure hasn't been filled with key data
// returned byte string is filled with zeros.
func (prv *PrivateKey) Export(out []byte) { _ = "STUB: not implemented"; return }

// Size returns size of the private key in bytes.
func (prv *PrivateKey) Size() int { _ = "STUB: not implemented"; return 0 }

// Size returns size of the shared secret.
func (prv *PrivateKey) SharedSecretSize() int { _ = "STUB: not implemented"; return 0 }

// Import clears content of the private key currently stored in the structure
// and imports key from octet string. In case of SIKE, the random value 'S'
// must be prepended to the value of actual private key (see SIKE spec for details).
// Function doesn't import public key value to PrivateKey object.
func (prv *PrivateKey) Import(input []byte) error { _ = "STUB: not implemented"; return nil }

// Generates random private key for SIDH or SIKE. Generated value is
// formed as little-endian integer from key-space <2^(e2-1)..2^e2 - 1>
// for KeyVariant_A or <2^(s-1)..2^s - 1>, where s = floor(log_2(3^e3)),
// for KeyVariant_B.
//
// Returns error in case user provided RNG fails.
func (prv *PrivateKey) Generate(rand io.Reader) error { _ = "STUB: not implemented"; return nil }

// Private key generation takes advantage of the fact that keyspace for secret
// key is (0, 2^x - 1), for some positive value of 'x' (see SIKE, 1.3.8).
// It means that all bytes in the secret key, but the last one, can take any
// value between <0x00,0xFF>. Similarly for the last byte, but generation
// needs to chop off some bits, to make sure generated value is an element of
// a key-space.

// Make sure scalar is SecretBitLen long. SIKE spec says that key
// space starts from 0, but I'm not comfortable with having low
// value scalars used for private keys. It is still secure as per
// table 5.1 in [SIKE].

// Generates public key.
func (prv *PrivateKey) GeneratePublicKey(pub *PublicKey) { _ = "STUB: not implemented"; return }

// Computes a SIDH shared secret. Function requires that pub has different
// KeyVariant than prv. Length of returned output is 2*ceil(log_2 P)/8),
// where P is a prime defining finite field.
//
// Caller must make sure key SIDH key pair is not used more than once.
func (prv *PrivateKey) DeriveSecret(ss []byte, pub *PublicKey) { _ = "STUB: not implemented"; return }
