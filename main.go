package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(ExitCodeOK)
	}

	os.Exit(ExitCodeError)
}

func run() error {
	s := &Std{
		Out: os.Stdout,
		Err: os.Stdin,
	}

	c := s.Get()

	if err := c.Execute(); err != nil {
		return err
	}

	return nil
}
