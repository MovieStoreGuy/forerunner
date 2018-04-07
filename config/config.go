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
	Cmds []string `json:"Commands" yaml:"Commands"`
}

// Load will take a given path map that to the Set struct
func Load(path string) (*Set, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Set
	err = yaml.Unmarshal(buff, &conf)
	if conf.Network == "" {
		conf.Network = "bridge"
	}
	return &conf, err
}

// Default creates the default config that can be used if nothing is defined
func Default() *Set {
	return &Set{
		Network:     "bridge",
		Environment: []string{"shit=poop"},
		Cmds:        []string{},
	}
}
