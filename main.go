package main

import (
	"fmt"
	"os"
)

func main() {

	nx := 200
	ny := 100

	data := getPPMHeader(nx, ny)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			v := newVec3From(float64(i)/float64(nx), float64(j)/float64(ny), float64(0.2))

			ir := int(255.99 * v.at(0))
			ig := int(255.99 * v.at(1))
			ib := int(255.99 * v.at(2))

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
