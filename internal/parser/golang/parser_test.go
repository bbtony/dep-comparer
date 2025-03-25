package golang

import (
	"os"
	"testing"
)

func TestTableParseGoMod(t *testing.T) {
	testCases := []struct {
		pathOfTestFile, expected string
	}{
		{"./../../testdata/go1.mod", "go.opentelemetry.io/collector"},
		{"./../../testdata/go2.mod", "k8s.io/kubernetes"},
		{"./../../testdata/go3.mod", "github.com/prometheus/prometheus"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.pathOfTestFile, func(t *testing.T) {
			data, err := os.ReadFile(testCase.pathOfTestFile)
			if err != nil {
				t.Fatal(err)
			}
			mod, err := parseGoMod(data)
			if err != nil {
				t.Fatal()
			}
			if string(mod.Name) != testCase.expected {
				t.Errorf("got %v, want %v", mod, testCase.expected)
			}
		})
	}
}
