package calcium_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name        string
		b           []byte
		errexpected bool
	}{
		{"Success", []byte(`tasks:
  - task:
    use: "test"
    run: |
      echo test`), false},
		{"Error", []byte(`hoge hoge`), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := calcium.New(tt.b)
			if err != nil && !tt.errexpected {
				t.Errorf("calcium.New return an unexpected error")
			}

			if err == nil && tt.errexpected {
				t.Errorf("calcium.New didn't return an expected error")
			}
		})
	}
}

func TestCalciumGetTask(t *testing.T) {
	tests := []struct {
		name        string
		ca          *calcium.Calcium
		use         string
		want        *calcium.Task
		errexpected bool
	}{
		{"Success", &calcium.Calcium{Tasks: calcium.Tasks([]*calcium.Task{&calcium.Task{Use: "test1"}, &calcium.Task{Use: "test"}})}, "test", &calcium.Task{Use: "test"}, false},
		{"Success", &calcium.Calcium{Tasks: calcium.Tasks([]*calcium.Task{&calcium.Task{Use: "error"}})}, "test", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ca.GetTask(tt.use)
			if err != nil && !tt.errexpected {
				t.Errorf("ca.GetTask(%s) return an unexpected error", tt.use)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ca.GetTask(%s) => got %v, but want %v", tt.use, got, tt.want)
			}
		})
	}
}

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
