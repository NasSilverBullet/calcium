package main

import (
	"fmt"
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
	c := &cli.CLI{
		In:   os.Stdin,
		Out:  os.Stdout,
		Err:  os.Stderr,
		Args: os.Args,
	}

	if err := c.Routes(); err != nil {
		return fmt.Errorf("Error: \n%w", err)
	}

	return nil
}
