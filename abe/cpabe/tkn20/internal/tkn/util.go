package tkn

import (
	"errors"

	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

var gtBaseVal *pairing.Gt

func init() {
	// This should really be a constant, but what can I do?
	g1 := pairing.G1Generator()
	g2 := pairing.G2Generator()
	gtBaseVal = pairing.Pair(g1, g2)
}

func ToScalar(n int) *pairing.Scalar { _ = "STUB: not implemented"; return nil }

func HashStringToScalar(key []byte, value string) *pairing.Scalar {
	_ = "STUB: not implemented"
	return nil
}

func appendLen16Prefixed(a []byte, b []byte) []byte { _ = "STUB: not implemented"; return nil }

func removeLen16Prefixed(data []byte) (next []byte, remainder []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var (
	appendLenPrefixed = appendLen16Prefixed
	removeLenPrefixed = removeLen16Prefixed
)

func appendLen32Prefixed(a []byte, b []byte) []byte { _ = "STUB: not implemented"; return nil }

func removeLen32Prefixed(data []byte) (next []byte, remainder []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func marshalBinarySortedMapMatrixG1(m map[string]*matrixG1) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalBinarySortedMapAttribute(m map[string]Attribute) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errBadMatrixSize       = errors.New("matrix inputs do not conform")
	errMatrixNonInvertible = errors.New("matrix has no inverse")
)
