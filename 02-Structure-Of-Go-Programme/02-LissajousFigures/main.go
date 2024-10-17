package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.White,
	color.Black,
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// Create a new source and a new rand instance
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	f, err := os.Create("output.gif") // Create an output file
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f, r)
}

func lissajous(out io.Writer, r *rand.Rand) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 250
		nFrames = 64
		delay   = 8
	)

	freq := r.Float64() * 3.0 // Use the rand instance to get a random frequency
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
