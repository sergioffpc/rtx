package camera

import "sergioffpc/rtx/pkg/rtx/cgmath"

type Camera interface {
	GenerateRay(x, y int) cgmath.Ray
	LookAt(from, to cgmath.Point3, up cgmath.Vector3)
}
