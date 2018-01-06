package physics

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertVectorComponents(t *testing.T, v Vector3, expComps []float64) {
	assert.Equal(t, expComps[0], v.X)
	assert.Equal(t, expComps[1], v.Y)
	assert.Equal(t, expComps[2], v.Z)
}

func TestConstructorReturnsExpectedVector3(t *testing.T) {
	assertVectorComponents(t, Vector3{}, []float64{0.0, 0.0, 0.0})
	assertVectorComponents(t, Vector3{1.0, 2.1, 3.1}, []float64{1.0, 2.1, 3.1})
	assertVectorComponents(t,
		Vector3{-1.0, 0.1, -43.1},
		[]float64{-1.0, 0.1, -43.1})
}

func TestAddReturnsExpectedVector3(t *testing.T) {
	v1 := Vector3{1.0, 1.0, 1.0}
	v2 := Vector3{2.0, 1.0, 3.0}
	v3 := v1.Add(v2)
	assert.Equal(t, Vector3{3.0, 2.0, 4.0}, v3)
}

func TestSubReturnsExpectedVector3(t *testing.T) {
	v1 := Vector3{1.0, 1.0, 1.0}
	v2 := Vector3{2.0, 1.0, 3.0}
	v3 := v1.Sub(v2)
	assert.Equal(t, Vector3{-1.0, 0.0, -2.0}, v3)
}

func TestMulReturnsExpectedVector3(t *testing.T) {
	v := Vector3{1.0, 1.0, 1.0}
	actual := v.Mul(3.0)
	assert.Equal(t, Vector3{3.0, 3.0, 3.0}, actual)

	actual = actual.Mul(-1.0)
	assert.Equal(t, Vector3{-3.0, -3.0, -3.0}, actual)
}

func TestDivReturnsExpectedVector3(t *testing.T) {
	v := Vector3{5.0, 10.0, 15.0}
	actual := v.Div(5.0)
	assert.Equal(t, Vector3{1.0, 2.0, 3.0}, actual)

	actual = actual.Div(-1.0)
	assert.Equal(t, Vector3{-1.0, -2.0, -3.0}, actual)
}

func TestMagnitudeReturnsExpectedValue(t *testing.T) {
	v := Vector3{1, 1, 1}
	assert.Equal(t, math.Sqrt(3), v.Magnitude())

	v = Vector3{2, 2, 2}
	assert.Equal(t, 2*math.Sqrt(3), v.Magnitude())
}

func TestNormalizeReturnsVector3WithMagnitudeOf1(t *testing.T) {
	v := Vector3{1, 1, 1}
	mag := v.Normalize().Magnitude()
	assert.InDelta(t, 1.0, mag, 0.01)

	v = Vector3{4, 12, -1}
	mag = v.Normalize().Magnitude()
	assert.InDelta(t, 1.0, mag, 0.01)
}

func TestDotReturnsExpectedFloatValue(t *testing.T) {
	v1 := Vector3{1, 0, 0}
	v2 := Vector3{1, 0, 0}
	assert.Equal(t, 1.0, v1.Dot(v2))

	v1 = Vector3{1, 0, 0}
	v2 = Vector3{-1, 0, 0}
	assert.Equal(t, -1.0, v1.Dot(v2))

	v1 = Vector3{1, 0, 0}
	v2 = Vector3{0, 1, 0}
	assert.Equal(t, 0.0, v1.Dot(v2))
}
