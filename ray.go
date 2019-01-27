package main

type ray struct {
	a *vec3
	b *vec3
}

func newRay() *ray {
	return &ray{}
}

func newRayFrom(a, b *vec3) *ray {
	return &ray{
		a: a,
		b: b,
	}
}

func (r *ray) origin() *vec3 {
	return r.a
}

func (r *ray) direction() *vec3 {
	return r.b
}

func (r *ray) set(origin, direction *vec3) {
	r.a = origin
	r.b = direction
}

func (r *ray) pointAtParam(t float64) *vec3 {
	p := vec3ScalarMul(r.b, t)
	p.add(r.a)

	return p
}
