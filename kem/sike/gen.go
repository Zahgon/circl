//go:build ignore
// +build ignore

// Autogenerates wrappers from templates to prevent too much duplicated code
// between the code for different modes.
package main

type Instance struct {
	Bits int
}

func (m Instance) Pkg() string { _ = "STUB: not implemented"; return "" }

func (m Instance) Name() string { _ = "STUB: not implemented"; return "" }

func (m Instance) Field() string { _ = "STUB: not implemented"; return "" }

var (
	Instances = []Instance{
		{Bits: 434},
		{Bits: 503},
		{Bits: 751},
	}
	TemplateWarning = "// Code generated from"
)

func main() {
	generatePackageFiles()
}

// Generates instance/sike.go from templates/pkg.templ.go
func generatePackageFiles() { _ = "STUB: not implemented"; return }

// Formating output code
