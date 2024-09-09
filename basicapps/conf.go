package main

import (
	"errors"
	"fmt"
	"os"
)

// Config contains the main app preferences
type config struct {
	// FilePath contains the absolute or relative path to a file for grep
	filePath string
	// KeyString is a substring we are looking for
	keyString string
	// IgnoreCase define the behavior about char's register: if ignoreCase is "true", we should ignore the register of the input
	ignoreCase string
}

// mustGetConfig do read config env variables and validation of the app input
// usually in go if function has a prefix "must/Must" - it means if will panic in case of any issue,
// like here.
func mustGetConfig() config {
	// TODO: add test
	config := loadConfig()
	if err := config.validate(); err != nil {
		panic(err)
	}

	return config
}

func loadConfig() config {
	// TODO: add test
	// var reader = bufio.NewReader(os.Stdin)
	// message, _ := reader.ReadString('\n')

	if len(os.Args[:]) > 1 {
		os.Setenv("KEY_STRING", os.Args[1])
		// TODO: logging
		fmt.Println("Warning: \"KEY_STRING\" environment variable got overwritten by command line argument")
	}
	return config{
		filePath:   os.Getenv("FILE_PATH"),
		keyString:  os.Getenv("KEY_STRING"),
		ignoreCase: os.Getenv("IGNORE_CASE"),
	}
}

func (c config) validate() error {
	// TODO: add test
	if c.filePath == "" {
		return errors.New("file path is empty")
	}

	if c.keyString == "" {
		return errors.New("nothing to search")
	}

	return nil
}
