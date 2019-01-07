package main

import (
	"fmt"
	"math"
)

type vec3 struct {
	e [3]float64
}

func newVec3() *vec3 {
	return &vec3{}
}

func newVec3From(e0, e1, e2 float64) *vec3 {
	return &vec3{e: [3]float64{e0, e1, e2}}
}

func (v *vec3) x() float64 { return v.e[0] }

func (v *vec3) y() float64 { return v.e[1] }

func (v *vec3) z() float64 { return v.e[2] }

func (v *vec3) r() float64 { return v.e[0] }

func (v *vec3) g() float64 { return v.e[1] }

func (v *vec3) b() float64 { return v.e[2] }

func negate(v *vec3) *vec3 {
	return newVec3From(-v.e[0], -v.e[1], -v.e[2])
}

func (v *vec3) at(i int) float64 { return v.e[i] }

func (v *vec3) length() float64 {
	return math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
}

func (v *vec3) squaredLength() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.e[0], v.e[1], v.e[2])
}

func (v *vec3) makeUnitVector() {
	var k = 1.0 / math.Sqrt(v.e[0]*v.e[0]+v.e[1]*v.e[1]+v.e[2]*v.e[2])
	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k
}

func add(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]+v2.e[0], v1.e[1]+v2.e[1], v1.e[2]+v2.e[2])
}

func sub(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
}

func mul(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]*v2.e[0], v1.e[1]*v2.e[1], v1.e[2]*v2.e[2])
}

func div(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]/v2.e[0], v1.e[1]/v2.e[1], v1.e[2]/v2.e[2])
}

func scalarMul(v *vec3, t float64) *vec3 {
	return newVec3From(v.e[0]*t, v.e[1]*t, v.e[2]*t)
}

func scalarDiv(v *vec3, t float64) *vec3 {
	return newVec3From(v.e[0]/t, v.e[1]/t, v.e[2]/t)
}

func dot(v1, v2 *vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func cross(v1, v2 *vec3) *vec3 {
	return newVec3From(
		v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		-(v1.e[0]*v2.e[2] - v1.e[2]*v2.e[0]),
		v1.e[0]*v2.e[1]-v1.e[1]*v2.e[0],
	)
}

func (v *vec3) add(v2 *vec3) {
	v.e[0] += v2.e[0]
	v.e[1] += v2.e[1]
	v.e[2] += v2.e[2]
}

func (v *vec3) mul(v2 *vec3) {
	v.e[0] *= v2.e[0]
	v.e[1] *= v2.e[1]
	v.e[2] *= v2.e[2]
}

func (v *vec3) div(v2 *vec3) {
	v.e[0] /= v2.e[0]
	v.e[1] /= v2.e[1]
	v.e[2] /= v2.e[2]
}
func (v *vec3) sub(v2 *vec3) {
	v.e[0] -= v2.e[0]
	v.e[1] -= v2.e[1]
	v.e[2] -= v2.e[2]
}

func (v *vec3) scalarMul(t float64) {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t
}

func (v *vec3) scalarDiv(t float64) {
	var k = 1.0 / t

	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k
}

func unitVector(v *vec3) *vec3 {
	return scalarDiv(v, v.length())
}
