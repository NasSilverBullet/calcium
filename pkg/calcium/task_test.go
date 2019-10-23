package calcium_test

import (
	"testing"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
)

func TestTaskParse(t *testing.T) {

	task1 := &calcium.Task{
		Use: "test",
		Flags: calcium.Flags(
			[]*calcium.Flag{
				{Name: "test1", Short: "t", Long: "test"},
			},
		),
		Run: "echo {{test1}}",
	}

	task2 := &calcium.Task{
		Use: "test",
		Flags: calcium.Flags(
			[]*calcium.Flag{
				{Name: "test1"},
			},
		),
		Run: "echo test",
	}

	tests := []struct {
		name        string
		task        *calcium.Task
		argflags    map[string]string
		want        string
		errexpected bool
	}{
		{"SuccessShort", task1, map[string]string{"-t": "test"}, "echo test", false},
		{"SuccessLong", task1, map[string]string{"--test": "test"}, "echo test", false},
		{"ErrorNoGivenFlags", task1, map[string]string{}, "", true},
		{"ErrorundefinedFlag", task2, map[string]string{"-t": "test"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.task.Parse(tt.argflags)
			if err != nil && !tt.errexpected {
				t.Errorf("task.Parse(%v) return an unexpected error: %v", tt.argflags, err)
			}

			if tt.want != got {
				t.Errorf("task.Parse(%v) => got %s, but want %s", tt.argflags, tt.want, got)
			}
		})
	}

}

func TestTaskUsage(t *testing.T) {

	task1 := &calcium.Task{
		Use: "test",
		Flags: calcium.Flags(
			[]*calcium.Flag{
				{Name: "test1", Short: "t", Long: "test", Description: "test description"},
			},
		),
		Run: "echo {{test1}}",
	}

	task2 := &calcium.Task{
		Use: "test",
		Flags: calcium.Flags(
			[]*calcium.Flag{},
		),
		Run: "echo {{test1}}",
	}

	tests := []struct {
		name string
		task *calcium.Task
		want string
	}{
		{"Success", task1, `Usage:
  ca run test [flags]

Flags:
  -t, --test   test description`},

		{"SuccessNoFlags", task2, `Usage:
  ca run test`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.Usage(); tt.want != got {
				t.Errorf("t.Usage() => want %s, got %s", tt.want, got)
			}
		})
	}
}
