package golang

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestTableParseGoMod(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		pathOfTestFile, expected string
	}{
		{"./../../../testdata/go1.mod", "go.opentelemetry.io/collector"},
		{"./../../../testdata/go2.mod", "k8s.io/kubernetes"},
		{"./../../../testdata/go3.mod", "github.com/prometheus/prometheus"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.pathOfTestFile, func(t *testing.T) {
			data, err := os.ReadFile(testCase.pathOfTestFile)
			if err != nil {
				t.Fatal(err)
			}
			name := strings.Split(testCase.pathOfTestFile, string(os.PathSeparator))
			mod, err := Parse(ctx, name[len(name)-1], data)
			if err != nil {
				t.Fatal()
			}
			if string(mod.DependencyPath) != testCase.expected {
				t.Errorf("got %v, want %v", mod, testCase.expected)
			}
		})
	}
}
