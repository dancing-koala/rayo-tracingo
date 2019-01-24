package main

import (
	"math"
	"math/rand"
)

func randomInUnitSphere() *vec3 {
	var p *vec3

	unitVec := newVec3From(1.0, 1.0, 1.0)

	for {
		p = vec3ScalarMul(
			vec3Sub(
				newVec3From(rand.Float64(), rand.Float64(), rand.Float64()),
				unitVec,
			),
			2.0,
		)

		if p.squaredLength() < 1.0 {
			return p
		}
	}
}

type material interface {
	scatter(inputRay *ray, record *hitRecord, attenuation *vec3, scatteredRay *ray) bool
}

type lambertian struct {
	albedo *vec3
}

func newLambertian(albedo *vec3) *lambertian {
	return &lambertian{
		albedo: albedo,
	}
}

func (l *lambertian) scatter(inputRay *ray, record *hitRecord, attenuation *vec3, scatteredRay *ray) bool {
	target := vec3Add(
		vec3Add(
			record.p,
			record.normal,
		),
		randomInUnitSphere(),
	)

	scatteredRay.a = record.p
	scatteredRay.b = vec3Sub(target, record.p)

	vec3Copy(l.albedo, attenuation)

	return true
}

type metal struct {
	albedo *vec3
	fuzz   float64
}

func newMetal(albedo *vec3, fuzz float64) *metal {

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

func (m *metal) scatter(inputRay *ray, record *hitRecord, attenuation *vec3, scatteredRay *ray) bool {
	reflected := reflect(unitVector(inputRay.direction()), record.normal)
	scatteredRay.set(
		record.p,
		vec3Add(
			reflected,
			vec3ScalarMul(randomInUnitSphere(), m.fuzz),
		),
	)
	vec3Copy(m.albedo, attenuation)

	return dot(scatteredRay.direction(), record.normal) > 0
}

type dielectric struct {
	refractIndex float64
}

func newDielectric(refractIndex float64) *dielectric {
	return &dielectric{
		refractIndex: refractIndex,
	}
}

func (d *dielectric) scatter(inputRay *ray, record *hitRecord, attenuation *vec3, scatteredRay *ray) bool {

	var niOverNt, cosine, reflectProb float64
	var outwardNormal *vec3
	reflected := reflect(inputRay.direction(), record.normal)
	refracted := newVec3()

	attenuation.set(1.0, 1.0, 1.0)

	if dot(inputRay.direction(), record.normal) > 0 {
		outwardNormal = vec3Negate(record.normal)
		niOverNt = d.refractIndex
		cosine = d.refractIndex * dot(inputRay.direction(), record.normal) / inputRay.direction().length()
	} else {
		outwardNormal = record.normal
		niOverNt = 1.0 / d.refractIndex
		cosine = -dot(inputRay.direction(), record.normal) / inputRay.direction().length()
	}

	if refract(inputRay.direction(), outwardNormal, niOverNt, refracted) {
		reflectProb = schlick(cosine, d.refractIndex)
	} else {
		reflectProb = 10
	}

	if reflectProb > rand.Float64() {
		scatteredRay.set(record.p, reflected)
	} else {
		scatteredRay.set(record.p, refracted)
	}

	return true
}

func refract(v, n *vec3, niOverNt float64, refracted *vec3) bool {
	uv := unitVector(v)
	dt := dot(uv, n)

	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)

	if discriminant > 0 {
		refracted.copyFrom(
			vec3Sub(
				vec3ScalarMul(
					vec3Sub(
						uv,
						vec3ScalarMul(n, dt),
					),
					niOverNt,
				),
				vec3ScalarMul(n, math.Sqrt(discriminant)),
			),
		)

		return true
	}

	return false
}

func schlick(cosine, refractIndex float64) float64 {
	r0 := (1.0 - refractIndex) / (1 + refractIndex)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1.0-cosine, 5.0)
}
