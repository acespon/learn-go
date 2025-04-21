package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("Expected %d but got %d", expected, sum)
	}
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
