package calcium

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Calcium has Task and serveral info
type Calcium struct {
	Version string `yaml:"version"`
	Tasks   `yaml:"tasks"`
}

// Tasks is Task's slice
type Tasks []*Task

// Task has shell script info
type Task struct {
	Description string `yaml:"task"`
	Use         string `yaml:"use"`
	Flags       `yaml:"flags"`
	Run         string `yaml:"run"`
}

// Flags is Flag's slice
type Flags []*Flag

// Flag is cli flag
type Flag struct {
	Name        string `yaml:"name"`
	Short       string `yaml:"short"`
	Long        string `yaml:"long"`
	Description string `yaml:"description"`
}

// New is Calcium constructor
func New(b []byte) (*Calcium, error) {
	c := &Calcium{
		Tasks: Tasks{},
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}

	return c, nil
}

// GetTask give Task from Task's 'use'
func (ca *Calcium) GetTask(use string) (*Task, error) {
	for _, t := range ca.Tasks {
		if t.Use != use {
			continue
		}

		return t, nil
	}

	return nil, fmt.Errorf("Task definition does not exist")
}
