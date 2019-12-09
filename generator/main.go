package main

//go:generate go run generate_templates.go

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type generator struct {
	agent *agent

	ddlfile     string
	outdir      string
	packageName string
}

func main() {
	g := &generator{}

	app := kingpin.New("choria_generage", "Generates golang stubs to a remote agent")
	app.Arg("ddl", "Full path to the JSON DDL to use for input").Required().ExistingFileVar(&g.ddlfile)
	app.Arg("output", "Full path to the output directory").Required().ExistingDirVar(&g.outdir)
	app.Flag("package", "Name of the package to generate").StringVar(&g.packageName)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	err := g.generateClient()
	if err != nil {
		panic(err)
	}
}
