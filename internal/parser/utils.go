package parser

import (
	"dep-comparer/internal/parser/types"
	"os"
	"strings"
)

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

// ConvertSummarizeDepToList - prepare list of dependencies for the next reports
func ConvertSummarizeDepToList(dependencies map[types.DependencyPath]struct{}) []string {
	l := make([]string, 0, len(dependencies))
	for path, _ := range dependencies {
		l = append(l, string(path))
	}

	return l
}

// GetNameOfDependencyFile - prepare name of dependency file from path
func GetNameOfDependencyFile(path string) (result string) {
	pathOfFile := strings.Split(path, string(os.PathSeparator))
	return pathOfFile[len(pathOfFile)-1]
}
