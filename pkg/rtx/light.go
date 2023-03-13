package rtx

type Light interface {
	Li(p Point3) Spectrum
}

type PointLight struct {
	I Spectrum
}

func (l PointLight) Li(p Point3) Spectrum {
	distanceSq := Point3.DistanceSq(p, Point3{})
	return Spectrum.DivFloat(l.I, distanceSq)
}
