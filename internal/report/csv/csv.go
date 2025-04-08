package csv

import (
	"dep-comparer/internal/parser/types"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	ByRows   = "Rows"
	ByColumn = "Column"
)

func NewReport(
	language string,
	listOfDependencies []string,
	listOfDevDependencies []string,
	order string,
	modules ...*types.Dependency,
) (report string, err error) {
	var res [][]string

	switch order {
	case ByColumn:
		res = make([][]string, 0, len(modules)+1) // plus 1 for headers
		res = prepareReportByColumn(language, listOfDependencies, modules...)
	case ByRows:
		res = make([][]string, 0, len(listOfDependencies)+2) // plus 2 for headers (modules and go version)
		res = prepareReportByRows(language, listOfDependencies, listOfDevDependencies, modules...)
	default:
		// default is ByRows
		res = make([][]string, 0, len(listOfDependencies)+2) // plus 2 for headers (modules and go version)
		res = prepareReportByRows(language, listOfDependencies, listOfDevDependencies, modules...)
	}

	report, err = writeAllToCSV(res)
	if err != nil {
		return report, err
	}

	return report, nil
}

// prepareReportByRows - make a report where columns like modules and rows like dependencies
// For example:
// --------------------------------------------
// modules 	  |  service  | service | service |
// go version |    1.20   |  1.21   |   1.19  |
// dep_1 	  |    v1.01  |  v1.01  |  v1.01  |
// dep_2      |   v0.0.1  |    -    |  v3.0   |
func prepareReportByRows(
	lang string,
	listOfDependencies []string,
	listOfDevDependencies []string,
	dependencies ...*types.Dependency,
) [][]string {
	res := make([][]string, 2, len(listOfDependencies)+2)

	// set a path of module
	// dependencies 	  |  service  | service | service |
	headers := make([]string, 1, len(dependencies)+1)
	headers[0] = "dependencies"

	// set a version of module
	// go version |    1.20   |  1.21   |   1.19  |
	versionHeaders := make([]string, 1, len(dependencies)+1)
	versionHeaders[0] = lang + " version"

	for _, m := range dependencies {
		headers = append(headers, string(m.DependencyPath))
		versionHeaders = append(versionHeaders, string(m.Version))
	}

	res[0] = headers
	res[1] = versionHeaders

	// dependencies
	for _, dep := range listOfDependencies {
		dependency := make([]string, 0, len(dependencies))
		dependency = append(dependency, dep)
		for _, m := range dependencies {
			// here we check dependency in module dependencies if there is then add to list and go next
			if v, ok := m.Dependencies[types.DependencyPath(dep)]; ok {
				dependency = append(dependency, string(v))
				continue
			}
			// skip dependency if there is not
			dependency = append(dependency, "-")
		}
		res = append(res, dependency)
	}

	if listOfDevDependencies != nil {
		// next index of rows and prepare head row of require-dev
		nextIndex := len(res) - 1
		reqDev := make([]string, 1, len(dependencies)+1)
		reqDev[0] = "require-dev"
		for _, dep := range dependencies {
			reqDev = append(reqDev, string(dep.DependencyPath))
		}
		res[nextIndex] = reqDev

		// devDependencies
		for _, dep := range listOfDevDependencies {
			dependency := make([]string, 0, len(dependencies))
			dependency = append(dependency, dep)
			for _, m := range dependencies {
				if v, ok := m.DevDependencies[types.DependencyPath(dep)]; ok {
					dependency = append(dependency, string(v))
					continue
				}
				dependency = append(dependency, "-")
			}

			res = append(res, dependency)
		}
	}

	return res
}

// prepareReportByColumn - make a report where columns like dependencies and rows like modules
// For example:
// ------------------------------------------------------
// module  | go version | dep_1 | dep_2 | dep_3 | dep_4 |
// service |    1.20    |   -   |   -   | v1.01 |	-   |
// service |    1.19    |   -   |  v3.0 | v1.0  |	-   |
func prepareReportByColumn(lang string, listOfDependencies []string, modules ...*types.Dependency) (res [][]string) {
	// make headers for csv report
	headers := make([]string, 2, len(listOfDependencies)+2)
	headers[0] = "module"
	headers[1] = lang + " version"
	headers = append(headers, listOfDependencies...)

	res = append(res, headers)

	// prepare structure of report
	for _, m := range modules {
		service := make([]string, 2, len(listOfDependencies)+2)
		service[0] = string(m.DependencyPath) // put path of module
		service[1] = string(m.Version)        // put  version of module

		for _, dep := range listOfDependencies {
			// here we check dependency in all dependencies if there is then go next
			if v, ok := m.Dependencies[types.DependencyPath(dep)]; ok {
				service = append(service, string(v))
				continue
			}
			service = append(service, "-")
		}

		res = append(res, service)
	}

	return res
}

// writeAllToCSV - generate new csv reports with all dependencies
func writeAllToCSV(records [][]string) (report string, err error) {
	report = "report-" + strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	f, err := os.Create(report)
	if err != nil {
		return "", err
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("could not close report: %v", err)
		}
	}()

	w := csv.NewWriter(f)

	err = w.WriteAll(records)
	if err != nil {
		return "", err
	}

	return report, err
}
