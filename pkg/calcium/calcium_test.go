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
