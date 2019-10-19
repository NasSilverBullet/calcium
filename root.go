package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	ca *calcium.Calcium
	tt *calcium.Task
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca",
		Short: "calcium",
		Args: func(cmd *cobra.Command, args []string) error {
			b, err := ioutil.ReadFile("testdata/calcium.yaml")
			if err != nil {
				return err
			}

			ca, err = calcium.Parse(b)
			if err != nil {
				return err
			}

			for _, t := range ca.Tasks {
				if t.Use != args[0] {
					continue
				}
				tt = t
			}
			if tt == nil {
				return fmt.Errorf("Task definition does not exist")
			}

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			tt.Run = fmt.Sprintf(tt.RunRaw, args[1])
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if tt.Description != "" {
				cmd.Printf("<<< %s >>>\n\n", tt.Description)
			}

			c := exec.Command("sh", "-c", tt.Run)

			out, err := c.CombinedOutput()
			if err != nil {
				return errors.WithStack(err)
			}

			cmd.Print(string(out))

			return nil
		},
	}

	return cmd
}
