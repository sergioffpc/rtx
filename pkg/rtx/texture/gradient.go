package texture

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type GradientTexture struct {
	Kd1, Kd2 color.Spectrum
}

func (t GradientTexture) D(p cgmath.Point3, uv cgmath.Point2) color.Spectrum {
	return color.Spectrum.Lerp(t.Kd1, t.Kd2, p.X-math.Floor(p.X))
}
