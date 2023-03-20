package film

import (
	"sergioffpc/rtx/pkg/rtx/camera"
	"sergioffpc/rtx/pkg/rtx/integrator"
	"sergioffpc/rtx/pkg/rtx/scene"
)

type Film interface {
	Render(scene *scene.Scene, integrator integrator.Integrator, camera camera.Camera) error
}
