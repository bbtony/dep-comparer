package php

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		pathOfTestFile, expected string
	}{
		{"./../../../testdata/composer_laravel.json", "laravel/framework"},
		{"./../../../testdata/composer_symfony.json", "symfony/symfony"},
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
