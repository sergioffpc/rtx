package rtx_test

import (
	"math"
	"sergioffpc/rtx/pkg/rtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector3Add(t *testing.T) {
	a, b := rtx.Vector3{3, -2, 5}, rtx.Vector3{-2, 3, 1}
	assert.Equal(t, rtx.Vector3{1, 1, 6}, rtx.Vector3.Add(a, b))
}

func TestVector3Cross(t *testing.T) {
	for _, tc := range []struct {
		a, b     rtx.Vector3
		expected rtx.Vector3
	}{
		{a: rtx.Vector3{1, 2, 3}, b: rtx.Vector3{2, 3, 4}, expected: rtx.Vector3{-1, 2, -1}},
		{a: rtx.Vector3{2, 3, 4}, b: rtx.Vector3{1, 2, 3}, expected: rtx.Vector3{1, -2, 1}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, rtx.Vector3.Cross(tc.a, tc.b))
		})
	}
}

func TestVector3DivFloat(t *testing.T) {
	assert.Equal(t, rtx.Vector3{0.5, -1, 1.5}, rtx.Vector3{1, -2, 3}.DivFloat(2))
}

func TestVector3Dot(t *testing.T) {
	a, b := rtx.Vector3{1, 2, 3}, rtx.Vector3{2, 3, 4}
	assert.Equal(t, 20.0, rtx.Vector3.Dot(a, b))
}

func TestVector3Len(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Vector3
		expected float64
	}{
		{a: rtx.Vector3{1, 0, 0}, expected: 1},
		{a: rtx.Vector3{0, 1, 0}, expected: 1},
		{a: rtx.Vector3{0, 0, 1}, expected: 1},
		{a: rtx.Vector3{1, 2, 3}, expected: math.Sqrt(14)},
		{a: rtx.Vector3{-1, -2, -3}, expected: math.Sqrt(14)},
		{a: rtx.Vector3{1, 2, 3}.Normalize(), expected: 1},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, rtx.Vector3.Len(tc.a))
		})
	}
}

func TestVector3MulFloat(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Vector3
		f        float64
		expected rtx.Vector3
	}{
		{a: rtx.Vector3{1, -2, 3}, f: 3.5, expected: rtx.Vector3{3.5, -7, 10.5}},
		{a: rtx.Vector3{1, -2, 3}, f: 0.5, expected: rtx.Vector3{0.5, -1, 1.5}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, rtx.Vector3.MulFloat(tc.a, tc.f))
		})
	}
}

func TestVector3Neg(t *testing.T) {
	assert.Equal(t, rtx.Vector3{-1, 2, -3}, rtx.Vector3{1, -2, 3}.Neg())
}

func TestVector3Normalize(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Vector3
		expected rtx.Vector3
	}{
		{a: rtx.Vector3{4, 0, 0}, expected: rtx.Vector3{1, 0, 0}},
		{a: rtx.Vector3{1, 2, 3}, expected: rtx.Vector3{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, rtx.Vector3.Normalize(tc.a))
		})
	}
}

func TestVector3Sub(t *testing.T) {
	a, b := rtx.Vector3{3, 2, 1}, rtx.Vector3{5, 6, 7}
	assert.Equal(t, rtx.Vector3{-2, -4, -6}, rtx.Vector3.Sub(a, b))
}

func TestRayPosition(t *testing.T) {
	for _, tc := range []struct {
		r        rtx.Ray
		t        float64
		expected rtx.Point3
	}{
		{r: rtx.Ray{rtx.Point3{2, 3, 4}, rtx.Vector3{1, 0, 0}, math.MaxFloat64}, t: 0, expected: rtx.Point3{2, 3, 4}},
		{r: rtx.Ray{rtx.Point3{2, 3, 4}, rtx.Vector3{1, 0, 0}, math.MaxFloat64}, t: 1, expected: rtx.Point3{3, 3, 4}},
		{r: rtx.Ray{rtx.Point3{2, 3, 4}, rtx.Vector3{1, 0, 0}, math.MaxFloat64}, t: -1, expected: rtx.Point3{1, 3, 4}},
		{r: rtx.Ray{rtx.Point3{2, 3, 4}, rtx.Vector3{1, 0, 0}, math.MaxFloat64}, t: 2.5, expected: rtx.Point3{4.5, 3, 4}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, rtx.Ray.Position(tc.r, tc.t))
		})
	}
}
