package integer

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	exp := 4

	if sum != exp {
		t.Errorf("Expected %q but got %q", exp, sum)
	}
}
