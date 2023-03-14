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

type GradientTexture struct {
	Kd1, Kd2 Spectrum
}

func (t GradientTexture) D(p Point3, uv Point2) Spectrum {
	return Spectrum.Lerp(t.Kd1, t.Kd2, p.X-math.Floor(p.X))
}

type RingTexture struct {
	Kd1, Kd2 Spectrum
}

func (t RingTexture) D(p Point3, uv Point2) Spectrum {
	if math.Mod(math.Floor(math.Sqrt(p.X*p.X+p.Z*p.Z)), 2) == 0 {
		return t.Kd1
	}
	return t.Kd2
}

type CheckerTexture struct {
	Kd1, Kd2 Spectrum
}

func (t CheckerTexture) D(p Point3, uv Point2) Spectrum {
	if math.Mod(math.Floor(p.X)+math.Floor(p.Y)+math.Floor(p.Z), 2) == 0 {
		return t.Kd1
	}
	return t.Kd2
}
