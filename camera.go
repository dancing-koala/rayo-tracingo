package main

type camera struct {
	origin          *vec3
	lowerLeftCorner *vec3
	horizontal      *vec3
	vertical        *vec3
}

func newCamera() *camera {
	return &camera{
		origin:          newVec3(),
		lowerLeftCorner: newVec3From(-2.0, -1.0, -1.0),
		horizontal:      newVec3From(4.0, 0.0, 0.0),
		vertical:        newVec3From(0.0, 2.0, 0.0),
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
