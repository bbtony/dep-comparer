package php

import (
	"context"
	"dep-comparer/internal/parser"
	"dep-comparer/internal/parser/types"
	"dep-comparer/internal/reader"
	"golang.org/x/sync/errgroup"
)

// Parse - main function of parsing files of dependencies
func Parse(
	ctx context.Context,
	listOfDepFiles []string,
) ([]*types.Dependency, error) {
	g, ctx := errgroup.WithContext(ctx)

	depParserRes := make(chan *types.Dependency, len(listOfDepFiles))

	modules := []*types.Dependency{}

	for _, depFile := range listOfDepFiles {
		g.Go(func() error {
			// read file
			data, err := reader.ReadFile(depFile)
			if err != nil {
				return err
			}

			// parse data
			module, err := parsePHPComposerJSON(parser.GetNameOfDependencyFile(depFile), data)
			if err != nil {
				return err
			}

			depParserRes <- module

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	close(depParserRes)

	for module := range depParserRes {
		modules = append(modules, module)
	}

	return modules, nil
}

func parsePHPComposerJSON(name string, data []byte) (*types.Dependency, error) {
	return nil, nil
}
