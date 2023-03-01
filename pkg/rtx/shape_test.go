package rtx_test

import (
	"math"
	"sergioffpc/rtx/pkg/rtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSphereShapeIntersect intersect rays with spheres.
func TestSphereShapeIntersect(t *testing.T) {
	for _, tc := range []struct {
		r   rtx.Ray
		hit bool
		t   float64
	}{
		// A ray intersects a sphere at two points.
		{r: rtx.Ray{rtx.Point3{0, 0, -5}, rtx.Vector3{0, 0, 1}, math.MaxFloat64}, hit: true, t: 4},
		// A ray intersects a sphere at a tangent.
		{r: rtx.Ray{rtx.Point3{0, 1, -5}, rtx.Vector3{0, 0, 1}, math.MaxFloat64}, hit: true, t: 5},
		// A ray misses a sphere.
		{r: rtx.Ray{rtx.Point3{0, 2, -5}, rtx.Vector3{0, 0, 1}, math.MaxFloat64}, hit: false, t: 0},
		/// A ray originates inside a sphere.
		{r: rtx.Ray{rtx.Point3{0, 0, 0}, rtx.Vector3{0, 0, 1}, math.MaxFloat64}, hit: true, t: 1},
		// A sphere is behind a ray.
		{r: rtx.Ray{rtx.Point3{0, 0, 5}, rtx.Vector3{0, 0, 1}, math.MaxFloat64}, hit: false, t: 0},
	} {
		t.Run("", func(t *testing.T) {
			hit, _, _, time := rtx.SphereShape{}.Intersect(tc.r)
			assert.Equal(t, tc.hit, hit)
			assert.Equal(t, tc.t, time)
		})
	}
}
