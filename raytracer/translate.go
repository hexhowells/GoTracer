package raytracer

import(
	"math"
)

type Translate struct {
	P Object
	Offset Vector
}


type RotateY struct {
	P Object
	Angle float64
	CosTheta float64
	SinTheta float64
}


func (t Translate) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	movedRay := Ray{r.Origin.Minus(t.Offset), r.Direction}

	if !t.P.hit(movedRay, tMin, tMax, rec) {
		return false
	}

	rec.P = rec.P.Add(t.Offset)
	rec.SetFaceNormal(movedRay, rec.Normal)

	return true
}


func NewRotateY(p Object, angle float64) RotateY {
	rt := RotateY{}

	rt.P = p
	rt.Angle = angle
	radians := DegreesToRadians(angle)
	rt.SinTheta = math.Sin(radians)
	rt.CosTheta = math.Cos(radians)

	return rt
}


func (ry RotateY) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	origin := r.Origin
	direction := r.Direction

	origin.X = ry.CosTheta*r.Origin.X - ry.SinTheta*r.Origin.Z
	origin.Z = ry.SinTheta*r.Origin.X + ry.CosTheta*r.Origin.Z

	direction.X = ry.CosTheta*r.Direction.X - ry.SinTheta*r.Direction.Z
	direction.Z = ry.SinTheta*r.Direction.X + ry.CosTheta*r.Direction.Z

	rotatedRay := Ray{origin, direction}

	if !ry.P.hit(rotatedRay, tMin, tMax, rec) {
		return false
	}

	p := rec.P
	normal := rec.Normal

	p.X = ry.CosTheta*rec.P.X + ry.SinTheta*rec.P.Z
	p.Z = -ry.SinTheta*rec.P.X + ry.CosTheta*rec.P.Z

	normal.X = ry.CosTheta*rec.Normal.X + ry.SinTheta*rec.Normal.Z
	normal.Z = -ry.SinTheta*rec.Normal.X + ry.CosTheta*rec.Normal.Z

	rec.P = p
	rec.SetFaceNormal(rotatedRay, normal)

	return true
}