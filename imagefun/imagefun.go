package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

// The agenda:

// struct
// errors
// panics (Aaaaaaaaa!)
// arrays & slices
// maps
// pointers

const (
	myConst = 42
)

type imgOptions struct {
	width    int
	height   int
	colors   map[int]*color.RGBA
	fileName string
}

var (
	colors = map[int]*color.RGBA{
		// RGBa: RED, GREEN, BLUE, ALPHA ch. (transparency)
		0: {R: 100, G: 200, B: 200, A: 0xff},
		1: {R: 70, G: 70, B: 21, A: 0xff},
		2: {R: 207, G: 70, B: 110, A: 0xff},
		3: {R: 78, G: 70, B: 207, A: 0xff},
		4: {R: 207, G: 205, B: 70, A: 0xff},
		5: {R: 177, G: 37, B: 180, A: 0xff},
	}
)

func init() {
	// We use random seed to take some seed for pseudo-random numbers algorithm
	rand.Seed(time.Now().UnixNano())
}

func main() {

	var w *int = flag.Int("w", 200, "image width in pixels")
	var h *int = flag.Int("h", 200, "image height in pixels")
	var fn *string = flag.String("f", "image.png", "file name")
	flag.Parse()

	var options imgOptions = imgOptions{
		width:    *w,
		height:   *h,
		colors:   colors,
		fileName: *fn,
	}

	// create an empty image
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{options.width, options.height}})
	// draw func that fill the image with some colors
	draw(img, options)

	// create a file
	f, err := os.Create(options.fileName)
	if err != nil {
		log.Fatal(err)
	}

	// write an image to a file
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

//func draw(img *image.RGBA, colors map[int]*color.RGBA) {
func draw(img *image.RGBA, options imgOptions) {
	// line by line go through the every pixel and fill that with some random color
	for x := 0; x < (options.width / 2); x++ {
		for y := 0; y < options.height; y++ {
			color := rand.Intn(5)
			// fill the left side firstly
			img.Set(x, y, options.colors[color])
			// fill the right side to make the image symmetric
			img.Set(options.width-x-1, y, options.colors[color])
		}
	}
}
