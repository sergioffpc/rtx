package integrator

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
	"sergioffpc/rtx/pkg/rtx/scene"
)

type Whitted struct {
	MaxDepth int
}

func (w Whitted) Render(scene *scene.Scene, ray cgmath.Ray) color.Spectrum {
	return w.li(scene, ray, w.MaxDepth)
}

func (w Whitted) li(scene *scene.Scene, ray cgmath.Ray, depth int) color.Spectrum {
	if depth <= 0 {
		return color.Spectrum{}
	}

	ok, isect := scene.Intersect(ray)
	if !ok {
		return color.Spectrum{}
	}

	var li color.Spectrum
	for _, l := range scene.Lights {
		light := l
		if !w.litted(scene, isect, light) {
			continue
		}

		li.AddAssign(isect.F(light))
	}

	return li.Add(w.reflection(scene, isect, depth))
}

func (Whitted) litted(scene *scene.Scene, isect scene.Interaction, primitive scene.LightPrimitive) bool {
	p := cgmath.Point3{}.Transform(primitive.LightToWorld)
	wi := cgmath.Point3.Sub(p, isect.P)
	return !scene.IntersectP(cgmath.Ray{
		// Slighty bump above the surface in the direction of the
		// normal to prevent self-shadowing.
		O:    cgmath.Point3.Add(isect.P, cgmath.Point3(isect.N).MulFloat(cgmath.Epsilon)),
		D:    wi.Normalize(),
		TMax: wi.Len(),
	})
}

func (w Whitted) reflection(scene *scene.Scene, isect scene.Interaction, depth int) color.Spectrum {
	rho := isect.Rho()
	if rho == 0 {
		return color.Spectrum{}
	}
	return w.li(scene, cgmath.Ray{
		O:    cgmath.Point3.Add(isect.P, cgmath.Point3(isect.N).MulFloat(cgmath.Epsilon)),
		D:    cgmath.Vector3.Reflect(isect.Wo.Neg(), isect.N),
		TMax: math.MaxFloat64,
	}, depth-1).MulFloat(rho)
}
