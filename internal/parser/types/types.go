package types

type DependencyVersion string

type DependencyPath string

type Name string

type Version string

type Dependency struct {
	DependencyPath
	Version
	Dependencies map[DependencyPath]DependencyVersion
}
