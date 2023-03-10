package rtx

type Light interface {
	Li() Spectrum
}

type PointLight struct {
	I Spectrum
}

func (l PointLight) Li() Spectrum {
	return l.I
}
