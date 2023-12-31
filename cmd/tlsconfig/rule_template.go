package main

import "text/template"

var generatedRuleTmpl = template.Must(template.New("generated").Parse(`
// New{{.Name}}TLSCheck creates a check for {{.Name}} TLS ciphers
// DO NOT EDIT - generated by tlsconfig tool
func New{{.Name}}TLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &insecureConfigTLS{
                MetaData: issue.MetaData{ID: id},
		requiredType: "crypto/tls.Config",
		MinVersion:   {{ .MinVersion }},
		MaxVersion:   {{ .MaxVersion }},
		goodCiphers: []string{
{{range $cipherName := .Ciphers }} "{{$cipherName}}",
{{end}}
		},
	}, []ast.Node{(*ast.CompositeLit)(nil), (*ast.AssignStmt)(nil)}
}
`))
