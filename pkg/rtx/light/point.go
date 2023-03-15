package light

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type PointLight struct {
	I color.Spectrum
}

func (l PointLight) Li(p cgmath.Point3) color.Spectrum {
	distanceSq := cgmath.Point3.DistanceSq(p, cgmath.Point3{})
	return color.Spectrum.DivFloat(l.I, distanceSq)
}
