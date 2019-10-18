package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	b, err := ioutil.ReadFile("testdata/calcium.yaml")
	if err != nil {
		return errors.WithStack(err)
	}

	ca, err := calcium.Parse(b)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, t := range ca.Tasks {
		cmd := exec.Command("sh", "-c", t.Run)

		out, err := cmd.CombinedOutput()
		if err != nil {
			return errors.WithStack(err)
		}

		fmt.Fprintln(os.Stdout, string(out))
	}

	return nil
}
