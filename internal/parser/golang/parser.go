package golang

import (
	"dep-comparer/internal/parser/types"

	"golang.org/x/mod/modfile"
)

func newDependency() *types.Dependency {
	return &types.Dependency{
		Dependencies: make(map[types.DependencyPath]types.DependencyVersion),
	}
}

// ParseGoMod - parse data of Golang dependencies file
func ParseGoMod(nameOfFile string, data []byte) (*types.Dependency, error) {

	file, err := modfile.Parse(nameOfFile, data, nil)
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
