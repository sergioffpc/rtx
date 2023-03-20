package camera_test

import (
	"fmt"
	"math"
	"sergioffpc/rtx/pkg/rtx/camera"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPerspectiveCameraGenerateRay constructing camera rays.
func TestPerspectiveCameraGenerateRay(t *testing.T) {
	for _, tc := range []struct {
		camera   camera.PerspectiveCamera
		x, y     int
		expected cgmath.Ray
	}{
		// Ray through the center of the canvas.
		{camera: camera.NewPerspectiveCamera(201, 101, math.Pi/2), x: 100, y: 50, expected: cgmath.Ray{
			O:    cgmath.Point3{},
			D:    cgmath.Vector3{X: 0, Y: 0, Z: -1},
			TMax: math.MaxFloat64,
		}},
		// Ray through a corner of the canvas.
		{camera: camera.NewPerspectiveCamera(201, 101, math.Pi/2), x: 0, y: 0, expected: cgmath.Ray{
			O:    cgmath.Point3{},
			D:    cgmath.Vector3{X: 0.6651864261194507, Y: 0.3325932130597254, Z: -0.6685123582500481},
			TMax: math.MaxFloat64,
		}},
	} {
		t.Run("", func(t *testing.T) {
			got := tc.camera.GenerateRay(tc.x, tc.y, cgmath.Point2{X: 0.5, Y: 0.5})
			assert.True(t, cgmath.Ray.Eq(tc.expected, got), fmt.Sprintf("expected: %v, got: %v", tc.expected, got))
		})
	}
}
