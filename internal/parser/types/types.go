package types

type Language int

type DependencyVersion string

type DependencyPath string

type Name string

type Version string

type Dependency struct {
	DependencyPath
	Version
	Dependencies    map[DependencyPath]DependencyVersion
	DevDependencies map[DependencyPath]DependencyVersion
}
