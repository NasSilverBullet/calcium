package cli

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
	"github.com/pkg/errors"
)

func (c *CLI) Run() error {
	if len(c.Args) < 3 {
		return errors.WithStack(fmt.Errorf("Please choose task"))
	}

	ca, err := calcium.New(c.Yaml)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := ca.GetTask(c.Args[2])
	if err != nil {
		return errors.WithStack(err)
	}

	fs, err := c.parseFlags()
	if err != nil {
		err = fmt.Errorf(`%w

%s`, err, t.Usage())
		return errors.WithStack(err)
	}

	script, err := t.Parse(fs)
	if err != nil {
		err = fmt.Errorf(`%w

%s`, err, t.Usage())
		return errors.WithStack(err)
	}

	if err := c.execute(script); err != nil {
		err = fmt.Errorf(`%w

%s`, err, t.Usage())
		return errors.WithStack(err)
	}

	return nil
}

func (c *CLI) parseFlags() (map[string]string, error) {
	flagMap := map[string]string{}

	if len(c.Args) < 4 {
		return flagMap, nil
	}

	argFlagSection := c.Args[3:]

	if len(argFlagSection)%2 != 0 {
		return nil, errors.WithStack(fmt.Errorf("Invalid flags passed"))
	}

	for i, a := range argFlagSection {
		if i%2 != 0 {
			continue
		}

		if strings.HasPrefix(a, "--") {
			flagMap[a] = argFlagSection[i+1]
			continue
		}

		if strings.HasPrefix(a, "-") {
			flagMap[a] = argFlagSection[i+1]
			continue
		}
	}

	return flagMap, nil
}

func (c *CLI) execute(s string) error {
	cmd := exec.Command("sh", "-c", s)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Fprint(c.Out, string(out))

	return nil
}
