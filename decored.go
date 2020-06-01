package main

import (
	"errors"
)

func decodeFromNum(str string) (res string, err error) {
	for _, letter := range str {
		if vowel, ok := decryptingMap[letter]; ok {
			res += string(vowel)
		} else {
			res += string(letter)
		}
	}
	return
}

func convertWordFromPigLng(str string) (res string, err error) {
	startPigSuffix := len(str) - 3
	pigSuffix := str[startPigSuffix:]
	if pigSuffix[1:] == pigLangSuffix {
		otherLetters := str[:startPigSuffix]
		firstLetter := string(pigSuffix[0])
		res = firstLetter + otherLetters
	} else {
		err = errors.New("This string isn't pig language")
	}
	return
}

func decodeFromPigLng(str string) (string, error) {
	return analyseString(str, convertWordFromPigLng)
}
