package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

type result struct {
	x, y    int
	r, g, b uint8
}

type job struct {
	x, y  int
	cam   *camera
	world hitableList
}

const (
	nx          = 800
	ny          = 400
	ns          = 100.0
	workerCount = 8
	aperture    = 1.0
	fov         = 60
)

func main() {

	runtime.GOMAXPROCS(2)

	world := randomWorld()

	lookfrom := newVec3From(0.0, 3.0, 5.0)
	lookat := newVec3From(0.0, 1.0, 0.0)
	distToFocus := vec3Sub(lookfrom, lookat).length()
	vup := newVec3From(0.0, 1.0, 0.0)

	cam := newCamera(
		lookfrom,
		lookat,
		vup,
		fov,
		float64(nx)/float64(ny),
		aperture,
		distToFocus,
	)

	wg := &sync.WaitGroup{}
	wg.Add(nx * ny)

	jobsChan := make(chan job, nx*ny)
	resultsChan := make(chan result, workerCount)

	startWorkers(jobsChan, resultsChan)

	createJobs(jobsChan, cam, world)

	go finalize(wg, resultsChan)

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	for res := range resultsChan {
		wg.Done()
		img.Set(res.x, res.y, color.RGBA{res.r, res.g, res.b, 0xff})
	}

	f, err := os.Create("result.png")

	checkErr(err)
	png.Encode(f, img)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getColor(random *rand.Rand, r *ray, hitables hitableList, depth int) *vec3 {
	record := &hitRecord{}

	if hitables.hit(r, 0.001, math.MaxFloat64, record) {

		scatteredRay := newRay()
		attenuation := newVec3()

		if depth < 50 && record.itemMaterial.scatter(random, r, record, attenuation, scatteredRay) {
			attenuation.mul(getColor(random, scatteredRay, hitables, depth+1))
			return attenuation
		}

		return attenuation
	}

	unitDirection := unitVector(r.direction())
	t := 0.5 * (unitDirection.y() + 1.0)

	final := newVec3From(1.0, 1.0, 1.0)
	final.scalarMul(1.0 - t)

	b := newVec3From(0.5, 0.7, 1.0)
	b.scalarMul(t)

	final.add(b)

	return final
}

func randomWorld() hitableList {

	world := make(hitableList, 1)
	world[0] = newSphere(
		newVec3From(0.0, -1000.0, 0.0),
		1000.0,
		newLambertian(newVec3From(0.1, 0.4, 0.2)),
	)

	var object hitable

	refVec := newVec3From(0.0, 0.0, 0.0)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for a := -5; a < 5; a++ {
		for b := -5; b < 5; b++ {
			materialChoice := r.Float64()
			center := newVec3From(float64(a)+0.9*r.Float64(), 0.2, float64(b)+0.9*r.Float64())

			if vec3Sub(center, refVec).length() > 0.9 {

				switch {
				case materialChoice < 0.8:
					object = newSphere(
						center,
						0.2,
						newLambertian(
							newVec3From(
								r.Float64()*r.Float64(),
								r.Float64()*r.Float64(),
								r.Float64()*r.Float64(),
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
								0.5*(1+r.Float64()),
								0.5*(1+r.Float64()),
								0.5*(1+r.Float64()),
							),
							0.5*r.Float64(),
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

	object = newSphere(newVec3From(-3.0, 1.0, 0.0), 1.0, newLambertian(newVec3From(0.4, 0.2, 0.1)))
	world = append(world, object)

	object = newSphere(newVec3From(3.0, 1.0, 0.0), 1.0, newMetal(newVec3From(0.7, 0.6, 0.5), 0.0))
	world = append(world, object)

	return world
}

func startWorkers(jobsChan chan job, resultsChan chan result) {
	for n := 0; n < workerCount; n++ {
		go worker(n, jobsChan, resultsChan)
	}
}

func worker(id int, jobsChan chan job, resultsChan chan result) {
	c := newVec3()

	random := rand.New(rand.NewSource(time.Now().Unix()))

	for j := range jobsChan {
		c.reset()

		for k := 0; k < int(ns); k++ {
			u := (float64(j.x) + random.Float64()) / float64(nx)
			v := (float64(j.y) + random.Float64()) / float64(ny)

			r := j.cam.getRay(random, u, v)

			c.add(getColor(random, r, j.world, 0))
		}

		c.scalarDiv(ns)
		c.sqrt()

		resultsChan <- result{
			x: j.x,
			y: j.y,
			r: uint8(255.99 * c.at(0)),
			g: uint8(255.99 * c.at(1)),
			b: uint8(255.99 * c.at(2)),
		}
	}
}

func finalize(wg *sync.WaitGroup, resultsChan chan result) {
	wg.Wait()
	close(resultsChan)
}

func createJobs(jobsChan chan job, cam *camera, world hitableList) {
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			newJob := job{
				x:     i,
				y:     j,
				cam:   cam,
				world: world,
			}

			jobsChan <- newJob
		}
	}

	close(jobsChan)
}
