package rtx

type Integrator interface {
	Render(scene *Scene, ray Ray) Spectrum
}

type Whitted struct{}

func (w Whitted) Render(scene *Scene, ray Ray) Spectrum {
	var lI Spectrum
	if ok, isect := scene.Intersect(ray); ok {
		for _, l := range scene.Lights {
			lP := Point3.Transform(Point3{}, l.ObjectToWorld)
			wi := Point3.Sub(lP, isect.P).Normalize()
			f := isect.Primitive.Material.F(isect.P, isect.N, wi, isect.Wo, l.Light.I(), isect.T)
			lI.AddAssign(f)
		}
	}

	return lI
}
