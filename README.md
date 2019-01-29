# Rayo Tracingo

A simple raytracer in Go.

![Rendering result](https://github.com/dancing-koala/rayo-tracingo/blob/master/result.png)


## Why

I wanted to learn about raytracing and improve my knowledge of the Go language.
This was a good project to learn about how to optimize Go code by doing less allocations, using concurrency & parallelism and discover some gotchas like using the `rand` package within multiple goroutines.


## AA Glitches

Below you can see some glitches that may happen if you mess up the anti-aliasing.

I think it is kind of trippy and cool.

![Glitch 1](https://github.com/dancing-koala/rayo-tracingo/blob/master/glitches/picture-20190112-1609.png)

![Glitch 2](https://github.com/dancing-koala/rayo-tracingo/blob/master/glitches/picture-20190112-1837.png)

![Glitch 3](https://github.com/dancing-koala/rayo-tracingo/blob/master/glitches/picture-20190113-0113.png)