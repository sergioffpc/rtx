package scene

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type Interaction struct {
	// P is where our ray intersects the object.
	P cgmath.Point3
	// N is the vector perpendicular to the surface at P.
	N cgmath.Normal3
	// Wo is the vector pointing from P to the origin of the ray.
	Wo cgmath.Vector3
	// T is the time of intersection.
	T float64
	// Uv is the two-dimensional texture coordinate.
	Uv cgmath.Point2
	// Primitive points to the geometric primitive intersectect by the ray.
	Primitive *GeometricPrimitive
}

func (i Interaction) F(primitive LightPrimitive) color.Spectrum {
	lightPos := cgmath.Point3{}.Transform(primitive.LightToWorld)
	wi := cgmath.Point3.Sub(lightPos, i.P).Normalize()
	li := primitive.Li(i.P)

	return i.Primitive.Material.F(
		i.P.Transform(i.Primitive.WorldToObject),
		i.N.Transform(i.Primitive.WorldToObject).Normalize(),
		i.Wo.Transform(i.Primitive.WorldToObject).Normalize(),
		i.Uv,
		wi.Transform(i.Primitive.WorldToObject).Normalize(),
		li,
	)
}

func (i Interaction) Rho() float64 {
	return i.Primitive.Material.Rho()
}
