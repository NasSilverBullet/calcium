package calcium_test

import (
	"testing"

	"github.com/NasSilverBullet/calcium/pkg/calcium"
)

func TestTasksUsage(t *testing.T) {
	tests := []struct {
		name  string
		tasks calcium.Tasks
		want  string
	}{
		{"HasTasks", calcium.Tasks{&calcium.Task{Use: "test"}, &calcium.Task{Use: "test2", Flags: calcium.Flags{&calcium.Flag{}}}}, "Usage:\n  ca run test\n  ca run test2 [Flags]"},
		{"NoTasks", calcium.Tasks{}, "Usage:\n  Tasks not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tasks.Usage(); tt.want != got {
				t.Errorf("tasks.Usage() => got %s, but want %s", got, tt.want)
			}
		})
	}
}
