package parser

import (
	"dep-comparer/internal/parser/types"
	"os"
	"strings"
)

// SummarizeModules - prepare map of dependencies from all files
func SummarizeModules(language types.Language, modules ...*types.Dependency) (
	dep map[types.DependencyPath]struct{}, devDep map[types.DependencyPath]struct{},
) {
	dep = make(map[types.DependencyPath]struct{})
	switch language {
	case PHP, JS:
		devDep = make(map[types.DependencyPath]struct{})
	default:
	}

	// TODO: need to use concurrency
	for _, el := range modules {
		for path := range el.Dependencies {
			if _, flag := dep[path]; !flag {
				dep[path] = struct{}{}
			}
		}
	}

	for _, el := range modules {
		for path := range el.DevDependencies {
			if _, flag := devDep[path]; !flag {
				devDep[path] = struct{}{}
			}
		}
	}

	return
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
