package camera

import (
	"math"
	"sergioffpc/rtx/pkg/rtx/cgmath"
)

type PerspectiveCamera struct {
	halfWidth, halfHeight   float64
	pixelWidth, pixelHeight float64
	transform               cgmath.Transform
}

func NewPerspectiveCamera(width, height int, fov float64) PerspectiveCamera {
	halfWidth := math.Tan(fov / 2)
	halfHeight := halfWidth / (float64(width) / float64(height))

	pixelWidth := (halfWidth * 2) / float64(width)
	pixelHeight := (halfHeight * 2) / float64(height)

	return PerspectiveCamera{
		halfWidth:   halfWidth,
		halfHeight:  halfHeight,
		pixelWidth:  pixelWidth,
		pixelHeight: pixelHeight,
		transform:   cgmath.IdentityTransform(),
	}
}

func (c PerspectiveCamera) GenerateRay(x, y int) cgmath.Ray {
	cx := c.halfWidth - (float64(x)+0.5)*c.pixelWidth
	cy := c.halfHeight - (float64(y)+0.5)*c.pixelHeight

	o := cgmath.Point3{}.Transform(c.transform.Inverse())
	px := cgmath.Point3{X: cx, Y: cy, Z: -1}.Transform(c.transform.Inverse())
	return cgmath.Ray{
		O:    o,
		D:    cgmath.Point3.Sub(px, o).Normalize(),
		TMax: math.MaxFloat64,
	}
}

func (c *PerspectiveCamera) LookAt(from, to cgmath.Point3, up cgmath.Vector3) {
	c.transform = cgmath.LookAtTransform(from, to, up)
}
