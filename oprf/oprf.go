// Package oprf provides Verifiable, Oblivious Pseudo-Random Functions.
//
// An Oblivious Pseudorandom Function (OPRFs) is a two-party protocol for
// computing the output of a PRF. One party (the server) holds the PRF secret
// key, and the other (the client) holds the PRF input.
//
// This package is compatible with the OPRF specification at RFC-9497 [1].
//
// # Protocol Overview
//
// This diagram shows the steps of the protocol that are common for all operation modes.
//
//	Client(info*)                               Server(sk, pk, info*)
//	=================================================================
//	finData, evalReq = Blind(input)
//
//	                            evalReq
//	                          ---------->
//
//	                            evaluation = Evaluate(evalReq, info*)
//
//	                           evaluation
//	                          <----------
//
//	output = Finalize(finData, evaluation, info*)
//
// # Operation Modes
//
// Each operation mode provides different properties to the PRF.
//
// Base Mode: Provides obliviousness to the PRF evaluation, i.e., it ensures
// that the server does not learn anything about the client's input and output
// during the Evaluation step.
//
// Verifiable Mode: Extends the Base mode allowing the client to verify that
// Server used the private key that corresponds to the public key.
//
// Partial Oblivious Mode: Extends the Verifiable mode by including shared
// public information to the PRF input.
//
// All three modes can perform batches of PRF evaluations, so passing an array
// of inputs will produce an array of outputs.
//
// # References
//
// [1] RFC-9497: https://www.rfc-editor.org/info/rfc9497
package oprf

import (
	"crypto"
	"errors"
	"hash"
	"io"

	"github.com/cloudflare/circl/group"
	"github.com/cloudflare/circl/zk/dleq"
)

const (
	version          = "OPRFV1-"
	finalizeDST      = "Finalize"
	hashToGroupDST   = "HashToGroup-"
	hashToScalarDST  = "HashToScalar-"
	deriveKeyPairDST = "DeriveKeyPair"
	infoLabel        = "Info"
)

type Mode = uint8

const (
	BaseMode             Mode = 0x00
	VerifiableMode       Mode = 0x01
	PartialObliviousMode Mode = 0x02
)

func isValidMode(m Mode) bool { _ = "STUB: not implemented"; return false }

type Suite interface {
	Identifier() string
	Group() group.Group
	Hash() crypto.Hash
	cannotBeImplementedExternally()
}

var (
	// SuiteRistretto255 represents the OPRF with Ristretto255 and SHA-512
	SuiteRistretto255 Suite = params{identifier: "ristretto255-SHA512", group: group.Ristretto255, hash: crypto.SHA512}
	// SuiteP256 represents the OPRF with P-256 and SHA-256.
	SuiteP256 Suite = params{identifier: "P256-SHA256", group: group.P256, hash: crypto.SHA256}
	// SuiteP384 represents the OPRF with P-384 and SHA-384.
	SuiteP384 Suite = params{identifier: "P384-SHA384", group: group.P384, hash: crypto.SHA384}
	// SuiteP521 represents the OPRF with P-521 and SHA-512.
	SuiteP521 Suite = params{identifier: "P521-SHA512", group: group.P521, hash: crypto.SHA512}
)

func GetSuite(identifier string) (Suite, error) { _ = "STUB: not implemented"; return *new(Suite), nil }

func NewClient(s Suite) Client { _ = "STUB: not implemented"; return *new(Client) }

func NewVerifiableClient(s Suite, server *PublicKey) VerifiableClient {
	_ = "STUB: not implemented"
	return *new(VerifiableClient)
}

func NewPartialObliviousClient(s Suite, server *PublicKey) PartialObliviousClient {
	_ = "STUB: not implemented"
	return *new(PartialObliviousClient)
}

func NewServer(s Suite, key *PrivateKey) Server { _ = "STUB: not implemented"; return *new(Server) }

func NewVerifiableServer(s Suite, key *PrivateKey) VerifiableServer {
	_ = "STUB: not implemented"
	return *new(VerifiableServer)
}

func NewPartialObliviousServer(s Suite, key *PrivateKey) PartialObliviousServer {
	_ = "STUB: not implemented"
	return *new(PartialObliviousServer)
}

type params struct {
	m          Mode
	group      group.Group
	hash       crypto.Hash
	identifier string
}

func (p params) cannotBeImplementedExternally() { _ = "STUB: not implemented"; return }

func (p params) String() string     { _ = "STUB: not implemented"; return "" }
func (p params) Group() group.Group { _ = "STUB: not implemented"; return *new(group.Group) }
func (p params) Hash() crypto.Hash  { _ = "STUB: not implemented"; return *new(crypto.Hash) }
func (p params) Identifier() string { _ = "STUB: not implemented"; return "" }

func (p params) getDST(name string) []byte { _ = "STUB: not implemented"; return nil }

func (p params) scalarFromInfo(info []byte) (group.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(group.Scalar), nil
}

func (p params) finalizeHash(h hash.Hash, input, info, element []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (p params) getDLEQParams() (out dleq.Params) {
	_ = "STUB: not implemented"
	return *new(dleq.Params)
}

func mustWrite(h io.Writer, bytes []byte) { _ = "STUB: not implemented"; return }

var (
	ErrInvalidSuite       = errors.New("oprf: invalid suite")
	ErrInvalidMode        = errors.New("oprf: invalid mode")
	ErrDeriveKeyPairError = errors.New("oprf: key pair derivation failed")
	ErrInvalidInput       = errors.New("oprf: invalid input")
	ErrInvalidSeed        = errors.New("oprf: invalid seed size")
	ErrInvalidInfo        = errors.New("oprf: invalid info")
	ErrInvalidProof       = errors.New("oprf: proof verification failed")
	ErrInverseZero        = errors.New("oprf: inverting a zero value")
	ErrNoKey              = errors.New("oprf: must provide a key")
)

type (
	Blind     = group.Scalar
	Blinded   = group.Element
	Evaluated = group.Element
)

// FinalizeData encapsulates data needed for Finalize step.
type FinalizeData struct {
	inputs  [][]byte
	blinds  []Blind
	evalReq *EvaluationRequest
}

// CopyBlinds copies the serialized blinds to use when deterministically
// invoking DeterministicBlind.
func (f FinalizeData) CopyBlinds() []Blind { _ = "STUB: not implemented"; return nil }

// EvaluationRequest contains the blinded elements to be evaluated by the Server.
type EvaluationRequest struct {
	Elements []Blinded
}

// Evaluation contains a list of elements produced during server's evaluation, and
// for verifiable modes it also includes a proof.
type Evaluation struct {
	Elements []Evaluated
	Proof    *dleq.Proof
}
