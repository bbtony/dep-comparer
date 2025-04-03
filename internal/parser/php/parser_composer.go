package php

import (
	"context"
	"dep-comparer/internal/parser/types"
	"encoding/json"
	"fmt"
)

type composer struct {
	Name       string            `json:"Name,omitempty"`
	Require    map[string]string `json:"require,omitempty"`
	RequireDev map[string]string `json:"require-dev,omitempty"`
}

func newDependency() *types.Dependency {
	return &types.Dependency{
		Dependencies: make(map[types.DependencyPath]types.DependencyVersion),
	}
}

func Parse(ctx context.Context, name string, data []byte) (*types.Dependency, error) {
	var c composer
	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	fmt.Println(c.Name)

	currentDep := newDependency()
	currentDep.DependencyPath = types.DependencyPath(c.Name)
	if phpVersion, ok := c.Require["php"]; ok {
		currentDep.Version = types.Version(phpVersion)
	}
	for key, version := range c.Require {
		currentDep.Dependencies[types.DependencyPath(key)] = types.DependencyVersion(version)
	}

	return currentDep, nil
}
