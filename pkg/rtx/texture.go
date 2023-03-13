package rtx

import "math"

type Texture interface {
	D(p Point3, uv Point2) Spectrum
}

type SolidTexture struct {
	Kd Spectrum
}

func (t SolidTexture) D(p Point3, uv Point2) Spectrum {
	return t.Kd
}

type StripeTexture struct {
	Kd1, Kd2 Spectrum
}

func (t StripeTexture) D(p Point3, uv Point2) Spectrum {
	if math.Mod(math.Floor(p.X), 2) == 0 {
		return t.Kd1
	}
	return t.Kd2
}
