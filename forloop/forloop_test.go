package forloop

import "testing"

func TestRepeat(t *testing.T) {

	expected := "aaaaa"
	repeated := Repeate("a", 5)

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeate("a", 5)
	}
}
