package rtx

type Light interface {
	I() Spectrum
}

type PointLight struct {
	Ia Spectrum
}

func (l PointLight) I() Spectrum {
	return l.Ia
}
