package main

import (
	"strings"
)

func encodeToNum(str string) (res string) {
	lowerStr := strings.ToLower(str)
	for i, letter := range lowerStr {
		if num, ok := encryptingMap[letter]; ok {
			res += string(num)
		} else {
			res += string(str[i])
		}
	}
	return
}

func convertWordToPigLng(str string) (res string, err error) {
	pigSuffix := string(str[0]) + pigLangSuffix
	otherLetters := str[1:]
	return otherLetters + pigSuffix, err
}

func encodeToPigLng(str string) string {
	res, _ := analyseString(str, convertWordToPigLng)
	return res
}
