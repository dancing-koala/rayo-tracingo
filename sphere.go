package main

import (
	"math"
)

type sphere struct {
	center       *vec3
	radius       float64
	itemMaterial material
}

func newSphere() *sphere {
	return &sphere{
		center: newVec3(),
	}
}

func newSphereFrom(center *vec3, radius float64, itemMaterial material) *sphere {
	return &sphere{
		center:       center,
		radius:       radius,
		itemMaterial: itemMaterial,
	}
}

func (s *sphere) hit(r *ray, tMin, tMax float64, record *hitRecord) bool {
	oc := vec3Sub(r.origin(), s.center)
	a := dot(r.direction(), r.direction())
	b := dot(oc, r.direction())
	c := dot(oc, oc) - s.radius*s.radius

	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a

		if tMin < temp && temp < tMax {
			record.t = temp
			record.p = r.pointAtParam(temp)
			record.normal = vec3ScalarDiv(
				vec3Sub(record.p, s.center),
				s.radius,
			)
			record.itemMaterial = s.itemMaterial
			return true
		}

		temp = (-b + math.Sqrt(discriminant)) / a

		if tMin < temp && temp < tMax {
			record.t = temp
			record.p = r.pointAtParam(temp)
			record.normal = vec3ScalarDiv(
				vec3Sub(record.p, s.center),
				s.radius,
			)
			record.itemMaterial = s.itemMaterial
			return true
		}

	}

	return false
}
