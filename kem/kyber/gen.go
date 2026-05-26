//go:build ignore
// +build ignore

// Autogenerates wrappers from templates to prevent too much duplicated code
// between the code for different modes.
package main

type Instance struct {
	Name string
}

func (m Instance) KemName() string { _ = "STUB: not implemented"; return "" }

func (m Instance) NIST() bool { _ = "STUB: not implemented"; return false }

func (m Instance) PkePkg() string { _ = "STUB: not implemented"; return "" }

func (m Instance) Pkg() string { _ = "STUB: not implemented"; return "" }

func (m Instance) PkgPath() string { _ = "STUB: not implemented"; return "" }

var (
	Instances = []Instance{
		{Name: "Kyber512"},
		{Name: "Kyber768"},
		{Name: "Kyber1024"},
		{Name: "ML-KEM-512"},
		{Name: "ML-KEM-768"},
		{Name: "ML-KEM-1024"},
	}
	TemplateWarning = "// Code generated from"
)

func main() {
	generatePackageFiles()
}

// Generates instance/kyber.go from templates/pkg.templ.go
func generatePackageFiles() { _ = "STUB: not implemented"; return }

// Formating output code
