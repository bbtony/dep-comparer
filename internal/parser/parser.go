package parser

import (
	"context"
	"dep-comparer/internal/reader"

	"golang.org/x/mod/modfile"
	"golang.org/x/sync/errgroup"
)

const GoMod = "go.mod"

type DependencyVersion string

type DependencyPath string

type ModulePath string

type GoVersion string

type Module struct {
	ModulePath
	GoVersion
	Dependencies map[DependencyPath]DependencyVersion
}

func newModule() *Module {
	return &Module{
		Dependencies: make(map[DependencyPath]DependencyVersion),
	}
}

// Parse - main function of parsing files of dependencies
func Parse(
	ctx context.Context,
	listOfDepFiles []string,
	goroutineCount int,
) ([]*Module, error) {
	g, ctx := errgroup.WithContext(ctx)

	depParserRes := make(chan *Module, len(listOfDepFiles))

	limiter := make(chan struct{}, goroutineCount)
	defer close(limiter)

	modules := []*Module{}

	for _, el := range listOfDepFiles {
		g.Go(func() error {
			// read file
			data, err := reader.ReadFile(el)
			if err != nil {
				return err
			}
			// parse data
			module, err := ParseGoMod(data)
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

// ParseGoMod - parse data of Golang dependencies file
func ParseGoMod(data []byte) (*Module, error) {

	file, err := modfile.Parse(GoMod, data, nil)
	if err != nil {
		return nil, err
	}

	currentModule := newModule()
	currentModule.ModulePath = ModulePath(file.Module.Mod.Path)
	currentModule.GoVersion = GoVersion(file.Go.Version)
	for _, el := range file.Require {
		currentModule.Dependencies[DependencyPath(el.Mod.Path)] = DependencyVersion(el.Mod.Version)
	}

	return currentModule, nil
}

// SummarizeModules - prepare map of dependencies from all files
func SummarizeModules(modules ...*Module) map[DependencyPath]struct{} {
	res := make(map[DependencyPath]struct{})
	for _, el := range modules {
		for path := range el.Dependencies {
			if _, flag := res[path]; !flag {
				res[path] = struct{}{}
			}
		}
	}

	return res
}

// ConvertSummarizeDepToList -
func ConvertSummarizeDepToList(dependencies map[DependencyPath]struct{}) []string {
	l := make([]string, 0, len(dependencies))
	for path, _ := range dependencies {
		l = append(l, string(path))
	}

	return l
}
