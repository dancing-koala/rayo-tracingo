package main

type hitRecord struct {
	t            float64
	p            *vec3
	normal       *vec3
	itemMaterial material
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

			copyHitRecord(tmpRec, record)
		}
	}

	return hitAnything
}

func copyHitRecord(src, dst *hitRecord) {
	dst.normal = src.normal
	dst.p = src.p
	dst.t = src.t
	dst.itemMaterial = src.itemMaterial
}
