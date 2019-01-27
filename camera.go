package main

import (
	"math"
	"math/rand"
)

var unitForDisk = newVec3From(1.0, 1.0, 0.0)

func randomInUnitDisk(r *rand.Rand) *vec3 {
	// p := newVec3()
	// v := unitForDisk

	// for {
	// 	p.set(r.Float64(), r.Float64(), 0.0)
	// 	p.scalarMul(2.0)
	// 	p.sub(v)

	// 	if dot(p, p) >= 1.0 {
	// 		return p
	// 	}
	// }

	return unitForDisk
}

type camera struct {
	origin          *vec3
	lowerLeftCorner *vec3
	horizontal      *vec3
	vertical        *vec3
	u, v, w         *vec3
	lensRadius      float64
}

func newCamera(lookfrom, lookat, vup *vec3, vfov, aspect, aperture, focusDist float64) *camera {

	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	w := vec3Sub(lookfrom, lookat)
	w.makeUnitVector()

	u := cross(w, vup)
	u.makeUnitVector()

	v := cross(w, u)

	llc := vec3Sub(lookfrom, vec3ScalarMul(u, focusDist*halfWidth))
	llc.sub(vec3ScalarMul(v, focusDist*halfHeight))
	llc.sub(vec3ScalarMul(w, focusDist))

	return &camera{
		origin:          lookfrom,
		lowerLeftCorner: llc,
		horizontal:      vec3ScalarMul(u, 2*focusDist*halfWidth),
		vertical:        vec3ScalarMul(v, 2*focusDist*halfHeight),
		u:               u,
		v:               v,
		w:               w,
		lensRadius:      aperture / 2.0,
	}
}

func (c *camera) getRay(r *rand.Rand, s, t float64) *ray {

	rd := randomInUnitDisk(r)
	rd.scalarMul(c.lensRadius)

	offset := vec3ScalarMul(c.u, rd.x())
	offset.add(vec3ScalarMul(c.v, rd.y()))

	newOrigin := vec3Add(c.origin, offset)

	direction := vec3ScalarMul(c.horizontal, s)
	direction.add(vec3ScalarMul(c.vertical, t))
	direction.add(c.lowerLeftCorner)
	direction.sub(newOrigin)

	return newRayFrom(newOrigin, direction)
}
