package rtx

type Interaction struct {
	P         Point3
	N         Normal3
	Wo        Vector3
	T         float64
	Primitive *GeometricPrimitive
}

type GeometricPrimitive struct {
	Shape         *Shape
	WorldToObject Transform
	ObjectToWorld Transform
}

func (p *GeometricPrimitive) intersect(rW Ray) (ok bool, isect Interaction) {
	rO := rW.Transform(p.WorldToObject)
	if hit, pO, nO, t := (*p.Shape).Intersect(rO); hit {
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

type Scene struct {
	Geometries []GeometricPrimitive
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
