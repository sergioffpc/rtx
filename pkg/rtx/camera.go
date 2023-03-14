package rtx

import "math"

type Camera interface {
	GenerateRay(x, y int) Ray
	LookAt(from, to Point3, up Vector3)
}

type PerspectiveCamera struct {
	halfWidth, halfHeight   float64
	pixelWidth, pixelHeight float64
	transform               Transform
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
		transform:   IdentityTransform(),
	}
}

func (c PerspectiveCamera) GenerateRay(x, y int) Ray {
	cx := c.halfWidth - (float64(x)+0.5)*c.pixelWidth
	cy := c.halfHeight - (float64(y)+0.5)*c.pixelHeight

	o := Point3{}.Transform(c.transform.Inverse())
	px := Point3{X: cx, Y: cy, Z: -1}.Transform(c.transform.Inverse())
	return Ray{
		O:    o,
		D:    Point3.Sub(px, o).Normalize(),
		TMax: math.MaxFloat64,
	}
}

func (c *PerspectiveCamera) LookAt(from, to Point3, up Vector3) {
	c.transform = LookAtTransform(from, to, up)
}
