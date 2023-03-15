package texture

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type RingTexture struct {
	Kd1, Kd2 color.Spectrum
}

func (t RingTexture) D(p cgmath.Point3, uv cgmath.Point2) color.Spectrum {
	if math.Mod(math.Floor(math.Sqrt(p.X*p.X+p.Z*p.Z)), 2) == 0 {
		return t.Kd1
	}
	return t.Kd2
}
