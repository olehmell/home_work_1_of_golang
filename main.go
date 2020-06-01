package main

import (
	"fmt"
	"os"
)

func main() {
	str := new(cryptoString)
	key := new(string)

	for {
		fmt.Println("Enter method:\n 1.Encode string\n 2.Decode string\n 3.Create new word on pig lang\n 4.Decode pig`s string\n 5.Exit")
		fmt.Scanln(key)
		switch *key {
		case "1":
			str.Encode(encodeToNum)
		case "2":
			str.Decode(decodeFromNum)
		case "3":
			str.Encode(encodeToPigLng)
		case "4":
			str.Decode(decodeFromPigLng)
		case "q":
			os.Exit(0)
		}

		str.Show()
	}
}
