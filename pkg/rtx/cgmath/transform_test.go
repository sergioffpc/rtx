package cgmath_test

import (
	"fmt"
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMatrix44Eq matrix equality.
func TestMatrix44Eq(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Matrix44
		b        cgmath.Matrix44
		expected bool
	}{
		{
			a: cgmath.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: cgmath.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			expected: true,
		},
		{
			a: cgmath.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: cgmath.Matrix44{
				2, 3, 4, 5,
				6, 7, 8, 9,
				8, 7, 6, 5,
				4, 3, 2, 1,
			},
			expected: false,
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, cgmath.Matrix44.Eq(tc.a, tc.b))
		})
	}
}

// TestMatrix44Inverse calculation the inverse of a matrix.
func TestMatrix44Inverse(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Matrix44
		expected cgmath.Matrix44
	}{
		{
			a: cgmath.Matrix44{
				-5, 2, 6, -8,
				1, -5, 1, 8,
				7, 7, -6, -7,
				1, -3, 7, 4,
			},
			expected: cgmath.Matrix44{
				0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216,
				-0.8082706766917293, -1.456766917293233, -0.44360902255639095, 0.5206766917293233,
				-0.07894736842105263, -0.22368421052631576, -0.05263157894736842, 0.19736842105263158,
				-0.5225563909774436, -0.8139097744360901, -0.3007518796992481, 0.306390977443609,
			},
		},
		{
			a:        cgmath.Identity(),
			expected: cgmath.Identity(),
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, cgmath.Matrix44.Eq(tc.expected, cgmath.Matrix44.Inverse(tc.a)))
		})
	}
}

// TestMatrix44Mul multiplying two matrices.
func TestMatrix44Mul(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Matrix44
		b        cgmath.Matrix44
		expected cgmath.Matrix44
	}{
		{
			a: cgmath.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: cgmath.Matrix44{
				-2, 1, 2, 3,
				3, 2, 1, -1,
				4, 3, 6, 5,
				1, 2, 7, 8,
			},
			expected: cgmath.Matrix44{
				20, 22, 50, 48,
				44, 54, 114, 108,
				40, 58, 110, 102,
				16, 26, 46, 42,
			},
		},
		{
			a: cgmath.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: cgmath.Identity(),
			expected: cgmath.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, cgmath.Matrix44.Eq(tc.expected, cgmath.Matrix44.Mul(tc.a, tc.b)))
		})
	}
}

// TestMatrix44Transpose transposing a matrix.
func TestMatrix44Transpose(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Matrix44
		expected cgmath.Matrix44
	}{
		{
			a: cgmath.Matrix44{
				0, 9, 3, 0,
				9, 8, 0, 8,
				1, 8, 5, 3,
				0, 0, 5, 8,
			},
			expected: cgmath.Matrix44{
				0, 9, 1, 0,
				9, 8, 8, 0,
				3, 0, 5, 5,
				0, 8, 3, 8,
			},
		},
		{
			a:        cgmath.Identity(),
			expected: cgmath.Identity(),
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, cgmath.Matrix44.Eq(tc.expected, cgmath.Matrix44.Transpose(tc.a)))
		})
	}
}

// TestTransformChain chained transformations must be applied in reverse order.
func TestTransformChain(t *testing.T) {
	rx := cgmath.RotateXTransform(math.Pi / 2)
	st := cgmath.ScaleTransform(5, 5, 5)
	tt := cgmath.TranslateTransform(10, 5, 7)
	assert.Equal(t, cgmath.Point3{15, 0, 7}, cgmath.Point3{1, 0, 1}.Transform(cgmath.ChainTransform(rx, st, tt)))
}

// ExampleTransform chained transformations must be applied in reverse order.
func ExampleTransform() {
	rx := cgmath.RotateXTransform(math.Pi / 2)
	st := cgmath.ScaleTransform(5, 5, 5)
	tt := cgmath.TranslateTransform(10, 5, 7)
	fmt.Printf("%v", cgmath.Point3{1, 0, 1}.Transform(cgmath.ChainTransform(rx, st, tt)))
	// Output: {15 0 7}
}

// TestTransformLookAtTransform view transformation.
func TestTransformLookAtTransform(t *testing.T) {
	for _, tc := range []struct {
		from, to cgmath.Point3
		up       cgmath.Vector3
		expected cgmath.Matrix44
	}{
		// The transformation matrix for the default orientation.
		{from: cgmath.Point3{0, 0, 0}, to: cgmath.Point3{0, 0, -1}, up: cgmath.Vector3{0, 1, 0}, expected: cgmath.Identity()},
		// A view transformation matrix looking in positive z direction.
		{from: cgmath.Point3{0, 0, 0}, to: cgmath.Point3{0, 0, 1}, up: cgmath.Vector3{0, 1, 0}, expected: cgmath.ScaleTransform(-1, 1, -1).M},
		// The view transformation moves the world.
		{from: cgmath.Point3{0, 0, 8}, to: cgmath.Point3{0, 0, 0}, up: cgmath.Vector3{0, 1, 0}, expected: cgmath.TranslateTransform(0, 0, -8).M},
		// An arbitrary view transformation.
		{from: cgmath.Point3{1, 2, 3}, to: cgmath.Point3{4, -2, 8}, up: cgmath.Vector3{1, 1, 0}, expected: cgmath.Matrix44{
			-0.4999999999999999, 0.4999999999999999, 0.7, -2.5999999999999996,
			0.7495331880577403, 0.6505382386916236, 0.07071067811865474, -2.2627416997969516,
			-0.4242640687119285, 0.565685424949238, -0.7071067811865475, 1.414213562373095,
			0, 0, 0, 1,
		}},
	} {
		t.Run("", func(t *testing.T) {
			got := cgmath.LookAtTransform(tc.from, tc.to, tc.up).M
			assert.True(t, cgmath.Matrix44.Eq(tc.expected, got), fmt.Sprintf("expected: %v, got: %v", tc.expected, got))
		})
	}
}

// TestTransformRotateXTransform rotating a point around the x axis.
func TestTransformRotateXTransform(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Point3
		rx       cgmath.Transform
		expected cgmath.Point3
	}{
		{a: cgmath.Point3{0, 1, 0}, rx: cgmath.RotateXTransform(math.Pi / 2), expected: cgmath.Point3{0, 0, 1}},
		{a: cgmath.Point3{0, 1, 0}, rx: cgmath.RotateXTransform(math.Pi / 4), expected: cgmath.Point3{0, math.Sqrt(2) / 2, math.Sqrt(2) / 2}},
	} {
		t.Run("", func(t *testing.T) {
			got := tc.a.Transform(tc.rx)
			assert.True(t, cgmath.Point3.Eq(tc.expected, got), "expected: %v, got: %v", tc.expected, got)
		})
	}
}

// TestTransformRotateYTransform rotating a point around the y axis.
func TestTransformRotateYTransform(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Point3
		ry       cgmath.Transform
		expected cgmath.Point3
	}{
		{a: cgmath.Point3{0, 0, 1}, ry: cgmath.RotateYTransform(math.Pi / 2), expected: cgmath.Point3{1, 0, 0}},
		{a: cgmath.Point3{0, 0, 1}, ry: cgmath.RotateYTransform(math.Pi / 4), expected: cgmath.Point3{math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2}},
	} {
		t.Run("", func(t *testing.T) {
			got := tc.a.Transform(tc.ry)
			assert.True(t, cgmath.Point3.Eq(tc.expected, got), "expected: %v, got: %v", tc.expected, got)
		})
	}
}

// TestTransformRotateZTransform rotating a point around the z axis.
func TestTransformRotateZTransform(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Point3
		rz       cgmath.Transform
		expected cgmath.Point3
	}{
		{a: cgmath.Point3{0, 1, 0}, rz: cgmath.RotateZTransform(math.Pi / 2), expected: cgmath.Point3{-1, 0, 0}},
		{a: cgmath.Point3{0, 1, 0}, rz: cgmath.RotateZTransform(math.Pi / 4), expected: cgmath.Point3{-math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}},
	} {
		t.Run("", func(t *testing.T) {
			got := tc.a.Transform(tc.rz)
			assert.True(t, cgmath.Point3.Eq(tc.expected, got), "expected: %v, got: %v", tc.expected, got)
		})
	}
}

// TestTransformScaleTransform a scaling matrix applied to a vector.
func TestTransformScaleTransform(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Vector3
		s        cgmath.Transform
		expected cgmath.Vector3
	}{
		{a: cgmath.Vector3{-4, 6, 8}, s: cgmath.ScaleTransform(2, 3, 4), expected: cgmath.Vector3{-8, 18, 32}},
		{a: cgmath.Vector3{-4, 6, 8}, s: cgmath.ScaleTransform(2, 3, 4).Inverse(), expected: cgmath.Vector3{-2, 2, 2}},
		{a: cgmath.Vector3{2, 3, 4}, s: cgmath.ScaleTransform(-1, 1, 1), expected: cgmath.Vector3{-2, 3, 4}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, cgmath.Vector3.Eq(tc.expected, tc.a.Transform(tc.s)))
		})
	}
}

// TestTransformTranslateTransform a translation matrix applied to a point.
func TestTransformTranslateTransform(t *testing.T) {
	for _, tc := range []struct {
		a        cgmath.Point3
		s        cgmath.Transform
		expected cgmath.Point3
	}{
		{a: cgmath.Point3{-3, 4, 5}, s: cgmath.TranslateTransform(5, -3, 2), expected: cgmath.Point3{2, 1, 7}},
		{a: cgmath.Point3{-3, 4, 5}, s: cgmath.TranslateTransform(5, -3, 2).Inverse(), expected: cgmath.Point3{-8, 7, 3}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, cgmath.Point3.Eq(tc.expected, tc.a.Transform(tc.s)))
		})
	}
}
