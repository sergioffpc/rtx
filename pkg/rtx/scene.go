package rtx

type Interaction struct {
	// P is where our ray intersects the object.
	P Point3
	// N is the vector perpendicular to the surface at P.
	N Normal3
	// Wo is the vector pointing from P to the origin of the ray.
	Wo Vector3
	// T is the time of intersection.
	T float64
	// Primitive points to the geometric primitive intersectect by the ray.
	Primitive *GeometricPrimitive
}

type GeometricPrimitive struct {
	Shape         Shape
	Material      Material
	ObjectToWorld Transform
	WorldToObject Transform
}

func (p *GeometricPrimitive) intersect(rW Ray) (ok bool, isect Interaction) {
	rO := rW.Transform(p.WorldToObject)
	if hit, pO, nO, t := p.Shape.Intersect(rO); hit {
		ok = true
		isect = Interaction{
			P:         pO.Transform(p.ObjectToWorld),
			N:         nO.Transform(p.ObjectToWorld).Normalize(),
			Wo:        rW.D.Neg(),
			T:         t,
			Primitive: p,
		}
	}

	return ok, isect
}

type LightPrimitive struct {
	Light         Light
	ObjectToWorld Transform
	WorldToObject Transform
}

type Scene struct {
	Geometries []GeometricPrimitive
	Lights     []LightPrimitive
}

func (s Scene) Intersect(rW Ray) (ok bool, nearest Interaction) {
	for _, e := range s.Geometries {
		g := e
		if hit, isect := g.intersect(rW); hit {
			rW.TMax = isect.T
			ok = true
			nearest = isect
		}
	}

	return ok, nearest
}
