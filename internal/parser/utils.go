package parser

import "dep-comparer/internal/parser/types"

// SummarizeModules - prepare map of dependencies from all files
func SummarizeModules(modules ...*types.Dependency) map[types.DependencyPath]struct{} {
	res := make(map[types.DependencyPath]struct{})
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
func ConvertSummarizeDepToList(dependencies map[types.DependencyPath]struct{}) []string {
	l := make([]string, 0, len(dependencies))
	for path, _ := range dependencies {
		l = append(l, string(path))
	}

	return l
}
