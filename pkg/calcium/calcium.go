package calcium

import (
	"fmt"
	"strings"

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

func (t *Task) Parse(argFlags map[string]string) (string, error) {
	script := t.Run

	for _, f := range t.Flags {
		mustache := fmt.Sprintf("{{%s}}", f.Name)

		if strings.Index(t.Run, mustache) < 0 {
			return "", fmt.Errorf("Can not find %s flag in run section", f.Name)
		}

		var parsed bool

		if af := argFlags[fmt.Sprintf("-%s", f.Short)]; af != "" {
			script = strings.ReplaceAll(script, mustache, af)
			parsed = true
		}

		if af := argFlags[fmt.Sprintf("--%s", f.Long)]; af != "" {
			script = strings.ReplaceAll(script, mustache, af)
			parsed = true
		}

		if !parsed {
			return "", fmt.Errorf("No argument : %s were given", f.Name)
		}
	}

	return script, nil
}
