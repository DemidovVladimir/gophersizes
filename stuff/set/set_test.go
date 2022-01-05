package set

import "testing"

func TestMakeSet(t *testing.T) {
	s := MakeSet()
	s.Add("John")

	if !s.Exists("John") {
		t.Fatalf("John does not exists, test is not passing...")
	}
}

func BenchmarkMakeSet10(t *testing.B) {
	for n := 0; n < t.N; n++ {
		s := MakeSet()
		s.Add("John")
	}
}
