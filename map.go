package main

var encryptingMap = map[rune]rune{
	'a': '1',
	'e': '2',
	'i': '3',
	'o': '4',
	'u': '5'}

var decryptingMap = map[rune]rune{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u'}

var pigLangSuffix = "ay"

var delimiters = ".,-;:?!"
