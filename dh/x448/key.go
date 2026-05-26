package x448

// Size is the length in bytes of a X448 key.
const Size = 56

// Key represents a X448 key.
type Key [Size]byte

func (k *Key) clamp(in *Key) *Key { _ = "STUB: not implemented"; return nil }

// isValidPubKey verifies if the public key is not a low-order point.
func (k *Key) isValidPubKey() bool { _ = "STUB: not implemented"; return false }

// KeyGen obtains a public key given a secret key.
func KeyGen(public, secret *Key) { _ = "STUB: not implemented"; return }

// Shared calculates Alice's shared key from Alice's secret key and Bob's
// public key returning true on success. A failure case happens when the public
// key is a low-order point, thus the shared key is all-zeros and the function
// returns false.
func Shared(shared, secret, public *Key) bool { _ = "STUB: not implemented"; return false }
