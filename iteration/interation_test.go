package interation

import (
	"strings"
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat("a",repeatCount)
	expected := strings.Repeat("a", repeatCount)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
	
}