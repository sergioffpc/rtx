package sampler

import (
	"math/rand"
	"sergioffpc/rtx/pkg/rtx/cgmath"
)

func Get2D() cgmath.Point2 {
	return cgmath.Point2{X: rand.Float64(), Y: rand.Float64()}
}
