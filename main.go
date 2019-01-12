package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
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

func color(r *ray, hitables hitableList, depth int) *vec3 {
	record := &hitRecord{}

	if hitables.hit(r, 0.001, math.MaxFloat64, record) {

		scatteredRay := newRay()
		attenuation := newVec3()

		if depth < 50 && record.itemMaterial.scatter(r, record, attenuation, scatteredRay) {
			return vec3Mul(
				attenuation,
				color(scatteredRay, hitables, depth+1),
			)
		} else {
			return newVec3()
		}
	}

	unitDirection := unitVector(r.direction())
	t := 0.5 * (unitDirection.y() + 1.0)

	return vec3Add(
		vec3ScalarMul(newVec3From(1.0, 1.0, 1.0), 1.0-t),
		vec3ScalarMul(newVec3From(0.5, 0.7, 1.0), t),
	)
}

func main() {
	nx, ny := 200, 100
	ns := 50.0

	data := getPPMHeader(nx, ny)

	world := make(hitableList, 4)
	world[0] = newSphereFrom(newVec3From(0, 0, -1.0), 0.5, newLambertianFrom(newVec3From(0.1, 0.2, 0.5)))
	world[1] = newSphereFrom(newVec3From(0, -100.5, -1.0), 100, newLambertianFrom(newVec3From(0.8, 0.8, 0.0)))
	world[2] = newSphereFrom(newVec3From(1.0, 0, -1.0), 0.5, newMetalFrom(newVec3From(0.8, 0.6, 0.2), 1.0))
	world[3] = newSphereFrom(newVec3From(-1.0, 0, -1.0), 0.5, newDielectricFrom(1.5))

	cam := newCamera()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			c := newVec3()

			for k := 0; k < int(ns); k++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := cam.getRay(u, v)

				c.add(color(r, world, 0))
			}

			c.scalarDiv(ns)

			c = newVec3From(math.Sqrt(c.r()), math.Sqrt(c.g()), math.Sqrt(c.b()))

			ir := int(255.99 * c.at(0))
			ig := int(255.99 * c.at(1))
			ib := int(255.99 * c.at(2))

			data += fmt.Sprintf("%d %d %d\n", ir, ig, ib)
		}
	}

	writeFile([]byte(data))
}

func getPPMHeader(nx, ny int) string {
	return fmt.Sprintf("P3\n%d %d\n255\n", nx, ny)
}

func writeFile(data []byte) {
	file, err := os.Create("picture.ppm")

	checkErr(err)

	defer file.Close()

	n, err := file.Write(data)

	checkErr(err)

	fmt.Printf("Wrote %d bytes\n", n)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
