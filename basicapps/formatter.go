package main

import (
	"fmt"
	"log"
	"regexp"
)

const (
	reset = "\033[0m"
	red   = "\033[31m"
)

type (
	formatter func(keyWord, line string, ignoreCase bool) string
	//formatter func(key, line string) string
)

func colorFormat(keyWord, line string, ignoreCase bool) string {
	if line == "" {
		return ""
	}
	var keyWord_re string
	if ignoreCase {
		keyWord_re = "(?i)" + keyWord
	} else {
		keyWord_re = keyWord
	}
	key, err := regexp.Compile(keyWord_re)
	if err != nil {
		log.Fatal(err)
	}

	replaced := key.ReplaceAllString(line, fmt.Sprintf("%s%s%s", red, key.FindString(line), reset)) //TODO: fix bug (line might contain more than one hits)
	return fmt.Sprintf("%s\n", replaced)
}
