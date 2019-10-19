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
	ts calcium.Tasks
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

			for _, a := range args {
				for _, t := range ca.Tasks {
					if t.Use != a {
						continue
					}
					ts = append(ts, t)
				}
			}

			if len(ts) == 0 {
				return fmt.Errorf("Task definition does not exist")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, t := range ts {
				if t.Description != "" {
					cmd.Printf("<<< %s >>>\n\n", t.Description)
				}

				c := exec.Command("sh", "-c", t.Run)

				out, err := c.CombinedOutput()
				if err != nil {
					return errors.WithStack(err)
				}

				cmd.Print(string(out))
			}

			return nil
		},
	}

	return cmd
}
