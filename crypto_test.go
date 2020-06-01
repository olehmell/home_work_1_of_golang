package main

import (
	"testing"
)

func TestKoans(t *testing.T) {
	testStr := new(cryptoString)
	testStr.Write("My name Oleh, and I lear Go in Yalantis School! Are you too?")
	testStr.Show()
	testStr.EncodeToNumber()
	testStr.Show("You encoded vowels to numbers:")
	testStr.EncodeToPigLng()
	testStr.Show("You encoded string to pig language:")
	testStr.DecodeFromNumber()
	testStr.Show("You decoded string from pig language")
	testStr.DecodeFromPigLng()
	testStr.Show("You decoded vowels from numbers")
}
