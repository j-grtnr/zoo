package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	options = imgOptions{
		width:    200,
		height:   150,
		colors:   colors,
		fileName: "test_image.png",
	}
	inputLines     = []string{"1111", "%%%", "//"}
	colorReprArray = [][]int{{0, 0, 0, 0}, {1, 1, 1}, {2, 2}}
)

func TestGetArrayFromLines(t *testing.T) {
	arr := getArrayFromLines(inputLines)

	for ii := range colorReprArray {
		for jj := range colorReprArray[ii] {
			if arr[ii][jj] != colorReprArray[ii][jj] {
				t.Errorf("mismatch found:\n %v \n %v", arr, colorReprArray)
				break
			}
		}
	}
}

func TestPutColors(t *testing.T) {
	arr := putColors(options, colorReprArray)

	//check whether each value in array is represented in options.colors
	for ii, line := range arr {
		for jj, color := range line {
			var contained bool = false
			for _, refColor := range options.colors {

				if color == refColor {
					contained = true
				}
			}
			if !contained {
				t.Errorf("wrong type of value in array  at (%d, %d)", ii, jj)
			}
		}
	}
}

func TestModifyUniform(t *testing.T) {
	arr := putColors(options, colorReprArray)
	arrUniform := modifyUniform(options, arr)
	width := len(arrUniform[0])
	for _, line := range arrUniform {
		assert.Equal(t, len(line), width, "lengths of pixel rows should be equal")
	}
}

// Assert image size to input width and height
func TestScaleContent(t *testing.T) {
	arr := putColors(options, colorReprArray)
	arrUniform := modifyUniform(options, arr)
	arrScaled := scaleContent(options, arrUniform)

	assert.Equal(t, len(arrScaled), options.height, "number of pixel rows does not fit image height")
	for _, row := range arrScaled {
		assert.Equal(t, len(row), options.width, "length of pixel row does not fit image width")
	}
}
