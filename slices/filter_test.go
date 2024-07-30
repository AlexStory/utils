package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("Filter works", func(t *testing.T) {
		{
			got := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
			want := []int{2, 4}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Filter test failed, got = %v; want %v", got, want)
			}
		}
	})

	t.Run("Filter doesnt mutate the original slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		Filter(slice, func(v int) bool { return v%2 == 0 })
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("Filter test failed, got = %v; want %v", slice, want)
		}
	})
}
