package rtx

import "math"

type Material interface {
	F(p Point3, n Normal3, wo Vector3, uv Point2, wi Vector3, i Spectrum) Spectrum
}

type PhongMaterial struct {
	// Ks is a specular reflection constant.
	Ks float64
	// Kd is a diffuse reflection constant.
	Kd float64
	// Ka is an ambient reflection constant.
	Ka float64
	// Alpha is a shininess constant.
	Alpha float64
	// Tex is the surface texture.
	Tex Texture
}

func (m PhongMaterial) F(p Point3, n Normal3, wo Vector3, uv Point2, wi Vector3, i Spectrum) Spectrum {
	s := m.Tex.D(p, uv)
	f := Spectrum.MulFloat(i, m.Ka).Mul(s)

	wiDn := Vector3.Dot(wi, Vector3(n))
	if wiDn <= 0 {
		return Spectrum{}
	}

	d := m.Kd * wiDn
	f.AddAssign(Spectrum.MulFloat(i, d).Mul(s))

	woDr := Vector3.Dot(wo, Vector3.Reflect(wi.Neg(), n))
	if woDr > 0 {
		s := m.Ks * math.Pow(woDr, m.Alpha)
		f.AddAssign(Spectrum.MulFloat(i, s))
	}

	return f
}
