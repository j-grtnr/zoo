package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type (
	checkFunc func(keyWord, line string) (string, error)
)

func checkLine(line string, keyWord string, checkFunc checkFunc, formatter formatter, ignoreCase bool) (string, error) {
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

	res, err := checkFunc(key, line)
	if err != nil {
		return "", fmt.Errorf("error while check line: %w: %s", err, line)
	}
	return formatter(keyWord, res, ignoreCase), nil //TODO: fix bug - formatter fails if IGNORE_CASE="true"
	//return formatter(key, res), nil
}

func containsCheck(keyWord, line string) (string, error) {
	// TODO: add test
	if strings.Contains(line, keyWord) {
		return line, nil
	}

	// usually "else" keyword of what we don't really need
	return "", nil
}

func containsCheckIgnoreCase(keyWord, line string) (string, error) {
	// TODO: add test
	if strings.Contains(strings.ToLower(line), strings.ToLower(keyWord)) {
		return line, nil
	}
	return "", nil
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
