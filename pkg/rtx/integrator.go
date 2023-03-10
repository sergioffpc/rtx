package rtx

type Integrator interface {
	Render(scene *Scene, ray Ray) Spectrum
}

type Whitted struct{}

func (w Whitted) Render(scene *Scene, ray Ray) Spectrum {
	var li Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		for _, l := range scene.Lights {
			li.AddAssign(isect.Shade(l))
		}
	}

	return li
}
