package calcium

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Calcium struct {
	Task *Task `yaml:"task"`
}

type Task struct {
	Use string `yaml:"use"`
	Run string `yaml:"run"`
}

func Parse(b []byte) (*Calcium, error) {
	c := &Calcium{
		Task: &Task{},
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
