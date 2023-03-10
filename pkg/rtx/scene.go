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

func (p *GeometricPrimitive) intersect(r Ray) (ok bool, isect Interaction) {
	if hit, pos, n, t := p.Shape.Intersect(r.Transform(p.WorldToObject)); hit {
		ok = true
		isect = Interaction{
			P:         pos.Transform(p.ObjectToWorld),
			N:         n.Transform(p.ObjectToWorld).Normalize(),
			Wo:        r.D.Neg(),
			T:         t,
			Primitive: p,
		}
	}

	return ok, isect
}

type LightPrimitive struct {
	Light        Light
	LightToWorld Transform
	WorldToLight Transform
}

type Scene struct {
	Geometries []GeometricPrimitive
	Lights     []LightPrimitive
}

func (s Scene) Intersect(r Ray) (ok bool, nearest Interaction) {
	for _, g := range s.Geometries {
		p := g
		if hit, isect := p.intersect(r); hit {
			r.TMax = isect.T
			ok = true
			nearest = isect
		}
	}

	return ok, nearest
}
