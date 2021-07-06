package raytracer


type Ray struct {
	Origin Vector
	Direction Vector
}


func (r *Ray) At(t float64) Vector {
	return r.Origin.Add(r.Direction.MultiplyN(t))
}