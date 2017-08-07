package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ChimeraCoder/gojson"
)

var (
	clientString = `{{.StructCode}}

var (
    {{.Route}}Endpoint = "/{{.Route}}"
    {{.Route}}Method = "{{.Method}}"
)

// Do wraps and returns the {{.Method}}'ing of {{.StructName}} to /{{.Route}}
func (o {{.StructName}}) Do() (i interface{}, err error) {
    return Client.Do({{.Route}}Method, {{.Route}}Endpoint, o)
}

`
)

type client struct {
	StructName string
	StructCode string
	Route      string
	Method     string
}

func main() {
	err := filepath.Walk("./schema", Walker)
	if err != nil {
		panic(err)
	}
}

func Walker(path string, info os.FileInfo, _ error) (err error) {
	if info.IsDir() {
		return
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	structName := InferStructName(path)
	fileName := InferFileName(path)

	output, err := gojson.Generate(f, gojson.ParseJson, structName, "tyk", []string{"json"}, false)
	if err != nil {
		panic(err)
	}

	r, m := pathElements(path) // duplication!!!

	c := client{
		StructName: structName,
		StructCode: string(output),
		Route:      r,
		Method:     m,
	}

	tmpl, err := template.New("test").Parse(clientString)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	err = tmpl.Execute(outputFile, c)
	if err != nil {
		panic(err)
	}

	return
}

func InferStructName(p string) string {
	r, p := pathElements(p)

	return fmt.Sprintf("%s%s",
		strings.Title(r),
		strings.Title(methodToPurpose(p)),
	)
}

func InferFileName(p string) string {
	r, p := pathElements(p)

	return fmt.Sprintf("client/%s_%s.go",
		r,
		methodToPurpose(p),
	)
}

func pathElements(p string) (route, purpose string) {
	pElems := strings.Split(p, "/")

	route = pElems[1]
	purpose = strings.Split(pElems[len(pElems)-1], ".")[0]

	return
}

func methodToPurpose(s string) string {
	switch s {
	case "post":
		return "create"
	}

	return s
}
