package cli

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
	"github.com/pkg/errors"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type CLI struct {
	In, Out, Err io.Writer
}

func (c *CLI) Run(args []string, yaml []byte) error {
	ca, err := calcium.New(yaml)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := ca.GetTask(args[1])
	if err != nil {
		return errors.WithStack(err)
	}

	fs, err := c.ParseFlags(args[2:])
	if err != nil {
		return errors.WithStack(err)
	}

	script, err := t.Parse(fs)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.Execute(script); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *CLI) ParseFlags(args []string) (map[string]string, error) {
	if len(args)%2 != 0 {
		return nil, errors.WithStack(fmt.Errorf("InValid Flags"))
	}

	flagMap := map[string]string{}

	for i, a := range args {
		if i%2 != 0 {
			continue
		}

		if strings.HasPrefix(a, "--") {
			flagMap[a] = args[i+1]
			continue
		}

		if strings.HasPrefix(a, "-") {
			flagMap[a] = args[i+1]
			continue
		}
	}

	return flagMap, nil
}

func (c *CLI) Execute(s string) error {
	cmd := exec.Command("sh", "-c", s)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Fprint(c.Out, string(out))

	return nil
}
