package material

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
	"sergioffpc/rtx/pkg/rtx/texture"
)

type PhongMaterial struct {
	// Ks is a specular reflection constant.
	Ks float64
	// Kd is a diffuse reflection constant.
	Kd float64
	// Ka is an ambient reflection constant.
	Ka float64
	// Kr is a reflective constant.
	Kr float64
	// Kt is a refractive constant.
	Kt float64
	// Shininess is a shininess constant.
	Shininess float64
	// Transparency is a transparency constant.
	Transparency float64
	// Tex is the surface texture.
	Tex texture.Texture
}

func (m PhongMaterial) F(p cgmath.Point3, n cgmath.Normal3, wo cgmath.Vector3, uv cgmath.Point2, wi cgmath.Vector3, i color.Spectrum) color.Spectrum {
	s := m.Tex.D(p, uv)
	f := color.Spectrum.MulFloat(i, m.Ka).Mul(s)
	wiDn := cgmath.Vector3.Dot(wi, cgmath.Vector3(n))
	if wiDn <= 0 {
		return color.Spectrum{}
	}

	d := m.Kd * wiDn
	f.AddAssign(color.Spectrum.MulFloat(i, d).Mul(s))

	woDr := cgmath.Vector3.Dot(wo, cgmath.Vector3.Reflect(wi.Neg(), n))
	if woDr > 0 {
		s := m.Ks * math.Pow(woDr, m.Shininess)
		f.AddAssign(color.Spectrum.MulFloat(i, s))
	}

	return f
}

func (m PhongMaterial) Rho() float64 { return m.Kr }
