package rtx

type Integrator interface {
	Render(scene *Scene, ray Ray) Spectrum
}

type Whitted struct{}

func (w Whitted) Render(scene *Scene, ray Ray) Spectrum {
	var li Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		for _, l := range scene.Lights {
			lightP := Point3{}.Transform(l.LightToWorld)

			hitLightV := Point3.Sub(lightP, isect.P)
			shadowR := Ray{
				// Slighty bump above the surface in the direction of the
				// normal to prevent self-shadowing.
				O:    Point3.Add(isect.P, Point3(isect.N).MulFloat(Epsilon)),
				D:    hitLightV.Normalize(),
				TMax: hitLightV.Len(),
			}
			if scene.IntersectP(shadowR) {
				continue
			}

			f := isect.F(l)
			li.AddAssign(f)
		}
	}

	return li
}
