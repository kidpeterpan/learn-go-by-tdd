package integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 4

	sum := Add(2, 2)

	assert.Equal(t, expected, sum, "want %d, got %d",expected,sum)
}
