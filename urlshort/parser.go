package urlshort

import (
	"gopkg.in/yaml.v3"
	"encoding/json"
)

type Parser interface{
	Parse(data []byte) ([]pathUrl, error)
}

type YAMLParser struct{}
func (p YAMLParser) Parse(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	} 

	return pathUrls, err
}

type JSONParser struct{}
func (j JSONParser) Parse(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, err
}

type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string	`yaml:"url" json:"url"`
}

func BuildPath(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return pathsToUrls
}