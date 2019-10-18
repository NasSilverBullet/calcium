package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca",
		Short: "calcium",
		RunE: func(cmd *cobra.Command, args []string) error {
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

				fmt.Fprint(os.Stdout, string(out))
			}

			return nil
		},
	}

	return cmd
}
