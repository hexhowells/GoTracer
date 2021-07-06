package raytracer


type Box struct {
	Sides []Object
	Min Vector
	Max Vector
}


func NewBox(p0 Vector, p1 Vector, mat Material) Box {
	box := Box{}
	box.Min = p0
	box.Max = p1

	box.Sides = append(box.Sides, XYRect{p0.X, p1.X, p0.Y, p1.Y, p1.Z, mat})
	box.Sides = append(box.Sides, XYRect{p0.X, p1.X, p0.Y, p1.Y, p0.Z, mat})

	box.Sides = append(box.Sides, XZRect{p0.X, p1.X, p0.Z, p1.Z, p1.Y, mat})
	box.Sides = append(box.Sides, XZRect{p0.X, p1.X, p0.Z, p1.Z, p0.Y, mat})

	box.Sides = append(box.Sides, YZRect{p0.Y, p1.Y, p0.Z, p1.Z, p1.X, mat})
	box.Sides = append(box.Sides, YZRect{p0.Y, p1.Y, p0.Z, p1.Z, p0.X, mat})

	return box
}


func (b Box) hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, side := range b.Sides {
		if side.hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}