package cli

import "fmt"

func (c *CLI) Usage() {
	fmt.Fprintf(c.Out, "Usage")
}
