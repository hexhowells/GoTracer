package raytracer

type Object interface {
	hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool
}