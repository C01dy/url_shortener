package urlshort

import (
	"reflect"
	"testing"
)

func TestParseYAML(t *testing.T) {
	yaml := `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution`

	t.Run("should parse YAML sucessfully", func(t *testing.T) {
		parsed, err := Parser.Parse(YAMLParser{}, []byte(yaml))
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}

		expected := []pathUrl{
			{Path: "/urlshort", URL: "https://github.com/gophercises/urlshort"},
			{Path: "/urlshort-final", URL: "https://github.com/gophercises/urlshort/tree/solution"},
		}

		if !reflect.DeepEqual(parsed, expected) {
			t.Errorf("parsed data does not match expected data")
		}

	})

	t.Run("should return error for invalid YAML", func(t *testing.T) {
		invalidYaml := `this: is: not: valid: yaml`
		_, err := Parser.Parse(YAMLParser{}, []byte(invalidYaml))

		if err == nil {
			t.Errorf("expected an error for invalid YAML, but got none")
		}
	})
}

func TestParseJSON(t *testing.T) {

	t.Run("should parse JSON successfully", func(t *testing.T) {
		json := `[
  			{ "path": "/mygithub", "url": "https://github.com/C01dy" }
		]`
		parsed, err := Parser.Parse(JSONParser{}, []byte(json))
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}

		expected := []pathUrl{
			{Path: "/mygithub",
				URL: "https://github.com/C01dy"}}

		if !reflect.DeepEqual(parsed, expected) {
			t.Errorf("parsed data does not match expected data")
		}
	})

	t.Run("should return error for invalid JSON", func(t *testing.T) {
		invalidJson := `[){ this: , not: valid JSON}]`

		_, err := Parser.Parse(JSONParser{}, []byte(invalidJson))

		if err == nil {
			t.Errorf("expected an error for invalid JSON, but got none")
		}
	})
}
