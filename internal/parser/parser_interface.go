package parser

import (
	"context"
	"dep-comparer/internal/parser/types"
)

type LanguageParserInterface interface {
	Parse(ctx context.Context, listOfDepFiles []string) ([]*types.Dependency, error)
}
