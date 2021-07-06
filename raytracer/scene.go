package raytracer


type Scene struct {
	Objects []Object
}


func (s *Scene) AddObject(obj Object) {
	s.Objects = append(s.Objects, obj)
}


func (s Scene) Hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, object := range s.Objects {
		if object.hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}

