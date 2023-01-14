package translate

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

var catalog = Catalog{
	language.English: Dictionary{
		458405366:  "First test",
		2703638123: "Second test",
	},
	language.Russian: Dictionary{
		458405366: "Первая проверка",
	},
}

func TestNewFromDir(t *testing.T) {
	tr, err := NewFromDir("testdata")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if !reflect.DeepEqual(tr.cat, catalog) {
		t.Fatalf("Expected catalog %v, got %v", catalog, tr.cat)
	}
}

func TestTranslate(t *testing.T) {
	type testCase struct {
		name string
		str  string
		lang language.Tag
		exp  string
	}

	var cases = []testCase{
		{"English1", "First test", language.English, "First test"},
		{"English2", "Second test", language.English, "Second test"},
		{"NotFound", "test", language.English, "test"},
		{"Russian", "First test", language.Russian, "Первая проверка"},
		{"InvalidTag", "Hello", language.Tag{}, "Hello"},
	}

	tr := New(catalog)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := tr.TranslateTo(tc.str, tc.lang)
			if out != tc.exp {
				t.Errorf("Expected '%s', got '%s'", tc.exp, out)
			}
		})
	}
}
