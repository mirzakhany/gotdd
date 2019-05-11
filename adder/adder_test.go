package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	sum := Adder(2, 2)
	expectation := 4

	if sum != expectation {
		t.Errorf("got %d , want %d", sum, expectation)
	}
}

func ExampleAdder() {
	sum := Adder(2, 4)
	fmt.Println(sum)
	// Output: 6
}
