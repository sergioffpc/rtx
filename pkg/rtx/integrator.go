package rtx

type Integrator interface {
	Render(scene *Scene, ray Ray) Spectrum
}

type Whitted struct{}

func (w Whitted) Render(scene *Scene, ray Ray) Spectrum {
	var li Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		for _, l := range scene.Lights {
			light := l
			if w.isIlluminated(scene, isect, light) {
				li.AddAssign(isect.F(light))
			}
		}
	}

	return li
}

func (Whitted) isIlluminated(scene *Scene, isect Interaction, primitive LightPrimitive) bool {
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
