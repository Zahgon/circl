package test

import (
	"encoding"
	"testing"
)

// ReportError reports an error if got is different from want.
func ReportError(t testing.TB, got, want interface{}, inputs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// CheckOk fails the test if result == false.
func CheckOk(result bool, msg string, t testing.TB) { _ = "STUB: not implemented"; return }

// checkErr fails on error condition. mustFail indicates whether err is expected
// to be nil or not.
func checkErr(t testing.TB, err error, mustFail bool, msg string) {
	_ = "STUB: not implemented"
	return
}

// CheckNoErr fails if err !=nil. Print msg as an error message.
func CheckNoErr(t testing.TB, err error, msg string) { _ = "STUB: not implemented"; return }

// CheckIsErr fails if err ==nil. Print msg as an error message.
func CheckIsErr(t testing.TB, err error, msg string) { _ = "STUB: not implemented"; return }

// CheckPanic returns true if call to function 'f' caused panic.
func CheckPanic(f func()) error { _ = "STUB: not implemented"; return nil }

func CheckMarshal(
	t *testing.T,
	x, y interface {
		encoding.BinaryMarshaler
		encoding.BinaryUnmarshaler
	},
) {
	_ = "STUB: not implemented"
	return
}

// []byte but is encoded in hex for JSON
type HexBytes []byte

func (b HexBytes) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *HexBytes) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func gunzip(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Like os.ReadFile, but gunzip first.
func ReadGzip(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
