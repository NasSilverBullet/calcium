package commands

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Command struct {
	Task *Task
}

type Task struct {
	Use string `yaml:"use"`
	Run string `yaml:"run"`
}

func Parse(b []byte) (*Command, error) {
	c := &Command{
		Task: &Task{},
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
