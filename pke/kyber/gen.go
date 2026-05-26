//go:build ignore
// +build ignore

// Autogenerates wrappers from templates to prevent too much duplicated code
// between the code for different modes.
package main

type Instance struct {
	Name           string
	K              int
	Eta1           int
	CiphertextSize int
	DU             int
	DV             int
}

func (m Instance) Pkg() string { _ = "STUB: not implemented"; return "" }

func (m Instance) Impl() string { _ = "STUB: not implemented"; return "" }

var (
	Instances = []Instance{
		{
			Name:           "Kyber512",
			Eta1:           3,
			K:              2,
			CiphertextSize: 768,
			DU:             10,
			DV:             4,
		},
		{
			Name:           "Kyber768",
			Eta1:           2,
			K:              3,
			CiphertextSize: 1088,
			DU:             10,
			DV:             4,
		},
		{
			Name:           "Kyber1024",
			Eta1:           2,
			K:              4,
			CiphertextSize: 1568,
			DU:             11,
			DV:             5,
		},
	}
	TemplateWarning = "// Code generated from"
)

func main() {
	generatePackageFiles()
	generateParamsFiles()
	generateSourceFiles()
}

// Generates instance/internal/params.go from templates/params.templ.go
func generateParamsFiles() { _ = "STUB: not implemented"; return }

// Formating output code

// Generates instance/kyber.go from templates/pkg.templ.go
func generatePackageFiles() { _ = "STUB: not implemented"; return }

// Copies kyber512 source files to other modes
func generateSourceFiles() { _ = "STUB: not implemented"; return }

// Ignore mode specific files.

// Read files

// Go over modes
