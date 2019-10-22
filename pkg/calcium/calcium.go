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

func (t *Task) Parse(givenFlags map[string]string) (string, error) {
	script := t.Run

	var taskFlagCount int

	for gf, v := range givenFlags {

		var isCorrectFlag bool

		for _, tf := range t.Flags {
			if strings.HasPrefix(gf, "-") && gf[1:] == tf.Short {
				script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), v)
				taskFlagCount++
				isCorrectFlag = true
				break
			}

			if strings.HasPrefix(gf, "--") && gf[2:] == tf.Long {
				script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), v)
				taskFlagCount++
				isCorrectFlag = true
				break
			}

			if !isCorrectFlag {
				return "", fmt.Errorf("Undefined %s flag in %s task", gf, t.Use)
			}
		}

	}

	if taskFlagCount < len(t.Flags) {
		return "", fmt.Errorf("Missing flags in %s task", t.Use)
	}

	return script, nil
}
