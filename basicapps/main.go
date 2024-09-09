package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// The agenda:
// 1. App structure
//		1.1 Go don't go with a packages in most cases (it's not Java or C++)
//		1.2 Different files working very well
//		1.3 Security and immutability?
// 2. Code style
//		2.1 Functions over all (Max 5 input args, understandable name, simplicity, single responsibility (it's very hard)) etc
//		2.2 Responsibility
// 		2.3	Create files, not packages
//		2.4 Don't use "else" and "must" functions
//		2.5 Always to check error and think about the future (but not too much)
//		2.6
// 3. Dependencies
//		3.1 Inject them all
//		3.2 Use constructors (what is it?)
//		3.3 99% of your code must have an input and output - for testing
//		3.4 Dependencies must be substitutable
// 4. Good practices
//		4.1 Don't use pointers if you are not 100% sure app really need that (or by performance reasons)
//		4.2 Write tests
//		4.3 Write comments where it is necessary
//		4.4 No nested ifelse trash
// 5. Make own grep

func main() {
	config := mustGetConfig()

	file, err := os.Open(config.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err().Error())
	}
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		ignoreCase, err := strconv.ParseBool(config.ignoreCase)
		if err != nil {
			log.Fatalf("%v\nFailed to interprete command line argument.", err)
		}

		detected, err := checkLine(line, config.keyString, colorFormat, ignoreCase)
		if err != nil {
			log.Fatal(err)
		}
		if detected != "" {
			fmt.Print(detected)
		}
	}
}
