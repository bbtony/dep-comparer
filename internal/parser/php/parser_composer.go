package php

import (
	"context"
	"dep-comparer/internal/parser/types"
	"encoding/json"
	"fmt"
)

type composer struct {
	Name string `json:"Name,omitempty"`
	Require map[string]string `json:"require,omitempty"`
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

// Parse - main function of parsing files of dependencies
//func Parse(
//	ctx context.Context,
//	listOfDepFiles []string,
//) ([]*types.Dependency, error) {
//	g, ctx := errgroup.WithContext(ctx)
//
//	depParserRes := make(chan *types.Dependency, len(listOfDepFiles))
//
//	modules := []*types.Dependency{}
//
//	for _, depFile := range listOfDepFiles {
//		g.Go(func() error {
//			// read file
//			data, err := reader.ReadFile(depFile)
//			if err != nil {
//				return err
//			}
//
//			// parse data
//			module, err := ParsePHPComposerJSON(parser.GetNameOfDependencyFile(depFile), data)
//			if err != nil {
//				return err
//			}
//
//			depParserRes <- module
//
//			return nil
//		})
//	}
//
//	if err := g.Wait(); err != nil {
//		return nil, err
//	}
//	close(depParserRes)
//
//	for module := range depParserRes {
//		modules = append(modules, module)
//	}
//
//	return modules, nil
//}
