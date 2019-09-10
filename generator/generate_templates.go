// +build ignore

package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type Template struct {
	Name string
	Body string
}

type Templates []Template

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

var templateWrapperTemplate = `
package main

var templates = map[string]string{
{{- range . }}
	"{{ .Name }}": "{{ .Body }}",
{{- end }}
}
`

func main() {
	files, err := ioutil.ReadDir("templates")
	panicIfErr(err)

	templates := Templates{}

	for _, file := range files {
		fname := strings.TrimSuffix(file.Name(), path.Ext(file.Name()))
		templ, err := ioutil.ReadFile(path.Join("templates", file.Name()))
		panicIfErr(err)

		templates = append(templates, Template{Name: fname, Body: base64.StdEncoding.EncodeToString(templ)})
	}

	t, err := template.New("templates").Parse(templateWrapperTemplate)
	panicIfErr(err)

	out, err := os.Create("templates.go")
	panicIfErr(err)
	defer out.Close()

	t.Execute(out, templates)
}
