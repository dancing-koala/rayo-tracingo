package main

type hitRecord struct {
	t      float64
	p      *vec3
	normal *vec3
}

type hitable interface {
	hit(r *ray, tMin, tMax float64, record *hitRecord) bool
}

type hitableList []hitable

func (h *hitableList) hit(r *ray, tMin, tMax float64, record *hitRecord) bool {
	tmpRec := &hitRecord{}
	hitAnything := false

	closestSoFar := tMax

	length := len(*h)

	for i := 0; i < length; i++ {
		if (*h)[i].hit(r, tMin, closestSoFar, tmpRec) {
			hitAnything = true
			closestSoFar = tmpRec.t

			record.normal = tmpRec.normal
			record.p = tmpRec.p
			record.t = tmpRec.t
		}
	}

	return hitAnything
}
