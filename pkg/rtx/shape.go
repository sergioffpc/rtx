package rtx

type Shape interface {
	Intersect(r Ray) (bool, Point3, Normal3, float64)
}

type SphereShape struct{}

func (s SphereShape) Intersect(ray Ray) (bool, Point3, Normal3, float64) {
	a := Vector3.Dot(ray.D, ray.D)
	b := 2 * Vector3.Dot(ray.D, Vector3(ray.O))
	c := Vector3.Dot(Vector3(ray.O), Vector3(ray.O)) - 1

	// Since this is a unit sphere, the normal vector will be normalized by
	// default for any point on the surface.
	switch ok, t0, t1 := QuadraticSolver(a, b, c); {
	case ok && t0 > 0 && t0 < ray.TMax:
		p := ray.Position(t0)
		n := Normal3(p)
		return true, p, n, t0
	case ok && t1 > 0 && t1 < ray.TMax:
		p := ray.Position(t1)
		n := Normal3(p)
		return true, p, n, t1
	default:
		return false, Point3{}, Normal3{}, 0
	}
}
