package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

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
		}

		return newVec3()
	}

	unitDirection := unitVector(r.direction())
	t := 0.5 * (unitDirection.y() + 1.0)

	return vec3Add(
		vec3ScalarMul(newVec3From(1.0, 1.0, 1.0), 1.0-t),
		vec3ScalarMul(newVec3From(0.5, 0.7, 1.0), t),
	)
}

func randomWorld() hitableList {

	world := make(hitableList, 1)
	world[0] = newSphere(
		newVec3From(0.0, -1000.0, 0.0),
		1000.0,
		newLambertian(newVec3From(0.5, 0.5, 0.5)),
	)

	refVec := newVec3From(4.0, 0.2, 0.0)

	var object hitable

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			materialChoice := rand.Float64()
			center := newVec3From(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if vec3Sub(center, refVec).length() > 0.9 {

				switch {
				case materialChoice < 0.8:
					object = newSphere(
						center,
						0.2,
						newLambertian(
							newVec3From(
								rand.Float64()*rand.Float64(),
								rand.Float64()*rand.Float64(),
								rand.Float64()*rand.Float64(),
							),
						),
					)

					world = append(world, object)
				case materialChoice < 0.95:
					object = newSphere(
						center,
						0.2,
						newMetal(
							newVec3From(
								0.5*(1+rand.Float64()),
								0.5*(1+rand.Float64()),
								0.5*(1+rand.Float64()),
							),
							0.5*rand.Float64(),
						),
					)

				default:
					object = newSphere(center, 0.2, newDielectric(1.5))
				}

			} else {
				object = nil
			}

			if object != nil {
				world = append(world, object)
			}

		}
	}

	object = newSphere(newVec3From(0.0, 1.0, 0.0), 1.0, newDielectric(1.5))
	world = append(world, object)

	object = newSphere(newVec3From(-4.0, 1.0, 0.0), 1.0, newLambertian(newVec3From(0.4, 0.2, 0.1)))
	world = append(world, object)

	object = newSphere(newVec3From(4.0, 1.0, 0.0), 0.5, newMetal(newVec3From(0.7, 0.6, 0.5), 0.0))
	world = append(world, object)

	return world
}

func main() {
	nx, ny := 200, 100
	ns := 100.0

	data := getPPMHeader(nx, ny)

	world := randomWorld()

	lookfrom := newVec3From(3, 3, 2)
	lookat := newVec3From(0.0, 0.0, -1.0)
	distToFocus := vec3Sub(lookfrom, lookat).length()
	aperture := 2.0

	cam := newCamera(
		lookfrom,
		lookat,
		newVec3From(0.0, 1.0, 0.0),
		20,
		float64(nx)/float64(ny),
		aperture,
		distToFocus,
	)

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
