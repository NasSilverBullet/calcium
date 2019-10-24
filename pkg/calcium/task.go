package calcium

import (
	"fmt"
	"strings"
)

// Parse Task's Run field from flags
func (t *Task) Parse(givenFlags map[string]string) (string, error) {
	script := t.Run

	checkGivenFlags := map[string]bool{}
	for gf := range givenFlags {
		checkGivenFlags[gf] = false
	}

	checkTaskFlags := map[string]bool{}
	for _, tf := range t.Flags {
		checkTaskFlags[tf.Name] = false
	}

	for _, tf := range t.Flags {
		if gf, gfv := "-"+tf.Short, givenFlags["-"+tf.Short]; gfv != "" {
			script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), gfv)
			checkGivenFlags[gf], checkTaskFlags[tf.Name] = true, true
			continue
		}

		if gf, gfv := "--"+tf.Long, givenFlags["--"+tf.Long]; gfv != "" {
			script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", tf.Name), gfv)
			checkGivenFlags[gf], checkTaskFlags[tf.Name] = true, true
			continue
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

// Usage give Task's usage
func (t *Task) Usage() string {
	m := fmt.Sprintf(`Usage:
  ca run %s`, t.Use)

	if len(t.Flags) <= 0 {
		return m
	}

	m += ` [flags]

Flags:`

	var maxShortLen, maxLongLen int

	for _, f := range t.Flags {
		if sum := len(f.Short); sum > maxShortLen {
			maxShortLen = sum
		}

		if sum := len(f.Long); sum > maxLongLen {
			maxLongLen = sum
		}
	}

	for _, f := range t.Flags {
		m += fmt.Sprintf("\n  -%s, ", f.Short)
		m += strings.Repeat(" ", maxShortLen-len(f.Short))

		m += fmt.Sprintf("--%s   ", f.Long)
		m += strings.Repeat(" ", maxLongLen-len(f.Long))

		m += f.Description
	}
	return m
}
