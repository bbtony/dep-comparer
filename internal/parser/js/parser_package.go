package js

import (
	"context"
	"dep-comparer/internal/parser/types"
	"encoding/json"
)

type parser struct{}

type packageJson struct {
	Name            string            `json:"Name,omitempty"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Parse(ctx context.Context, name string, data []byte) (*types.Dependency, error) {
	var pJSON packageJson
	err := json.Unmarshal(data, &pJSON)
	if err != nil {
		return nil, err
	}

	currentDep := newDependency()
	currentDep.DependencyPath = types.DependencyPath(pJSON.Name)
	//if phpVersion, ok := pJSON.Dependencies["js"]; ok {
	//	currentDep.Version = types.Version(phpVersion)
	//}

	for key, version := range pJSON.Dependencies {
		currentDep.Dependencies[types.DependencyPath(key)] = types.DependencyVersion(version)
	}

	for key, dev := range pJSON.DevDependencies {
		currentDep.DevDependencies[types.DependencyPath(key)] = types.DependencyVersion(dev)
	}

	return currentDep, nil
}

func newDependency() *types.Dependency {
	return &types.Dependency{
		Dependencies:    make(map[types.DependencyPath]types.DependencyVersion),
		DevDependencies: make(map[types.DependencyPath]types.DependencyVersion),
	}
}
