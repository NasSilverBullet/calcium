package cli_test

import (
	"bytes"
	"testing"

	"github.com/NasSilverBullet/calcium/cmd/cli"
)

func getCLIYaml() []byte {
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

func TestCLIRoutes(t *testing.T) {

	tests := []struct {
		name        string
		args        []string
		yaml        []byte
		errexpected bool
	}{
		{"SuccessUsage", []string{"calcium"}, getCLIYaml(), false},
		{"SuccessRun", []string{"calcium", "run", "test1"}, getCLIYaml(), false},
		{"InvalidRunError", []string{"calcium", "run", "invalid"}, getCLIYaml(), true},
		{"InvalidCommand", []string{"calcium", "invalid"}, getCLIYaml(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cli.CLI{
				&bytes.Buffer{},
				&bytes.Buffer{},
				&bytes.Buffer{},
				tt.args,
				cli.YamlFunc(func() ([]byte, error) { return tt.yaml, nil }),
			}
			if err := c.Routes(); err != nil && !tt.errexpected {
				t.Errorf("Unexpected error : %v", err)
			}

		})
	}
}
