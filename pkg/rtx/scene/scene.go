package scene

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
)

type Scene struct {
	Geometries []GeometricPrimitive
	Lights     []LightPrimitive
}

func (s Scene) Intersect(ray cgmath.Ray) (ok bool, nearest Interaction) {
	for _, g := range s.Geometries {
		geometry := g
		if hit, isect := geometry.intersect(ray); hit {
			ray.TMax = isect.T
			ok = true
			nearest = isect
		}
	}

	return ok, nearest
}

func (s Scene) IntersectP(ray cgmath.Ray) bool {
	for _, g := range s.Geometries {
		geometry := g
		if ok, _ := geometry.Shape.IntersectP(ray.Transform(geometry.WorldToObject)); ok {
			return true
		}
	}
	return false
}
