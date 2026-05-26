//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"math/big"
	"os"
	"path"
	"strings"
	"text/template"
)

type Fp struct {
	Name          string
	Bits          uint
	NumRootsUnity uint
	NumInverseInt uint
	Modulus       string
	Generator     string
}

func (f Fp) NumUint64() int { _ = "STUB: not implemented"; return 0 }

func (f Fp) NumUint8() int { _ = "STUB: not implemented"; return 0 }

func (f Fp) NeedsMasking() bool { _ = "STUB: not implemented"; return false }

func (f Fp) OrderConst() (s string) { _ = "STUB: not implemented"; return "" }

func (f Fp) OrderVar() string { _ = "STUB: not implemented"; return "" }

func (f Fp) Prime() *big.Int { _ = "STUB: not implemented"; return nil }

func (f Fp) Half() string { _ = "STUB: not implemented"; return "" }

func (f Fp) RSquare() string { _ = "STUB: not implemented"; return "" }

func (f Fp) RootsOfUnity() (s string) { _ = "STUB: not implemented"; return "" }

func (f Fp) InverseInt() (s string) { _ = "STUB: not implemented"; return "" }

func (f Fp) toMont(x, p *big.Int) { _ = "STUB: not implemented"; return }

func baseTwo64(v *big.Int) (d []*big.Int) { _ = "STUB: not implemented"; return nil }

func printDigits(n *big.Int) (s string) { _ = "STUB: not implemented"; return "" }

func main() {
	const TemplateWarning = "// Code generated from"

	fields := []Fp{
		{
			Name:          "Fp64",
			Bits:          64,
			NumRootsUnity: 32,
			NumInverseInt: 8,
			// Modulus: 2^32 * 4294967295 + 1
			Modulus: "ffffffff00000001",
			// Generator: 7^4294967295
			Generator: "185629dcda58878c",
		},
		{
			Name:          "Fp128",
			Bits:          128,
			NumRootsUnity: 66,
			NumInverseInt: 8,
			// Modulus: 2^66 * 4611686018427387897 + 1
			Modulus: "ffffffffffffffe40000000000000001",
			// Generator: 7^4611686018427387897
			Generator: "6d278fbf4f60228b1f9b2759c5109f06",
		},
	}

	for _, file := range []string{"fp_test", "fp", "vector", "poly"} {
		tName := "templates/" + file + ".go.tmpl"
		tl, err := template.
			New(path.Base(tName)).
			Funcs(template.FuncMap{"ToLower": strings.ToLower}).
			ParseFiles(tName)
		if err != nil {
			panic(err)
		}

		for _, f := range fields {
			buf := new(bytes.Buffer)
			err := tl.Execute(buf, f)
			if err != nil {
				panic(err)
			}

			code := buf.Bytes()
			code, err = format.Source(code)
			if err != nil {
				panic("error formating code")
			}

			res := string(code)
			offset := strings.Index(res, TemplateWarning)
			if offset == -1 {
				panic("Missing template warning")
			}

			folder := strings.ToLower(f.Name)
			fileName := file
			if file == "fp_test" {
				folder = "."
				fileName = strings.ToLower(f.Name) + "_test"
			}

			path := fmt.Sprintf("%v/%v.go", folder, fileName)
			err = os.WriteFile(path, []byte(res[offset:]), 0o600)
			if err != nil {
				panic(err)
			}
		}
	}
}
