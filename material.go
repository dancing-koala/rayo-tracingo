package main

type material interface {
	scatter(inputRay *ray, record *hitRecord, atternuation *vec3, scatteredRay *ray) bool
}

type lambertian struct {
	albedo *vec3
}

func newLambertian() *lambertian {
	return &lambertian{
		albedo: newVec3(),
	}
}

func newLambertianFrom(albedo *vec3) *lambertian {
	return &lambertian{
		albedo: albedo,
	}
}

func (l *lambertian) scatter(inputRay *ray, record *hitRecord, atternuation *vec3, scatteredRay *ray) bool {
	target := vec3Add(
		vec3Add(
			record.p,
			record.normal,
		),
		randomInUnitSphere(),
	)

	scatteredRay.a = record.p
	scatteredRay.b = vec3Sub(target, record.p)

	vec3Copy(l.albedo, atternuation)

	return true
}

type metal struct {
	albedo *vec3
	fuzz   float64
}

func newMetal() *metal {
	return &metal{
		albedo: newVec3(),
	}
}

func newMetalFrom(albedo *vec3, fuzz float64) *metal {

	if fuzz > 1.0 {
		fuzz = 1.0
	}

	return &metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

func reflect(v, n *vec3) *vec3 {
	return vec3Sub(
		v,
		vec3ScalarMul(
			n,
			2.0*dot(v, n),
		),
	)
}

func (m *metal) scatter(inputRay *ray, record *hitRecord, atternuation *vec3, scatteredRay *ray) bool {
	reflected := reflect(unitVector(inputRay.direction()), record.normal)
	scatteredRay.a = record.p
	scatteredRay.b = vec3Add(
		reflected,
		vec3ScalarMul(randomInUnitSphere(), m.fuzz),
	)
	vec3Copy(m.albedo, atternuation)

	return dot(scatteredRay.direction(), record.normal) > 0
}
