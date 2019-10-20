package cli

import "fmt"

func (c *CLI) Usage() error {
	fmt.Fprintf(c.Out, "hoge")
	return nil
}
