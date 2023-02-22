package rtx_test

import (
	"math"
	"sergioffpc/rtx/pkg/rtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestTransformRotateXTransform(t *testing.T) {
	for _, tc := range []struct {
		v        rtx.Vector3
		rx       rtx.Transform
		expected rtx.Vector3
	}{
		{v: rtx.Vector3{0, 1, 0}, rx: rtx.RotateXTransform(math.Pi / 2), expected: rtx.Vector3{0, 0, 1}},
		{v: rtx.Vector3{0, 1, 0}, rx: rtx.RotateXTransform(math.Pi / 4), expected: rtx.Vector3{0, math.Sqrt(2) / 2, math.Sqrt(2) / 2}},
		{v: rtx.Vector3{0, 1, 0}, rx: rtx.RotateXTransform(math.Pi / 4).Inverse(), expected: rtx.Vector3{0, math.Sqrt(2) / 2, -math.Sqrt(2) / 2}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Vector3.Eq(tc.expected, tc.v.Transform(tc.rx)))
		})
	}
}

func TestTransformRotateYTransform(t *testing.T) {
	for _, tc := range []struct {
		v        rtx.Vector3
		ry       rtx.Transform
		expected rtx.Vector3
	}{
		{v: rtx.Vector3{0, 0, 1}, ry: rtx.RotateYTransform(math.Pi / 2), expected: rtx.Vector3{1, 0, 0}},
		{v: rtx.Vector3{0, 0, 1}, ry: rtx.RotateYTransform(math.Pi / 4), expected: rtx.Vector3{math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2}},
		{v: rtx.Vector3{0, 0, 1}, ry: rtx.RotateYTransform(math.Pi / 4).Inverse(), expected: rtx.Vector3{-math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Vector3.Eq(tc.expected, tc.v.Transform(tc.ry)))
		})
	}
}

func TestTransformScaleTransform(t *testing.T) {
	for _, tc := range []struct {
		p        rtx.Point3
		s        rtx.Transform
		expected rtx.Point3
	}{
		{p: rtx.Point3{-4, 6, 8}, s: rtx.ScaleTransform(2, 3, 4), expected: rtx.Point3{-8, 18, 32}},
		{p: rtx.Point3{-4, 6, 8}, s: rtx.ScaleTransform(2, 3, 4).Inverse(), expected: rtx.Point3{-2, 2, 2}},
		{p: rtx.Point3{2, 3, 4}, s: rtx.ScaleTransform(-1, 1, 1), expected: rtx.Point3{-2, 3, 4}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Point3.Eq(tc.expected, tc.p.Transform(tc.s)))
		})
	}
}

func TestTransformTranslateTransform(t *testing.T) {
	for _, tc := range []struct {
		p        rtx.Point3
		s        rtx.Transform
		expected rtx.Point3
	}{
		{p: rtx.Point3{-3, 4, 5}, s: rtx.TranslateTransform(5, -3, 2), expected: rtx.Point3{2, 1, 7}},
		{p: rtx.Point3{-3, 4, 5}, s: rtx.TranslateTransform(5, -3, 2).Inverse(), expected: rtx.Point3{-8, 7, 3}},
	} {
		t.Run("", func(t *testing.T) {
			assert.True(t, rtx.Point3.Eq(tc.expected, tc.p.Transform(tc.s)))
		})
	}
}
