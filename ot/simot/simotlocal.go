package simot

import (
	"github.com/cloudflare/circl/group"
)

const keyLength = 16

// AES GCM encryption, we don't need to pad because our input is fixed length
// Need to use authenticated encryption to defend against tampering on ciphertext
// Input: key, plaintext message
// Output: ciphertext
func aesEncGCM(key, plaintext []byte) []byte { _ = "STUB: not implemented"; return nil }

// AES GCM decryption
// Input: key, ciphertext message
// Output: plaintext
func aesDecGCM(key, ciphertext []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Initialization

// Input: myGroup, the group we operate in
// Input: m0, m1 the 2 message of the sender
// Input: index, the index of this SimOT
// Output: A = [a]G, a the sender randomness
func (sender *Sender) InitSender(myGroup group.Group, m0, m1 []byte, index int) group.Element {
	_ = "STUB: not implemented"
	return *new(group.Element)
}

// Round 1

// ---- sender should send A to receiver ----

// Input: myGroup, the group we operate in
// Input: choice, the receiver choice bit
// Input: index, the index of this SimOT
// Input: A, from sender
// Output: B = [b]G if c == 0, B = A+[b]G if c == 1 (Implementation in constant time). b, the receiver randomness
func (receiver *Receiver) Round1Receiver(myGroup group.Group, choice int, index int, A group.Element) group.Element {
	_ = "STUB: not implemented"
	return *new(group.Element)
}

// Round 2

// ---- receiver should send B to sender ----

// Input: B from the receiver
// Output: e0, e1, encryption of m0 and m1 under key k0, k1
func (sender *Sender) Round2Sender(B group.Element) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hash the whole transcript A|B|...

// Round 3

// ---- sender should send e0, e1 to receiver ----

// Input: e0, e1: encryption of m0 and m1 from the sender
// Input: choice, choice bit of receiver
// Choose e0 or e1 based on choice bit in constant time
func (receiver *Receiver) Round3Receiver(e0, e1 []byte, choice int) error {
	_ = "STUB: not implemented"
	return nil
}

// If c == 1, copy e1

// If c == 0, copy e0

// Hash the whole transcript so far

// kR, decryption key of mc

func (receiver *Receiver) Returnmc() []byte { _ = "STUB: not implemented"; return nil }

func (sender *Sender) Returne0e1() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (sender *Sender) Returnm0m1() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }
