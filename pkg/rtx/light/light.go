package light

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type Light interface {
	Li(p cgmath.Point3) color.Spectrum
}
