package rtx

import "math"

type Normal3 struct{ X, Y, Z float64 }

func (n Normal3) Add(m Normal3) Normal3 { return Normal3{X: n.X + m.X, Y: n.Y + m.Y, Z: n.Z + m.Z} }

func (n *Normal3) AddAssign(m Normal3) {
	n.X += m.X
	n.Y += m.Y
	n.Z += m.Z
}

func (n *Normal3) DivAssignFloat(f float64) { n.MulAssignFloat(f) }

func (n Normal3) DivFloat(f float64) Normal3 { return Normal3.MulFloat(n, 1/f) }

func (n Normal3) Dot(m Normal3) float64 { return n.X*m.X + n.Y*m.Y + n.Z*m.Z }

func (n Normal3) Eq(m Normal3) bool {
	return EqualFloat(n.X, m.X) && EqualFloat(n.Y, m.Y) && EqualFloat(n.Z, m.Z)
}

func (n Normal3) Len() float64 { return math.Sqrt(Normal3.Dot(n, n)) }

func (n *Normal3) MulAssignFloat(f float64) {
	n.X *= f
	n.Y *= f
	n.Z *= f
}

func (n Normal3) MulFloat(f float64) Normal3 { return Normal3{X: n.X * f, Y: n.Y * f, Z: n.Z * f} }

func (n Normal3) Neg() Normal3 { return Normal3.MulFloat(n, -1) }

func (n Normal3) Normalize() Normal3 { return Normal3.DivFloat(n, n.Len()) }

func (n Normal3) Sub(m Normal3) Normal3 { return Normal3{X: n.X - m.X, Y: n.Y - m.Y, Z: n.Z - m.Z} }

func (n *Normal3) SubAssign(m Normal3) {
	n.X -= m.X
	n.Y -= m.Y
	n.Z -= m.Z
}

type Point3 struct{ X, Y, Z float64 }

func (p Point3) Add(q Point3) Point3 { return Point3{X: p.X + q.X, Y: p.Y + q.Y, Z: p.Z + q.Z} }

func (p Point3) AddVector(v Vector3) Point3 { return Point3{X: p.X + v.X, Y: p.Y + v.Y, Z: p.Z + v.Z} }

func (p *Point3) AddAssign(q Point3) {
	p.X += q.X
	p.Y += q.Y
	p.Z += q.Z
}

func (p Point3) Distance(q Point3) float64 { return Point3.Sub(p, q).Len() }

func (p Point3) DistanceSq(q Point3) float64 { return Point3.Sub(p, q).LenSq() }

func (p *Point3) DivAssignFloat(f float64) { p.MulAssignFloat(f) }

func (p Point3) DivFloat(f float64) Point3 { return Point3.MulFloat(p, 1/f) }

func (p Point3) Eq(q Point3) bool {
	return EqualFloat(p.X, q.X) && EqualFloat(p.Y, q.Y) && EqualFloat(p.Z, q.Z)
}

func (p *Point3) MulAssignFloat(f float64) {
	p.X *= f
	p.Y *= f
	p.Z *= f
}

func (p Point3) MulFloat(f float64) Point3 { return Point3{X: p.X * f, Y: p.Y * f, Z: p.Z * f} }

func (p Point3) Neg() Point3 { return Point3.MulFloat(p, -1) }

func (p Point3) Sub(q Point3) Vector3 { return Vector3{X: p.X - q.X, Y: p.Y - q.Y, Z: p.Z - q.Z} }

func (p Point3) SubVector(v Vector3) Point3 { return Point3{X: p.X - v.X, Y: p.Y - v.Y, Z: p.Z - v.Z} }

type Ray struct {
	O    Point3
	D    Vector3
	TMax float64
}

func (r Ray) Eq(s Ray) bool {
	return r.O.Eq(s.O) && r.D.Eq(s.D) && r.TMax == s.TMax
}

func (r Ray) Position(t float64) Point3 { return Point3.AddVector(r.O, Vector3.MulFloat(r.D, t)) }

type Vector3 struct{ X, Y, Z float64 }

func (v Vector3) Add(u Vector3) Vector3 { return Vector3{X: v.X + u.X, Y: v.Y + u.Y, Z: v.Z + u.Z} }

func (v *Vector3) AddAssign(u Vector3) {
	v.X += u.X
	v.Y += u.Y
	v.Z += u.Z
}

func (v Vector3) Cross(u Vector3) Vector3 {
	return Vector3{X: v.Y*u.Z - v.Z*u.Y, Y: v.Z*u.X - v.X*u.Z, Z: v.X*u.Y - v.Y*u.X}
}

func (v *Vector3) DivAssignFloat(f float64) { v.MulAssignFloat(f) }

func (v Vector3) DivFloat(f float64) Vector3 { return Vector3.MulFloat(v, 1/f) }

func (v Vector3) Dot(u Vector3) float64 { return v.X*u.X + v.Y*u.Y + v.Z*u.Z }

func (v Vector3) Eq(u Vector3) bool {
	return EqualFloat(v.X, u.X) && EqualFloat(v.Y, u.Y) && EqualFloat(v.Z, u.Z)
}

func (v Vector3) Len() float64 { return math.Sqrt(v.LenSq()) }

func (v Vector3) LenSq() float64 { return Vector3.Dot(v, v) }

func (v *Vector3) MulAssignFloat(f float64) {
	v.X *= f
	v.Y *= f
	v.Z *= f
}

func (v Vector3) MulFloat(f float64) Vector3 { return Vector3{X: v.X * f, Y: v.Y * f, Z: v.Z * f} }

func (v Vector3) Neg() Vector3 { return Vector3.MulFloat(v, -1) }

func (v Vector3) Normalize() Vector3 { return Vector3.DivFloat(v, v.Len()) }

func (v Vector3) Reflect(n Normal3) Vector3 {
	f := 2 * Vector3.Dot(v, Vector3(n))
	return Vector3.Sub(v, Vector3.MulFloat(Vector3(n), f))
}

func (v Vector3) Sub(u Vector3) Vector3 { return Vector3{X: v.X - u.X, Y: v.Y - u.Y, Z: v.Z - u.Z} }

func (v *Vector3) SubAssign(u Vector3) {
	v.X -= u.X
	v.Y -= u.Y
	v.Z -= u.Z
}
