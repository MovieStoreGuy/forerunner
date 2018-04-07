package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Set is an container of variables that are required at various points of
// the Runner lifetime
type Set struct {
	// Variables to use that will passed onto Docker
	Network     string   `json:"Network,omitempty" yaml:"Network,omitempty"`
	Environment []string `json:"Environment,omitempty" yaml:"Environment,omitempty"`
	// Variables that will be passed onto the Spartan
	Cmds []([]string) `json:"Commands" yaml"Commands"`
}

// Load will take a given path map that to the Set struct
func (s *Set) Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	buff, err := ioutil.ReadAll(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buff, s)
}
