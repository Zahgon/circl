// Package dl provides a Schnorr NIZK discrete-log proof.
//
// This package implements a Schnorr NIZK discrete-log proof obtained from the
// interactive Schnorr identification scheme through a Fiat-Shamir transformation.
//
// Given (k,G,kG) the Prove function returns a Proof struct attesting that
// kG = [k]G, which can be validated using the Verify function.
//
// The userID label is a unique identifier for the prover.
//
// The otherInfo label is defined to allow flexible inclusion of contextual
// information in the Schnorr NIZK proof.
// The otherInfo is also used as a domain separation tag (dst) for the hash
// to scalar function.
//
// Reference: https://datatracker.ietf.org/doc/html/rfc8235
package dl

import (
	"io"

	"github.com/cloudflare/circl/group"
)

type Proof struct {
	V group.Element
	R group.Scalar
}

func calcChallenge(myGroup group.Group, G, V, A group.Element, userID, otherInfo []byte) group.Scalar {
	_ = "STUB: not implemented"
	// Hash transcript (G | V | A | UserID | OtherInfo) to get the random coin.
	return *new(group.Scalar)
}

// Prove returns a proof attesting that kG = [k]G.
func Prove(myGroup group.Group, G, kG group.Element, k group.Scalar, userID, otherInfo []byte, rnd io.Reader) Proof {
	_ = "STUB: not implemented"
	return *new(Proof)
}

// Verify checks whether the proof attests that kG = [k]G.
func Verify(myGroup group.Group, G, kG group.Element, p Proof, userID, otherInfo []byte) bool {
	_ = "STUB: not implemented"
	return false
}
