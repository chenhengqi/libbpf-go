package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

const codeTmpl = `package libbpf

// #include <bpf/{{.Header}}>
import "C"

{{range $API := .APIs}}
	{{$API}}
{{end}}
`

type codeDef struct {
	Header string
	APIs   []string
}

func gen(defs [][]byte, srcFile, dest string) {
	destDir, err := filepath.Abs(dest)
	if err != nil {
		panic(err)
	}
	file := strings.TrimSuffix(srcFile, filepath.Ext(srcFile))
	destFile := filepath.Join(destDir, file+".go")
	codeDef := &codeDef{
		Header: srcFile,
	}
	for _, def := range defs {
		codeDef.APIs = append(codeDef.APIs, string(def))
	}

	t := template.Must(template.New("code").Parse(codeTmpl))
	buf := bytes.Buffer{}
	err = t.Execute(&buf, &codeDef)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(destFile, buf.Bytes(), 0644)
}
