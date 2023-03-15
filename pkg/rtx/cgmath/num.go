package cgmath

import "math"

const Epsilon float64 = 2.2204460e-8

func Clamp(f, lo, hi float64) float64 {
	switch {
	case f < lo:
		return lo
	case f > hi:
		return hi
	default:
		return f
	}
}

func Degrees(rad float64) float64 { return (180 / math.Pi) * rad }

func EqualFloat(a, b float64) bool { return math.Abs(a-b) < Epsilon }

func QuadraticSolver(a, b, c float64) (ok bool, t0 float64, t1 float64) {
	switch discrim := b*b - 4*a*c; {
	case discrim > 0:
		sqrtDiscrim := math.Sqrt(discrim)
		ok = true
		t0 = (-b - sqrtDiscrim) / (2 * a)
		t1 = (-b + sqrtDiscrim) / (2 * a)
	case discrim == 0:
		ok = true
		t0 = -b / (2 * a)
		t1 = t0
	}

	return ok, math.Min(t0, t1), math.Max(t0, t1)
}

func Radians(deg float64) float64 { return (math.Pi / 180) * deg }
