package curve4q

// Size is the size in bytes of keys.
const Size = 32

// Key represents a public or private key of FourQ.
type Key [Size]byte

// KeyGen calculates a public key k from a secret key.
func KeyGen(public, secret *Key) { _ = "STUB: not implemented"; return }

// Shared calculates a shared key k from Alice's secret and Bob's public key.
// Returns true on success.
func Shared(shared, secret, public *Key) bool { _ = "STUB: not implemented"; return false }
