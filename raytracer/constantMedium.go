package raytracer


import (
	"math"
)


type ConstantMedium struct {
	Obj Object
	Density float64
	Colour Colour
}


func (cm ConstantMedium) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	rec1 := HitRecord{}
	rec2 := HitRecord{}

	if !cm.Obj.hit(r, math.Inf(-1), math.Inf(1), &rec1) {
		return false
	}

	if !cm.Obj.hit(r, rec1.T+0.0001, math.Inf(1), &rec2) {
		return false
	}

	if rec1.T < tMin {
		rec1.T = tMin
	}
	if rec2.T > tMax {
		rec2.T = tMax
	}

	if rec1.T >= rec2.T {
		return false
	}

	if rec1.T < 0 {
		rec1.T = 0
	}

	rayLength := r.Direction.Length()
	distanceInsideBoundary := (rec2.T - rec1.T) * rayLength
	negInvDensity := -1 / cm.Density
	hitDistance := negInvDensity * math.Log(RandomFloat64())

	if hitDistance > distanceInsideBoundary {
		return false
	}

	rec.T = rec1.T + hitDistance / rayLength
	rec.P = r.At(rec.T)
	rec.Normal = Vector{1, 0, 0}
	rec.FrontFace = true
	rec.Mat = Material{IsotropicScatter, cm.Colour, 0.0, false}

	return true
}