package material

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type Material interface {
	F(p cgmath.Point3, n cgmath.Normal3, wo cgmath.Vector3, uv cgmath.Point2, wi cgmath.Vector3, i color.Spectrum) color.Spectrum
	Rho() float64
}
