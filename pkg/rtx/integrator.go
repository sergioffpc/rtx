package rtx

type Integrator interface {
	Render(scene *Scene, ray Ray) Spectrum
}

type Whitted struct{}

func (w Whitted) Render(scene *Scene, ray Ray) Spectrum {
	var li Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		for _, l := range scene.Lights {
			lP := Point3.Transform(Point3{}, l.LightToWorld)
			wi := Point3.Sub(lP, isect.P).Normalize()
			f := isect.Primitive.Material.F(isect.P, isect.N, wi, isect.Wo, l.Light.Li(), isect.T)
			li.AddAssign(f)
		}
	}

	return li
}
