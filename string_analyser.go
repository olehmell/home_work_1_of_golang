package main

import "strings"

func analyseString(str string, method func(str string) (string, error)) (res string, err error) {
	words := strings.Split(str, " ")
	for i, v := range words {
		lastIndex := len(v) - 1
		lastChar := rune(v[lastIndex])
		if strings.ContainsRune(delimiters, lastChar) {
			res, _err := method(v[:lastIndex])
			words[i] = res + string(lastChar)
			err = _err
		} else {
			res, _err := method(v)
			words[i] = res
			err = _err
		}
	}
	return strings.Join(words, " "), err
}
