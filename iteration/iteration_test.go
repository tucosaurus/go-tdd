package iteration

import "testing"

func TestIter(t *testing.T) {
	output := Repeat("loki")
	expected := "lokilokilokilokiloki"

	if output != expected {
		t.Errorf("expected '%s' but got '%s'", expected, output)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
