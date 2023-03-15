package scene

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
	"sergioffpc/rtx/pkg/rtx/light"
	"sergioffpc/rtx/pkg/rtx/material"
	"sergioffpc/rtx/pkg/rtx/shape"
)

type GeometricPrimitive struct {
	Label         string
	Shape         shape.Shape
	Material      material.Material
	ObjectToWorld cgmath.Transform
	WorldToObject cgmath.Transform
}

func (g *GeometricPrimitive) intersect(ray cgmath.Ray) (ok bool, isect Interaction) {
	if hit, hP, hN, hT := g.Shape.Intersect(ray.Transform(g.WorldToObject)); hit {
		ok = true
		isect = Interaction{
			P:         hP.Transform(g.ObjectToWorld),
			N:         hN.Transform(g.ObjectToWorld),
			Wo:        ray.D.Neg(),
			T:         hT,
			Primitive: g,
		}
	}

	return ok, isect
}

type LightPrimitive struct {
	Label        string
	Light        light.Light
	LightToWorld cgmath.Transform
	WorldToLight cgmath.Transform
}

func (l LightPrimitive) Li(p cgmath.Point3) color.Spectrum {
	return l.Light.Li(p.Transform(l.WorldToLight))
}
