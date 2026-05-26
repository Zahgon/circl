//go:build ignore
// +build ignore

// Autogenerates wrappers from templates to prevent too much duplicated code
// between the code for different modes.
package main

import (
	"encoding/asn1"

	"github.com/cloudflare/circl/sign/internal/dilithium/params"
)

type Mode struct {
	Name          string
	K             int
	L             int
	Eta           int
	DoubleEtaBits int
	Omega         int
	Tau           int
	Gamma1Bits    int
	Gamma2        int
	TRSize        int
	CTildeSize    int
	Oid           asn1.ObjectIdentifier
}

func (m Mode) Pkg() string { _ = "STUB: not implemented"; return "" }

func (m Mode) PkgPath() string { _ = "STUB: not implemented"; return "" }

func (m Mode) Impl() string { _ = "STUB: not implemented"; return "" }

func (m Mode) Mode() string { _ = "STUB: not implemented"; return "" }

func (m Mode) NIST() bool { _ = "STUB: not implemented"; return false }

// https://csrc.nist.gov/Projects/computer-security-objects-register/algorithm-registration
func (m Mode) OidGo() string { _ = "STUB: not implemented"; return "" }

var (
	Modes = []Mode{
		{
			Name:          "Dilithium2",
			K:             4,
			L:             4,
			Eta:           2,
			DoubleEtaBits: 3,
			Omega:         80,
			Tau:           39,
			Gamma1Bits:    17,
			Gamma2:        (params.Q - 1) / 88,
			TRSize:        32,
			CTildeSize:    32,
		},
		{
			Name:          "Dilithium3",
			K:             6,
			L:             5,
			Eta:           4,
			DoubleEtaBits: 4,
			Omega:         55,
			Tau:           49,
			Gamma1Bits:    19,
			Gamma2:        (params.Q - 1) / 32,
			TRSize:        32,
			CTildeSize:    32,
		},
		{
			Name:          "Dilithium5",
			K:             8,
			L:             7,
			Eta:           2,
			DoubleEtaBits: 3,
			Omega:         75,
			Tau:           60,
			Gamma1Bits:    19,
			Gamma2:        (params.Q - 1) / 32,
			TRSize:        32,
			CTildeSize:    32,
		},
		{
			Name:          "ML-DSA-44",
			K:             4,
			L:             4,
			Eta:           2,
			DoubleEtaBits: 3,
			Omega:         80,
			Tau:           39,
			Gamma1Bits:    17,
			Gamma2:        (params.Q - 1) / 88,
			TRSize:        64,
			CTildeSize:    32,
			Oid:           asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 3, 17},
		},
		{
			Name:          "ML-DSA-65",
			K:             6,
			L:             5,
			Eta:           4,
			DoubleEtaBits: 4,
			Omega:         55,
			Tau:           49,
			Gamma1Bits:    19,
			Gamma2:        (params.Q - 1) / 32,
			TRSize:        64,
			CTildeSize:    48,
			Oid:           asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 3, 18},
		},
		{
			Name:          "ML-DSA-87",
			K:             8,
			L:             7,
			Eta:           2,
			DoubleEtaBits: 3,
			Omega:         75,
			Tau:           60,
			Gamma1Bits:    19,
			Gamma2:        (params.Q - 1) / 32,
			TRSize:        64,
			CTildeSize:    64,
			Oid:           asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 3, 19},
		},
	}
	TemplateWarning = "// Code generated from"
)

func main() {
	generateModePackageFiles()
	generateACVPTest()
	generateParamsFiles()
	generateSourceFiles()
}

// Generates modeX/internal/params.go from templates/params.templ.go
func generateParamsFiles() { _ = "STUB: not implemented"; return }

// Formating output code

// Generates modeX/dilithium.go from templates/pkg.templ.go
func generateModePackageFiles() { _ = "STUB: not implemented"; return }

// Generates modeX/dilithium.go from templates/pkg.templ.go
func generateACVPTest() { _ = "STUB: not implemented"; return }

// Copies mode3 source files to other modes
func generateSourceFiles() { _ = "STUB: not implemented"; return }

// Ignore mode specific files.

// Read files

// Go over modes
