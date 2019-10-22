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

	checkTaskFlags := map[string]bool{}
	for _, tf := range t.Flags {
		checkTaskFlags[tf.Name] = false
	}

	for gf, v := range givenFlags {
		var isCorrectGivenFlag bool

		for _, tf := range t.Flags {
			if strings.HasPrefix(gf, "-") && gf[1:] == tf.Short {
				script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), v)
				checkTaskFlags[tf.Name] = true
				isCorrectGivenFlag = true
				break
			}

			if strings.HasPrefix(gf, "--") && gf[2:] == tf.Long {
				script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), v)
				checkTaskFlags[tf.Name] = true
				isCorrectGivenFlag = true
				break
			}
		}

		if !isCorrectGivenFlag {
			return "", fmt.Errorf("Undefined %s flag in %s task", gf, t.Use)
		}
	}

	noGivenFlags := []string{}

	for tfName, isGiven := range checkTaskFlags {
		if !isGiven {
			noGivenFlags = append(noGivenFlags, tfName)
		}
	}

	if len(noGivenFlags) > 0 {
		return "", fmt.Errorf("Missing flags: %v ", noGivenFlags)
	}

	return script, nil
}
