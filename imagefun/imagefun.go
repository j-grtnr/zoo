package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strings"
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
	widthDefault    = 200
	heightDefault   = 200
	fileNameDefault = "image.png"
)

var (
	colors = map[int]*color.RGBA{
		// RGBa: RED, GREEN, BLUE, ALPHA ch. (transparency)
		0:  {R: 100, G: 200, B: 200, A: 0xff},
		1:  {R: 70, G: 70, B: 21, A: 0xff},
		2:  {R: 207, G: 70, B: 110, A: 0xff},
		3:  {R: 78, G: 70, B: 207, A: 0xff},
		4:  {R: 207, G: 205, B: 70, A: 0xff},
		5:  {R: 177, G: 37, B: 180, A: 0xff},
		6:  {R: 225, G: 39, B: 232, A: 0xff},
		7:  {R: 255, G: 0, B: 0, A: 0xff},
		8:  {R: 0, G: 255, B: 0, A: 0xff},
		9:  {R: 0, G: 0, B: 255, A: 0xff},
		10: {R: 255, G: 128, B: 0, A: 0xff},
		11: {R: 255, G: 0, B: 127, A: 0xff},
	}
	err   error
	lines []string
)

type imgOptions struct {
	width    int
	height   int
	colors   map[int]*color.RGBA
	fileName string
	content  [][]*color.RGBA
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	var w *int = flag.Int("w", widthDefault, "image width in pixels")
	var h *int = flag.Int("h", heightDefault, "image height in pixels")
	var fn *string = flag.String("f", fileNameDefault, "file name")
	var isFlag *bool = flag.Bool("flag", false, "draw flag image based on user input")
	flag.Parse()

	var options imgOptions = imgOptions{
		width:    *w,
		height:   *h,
		colors:   colors,
		fileName: *fn,
	}

	// create an empty image
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{options.width, options.height}})

	if *isFlag {
		err, lines = read_lines()

		if err != nil {
			log.Fatal(err)
		}
		//process input
		inputArrayWithColorRepresentation := GetArray(lines)
		inputArray := PutColors(options, inputArrayWithColorRepresentation)
		inputArrayUniform := modifyUniform(options, inputArray)
		options.content = scaleContent(options, inputArrayUniform)
	}

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

func draw(img *image.RGBA, options imgOptions) {
	// line by line go through the every pixel and fill that with some random color
	var asFlag bool = (options.content != nil)
	var pixelColor *color.RGBA

	for x := 0; x < (options.width / 2); x++ {
		for y := 0; y < options.height; y++ {
			if asFlag {
				pixelColor = options.content[y][x]
				img.Set(x, y, pixelColor)
				pixelColor = options.content[y][options.width-x-1]
				img.Set(options.width-x-1, y, pixelColor)
				//
			} else {
				colorIndex := rand.Intn(len(options.colors))
				pixelColor = options.colors[colorIndex]
				// fill the left side firstly
				img.Set(x, y, pixelColor)
				// fill the right side to make the image symmetric
				img.Set(options.width-x-1, y, pixelColor)
			}
		}
	}
}

func read_lines() (error, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter flag representation: ")
	var lines []string
	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		line := scanner.Text()
		if len(line) != 0 {
			lines = append(lines, line)
		} else {
			// exit if user entered an empty string
			break
		}
	}
	return scanner.Err(), lines
}

func GetArray(lines []string) [][]int {
	// convert lines to 2D array with int

	var array [][]int
	var stack []string
	var colorRepresenter int = -1
	var colorMap = make(map[string]int)

	for x := 0; x < len(lines); x++ {
		elements := strings.Split(lines[x], "")
		var mask []int
		for _, element := range elements {
			if !IsIn(element, stack) {
				colorRepresenter++
				colorMap[element] = colorRepresenter
				stack = append(stack, element)
			}
			mask = append(mask, colorMap[element])
		}
		array = append(array, mask)
	}
	return array
}

func PutColors(options imgOptions, colorReprArray [][]int) (colorArray [][]*color.RGBA) {
	// shuffle order of colors to pick
	colorPicker := make([]int, len(options.colors))
	for i := range colorPicker {
		colorPicker[i] = i
	}
	rand.Shuffle(len(colorPicker), func(i, j int) {
		colorPicker[i], colorPicker[j] = colorPicker[j], colorPicker[i]
	})

	for _, values := range colorReprArray {
		var row []*color.RGBA
		for _, value := range values {
			row = append(row, options.colors[colorPicker[value]%len(options.colors)])
		}
		colorArray = append(colorArray, row)
	}
	return colorArray
}

func modifyUniform(options imgOptions, inputArray [][]*color.RGBA) (scaledArray [][]*color.RGBA) {
	//correct dimensions to be symmetric

	var maxWidth int
	for _, row := range inputArray {
		if maxWidth < len(row) {
			maxWidth = len(row)
		}
	}

	//fill array to uniform dimensions
	var arrayUniform [][]*color.RGBA
	for ii, row := range inputArray {
		var rowUniform []*color.RGBA
		for jj := 0; jj < maxWidth; jj++ {
			if jj < len(row) {
				rowUniform = append(rowUniform, inputArray[ii][jj])
			} else {
				rowUniform = append(rowUniform, inputArray[ii][len(row)-1])
			}
		}
		arrayUniform = append(arrayUniform, rowUniform)
	}
	return arrayUniform
}

func scaleContent(options imgOptions, inputArray [][]*color.RGBA) (scaledArray [][]*color.RGBA) {

	var inputWidth int = len(inputArray[0])
	var wScaleFactor int = options.width / inputWidth
	var hScaleFactor int = options.height / len(inputArray)

	for hh := 0; hh < options.height; hh++ {
		var scaledRow []*color.RGBA
		switch {
		case hh < len(inputArray)*hScaleFactor:
			for ww := 0; ww < options.width; ww++ {
				if ww < inputWidth*wScaleFactor {
					scaledRow = append(scaledRow, inputArray[hh/hScaleFactor][ww/wScaleFactor])
				} else {
					scaledRow = append(scaledRow, inputArray[hh/hScaleFactor][inputWidth-1])
				}
			}
		case hh >= len(inputArray)*hScaleFactor:
			scaledRow = scaledArray[hh-1]
		}
		scaledArray = append(scaledArray, scaledRow)
	}
	return scaledArray
}

func IsIn(element string, stack []string) (val bool) {
	for _, ee := range stack {
		if ee == element {
			val = true
			break
		}
	}
	return val
}
