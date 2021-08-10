package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := repeat("a")
	exp := "aaaaa"
	if repeated != exp {
		t.Errorf("Expected %q but got %q", exp, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}

func TestExampleRepeat(t *testing.T) {
	actual := exampleRepeat("c", 10)
	exp := "cccccccccc"
	if actual != exp {
		t.Errorf("Expected %q but got %q", exp, actual)
	}
}

func BenchmarkExampleRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exampleRepeat("c", 10)
	}
}
