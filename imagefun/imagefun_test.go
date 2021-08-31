package main

import (
	"testing"
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

func TestGetArray(t *testing.T) {
	arr := getArray(inputLines)

	for ii := range colorReprArray {
		for jj := range colorReprArray[ii] {
			if arr[ii][jj] != colorReprArray[ii][jj] {
				t.Fail()
				t.Logf("mismatch found:\n %v \n %v", arr, colorReprArray)
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
				t.Fail()
				t.Logf("wrong type of value in array  at (%d, %d)", ii, jj)
			}
		}
	}
}

func TestModifyUniform(t *testing.T) {
	arr := putColors(options, colorReprArray)
	arrUniform := modifyUniform(options, arr)
	width := len(arrUniform[0])
	for _, line := range arrUniform {
		if len(line) != width {
			t.Fail()
			t.Logf("wrong length of line in array: is %d, should be %d", len(line), width)
		}
	}
}

func TestScaleContent(t *testing.T) {
	arr := putColors(options, colorReprArray)
	arrUniform := modifyUniform(options, arr)
	arrScaled := scaleContent(options, arrUniform)

	if len(arrScaled) != options.height {
		t.Fail()
		t.Logf("wrong height of array")
	}
	for ii, row := range arrScaled {
		if len(row) != options.width {
			t.Fail()
			t.Logf("wrong width of array in row %d", ii)
		}
	}

}
