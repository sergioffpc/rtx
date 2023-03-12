package rtx

import "math"

type Shape interface {
	Intersect(r Ray) (bool, Point3, Normal3, float64)
	IntersectP(r Ray) (bool, float64)
}

type SphereShape struct{}

func (s SphereShape) Intersect(ray Ray) (bool, Point3, Normal3, float64) {
	if ok, t := s.IntersectP(ray); ok {
		p := ray.Position(t)

		// Since this is a unit sphere, the normal vector will be normalized by
		// default for any point on the surface.
		n := Normal3(p)

		return true, p, n, t
	}
	return false, Point3{}, Normal3{}, 0
}

func (SphereShape) IntersectP(ray Ray) (bool, float64) {
	a := Vector3.Dot(ray.D, ray.D)
	b := 2 * Vector3.Dot(ray.D, Vector3(ray.O))
	c := Vector3.Dot(Vector3(ray.O), Vector3(ray.O)) - 1

	switch ok, t0, t1 := QuadraticSolver(a, b, c); {
	case ok && t0 > 0 && t0 < ray.TMax:
		return true, t0
	case ok && t1 > 0 && t1 < ray.TMax:
		return true, t1
	default:
		return false, 0
	}
}

type PlaneShape struct{}

func (p PlaneShape) Intersect(ray Ray) (bool, Point3, Normal3, float64) {
	if ok, t := p.IntersectP(ray); ok {
		return true, ray.Position(t), Normal3{X: 0, Y: 1, Z: 0}, t
	}
	return false, Point3{}, Normal3{}, 0
}

func (PlaneShape) IntersectP(ray Ray) (bool, float64) {
	if math.Abs(ray.D.Y) < Epsilon {
		return false, 0
	}
	t := -ray.O.Y / ray.D.Y
	return t > 0 && t < ray.TMax, t
}
