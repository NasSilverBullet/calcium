package cli

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type CLI struct {
	In, Out, Err io.Writer
	Args         []string
	Yaml         []byte
}

func (c *CLI) Routes() error {
	if len(c.Args) < 2 {
		c.Usage()
		return nil
	}

	switch c.Args[1] {
	case "run":
		if err := c.Run(); err != nil {
			return errors.WithStack(err)
		}
	default:
		c.Usage()
		return fmt.Errorf("Undefined command : %s", c.Args[1])
	}
	return nil
}
