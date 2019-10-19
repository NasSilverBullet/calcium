package cli

import (
	"fmt"
	"io"
	"io/ioutil"
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

func (c *CLI) Run(args []string) error {
	ca, err := c.ParseCalcium()
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

	if err := c.Execute(t, fs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *CLI) ParseCalcium() (*calcium.Calcium, error) {
	b, err := ioutil.ReadFile("testdata/calcium.yaml")
	if err != nil {
		return nil, err
	}

	ca, err := calcium.New(b)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

func (c *CLI) ParseFlags(args []string) (map[string]string, error) {
	if len(args)%2 != 0 {
		return nil, errors.WithStack(fmt.Errorf("InValid Flags"))
	}

	var flagMap map[string]string

	for i, a := range args {
		if i%2 != 0 {
			continue
		}

		if strings.HasPrefix(a, "--") {
			flagMap[strings.Replace(a, "--", "", 1)] = args[i+1]
			continue
		}

		if strings.HasPrefix(a, "-") {
			flagMap[strings.Replace(a, "-", "", 1)] = args[i+1]
			continue
		}
	}

	return flagMap, nil
}

func (c *CLI) Execute(t *calcium.Task, fs map[string]string) error {
	if t == nil {
		return errors.WithStack(fmt.Errorf("No tasks"))
	}

	if t.Description != "" {
		fmt.Fprintf(c.Out, "<<< %s >>>\n\n", t.Description)
	}

	// TODO parse flag
	cmd := exec.Command("sh", "-c", t.Run)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Fprint(c.Out, string(out))

	return nil
}
