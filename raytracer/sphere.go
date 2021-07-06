package raytracer

import (
	"math"
)

type Sphere struct {
	Center Vector
	Radius float64
	Mat Material
}


func (s Sphere) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	oc := r.Origin.Minus(s.Center)

	a := r.Direction.LengthSquared()
	half_b := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (-half_b - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-half_b + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	rec.Normal = rec.P.Minus(s.Center).DivideN(s.Radius)
	outwardNormal := rec.P.Minus(s.Center).DivideN(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = s.Mat

	return true
}