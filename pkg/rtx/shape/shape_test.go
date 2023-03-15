package shape_test

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/shape"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSphereShapeIntersect intersect rays with spheres.
func TestSphereShapeIntersect(t *testing.T) {
	for _, tc := range []struct {
		r   cgmath.Ray
		hit bool
		t   float64
	}{
		// A ray intersects a sphere at two points.
		{r: cgmath.Ray{O: cgmath.Point3{X: 0, Y: 0, Z: -5}, D: cgmath.Vector3{X: 0, Y: 0, Z: 1}, TMax: math.MaxFloat64}, hit: true, t: 4},
		// A ray intersects a sphere at a tangent.
		{r: cgmath.Ray{O: cgmath.Point3{X: 0, Y: 1, Z: -5}, D: cgmath.Vector3{X: 0, Y: 0, Z: 1}, TMax: math.MaxFloat64}, hit: true, t: 5},
		// A ray misses a sphere.
		{r: cgmath.Ray{O: cgmath.Point3{X: 0, Y: 2, Z: -5}, D: cgmath.Vector3{X: 0, Y: 0, Z: 1}, TMax: math.MaxFloat64}, hit: false, t: 0},
		/// A ray originates inside a sphere.
		{r: cgmath.Ray{O: cgmath.Point3{X: 0, Y: 0, Z: 0}, D: cgmath.Vector3{X: 0, Y: 0, Z: 1}, TMax: math.MaxFloat64}, hit: true, t: 1},
		// A sphere is behind a ray.
		{r: cgmath.Ray{O: cgmath.Point3{X: 0, Y: 0, Z: 5}, D: cgmath.Vector3{X: 0, Y: 0, Z: 1}, TMax: math.MaxFloat64}, hit: false, t: 0},
	} {
		t.Run("", func(t *testing.T) {
			hit, _, _, time := shape.SphereShape{}.Intersect(tc.r)
			assert.Equal(t, tc.hit, hit)
			assert.Equal(t, tc.t, time)
		})
	}
}
