package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func checkLine(line string, keyWord string, formatter formatter, ignoreCase bool) (string, error) {
	var keyWord_re string
	if ignoreCase {
		keyWord_re = `(?i)` + keyWord
	} else {
		keyWord_re = keyWord
	}
	key, err := regexp.Compile(keyWord_re)
	if err != nil {
		log.Fatal(err)
	}

	res, err := containsCheckRegexp(key, line)
	if err != nil {
		return "", fmt.Errorf("error while check line: %w: %s", err, line)
	}
	return formatter(keyWord, res, ignoreCase), nil
}

func containsCheckRegexp(key *regexp.Regexp, line string) (string, error) {
	// TODO: add test
	for _, word := range strings.Fields(strings.TrimSpace(line)) {
		found := key.MatchString(word) //bool
		if found {
			return line, nil
			break
		}
	}
	return "", nil
}
