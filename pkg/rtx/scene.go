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
	// Uv is the two-dimensional texture coordinate.
	Uv Point2
	// Primitive points to the geometric primitive intersectect by the ray.
	Primitive *GeometricPrimitive
}

func (i Interaction) F(l LightPrimitive) Spectrum {
	return i.Primitive.Material.F(
		i.P.Transform(i.Primitive.WorldToObject),
		i.N.Transform(i.Primitive.WorldToObject),
		i.Wo,
		i.Uv,
		Point3.Sub(Point3{}.Transform(l.LightToWorld), i.P).Normalize(),
		l.Li(i.P),
	)
}

type GeometricPrimitive struct {
	Shape         Shape
	Material      Material
	ObjectToWorld Transform
	WorldToObject Transform
}

func (g *GeometricPrimitive) intersect(r Ray) (ok bool, isect Interaction) {
	if hit, hP, hN, hT := g.Shape.Intersect(r.Transform(g.WorldToObject)); hit {
		ok = true
		isect = Interaction{
			P:         hP.Transform(g.ObjectToWorld),
			N:         hN.Transform(g.ObjectToWorld),
			Wo:        r.D.Neg(),
			T:         hT,
			Primitive: g,
		}
	}

	return ok, isect
}

type LightPrimitive struct {
	Light        Light
	LightToWorld Transform
	WorldToLight Transform
}

func (l LightPrimitive) Li(p Point3) Spectrum {
	return l.Light.Li(p.Transform(l.WorldToLight))
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

func (s Scene) IntersectP(r Ray) bool {
	for _, g := range s.Geometries {
		if ok, _ := g.Shape.IntersectP(r.Transform(g.WorldToObject)); ok {
			return true
		}
	}
	return false
}
