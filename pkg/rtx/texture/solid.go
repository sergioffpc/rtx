package texture

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type SolidTexture struct {
	Kd color.Spectrum
}

func (t SolidTexture) D(p cgmath.Point3, uv cgmath.Point2) color.Spectrum {
	return t.Kd
}
