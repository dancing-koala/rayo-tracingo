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

func (v *vec3) set(x, y, z float64) {
	v.e[0] = x
	v.e[1] = y
	v.e[2] = z
}

func (v *vec3) copyFrom(src *vec3) {
	v.e[0] = src.e[0]
	v.e[1] = src.e[1]
	v.e[2] = src.e[2]
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

func vec3Negate(v *vec3) *vec3 {
	return newVec3From(-v.e[0], -v.e[1], -v.e[2])
}

func vec3Add(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]+v2.e[0], v1.e[1]+v2.e[1], v1.e[2]+v2.e[2])
}

func vec3Sub(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
}

func vec3Mul(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]*v2.e[0], v1.e[1]*v2.e[1], v1.e[2]*v2.e[2])
}

func vec3Div(v1, v2 *vec3) *vec3 {
	return newVec3From(v1.e[0]/v2.e[0], v1.e[1]/v2.e[1], v1.e[2]/v2.e[2])
}

func vec3ScalarMul(v *vec3, t float64) *vec3 {
	return newVec3From(v.e[0]*t, v.e[1]*t, v.e[2]*t)
}

func vec3ScalarDiv(v *vec3, t float64) *vec3 {
	return newVec3From(v.e[0]/t, v.e[1]/t, v.e[2]/t)
}

func vec3Copy(src, dst *vec3) {
	dst.e[0] = src.e[0]
	dst.e[1] = src.e[1]
	dst.e[2] = src.e[2]
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

func (v *vec3) add(v2 *vec3) *vec3 {
	v.e[0] += v2.e[0]
	v.e[1] += v2.e[1]
	v.e[2] += v2.e[2]

	return v
}

func (v *vec3) mul(v2 *vec3) *vec3 {
	v.e[0] *= v2.e[0]
	v.e[1] *= v2.e[1]
	v.e[2] *= v2.e[2]

	return v
}

func (v *vec3) div(v2 *vec3) *vec3 {
	v.e[0] /= v2.e[0]
	v.e[1] /= v2.e[1]
	v.e[2] /= v2.e[2]

	return v
}
func (v *vec3) sub(v2 *vec3) *vec3 {
	v.e[0] -= v2.e[0]
	v.e[1] -= v2.e[1]
	v.e[2] -= v2.e[2]

	return v
}

func (v *vec3) scalarMul(t float64) *vec3 {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t

	return v
}

func (v *vec3) scalarDiv(t float64) *vec3 {
	var k = 1.0 / t

	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k

	return v
}

func unitVector(v *vec3) *vec3 {
	return vec3ScalarDiv(v, v.length())
}
