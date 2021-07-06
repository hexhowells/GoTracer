package raytracer

type HitRecord struct {
	P Vector
	Normal Vector
	T float64
	FrontFace bool
	Mat Material
}


func (h *HitRecord) SetFaceNormal(r Ray, outwardNormal Vector) {
	h.FrontFace = outwardNormal.Dot(r.Direction) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.MultiplyN(-1.0)
	}
}