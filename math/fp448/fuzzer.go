//go:build gofuzz
// +build gofuzz

// How to run the fuzzer:
//
//	$ go get -u github.com/dvyukov/go-fuzz/go-fuzz
//	$ go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
//	$ go-fuzz-build -libfuzzer -func FuzzReduction -o lib.a
//	$ clang -fsanitize=fuzzer lib.a -o fu.exe
//	$ ./fu.exe
package fp448

// FuzzReduction is a fuzzer target for red64 function, which reduces t
// (112 bits) to a number t' (56 bits) congruent modulo p448.
func FuzzReduction(data []byte) int { _ = "STUB: not implemented"; return 0 }

// 2^448

// 2^448-1

// 2^224+1
