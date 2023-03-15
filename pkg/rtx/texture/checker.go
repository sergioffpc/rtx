package texture

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
)

type CheckerTexture struct {
	Kd1, Kd2 color.Spectrum
}

func (t CheckerTexture) D(p cgmath.Point3, uv cgmath.Point2) color.Spectrum {
	if math.Mod(math.Floor(p.X)+math.Floor(p.Y)+math.Floor(p.Z), 2) == 0 {
		return t.Kd1
	}
	return t.Kd2
}
