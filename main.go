package main

import (
	"fmt"
	"os"
)

func color(r *ray) *vec3 {
	unitDirection := unitVector(r.direction())
	t := 0.5 * (unitDirection.y() + 1.0)
	return add(
		scalarMul(newVec3From(1.0, 1.0, 1.0), 1.0-t),
		scalarMul(newVec3From(0.5, 0.7, 1.0), t),
	)
}

func main() {
	nx := 200
	ny := 100

	data := getPPMHeader(nx, ny)

	lowerLeftCorner := newVec3From(-2.0, -1.0, -1.0)
	horizontal := newVec3From(4.0, 0.0, 0.0)
	vertical := newVec3From(0.0, 2.0, 0.0)
	origin := newVec3()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			r := newRayFrom(
				origin,
				add(
					lowerLeftCorner,
					add(
						scalarMul(horizontal, u),
						scalarMul(vertical, v),
					),
				),
			)

			c := color(r)

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
