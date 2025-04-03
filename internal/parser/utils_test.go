package parser

import (
	"dep-comparer/internal/parser/types"
	"testing"
)

func TestGetNameOfDependencyFile(t *testing.T) {
	testCases := []struct {
		pathOfFile, expectedName string
	}{
		{"testdata/go1.mod", "go1.mod"},
		{"testdata/go2.mod", "go2.mod"},
		{"testdata/go3.mod", "go3.mod"},
		{"testdata/composer_laravel.json", "composer_laravel.json"},
		{"testdata/package.json", "package.json"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.pathOfFile, func(t *testing.T) {
			name := GetNameOfDependencyFile(testCase.pathOfFile)
			if name != testCase.expectedName {
				t.Errorf("GetNameOfDependencyFile: expected %s, got %s", testCase.expectedName, name)
			}
		})
	}
}

func TestConvertSummarizeDepToList(t *testing.T) {

	mapOfDep := map[types.DependencyPath]struct{}{
		"github.com/docker/go-units":                                                  struct{}{},
		"github.com/onsi/gomega":                                                      struct{}{},
		"github.com/jonboulle/clockwork":                                              struct{}{},
		"github.com/vultr/govultr/v2":                                                 struct{}{},
		"github.com/Azure/azure-sdk-for-go/sdk/internal":                              struct{}{},
		"github.com/google/go-querystring":                                            struct{}{},
		"github.com/cenkalti/backoff/v4":                                              struct{}{},
		"go.opentelemetry.io/otel/sdk":                                                struct{}{},
		"github.com/containerd/ttrpc":                                                 struct{}{},
		"github.com/ghodss/yaml":                                                      struct{}{},
		"go.opentelemetry.io/collector/pdata":                                         struct{}{},
		"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc": struct{}{},
		"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5": struct{}{},
		"github.com/Microsoft/go-winio":                                               struct{}{},
		"github.com/spf13/pflag":                                                      struct{}{},
		"github.com/vishvananda/netns":                                                struct{}{},
		"github.com/oklog/ulid":                                                       struct{}{},
		"github.com/mattn/go-colorable":                                               struct{}{},
		"github.com/stoewer/go-strcase":                                               struct{}{},
		"gopkg.in/square/go-jose.v2":                                                  struct{}{},
		"github.com/lithammer/dedent":                                                 struct{}{},
		"github.com/prometheus/alertmanager":                                          struct{}{},
		"go.opentelemetry.io/otel/exporters/otlp/otlptrace":                           struct{}{},
		"go.opentelemetry.io/otel/trace":                                              struct{}{},
		"github.com/ovh/go-ovh":                                                       struct{}{},
		"sigs.k8s.io/yaml":                                                            struct{}{},
		"golang.org/x/mod":                                                            struct{}{},
		"github.com/armon/circbuf":                                                    struct{}{},
	}

	listOfDeps := ConvertSummarizeDepToList(mapOfDep)

	for _, dep := range listOfDeps {
		if _, ok := mapOfDep[types.DependencyPath(dep)]; ok {
			delete(mapOfDep, types.DependencyPath(dep))
		}
	}

	if len(mapOfDep) != 0 {
		t.Errorf("mapOfDep not empty: %v", mapOfDep)
	}

}
