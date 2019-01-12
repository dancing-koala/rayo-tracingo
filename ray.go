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

func (r *ray) pointAtParam(t float64) *vec3 {
	return vec3Add(r.a, vec3ScalarMul(r.b, t))
}
