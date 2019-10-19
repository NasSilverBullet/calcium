package main

import (
	"fmt"
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

	if err := c.Run(os.Args); err != nil {
		return err
	}

	return nil
}
