package lists

import (
	"reflect"
	"testing"
)

func TestLinked(t *testing.T) {
	tests := map[string]struct {
		data []int
		list []int
	}{
		"simple": {
			data: []int{1, 3, 20, 4, 1, 0},
			list: []int{1, 3, 20, 4, 1, 0},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			linked := NewLinked()
			for _, data := range test.data {
				linked.Insert(data)
			}
			if !reflect.DeepEqual(test.list, linked.AsSlice()) {
				t.Errorf("Unexpected result: expecting %#v, got %#v", test.list, linked.AsSlice())
			}
		})
	}
}
