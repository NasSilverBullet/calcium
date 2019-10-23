package cli_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/NasSilverBullet/calcium/cmd/cli"
)

func getRunYaml() []byte {
	return []byte(`version: 1

tasks:

  - task:
    use: test1
    run: |
      echo test

  - task:
    use: test2
    flags:
      - name: value
        short: v
        long: val
    run: |
      echo {{value}}

  - task:
    use: fail
    run: |
      invalid script`)
}

func getRunInValidYaml() []byte {
	return []byte(`hoge hoge`)
}

func TestCLIRun(t *testing.T) {

	tests := []struct {
		name           string
		args           []string
		yaml           []byte
		parseyamlerror error
		errexpected    bool
	}{
		{"Success", []string{"calcium", "run", "test1"}, getRunYaml(), nil, false},
		{"SuccessFlags", []string{"calcium", "run", "test2", "-v", "test"}, getRunYaml(), nil, false},
		{"NoTaskChosenError", []string{"calcium", "run"}, getRunYaml(), nil, true},
		{"FileNotFoundError", []string{"calcium", "run", "invalid"}, getRunYaml(), fmt.Errorf("error"), true},
		{"InValidYamlError", []string{"calcium", "run", "invalid"}, getRunInValidYaml(), nil, true},
		{"InvalidTaskError", []string{"calcium", "run", "invalid"}, getRunYaml(), nil, true},
		{"InvalidFlagsError", []string{"calcium", "run", "test2", "-v"}, getRunYaml(), nil, true},
		{"NoFlagGivenError", []string{"calcium", "run", "test2", "--invalid", "invalid"}, getRunYaml(), nil, true},
		{"ScriptFailError", []string{"calcium", "run", "fail", "", "invalid"}, getRunYaml(), nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cli.CLI{
				&bytes.Buffer{},
				&bytes.Buffer{},
				&bytes.Buffer{},
				tt.args,
				cli.YamlFunc(func() ([]byte, error) { return tt.yaml, tt.parseyamlerror }),
			}
			if err := c.Run(); err != nil && !tt.errexpected {
				t.Errorf("Unexpected error : %v", err)
			}

		})
	}
}
