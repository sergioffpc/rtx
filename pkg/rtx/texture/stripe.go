package texture

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type StripeTexture struct {
	Kd1, Kd2 color.Spectrum
}

func (t StripeTexture) D(p cgmath.Point3, uv cgmath.Point2) color.Spectrum {
	if math.Mod(math.Floor(p.X), 2) == 0 {
		return t.Kd1
	}
	return t.Kd2
}
