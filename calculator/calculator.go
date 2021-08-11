package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		inputs    []string
		operandA  int
		operandB  int
		operation string
		result    int
		err       error
	)

	if len(os.Args) != 2 {
		log.Fatal("Wrong number of command line arguments. Example usage: ./calculator \"1 + 1\"")
	}
	inputs = strings.Fields(os.Args[1])
	operandA, err = strconv.Atoi(inputs[0])
	if err != nil {
		log.Fatal(err)
	}
	operation = inputs[1]
	operandB, err = strconv.Atoi(inputs[2])
	if err != nil {
		log.Fatal(err)
	}

	switch operation {
	case "+":
		result = operandA + operandB
	case "-":
		result = operandA - operandB
	case "*":
		result = operandA * operandB
	case "/":
		result = operandA / operandB
	case "**":
		res := math.Pow(float64(operandA), float64(operandB))
		result = int(res)
	default:
		log.Fatal("Invalid operation.")
	}

	fmt.Println(result)

}
