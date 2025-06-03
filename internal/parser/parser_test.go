package parser

import (
	"dep-comparer/internal/parser/types"
	"testing"
)

func TestGetLanguageTypeByName(t *testing.T) {
	tCases := []struct {
		nameOfProgrammingLanguage string
		expected                  types.Language
	}{
		{"Go", types.Language(0)},
		{"Golang", types.Language(0)},
		{"GoLang", types.Language(0)},
		{"PHP", types.Language(1)},
		{"JavaScript", types.Language(2)},
		{"JS", types.Language(2)},
		{"Java script", types.Language(2)},
	}

	for _, tc := range tCases {
		t.Run(tc.nameOfProgrammingLanguage, func(t *testing.T) {
			res, err := GetLanguageTypeByName(tc.nameOfProgrammingLanguage)
			if err != nil {
				t.Errorf("GetLanguageTypeByName error: %s", err)
			}
			if res != tc.expected {
				t.Errorf("GetLanguageTypeByName returned %v, expected %v", res, tc.expected)
			}
		})
	}
}

func TestGetLanguageNameByType(t *testing.T) {
	tCases := []struct {
		programmLanguageType types.Language
		expected             string
	}{
		{types.Language(0), "go"},
		{types.Language(1), "php"},
		{types.Language(2), "js"},
	}
	for _, tc := range tCases {
		t.Run(tc.expected, func(t *testing.T) {
			res, err := GetLanguageNameByType(tc.programmLanguageType)
			if err != nil {
				t.FailNow()
			}
			if res != tc.expected {
				t.Errorf("GetLanguageNameByType returned %v, expected %v", res, tc.expected)
			}
		})
	}
}
