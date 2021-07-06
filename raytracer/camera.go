package raytracer

import (
	"math"
)

type Camera struct {
	AspectRatio float64
	ViewportHeight float64
	ViewportWidth float64
	FocalLength float64
	Origin Vector
	Horizontal Vector
	Vertical Vector
	LowerLeftCorner Vector
	LensRadius float64
	U Vector
	V Vector
	W Vector
}


func NewCamera(lookFrom Vector, lookAt Vector, viewUp Vector, vFov float64, aspectRatio float64, aperture float64, focusDist float64) Camera {
	cam := Camera{}

	theta := DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)

	cam.AspectRatio = aspectRatio
	cam.ViewportHeight = h * 2.0
	cam.ViewportWidth = cam.AspectRatio * cam.ViewportHeight
	
	cam.W = lookFrom.Minus(lookAt).UnitVector()
	cam.U = viewUp.Cross(cam.W).UnitVector()
	cam.V = cam.W.Cross(cam.U)

	cam.Origin = lookFrom
	cam.Horizontal = cam.U.MultiplyN(cam.ViewportWidth).MultiplyN(focusDist)
	cam.Vertical = cam.V.MultiplyN(cam.ViewportHeight).MultiplyN(focusDist)
	cam.LowerLeftCorner = cam.Origin.Minus(cam.Horizontal.DivideN(2.0)).Minus(cam.Vertical.DivideN(2)).Minus(cam.W.MultiplyN(focusDist))

	cam.LensRadius = aperture / 2.0

	return cam
}


func (cam Camera) GetRay(u float64, v float64) Ray {
	rd := RandomInUnitDisk().MultiplyN(cam.LensRadius)
	offset := cam.U.MultiplyN(rd.X).Add(cam.V.MultiplyN(rd.Y))

	direction := cam.LowerLeftCorner.Add(cam.Horizontal.MultiplyN(u)).Add(cam.Vertical.MultiplyN(v)).Minus(cam.Origin).Minus(offset)

	return Ray{cam.Origin.Add(offset), direction}
}