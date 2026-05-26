package oprf

import (
	"github.com/cloudflare/circl/group"
)

type server struct {
	params
	privateKey *PrivateKey
}

type Server struct{ server }

type VerifiableServer struct{ server }

type PartialObliviousServer struct{ server }

func (s server) PublicKey() *PublicKey { _ = "STUB: not implemented"; return nil }

func (s server) evaluate(elements []Blinded, secret Blind) []Evaluated {
	_ = "STUB: not implemented"
	return nil
}

func (s Server) Evaluate(req *EvaluationRequest) (*Evaluation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s VerifiableServer) Evaluate(req *EvaluationRequest) (*Evaluation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s PartialObliviousServer) Evaluate(req *EvaluationRequest, info []byte) (*Evaluation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s server) secretFromInfo(info []byte) (t, tInv group.Scalar, err error) {
	_ = "STUB: not implemented"
	return *new(group.Scalar), *new(group.Scalar), nil
}

func (s server) fullEvaluate(input, info []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s Server) FullEvaluate(input []byte) (output []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s VerifiableServer) FullEvaluate(input []byte) (output []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s PartialObliviousServer) FullEvaluate(input, info []byte) (output []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s server) verifyFinalize(input, info, expectedOutput []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Server) VerifyFinalize(input, expectedOutput []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (s VerifiableServer) VerifyFinalize(input, expectedOutput []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (s PartialObliviousServer) VerifyFinalize(input, info, expectedOutput []byte) bool {
	_ = "STUB: not implemented"
	return false
}
