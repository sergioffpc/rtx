package rtx_test

import (
	"fmt"
	"math"
	"sergioffpc/rtx/pkg/rtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPerspectiveCameraGenerateRay constructing camera rays.
func TestPerspectiveCameraGenerateRay(t *testing.T) {
	for _, tc := range []struct {
		camera   rtx.PerspectiveCamera
		x, y     int
		expected rtx.Ray
	}{
		// Ray through the center of the canvas.
		{camera: rtx.NewPerspectiveCamera(201, 101, math.Pi/2), x: 100, y: 50, expected: rtx.Ray{O: rtx.Point3{}, D: rtx.Vector3{0, 0, -1}, TMax: math.MaxFloat64}},
		// Ray through a corner of the canvas.
		{camera: rtx.NewPerspectiveCamera(201, 101, math.Pi/2), x: 0, y: 0, expected: rtx.Ray{O: rtx.Point3{}, D: rtx.Vector3{0.6651864261194507, 0.3325932130597254, -0.6685123582500481}, TMax: math.MaxFloat64}},
	} {
		t.Run("", func(t *testing.T) {
			got := tc.camera.GenerateRay(tc.x, tc.y)
			assert.True(t, rtx.Ray.Eq(tc.expected, got), fmt.Sprintf("expected: %v, got: %v", tc.expected, got))
		})
	}
}
