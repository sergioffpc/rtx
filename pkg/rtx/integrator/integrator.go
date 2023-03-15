package integrator

import (
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
	"sergioffpc/rtx/pkg/rtx/scene"
)

type Integrator interface {
	Render(scene *scene.Scene, ray cgmath.Ray) color.Spectrum
}
