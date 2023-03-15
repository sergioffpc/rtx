package texture

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type Texture interface {
	D(p cgmath.Point3, uv cgmath.Point2) color.Spectrum
}
