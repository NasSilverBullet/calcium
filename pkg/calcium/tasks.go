package calcium

import (
	"fmt"
	"strings"
)

// Usage give Tasks's usage
func (ts Tasks) Usage() string {
	m := "Usage:"

	var maxTaskLen int

	for _, t := range ts {
		if l := len(t.Use); l > maxTaskLen {
			maxTaskLen = l
		}
	}

	if len(ts) < 1 {
		m += "\n  Tasks not found"
		return m
	}

	for _, t := range ts {
		m += fmt.Sprintf("\n  ca run %s", t.Use)
		if len(t.Flags) < 1 {
			continue
		}
		m += strings.Repeat(" ", maxTaskLen-len(t.Use))
		m += " [Flags]"
	}
	return m
}
