package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := repeat("a")
	exp := "aaaaa"
	if repeated != exp {
		t.Errorf("Expected %q but got %q", exp, repeated)
	}
}
