package golang

import (
	"context"
	"dep-comparer/internal/parser/types"
	"dep-comparer/internal/reader"
	"golang.org/x/mod/modfile"
	"golang.org/x/sync/errgroup"
)

const GoMod = "go.mod"

func newDependency() *types.Dependency {
	return &types.Dependency{
		Dependencies: make(map[types.DependencyPath]types.DependencyVersion),
	}
}

// Parse - main function of parsing files of dependencies
func Parse(
	ctx context.Context,
	listOfDepFiles []string,
) ([]*types.Dependency, error) {
	g, ctx := errgroup.WithContext(ctx)

	depParserRes := make(chan *types.Dependency, len(listOfDepFiles))

	modules := []*types.Dependency{}

	for _, el := range listOfDepFiles {
		g.Go(func() error {
			// read file
			data, err := reader.ReadFile(el)
			if err != nil {
				return err
			}
			// parse data
			module, err := parseGoMod(data)
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

// parseGoMod - parse data of Golang dependencies file
func parseGoMod(data []byte) (*types.Dependency, error) {

	file, err := modfile.Parse(GoMod, data, nil)
	if err != nil {
		return nil, err
	}

	currentDep := newDependency()
	currentDep.DependencyPath = types.DependencyPath(file.Module.Mod.Path)
	currentDep.Version = types.Version(file.Go.Version)
	for _, el := range file.Require {
		currentDep.Dependencies[types.DependencyPath(el.Mod.Path)] = types.DependencyVersion(el.Mod.Version)
	}

	return currentDep, nil
}
