package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	// Test: Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }) == []int{2, 4}
	{
		got := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
		want := []int{2, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Filter test failed, got = %v; want %v", got, want)
		}
	}
}
