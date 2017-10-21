package buildfile

import yaml "gopkg.in/yaml.v2"

type TestContent struct {
	Name string `yaml:"name"`
}

// BuildFile is the struct with the content of .cino.ym;
type BuildFile struct {
	Main  string        `yaml:"main"`
	Tests []TestContent `yaml:"tests"`
}

// Parse parses a yml file to a BuildFile
func Parse(content []byte) (BuildFile, error) {
	out := BuildFile{}
	err := yaml.Unmarshal(content, &out)

	return out, err
}