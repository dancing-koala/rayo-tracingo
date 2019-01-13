package main

import "math"

type camera struct {
	origin          *vec3
	lowerLeftCorner *vec3
	horizontal      *vec3
	vertical        *vec3
}

func newCamera(lookfrom, lookat, vup *vec3, vfov, aspect float64) *camera {

	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	w := unitVector(vec3Sub(lookfrom, lookat))
	u := unitVector(cross(vup, w))
	v := cross(w, u)

	llc := newVec3()
	llc.copyFrom(lookfrom)
	llc.sub(vec3ScalarMul(u, halfWidth)).sub(vec3ScalarMul(v, halfHeight)).sub(w)

	return &camera{
		origin:          lookfrom,
		lowerLeftCorner: llc,
		horizontal:      vec3ScalarMul(u, 2.0*halfWidth),
		vertical:        vec3ScalarMul(v, 2.0*halfHeight),
	}
}

func (c *camera) getRay(u, v float64) *ray {
	return newRayFrom(
		c.origin,
		vec3Sub(
			vec3Add(
				c.lowerLeftCorner,
				vec3Add(
					vec3ScalarMul(c.horizontal, u),
					vec3ScalarMul(c.vertical, v),
				),
			),
			c.origin,
		),
	)
}
