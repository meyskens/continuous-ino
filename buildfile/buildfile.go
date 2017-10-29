package buildfile

import yaml "gopkg.in/yaml.v2"

// TestContent is the content of a single test
type TestContent struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
}

// BuildFile is the struct with the content of .cino.ym;
type BuildFile struct {
	Main  string        `yaml:"main"`
	Baud  int           `yaml:"baud"`
	Tests []TestContent `yaml:"tests"`
}

// New returns an empty BuildFile with the defaults set
func New() BuildFile {
	return BuildFile{
		Baud: 9600,
	}
}

// Parse parses a yml file to a BuildFile
func Parse(content []byte) (BuildFile, error) {
	out := New()
	err := yaml.Unmarshal(content, &out)

	return out, err
}
