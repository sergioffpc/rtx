package color

import "sergioffpc/rtx/pkg/rtx/cgmath"

type Spectrum struct{ R, G, B float64 }

func (s Spectrum) Add(t Spectrum) Spectrum { return Spectrum{R: s.R + t.R, G: s.G + t.G, B: s.B + t.B} }

func (s *Spectrum) AddAssign(t Spectrum) {
	s.R += t.R
	s.G += t.G
	s.B += t.B
}

func (s Spectrum) Clamp(lo, hi float64) Spectrum {
	return Spectrum{R: cgmath.Clamp(s.R, lo, hi), G: cgmath.Clamp(s.G, lo, hi), B: cgmath.Clamp(s.B, lo, hi)}
}

func (s Spectrum) Div(t Spectrum) Spectrum { return Spectrum{R: s.R / t.R, G: s.G / t.G, B: s.B / t.B} }

func (s Spectrum) DivFloat(f float64) Spectrum { return Spectrum.MulFloat(s, 1/f) }

func (s Spectrum) Eq(t Spectrum) bool {
	return cgmath.EqualFloat(s.R, t.R) && cgmath.EqualFloat(s.G, t.G) && cgmath.EqualFloat(s.B, t.B)
}

func (s Spectrum) Lerp(t Spectrum, f float64) Spectrum {
	return Spectrum.Add(s.MulFloat(1-f), t.MulFloat(f))
}

func (s Spectrum) Mul(t Spectrum) Spectrum { return Spectrum{R: s.R * t.R, G: s.G * t.G, B: s.B * t.B} }

func (s Spectrum) MulFloat(f float64) Spectrum { return Spectrum{R: s.R * f, G: s.G * f, B: s.B * f} }

func (s Spectrum) Sub(t Spectrum) Spectrum { return Spectrum{R: s.R - t.R, G: s.G - t.G, B: s.B - t.B} }

func (s *Spectrum) SubAssign(t Spectrum) {
	s.R -= t.R
	s.G -= t.G
	s.B -= t.B
}
