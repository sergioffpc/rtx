package rtx_test

import (
	"math"
	"sergioffpc/rtx/pkg/rtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMatrix44Eq matrix equality.
func TestMatrix44Eq(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Matrix44
		b        rtx.Matrix44
		expected bool
	}{
		{
			a: rtx.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: rtx.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			expected: true,
		},
		{
			a: rtx.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: rtx.Matrix44{
				2, 3, 4, 5,
				6, 7, 8, 9,
				8, 7, 6, 5,
				4, 3, 2, 1,
			},
			expected: false,
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, rtx.Matrix44.Eq(tc.a, tc.b))
		})
	}
}

// TestMatrix44Inverse calculation the inverse of a matrix.
func TestMatrix44Inverse(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Matrix44
		expected rtx.Matrix44
	}{
		{
			a: rtx.Matrix44{
				-5, 2, 6, -8,
				1, -5, 1, 8,
				7, 7, -6, -7,
				1, -3, 7, 4,
			},
			expected: rtx.Matrix44{
				0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216,
				-0.8082706766917293, -1.456766917293233, -0.44360902255639095, 0.5206766917293233,
				-0.07894736842105263, -0.22368421052631576, -0.05263157894736842, 0.19736842105263158,
				-0.5225563909774436, -0.8139097744360901, -0.3007518796992481, 0.306390977443609,
			},
		},
		{
			a:        rtx.Identity(),
			expected: rtx.Identity(),
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Matrix44.Eq(tc.expected, rtx.Matrix44.Inverse(tc.a)))
		})
	}
}

// TestMatrix44Mul multiplying two matrices.
func TestMatrix44Mul(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Matrix44
		b        rtx.Matrix44
		expected rtx.Matrix44
	}{
		{
			a: rtx.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: rtx.Matrix44{
				-2, 1, 2, 3,
				3, 2, 1, -1,
				4, 3, 6, 5,
				1, 2, 7, 8,
			},
			expected: rtx.Matrix44{
				20, 22, 50, 48,
				44, 54, 114, 108,
				40, 58, 110, 102,
				16, 26, 46, 42,
			},
		},
		{
			a: rtx.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
			b: rtx.Identity(),
			expected: rtx.Matrix44{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			},
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Matrix44.Eq(tc.expected, rtx.Matrix44.Mul(tc.a, tc.b)))
		})
	}
}

// TestMatrix44Transpose transposing a matrix.
func TestMatrix44Transpose(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Matrix44
		expected rtx.Matrix44
	}{
		{
			a: rtx.Matrix44{
				0, 9, 3, 0,
				9, 8, 0, 8,
				1, 8, 5, 3,
				0, 0, 5, 8,
			},
			expected: rtx.Matrix44{
				0, 9, 1, 0,
				9, 8, 8, 0,
				3, 0, 5, 5,
				0, 8, 3, 8,
			},
		},
		{
			a:        rtx.Identity(),
			expected: rtx.Identity(),
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Matrix44.Eq(tc.expected, rtx.Matrix44.Transpose(tc.a)))
		})
	}
}

func TestTransformChain(t *testing.T) {
	rx := rtx.RotateXTransform(math.Pi / 2)
	st := rtx.ScaleTransform(5, 5, 5)
	tt := rtx.TranslateTransform(10, 5, 7)
	assert.Equal(t, rtx.Point3{15, 0, 7}, rtx.Point3{1, 0, 1}.Transform(rx, st, tt))
}

// ExampleTransform chained transformations must be applied in reverse order.
// Output: rtx.Point3{15, 0, 7}
func ExampleTransform() {
	rx := rtx.RotateXTransform(math.Pi / 2)
	st := rtx.ScaleTransform(5, 5, 5)
	tt := rtx.TranslateTransform(10, 5, 7)
	rtx.Point3{1, 0, 1}.Transform(rx, st, tt)
}

// TestTransformRotateXTransform rotating a point around the x axis.
func TestTransformRotateXTransform(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Point3
		rx       rtx.Transform
		expected rtx.Point3
	}{
		{a: rtx.Point3{0, 1, 0}, rx: rtx.RotateXTransform(math.Pi / 2), expected: rtx.Point3{0, 0, 1}},
		{a: rtx.Point3{0, 1, 0}, rx: rtx.RotateXTransform(math.Pi / 4), expected: rtx.Point3{0, math.Sqrt(2) / 2, math.Sqrt(2) / 2}},
		{a: rtx.Point3{0, 1, 0}, rx: rtx.RotateXTransform(math.Pi / 4).Inverse(), expected: rtx.Point3{0, math.Sqrt(2) / 2, -math.Sqrt(2) / 2}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Point3.Eq(tc.expected, tc.a.Transform(tc.rx)))
		})
	}
}

// TestTransformRotateYTransform rotating a point around the y axis.
func TestTransformRotateYTransform(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Point3
		ry       rtx.Transform
		expected rtx.Point3
	}{
		{a: rtx.Point3{0, 0, 1}, ry: rtx.RotateYTransform(math.Pi / 2), expected: rtx.Point3{1, 0, 0}},
		{a: rtx.Point3{0, 0, 1}, ry: rtx.RotateYTransform(math.Pi / 4), expected: rtx.Point3{math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2}},
		{a: rtx.Point3{0, 0, 1}, ry: rtx.RotateYTransform(math.Pi / 4).Inverse(), expected: rtx.Point3{-math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Point3.Eq(tc.expected, tc.a.Transform(tc.ry)))
		})
	}
}

// TestTransformRotateZTransform rotating a point around the z axis.
func TestTransformRotateZTransform(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Point3
		rz       rtx.Transform
		expected rtx.Point3
	}{
		{a: rtx.Point3{0, 1, 0}, rz: rtx.RotateZTransform(math.Pi / 2), expected: rtx.Point3{-1, 0, 0}},
		{a: rtx.Point3{0, 1, 0}, rz: rtx.RotateZTransform(math.Pi / 4), expected: rtx.Point3{-math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}},
		{a: rtx.Point3{0, 1, 0}, rz: rtx.RotateZTransform(math.Pi / 4).Inverse(), expected: rtx.Point3{0, 1, 0}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Point3.Eq(tc.expected, tc.a.Transform(tc.rz)))
		})
	}
}

// TestTransformScaleTransform a scaling matrix applied to a vector.
func TestTransformScaleTransform(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Vector3
		s        rtx.Transform
		expected rtx.Vector3
	}{
		{a: rtx.Vector3{-4, 6, 8}, s: rtx.ScaleTransform(2, 3, 4), expected: rtx.Vector3{-8, 18, 32}},
		{a: rtx.Vector3{-4, 6, 8}, s: rtx.ScaleTransform(2, 3, 4).Inverse(), expected: rtx.Vector3{-2, 2, 2}},
		{a: rtx.Vector3{2, 3, 4}, s: rtx.ScaleTransform(-1, 1, 1), expected: rtx.Vector3{-2, 3, 4}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Vector3.Eq(tc.expected, tc.a.Transform(tc.s)))
		})
	}
}

// TestTransformTranslateTransform a translation matrix applied to a point.
func TestTransformTranslateTransform(t *testing.T) {
	for _, tc := range []struct {
		a        rtx.Point3
		s        rtx.Transform
		expected rtx.Point3
	}{
		{a: rtx.Point3{-3, 4, 5}, s: rtx.TranslateTransform(5, -3, 2), expected: rtx.Point3{2, 1, 7}},
		{a: rtx.Point3{-3, 4, 5}, s: rtx.TranslateTransform(5, -3, 2).Inverse(), expected: rtx.Point3{-8, 7, 3}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Point3.Eq(tc.expected, tc.a.Transform(tc.s)))
		})
	}
}
