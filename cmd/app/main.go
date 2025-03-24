package main

import (
	"context"
	"dep-comparer/internal/parser"
	"dep-comparer/internal/report/csv"
	"dep-comparer/internal/report/dot"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()

	dotFlag := flag.Bool("dot", false, "Generate dot file")
	flag.Parse()
	listOfDepFiles := flag.Args()

	if len(listOfDepFiles) == 0 {
		slog.Info("no files to diff")
		os.Exit(0)
	}

	modules, err := parser.Parse(ctx, listOfDepFiles, 5) // TODO: 5 maybe customized by params in the future
	if err != nil {
		log.Fatal(err)
	}

	listOfDeps := parser.ConvertSummarizeDepToList(parser.SummarizeModules(modules...))
	path, err := csv.NewReport(listOfDeps, csv.ByRows, modules...)
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
