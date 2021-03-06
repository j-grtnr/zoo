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

	// Create an empty image.
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{options.width, options.height}})

	// Create image content from user input if that flag is used.
	if *isFlag {
		fmt.Print("Enter flag representation: ")

		lines, err := read_lines()

		if err != nil {
			log.Fatal(err)
		}

		options.content = processInput(lines, options)
	}

	// Fill the image with colors.
	draw(img, options)

	// Create a file.
	f, err := os.Create(options.fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Write an image to a file.
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

// Prepare the user input to fit into the image and translate to colors.
func processInput(lines []string, options imgOptions) [][]*color.RGBA {
	inputArrayWithColorRepresentation := getArrayFromLines(lines)
	inputArray := putColors(options, inputArrayWithColorRepresentation)
	inputArrayUniform := modifyUniform(options, inputArray)
	return scaleContent(options, inputArrayUniform)
}

// Draw an image by setting colors for each pixel.
func draw(img *image.RGBA, options imgOptions) {
	var asFlag bool = (options.content != nil)
	if asFlag {
		drawContent(img, options)
	}
	if !asFlag {
		drawRandomColorSymmetricLayout(img, options)
	}
}

// Draws an image by choosing random colours for the left side from the given colours
// that are provided in imgOptions.
// Left side pixel colours are mirrored to the right side.
func drawRandomColorSymmetricLayout(img *image.RGBA, options imgOptions) {
	var pixelColor *color.RGBA
	for x := 0; x < (options.width / 2); x++ {
		for y := 0; y < options.height; y++ {
			colorIndex := rand.Intn(len(options.colors))
			pixelColor = options.colors[colorIndex]
			// fill the left side first
			img.Set(x, y, pixelColor)
			// fill the right side to make the image symmetric
			img.Set(options.width-x-1, y, pixelColor)
		}
	}
}

// Draws an image from imgOptions content.
func drawContent(img *image.RGBA, options imgOptions) {
	for x := 0; x < (options.width); x++ {
		for y := 0; y < options.height; y++ {
			img.Set(x, y, options.content[y][x])
		}
	}
}

// Reads user input line by line as strings.
// Finishes after an empty line was given.
func read_lines() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		// reads user input until \n by default
		scanner.Scan()
		// holds the string that was scanned
		line := scanner.Text()
		if len(line) != 0 {
			lines = append(lines, line)
		} else {
			// exit if user entered an empty string
			break
		}
	}
	return lines, scanner.Err()
}

// Convert lines to 2D array with Integers representing input characters.
func getArrayFromLines(lines []string) (array [][]int) {
	var colorRepresenter int = -1
	var colorMap = make(map[string]int)

	for _, line := range lines {
		elements := strings.Split(line, "")
		var mask []int
		for _, element := range elements {
			_, exists := colorMap[element]
			if !exists {
				colorRepresenter++
				colorMap[element] = colorRepresenter
			}
			mask = append(mask, colorMap[element])
		}
		array = append(array, mask)
	}
	return array
}

// Translates 2D Array of Integers into 2D Array of color RGBA values.
// Each distinct Integer gets a distinct color value until the given colours are used off.
// Colors are used multiple times if the given colours are not enough.
func putColors(options imgOptions, colorReprArray [][]int) (colorArray [][]*color.RGBA) {

	var colorPicker []int = getShuffleRange(len(options.colors))

	for _, values := range colorReprArray {
		var row []*color.RGBA
		for _, value := range values {
			row = append(row, options.colors[colorPicker[value%len(options.colors)]])
		}
		colorArray = append(colorArray, row)
	}
	return colorArray
}

// Correct dimensions of 2D Array to be uniform. Each row should have the same length,
// so do the columns.
func modifyUniform(options imgOptions, inputArray [][]*color.RGBA) (arrayUniform [][]*color.RGBA) {
	var maxWidth int
	for _, row := range inputArray {
		if maxWidth < len(row) {
			maxWidth = len(row)
		}
	}

	//fill array to uniform dimensions
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

// Scale the size and content of 2D Array according to width and height parameters given in imgOptions.
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

// Get a slice of Integers with shuffled values from 0 to number argument minus one.
func getShuffleRange(number int) []int {
	shuffled := make([]int, number)
	for i := range shuffled {
		shuffled[i] = i
	}
	rand.Shuffle(number, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
