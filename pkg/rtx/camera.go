package rtx

import "math"

type Camera struct {
	Width, Height int
}

func NewCamera(width, height int) Camera {
	return Camera{
		Width:  width,
		Height: height,
	}
}

func (c Camera) GenerateRay(x, y int) Ray {
	aspect := float64(c.Width) / float64(c.Height)
	dx := (-0.5 + (float64(x) / float64(c.Width))) * aspect
	dy := -0.5 + (float64(y) / float64(c.Height))
	return Ray{
		O:    Point3{},
		D:    Vector3{X: dx, Y: dy, Z: 1}.Normalize(),
		TMax: math.MaxFloat64,
	}
}
