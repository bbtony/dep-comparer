package parser

import (
	"context"
	"dep-comparer/internal/parser/golang"
	"dep-comparer/internal/parser/js"
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

type LanguageParser interface {
	Parse(ctx context.Context, nameOfFile string, data []byte) (*types.Dependency, error)
}

type Parser struct {
	LanguageType types.Language
	LanguageParser
}

func New(LanguageType types.Language) (*Parser, error) {
	var langParser LanguageParser

	switch LanguageType {
	case Golang:
		langParser = golang.NewParser()
	case PHP:
		langParser = php.NewParser()
	case JS:
		langParser = js.NewParser()
	default:
		return nil, fmt.Errorf("this is not a supported programming language")
	}

	return &Parser{
		LanguageType:   LanguageType,
		LanguageParser: langParser,
	}, nil
}

func (p *Parser) Parse(ctx context.Context, listOfDepFiles []string) ([]*types.Dependency, error) {
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

			var module *types.Dependency
			module, err = p.LanguageParser.Parse(ctx, depFile, data)
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

func GetLanguageNameByType(t types.Language) (string, error) {
	switch t {
	case Golang:
		return "go", nil
	case PHP:
		return "php", nil
	case JS:
		return "js", nil
	default:
		return "", fmt.Errorf("unknown type")
	}
}
