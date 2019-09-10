package main

//go:generate go run generate_templates.go

import (
	// "fmt"
	"encoding/base64"
	"fmt"
	"github.com/choria-io/go-choria/choria"
	addl "github.com/choria-io/mcorpc-agent-provider/mcorpc/ddl/agent"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type Agent struct {
	Package string
	DDL     *addl.DDL
	RawDDL  string // raw text of the JSON DDL file
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"base64encode": func(v string) string { return base64.StdEncoding.EncodeToString([]byte(v)) },
		"Capitalize":   strings.Title,
		"ToLower":      strings.ToLower,
		"SnakeToCamel": func(v string) string {
			parts := strings.Split(v, "_")
			out := []string{}
			for _, s := range parts {
				out = append(out, strings.Title(s))
			}

			return strings.Join(out, "")
		},
		"ChoriaTypeToGoType": func(v string) string {
			switch v {
			case "string", "list":
				return "string"
			case "integer":
				return "int64"
			case "number", "float":
				return "float64"
			case "boolean":
				return "bool"
			case "hash":
				return "map[string]interface{}"
			case "array":
				return "[]interface{}"
			default:
				return "interface{}"
			}
		},
		"ChoriaTypeToValOfType": func(v string) string {
			switch v {
			case "string", "list":
				return "val.(string)"
			case "integer":
				return "val.(int64)"
			case "number", "float":
				return "val.(float64)"
			case "boolean":
				return "val.(bool)"
			case "hash":
				return "val.(map[string]interface{}"
			case "array":
				return "val.([]interface{})"
			default:
				return "val.(interface{})"
			}
		},
	}
}

func writeActions(agent *Agent, parent string) {
	code, err := base64.StdEncoding.DecodeString(templates["action"])
	panicIfError(err)

	type action struct {
		Agent             *Agent
		Package           string
		AgentName         string
		ActionName        string
		ActionDescription string
		OutputNames       []string
		InputNames        []string
		Inputs            map[string]*addl.ActionInputItem
		Outputs           map[string]*addl.ActionOutputItem
	}

	for _, actname := range agent.DDL.ActionNames() {
		actint, err := agent.DDL.ActionInterface(actname)
		panicIfError(err)

		out, err := os.Create(path.Join(parent, fmt.Sprintf("action_%s.go", actint.Name)))
		panicIfError(err)

		act := &action{
			Agent:             agent,
			Package:           agent.Package,
			AgentName:         agent.DDL.Metadata.Name,
			ActionName:        actint.Name,
			ActionDescription: actint.Description,
			InputNames:        actint.InputNames(),
			OutputNames:       actint.OutputNames(),
			Inputs:            actint.Input,
			Outputs:           actint.Output,
		}

		tpl := template.Must(template.New(actint.Name).Funcs(funcMap()).Parse(string(code)))
		err = tpl.Execute(out, act)
		panicIfError(err)
	}
}

func writeBasics(agent *Agent, parent string) {
	for _, t := range []string{"resultdetails", "requestor", "ddl", "discover", "rpcoptions", "client", "initoptions", "logging"} {
		out, err := os.Create(path.Join(parent, t+".go"))
		panicIfError(err)

		code, err := base64.StdEncoding.DecodeString(templates[t])
		panicIfError(err)

		tpl := template.Must(template.New(t).Funcs(funcMap()).Parse(string(code)))

		err = tpl.Execute(out, agent)
		panicIfError(err)
	}
}

func main() {
	ddlfile := ""
	outdir := ""

	app := kingpin.New("choria_generage", "Generates golang stubs to a remote agent")
	app.Arg("ddl", "Full path to the JSON DDL to use for input").ExistingFileVar(&ddlfile)
	app.Arg("output", "Full path to the output directory").ExistingDirVar(&outdir)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	_, err := choria.New(choria.UserConfig())
	panicIfError(err)

	agent := &Agent{}

	agent.DDL, err = addl.New(ddlfile)
	panicIfError(err)

	raw, err := ioutil.ReadFile(ddlfile)
	panicIfError(err)
	agent.RawDDL = string(raw)

	agent.Package = strings.ToLower(agent.DDL.Metadata.Name) + "client"

	if agent.DDL == nil {
		panic("could not find any DDL")
	}

	writeActions(agent, outdir)
	writeBasics(agent, outdir)
}
