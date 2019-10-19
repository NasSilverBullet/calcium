package main

import (
	"io"

	"github.com/spf13/cobra"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type Std struct {
	Out, Err io.Writer
}

func (s *Std) Get() *cobra.Command {
	c := NewRootCmd()
	c.SetOut(s.Out)
	c.SetErr(s.Err)
	return c
}
