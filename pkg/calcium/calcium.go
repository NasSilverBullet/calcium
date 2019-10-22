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
	Name        string `yaml:"name"`
	Short       string `yaml:"short"`
	Long        string `yaml:"long"`
	Description string `yaml:"description"`
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

	checkGivenFlags := map[string]bool{}
	for gf, _ := range givenFlags {
		checkGivenFlags[gf] = false
	}

	checkTaskFlags := map[string]bool{}
	for _, tf := range t.Flags {
		checkTaskFlags[tf.Name] = false
	}

	for gf, v := range givenFlags {
		for _, tf := range t.Flags {
			if strings.HasPrefix(gf, "-") && gf[1:] == tf.Short {
				script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), v)
				checkGivenFlags[gf], checkTaskFlags[tf.Name] = true, true
				break
			}

			if strings.HasPrefix(gf, "--") && gf[2:] == tf.Long {
				script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), v)
				checkGivenFlags[gf], checkTaskFlags[tf.Name] = true, true
				break
			}
		}
	}

	var errMessage string

	noGivenFlags := []string{}
	for tfName, isGiven := range checkTaskFlags {
		if !isGiven {
			noGivenFlags = append(noGivenFlags, tfName)
		}
	}
	if len(noGivenFlags) > 0 {
		errMessage += fmt.Sprintf("Missing flags: %v ", noGivenFlags)
	}

	undefinedFlags := []string{}
	for gfName, isDefined := range checkGivenFlags {
		if !isDefined {
			undefinedFlags = append(undefinedFlags, gfName)
		}
	}
	if len(undefinedFlags) > 0 {
		if len(errMessage) > 0 {
			errMessage += "\n"
		}
		errMessage += fmt.Sprintf("Undefined flags: %v ", undefinedFlags)
	}

	if len(errMessage) > 0 {
		return "", fmt.Errorf(errMessage)
	}

	return script, nil
}

func (t *Task) Usage() string {
	m := fmt.Sprintf(`Usage:
  calcium run %s`, t.Use)

	if len(t.Flags) <= 0 {
		return m
	}

	m += ` [flags]

Flags:`

	for _, f := range t.Flags {
		m += fmt.Sprintf("\n  -%s, --%s      %s", f.Short, f.Long, f.Description)
	}
	return m
}
