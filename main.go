package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	c := getCmds()
	if err := c.Execute(); err != nil {
		return err
	}
	return nil
}

func getCmds() *cobra.Command {
	c := NewRootCmd()
	return c
}
