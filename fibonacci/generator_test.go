package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratesExpectedSequence(t *testing.T) {
	f := fibonacci()
	assert.Equal(t, f(), 0)
	assert.Equal(t, f(), 1)
	assert.Equal(t, f(), 1)
	assert.Equal(t, f(), 2)
	assert.Equal(t, f(), 3)
	assert.Equal(t, f(), 5)
	assert.Equal(t, f(), 8)
	assert.Equal(t, f(), 13)
	assert.Equal(t, f(), 21)
	assert.Equal(t, f(), 34)
}
