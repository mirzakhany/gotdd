package arrayops

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {

	t.Run("test with fixed size slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := ArraySum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("test with any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := ArraySum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{5, 6})
	want := []int{3, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v , want %v", got, want)
	}

}
