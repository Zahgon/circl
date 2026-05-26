package oprf

import (
	"github.com/cloudflare/circl/group"
)

type client struct{ params }

type Client struct {
	client
}

type VerifiableClient struct {
	client
	pkS *PublicKey
}

type PartialObliviousClient struct {
	client
	pkS *PublicKey
}

func (c client) Blind(inputs [][]byte) (*FinalizeData, *EvaluationRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c client) DeterministicBlind(inputs [][]byte, blinds []Blind) (*FinalizeData, *EvaluationRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c client) blind(inputs [][]byte, blinds []Blind) (*FinalizeData, *EvaluationRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c client) unblind(serUnblindeds [][]byte, blindeds []group.Element, blind []Blind) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c client) validate(f *FinalizeData, e *Evaluation) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c client) finalize(f *FinalizeData, e *Evaluation, info []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Finalize(f *FinalizeData, e *Evaluation) (outputs [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c VerifiableClient) Finalize(f *FinalizeData, e *Evaluation) (outputs [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c PartialObliviousClient) Finalize(f *FinalizeData, e *Evaluation, info []byte) (outputs [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c PartialObliviousClient) pointFromInfo(info []byte) (group.Element, error) {
	_ = "STUB: not implemented"
	return *new(group.Element), nil
}
