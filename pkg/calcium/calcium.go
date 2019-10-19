package calcium

import (
	"fmt"

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
	Flags       `yaml:"flags"`
	Run         string `yaml:"run"`
}

type Flags []*Flag

type Flag struct {
	Name  string `yaml:"name"`
	Short string `yaml:"short"`
	Long  string `yaml:"long"`
	Type  string `yaml:"type"`
}

func New(b []byte) (*Calcium, error) {
	c := &Calcium{
		Tasks: Tasks{},
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (ca *Calcium) GetTask(use string) (*Task, error) {
	for _, t := range ca.Tasks {
		if t.Use != use {
			continue
		}

		return t, nil
	}

	return nil, fmt.Errorf("Task definition does not exist")
}

func (t *Task) Parse() (string, error) {
	return "", nil
}
