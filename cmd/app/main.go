package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"dep-comparer/internal/parser"
	"dep-comparer/internal/report/csv"
	"dep-comparer/internal/report/dot"
)

func main() {
	ctx := context.Background()
	var language string
	dotFlag := flag.Bool("dot", false, "Generate dot file")
	flag.StringVar(&language, "l", "", "Language to use")
	flag.Parse()
	listOfDepFiles := flag.Args()

	if len(listOfDepFiles) == 0 {
		slog.Info("no files to diff")
		os.Exit(0)
	}

	programLanguage, err := parser.GetLanguageTypeByName(language)
	if err != nil {
		slog.Error("could not determine language", "language", err)
		os.Exit(1)
	}

	p := parser.New()
	modules, err := p.Parse(ctx, programLanguage, listOfDepFiles)
	if err != nil {
		log.Fatal(err)
	}

	dep, devDep := parser.SummarizeModules(programLanguage, modules...)
	listOfDeps := parser.ConvertSummarizeDepToList(dep)
	var listOfDevDeps []string
	if devDep != nil {
		listOfDevDeps = parser.ConvertSummarizeDepToList(devDep)
	}
	path, err := csv.NewReport(language, listOfDeps, listOfDevDeps, csv.ByRows, modules...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\u001b[32m", path, "\u001b[0m") // Green color of report's name

	var pathOfDotFile string
	if *dotFlag {
		pathOfDotFile, err = dot.NewReport(modules...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\u001b[32m", pathOfDotFile, "\u001b[0m") // Green color of report's name
	}

}
