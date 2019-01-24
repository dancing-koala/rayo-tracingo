package main

import (
	"math"
	"math/rand"
)

func randomInUnitDisk(r *rand.Rand) *vec3 {
	var p *vec3

	v := newVec3From(1.0, 1.0, 0.0)

	for {
		p = vec3Sub(
			vec3ScalarMul(
				newVec3From(r.Float64(), r.Float64(), 0.0),
				2.0,
			),
			v,
		)

		if dot(p, p) >= 1.0 {
			return p
		}
	}
}

type camera struct {
	origin          *vec3
	lowerLeftCorner *vec3
	horizontal      *vec3
	vertical        *vec3
	u               *vec3
	v               *vec3
	w               *vec3
	lensRadius      float64
}

func newCamera(lookfrom, lookat, vup *vec3, vfov, aspect, aperture, focusDist float64) *camera {

	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	w := unitVector(vec3Sub(lookfrom, lookat))
	u := unitVector(cross(vup, w))
	v := cross(w, u)

	llc := vec3Sub(
		vec3Sub(
			vec3Sub(
				lookfrom,
				vec3ScalarMul(u, focusDist*halfWidth),
			),
			vec3ScalarMul(v, focusDist*halfHeight),
		),
		vec3ScalarMul(w, focusDist),
	)

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

	rd := vec3ScalarMul(randomInUnitDisk(r), c.lensRadius)
	offset := vec3Add(
		vec3ScalarMul(c.u, rd.x()),
		vec3ScalarMul(c.v, rd.y()),
	)

	return newRayFrom(
		vec3Add(c.origin, offset),
		vec3Sub(
			vec3Add(
				c.lowerLeftCorner,
				vec3Add(
					vec3ScalarMul(c.horizontal, s),
					vec3ScalarMul(c.vertical, t),
				),
			),
			vec3Add(
				c.origin,
				offset,
			),
		),
	)
}
