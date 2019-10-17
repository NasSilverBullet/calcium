package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NasSilverBullet/calicium/pkg/commands"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	b, err := ioutil.ReadFile("testdata/calcium.yaml")
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(b)

	c, err := commands.Parse(b)
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Println(c)
	return nil
}
