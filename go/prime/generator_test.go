package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratesExpectedPrimeNumbers(t *testing.T) {
	f := primeNumberGenerator()
	assert.Equal(t, 2, f())
	assert.Equal(t, 3, f())
	assert.Equal(t, 5, f())
	assert.Equal(t, 7, f())
	assert.Equal(t, 11, f())
	assert.Equal(t, 13, f())
	assert.Equal(t, 17, f())
	assert.Equal(t, 19, f())
	assert.Equal(t, 23, f())
	assert.Equal(t, 29, f())
	assert.Equal(t, 31, f())
	assert.Equal(t, 37, f())
	assert.Equal(t, 41, f())
	assert.Equal(t, 43, f())
}
