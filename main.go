package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NasSilverBullet/calcium/cmd/cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(cli.ExitCodeOK)
	}

	os.Exit(cli.ExitCodeError)
}

func run() error {
	c := &cli.CLI{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}

	yaml, err := ioutil.ReadFile("calcium.yml")
	if err != nil {
		return fmt.Errorf("cannot find calcium.yml, Please place")
	}

	if err := c.Run(os.Args, yaml); err != nil {
		return err
	}

	return nil
}
