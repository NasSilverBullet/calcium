package cli

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

const CalciumFile = "calcium.yml"

type CLI struct {
	In, Out, Err io.Writer
	Args         []string
	Yaml
}

type Yaml interface {
	Read() ([]byte, error)
}

type YamlFunc func() ([]byte, error)

func (f YamlFunc) Read() ([]byte, error) {
	return f()
}

func (c *CLI) Read() ([]byte, error) {
	if c.Yaml != nil {
		return c.Yaml.Read()
	}

	b, err := ioutil.ReadFile(CalciumFile)
	if err != nil {
		return nil, fmt.Errorf("Error: \ncannot find %s, Please create", CalciumFile)
	}

	return b, nil
}

func (c *CLI) Routes() error {
	if len(c.Args) < 2 {
		fmt.Fprintln(c.Out, c.Usage())
		return nil
	}

	switch c.Args[1] {
	case "run":
		if err := c.Run(); err != nil {
			return errors.WithStack(err)
		}
	default:
		return fmt.Errorf(`Undefined command: %s

%s`, c.Args[1], c.Usage())
	}
	return nil
}
