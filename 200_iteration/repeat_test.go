package iteration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	character := "a"
	expected := "aaaaa"

	actual := Repeat(character)

	assert.Equal(t,expected,actual, "want %q, got %q",expected,actual)
}

// BenchmarkRepeat execute by go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N will appear in log as execute times ~10000000
		Repeat("a")
	}
}