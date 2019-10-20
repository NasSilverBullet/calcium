package cli

import (
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
		if err := c.Usage(); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	switch c.Args[1] {
	case "run":
		if err := c.Run(); err != nil {
			return errors.WithStack(err)
		}
	default:
		if err := c.Usage(); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
