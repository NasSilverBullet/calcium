package cli_test

import (
	"bytes"
	"testing"

	"github.com/NasSilverBullet/calcium/cmd/cli"
)

func getRunYaml() []byte {
	return []byte(`version: 1

tasks:

  - task:
    use: "test1"
    run: |
      echo test

  - task:
    use: "test2"
    flags:
      - name: value
        short: v
        long: val
    run: |
      echo {{value}}`)
}

func getRunInValidYaml() []byte {
	return []byte(`hoge hoge`)
}

func TestCLIRun(t *testing.T) {

	tests := []struct {
		name        string
		args        []string
		yaml        []byte
		errexpected bool
	}{
		{"Success", []string{"calcium", "run", "test1"}, getRunYaml(), false},
		{"SuccessFlags", []string{"calcium", "run", "test2", "-v", "test"}, getRunYaml(), false},
		{"NoTaskChosenError", []string{"calcium", "run"}, getRunYaml(), true},
		{"InvalidYamlError", []string{"calcium", "run", "invalid"}, getRunInValidYaml(), true},
		{"InvalidTaskError", []string{"calcium", "run", "invalid"}, getRunYaml(), true},
		{"InvalidFlagsError", []string{"calcium", "run", "test2", "-v"}, getRunYaml(), true},
		{"NoFlagGivenError", []string{"calcium", "run", "test2", "--invalid", "invalid"}, getRunYaml(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cli.CLI{
				&bytes.Buffer{},
				&bytes.Buffer{},
				&bytes.Buffer{},
				tt.args,
				tt.yaml,
			}
			if err := c.Run(); err != nil && !tt.errexpected {
				t.Errorf("Unexpected error : %v", err)
			}

		})
	}
}
