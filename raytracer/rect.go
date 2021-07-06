package raytracer

type XYRect struct {
	X0 float64
	X1 float64
	Y0 float64
	Y1 float64
	K float64
	Mat Material
}


type XZRect struct {
	X0 float64
	X1 float64
	Z0 float64
	Z1 float64
	K float64
	Mat Material
}


type YZRect struct {
	Y0 float64
	Y1 float64
	Z0 float64
	Z1 float64
	K float64
	Mat Material
}


func (rect XYRect) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	t := (rect.K-r.Origin.Z) / (r.Direction.Z)

	if t < tMin || t > tMax {
		return false
	}

	x := r.Origin.X + t*r.Direction.X
	y := r.Origin.Y + t*r.Direction.Y

	if x < rect.X0 || x > rect.X1 || y < rect.Y0 || y > rect.Y1 {
		return false
	}

	rec.T = t
	outwardNormal := Vector{0.0, 0.0, 1.0}
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = rect.Mat
	rec.P = r.At(t)

	return true
}


func (rect XZRect) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	t := (rect.K-r.Origin.Y) / (r.Direction.Y)

	if t < tMin || t > tMax {
		return false
	}

	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	if x < rect.X0 || x > rect.X1 || z < rect.Z0 || z > rect.Z1 {
		return false
	}

	rec.T = t
	outwardNormal := Vector{0.0, 1.0, 0.0}
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = rect.Mat
	rec.P = r.At(t)

	return true
}


func (rect YZRect) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	t := (rect.K-r.Origin.X) / (r.Direction.X)

	if t < tMin || t > tMax {
		return false
	}

	y := r.Origin.Y + t*r.Direction.Y
	z := r.Origin.Z + t*r.Direction.Z

	if y < rect.Y0 || y > rect.Y1 || z < rect.Z0 || z > rect.Z1 {
		return false
	}

	rec.T = t
	outwardNormal := Vector{1.0, 0.0, 0.0}
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = rect.Mat
	rec.P = r.At(t)

	return true
}