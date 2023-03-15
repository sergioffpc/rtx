package rtx

import "math"

type Integrator interface {
	Render(scene *Scene, ray Ray) Spectrum
}

type Whitted struct {
	MaxDepth int
}

func (w Whitted) Render(scene *Scene, ray Ray) Spectrum {
	return w.li(scene, ray, w.MaxDepth)
}

func (w Whitted) li(scene *Scene, ray Ray, depth int) Spectrum {
	if depth <= 0 {
		return Spectrum{}
	}

	ok, isect := scene.Intersect(ray)
	if !ok {
		return Spectrum{}
	}

	var li Spectrum
	for _, l := range scene.Lights {
		light := l
		if !w.litted(scene, isect, light) {
			continue
		}

		li.AddAssign(isect.F(light))
	}

	return li.Add(w.reflection(scene, isect, depth))
}

func (Whitted) litted(scene *Scene, isect Interaction, primitive LightPrimitive) bool {
	p := Point3{}.Transform(primitive.LightToWorld)
	wi := Point3.Sub(p, isect.P)
	return !scene.IntersectP(Ray{
		// Slighty bump above the surface in the direction of the
		// normal to prevent self-shadowing.
		O:    Point3.Add(isect.P, Point3(isect.N).MulFloat(Epsilon)),
		D:    wi.Normalize(),
		TMax: wi.Len(),
	})
}

func (w Whitted) reflection(scene *Scene, isect Interaction, depth int) Spectrum {
	rho := isect.Rho()
	if rho == 0 {
		return Spectrum{}
	}
	return w.li(scene, Ray{
		O:    Point3.Add(isect.P, Point3(isect.N).MulFloat(Epsilon)),
		D:    Vector3.Reflect(isect.Wo.Neg(), isect.N),
		TMax: math.MaxFloat64,
	}, depth-1).MulFloat(rho)
}
