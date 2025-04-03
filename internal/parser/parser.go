package parser

import (
	"context"
	"dep-comparer/internal/parser/golang"
	"dep-comparer/internal/parser/php"
	"dep-comparer/internal/parser/types"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"strings"
)

const (
	Golang types.Language = iota
	PHP
	JS
)

type LanguageParserInterface interface {
	Parse(ctx context.Context, listOfDepFiles []string, LanguageType types.Language) ([]*types.Dependency, error)
}
type parser struct{}

func New() *parser {
	return &parser{}
}

func (p *parser) Parse(ctx context.Context, listOfDepFiles []string, LanguageType types.Language) ([]*types.Dependency, error) {
	g, ctx := errgroup.WithContext(ctx)

	depParserRes := make(chan *types.Dependency, len(listOfDepFiles))

	modules := []*types.Dependency{}

	for _, depFile := range listOfDepFiles {
		g.Go(func() error {
			// read file
			data, err := os.ReadFile(depFile)
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}

			var module *types.Dependency
			// parse data
			switch LanguageType {
			case Golang:
				module, err = golang.Parse(ctx, GetNameOfDependencyFile(depFile), data)
				if err != nil {
					return err
				}
			case PHP:

				module, err = php.Parse(ctx, GetNameOfDependencyFile(depFile), data)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("this is not a supported programming language")
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

func GetLanguageTypeByName(nameofProgrammingLanguage string) (types.Language, error) {
	switch strings.ToLower(nameofProgrammingLanguage) {
	case "go", "golang":
		return Golang, nil
	case "php":
		return PHP, nil
	case "java script", "javascript", "js", "java-script":
		return JS, nil
	default:
		return types.Language(0), fmt.Errorf("unknown language")
	}
}
