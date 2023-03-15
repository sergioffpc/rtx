package cgmath_test

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPoint3Sub substracting two points.
func TestPoint3Sub(t *testing.T) {
	a, b := cgmath.Point3{3, 2, 1}, cgmath.Point3{5, 6, 7}
	assert.Equal(t, cgmath.Vector3{-2, -4, -6}, cgmath.Point3.Sub(a, b))
}

// TestPoint3SubVector substracting a vector from a point.
func TestPoint3SubVector(t *testing.T) {
	a, b := cgmath.Point3{3, 2, 1}, cgmath.Vector3{5, 6, 7}
	assert.Equal(t, cgmath.Point3{-2, -4, -6}, cgmath.Point3.SubVector(a, b))
}

// TestVector3Add adding two vectors.
func TestVector3Add(t *testing.T) {
	a, b := cgmath.Vector3{3, -2, 5}, cgmath.Vector3{-2, 3, 1}
	assert.Equal(t, cgmath.Vector3{1, 1, 6}, cgmath.Vector3.Add(a, b))
}

// TestVector3Cross the cross product of two vectors.
func TestVector3Cross(t *testing.T) {
	for _, tc := range []struct {
		a, b     cgmath.Vector3
		expected cgmath.Vector3
	}{
		{a: cgmath.Vector3{1, 2, 3}, b: cgmath.Vector3{2, 3, 4}, expected: cgmath.Vector3{-1, 2, -1}},
		{a: cgmath.Vector3{2, 3, 4}, b: cgmath.Vector3{1, 2, 3}, expected: cgmath.Vector3{1, -2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, cgmath.Vector3.Cross(tc.a, tc.b))
		})
	}
}

// TestVector3DivFloat dividing a vector by a scalar.
func TestVector3DivFloat(t *testing.T) {
	assert.Equal(t, cgmath.Vector3{0.5, -1, 1.5}, cgmath.Vector3{1, -2, 3}.DivFloat(2))
}

// TestVector3Dot the dot product of two vectors.
func TestVector3Dot(t *testing.T) {
	a, b := cgmath.Vector3{1, 2, 3}, cgmath.Vector3{2, 3, 4}
	assert.Equal(t, 20.0, cgmath.Vector3.Dot(a, b))
}

// TestVector3Len computing the length of a vector.
func TestVector3Len(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Vector3
		expected float64
	}{
		{a: cgmath.Vector3{1, 0, 0}, expected: 1},
		{a: cgmath.Vector3{0, 1, 0}, expected: 1},
		{a: cgmath.Vector3{0, 0, 1}, expected: 1},
		{a: cgmath.Vector3{1, 2, 3}, expected: math.Sqrt(14)},
		{a: cgmath.Vector3{-1, -2, -3}, expected: math.Sqrt(14)},
		{a: cgmath.Vector3{1, 2, 3}.Normalize(), expected: 1},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, cgmath.Vector3.Len(tc.a))
		})
	}
}

// TestVector3MulFloat multiplying a vector by a scalar.
func TestVector3MulFloat(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Vector3
		f        float64
		expected cgmath.Vector3
	}{
		{a: cgmath.Vector3{1, -2, 3}, f: 3.5, expected: cgmath.Vector3{3.5, -7, 10.5}},
		{a: cgmath.Vector3{1, -2, 3}, f: 0.5, expected: cgmath.Vector3{0.5, -1, 1.5}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, cgmath.Vector3.MulFloat(tc.a, tc.f))
		})
	}
}

// TestVector3Neg negating a vector.
func TestVector3Neg(t *testing.T) {
	assert.Equal(t, cgmath.Vector3{-1, 2, -3}, cgmath.Vector3{1, -2, 3}.Neg())
}

// TestVector3Normalize normalizing a vector.
func TestVector3Normalize(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Vector3
		expected cgmath.Vector3
	}{
		{a: cgmath.Vector3{4, 0, 0}, expected: cgmath.Vector3{1, 0, 0}},
		{a: cgmath.Vector3{1, 2, 3}, expected: cgmath.Vector3{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, cgmath.Vector3.Normalize(tc.a))
		})
	}
}

// TestVector3Reflect reflecting vectors.
func TestVector3Reflect(t *testing.T) {
	assert.Equal(t, cgmath.Vector3{1, 1, 0}, cgmath.Vector3.Reflect(cgmath.Vector3{1, -1, 0}, cgmath.Normal3{0, 1, 0}))
}

// TestVector3Sub substracting two vectors.
func TestVector3Sub(t *testing.T) {
	a, b := cgmath.Vector3{3, 2, 1}, cgmath.Vector3{5, 6, 7}
	assert.Equal(t, cgmath.Vector3{-2, -4, -6}, cgmath.Vector3.Sub(a, b))
}

func TestRayPosition(t *testing.T) {
	for _, tc := range []struct {
		r        cgmath.Ray
		t        float64
		expected cgmath.Point3
	}{
		{r: cgmath.Ray{cgmath.Point3{2, 3, 4}, cgmath.Vector3{1, 0, 0}, math.MaxFloat64}, t: 0, expected: cgmath.Point3{2, 3, 4}},
		{r: cgmath.Ray{cgmath.Point3{2, 3, 4}, cgmath.Vector3{1, 0, 0}, math.MaxFloat64}, t: 1, expected: cgmath.Point3{3, 3, 4}},
		{r: cgmath.Ray{cgmath.Point3{2, 3, 4}, cgmath.Vector3{1, 0, 0}, math.MaxFloat64}, t: -1, expected: cgmath.Point3{1, 3, 4}},
		{r: cgmath.Ray{cgmath.Point3{2, 3, 4}, cgmath.Vector3{1, 0, 0}, math.MaxFloat64}, t: 2.5, expected: cgmath.Point3{4.5, 3, 4}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, cgmath.Ray.Position(tc.r, tc.t))
		})
	}
}
