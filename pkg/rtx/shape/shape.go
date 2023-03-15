package shape

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
)

type Shape interface {
	Intersect(r cgmath.Ray) (bool, cgmath.Point3, cgmath.Normal3, float64)
	IntersectP(r cgmath.Ray) (bool, float64)
}

type SphereShape struct{}

func (s SphereShape) Intersect(ray cgmath.Ray) (bool, cgmath.Point3, cgmath.Normal3, float64) {
	if ok, t := s.IntersectP(ray); ok {
		p := ray.Position(t)

		// Since this is a unit sphere, the normal vector will be normalized by
		// default for any point on the surface.
		n := cgmath.Normal3(p)

		return true, p, n, t
	}
	return false, cgmath.Point3{}, cgmath.Normal3{}, 0
}

func (SphereShape) IntersectP(ray cgmath.Ray) (bool, float64) {
	a := cgmath.Vector3.Dot(ray.D, ray.D)
	b := 2 * cgmath.Vector3.Dot(ray.D, cgmath.Vector3(ray.O))
	c := cgmath.Vector3.Dot(cgmath.Vector3(ray.O), cgmath.Vector3(ray.O)) - 1

	switch ok, t0, t1 := cgmath.QuadraticSolver(a, b, c); {
	case ok && t0 > 0 && t0 < ray.TMax:
		return true, t0
	case ok && t1 > 0 && t1 < ray.TMax:
		return true, t1
	default:
		return false, 0
	}
}

type PlaneShape struct{}

func (p PlaneShape) Intersect(ray cgmath.Ray) (bool, cgmath.Point3, cgmath.Normal3, float64) {
	if ok, t := p.IntersectP(ray); ok {
		return true, ray.Position(t), cgmath.Normal3{X: 0, Y: 1, Z: 0}, t
	}
	return false, cgmath.Point3{}, cgmath.Normal3{}, 0
}

func (PlaneShape) IntersectP(ray cgmath.Ray) (bool, float64) {
	if math.Abs(ray.D.Y) < cgmath.Epsilon {
		return false, 0
	}
	t := -ray.O.Y / ray.D.Y
	return t > 0 && t < ray.TMax, t
}
