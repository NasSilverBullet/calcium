package calcium

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Calcium struct {
	Version string `yaml:"version"`
	Tasks   `yaml:"tasks"`
}

type Tasks []*Task

type Task struct {
	Description string `yaml:"task"`
	Use         string `yaml:"use"`
	RunRaw      string `yaml:"run"`
	Run         string `yaml:"_run"`
}

func Parse(b []byte) (*Calcium, error) {
	c := &Calcium{
		Tasks: Tasks{},
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
