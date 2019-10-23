package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NasSilverBullet/calcium/cmd/cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cli.ExitCodeError)
	}

	os.Exit(cli.ExitCodeOK)
}

func run() error {
	yaml, err := ioutil.ReadFile("calcium.yml")
	if err != nil {
		return fmt.Errorf("Error: \ncannot find calcium.yml, Please create")
	}

	c := &cli.CLI{
		In:   os.Stdin,
		Out:  os.Stdout,
		Err:  os.Stderr,
		Args: os.Args,
		Yaml: yaml,
	}

	if err := c.Routes(); err != nil {
		return fmt.Errorf("Error: \n%w", err)
	}

	return nil
}
